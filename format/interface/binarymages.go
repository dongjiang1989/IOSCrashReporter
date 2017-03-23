package Interface

import (
	"common/util"
	"errors"
	"fmt"
	"strings"

	log "github.com/cihub/seelog"
)

func NewInitImage(str string, lp64 bool) (*Image, error) {
	strs := Utils.FilterArray(strings.Split(strings.Trim(str, " "), " "))
	if len(strs) != 7 {
		log.Error(str)
		return nil, errors.New("strings split err! str:" + str)
	} else {
		begin, err := Utils.StringToUint64(strings.Trim(strs[0], " "))
		if err != nil {
			return nil, errors.New("ImageBaseAddressBegin can not to uint64 err! str:" + strs[0])
		}

		end, err := Utils.StringToUint64(strings.Trim(strs[2], " "))
		if err != nil {
			return nil, errors.New("ImageBaseAddressEnd can not to uint64 err! str:" + strs[2])
		}

		uuid := strings.TrimRight(strings.TrimLeft(strings.Trim(strs[5], " "), "<"), ">")
		if len(uuid) != 32 {
			return nil, errors.New("Get uuid err! uuid:" + uuid)
		}
		binarydesignator := " "
		imagename := strings.Trim(strs[3], " ")
		if imagename != "" && imagename[0] == '+' {
			binarydesignator = "+"
			imagename = strings.TrimLeft(strings.Trim(strs[3], " "), "+")
		}

		return &Image{PRIX64: lp64, ImageBaseAddressBegin: begin, ImageBaseAddressEnd: end, BinaryDesignator: binarydesignator, ImageName: imagename, ArchName: strings.Trim(strs[4], " "), Uuid: uuid, ImagePath: strings.Trim(strs[6], " ")}, nil
	}

	return nil, errors.New("no Image Object!")
}

// fmt : @"%18#" PRIx64 " - %18#" PRIx64 " %@%@ %@  <%@> %@\n"
// demo like:
//     0x1000b0000 - 0x10124ffff +KDaijiaDriver arm64 <70eb26f62f0f35cdbc711cf4cdec7a71> /private/var/mobile/Containers/Bundle/Application/AAD091BC-CAC2-4E5A-9501-F00DEB6818D1/KDaijiaDriver.app/KDaijiaDriver
type Image struct {
	PRIX64                bool
	ImageBaseAddressBegin uint64
	ImageBaseAddressEnd   uint64
	BinaryDesignator      string //"+" or ""
	ImageName             string
	ArchName              string
	Uuid                  string
	ImagePath             string
}

func (i *Image) String() string {
	if i.PRIX64 {
		return fmt.Sprintf("%#18x - %#18x %s%s %s  <%s> %s", i.ImageBaseAddressBegin, i.ImageBaseAddressEnd, i.BinaryDesignator, i.ImageName, i.ArchName, i.Uuid, i.ImagePath)
	} else {
		return fmt.Sprintf("%#10x - %#10x %s%s %s  <%s> %s", i.ImageBaseAddressBegin, i.ImageBaseAddressEnd, i.BinaryDesignator, i.ImageName, i.ArchName, i.Uuid, i.ImagePath)
	}
}

// -------------------------------------------------------------------------------------------------------------------------------------------------------
func NewInitBinaryImages(str string, lp64 bool) (*BinaryImages, error) {
	uuids := make(map[string]string, 0)
	bi := BinaryImages{UUidMaps: uuids}

	strs := strings.Split(str, "\n")
	for _, item := range strs {
		item := strings.Trim(item, " ")
		if item == "Binary Images:" {
			continue
		}

		if item != "" {
			image, err := NewInitImage(item, lp64)
			if err == nil {
				bi.Images = append(bi.Images, image)
			}
		}
	}
	if len(bi.Images) <= 0 {
		return nil, errors.New("Can not find Image!")
	}
	return &bi, nil
}

type BinaryImages struct {
	Images   []*Image
	UUidMaps map[string]string // moudle : uuid
}

func (bi *BinaryImages) Uuids() map[string]string {
	if len(bi.UUidMaps) <= 0 {
		ret := make(map[string]string)
		for _, image := range bi.Images {
			if image.Uuid != "" {
				ret[image.ImageName] = image.Uuid
			}
		}
		bi.UUidMaps = ret
	}
	return bi.UUidMaps
}

func (bi *BinaryImages) ModuleUuid() string {
	for _, image := range bi.Images {
		if image.BinaryDesignator == "+" {
			return image.Uuid
		}
	}
	return ""
}

func (bi *BinaryImages) ModuleName() string {
	for _, image := range bi.Images {
		if image.BinaryDesignator == "+" {
			return image.ImageName
		}
	}
	return ""
}

func (bi *BinaryImages) String() string {
	ret := "Binary Images:\n"
	for _, image := range bi.Images {
		ret = ret + image.String() + "\n"
	}
	return strings.TrimRight(ret, "\n")
}
