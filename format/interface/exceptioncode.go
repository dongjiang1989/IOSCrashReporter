package Interface

import (
	"common/util"
	"errors"
	"fmt"
	"strings"

	log "github.com/cihub/seelog"
)

//Exception Type:  SIGABRT
//Exception Codes: #0 at 0x1983131e0
//Crashed Thread:  33
func NewInitExceptionCode(str string) (*ExceptionCode, error) {
	strs := Utils.FilterArray(strings.Split(strings.Trim(str, " "), "\n"))
	if len(strs) != 3 {
		log.Error(str)
		return nil, errors.New("strings split err! str:" + str)
	} else {
		ah := &ExceptionCode{}
		for _, item := range strs {
			infoItem := strings.SplitN(item, ":", 2)
			if len(infoItem) != 2 {
				log.Error("strings.SplitN error! str:" + item)
				continue
			}

			key, value := strings.Trim(infoItem[0], " "), strings.Trim(infoItem[1], " ")
			switch key {
			case "Exception Type":
				ah.ExceptionType = value
			case "Exception Codes":
				ah.ExceptionCodes = value
			case "Crashed Thread":
				ah.CrashedThread = value
			default:
				continue
			}
		}

		return ah, nil
	}

	return nil, errors.New("no ExceptionCode Object!")
}

type ExceptionCode struct {
	ExceptionType  string
	ExceptionCodes string
	CrashedThread  string
}

func (e *ExceptionCode) String() string {
	ret := ""
	ret = ret + fmt.Sprintf("Exception Type:  %s\n", e.ExceptionType)
	ret = ret + fmt.Sprintf("Exception Codes: %s\n", e.ExceptionCodes)
	ret = ret + fmt.Sprintf("Crashed Thread:  %s\n", e.CrashedThread)
	return ret
}

//TODO maybe
func (e *ExceptionCode) ToErrTagTitle() string {
	ret := ""
	ret = ret + fmt.Sprintf("Exception Type:  %s Exception Codes: %s", e.ExceptionType, e.ExceptionCodes)
	return ret
}

func (e *ExceptionCode) ToErrTagTitleString(offset string) string {
	ret := ""
	ret = ret + fmt.Sprintf("Exception Type:  %s Exception Codes: %s Offset: %s", e.ExceptionType, e.ExceptionCodes, offset)
	return ret
}

func (e *ExceptionCode) ToErrTagType() string {
	return fmt.Sprintf("Exception Type:  %s", e.ExceptionType)
}
