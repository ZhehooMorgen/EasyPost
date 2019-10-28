package whiteBox_test

import (
	"backend/userAction/querys"
	"context"
	"testing"
)

func Test_ExistAccountTitle(t *testing.T){
	querys.InitMongoDBConnection()
	if ok,err:= querys.ExistAccountTitle(context.TODO(),"Root");!ok||err!=nil{
		t.Fatal("query result not right")
	}
}

func Test_CreateAccount(t *testing.T){
	querys.InitMongoDBConnection()
	if _,err:= querys.CreateAccount(context.TODO(),"test1","1");err!=nil{
		t.Fatal("query result not right")
	}
	if _,err:= querys.CreateAccount(context.TODO(),"test2","1");err!=nil{
		t.Fatal("query result not right")
	}
}