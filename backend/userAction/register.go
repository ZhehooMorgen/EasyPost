package userAction

import (
	"backend/helper/httpHelper"
	"backend/userAction/querys"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

func registerService(w http.ResponseWriter, req *http.Request) {
	var ctx, cancelCtx = context.WithTimeout(context.Background(), time.Second*3)
	defer cancelCtx()
	var regData registerData
	logInfo, _, err := httpHelper.ReqPreProcess(http.MethodPost, w, req, true, &regData)
	fmt.Println(logInfo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	if ok, err := querys.ExistAccountTitle(ctx, regData.UserName); !ok || err != nil {
		w.WriteHeader(403)
		if _, e := w.Write([]byte("Already exist")); e != nil {
			w.WriteHeader(500)
		}
		return
	}

}

type registerData struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type registerResult struct {
	UserID string `json:"userID"`
}
