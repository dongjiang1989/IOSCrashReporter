package Interface

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInitExceptionCode(t *testing.T) {
	assert := assert.New(t)
	str := `Exception Type:  SIGABRT
Exception Codes: #0 at 0x1983131e0
Crashed Thread:  33
`
	imaga, err := NewInitExceptionCode(str)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.ExceptionCodes, "#0 at 0x1983131e0", "is not equal")
	assert.Equal(imaga.ExceptionType, "SIGABRT", "is not equal")
	assert.Equal(imaga.CrashedThread, "33", "is not equal")

	assert.Equal(imaga.String(), str, "is not equal!")
}
