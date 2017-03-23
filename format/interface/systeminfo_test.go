package Interface

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInitSystemInfo(t *testing.T) {
	assert := assert.New(t)
	str := `Date/Time:       2017-01-06T09:52:42Z
Launch Time:     2017-01-06T09:52:33Z
OS Version:      iPhone OS 9.0.2 (13A452)
Report Version:  104
`
	imaga, err := NewInitSystemInfo(str)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.DateTime.Format("2006-01-02T15:04:05Z"), "2017-01-06T09:52:42Z", "is not equal")
	assert.Equal(imaga.LaunchTime.Format("2006-01-02T15:04:05Z"), "2017-01-06T09:52:33Z", "is not equal")
	assert.Equal(imaga.OSVersion, "iPhone OS 9.0.2 (13A452)", "is not equal")
	assert.Equal(imaga.ReportVersion, "104", "is not equal")

	assert.Equal(imaga.String(), str, "is not equal!")
}
