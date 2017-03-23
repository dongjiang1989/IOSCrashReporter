package Format

import (
	"common/util"
	"encoding/json"
	"fmt"

	log "github.com/cihub/seelog"
)

//type ITFHttpDoInfo interface {
//	GetModuleName(string, string) string // hit count
//	GetParseSymbol(interface{}, int, string) ([]byte, error)
//}

//// HttpDo
//type HttpDoInfo struct {
//}

// get module name
func GetModuleName(uuid string, classname string) string {
	head := make(map[string]string)
	params := make(map[string]string)
	params["uuid"] = uuid
	params["classname"] = classname

	returnByte, err := Utils.HttpJsonDo("http://crashapi.xxxxxxxxx.com/crash/api/getmodulename", head, params)
	if err != nil {
		log.Error("Crash_Analysis : ", uuid, ", err :", err.Error())
		return ""
	}

	ret := make(map[string]interface{}, 0)
	err = json.Unmarshal(returnByte, &ret)
	if err != nil {
		log.Error("Crash_Analysis : json Unmarshal err! err :%s", uuid, err.Error())
		return ""
	}

	retStr, ok := ret["data"].(string)
	if ok {
		return retStr
	}
	return ""
}

// get parse symbol
func GetParseSymbol(uuidMap interface{}, app_id int, msgid string) ([]byte, error) {
	params := make(map[string]string)
	offsetBytes, _ := json.Marshal(uuidMap)
	params["offsets"] = string(offsetBytes)
	params["app_id"] = fmt.Sprintf("%d", app_id)

	head := make(map[string]string)
	head["Content-Type"] = "application/json"
	return Utils.HttpJsonDo("http://crashapi.xxxxxxxxx.com/crash/api/parselatestsymbol", head, params)
}
