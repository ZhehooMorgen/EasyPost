package blackBox

import (
	"backend/helper/httpHelper"
	"testing"
)

func TestGen_ReqDataParseFail(t *testing.T){
	if httpHelper.ReqDataParseFail(true,nil).ErrorCode()!=500{
		t.Fatal("error code gen not correct")
	}
	if httpHelper.ReqDataParseFail(false,nil).ErrorCode()!=400{
		t.Fatal("error code gen not correct")
	}
}

func TestGen_HTTPMethodWrong(t *testing.T){
	if httpHelper.HTTPMethodWrong(nil).ErrorCode()!=405{
		t.Fatal("error code gen not correct")
	}
}

func TestGen_HTTPWriteFail(t *testing.T){
	if httpHelper.HTTPWriteFail(nil).ErrorCode()!=20{
		t.Fatal("error code gen not correct")
	}
}

func TestGenHTTPReadFail(t *testing.T){
	if httpHelper.HTTPReadFail(nil).ErrorCode()!=500{
		t.Fatal("error code gen not correct")
	}
}