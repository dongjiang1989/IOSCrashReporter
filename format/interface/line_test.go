package Interface

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInitLine(t *testing.T) {
	assert := assert.New(t)
	str := `0   KDaijiaDriver                        0x100cb4aa0 0x100090000 + 12733088
`
	imaga, err := NewInitLine(str)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.BaseAddress, uint64(4295557120), "is not equal")
	assert.Equal(imaga.FrameIndex, int64(0), "is not equal")
	assert.Equal(imaga.FrameInfoPointer, uint64(4308290208), "is not equal")
	assert.Equal(imaga.ImageName, "KDaijiaDriver", "is not equal")
	assert.Equal(imaga.PcOffset, int64(12733088), "is not equal")

	assert.Equal(imaga.String(false), str, "is not equal!")
}

func Test_NewInitLine2(t *testing.T) {
	assert := assert.New(t)
	str := `1   KDaijiaDriver                        0x100cb3194 0x100090000 + 12726676
`
	imaga, err := NewInitLine(str)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.BaseAddress, uint64(4295557120), "is not equal")
	assert.Equal(imaga.FrameIndex, int64(1), "is not equal")
	assert.Equal(imaga.FrameInfoPointer, uint64(4308283796), "is not equal")
	assert.Equal(imaga.ImageName, "KDaijiaDriver", "is not equal")
	assert.Equal(imaga.PcOffset, int64(12726676), "is not equal")

	assert.Equal(imaga.String(false), str, "is not equal!")
}
