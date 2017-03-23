package Interface

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInitApplicationHeader(t *testing.T) {
	assert := assert.New(t)
	str := `Incident Identifier: 15BF6315-3B8D-48DD-A671-FCCF16581847
CrashReporter Key:   A4E559D3-5F14-4EEB-A802-122D3FFCE1CF
Hardware Model:      iPhone8,1
Process:         DiSpecialDriver [1750]
Path:            /var/mobile/Containers/Bundle/Application/AB330C81-3D29-46F5-B1E1-D353760E79CC/DiSpecialDriver.app/DiSpecialDriver
Identifier:      com.xxx.DiSpecialDrivers
Version:         2.7.3 (2.7.3.1612211140)
Code Type:       ARM-64
Parent Process:  ??? [1]
`
	imaga, err := NewInitApplicationHeader(str)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.Identifier, "com.xxx.DiSpecialDrivers", "is not equal")
	assert.Equal(imaga.IncidentIdentifier, "15BF6315-3B8D-48DD-A671-FCCF16581847", "is not equal")
	assert.Equal(imaga.CrashReporterKey, "A4E559D3-5F14-4EEB-A802-122D3FFCE1CF", "is not equal")
	assert.Equal(imaga.HardwareModel, "iPhone8,1", "is not equal")
	assert.Equal(imaga.Process, "DiSpecialDriver [1750]", "is not equal")
	assert.Equal(imaga.Path, "/var/mobile/Containers/Bundle/Application/AB330C81-3D29-46F5-B1E1-D353760E79CC/DiSpecialDriver.app/DiSpecialDriver", "is not equal")
	assert.Equal(imaga.Version, "2.7.3 (2.7.3.1612211140)", "is not equal")
	assert.Equal(imaga.CodeType, "ARM-64", "is not equal")
	assert.Equal(imaga.ParentProcess, "??? [1]", "is not equal")

	assert.Equal(imaga.String(), str, "is not equal!")
}
