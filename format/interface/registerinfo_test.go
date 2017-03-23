package Interface

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInitRegistersInfo(t *testing.T) {
	assert := assert.New(t)
	str := `Thread 0 crashed with ARM Thread State:
    pc: 0x19a08a87c     fp: 0x16fda1e10     sp: 0x16fda1dc0     x0: 0x00000000 
    x1: 0x00000000     x2: 0x00000001     x3: 0x199f9db8b     x4: 0x16fda1ea8 
    x5: 0x00000000     x6: 0x00000000     x7: 0x00000000     x8: 0x1a09e2050 
    x9: 0x1a09e6090    x10: 0x000019b6    x11: 0x1a36f96c2    x12: 0x1a36f96c2 
   x13: 0x0000000f    x14: 0x8000001f    x15: 0x80000023    x16: 0x00000025 
   x17: 0x101159148    x18: 0x00000000    x19: 0x0000000b    x20: 0x00000002 
   x21: 0x16fda1ea8    x22: 0x16fda1e40    x23: 0x101532480    x24: 0x18abc188c 
   x25: 0xda00a4f596311095    x26: 0x00000001    x27: 0x18abc1add    x28: 0x00000000 
    lr: 0x100d7fbf8   cpsr: 0x00000000 
`
	imaga, err := NewInitRegistersInfo(str, false)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(len(imaga.Pairs), 34, "is not equal")

	assert.Equal(imaga.String(), str, "is not equal!")
}
