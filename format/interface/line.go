package Interface

import (
	"common/util"
	"errors"
	"fmt"
	"strconv"
	"strings"

	log "github.com/cihub/seelog"
)

func NewInitLine(str string) (*Line, error) {
	strs := Utils.FilterArray(strings.SplitN(strings.Trim(strings.Trim(str, " "), "\n"), " + ", 2))
	if len(strs) != 2 {
		log.Error(str)
		return nil, errors.New("strings split err! str:" + str)
	} else {
		if strs[1] != "" {
			line := &Line{}
			offset, err := strconv.ParseInt(strs[1], 10, 64) // offset 10 进制
			if err != nil {
				log.Warn("get offset err! str:" + str + " err:" + err.Error())
				line.PcOffset = 0
				line.ParseStr = strs[1]
			} else {
				line.PcOffset = offset
				line.ParseStr = "" //init "" string
			}

			items := Utils.FilterArray(strings.Split(strs[0], " "))
			if len(items) != 4 {
				log.Error(strs[0])
				return nil, errors.New("strings split err! str:" + strs[0])
			} else {
				FrameIndex, err := Utils.StringToInt64(items[0])
				if err != nil {
					return nil, errors.New("FrameIndex can not to int64 err! str:" + strs[0])
				}

				ImageName := items[1]

				FrameInfoPointer, err := Utils.StringToUint64(items[2])
				if err != nil {
					return nil, errors.New("FrameInfoPointer can not to uint64 err! str:" + strs[0])
				}

				BaseAddress, err := Utils.StringToUint64(items[3])
				if err != nil {
					return nil, errors.New("BaseAddress can not to uint64 err! str:" + strs[0])
				}

				line.FrameIndex = FrameIndex
				line.FrameInfoPointer = FrameInfoPointer
				line.BaseAddress = BaseAddress
				line.ImageName = ImageName
				return line, nil
			}

		} else {
			log.Error("offset is not find! str:" + str)
			return nil, errors.New("offset is not find! str:" + str)
		}
	}

	return nil, errors.New("no Line Object!")
}

//[NSString stringWithFormat: @"%-4ld%-35S 0x%0*" PRIx64 " %@\n",
//          (long) frameIndex,
//          (const uint16_t *)[imageName cStringUsingEncoding: NSUTF16StringEncoding],
//          lp64 ? 16 : 8, frameInfo.instructionPointer,
//          symbolString];
//"0x%" PRIx64 " + %" PRId64

//like: 2 libdispatch.dylib 0x1941ca20c 0x1941c4000 + 25100
type Line struct {
	FrameIndex       int64
	ImageName        string
	FrameInfoPointer uint64
	BaseAddress      uint64
	PcOffset         int64
	//Uuid             string

	//paste string
	ParseStr string
}

//TODO: 8 36
func (l *Line) String(isParse bool) string {
	ret := ""
	if isParse && l.ParseStr != "" {
		ret = ret + fmt.Sprintf("%-4d%-36s 0x%0*x 0x%x + %s\n", l.FrameIndex, l.ImageName, 8, l.FrameInfoPointer, l.BaseAddress, l.ParseStr)

	} else {
		ret = ret + fmt.Sprintf("%-4d%-36s 0x%0*x 0x%x + %d\n", l.FrameIndex, l.ImageName, 8, l.FrameInfoPointer, l.BaseAddress, l.PcOffset)
	}
	return ret
}

func (l *Line) SetParseStr(str string) {
	l.ParseStr = strings.Trim(strings.Trim(str, " "), "\n")
}
