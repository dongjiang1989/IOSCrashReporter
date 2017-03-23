package Format

import (
	"common/util"
	"encoding/json"
	"errors"
	"format/interface"
	"strconv"
	"strings"

	log "github.com/cihub/seelog"
)

func NewInitCrashReport(str string) (*CrashReport, error) {
	strs := Utils.FilterArray(strings.Split(strings.Trim(str, " "), "\n\n"))
	//TODO: maybe bug
	cr := &CrashReport{Application: nil, System: nil, ExceptionCode: nil, UncaughtException: nil, LastExceptionBacktrace: nil, Registers: nil, BinaryImages: nil}
	cr.Threads = &Interface.ThreadsInfos{}
	if len(strs) >= 7 {
		for _, item := range strs {
			if strings.Contains(item, "Incident Identifier:") {
				ah, err := Interface.NewInitApplicationHeader(item)
				if err != nil {
					log.Error(err)
					return nil, err
				}
				cr.Application = ah
			}

			if strings.Contains(item, "Date/Time:  ") {
				si, err := Interface.NewInitSystemInfo(item)
				if err != nil {
					log.Error(err)
					return nil, err
				}
				cr.System = si
			}

			if strings.Contains(item, "Exception Codes: ") {
				ec, err := Interface.NewInitExceptionCode(item)
				if err != nil {
					log.Error(err)
					return nil, err
				}
				cr.ExceptionCode = ec
			}

			if strings.Contains(item, "Application Specific Information:") {
				ue, err := Interface.NewInitUncaughtException(item)
				if err != nil {
					log.Error(err)
					return nil, err
				}
				cr.UncaughtException = ue
			}

			if strings.Contains(item, "Last Exception Backtrace:") {
				bl, err := Interface.NewInitLastExceptionBacktrace(item)
				if err != nil {
					log.Error(err)
					return nil, err
				}
				cr.LastExceptionBacktrace = bl
			}

			if strings.Contains(item, " Thread State:") {
				ri, err := Interface.NewInitRegistersInfo(item, LP64(cr.Application.CodeType))
				if err != nil {
					log.Error(err)
					return nil, err
				}
				cr.Registers = ri
			} else {
				if strings.Contains(item, "Thread ") { //TODO
					ti, err := Interface.NewInitThreadsInfo(item)
					if err != nil {
						log.Error(err)
						continue
					}
					cr.Threads.Threads = append(cr.Threads.Threads, ti)
				}
			}

			if strings.Contains(item, "Binary Images:") {

				bi, err := Interface.NewInitBinaryImages(item, LP64(cr.Application.CodeType))
				if err != nil {
					log.Error(err)
					return nil, err
				}
				cr.BinaryImages = bi
			}
		}
	} else {
		log.Error("str len is :", len(strs), " strs:", strs)
		return nil, errors.New("str len is :" + strconv.Itoa(len(strs)))
	}
	return cr, nil
}

func LP64(str string) bool {
	lp64 := true

	if str == "ARM" || str == "X86" || str == "PPC" {
		lp64 = false
	}

	return lp64
}

type CrashReport struct {
	IsParsed    bool
	IsAllParsed bool

	Application            *Interface.ApplicationHeader
	System                 *Interface.SystemInfo
	ExceptionCode          *Interface.ExceptionCode
	UncaughtException      *Interface.UncaughtException
	LastExceptionBacktrace *Interface.LastExceptionBacktrace
	Threads                *Interface.ThreadsInfos
	Registers              *Interface.RegistersInfo
	BinaryImages           *Interface.BinaryImages
}

func (cr *CrashReport) String(isParse bool) string {
	ret := ""
	if cr.Application != nil {
		ret = ret + cr.Application.String()
		ret = ret + "\n"
	}

	if cr.System != nil {
		ret = ret + cr.System.String()
		ret = ret + "\n"
	}

	if cr.ExceptionCode != nil {
		ret = ret + cr.ExceptionCode.String()
		ret = ret + "\n"
	}

	if cr.UncaughtException != nil {
		ret = ret + cr.UncaughtException.String()
		ret = ret + "\n"
	}

	if cr.LastExceptionBacktrace != nil {
		ret = ret + cr.LastExceptionBacktrace.String(isParse)
		ret = ret + "\n"
	}

	for _, thread := range cr.Threads.Threads {
		ret = ret + thread.String(isParse)
		ret = ret + "\n"
	}

	if cr.Registers != nil {
		ret = ret + cr.Registers.String()
		ret = ret + "\n"
	}

	if cr.BinaryImages != nil {
		ret = ret + cr.BinaryImages.String()
	}

	return ret
}

//imageName : [value, value, ...]
func (cr *CrashReport) GetUUids() map[string]string {
	return cr.BinaryImages.Uuids()
}

func (cr *CrashReport) GetModuleName() string {
	return cr.BinaryImages.ModuleName()
}

//uuid : [value, value, ...]
func (cr *CrashReport) UuidMaps() map[string][]int64 {
	ret := make(map[string][]int64, 0)
	if cr.LastExceptionBacktrace != nil {
		for _, item := range cr.LastExceptionBacktrace.LastExceptionBacktrace {
			_, ok := ret[item.ImageName]
			if !ok {
				ret[item.ImageName] = make([]int64, 0)
			}
			ret[item.ImageName] = append(ret[item.ImageName], item.PcOffset)

		}
	}

	if cr.Threads != nil {
		for _, thread := range cr.Threads.Threads {
			for _, item := range thread.Lines {
				_, ok := ret[item.ImageName]
				if !ok {
					ret[item.ImageName] = make([]int64, 0)
				}
				ret[item.ImageName] = append(ret[item.ImageName], item.PcOffset)
			}
		}
	}

	uuidMap := cr.BinaryImages.Uuids()
	retNew := make(map[string][]int64)

	for imageName, value := range ret {
		uuid, ok := uuidMap[imageName]
		if !ok {
			log.Warn("Can not find imageName:" + imageName + " in crash binary images!")
			continue
		}

		retNew[uuid] = Utils.RemoveDuplicates(value)
	}

	return retNew
}

type SymbolRespose struct {
	Hit   int64                        `json:"hit"`
	Total int64                        `json:"total"`
	Msg   string                       `json:"string"`
	Ret   int64                        `json:"ret"`
	Data  map[string]map[string]string `json:"data"`
}

func (cr *CrashReport) ParseToString(appid int, msgid string) error {
	retD, err := GetParseSymbol(cr.UuidMaps(), appid, msgid)
	if err != nil {
		return err
	}

	data := new(SymbolRespose)
	err = json.Unmarshal(retD, &data)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	//完整解析
	if data.Hit+1 >= data.Total {
		cr.IsAllParsed = true
		cr.IsParsed = true
	} else {
		cr.IsAllParsed = false
		infoUuids, ok := data.Data[cr.BinaryImages.ModuleUuid()]
		if ok && len(infoUuids) > 0 {
			cr.IsParsed = true
		} else {
			cr.IsParsed = false
		}
	}

	if cr.LastExceptionBacktrace != nil {
		for _, item := range cr.LastExceptionBacktrace.LastExceptionBacktrace {
			uuid, ok := cr.BinaryImages.Uuids()[item.ImageName]
			if !ok {
				continue
			} else {
				info, ok := data.Data[uuid]
				if !ok {
					continue
				} else {
					tmpdata, ok := info[strconv.FormatInt(item.PcOffset, 10)]
					if !ok {
						item.ParseStr = strconv.FormatInt(item.PcOffset, 10)
					} else {
						item.ParseStr = tmpdata

					}
				}
			}
		}
	}

	if cr.Threads != nil {
		for _, thread := range cr.Threads.Threads {
			for _, item := range thread.Lines {
				uuid, ok := cr.BinaryImages.Uuids()[item.ImageName]
				if !ok {
					continue
				} else {
					info, ok := data.Data[uuid]
					if !ok {
						continue
					} else {
						tmpdata, ok := info[strconv.FormatInt(item.PcOffset, 10)]
						if !ok {
							item.ParseStr = strconv.FormatInt(item.PcOffset, 10)
						} else {
							item.ParseStr = tmpdata
						}
					}
				}
			}
		}
	}

	return nil
}

func (cr *CrashReport) GetCrashStack() []*Interface.Line {
	if cr.LastExceptionBacktrace != nil {
		return cr.LastExceptionBacktrace.LastExceptionBacktrace
	} else {
		return cr.Threads.CrashStack()
	}
	return nil
}

func (cr *CrashReport) CorruptThread() string {
	ret := ""
	if cr.UncaughtException != nil {
		ret = ret + cr.UncaughtException.String()
		ret = ret + "\n"
	}

	if cr.LastExceptionBacktrace != nil {
		ret = ret + cr.LastExceptionBacktrace.String(true)
	} else {
		thread, _ := cr.Threads.CrashThread()
		ret = ret + thread.String(true)
	}

	return ret
}

func (cr *CrashReport) ToErrMsg() string {
	if cr.UncaughtException != nil {
		return cr.UncaughtException.ApplicationSpecificInformation
	}
	return ""
}
