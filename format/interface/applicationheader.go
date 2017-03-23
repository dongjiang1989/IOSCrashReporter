package Interface

import (
	"common/util"
	"errors"
	"fmt"
	"strings"

	log "github.com/cihub/seelog"
)

//Incident Identifier: 15BF6315-3B8D-48DD-A671-FCCF16581847
//CrashReporter Key:   A4E559D3-5F14-4EEB-A802-122D3FFCE1CF
//Hardware Model:      iPhone8,1
//Version:         2.7.3 (2.7.3.1612211140)
//Code Type:       ARM-64
//Parent Process:  ??? [1]

func NewInitApplicationHeader(str string) (*ApplicationHeader, error) {
	strs := Utils.FilterArray(strings.Split(strings.Trim(str, " "), "\n"))
	if len(strs) != 9 {
		log.Error(str)
		return nil, errors.New("strings split err! str:" + str)
	} else {
		ah := &ApplicationHeader{}
		for _, item := range strs {
			infoItem := strings.SplitN(item, ":", 2)
			if len(infoItem) != 2 {
				log.Error("strings.SplitN error! str:" + item)
				continue
			}
			key, value := strings.Trim(infoItem[0], " "), strings.Trim(infoItem[1], " ")
			switch key {
			case "Incident Identifier":
				ah.IncidentIdentifier = value
			case "CrashReporter Key":
				ah.CrashReporterKey = value
			case "Hardware Model":
				ah.HardwareModel = value
			case "Process":
				ah.Process = value
			case "Path":
				ah.Path = value
			case "Identifier":
				ah.Identifier = value
			case "Version":
				ah.Version = value
			case "Code Type":
				ah.CodeType = value
			case "Parent Process":
				ah.ParentProcess = value
			default:
				continue
			}
		}

		return ah, nil
	}

	return nil, errors.New("no ApplicationHeader Object!")
}

type ApplicationHeader struct {
	IncidentIdentifier string
	CrashReporterKey   string
	HardwareModel      string
	Process            string
	Path               string
	Identifier         string
	Version            string
	CodeType           string
	ParentProcess      string
}

func (ah *ApplicationHeader) String() string {
	ret := ""
	ret = ret + fmt.Sprintf("Incident Identifier: %s\n", ah.IncidentIdentifier)
	ret = ret + fmt.Sprintf("CrashReporter Key:   %s\n", ah.CrashReporterKey)
	ret = ret + fmt.Sprintf("Hardware Model:      %s\n", ah.HardwareModel)
	ret = ret + fmt.Sprintf("Process:         %s\n", ah.Process)
	ret = ret + fmt.Sprintf("Path:            %s\n", ah.Path)
	ret = ret + fmt.Sprintf("Identifier:      %s\n", ah.Identifier)
	ret = ret + fmt.Sprintf("Version:         %s\n", ah.Version)
	ret = ret + fmt.Sprintf("Code Type:       %s\n", ah.CodeType)
	ret = ret + fmt.Sprintf("Parent Process:  %s\n", ah.ParentProcess)
	return ret
}
