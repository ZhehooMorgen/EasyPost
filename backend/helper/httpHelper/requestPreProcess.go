package httpHelper

import (
	"backend/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReqPreProcess(method string, w http.ResponseWriter, req *http.Request, cors bool, data interface{}) (reqInfo string, rawData []byte, err util.Err) {
	reqInfo = fmt.Sprint("New request :", req.Method, req.Host, req.RequestURI)
	if cors {
		CORS(w)
	}
	if method != "" && req.Method != method {
		dataBytes, _ := ioutil.ReadAll(req.Body)
		return reqInfo, dataBytes, HTTPMethodWrong(nil)
	}
	dataBytes, e := ioutil.ReadAll(req.Body)
	if data != nil {
		if e != nil {
			return reqInfo, dataBytes, HTTPReadFail(e)
		}
		e = json.Unmarshal(dataBytes, data)
		if e != nil {
			if _, ok := e.(*json.InvalidUnmarshalError); ok {
				return reqInfo, dataBytes, ReqDataParseFail(true, e)
			} else {
				return reqInfo, dataBytes, ReqDataParseFail(false, e)
			}
		}
		return reqInfo, dataBytes, nil
	} else {
		return reqInfo, dataBytes, nil
	}

}
