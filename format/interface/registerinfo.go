package Interface

import (
	"common/util"
	"errors"
	//"strconv"
	"fmt"
	"strings"

	log "github.com/cihub/seelog"
)

//Thread 0 crashed with ARM Thread State:
//    pc: 0x19a08a87c     fp: 0x16fda1e10     sp: 0x16fda1dc0     x0: 0x00000000
//    x1: 0x00000000     x2: 0x00000001     x3: 0x199f9db8b     x4: 0x16fda1ea8
//    x5: 0x00000000     x6: 0x00000000     x7: 0x00000000     x8: 0x1a09e2050
//    x9: 0x1a09e6090    x10: 0x000019b6    x11: 0x1a36f96c2    x12: 0x1a36f96c2
//   x13: 0x0000000f    x14: 0x8000001f    x15: 0x80000023    x16: 0x00000025
//   x17: 0x101159148    x18: 0x00000000    x19: 0x0000000b    x20: 0x00000002
//   x21: 0x16fda1ea8    x22: 0x16fda1e40    x23: 0x101532480    x24: 0x18abc188c
//   x25: 0xda00a4f596311095    x26: 0x00000001    x27: 0x18abc1add    x28: 0x00000000
//    lr: 0x100d7fbf8   cpsr: 0x00000000

func NewInitRegistersInfo(str string, lp64 bool) (*RegistersInfo, error) {
	strs := Utils.FilterArray(strings.Split(strings.Trim(str, " "), "\n"))
	if len(strs) < 1 {
		log.Error(str)
		return nil, errors.New("strings split err! str:" + str)
	} else {
		sis := &RegistersInfo{}
		sis.PRIX64 = lp64
		sis.ThreadState = strs[0]

		for _, item := range strs[1:] {
			infoItems := Utils.FilterArray(strings.Split(strings.Trim(item, " "), "   ")) // 3 space
			for _, index := range infoItems {
				kvs := Utils.FilterArray(strings.Split(strings.Trim(index, " "), ": "))
				if len(kvs) != 2 {
					log.Error("Is not pairs! pair:", kvs)
					continue
				}
				si := &RegisterPair{}
				si.RegName = kvs[0]
				si.RegValue, _ = Utils.StringToUint64(kvs[1])

				sis.Pairs = append(sis.Pairs, si)
			}
		}

		return sis, nil
	}

	return nil, errors.New("no RegistersInfo Object!")
}

type RegisterPair struct {
	RegName  string
	RegValue uint64
}

func (rp *RegisterPair) String() string {
	return rp.RegName + ": " + Utils.Uint64ToString(rp.RegValue)
}

type RegistersInfo struct {
	PRIX64      bool
	ThreadState string
	Pairs       []*RegisterPair
}

func (sis *RegistersInfo) String() string {
	ret := ""
	ret = ret + sis.ThreadState + "\n"
	regColumn := 0
	for _, item := range sis.Pairs {
		if sis.PRIX64 {
			ret = ret + fmt.Sprintf("%6s: 0x%016x ", item.RegName, item.RegValue)

		} else {
			ret = ret + fmt.Sprintf("%6s: 0x%08x ", item.RegName, item.RegValue)

		}
		regColumn++
		if regColumn == 4 {
			ret = ret + "\n"
			regColumn = 0
		}
	}
	ret = ret + "\n" //last one
	return ret
}
