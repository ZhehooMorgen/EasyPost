package resourceScheduler

import (
	"backend/util"
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	Name    string
	Balance int
	ID      uuid.UUID
}

func (b *BankAccount) Definition() (Type, uuid.UUID) {
	return "back account", b.ID
}

func Pay(bank Scheduler, from *BankAccount, to *BankAccount, amount int, duration time.Duration) util.Err {
	return bank.Request(context.Background(), func() {
		fmt.Println("from ", from.Name, " to ", to.Name, " start")
		if &from.Balance != &to.Balance {
			f := from.Balance - amount
			t := to.Balance + amount
			time.Sleep(duration * time.Millisecond)
			from.Balance = f
			to.Balance = t
		} else {
			b := from.Balance
			time.Sleep(duration * time.Millisecond)
			to.Balance = b
		}
		fmt.Println("from ", from.Name, " to ", to.Name, " finish")
	}, from, to)
}

func TestBasicScheduler_RegRes(t *testing.T) {
	if (&BasicScheduler{}).RegRes(nil) == nil {
		t.Fatal()
	}
	var scheduler Scheduler = NewBasicScheduler()
	if scheduler.RegRes(&BankAccount{
		Name:    "Bob",
		Balance: 0,
		ID:      uuid.UUID{},
	}) == nil {
		t.Fatal()
	}
	if scheduler.RegRes(&BankAccount{
		Name:    "Bob",
		Balance: 0,
		ID:      uuid.NewV4(),
	}) != nil {
		t.Fatal()
	}

}

func TestBasicScheduler_Request(t *testing.T) {
	var scheduler Scheduler = NewBasicScheduler()
	Bob := &BankAccount{
		Name:    "Bob",
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	Bill := &BankAccount{
		Name:    "Bill",
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	Error := &BankAccount{
		Name:    "Bob",
		Balance: 0,
		ID:      uuid.Nil,
	}
	if scheduler.Request(context.Background(), func() {}, Bob, Bill) == nil {
		t.Fatal()
	}
	_ = scheduler.RegRes(Bob)
	_ = scheduler.RegRes(Bill)
	if scheduler.Request(context.Background(), func() {}, Bob, Bill) != nil {
		t.Fatal()
	}
	if scheduler.Request(context.Background(), func() {}, Bob, Error) == nil {
		t.Fatal()
	}
}

func TestBasicScheduler(t *testing.T) {
	var bank Scheduler = NewBasicScheduler()
	Bob := &BankAccount{
		Name:    "Bob",
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	Bill := &BankAccount{
		Name:    "Bill",
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	Simon := &BankAccount{
		Name:    "Simon",
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	Dave := &BankAccount{
		Name:    "Dave",
		Balance: 0,
		ID:      uuid.NewV4(),
	}
	_ = bank.RegRes(Bob)
	_ = bank.RegRes(Bill)
	_ = bank.RegRes(Simon)
	_ = bank.RegRes(Dave)
	var testCase = []struct {
		From     *BankAccount
		To       *BankAccount
		Amount   int
		Duration time.Duration
	}{
		{Bob, Bill, 2, 110},
		{Simon, Dave, 43, 154},
		{Dave, Bob, 564, 164},
		{Bill, Simon, 28, 104},
		{Bill, Bill, 83, 139},
		{Bob, Bob, 53, 195},
		{Simon, Simon, 47, 103},
		{Dave, Dave, 9, 100},
	}
	wait := sync.WaitGroup{}
	wait.Add(7)
	for _, item := range testCase {
		var local = item
		go func() {
			Pay(bank, local.From, local.To, local.Amount, local.Duration)
			wait.Done()
		}()
	}
	go func() {
		ctx, _ := context.WithTimeout(context.Background(), 1)
		bank.Request(ctx, func() {}, Simon, Bob, Bill, Dave)
	}()
	wait.Wait()
	if Bob.Balance+Bill.Balance+Simon.Balance+Dave.Balance != 0 {
		t.Fatal(Bob.Balance + Bill.Balance + Simon.Balance + Dave.Balance)
	}
}
