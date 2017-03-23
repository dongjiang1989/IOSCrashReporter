package Interface

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInitThreadsInfo(t *testing.T) {
	assert := assert.New(t)
	str := `Thread 10 Crashed:
0   libsystem_kernel.dylib               0x19a08a87c 0x19a070000 + 108668
1   libsystem_platform.dylib             0x19a14d94c 0x19a148000 + 22860
2   KDaijiaDriver                        0x10090a920 0x10005c000 + 9103648
3   KDaijiaDriver                        0x100904460 0x10005c000 + 9077856
4   KDaijiaDriver                        0x100904368 0x10005c000 + 9077608
5   OTRLocation.dylib                    0x103de1050 0x103ddc000 + 20560
6   OTRLocation.dylib                    0x103de1050 0x103ddc000 + 20560
7   GpsHookLibrary.dylib                 0x103dc5fe8 0x103dc4000 + 8168
8   CoreLocation                         0x1853f0850 0x1853e4000 + 51280
9   CoreLocation                         0x1853ecab0 0x1853e4000 + 35504
10  CoreLocation                         0x1853e6e90 0x1853e4000 + 11920
11  CoreFoundation                       0x184c1442c 0x184b38000 + 902188
12  CoreFoundation                       0x184c13d64 0x184b38000 + 900452
13  CoreFoundation                       0x184c121ac 0x184b38000 + 893356
14  CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
15  GraphicsServices                     0x18fd7c088 0x18fd70000 + 49288
16  UIKit                                0x18a258ffc 0x18a1d8000 + 528380
17  KDaijiaDriver                        0x1002771f8 0x10005c000 + 2208248
18  libdyld.dylib                        0x199f6e8b8 0x199f6c000 + 10424
`
	imaga, err := NewInitThreadsInfo(str)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.IsCrashed, true, "is not equal")
	assert.Equal(imaga.ThreadNum, int64(10), "is not equal")
	assert.Equal(len(imaga.Lines), 19, "is not equal")

	assert.Equal(imaga.String(false), str, "is not equal!")
}

func Test_NewInitThreadsInfo1(t *testing.T) {
	assert := assert.New(t)
	str := `Thread 1:
0   libsystem_kernel.dylib               0x19a08c4fc 0x19a070000 + 115964
1   libdispatch.dylib                    0x199f3f874 0x199f3c000 + 14452
`
	imaga, err := NewInitThreadsInfo(str)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(imaga.IsCrashed, false, "is not equal")
	assert.Equal(imaga.ThreadNum, int64(1), "is not equal")
	assert.Equal(len(imaga.Lines), 2, "is not equal")

	assert.Equal(imaga.String(false), str, "is not equal!")
}

//Thread 0 Crashed:
//0   libsystem_kernel.dylib               0x19a08a87c 0x19a070000 + 108668
//1   libsystem_platform.dylib             0x19a14d94c 0x19a148000 + 22860
//2   KDaijiaDriver                        0x10090a920 0x10005c000 + 9103648
//3   KDaijiaDriver                        0x100904460 0x10005c000 + 9077856
//4   KDaijiaDriver                        0x100904368 0x10005c000 + 9077608
//5   OTRLocation.dylib                    0x103de1050 0x103ddc000 + 20560
//6   OTRLocation.dylib                    0x103de1050 0x103ddc000 + 20560
//7   GpsHookLibrary.dylib                 0x103dc5fe8 0x103dc4000 + 8168
//8   CoreLocation                         0x1853f0850 0x1853e4000 + 51280
//9   CoreLocation                         0x1853ecab0 0x1853e4000 + 35504
//10  CoreLocation                         0x1853e6e90 0x1853e4000 + 11920
//11  CoreFoundation                       0x184c1442c 0x184b38000 + 902188
//12  CoreFoundation                       0x184c13d64 0x184b38000 + 900452
//13  CoreFoundation                       0x184c121ac 0x184b38000 + 893356
//14  CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//15  GraphicsServices                     0x18fd7c088 0x18fd70000 + 49288
//16  UIKit                                0x18a258ffc 0x18a1d8000 + 528380
//17  KDaijiaDriver                        0x1002771f8 0x10005c000 + 2208248
//18  libdyld.dylib                        0x199f6e8b8 0x199f6c000 + 10424

//Thread 1:
//0   libsystem_kernel.dylib               0x19a08c4fc 0x19a070000 + 115964
//1   libdispatch.dylib                    0x199f3f874 0x199f3c000 + 14452

//Thread 2:
//0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
//1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
//2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
//3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//4   WebCore                              0x19672a54c 0x1966b4000 + 484684
//5   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//6   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//7   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 3:
//0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
//1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
//2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
//3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//4   CFNetwork                            0x1843d9b84 0x18432c000 + 711556
//5   Foundation                           0x185b9fc80 0x185aac000 + 998528
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 4:
//0   libsystem_kernel.dylib               0x19a08b368 0x19a070000 + 111464
//1   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//2   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//3   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 5:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
//2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
//3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
//4   KDaijiaDriver                        0x100cfa558 0x10005c000 + 13231448
//5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 6:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
//2   KDaijiaDriver                        0x100ca2e80 0x10005c000 + 12873344
//3   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//4   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//5   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 7:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
//2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
//3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
//4   KDaijiaDriver                        0x100d1deec 0x10005c000 + 13377260
//5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 8:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
//2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
//3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
//4   KDaijiaDriver                        0x100d1deec 0x10005c000 + 13377260
//5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 9:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
//2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
//3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
//4   KDaijiaDriver                        0x100d1deec 0x10005c000 + 13377260
//5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 10:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
//2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
//3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
//4   KDaijiaDriver                        0x100d1deec 0x10005c000 + 13377260
//5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 11:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
//2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
//3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
//4   KDaijiaDriver                        0x100d1deec 0x10005c000 + 13377260
//5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 12:
//0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
//1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
//2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
//3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//4   libAVFAudio.dylib                    0x1832a9810 0x183264000 + 284688
//5   libAVFAudio.dylib                    0x18327e384 0x183264000 + 107396
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 13:
//0   libsystem_kernel.dylib               0x19a070a7c 0x19a070000 + 2684
//1   MediaToolbox                         0x187474aa4 0x187470000 + 19108
//2   CoreMedia                            0x1854c6a70 0x185484000 + 273008
//3   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//4   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//5   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 14:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x1009fc0f4 0x10005c000 + 10092788
//2   KDaijiaDriver                        0x100a11cfc 0x10005c000 + 10181884
//3   KDaijiaDriver                        0x100a11cd4 0x10005c000 + 10181844
//4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 15:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x1009fc0f4 0x10005c000 + 10092788
//2   KDaijiaDriver                        0x100a11cfc 0x10005c000 + 10181884
//3   KDaijiaDriver                        0x100a11cd4 0x10005c000 + 10181844
//4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 16:
//0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
//1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
//2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
//3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//4   Foundation                           0x185ab92bc 0x185aac000 + 53948
//5   Foundation                           0x185b0e8f4 0x185aac000 + 403700
//6   KDaijiaDriver                        0x100491610 0x10005c000 + 4412944
//7   Foundation                           0x185b9fc80 0x185aac000 + 998528
//8   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//9   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//10  libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 17:
//0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
//1   KDaijiaDriver                        0x100d6f3b0 0x10005c000 + 13710256
//2   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//3   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//4   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 18:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   libc++.1.dylib                       0x198d5c074 0x198d54000 + 32884
//2   JavaScriptCore                       0x1865bcad8 0x186288000 + 3361496
//3   JavaScriptCore                       0x1865bcb70 0x186288000 + 3361648
//4   JavaScriptCore                       0x186292824 0x186288000 + 43044
//5   JavaScriptCore                       0x186292734 0x186288000 + 42804
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 19:
//0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
//1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
//2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
//3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//4   Foundation                           0x185ab92bc 0x185aac000 + 53948
//5   KDaijiaDriver                        0x100542ce8 0x10005c000 + 5139688
//6   Foundation                           0x185b9fc80 0x185aac000 + 998528
//7   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//8   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//9   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 20:
//0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
//1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
//2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
//3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//4   Foundation                           0x185ab92bc 0x185aac000 + 53948
//5   KDaijiaDriver                        0x100904f90 0x10005c000 + 9080720
//6   Foundation                           0x185b9fc80 0x185aac000 + 998528
//7   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//8   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//9   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 21:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x100954d58 0x10005c000 + 9407832
//2   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
//3   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//4   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//5   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 22:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x10095563c 0x10005c000 + 9410108
//2   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
//3   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//4   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//5   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 23:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x10090f1c0 0x10005c000 + 9122240
//2   KDaijiaDriver                        0x10093488c 0x10005c000 + 9275532
//3   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
//4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 24:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x10093c948 0x10005c000 + 9308488
//2   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
//3   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//4   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//5   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 25:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x10090f1c0 0x10005c000 + 9122240
//2   KDaijiaDriver                        0x1009213c8 0x10005c000 + 9196488
//3   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
//4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 26:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x10090f1c0 0x10005c000 + 9122240
//2   KDaijiaDriver                        0x1009509dc 0x10005c000 + 9390556
//3   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
//4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 27:
//0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
//1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
//2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
//3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//4   Foundation                           0x185ab92bc 0x185aac000 + 53948
//5   KDaijiaDriver                        0x100904f90 0x10005c000 + 9080720
//6   Foundation                           0x185b9fc80 0x185aac000 + 998528
//7   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//8   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//9   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 28:
//0   libsystem_kernel.dylib               0x19a08b440 0x19a070000 + 111680
//1   libsystem_c.dylib                    0x19a012f60 0x199f9c000 + 487264
//2   KDaijiaDriver                        0x1009868e4 0x10005c000 + 9611492
//3   Foundation                           0x185b9fc80 0x185aac000 + 998528
//4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 29:
//0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
//1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
//2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
//3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//4   AudioToolbox                         0x183e1e808 0x183de4000 + 239624
//5   AudioToolbox                         0x183e10ba4 0x183de4000 + 183204
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 30:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
//2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
//3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
//4   KDaijiaDriver                        0x100d14b3c 0x10005c000 + 13339452
//5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
//6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 31:
//0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
//1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
//2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
//3   KDaijiaDriver                        0x100ce077c 0x10005c000 + 13125500
//4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 32:
//0   libsystem_kernel.dylib               0x19a08bb6c 0x19a070000 + 113516
//1   libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

//Thread 33:
//0   libsystem_kernel.dylib               0x19a08bb6c 0x19a070000 + 113516
//1   libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

//Thread 34:
//0   libsystem_kernel.dylib               0x19a070a94 0x19a070000 + 2708
//1   CoreLocation                         0x1853e6d94 0x1853e4000 + 11668
//2   CoreLocation                         0x1853ebae0 0x1853e4000 + 31456
//3   CoreLocation                         0x185425c90 0x1853e4000 + 269456
//4   CoreLocation                         0x185423b7c 0x1853e4000 + 260988
//5   CoreLocation                         0x185423ab0 0x1853e4000 + 260784
//6   CoreLocation                         0x185424cfc 0x1853e4000 + 265468
//7   libxpc.dylib                         0x19a1853b0 0x19a180000 + 21424
//8   libxpc.dylib                         0x19a183158 0x19a180000 + 12632
//9   libdispatch.dylib                    0x199f3d7a4 0x199f3c000 + 6052
//10  libdispatch.dylib                    0x199f41a90 0x199f3c000 + 23184
//11  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//12  libdispatch.dylib                    0x199f40ba4 0x199f3c000 + 19364
//13  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//14  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
//15  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//16  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
//17  libdispatch.dylib                    0x199f4b5bc 0x199f3c000 + 62908
//18  libdispatch.dylib                    0x199f4b2dc 0x199f3c000 + 62172
//19  libsystem_pthread.dylib              0x19a151470 0x19a150000 + 5232
//20  libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

//Thread 35:
//0   libsystem_kernel.dylib               0x19a070a94 0x19a070000 + 2708
//1   CoreLocation                         0x1853e6d94 0x1853e4000 + 11668
//2   CoreLocation                         0x1853eb808 0x1853e4000 + 30728
//3   CoreLocation                         0x185425f28 0x1853e4000 + 270120
//4   CoreLocation                         0x185423b7c 0x1853e4000 + 260988
//5   CoreLocation                         0x185423ab0 0x1853e4000 + 260784
//6   CoreLocation                         0x185424cfc 0x1853e4000 + 265468
//7   libxpc.dylib                         0x19a1853b0 0x19a180000 + 21424
//8   libxpc.dylib                         0x19a183158 0x19a180000 + 12632
//9   libdispatch.dylib                    0x199f3d7a4 0x199f3c000 + 6052
//10  libdispatch.dylib                    0x199f41a90 0x199f3c000 + 23184
//11  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//12  libdispatch.dylib                    0x199f40ba4 0x199f3c000 + 19364
//13  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//14  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
//15  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//16  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
//17  libdispatch.dylib                    0x199f3d6a8 0x199f3c000 + 5800
//18  libdispatch.dylib                    0x199f4bb40 0x199f3c000 + 64320
//19  libdispatch.dylib                    0x199f4b2dc 0x199f3c000 + 62172
//20  libsystem_pthread.dylib              0x19a151470 0x19a150000 + 5232
//21  libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

//Thread 36:
//0   libsystem_kernel.dylib               0x19a070a94 0x19a070000 + 2708
//1   CoreLocation                         0x1853e6d94 0x1853e4000 + 11668
//2   CoreLocation                         0x1853eb808 0x1853e4000 + 30728
//3   CoreLocation                         0x185425f28 0x1853e4000 + 270120
//4   CoreLocation                         0x185423b7c 0x1853e4000 + 260988
//5   CoreLocation                         0x185423ab0 0x1853e4000 + 260784
//6   CoreLocation                         0x185424cfc 0x1853e4000 + 265468
//7   libxpc.dylib                         0x19a1853b0 0x19a180000 + 21424
//8   libxpc.dylib                         0x19a183158 0x19a180000 + 12632
//9   libdispatch.dylib                    0x199f3d7a4 0x199f3c000 + 6052
//10  libdispatch.dylib                    0x199f41a90 0x199f3c000 + 23184
//11  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//12  libdispatch.dylib                    0x199f40ba4 0x199f3c000 + 19364
//13  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//14  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
//15  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//16  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
//17  libdispatch.dylib                    0x199f3d6a8 0x199f3c000 + 5800
//18  libdispatch.dylib                    0x199f4bb40 0x199f3c000 + 64320
//19  libdispatch.dylib                    0x199f4b2dc 0x199f3c000 + 62172
//20  libsystem_pthread.dylib              0x19a151470 0x19a150000 + 5232
//21  libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

//Thread 37:
//0   libsystem_kernel.dylib               0x19a08bb6c 0x19a070000 + 113516
//1   libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

//Thread 38:
//0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
//1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
//2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
//3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
//4   Foundation                           0x185ab92bc 0x185aac000 + 53948
//5   KDaijiaDriver                        0x100904f90 0x10005c000 + 9080720
//6   Foundation                           0x185b9fc80 0x185aac000 + 998528
//7   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
//8   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
//9   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

//Thread 39:
//0   libsystem_kernel.dylib               0x19a070a94 0x19a070000 + 2708
//1   CoreLocation                         0x1853e6d94 0x1853e4000 + 11668
//2   CoreLocation                         0x1853eb808 0x1853e4000 + 30728
//3   CoreLocation                         0x185425f28 0x1853e4000 + 270120
//4   CoreLocation                         0x185423b7c 0x1853e4000 + 260988
//5   CoreLocation                         0x185423ab0 0x1853e4000 + 260784
//6   CoreLocation                         0x185424cfc 0x1853e4000 + 265468
//7   libxpc.dylib                         0x19a1853b0 0x19a180000 + 21424
//8   libxpc.dylib                         0x19a183158 0x19a180000 + 12632
//9   libdispatch.dylib                    0x199f3d7a4 0x199f3c000 + 6052
//10  libdispatch.dylib                    0x199f41a90 0x199f3c000 + 23184
//11  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//12  libdispatch.dylib                    0x199f40ba4 0x199f3c000 + 19364
//13  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//14  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
//15  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
//16  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
//17  libdispatch.dylib                    0x199f3d6a8 0x199f3c000 + 5800
//18  libdispatch.dylib                    0x199f4bb40 0x199f3c000 + 64320
//19  libdispatch.dylib                    0x199f4b2dc 0x199f3c000 + 62172
//20  libsystem_pthread.dylib              0x19a151470 0x19a150000 + 5232
//21  libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

//Thread 40:
//0   libsystem_kernel.dylib               0x19a08bb6c 0x19a070000 + 113516
//1   libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

//Thread 41:
//0   libsystem_kernel.dylib               0x19a08bb6c 0x19a070000 + 113516
//1   libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

func Test_NewInitThreadsInfos(t *testing.T) {
	assert := assert.New(t)
	str := `Thread 0 Crashed:
0   libsystem_kernel.dylib               0x19a08a87c 0x19a070000 + 108668
1   libsystem_platform.dylib             0x19a14d94c 0x19a148000 + 22860
2   KDaijiaDriver                        0x10090a920 0x10005c000 + 9103648
3   KDaijiaDriver                        0x100904460 0x10005c000 + 9077856
4   KDaijiaDriver                        0x100904368 0x10005c000 + 9077608
5   OTRLocation.dylib                    0x103de1050 0x103ddc000 + 20560
6   OTRLocation.dylib                    0x103de1050 0x103ddc000 + 20560
7   GpsHookLibrary.dylib                 0x103dc5fe8 0x103dc4000 + 8168
8   CoreLocation                         0x1853f0850 0x1853e4000 + 51280
9   CoreLocation                         0x1853ecab0 0x1853e4000 + 35504
10  CoreLocation                         0x1853e6e90 0x1853e4000 + 11920
11  CoreFoundation                       0x184c1442c 0x184b38000 + 902188
12  CoreFoundation                       0x184c13d64 0x184b38000 + 900452
13  CoreFoundation                       0x184c121ac 0x184b38000 + 893356
14  CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
15  GraphicsServices                     0x18fd7c088 0x18fd70000 + 49288
16  UIKit                                0x18a258ffc 0x18a1d8000 + 528380
17  KDaijiaDriver                        0x1002771f8 0x10005c000 + 2208248
18  libdyld.dylib                        0x199f6e8b8 0x199f6c000 + 10424

Thread 1:
0   libsystem_kernel.dylib               0x19a08c4fc 0x19a070000 + 115964
1   libdispatch.dylib                    0x199f3f874 0x199f3c000 + 14452

Thread 2:
0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
4   WebCore                              0x19672a54c 0x1966b4000 + 484684
5   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
6   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
7   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 3:
0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
4   CFNetwork                            0x1843d9b84 0x18432c000 + 711556
5   Foundation                           0x185b9fc80 0x185aac000 + 998528
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 4:
0   libsystem_kernel.dylib               0x19a08b368 0x19a070000 + 111464
1   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
2   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
3   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 5:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
4   KDaijiaDriver                        0x100cfa558 0x10005c000 + 13231448
5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 6:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
2   KDaijiaDriver                        0x100ca2e80 0x10005c000 + 12873344
3   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
4   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
5   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 7:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
4   KDaijiaDriver                        0x100d1deec 0x10005c000 + 13377260
5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 8:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
4   KDaijiaDriver                        0x100d1deec 0x10005c000 + 13377260
5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 9:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
4   KDaijiaDriver                        0x100d1deec 0x10005c000 + 13377260
5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 10:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
4   KDaijiaDriver                        0x100d1deec 0x10005c000 + 13377260
5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 11:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
4   KDaijiaDriver                        0x100d1deec 0x10005c000 + 13377260
5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 12:
0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
4   libAVFAudio.dylib                    0x1832a9810 0x183264000 + 284688
5   libAVFAudio.dylib                    0x18327e384 0x183264000 + 107396
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 13:
0   libsystem_kernel.dylib               0x19a070a7c 0x19a070000 + 2684
1   MediaToolbox                         0x187474aa4 0x187470000 + 19108
2   CoreMedia                            0x1854c6a70 0x185484000 + 273008
3   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
4   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
5   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 14:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x1009fc0f4 0x10005c000 + 10092788
2   KDaijiaDriver                        0x100a11cfc 0x10005c000 + 10181884
3   KDaijiaDriver                        0x100a11cd4 0x10005c000 + 10181844
4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 15:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x1009fc0f4 0x10005c000 + 10092788
2   KDaijiaDriver                        0x100a11cfc 0x10005c000 + 10181884
3   KDaijiaDriver                        0x100a11cd4 0x10005c000 + 10181844
4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 16:
0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
4   Foundation                           0x185ab92bc 0x185aac000 + 53948
5   Foundation                           0x185b0e8f4 0x185aac000 + 403700
6   KDaijiaDriver                        0x100491610 0x10005c000 + 4412944
7   Foundation                           0x185b9fc80 0x185aac000 + 998528
8   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
9   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
10  libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 17:
0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
1   KDaijiaDriver                        0x100d6f3b0 0x10005c000 + 13710256
2   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
3   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
4   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 18:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   libc++.1.dylib                       0x198d5c074 0x198d54000 + 32884
2   JavaScriptCore                       0x1865bcad8 0x186288000 + 3361496
3   JavaScriptCore                       0x1865bcb70 0x186288000 + 3361648
4   JavaScriptCore                       0x186292824 0x186288000 + 43044
5   JavaScriptCore                       0x186292734 0x186288000 + 42804
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 19:
0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
4   Foundation                           0x185ab92bc 0x185aac000 + 53948
5   KDaijiaDriver                        0x100542ce8 0x10005c000 + 5139688
6   Foundation                           0x185b9fc80 0x185aac000 + 998528
7   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
8   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
9   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 20:
0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
4   Foundation                           0x185ab92bc 0x185aac000 + 53948
5   KDaijiaDriver                        0x100904f90 0x10005c000 + 9080720
6   Foundation                           0x185b9fc80 0x185aac000 + 998528
7   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
8   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
9   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 21:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x100954d58 0x10005c000 + 9407832
2   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
3   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
4   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
5   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 22:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x10095563c 0x10005c000 + 9410108
2   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
3   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
4   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
5   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 23:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x10090f1c0 0x10005c000 + 9122240
2   KDaijiaDriver                        0x10093488c 0x10005c000 + 9275532
3   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 24:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x10093c948 0x10005c000 + 9308488
2   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
3   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
4   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
5   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 25:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x10090f1c0 0x10005c000 + 9122240
2   KDaijiaDriver                        0x1009213c8 0x10005c000 + 9196488
3   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 26:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x10090f1c0 0x10005c000 + 9122240
2   KDaijiaDriver                        0x1009509dc 0x10005c000 + 9390556
3   KDaijiaDriver                        0x10096151c 0x10005c000 + 9458972
4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 27:
0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
4   Foundation                           0x185ab92bc 0x185aac000 + 53948
5   KDaijiaDriver                        0x100904f90 0x10005c000 + 9080720
6   Foundation                           0x185b9fc80 0x185aac000 + 998528
7   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
8   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
9   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 28:
0   libsystem_kernel.dylib               0x19a08b440 0x19a070000 + 111680
1   libsystem_c.dylib                    0x19a012f60 0x199f9c000 + 487264
2   KDaijiaDriver                        0x1009868e4 0x10005c000 + 9611492
3   Foundation                           0x185b9fc80 0x185aac000 + 998528
4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 29:
0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
4   AudioToolbox                         0x183e1e808 0x183de4000 + 239624
5   AudioToolbox                         0x183e10ba4 0x183de4000 + 183204
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 30:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
3   KDaijiaDriver                        0x100cdf56c 0x10005c000 + 13120876
4   KDaijiaDriver                        0x100d14b3c 0x10005c000 + 13339452
5   KDaijiaDriver                        0x100ce089c 0x10005c000 + 13125788
6   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
7   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
8   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 31:
0   libsystem_kernel.dylib               0x19a08af48 0x19a070000 + 110408
1   KDaijiaDriver                        0x100cdf33c 0x10005c000 + 13120316
2   KDaijiaDriver                        0x100cdf66c 0x10005c000 + 13121132
3   KDaijiaDriver                        0x100ce077c 0x10005c000 + 13125500
4   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
5   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
6   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 32:
0   libsystem_kernel.dylib               0x19a08bb6c 0x19a070000 + 113516
1   libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

Thread 33:
0   libsystem_kernel.dylib               0x19a08bb6c 0x19a070000 + 113516
1   libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

Thread 34:
0   libsystem_kernel.dylib               0x19a070a94 0x19a070000 + 2708
1   CoreLocation                         0x1853e6d94 0x1853e4000 + 11668
2   CoreLocation                         0x1853ebae0 0x1853e4000 + 31456
3   CoreLocation                         0x185425c90 0x1853e4000 + 269456
4   CoreLocation                         0x185423b7c 0x1853e4000 + 260988
5   CoreLocation                         0x185423ab0 0x1853e4000 + 260784
6   CoreLocation                         0x185424cfc 0x1853e4000 + 265468
7   libxpc.dylib                         0x19a1853b0 0x19a180000 + 21424
8   libxpc.dylib                         0x19a183158 0x19a180000 + 12632
9   libdispatch.dylib                    0x199f3d7a4 0x199f3c000 + 6052
10  libdispatch.dylib                    0x199f41a90 0x199f3c000 + 23184
11  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
12  libdispatch.dylib                    0x199f40ba4 0x199f3c000 + 19364
13  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
14  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
15  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
16  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
17  libdispatch.dylib                    0x199f4b5bc 0x199f3c000 + 62908
18  libdispatch.dylib                    0x199f4b2dc 0x199f3c000 + 62172
19  libsystem_pthread.dylib              0x19a151470 0x19a150000 + 5232
20  libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

Thread 35:
0   libsystem_kernel.dylib               0x19a070a94 0x19a070000 + 2708
1   CoreLocation                         0x1853e6d94 0x1853e4000 + 11668
2   CoreLocation                         0x1853eb808 0x1853e4000 + 30728
3   CoreLocation                         0x185425f28 0x1853e4000 + 270120
4   CoreLocation                         0x185423b7c 0x1853e4000 + 260988
5   CoreLocation                         0x185423ab0 0x1853e4000 + 260784
6   CoreLocation                         0x185424cfc 0x1853e4000 + 265468
7   libxpc.dylib                         0x19a1853b0 0x19a180000 + 21424
8   libxpc.dylib                         0x19a183158 0x19a180000 + 12632
9   libdispatch.dylib                    0x199f3d7a4 0x199f3c000 + 6052
10  libdispatch.dylib                    0x199f41a90 0x199f3c000 + 23184
11  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
12  libdispatch.dylib                    0x199f40ba4 0x199f3c000 + 19364
13  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
14  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
15  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
16  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
17  libdispatch.dylib                    0x199f3d6a8 0x199f3c000 + 5800
18  libdispatch.dylib                    0x199f4bb40 0x199f3c000 + 64320
19  libdispatch.dylib                    0x199f4b2dc 0x199f3c000 + 62172
20  libsystem_pthread.dylib              0x19a151470 0x19a150000 + 5232
21  libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

Thread 36:
0   libsystem_kernel.dylib               0x19a070a94 0x19a070000 + 2708
1   CoreLocation                         0x1853e6d94 0x1853e4000 + 11668
2   CoreLocation                         0x1853eb808 0x1853e4000 + 30728
3   CoreLocation                         0x185425f28 0x1853e4000 + 270120
4   CoreLocation                         0x185423b7c 0x1853e4000 + 260988
5   CoreLocation                         0x185423ab0 0x1853e4000 + 260784
6   CoreLocation                         0x185424cfc 0x1853e4000 + 265468
7   libxpc.dylib                         0x19a1853b0 0x19a180000 + 21424
8   libxpc.dylib                         0x19a183158 0x19a180000 + 12632
9   libdispatch.dylib                    0x199f3d7a4 0x199f3c000 + 6052
10  libdispatch.dylib                    0x199f41a90 0x199f3c000 + 23184
11  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
12  libdispatch.dylib                    0x199f40ba4 0x199f3c000 + 19364
13  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
14  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
15  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
16  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
17  libdispatch.dylib                    0x199f3d6a8 0x199f3c000 + 5800
18  libdispatch.dylib                    0x199f4bb40 0x199f3c000 + 64320
19  libdispatch.dylib                    0x199f4b2dc 0x199f3c000 + 62172
20  libsystem_pthread.dylib              0x19a151470 0x19a150000 + 5232
21  libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

Thread 37:
0   libsystem_kernel.dylib               0x19a08bb6c 0x19a070000 + 113516
1   libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

Thread 38:
0   libsystem_kernel.dylib               0x19a070a40 0x19a070000 + 2624
1   CoreFoundation                       0x184c14108 0x184b38000 + 901384
2   CoreFoundation                       0x184c11e0c 0x184b38000 + 892428
3   CoreFoundation                       0x184b40ca0 0x184b38000 + 36000
4   Foundation                           0x185ab92bc 0x185aac000 + 53948
5   KDaijiaDriver                        0x100904f90 0x10005c000 + 9080720
6   Foundation                           0x185b9fc80 0x185aac000 + 998528
7   libsystem_pthread.dylib              0x19a153b28 0x19a150000 + 15144
8   libsystem_pthread.dylib              0x19a153a8c 0x19a150000 + 14988
9   libsystem_pthread.dylib              0x19a151028 0x19a150000 + 4136

Thread 39:
0   libsystem_kernel.dylib               0x19a070a94 0x19a070000 + 2708
1   CoreLocation                         0x1853e6d94 0x1853e4000 + 11668
2   CoreLocation                         0x1853eb808 0x1853e4000 + 30728
3   CoreLocation                         0x185425f28 0x1853e4000 + 270120
4   CoreLocation                         0x185423b7c 0x1853e4000 + 260988
5   CoreLocation                         0x185423ab0 0x1853e4000 + 260784
6   CoreLocation                         0x185424cfc 0x1853e4000 + 265468
7   libxpc.dylib                         0x19a1853b0 0x19a180000 + 21424
8   libxpc.dylib                         0x19a183158 0x19a180000 + 12632
9   libdispatch.dylib                    0x199f3d7a4 0x199f3c000 + 6052
10  libdispatch.dylib                    0x199f41a90 0x199f3c000 + 23184
11  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
12  libdispatch.dylib                    0x199f40ba4 0x199f3c000 + 19364
13  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
14  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
15  libdispatch.dylib                    0x199f498c0 0x199f3c000 + 55488
16  libdispatch.dylib                    0x199f411ac 0x199f3c000 + 20908
17  libdispatch.dylib                    0x199f3d6a8 0x199f3c000 + 5800
18  libdispatch.dylib                    0x199f4bb40 0x199f3c000 + 64320
19  libdispatch.dylib                    0x199f4b2dc 0x199f3c000 + 62172
20  libsystem_pthread.dylib              0x19a151470 0x19a150000 + 5232
21  libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

Thread 40:
0   libsystem_kernel.dylib               0x19a08bb6c 0x19a070000 + 113516
1   libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

Thread 41:
0   libsystem_kernel.dylib               0x19a08bb6c 0x19a070000 + 113516
1   libsystem_pthread.dylib              0x19a151020 0x19a150000 + 4128

`
	imaga, err := NewInitThreadsInfos(str)
	log.Println(imaga, err)
	assert.Nil(err)

	assert.Equal(len(imaga.Threads), 42, "is not equal")

	assert.Equal(imaga.String(false), str, "is not equal!")
}
