package Interface

import (
	"common/util"
	"errors"
	"fmt"
	"strings"
	"time"

	log "github.com/cihub/seelog"
)

//Date/Time:       2017-01-06T09:52:42Z
//Launch Time:     2017-01-06T09:52:33Z
//OS Version:      iPhone OS 9.0.2 (13A452)
//Report Version:  104

func NewInitSystemInfo(str string) (*SystemInfo, error) {
	strs := Utils.FilterArray(strings.Split(strings.Trim(str, " "), "\n"))
	if len(strs) != 4 {
		log.Error(str)
		return nil, errors.New("strings split err! str:" + str)
	} else {
		si := &SystemInfo{}
		for _, item := range strs {
			infoItem := strings.SplitN(item, ":", 2)
			if len(infoItem) != 2 {
				log.Error("strings.SplitN error! str:" + item)
				continue
			}
			key, value := strings.Trim(infoItem[0], " "), strings.Trim(infoItem[1], " ")
			switch key {
			case "Date/Time":
				loc, _ := time.LoadLocation("Local")
				si.DateTime, _ = time.ParseInLocation("2006-01-02T15:04:05Z", value, loc)
			case "Launch Time":
				loc, _ := time.LoadLocation("Local")
				si.LaunchTime, _ = time.ParseInLocation("2006-01-02T15:04:05Z", value, loc)
			case "OS Version":
				si.OSVersion = value
			case "Report Version":
				si.ReportVersion = value
			default:
				continue
			}
		}

		return si, nil
	}

	return nil, errors.New("no SystemInfo Object!")
}

type SystemInfo struct {
	DateTime      time.Time
	LaunchTime    time.Time
	OSVersion     string
	ReportVersion string
}

func (si *SystemInfo) String() string {
	ret := ""
	ret = ret + fmt.Sprintf("Date/Time:       %s\n", si.DateTime.Format("2006-01-02T15:04:05Z"))
	ret = ret + fmt.Sprintf("Launch Time:     %s\n", si.LaunchTime.Format("2006-01-02T15:04:05Z"))
	ret = ret + fmt.Sprintf("OS Version:      %s\n", si.OSVersion)
	ret = ret + fmt.Sprintf("Report Version:  %s\n", si.ReportVersion)
	return ret
}
