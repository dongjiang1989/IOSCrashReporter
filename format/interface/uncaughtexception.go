package Interface

import (
	"common/util"
	"errors"
	"strings"

	log "github.com/cihub/seelog"
)

//strA:
//Application Specific Information:
//*** Terminating app due to uncaught exception 'UninitializedMessage', reason: ''

//strB:
//Last Exception Backtrace:
//0   CoreFoundation                       0x00000001824e8f5c 0x1823c4000 + ___exceptionPreprocess
//1   libobjc.A.dylib                      0x00000001979d7f80 0x1979d0000 + _objc_exception_throw
//2   DiSpecialDriver                      0x000000010099c2a0 0x100034000 + -[PBGeneratedMessageBuilder checkInitialized]	GeneratedMessageBuilder.m:46
//3   DiSpecialDriver                      0x000000010031306c 0x100034000 + -[CdntSvrDownReqBuilder build]	Didi_protocol.pb.m:20660
//4   DiSpecialDriver                      0x0000000100312260 0x100034000 + +[CdntSvrDownReq parseFromData:]	Didi_protocol.pb.m:20566
//5   DiSpecialDriver                      0x00000001000fcfc8 0x100034000 + -[MsgTypeKMsgTypeCdntSvrDownReqRecver recv:data:]	MsgTypeKMsgTypeCdntSvrDownReqRecver.m:17
//6   DiSpecialDriver                      0x00000001005862ec 0x100034000 + _Z12RecvCallBackjyPKcj	DSLongLink.mm:57
//7   DiSpecialDriver                      0x0000000100d4db44 0x100034000 + _ZN8DidiPush28ConnConnectionHandleRecvDataEPNS_17tagConnConnectionEPKvj	conn_connection.cpp:1392
//8   DiSpecialDriver                      0x0000000100d4d864 0x100034000 + _ZN8DidiPush18ConnConnectionReadEisPNS_17tagConnConnectionE	conn_connection.cpp:1249
//9   DiSpecialDriver                      0x0000000100d54c44 0x100034000 + _event_base_loop
//10  DiSpecialDriver                      0x0000000100d00468 0x100034000 + _ZN8DidiPush9StartLoopEPNS_10tagPushSdkE	pushsdk.cpp:380
//11  Foundation                           0x0000000183427138 0x183334000 + ___NSThread__start__
//12  libsystem_pthread.dylib              0x00000001983dbb3c 0x1983d8000 + __pthread_body
//13  libsystem_pthread.dylib              0x00000001983dbaa0 0x1983d8000 + __pthread_start
//14  libsystem_pthread.dylib              0x00000001983d9030 0x1983d8000 + _thread_start
func NewInitUncaughtException(str string) (*UncaughtException, error) {
	strs := Utils.FilterArray(strings.Split(strings.Trim(str, " "), "\n"))
	ue := &UncaughtException{}
	if len(strs) < 2 {
		ue.ApplicationSpecificInformation = ""
	} else {
		if strings.Contains(strs[0], "Application Specific Information:") {
			ue.ApplicationSpecificInformation = strs[1]
		} else {
			log.Error("Can not find 'Application Specific Information:'!  str:" + strs[0])
			return nil, errors.New("Can not find 'Application Specific Information:'!  str:" + strs[0])
		}
	}
	return ue, nil
}

type UncaughtException struct {
	ApplicationSpecificInformation string //Application Specific Information:
}

func NewInitLastExceptionBacktrace(str string) (*LastExceptionBacktrace, error) {
	strs := Utils.FilterArray(strings.Split(strings.Trim(str, " "), "\n"))
	lb := &LastExceptionBacktrace{}

	if len(strs) < 2 {
		lb.LastExceptionBacktrace = nil
	} else if strings.Contains(strs[0], "Last Exception Backtrace:") {
		for _, item := range strs[1:] {
			l, err := NewInitLine(item)
			if err != nil {
				continue
			}
			lb.LastExceptionBacktrace = append(lb.LastExceptionBacktrace, l)
		}
	}

	return lb, nil
}

type LastExceptionBacktrace struct {
	LastExceptionBacktrace []*Line //Last Exception Backtrace:
}

func (ue *LastExceptionBacktrace) ModuleInfo() map[string][]int64 {
	ret := make(map[string][]int64)
	for _, item := range ue.LastExceptionBacktrace {
		dict, ok := ret[item.ImageName]
		if ok {
			dict = append(dict, item.PcOffset)
		} else {
			ret[item.ImageName] = make([]int64, 0)
			ret[item.ImageName] = append(ret[item.ImageName], item.PcOffset)
		}
	}
	return ret
}

func (ue *UncaughtException) String() string {
	ret := ""
	if ue.ApplicationSpecificInformation != "" {
		ret = ret + "Application Specific Information:\n"
		ret = ret + ue.ApplicationSpecificInformation + "\n"
	}
	return ret
}

func (lb *LastExceptionBacktrace) String(isParse bool) string {
	ret := ""
	if lb.LastExceptionBacktrace != nil {
		ret = ret + "Last Exception Backtrace:\n"
		for _, item := range lb.LastExceptionBacktrace {
			ret = ret + item.String(isParse)
		}
	}
	return ret
}
