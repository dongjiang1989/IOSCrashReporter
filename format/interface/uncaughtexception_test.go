package Interface

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInitUncaughtException(t *testing.T) {
	assert := assert.New(t)
	str1 := `Application Specific Information:
*** Terminating app due to uncaught exception 'UninitializedMessage', reason: ''
`
	str2 := `Last Exception Backtrace:
0   KDaijiaDriver                        0x100cb4aa0 0x100090000 + 12733088
1   KDaijiaDriver                        0x100cb3194 0x100090000 + 12726676
2   KDaijiaDriver                        0x100caa1f0 0x100090000 + 12689904
3   KDaijiaDriver                        0x100cb1ce4 0x100090000 + 12721380
4   Foundation                           0x1830a81cc 0x18308c000 + 115148
5   Foundation                           0x183169f28 0x18308c000 + 909096
6   libdispatch.dylib                    0x1940fd954 0x1940fc000 + 6484
7   libdispatch.dylib                    0x1941080a4 0x1940fc000 + 49316
8   libdispatch.dylib                    0x194100a5c 0x1940fc000 + 19036
9   libdispatch.dylib                    0x19410a318 0x1940fc000 + 58136
10  libdispatch.dylib                    0x19410bc4c 0x1940fc000 + 64588
11  libsystem_pthread.dylib              0x1942dd22c 0x1942dc000 + 4652
12  libsystem_pthread.dylib              0x1942dcef0 0x1942dc000 + 3824
`
	imaga, err := NewInitUncaughtException(str1)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.ApplicationSpecificInformation, "*** Terminating app due to uncaught exception 'UninitializedMessage', reason: ''", "is not equal")
	assert.Equal(imaga.String(), str1, "is not equal!")

	imaga1, err := NewInitLastExceptionBacktrace(str2)
	assert.Equal(len(imaga1.LastExceptionBacktrace), 13, "is not equal")

	assert.Equal(imaga1.String(false), str2, "is not equal!")
}

func Test_NewInitUncaughtException1(t *testing.T) {
	assert := assert.New(t)
	str := `Application Specific Information:
*** Terminating app due to uncaught exception 'UninitializedMessage', reason: ''
`
	imaga, err := NewInitUncaughtException(str)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.ApplicationSpecificInformation, "*** Terminating app due to uncaught exception 'UninitializedMessage', reason: ''", "is not equal")

	assert.Equal(imaga.String(), str, "is not equal!")
}
