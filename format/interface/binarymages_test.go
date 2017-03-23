package Interface

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInitImage(t *testing.T) {
	assert := assert.New(t)
	str := "       0x100034000 -        0x101757fff +DiSpecialDriver arm64  <1d946cdf81813082b3cd4e19468b578b> /var/mobile/Containers/Bundle/Application/AB330C81-3D29-46F5-B1E1-D353760E79CC/DiSpecialDriver.app/DiSpecialDriver"
	imaga, err := NewInitImage(str, true)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.ImageBaseAddressBegin, uint64(4295180288), "is not equal")
	assert.Equal(imaga.ImageBaseAddressEnd, uint64(4319444991), "is not equal")
	assert.Equal(imaga.BinaryDesignator, "+", "is not equal")
	assert.Equal(imaga.ImageName, "DiSpecialDriver", "is not equal")
	assert.Equal(imaga.ArchName, "arm64", "is not equal")
	assert.Equal(imaga.Uuid, "1d946cdf81813082b3cd4e19468b578b", "is not equal")
	assert.Equal(imaga.ImagePath, "/var/mobile/Containers/Bundle/Application/AB330C81-3D29-46F5-B1E1-D353760E79CC/DiSpecialDriver.app/DiSpecialDriver", "is not equal")

	assert.Equal(imaga.String(), str, "is not equal!")
}

func Test_NewInitImage2(t *testing.T) {
	assert := assert.New(t)
	str := "       0x101e08000 -        0x101e0bfff  MobileSubstrate.dylib arm64  <3134cfb2f722310ea2c742ae4dc131ab> /Library/MobileSubstrate/MobileSubstrate.dylib"
	imaga, err := NewInitImage(str, true)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.ImageBaseAddressBegin, uint64(4326457344), "is not equal")
	assert.Equal(imaga.ImageBaseAddressEnd, uint64(4326473727), "is not equal")
	assert.Equal(imaga.BinaryDesignator, " ", "is not equal")
	assert.Equal(imaga.ImageName, "MobileSubstrate.dylib", "is not equal")
	assert.Equal(imaga.ArchName, "arm64", "is not equal")
	assert.Equal(imaga.Uuid, "3134cfb2f722310ea2c742ae4dc131ab", "is not equal")
	assert.Equal(imaga.ImagePath, "/Library/MobileSubstrate/MobileSubstrate.dylib", "is not equal")

	assert.Equal(imaga.String(), str, "is not equal!")
}

func Test_NewInitImageErr(t *testing.T) {
	assert := assert.New(t)
	str := "adf 0x1000b0000 - 0x10124ffff +KDaijiaDriver arm64 <70eb26f62f0f35cdbc711cf4cdec7a71> /private/var/mobile/Containers/Bundle/Application/AAD091BC-CAC2-4E5A-9501-F00DEB6818D1/KDaijiaDriver.app/KDaijiaDriver"
	imaga, err := NewInitImage(str, true)
	log.Println(imaga, err)
	assert.NotNil(err)
	assert.Nil(imaga)
}

func Test_NewInitBinaryImages(t *testing.T) {
	assert := assert.New(t)
	str := `Binary Images:
       0x100034000 -        0x101757fff +DiSpecialDriver arm64  <1d946cdf81813082b3cd4e19468b578b> /var/mobile/Containers/Bundle/Application/AB330C81-3D29-46F5-B1E1-D353760E79CC/DiSpecialDriver.app/DiSpecialDriver
       0x101e08000 -        0x101e0bfff  MobileSubstrate.dylib arm64  <3134cfb2f722310ea2c742ae4dc131ab> /Library/MobileSubstrate/MobileSubstrate.dylib
       0x1020fc000 -        0x1020fffff  SubstrateLoader.dylib arm64  <54645dc0321231d88a022fd67a793278> /Library/Frameworks/CydiaSubstrate.framework/Libraries/SubstrateLoader.dylib
       0x102128000 -        0x102173fff  Flex.dylib arm64  <a1ffec6f2e3039f5a9642b5dba4abfda> /Library/MobileSubstrate/DynamicLibraries/Flex.dylib`

	imagas, err := NewInitBinaryImages(str, true)
	log.Println(imagas, err)
	assert.Nil(err)
	assert.Equal(len(imagas.Images), 4)

	assert.Equal(imagas.String(), str)
	assert.Equal(imagas.Uuids(), map[string]string{"DiSpecialDriver": "1d946cdf81813082b3cd4e19468b578b", "MobileSubstrate.dylib": "3134cfb2f722310ea2c742ae4dc131ab", "SubstrateLoader.dylib": "54645dc0321231d88a022fd67a793278", "Flex.dylib": "a1ffec6f2e3039f5a9642b5dba4abfda"}, "is not equal!")
}
