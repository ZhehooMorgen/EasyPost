package resourceScheduler

import (
	"backend/util"
	"context"
	uuid "github.com/satori/go.uuid"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	Name string
	Balance int
	ID uuid.UUID
}

func (b *BankAccount)Definition()(Type,uuid.UUID){
	return "back account",b.ID
}

func Pay(bank Scheduler,from *BankAccount, to *BankAccount, amount int,duration time.Duration)util.Err{
	 return bank.Request(context.Background(), func() {
		from.Balance-=amount
		time.Sleep(duration*time.Millisecond)
		to.Balance+=amount
	},from,to)
}

func TestBasicScheduler_RegRes(t *testing.T) {
	if ((&BasicScheduler{}).RegRes(nil)==nil){
		t.Fatal()
	}
	var scheduler Scheduler = NewBasicScheduler()
	if scheduler.RegRes(&BankAccount{
		Name:    "Bob",
		Balance: 0,
		ID:      uuid.UUID{},
	})==nil{
		t.Fatal()
	}
	if scheduler.RegRes(&BankAccount{
		Name:    "Bob",
		Balance: 0,
		ID:      uuid.NewV4(),
	})!=nil{
		t.Fatal()
	}

}

func TestBasicScheduler_Request(t *testing.T) {
	var scheduler Scheduler = NewBasicScheduler()
	Bob:=&BankAccount{
		Name:    "Bob",
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	Bill:=&BankAccount{
		Name:    "Bill",
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	Error:=&BankAccount{
		Name:    "Bob",
		Balance: 0,
		ID:      uuid.Nil,
	}
	if scheduler.Request(context.Background(), func() {},Bob,Bill)==nil{
		t.Fatal()
	}
	_=scheduler.RegRes(Bob)
	_=scheduler.RegRes(Bill)
	if scheduler.Request(context.Background(), func() {},Bob,Bill)!=nil{
		t.Fatal()
	}
	if scheduler.Request(context.Background(), func() {},Bob,Error)==nil{
		t.Fatal()
	}
}

func TestBasicScheduler(t *testing.T) {
	var bank Scheduler = NewBasicScheduler()
	Bob:=&BankAccount{
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	Bill:=&BankAccount{
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	Simon:=&BankAccount{
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	Dave:=&BankAccount{
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	_= bank.RegRes(Bob)
	_= bank.RegRes(Bill)
	_= bank.RegRes(Simon)
	_= bank.RegRes(Dave)
	wait :=sync.WaitGroup{}
	wait.Add(7)
	go func() {
		Pay(bank,Bob,Bill,2,110)
		wait.Done()
	}()
	go func() {
		Pay(bank,Simon,Dave,43,154)
		wait.Done()
	}()
	go func() {
		Pay(bank,Dave,Bob,564,164)
		wait.Done()
	}()
	go func() {
		Pay(bank,Bill,Simon,28,104)
		wait.Done()
	}()
	go func() {
		Pay(bank,Bill,Bill,83,139)
		wait.Done()
	}()
	go func() {
		Pay(bank,Bob,Bob,53,195)
		wait.Done()
	}()
	go func() {
		Pay(bank,Simon,Simon,47,103)
		wait.Done()
	}()
	go func() {
		ctx,_:=context.WithTimeout(context.Background(),0)
		bank.Request(ctx, func() {},Simon,Bob,Bill,Dave)
	}()
	wait.Wait()
	if Bob.Balance+Bill.Balance+Simon.Balance+Dave.Balance!=0{
		t.Fatal()
	}
}