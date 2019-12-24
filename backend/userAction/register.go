package userAction

import (
	"backend/dataBase"
	"backend/helper/httpHelper"
	"backend/resourceScheduler/schs"
	"backend/userAction/querys"
	"context"
	"encoding/json"
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
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	schs.Scheduler.Request(ctx, func() {
		if exist, err := querys.ExistAccountTitle(ctx, regData.UserName); exist || err != nil {
			if err != nil {
				w.WriteHeader(500)
				return
			}
			w.WriteHeader(403)
			if _, e := w.Write([]byte("Already exist")); e != nil {
				w.WriteHeader(500)
			}
			return
		}
		if newUserID, err := querys.CreateAccount(ctx, regData.UserName, regData.Password); err != nil {
			w.WriteHeader(500)
			return
		} else {
			if data, e := json.Marshal(registerResult{newUserID}); e != nil {
				w.WriteHeader(500)
			} else if _, e = w.Write(data); e != nil {
				w.WriteHeader(500)
			}
			return
		}
	},dataBase.Accounts,dataBase.Settings)
}

type registerData struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type registerResult struct {
	UserID string `json:"userID"`
}
