package Interface

import (
	"common/util"
	"errors"
	"fmt"
	"strconv"
	"strings"

	log "github.com/cihub/seelog"
)

//Thread 0 Crashed:
//0   libsystem_kernel.dylib               0x19a08a87c 0x19a070000 + 108668
//1   libsystem_platform.dylib             0x19a14d94c 0x19a148000 + 22860
//2   KDaijiaDriver                        0x10090a920 0x10005c000 + 9103648
//3   KDaijiaDriver                        0x100904460 0x10005c000 + 9077856
//4   KDaijiaDriver                        0x100904368 0x10005c000 + 9077608
//5   OTRLocation.dylib                    0x103de1050 0x103ddc000 + 20560
//6   OTRLocation.dylib                    0x103de1050 0x103ddc000 + 20560
//7   GpsHookLibrary.dylib                 0x103dc5fe8 0x103dc4000 + 8168
//8   CoreLocation                         0x1853f0850 0x1853e4000 + 51280
//9   CoreLocation                         0x1853ecab0 0x1853e4000 + 35504
//10  CoreLocation                         0x1853e6e90 0x1853e4000 + 11920
//11  CoreFoundation                       0x184c1442c 0x184b38000 + 902188
//12  CoreFoundation                       0x184c13d64 0x184b38000 + 900452
//13  CoreFoundation                       0x184c121ac 0x184b38000 + 893356
//14  CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//15  GraphicsServices                     0x18fd7c088 0x18fd70000 + 49288
//16  UIKit                                0x18a258ffc 0x18a1d8000 + 528380
//17  KDaijiaDriver                        0x1002771f8 0x10005c000 + 2208248
//18  libdyld.dylib                        0x199f6e8b8 0x199f6c000 + 10424
func NewInitThreadsInfo(str string) (*ThreadsInfo, error) {
	strs := Utils.FilterArray(strings.Split(strings.Trim(str, " "), "\n"))
	if len(strs) < 2 {
		log.Error(str)
		return nil, errors.New("strings split err! len:" + strconv.Itoa(len(strs)) + " str:" + str)
	} else {
		ue := &ThreadsInfo{}

		//TODO: maybe bug
		if strings.Contains(strs[0], "Thread ") {
			if strings.Contains(strs[0], "Crashed:") {
				ue.IsCrashed = true
			} else {
				ue.IsCrashed = false
			}
			items := Utils.FilterArray(strings.Split(strs[0], " "))
			if len(items) == 3 {
				ue.ThreadNum, _ = Utils.StringToInt64(items[1])
			} else if len(items) == 2 {
				ue.ThreadNum, _ = Utils.StringToInt64(strings.Trim(items[1], ":"))
			} else {
				log.Error("Can not trans to ThreadNum! str:" + strs[0])
				return nil, errors.New("Can not trans to ThreadNum! str:" + strs[0])
			}
		} else {
			log.Error("Can not find 'Thread .* xxx:'!  str:" + strs[0])
			return nil, errors.New("Can not find 'Thread .* xxx:'!  str:" + strs[0])
		}

		for _, item := range strs[1:] {
			l, err := NewInitLine(item)
			if err != nil {
				continue
			}
			ue.Lines = append(ue.Lines, l)
		}

		return ue, nil
	}

	return nil, errors.New("no ThreadsInfo Object!")
}

type ThreadsInfo struct {
	IsCrashed bool
	ThreadNum int64
	Lines     []*Line
}

func (ti *ThreadsInfo) ModuleInfo() map[string][]int64 {
	ret := make(map[string][]int64)
	for _, item := range ti.Lines {
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

func (ti *ThreadsInfo) String(isParse bool) string {
	ret := ""
	if ti.IsCrashed {
		ret = ret + fmt.Sprintf("Thread %s Crashed:\n", Utils.Int64ToString(ti.ThreadNum, 10))
	} else {
		ret = ret + fmt.Sprintf("Thread %s:\n", Utils.Int64ToString(ti.ThreadNum, 10))
	}
	for _, item := range ti.Lines {
		ret = ret + item.String(isParse)
	}
	return ret
}

//--------------------------------------------------------------------

func NewInitThreadsInfos(str string) (*ThreadsInfos, error) {
	strs := Utils.FilterArray(strings.Split(strings.Trim(str, " "), "\n\n"))
	if len(strs) <= 1 {
		log.Error(str)
		return nil, errors.New("strings split err! len:" + strconv.Itoa(len(strs)) + " str:" + str)
	} else {
		ues := &ThreadsInfos{}

		for _, item := range strs {
			l, err := NewInitThreadsInfo(item)
			if err != nil {
				continue
			}
			ues.Threads = append(ues.Threads, l)
		}

		return ues, nil
	}

	return nil, errors.New("no ThreadsInfos Object!")
}

type ThreadsInfos struct {
	Threads []*ThreadsInfo
}

func (tis *ThreadsInfos) String(isParse bool) string {
	ret := ""
	for _, thread := range tis.Threads {
		ret = ret + thread.String(isParse) + "\n"
	}
	return ret
}

func (tis *ThreadsInfos) CrashStack() []*Line {
	for _, thread := range tis.Threads {
		if thread.IsCrashed {
			return thread.Lines
		}
	}
	return nil
}

func (tis *ThreadsInfos) CrashThread() (*ThreadsInfo, bool) {
	for _, thread := range tis.Threads {
		if thread.IsCrashed {
			return thread, true
		}
	}
	return nil, true
}
