package Format

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInitCrashReportException(t *testing.T) {
	assert := assert.New(t)
	str := `Incident Identifier: C5BBAE93-8CFC-4512-A0BF-FF072D65A28E
CrashReporter Key:   5DEBFD59-DEC6-4181-A37D-0CE8AEC7F473
Hardware Model:      iPhone6,1
Process:         KDaijiaDriver [916]
Path:            /var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/KDaijiaDriver
Identifier:      com.xxxx.liangjian
Version:         1567 (1567)
Code Type:       ARM
Parent Process:  launchd [1]

Date/Time:       2017-01-05T15:45:34Z
Launch Time:     2017-01-05T12:50:13Z
OS Version:      iPhone OS 9.1 (13B143)
Report Version:  104

Exception Type:  SIGSEGV
Exception Codes: SEGV_ACCERR at 0x0
Crashed Thread:  0

Thread 0 Crashed:
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

Thread 0 crashed with ARM Thread State:
    pc: 0x19a08a87c     fp: 0x16fda1e10     sp: 0x16fda1dc0     x0: 0x00000000 
    x1: 0x00000000     x2: 0x00000001     x3: 0x199f9db8b     x4: 0x16fda1ea8 
    x5: 0x00000000     x6: 0x00000000     x7: 0x00000000     x8: 0x1a09e2050 
    x9: 0x1a09e6090    x10: 0x000019b6    x11: 0x1a36f96c2    x12: 0x1a36f96c2 
   x13: 0x0000000f    x14: 0x8000001f    x15: 0x80000023    x16: 0x00000025 
   x17: 0x101159148    x18: 0x00000000    x19: 0x0000000b    x20: 0x00000002 
   x21: 0x16fda1ea8    x22: 0x16fda1e40    x23: 0x101532480    x24: 0x18abc188c 
   x25: 0xda00a4f596311095    x26: 0x00000001    x27: 0x18abc1add    x28: 0x00000000 
    lr: 0x100d7fbf8   cpsr: 0x00000000 

Binary Images:
0x10005c000 - 0x101157fff +KDaijiaDriver arm64  <3f043d63e7503e9688fca40afa778931> /var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/KDaijiaDriver
0x10166c000 - 0x10166ffff  MobileSubstrate.dylib arm64  <3134cfb2f722310ea2c742ae4dc131ab> /Library/MobileSubstrate/MobileSubstrate.dylib
0x101690000 - 0x1018b3fff  libswiftCore.dylib arm64  <0969696c7f0b3c9e8eef73014242042e> /private/var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/Frameworks/libswiftCore.dylib
0x101b60000 - 0x101b6ffff  libswiftCoreGraphics.dylib arm64  <7198afd2badc33cba52d66e2a526bb1c> /private/var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/Frameworks/libswiftCoreGraphics.dylib
0x101b98000 - 0x101b9ffff  libswiftCoreImage.dylib arm64  <4e476f8e88ca34e39f9db870b8ebc7a1> /private/var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/Frameworks/libswiftCoreImage.dylib
0x101bac000 - 0x101bb3fff  libswiftDarwin.dylib arm64  <3c4418f9722635ddba9748a460382966> /private/var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/Frameworks/libswiftDarwin.dylib
0x101bd4000 - 0x101beffff  libswiftDispatch.dylib arm64  <43701ffffad133d29454d60a35c852c9> /private/var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/Frameworks/libswiftDispatch.dylib
0x101c44000 - 0x101d07fff  libswiftFoundation.dylib arm64  <9023e16cf7c03bd1af94d1fe66fac0ed> /private/var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/Frameworks/libswiftFoundation.dylib
0x101e10000 - 0x101e17fff  libswiftObjectiveC.dylib arm64  <fd15443960fd39e99b9662aa85746226> /private/var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/Frameworks/libswiftObjectiveC.dylib
0x101e3c000 - 0x101e43fff  libswiftQuartzCore.dylib arm64  <3724aa8e874837ccbd94e5ec0f530c91> /private/var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/Frameworks/libswiftQuartzCore.dylib
0x101e4c000 - 0x101e57fff  libswiftUIKit.dylib arm64  <72abc5d06b00391ebff2f6bde050ebaf> /private/var/mobile/Containers/Bundle/Application/8E3BAAE6-891C-47D6-BD9D-124D9CE3485B/KDaijiaDriver.app/Frameworks/libswiftUIKit.dylib
0x101ea8000 - 0x101eaffff  TEGPS.dylib arm64  <cdb4cf44266e3789879670a5c92d68df> /Library/MobileSubstrate/DynamicLibraries/TEGPS.dylib
0x101eb8000 - 0x101ebffff  TEMain.dylib arm64  <a989fc237acd392a89e154ce71083daa> /Library/MobileSubstrate/DynamicLibraries/TEMain.dylib
0x101ec8000 - 0x101ed3fff  xCon.dylib arm64  <4b4e7d71e56036f28e9a5c5269fa610e> /Library/MobileSubstrate/DynamicLibraries/xCon.dylib
0x1020ec000 - 0x1020effff  SubstrateLoader.dylib arm64  <54645dc0321231d88a022fd67a793278> /Library/Frameworks/CydiaSubstrate.framework/Libraries/SubstrateLoader.dylib
0x102118000 - 0x10211bfff  FakeClockUp.dylib arm64  <551c316267323724b2d56eb145c33ba5> /Library/MobileSubstrate/DynamicLibraries/FakeClockUp.dylib
0x102164000 - 0x1021affff  Flex.dylib arm64  <a1ffec6f2e3039f5a9642b5dba4abfda> /Library/MobileSubstrate/DynamicLibraries/Flex.dylib
0x103d5c000 - 0x103d63fff  librocketbootstrap.dylib arm64  <32890c4f316e307a918a53552e347b46> /usr/lib/librocketbootstrap.dylib
0x103d6c000 - 0x103d83fff  libsubstrate.dylib arm64  <9da13c499976324293cbf1ad36f5a052> /usr/lib/libsubstrate.dylib
0x103d8c000 - 0x103d9ffff  GPSTravellerTweakVIP.dylib arm64  <5c51270d43023645860a1c906978f54d> /Library/MobileSubstrate/DynamicLibraries/GPSTravellerTweakVIP.dylib
0x103dc4000 - 0x103dcbfff  GpsHookLibrary.dylib arm64  <05d17b42f2fc3418b62aed824cd17271> /Library/MobileSubstrate/DynamicLibraries/GpsHookLibrary.dylib
0x103ddc000 - 0x103de3fff  OTRLocation.dylib arm64  <ce6d0a7164ef312bb9ba01fd69148885> /Library/MobileSubstrate/DynamicLibraries/OTRLocation.dylib
0x182e14000 - 0x182febfff  RawCamera arm64  <59145b539e453a43af15da59d68d3537> /System/Library/CoreServices/RawCamera.bundle/RawCamera
0x183008000 - 0x1830ccfff  AGXGLDriver arm64  <3541e84ba1d23dd2a24c42798b7a0fbb> /System/Library/Extensions/AGXGLDriver.bundle/AGXGLDriver
0x1830e8000 - 0x183261fff  AVFoundation arm64  <62f9cc16934936bdbae949eff468326a> /System/Library/Frameworks/AVFoundation.framework/AVFoundation
0x183264000 - 0x1832dcfff  libAVFAudio.dylib arm64  <85d9da197673373dad866aa3adafb789> /System/Library/Frameworks/AVFoundation.framework/libAVFAudio.dylib
0x18332c000 - 0x18332cfff  Accelerate arm64  <52640254a7693f1a8c93b76f152d9a01> /System/Library/Frameworks/Accelerate.framework/Accelerate
0x183330000 - 0x183342fff  libCGInterfaces.dylib arm64  <6a7f965860a937b0be0d9b7e236e4e94> /System/Library/Frameworks/Accelerate.framework/Frameworks/vImage.framework/Libraries/libCGInterfaces.dylib
0x183344000 - 0x1835a0fff  vImage arm64  <a03cc73bb24b378dbe34367137d0dede> /System/Library/Frameworks/Accelerate.framework/Frameworks/vImage.framework/vImage
0x1835a4000 - 0x18364ffff  libBLAS.dylib arm64  <60ad0764403e345d8052faa2af69b0e2> /System/Library/Frameworks/Accelerate.framework/Frameworks/vecLib.framework/libBLAS.dylib
0x183650000 - 0x1839b5fff  libLAPACK.dylib arm64  <e580decf91913736ae6eeb322862a7ce> /System/Library/Frameworks/Accelerate.framework/Frameworks/vecLib.framework/libLAPACK.dylib
0x1839b8000 - 0x1839ccfff  libLinearAlgebra.dylib arm64  <9b0f1a9d4b0732a99de47e6a432afbf8> /System/Library/Frameworks/Accelerate.framework/Frameworks/vecLib.framework/libLinearAlgebra.dylib
0x1839d0000 - 0x1839dffff  libSparseBLAS.dylib arm64  <4a63abb231bb311d81eaa09cc4f1bde2> /System/Library/Frameworks/Accelerate.framework/Frameworks/vecLib.framework/libSparseBLAS.dylib
0x1839e0000 - 0x183a4cfff  libvDSP.dylib arm64  <0ce42521abc03db78fa29dfc23b655c4> /System/Library/Frameworks/Accelerate.framework/Frameworks/vecLib.framework/libvDSP.dylib
0x183a50000 - 0x183a72fff  libvMisc.dylib arm64  <71a81d7189ca3ee8b646cb3fef50a781> /System/Library/Frameworks/Accelerate.framework/Frameworks/vecLib.framework/libvMisc.dylib
0x183a74000 - 0x183a74fff  vecLib arm64  <d70ec799b3ed3850a50a89ef334c6f73> /System/Library/Frameworks/Accelerate.framework/Frameworks/vecLib.framework/vecLib
0x183a78000 - 0x183aaffff  Accounts arm64  <d450c8357bed3844af4977c553405f72> /System/Library/Frameworks/Accounts.framework/Accounts
0x183ab4000 - 0x183b2ffff  AddressBook arm64  <50989572ec193435a3be9a5b5d4b049a> /System/Library/Frameworks/AddressBook.framework/AddressBook
0x183b30000 - 0x183c05fff  AddressBookUI arm64  <d4f59dc6959e31a1be6ee551a1bd4349> /System/Library/Frameworks/AddressBookUI.framework/AddressBookUI
0x183c08000 - 0x183c1bfff  AssetsLibrary arm64  <248d6b2b00b635bc891dd23b5e9fd582> /System/Library/Frameworks/AssetsLibrary.framework/AssetsLibrary
0x183c1c000 - 0x183de2fff  AudioCodecs arm64  <901619a8d1db30cbb01f38ec786590b7> /System/Library/Frameworks/AudioToolbox.framework/AudioCodecs
0x183de4000 - 0x18410afff  AudioToolbox arm64  <c5d82c44162134029e93db6b2d932258> /System/Library/Frameworks/AudioToolbox.framework/AudioToolbox
0x18432c000 - 0x1845b9fff  CFNetwork arm64  <8dfb8c9f8cf432929e7bdc050b518fac> /System/Library/Frameworks/CFNetwork.framework/CFNetwork
0x1845bc000 - 0x18466bfff  CloudKit arm64  <58c8abfb8f42382f83b0f473a508fad0> /System/Library/Frameworks/CloudKit.framework/CloudKit
0x18466c000 - 0x1846f3fff  Contacts arm64  <11560b1bad78334daacc38851f7e5182> /System/Library/Frameworks/Contacts.framework/Contacts
0x1846f4000 - 0x1847f2fff  ContactsUI arm64  <1050eff6141b3013b44bc9c823872b1d> /System/Library/Frameworks/ContactsUI.framework/ContactsUI
0x1847f4000 - 0x1848affff  CoreAudio arm64  <3e2a210c14fe3aea910a71d076244cb3> /System/Library/Frameworks/CoreAudio.framework/CoreAudio
0x1848c8000 - 0x1848e9fff  CoreBluetooth arm64  <a1b17bfbccd23ca989973171b638b575> /System/Library/Frameworks/CoreBluetooth.framework/CoreBluetooth
0x1848ec000 - 0x184b35fff  CoreData arm64  <8c6127652dce3ab8adf9e0ce964c67d2> /System/Library/Frameworks/CoreData.framework/CoreData
0x184b38000 - 0x184eaffff  CoreFoundation arm64  <ff1340241cb339f88865829dd82f7638> /System/Library/Frameworks/CoreFoundation.framework/CoreFoundation
0x184eb0000 - 0x18501dfff  CoreGraphics arm64  <be820db45b2b3c1587783e96e4976102> /System/Library/Frameworks/CoreGraphics.framework/CoreGraphics
0x185028000 - 0x18502afff  libCGXType.A.dylib arm64  <5ef23d77e9e138cb94ab74822bb3f654> /System/Library/Frameworks/CoreGraphics.framework/Resources/libCGXType.A.dylib
0x185220000 - 0x18523efff  libRIP.A.dylib arm64  <cfcc47562efe39fa993e2d9c09785301> /System/Library/Frameworks/CoreGraphics.framework/Resources/libRIP.A.dylib
0x185240000 - 0x1853e2fff  CoreImage arm64  <79f96b1eb92e30e386d23a390140b6ea> /System/Library/Frameworks/CoreImage.framework/CoreImage
0x1853e4000 - 0x18544bfff  CoreLocation arm64  <161413ea48003214bb2aa2fc980c11a5> /System/Library/Frameworks/CoreLocation.framework/CoreLocation
0x185484000 - 0x185571fff  CoreMedia arm64  <bec9f808bb4a32a0bd0e011e7331e1a0> /System/Library/Frameworks/CoreMedia.framework/CoreMedia
0x185574000 - 0x18567efff  CoreMotion arm64  <33934564729c335c8ec7537668c4590d> /System/Library/Frameworks/CoreMotion.framework/CoreMotion
0x185680000 - 0x1856b8fff  CoreSpotlight arm64  <b35bc937bb253e07bc14ea0f6f2e3306> /System/Library/Frameworks/CoreSpotlight.framework/CoreSpotlight
0x1856bc000 - 0x18572efff  CoreTelephony arm64  <fd3ad09f419c3f5f96e17e7037172766> /System/Library/Frameworks/CoreTelephony.framework/CoreTelephony
0x185730000 - 0x18585bfff  CoreText arm64  <d80bf19c023d3fbd9bb5cd4769134622> /System/Library/Frameworks/CoreText.framework/CoreText
0x18585c000 - 0x18587cfff  CoreVideo arm64  <ec6ce9482a57346f8535cee4ba5644aa> /System/Library/Frameworks/CoreVideo.framework/CoreVideo
0x185880000 - 0x185910fff  EventKit arm64  <d7c8d0c69fef3a7c8c8340f16b76b00f> /System/Library/Frameworks/EventKit.framework/EventKit
0x185aac000 - 0x185d19fff  Foundation arm64  <ac3005793e833b088c2c4b69c11329d6> /System/Library/Frameworks/Foundation.framework/Foundation
0x185d1c000 - 0x185d4afff  GLKit arm64  <069efc428ff5308081f4f095bfd36e8b> /System/Library/Frameworks/GLKit.framework/GLKit
0x185d4c000 - 0x185d75fff  GSS arm64  <d5528ff4fbee3c57b38ccfc919b4a307> /System/Library/Frameworks/GSS.framework/GSS
0x185ef0000 - 0x185f60fff  IOKit arm64  <dd9b3af7bc4a368998b0317b8b6a8a8d> /System/Library/Frameworks/IOKit.framework/Versions/A/IOKit
0x185f64000 - 0x186285fff  ImageIO arm64  <c6363beea9073532b4bff8d2d5220fd7> /System/Library/Frameworks/ImageIO.framework/ImageIO
0x186288000 - 0x18686cfff  JavaScriptCore arm64  <acec26eb63153b649c17b2daf7a8952b> /System/Library/Frameworks/JavaScriptCore.framework/JavaScriptCore
0x187080000 - 0x1871fefff  MapKit arm64  <b0f1adea7982376986b580eb16709139> /System/Library/Frameworks/MapKit.framework/MapKit
0x187200000 - 0x187206fff  MediaAccessibility arm64  <901d649a3032323c98f806768c3e849f> /System/Library/Frameworks/MediaAccessibility.framework/MediaAccessibility
0x187208000 - 0x18746efff  MediaPlayer arm64  <3cfbc6cb1c27353380e324bd1e01b381> /System/Library/Frameworks/MediaPlayer.framework/MediaPlayer
0x187470000 - 0x187887fff  MediaToolbox arm64  <8de7937b4cee31a49914824f8222bfc8> /System/Library/Frameworks/MediaToolbox.framework/MediaToolbox
0x187888000 - 0x187962fff  MessageUI arm64  <9cba641275b235d5bf00f25be88c0bc5> /System/Library/Frameworks/MessageUI.framework/MessageUI
0x187964000 - 0x1879aafff  Metal arm64  <bbba4be3d7153e049eecff44457d12dc> /System/Library/Frameworks/Metal.framework/Metal
0x1879c4000 - 0x187a0ffff  MetalPerformanceShaders arm64  <3406a595bca93b138b8f9db0ee84f218> /System/Library/Frameworks/MetalPerformanceShaders.framework/MetalPerformanceShaders
0x187a10000 - 0x187ad1fff  MobileCoreServices arm64  <981b707bd8d03fa1923ee0cf2938154e> /System/Library/Frameworks/MobileCoreServices.framework/MobileCoreServices
0x187ad4000 - 0x187eb4fff  ModelIO arm64  <e9c84d57a6543c97bcd33e255d7e777a> /System/Library/Frameworks/ModelIO.framework/ModelIO
0x187f34000 - 0x188058fff  NetworkExtension arm64  <4966900fbeb431659a45b8ac355932b3> /System/Library/Frameworks/NetworkExtension.framework/NetworkExtension
0x188078000 - 0x1880adfff  OpenAL arm64  <052218babcbf3f94a1964270a2fe6cc5> /System/Library/Frameworks/OpenAL.framework/OpenAL
0x1880b0000 - 0x18819bfff  GLEngine arm64  <4ed2122627173e6b878d4e31906f5c1a> /System/Library/Frameworks/OpenGLES.framework/GLEngine.bundle/GLEngine
0x18819c000 - 0x1881a6fff  OpenGLES arm64  <188577111b3f35e8b397449355025e7b> /System/Library/Frameworks/OpenGLES.framework/OpenGLES
0x1881ac000 - 0x1881adfff  libCVMSPluginSupport.dylib arm64  <6554f91d84c93677acd5f898a3eb646f> /System/Library/Frameworks/OpenGLES.framework/libCVMSPluginSupport.dylib
0x1881b0000 - 0x1881b3fff  libCoreFSCache.dylib arm64  <bb098178791e3458bf7571bf1a3e13a6> /System/Library/Frameworks/OpenGLES.framework/libCoreFSCache.dylib
0x1881b4000 - 0x1881b8fff  libCoreVMClient.dylib arm64  <1a29dc38de6d3e4884876e8fbe02a11e> /System/Library/Frameworks/OpenGLES.framework/libCoreVMClient.dylib
0x1881bc000 - 0x1881c6fff  libGFXShared.dylib arm64  <0b1c85c5615b3e448be01f175cf528ac> /System/Library/Frameworks/OpenGLES.framework/libGFXShared.dylib
0x1881c8000 - 0x18820efff  libGLImage.dylib arm64  <7d88e13e8f81365cb07175cc27c56fc3> /System/Library/Frameworks/OpenGLES.framework/libGLImage.dylib
0x188210000 - 0x18837cfff  libGLProgrammability.dylib arm64  <99c9b36ae3e63e2b96225d7d0fa0b41e> /System/Library/Frameworks/OpenGLES.framework/libGLProgrammability.dylib
0x1899dc000 - 0x189b7efff  QuartzCore arm64  <40ea419d9c5631e0901ea9835bf602b9> /System/Library/Frameworks/QuartzCore.framework/QuartzCore
0x189b80000 - 0x189bd4fff  QuickLook arm64  <9bfe3d0793793111b66e75cc6b27a639> /System/Library/Frameworks/QuickLook.framework/QuickLook
0x189efc000 - 0x189f69fff  Security arm64  <1550eab724483ebb974565feb290ff26> /System/Library/Frameworks/Security.framework/Security
0x18a16c000 - 0x18a1c5fff  SystemConfiguration arm64  <d8a46ce0c1bb3c569a6d0475ee41663c> /System/Library/Frameworks/SystemConfiguration.framework/SystemConfiguration
0x18a1d8000 - 0x18adc6fff  UIKit arm64  <2cc71901f4823476a51ab0b06010d56c> /System/Library/Frameworks/UIKit.framework/UIKit
0x18adc8000 - 0x18ae48fff  VideoToolbox arm64  <d831910c269a3cf9992a60a6e996315d> /System/Library/Frameworks/VideoToolbox.framework/VideoToolbox
0x18aec8000 - 0x18b1adfff  WebKit arm64  <727ad514fb873f62a177304c7563ee40> /System/Library/Frameworks/WebKit.framework/WebKit
0x18b548000 - 0x18b575fff  ACTFramework arm64  <38fa7495c23838f5aa5b0648e0a00099> /System/Library/PrivateFrameworks/ACTFramework.framework/ACTFramework
0x18b578000 - 0x18b57bfff  AGXCompilerConnection arm64  <c2374341ef3f34f0840e687b132f0f4c> /System/Library/PrivateFrameworks/AGXCompilerConnection.framework/AGXCompilerConnection
0x18b5f4000 - 0x18b5fafff  AOSNotification arm64  <41c7e4984b223f5bbd293b4493c45875> /System/Library/PrivateFrameworks/AOSNotification.framework/AOSNotification
0x18b800000 - 0x18b863fff  AccountsDaemon arm64  <173c21691d9d37c9b7dce02c007cd881> /System/Library/PrivateFrameworks/AccountsDaemon.framework/AccountsDaemon
0x18b88c000 - 0x18b891fff  AggregateDictionary arm64  <de33026f793c38a0a9f78618007ac2f5> /System/Library/PrivateFrameworks/AggregateDictionary.framework/AggregateDictionary
0x18bd70000 - 0x18bdb7fff  AppSupport arm64  <f59084c396713b78a564798f054a7abe> /System/Library/PrivateFrameworks/AppSupport.framework/AppSupport
0x18bdb8000 - 0x18be14fff  AppleAccount arm64  <8277c69507be311d84bbfe7f0d4a176a> /System/Library/PrivateFrameworks/AppleAccount.framework/AppleAccount
0x18bf1c000 - 0x18bf36fff  AppleIDSSOAuthentication arm64  <0b4ed9a507cb3e2d8de1e2e751b1b86c> /System/Library/PrivateFrameworks/AppleIDSSOAuthentication.framework/AppleIDSSOAuthentication
0x18bf38000 - 0x18bf79fff  AppleJPEG arm64  <dbd19fb3cc493bb3850525a1c22bdb20> /System/Library/PrivateFrameworks/AppleJPEG.framework/AppleJPEG
0x18bf90000 - 0x18bfa4fff  ApplePushService arm64  <e95e3cc88c8c31ad8fcfd61dd1658b92> /System/Library/PrivateFrameworks/ApplePushService.framework/ApplePushService
0x18bfe8000 - 0x18bff3fff  AssertionServices arm64  <e3128601dc283abea23640adc0771aae> /System/Library/PrivateFrameworks/AssertionServices.framework/AssertionServices
0x18bff4000 - 0x18bff9fff  AssetCacheServices arm64  <be0d66d1fe02329f8d525b6471360d75> /System/Library/PrivateFrameworks/AssetCacheServices.framework/AssetCacheServices
0x18bffc000 - 0x18c01bfff  AssetsLibraryServices arm64  <82ac976666913fbfba7e0e2526588a17> /System/Library/PrivateFrameworks/AssetsLibraryServices.framework/AssetsLibraryServices
0x18c0ac000 - 0x18c0cffff  AuthKit arm64  <7fcde3970bd4339aa7f2e09782e23f06> /System/Library/PrivateFrameworks/AuthKit.framework/AuthKit
0x18c0ec000 - 0x18c0edfff  BTLEAudioController arm64  <30ebc78105fe3c8d909502df1676710a> /System/Library/PrivateFrameworks/BTLEAudioController.framework/BTLEAudioController
0x18c0f0000 - 0x18c111fff  BackBoardServices arm64  <00ee611f1aa3377b8b4c480529b75373> /System/Library/PrivateFrameworks/BackBoardServices.framework/BackBoardServices
0x18c118000 - 0x18c160fff  BaseBoard arm64  <1077fc60c413310fb01b767da20cdeb0> /System/Library/PrivateFrameworks/BaseBoard.framework/BaseBoard
0x18c164000 - 0x18c166fff  BaseBoardUI arm64  <e7b65b61797836b292529a34ad276d01> /System/Library/PrivateFrameworks/BaseBoardUI.framework/BaseBoardUI
0x18c194000 - 0x18c1c4fff  Bom arm64  <2350e797709f3803951c1201731bc263> /System/Library/PrivateFrameworks/Bom.framework/Bom
0x18c1e8000 - 0x18c266fff  BulletinBoard arm64  <84a13c31d77337ceaa93e016abef09ee> /System/Library/PrivateFrameworks/BulletinBoard.framework/BulletinBoard
0x18c2d0000 - 0x18c2dffff  CacheDelete arm64  <2e18dda7cba63946a3d7f44f2bfd5ff8> /System/Library/PrivateFrameworks/CacheDelete.framework/CacheDelete
0x18c348000 - 0x18c379fff  CalendarDaemon arm64  <7d47f46a065e3fdaaaa124c41f6652f7> /System/Library/PrivateFrameworks/CalendarDaemon.framework/CalendarDaemon
0x18c37c000 - 0x18c3fdfff  CalendarDatabase arm64  <b632712ae4c13d1593ad8840d1123da0> /System/Library/PrivateFrameworks/CalendarDatabase.framework/CalendarDatabase
0x18c400000 - 0x18c44cfff  CalendarFoundation arm64  <0d3a5bca2a3839d0b46da47078d7403d> /System/Library/PrivateFrameworks/CalendarFoundation.framework/CalendarFoundation
0x18c4a4000 - 0x18c58cfff  CameraKit arm64  <7b11b65de13e34c484b0e62bbebfc873> /System/Library/PrivateFrameworks/CameraKit.framework/CameraKit
0x18c6ec000 - 0x18c6f6fff  CaptiveNetwork arm64  <62bf35338c1838f7a2f5bbe498d5aeda> /System/Library/PrivateFrameworks/CaptiveNetwork.framework/CaptiveNetwork
0x18c704000 - 0x18c883fff  Celestial arm64  <1438fd17e7f231358dc2dbae563b60ae> /System/Library/PrivateFrameworks/Celestial.framework/Celestial
0x18c8b4000 - 0x18c8bafff  CertUI arm64  <358e2269f61f309f9d7e60570cdf1d4f> /System/Library/PrivateFrameworks/CertUI.framework/CertUI
0x18ca20000 - 0x18ca44fff  ChunkingLibrary arm64  <26b6f9f5eaf83c21b772e85b6811b407> /System/Library/PrivateFrameworks/ChunkingLibrary.framework/ChunkingLibrary
0x18cf38000 - 0x18d00bfff  CloudPhotoLibrary arm64  <a51b5830756f3fa5b92370b283ed4018> /System/Library/PrivateFrameworks/CloudPhotoLibrary.framework/CloudPhotoLibrary
0x18d01c000 - 0x18d074fff  ColorSync arm64  <536b7a0de1c93bc1814d5e3c72248765> /System/Library/PrivateFrameworks/ColorSync.framework/ColorSync
0x18d078000 - 0x18d07afff  CommonAuth arm64  <bce5a6980e1f38429ae4db24331c600d> /System/Library/PrivateFrameworks/CommonAuth.framework/CommonAuth
0x18d07c000 - 0x18d08efff  CommonUtilities arm64  <657e5636a6a33af594cedf30a56b90ad> /System/Library/PrivateFrameworks/CommonUtilities.framework/CommonUtilities
0x18d090000 - 0x18d094fff  CommunicationsFilter arm64  <e77a88ae5e583928a36f9c7518d83b39> /System/Library/PrivateFrameworks/CommunicationsFilter.framework/CommunicationsFilter
0x18d1d8000 - 0x18d1ddfff  ConstantClasses arm64  <02ec6c1746673bee89dfa346cb2c0178> /System/Library/PrivateFrameworks/ConstantClasses.framework/ConstantClasses
0x18d1e0000 - 0x18d214fff  ContactsAutocomplete arm64  <e8d04afb04dc330d82e801cb796751e1> /System/Library/PrivateFrameworks/ContactsAutocomplete.framework/ContactsAutocomplete
0x18d218000 - 0x18d240fff  ContactsFoundation arm64  <1cb3f905bba73a33867c99d7a6648109> /System/Library/PrivateFrameworks/ContactsFoundation.framework/ContactsFoundation
0x18d244000 - 0x18d285fff  ContentIndex arm64  <a45e1d7d37e13338b99a7750e9ee1753> /System/Library/PrivateFrameworks/ContentIndex.framework/ContentIndex
0x18d288000 - 0x18d28ffff  CoreAUC arm64  <86b964431d933dffafaf641bd308a200> /System/Library/PrivateFrameworks/CoreAUC.framework/CoreAUC
0x18d324000 - 0x18d37ffff  CoreDAV arm64  <4005dd37faf13f708078812254e86087> /System/Library/PrivateFrameworks/CoreDAV.framework/CoreDAV
0x18d380000 - 0x18d432fff  CoreDuet arm64  <1304cc1bd4e63fe69dccc7427e59721d> /System/Library/PrivateFrameworks/CoreDuet.framework/CoreDuet
0x18d434000 - 0x18d449fff  CoreDuetDaemonProtocol arm64  <39ab9e176e923e15b80c312f75028189> /System/Library/PrivateFrameworks/CoreDuetDaemonProtocol.framework/CoreDuetDaemonProtocol
0x18d450000 - 0x18d452fff  CoreDuetDebugLogging arm64  <9c878b0038e337788488f19dae735f45> /System/Library/PrivateFrameworks/CoreDuetDebugLogging.framework/CoreDuetDebugLogging
0x18d630000 - 0x18d75cfff  CoreMediaStream arm64  <39f1a53720233fbf9bacb9cc35d7e684> /System/Library/PrivateFrameworks/CoreMediaStream.framework/CoreMediaStream
0x18d808000 - 0x18d8b4fff  CorePDF arm64  <c6f9ca0157cb3f5d95156d01d96981c7> /System/Library/PrivateFrameworks/CorePDF.framework/CorePDF
0x18d96c000 - 0x18d976fff  CoreRecents arm64  <6b06a0f74df23d8f8ae09644d880b46e> /System/Library/PrivateFrameworks/CoreRecents.framework/CoreRecents
0x18da18000 - 0x18da3dfff  CoreServicesInternal arm64  <c32cef197d80355fa40ab617f1cb6a21> /System/Library/PrivateFrameworks/CoreServicesInternal.framework/CoreServicesInternal
0x18dc98000 - 0x18dc9ffff  CoreTime arm64  <15a0338b606c3ecd839d90c21a8787a1> /System/Library/PrivateFrameworks/CoreTime.framework/CoreTime
0x18dca0000 - 0x18dd57fff  CoreUI arm64  <3d4a9c364fa53e568570a3bb0527d35d> /System/Library/PrivateFrameworks/CoreUI.framework/CoreUI
0x18ddf8000 - 0x18de07fff  CrashReporterSupport arm64  <38c3907ec236378c9274d5cc695758c5> /System/Library/PrivateFrameworks/CrashReporterSupport.framework/CrashReporterSupport
0x18de24000 - 0x18de2afff  DAAPKit arm64  <3faf108a39d7300bacc4ba94afdc1ec8> /System/Library/PrivateFrameworks/DAAPKit.framework/DAAPKit
0x18de2c000 - 0x18de37fff  DCIMServices arm64  <11056ded77643ed2af081ed50c9f18bf> /System/Library/PrivateFrameworks/DCIMServices.framework/DCIMServices
0x18de38000 - 0x18de86fff  DataAccess arm64  <c0000603c6363bf8b723290edd85c467> /System/Library/PrivateFrameworks/DataAccess.framework/DataAccess
0x18e0e0000 - 0x18e105fff  DataAccessExpress arm64  <fb2f3d3f812839cd881a07ab3ebf1266> /System/Library/PrivateFrameworks/DataAccessExpress.framework/DataAccessExpress
0x18e188000 - 0x18e18ffff  DataMigration arm64  <2421fe9683ce3feeb52d2a5fa5f37017> /System/Library/PrivateFrameworks/DataMigration.framework/DataMigration
0x18e1a4000 - 0x18e1a5fff  DiagnosticLogCollection arm64  <cb19948e2c5e3ee68b8c5e94ae627865> /System/Library/PrivateFrameworks/DiagnosticLogCollection.framework/DiagnosticLogCollection
0x18e1a8000 - 0x18e1cdfff  DictionaryServices arm64  <56e01196bffd33d0a65e9fc5f65d2113> /System/Library/PrivateFrameworks/DictionaryServices.framework/DictionaryServices
0x18e230000 - 0x18e256fff  EAP8021X arm64  <5bfc0fd6474a371cbe1bcb655abb5d52> /System/Library/PrivateFrameworks/EAP8021X.framework/EAP8021X
0x18e294000 - 0x18e29cfff  FMCoreLite arm64  <b274507c0b263b7299b68d026ac45478> /System/Library/PrivateFrameworks/FMCoreLite.framework/FMCoreLite
0x18e324000 - 0x18e326fff  FTClientServices arm64  <4d2edfedef7c3701addce30df4aba1ee> /System/Library/PrivateFrameworks/FTClientServices.framework/FTClientServices
0x18e328000 - 0x18e35cfff  FTServices arm64  <87e7cd10e9a8396ca1fab5e8d978aea5> /System/Library/PrivateFrameworks/FTServices.framework/FTServices
0x18e360000 - 0x18e78cfff  FaceCore arm64  <0e1fceb16e743246abf1db4f3a55486e> /System/Library/PrivateFrameworks/FaceCore.framework/FaceCore
0x18e834000 - 0x18e834fff  FontServices arm64  <a6e696a4bda73bd4bfa65967cec1ef7a> /System/Library/PrivateFrameworks/FontServices.framework/FontServices
0x18e838000 - 0x18e918fff  libFontParser.dylib arm64  <d81bb3bffbd330fea27999ccbf9a804b> /System/Library/PrivateFrameworks/FontServices.framework/libFontParser.dylib
0x18e91c000 - 0x18e927fff  libGSFontCache.dylib arm64  <f566ea2f927e38b4b40856bb1ce21efa> /System/Library/PrivateFrameworks/FontServices.framework/libGSFontCache.dylib
0x18ea54000 - 0x18ea89fff  FrontBoardServices arm64  <56b81d3b3f943e7792fa68eb22985626> /System/Library/PrivateFrameworks/FrontBoardServices.framework/FrontBoardServices
0x18ec80000 - 0x18ec80fff  libmetal_timestamp.dylib arm64  <1ba5e8e5c3dd3877b4ff5a174d5c0a52> /System/Library/PrivateFrameworks/GPUCompiler.framework/libmetal_timestamp.dylib
0x18f3fc000 - 0x18f405fff  libGPUSupportMercury.dylib arm64  <eab387d6cb1c34df9038b25c4d9e63dd> /System/Library/PrivateFrameworks/GPUSupport.framework/libGPUSupportMercury.dylib
0x18f8f4000 - 0x18f90cfff  GenerationalStorage arm64  <594aac3acac93869972dc324fca28846> /System/Library/PrivateFrameworks/GenerationalStorage.framework/GenerationalStorage
0x18f910000 - 0x18fd6dfff  GeoServices arm64  <108c937f80493cfe8f66ac90bd85301c> /System/Library/PrivateFrameworks/GeoServices.framework/GeoServices
0x18fd70000 - 0x18fd84fff  GraphicsServices arm64  <3d44f5bf11b43713a8917376fe2fafab> /System/Library/PrivateFrameworks/GraphicsServices.framework/GraphicsServices
0x18fd9c000 - 0x18fd9cfff  HangTracer arm64  <6b11850315d1371da3e38b4f96181b3d> /System/Library/PrivateFrameworks/HangTracer.framework/HangTracer
0x190018000 - 0x190085fff  Heimdal arm64  <66f66f994da93424a44226dd6a19b39c> /System/Library/PrivateFrameworks/Heimdal.framework/Heimdal
0x1902a0000 - 0x19033cfff  HomeSharing arm64  <a0a68fb4b60339d9a3efe48794914d25> /System/Library/PrivateFrameworks/HomeSharing.framework/HomeSharing
0x1903ac000 - 0x19041bfff  IDS arm64  <e4d6fbfcc2d93df1b3fd2642192c739e> /System/Library/PrivateFrameworks/IDS.framework/IDS
0x19041c000 - 0x19044ffff  IDSFoundation arm64  <ee40963e1c973caea107a8dc976c5b9d> /System/Library/PrivateFrameworks/IDSFoundation.framework/IDSFoundation
0x19064c000 - 0x1906bbfff  IMFoundation arm64  <e36f4b9415a3387f8871f2ed8decd6a5> /System/Library/PrivateFrameworks/IMFoundation.framework/IMFoundation
0x1906c4000 - 0x1906c8fff  IOAccelerator arm64  <c735d9fc20763bfcbc73da56aa920011> /System/Library/PrivateFrameworks/IOAccelerator.framework/IOAccelerator
0x1906d0000 - 0x1906d6fff  IOMobileFramebuffer arm64  <e8239f57f4543dca826af569cda5610d> /System/Library/PrivateFrameworks/IOMobileFramebuffer.framework/IOMobileFramebuffer
0x1906d8000 - 0x1906defff  IOSurface arm64  <95b386632f0a3a738204e82c58777d8d> /System/Library/PrivateFrameworks/IOSurface.framework/IOSurface
0x1906e0000 - 0x1906e2fff  IOSurfaceAccelerator arm64  <8405accb54a434c4a53a06a428dfbde8> /System/Library/PrivateFrameworks/IOSurfaceAccelerator.framework/IOSurfaceAccelerator
0x190ac4000 - 0x190acafff  IncomingCallFilter arm64  <8890f055ecdf319382165c8d3ce46779> /System/Library/PrivateFrameworks/IncomingCallFilter.framework/IncomingCallFilter
0x190ad0000 - 0x190adcfff  IntlPreferences arm64  <976063bcf8bd33ec9273048b7fc353a6> /System/Library/PrivateFrameworks/IntlPreferences.framework/IntlPreferences
0x190b10000 - 0x190b91fff  LanguageModeling arm64  <d16f52bf27ac371786ab84b349a6a54e> /System/Library/PrivateFrameworks/LanguageModeling.framework/LanguageModeling
0x190b94000 - 0x190bb5fff  LatentSemanticMapping arm64  <6ba9f6e6f28631968b196028cd7c07a5> /System/Library/PrivateFrameworks/LatentSemanticMapping.framework/LatentSemanticMapping
0x190bc8000 - 0x190c0efff  MIME arm64  <20c9175ed7193b23805e65a56c39d66a> /System/Library/PrivateFrameworks/MIME.framework/MIME
0x190c10000 - 0x190c85fff  MMCS arm64  <930285d52a343db39ef3f5fd2b58c279> /System/Library/PrivateFrameworks/MMCS.framework/MMCS
0x190d10000 - 0x190d1ffff  MailServices arm64  <d3237ecb8ec032729a6362507a939bdf> /System/Library/PrivateFrameworks/MailServices.framework/MailServices
0x190d70000 - 0x190e2ffff  ManagedConfiguration arm64  <ebb03b31fb0c38568824a352747ea945> /System/Library/PrivateFrameworks/ManagedConfiguration.framework/ManagedConfiguration
0x190e50000 - 0x190e51fff  Marco arm64  <8b4b6248dee133ccb3159cc218a37abf> /System/Library/PrivateFrameworks/Marco.framework/Marco
0x191348000 - 0x191369fff  MediaRemote arm64  <63623e17af913ebfa42299cd97b94e11> /System/Library/PrivateFrameworks/MediaRemote.framework/MediaRemote
0x19136c000 - 0x191386fff  MediaServices arm64  <9099e11181d13565b04b3ae714e4f1f3> /System/Library/PrivateFrameworks/MediaServices.framework/MediaServices
0x1913ec000 - 0x191407fff  MediaStream arm64  <b92a35de747d3456a74b6dea7098ddc4> /System/Library/PrivateFrameworks/MediaStream.framework/MediaStream
0x191478000 - 0x191595fff  Message arm64  <b0066d552bfa39e7bd2bd0aafe5a6854> /System/Library/PrivateFrameworks/Message.framework/Message
0x1915a0000 - 0x1915a3fff  MessageSupport arm64  <8869abbe17d8324f9f4ed17b880b6756> /System/Library/PrivateFrameworks/MessageSupport.framework/MessageSupport
0x19161c000 - 0x19162bfff  MobileAsset arm64  <443ea8a79c8434758da8cfc2b0a3dc6c> /System/Library/PrivateFrameworks/MobileAsset.framework/MobileAsset
0x191694000 - 0x19169cfff  MobileInstallation arm64  <ef3d2f0aa11a34f4ab3f69c85e0f32b5> /System/Library/PrivateFrameworks/MobileInstallation.framework/MobileInstallation
0x1916a0000 - 0x1916affff  MobileKeyBag arm64  <abc5481725ec3750a66150c84028740a> /System/Library/PrivateFrameworks/MobileKeyBag.framework/MobileKeyBag
0x1916b8000 - 0x1916b8fff  MobileObliteration arm64  <5f5b5d439233383cb0400c157b1154f9> /System/Library/PrivateFrameworks/MobileObliteration.framework/MobileObliteration
0x1916c4000 - 0x1917f0fff  MobileSpotlightIndex arm64  <9386f24885a53b84ba16f0055a4d0662> /System/Library/PrivateFrameworks/MobileSpotlightIndex.framework/MobileSpotlightIndex
0x191824000 - 0x191827fff  MobileSystemServices arm64  <5c929aaa3b9b3baa9b40cb3bde3fabc6> /System/Library/PrivateFrameworks/MobileSystemServices.framework/MobileSystemServices
0x191858000 - 0x19186afff  MobileWiFi arm64  <0fcf3d4ee9f0363ebbb842cd46177923> /System/Library/PrivateFrameworks/MobileWiFi.framework/MobileWiFi
0x1918d0000 - 0x191b19fff  MusicLibrary arm64  <db75eeb7b14436e2a978cb8c73e16a05> /System/Library/PrivateFrameworks/MusicLibrary.framework/MusicLibrary
0x191ce4000 - 0x191ceafff  Netrb arm64  <e65c49ffac19355ca6b4dff785c3c25d> /System/Library/PrivateFrameworks/Netrb.framework/Netrb
0x191cec000 - 0x191d0ffff  Network arm64  <991e2a8f01d430c591ef5ef35f742d79> /System/Library/PrivateFrameworks/Network.framework/Network
0x191d10000 - 0x191d29fff  NetworkStatistics arm64  <a5f61c0bdcfd33a48d1af176e8f5b18e> /System/Library/PrivateFrameworks/NetworkStatistics.framework/NetworkStatistics
0x191d2c000 - 0x191d71fff  Notes arm64  <f5e839eb8d683ef8b9983870c87b86c1> /System/Library/PrivateFrameworks/Notes.framework/Notes
0x191ed8000 - 0x191ee3fff  NotificationsUI arm64  <f4b300ef13ed3f7d85a801c7b3da1c87> /System/Library/PrivateFrameworks/NotificationsUI.framework/NotificationsUI
0x191ee4000 - 0x191ee6fff  OAuth arm64  <db5840bdb4d53c68a23864f02543b949> /System/Library/PrivateFrameworks/OAuth.framework/OAuth
0x192848000 - 0x19284cfff  ParsecSubscriptionServiceSupport arm64  <7e2ba213e4fa3d89942d6b8cac721a41> /System/Library/PrivateFrameworks/ParsecSubscriptionServiceSupport.framework/ParsecSubscriptionServiceSupport
0x192918000 - 0x19293ffff  Pegasus arm64  <23c37a7f431f3500b6755496976e8e5c> /System/Library/PrivateFrameworks/Pegasus.framework/Pegasus
0x192984000 - 0x1929abfff  PersistentConnection arm64  <1fbe98fa40d73b068f637b2423b42547> /System/Library/PrivateFrameworks/PersistentConnection.framework/PersistentConnection
0x192b08000 - 0x192e45fff  PhotoLibraryServices arm64  <8efb83921ee430fbae7fb5e50acd4086> /System/Library/PrivateFrameworks/PhotoLibraryServices.framework/PhotoLibraryServices
0x192e48000 - 0x192e54fff  PhotosFormats arm64  <cc22c2c568c13dcaa1fb2d3d00c9eee7> /System/Library/PrivateFrameworks/PhotosFormats.framework/PhotosFormats
0x192ea0000 - 0x192eeefff  PhysicsKit arm64  <cddf26fc97c93796b39e8bd595eb3328> /System/Library/PrivateFrameworks/PhysicsKit.framework/PhysicsKit
0x192ef0000 - 0x192f0bfff  PlugInKit arm64  <f7e0a7b1866e32e6b0c32dfb9d449644> /System/Library/PrivateFrameworks/PlugInKit.framework/PlugInKit
0x192f0c000 - 0x192f19fff  PowerLog arm64  <a11b5db28ff13dec84eeb2c4ac12ae7d> /System/Library/PrivateFrameworks/PowerLog.framework/PowerLog
0x1934c4000 - 0x1934c8fff  ProgressUI arm64  <f397834e4c73333c9aff2941bd4a3087> /System/Library/PrivateFrameworks/ProgressUI.framework/ProgressUI
0x1934cc000 - 0x193593fff  ProofReader arm64  <acb94fd9b3cb33f88fbdf842dd112e23> /System/Library/PrivateFrameworks/ProofReader.framework/ProofReader
0x193594000 - 0x1935bafff  ProtectedCloudStorage arm64  <195dbd0ef29930d187f0dcc43c0aac6b> /System/Library/PrivateFrameworks/ProtectedCloudStorage.framework/ProtectedCloudStorage
0x1935bc000 - 0x1935cdfff  ProtocolBuffer arm64  <935437d4808b3429bd35249d42b6acb4> /System/Library/PrivateFrameworks/ProtocolBuffer.framework/ProtocolBuffer
0x1935d0000 - 0x193606fff  PrototypeTools arm64  <7d7de2070bb337c4a861037d3b44a543> /System/Library/PrivateFrameworks/PrototypeTools.framework/PrototypeTools
0x19360c000 - 0x193687fff  Quagga arm64  <e273664815173f6881fbbdb570ca4463> /System/Library/PrivateFrameworks/Quagga.framework/Quagga
0x193688000 - 0x193695fff  QuickLookThumbnailing arm64  <039fbfed2df033b4b60e32e5d1f87953> /System/Library/PrivateFrameworks/QuickLookThumbnailing.framework/QuickLookThumbnailing
0x19400c000 - 0x194043fff  SpringBoardFoundation arm64  <8b1ecc3b35a33dbcb9e8d1398b591f69> /System/Library/PrivateFrameworks/SpringBoardFoundation.framework/SpringBoardFoundation
0x194044000 - 0x194069fff  SpringBoardServices arm64  <3f550de7bba534b2a4a64b20da1ffbaa> /System/Library/PrivateFrameworks/SpringBoardServices.framework/SpringBoardServices
0x19406c000 - 0x194083fff  SpringBoardUI arm64  <7087ad2d30103dc68956211bcce995c7> /System/Library/PrivateFrameworks/SpringBoardUI.framework/SpringBoardUI
0x194084000 - 0x1940aafff  SpringBoardUIServices arm64  <de3be9d3f6a83bd39269b00eb9530dde> /System/Library/PrivateFrameworks/SpringBoardUIServices.framework/SpringBoardUIServices
0x194504000 - 0x1946eafff  StoreServices arm64  <51ea660f684930109164f4304d853af3> /System/Library/PrivateFrameworks/StoreServices.framework/StoreServices
0x19496c000 - 0x194970fff  TCC arm64  <fd45083d4b3a3ab497b354972aba1e71> /System/Library/PrivateFrameworks/TCC.framework/TCC
0x194984000 - 0x1949c9fff  TelephonyUI arm64  <76a1b6c1763638c38a1b03555875ec4a> /System/Library/PrivateFrameworks/TelephonyUI.framework/TelephonyUI
0x1949cc000 - 0x194a2ffff  TelephonyUtilities arm64  <6b91804c626e3c1c9f17f6ad0c1043a1> /System/Library/PrivateFrameworks/TelephonyUtilities.framework/TelephonyUtilities
0x195a0c000 - 0x195a3afff  TextInput arm64  <5d345ffb8a5c3ff495392354b7d5b53c> /System/Library/PrivateFrameworks/TextInput.framework/TextInput
0x195a94000 - 0x195ab4fff  ToneLibrary arm64  <675ca248d40c38d99bcc3bffd1d35d93> /System/Library/PrivateFrameworks/ToneLibrary.framework/ToneLibrary
0x195b38000 - 0x195c0efff  UIFoundation arm64  <e3f9f4859dd3329aafec5954d4a9f7a6> /System/Library/PrivateFrameworks/UIFoundation.framework/UIFoundation
0x195c48000 - 0x195c4bfff  UserFS arm64  <ffbf7d0d39f93d88a5b6577e46d4ea95> /System/Library/PrivateFrameworks/UserFS.framework/UserFS
0x195c74000 - 0x195c7efff  UserNotificationServices arm64  <549d8037fd1932baaf98b1278bb68150> /System/Library/PrivateFrameworks/UserNotificationServices.framework/UserNotificationServices
0x195ca0000 - 0x1962f4fff  VectorKit arm64  <504ab976b3513f9d8fc978aad072de1b> /System/Library/PrivateFrameworks/VectorKit.framework/VectorKit
0x196664000 - 0x196699fff  WebBookmarks arm64  <82b229f857cb3fd194c9752d5b565a5e> /System/Library/PrivateFrameworks/WebBookmarks.framework/WebBookmarks
0x19669c000 - 0x1966b2fff  WebContentAnalysis arm64  <0d5eceb067be3e70b1b9d9f2120004dc> /System/Library/PrivateFrameworks/WebContentAnalysis.framework/WebContentAnalysis
0x1966b4000 - 0x19765bfff  WebCore arm64  <3da99f2e6511384e9f08fdbf3b3c787f> /System/Library/PrivateFrameworks/WebCore.framework/WebCore
0x1976c4000 - 0x1977a1fff  WebKitLegacy arm64  <8da5b6e24d283db3825cbb6bf467f1bc> /System/Library/PrivateFrameworks/WebKitLegacy.framework/WebKitLegacy
0x19789c000 - 0x1978a3fff  XPCKit arm64  <f091573181193bc09bf380529e62e7c7> /System/Library/PrivateFrameworks/XPCKit.framework/XPCKit
0x197b00000 - 0x197b31fff  iCalendar arm64  <48229e07ccc53012b49439071dea9434> /System/Library/PrivateFrameworks/iCalendar.framework/iCalendar
0x197b60000 - 0x197ba3fff  iTunesStore arm64  <0a42c2518c913cfba2692473882a73a9> /System/Library/PrivateFrameworks/iTunesStore.framework/iTunesStore
0x197d7c000 - 0x197d9bfff  vCard arm64  <0a45ec8021753b1a8354c3fa82193d6c> /System/Library/PrivateFrameworks/vCard.framework/vCard
0x1984f8000 - 0x198502fff  libAccessibility.dylib arm64  <dd729d022e8b3e308fdad21285afa572> /usr/lib/libAccessibility.dylib
0x19882c000 - 0x198846fff  libCRFSuite.dylib arm64  <f507a48362643e589123ea19027d3790> /usr/lib/libCRFSuite.dylib
0x198854000 - 0x19885ffff  libChineseTokenizer.dylib arm64  <50ca891f1a093ee598ab3e27ea8f2354> /usr/lib/libChineseTokenizer.dylib
0x198afc000 - 0x198b1bfff  libMobileGestalt.dylib arm64  <f0b24095dbf53e2aa936136599e23796> /usr/lib/libMobileGestalt.dylib
0x198b50000 - 0x198b51fff  libSystem.B.dylib arm64  <63540f431e023138b9e2509de35049c7> /usr/lib/libSystem.B.dylib
0x198be8000 - 0x198c45fff  libTelephonyUtilDynamic.dylib arm64  <7462740af41d3a4e91f254bba9cd8d7c> /usr/lib/libTelephonyUtilDynamic.dylib
0x198cc8000 - 0x198cf2fff  libarchive.2.dylib arm64  <22a14b12726d3004a2e63d309142126f> /usr/lib/libarchive.2.dylib
0x198d30000 - 0x198d40fff  libbsm.0.dylib arm64  <8a78e1a8ac803453b860da3bd66f0389> /usr/lib/libbsm.0.dylib
0x198d44000 - 0x198d51fff  libbz2.1.0.dylib arm64  <acd2d5f91e4530459212f6d3dcf0fc0a> /usr/lib/libbz2.1.0.dylib
0x198d54000 - 0x198da7fff  libc++.1.dylib arm64  <87ad49d1d16936358965927dc98000bf> /usr/lib/libc++.1.dylib
0x198da8000 - 0x198dc7fff  libc++abi.dylib arm64  <434eb18a8b293ac1a443def910a6bcd8> /usr/lib/libc++abi.dylib
0x198dcc000 - 0x198dddfff  libcmph.dylib arm64  <3cb53ff418033b56a65f3c147cc00046> /usr/lib/libcmph.dylib
0x198de0000 - 0x198df7fff  libcompression.dylib arm64  <c16ca5f082b13cdfa9ab6cbd1320cb41> /usr/lib/libcompression.dylib
0x198f24000 - 0x198f2cfff  libcupolicy.dylib arm64  <8d8235e8d34339d69cb2e9bb2f526ca4> /usr/lib/libcupolicy.dylib
0x198f54000 - 0x198f54fff  libenergytrace.dylib arm64  <661a93b7c7d73228b027334b91a3f13e> /usr/lib/libenergytrace.dylib
0x198f68000 - 0x198f86fff  libextension.dylib arm64  <2304b2a80c683e8b855aae5864779504> /usr/lib/libextension.dylib
0x198fb4000 - 0x198fb9fff  libheimdal-asn1.dylib arm64  <808afe7ade97315eb146e6f2736bd956> /usr/lib/libheimdal-asn1.dylib
0x198fbc000 - 0x1990aefff  libiconv.2.dylib arm64  <4502eec3f81a39839a62939338a43826> /usr/lib/libiconv.2.dylib
0x1990b0000 - 0x1992b5fff  libicucore.A.dylib arm64  <fd01421bf4403da595b9f542e85c64d8> /usr/lib/libicucore.A.dylib
0x1992c4000 - 0x1992c5fff  liblangid.dylib arm64  <deddf5b59914302d86fe971e0d672e30> /usr/lib/liblangid.dylib
0x1992c8000 - 0x1992d4fff  liblockdown.dylib arm64  <184cbed1a307341ea2044f4a07baf10f> /usr/lib/liblockdown.dylib
0x1992d8000 - 0x1992f1fff  liblzma.5.dylib arm64  <ed4783db6d2f3f8e8cda2fd54cfc4d69> /usr/lib/liblzma.5.dylib
0x1992f4000 - 0x19930bfff  libmarisa.dylib arm64  <6f44ff9c5945354f8693c0c10cecf7ce> /usr/lib/libmarisa.dylib
0x199444000 - 0x1996bdfff  libmecabra.dylib arm64  <019b92a4b1db312b8d265d2f76078a9f> /usr/lib/libmecabra.dylib
0x1996c0000 - 0x1996dafff  libmis.dylib arm64  <48cc61b0181932ccad34166117137204> /usr/lib/libmis.dylib
0x199710000 - 0x199717fff  libnetwork.dylib arm64  <e8c127f270f236439fc45aeb0cfcd420> /usr/lib/libnetwork.dylib
0x199718000 - 0x199a84fff  libobjc.A.dylib arm64  <c1da9181c3ff3ce48b28fc0d83601d17> /usr/lib/libobjc.A.dylib
0x199bb4000 - 0x199bcffff  libresolv.9.dylib arm64  <543f10bcbd6e335fa8c9b27343b9ed26> /usr/lib/libresolv.9.dylib
0x199c04000 - 0x199cecfff  libsqlite3.dylib arm64  <45f2742f39393b2ca3320766fa68f219> /usr/lib/libsqlite3.dylib
0x199cf0000 - 0x199d41fff  libstdc++.6.dylib arm64  <3bde6c999aeb3517bd9bb2014f982b63> /usr/lib/libstdc++.6.dylib
0x199d44000 - 0x199d75fff  libtidy.A.dylib arm64  <e8c57f4ed0b9330f80f3768577de9786> /usr/lib/libtidy.A.dylib
0x199d88000 - 0x199e72fff  libxml2.2.dylib arm64  <9cc257d7bd4f3496b4bf8db52a07d838> /usr/lib/libxml2.2.dylib
0x199e74000 - 0x199e9dfff  libxslt.1.dylib arm64  <d4b2385d674b3b03b408e7fc37ffc346> /usr/lib/libxslt.1.dylib
0x199ea0000 - 0x199eb1fff  libz.1.dylib arm64  <546b170fa0203db79dceef6d992503c9> /usr/lib/libz.1.dylib
0x199eb4000 - 0x199eb8fff  libcache.dylib arm64  <05927d2648d93593a48c6cd9c92c1085> /usr/lib/system/libcache.dylib
0x199ebc000 - 0x199ec7fff  libcommonCrypto.dylib arm64  <ea6de92f241831faadcda6c1805b09e5> /usr/lib/system/libcommonCrypto.dylib
0x199ec8000 - 0x199ecbfff  libcompiler_rt.dylib arm64  <19dc6f7f27d13adba2d40c063fe83269> /usr/lib/system/libcompiler_rt.dylib
0x199ed8000 - 0x199edffff  libcopyfile.dylib arm64  <b90fc57e270c3143a77b20bedebd4528> /usr/lib/system/libcopyfile.dylib
0x199ee0000 - 0x199f3bfff  libcorecrypto.dylib arm64  <207c7c665c5a3ceb9ee85df6b48fdcf2> /usr/lib/system/libcorecrypto.dylib
0x199f3c000 - 0x199f6afff  libdispatch.dylib arm64  <3522a9c2bffa3fb59fe97869e154f373> /usr/lib/system/libdispatch.dylib
0x199f6c000 - 0x199f6efff  libdyld.dylib arm64  <6e6e8e5fd98c3a03bcd0bb902757c758> /usr/lib/system/libdyld.dylib
0x199f70000 - 0x199f70fff  liblaunch.dylib arm64  <587dfce9b9e03bc298a44f448e8c2b87> /usr/lib/system/liblaunch.dylib
0x199f74000 - 0x199f78fff  libmacho.dylib arm64  <94f6326595d13d8ea79ccc3278d03e99> /usr/lib/system/libmacho.dylib
0x199f7c000 - 0x199f7dfff  libremovefile.dylib arm64  <0bc46632c176374ab87bfc8cc169eca3> /usr/lib/system/libremovefile.dylib
0x199f80000 - 0x199f96fff  libsystem_asl.dylib arm64  <6dc388aafcad3bbb93f73fb2b4b7b494> /usr/lib/system/libsystem_asl.dylib
0x199f98000 - 0x199f99fff  libsystem_blocks.dylib arm64  <ca4becaf737b36c8b0eb3eda80136fa9> /usr/lib/system/libsystem_blocks.dylib
0x199f9c000 - 0x19a01cfff  libsystem_c.dylib arm64  <3921694681f93149adaaa6b535b135a3> /usr/lib/system/libsystem_c.dylib
0x19a020000 - 0x19a023fff  libsystem_configuration.dylib arm64  <a8321ff060223f61abce4fa73fa78399> /usr/lib/system/libsystem_configuration.dylib
0x19a024000 - 0x19a026fff  libsystem_containermanager.dylib arm64  <e0bbf49de038316bbecaacca1d00465c> /usr/lib/system/libsystem_containermanager.dylib
0x19a028000 - 0x19a029fff  libsystem_coreservices.dylib arm64  <71ea857e1b363ea9a2cfff73b122cb42> /usr/lib/system/libsystem_coreservices.dylib
0x19a02c000 - 0x19a03ffff  libsystem_coretls.dylib arm64  <5ffec8b142d93f8285276ccfaf3d7ae9> /usr/lib/system/libsystem_coretls.dylib
0x19a040000 - 0x19a048fff  libsystem_dnssd.dylib arm64  <f8f1e9fe49573e9bb11a040d5056e22c> /usr/lib/system/libsystem_dnssd.dylib
0x19a04c000 - 0x19a06dfff  libsystem_info.dylib arm64  <a3b8468e730832a3853b02692e97de17> /usr/lib/system/libsystem_info.dylib
0x19a070000 - 0x19a090fff  libsystem_kernel.dylib arm64  <bdb28b7fc5f738dfa218ee207a70a7d6> /usr/lib/system/libsystem_kernel.dylib
0x19a094000 - 0x19a0b0fff  libsystem_m.dylib arm64  <b003183df8e233c09e6d9b00535cb382> /usr/lib/system/libsystem_m.dylib
0x19a0b4000 - 0x19a0cdfff  libsystem_malloc.dylib arm64  <dd15a9e51174390cba785a942975fcbb> /usr/lib/system/libsystem_malloc.dylib
0x19a0d0000 - 0x19a12cfff  libsystem_network.dylib arm64  <72998f05a7583613b92ddd1aba383615> /usr/lib/system/libsystem_network.dylib
0x19a130000 - 0x19a138fff  libsystem_networkextension.dylib arm64  <53444d80dd2c322c909919876378866c> /usr/lib/system/libsystem_networkextension.dylib
0x19a13c000 - 0x19a146fff  libsystem_notify.dylib arm64  <14b89d79d4b93ef29d22af53f0452149> /usr/lib/system/libsystem_notify.dylib
0x19a148000 - 0x19a14dfff  libsystem_platform.dylib arm64  <274d7c0b6cb03b5c8759dffb5eb957f6> /usr/lib/system/libsystem_platform.dylib
0x19a150000 - 0x19a158fff  libsystem_pthread.dylib arm64  <c2dc25304a05357191b5dac769fdc530> /usr/lib/system/libsystem_pthread.dylib
0x19a15c000 - 0x19a15efff  libsystem_sandbox.dylib arm64  <5bf90157ef1733749db7014d51a40c83> /usr/lib/system/libsystem_sandbox.dylib
0x19a160000 - 0x19a170fff  libsystem_trace.dylib arm64  <4a2000f055363e9ab20b356960640d52> /usr/lib/system/libsystem_trace.dylib
0x19a174000 - 0x19a179fff  libunwind.dylib arm64  <fe998822ac253ec98dbc9a8211331368> /usr/lib/system/libunwind.dylib
0x19a17c000 - 0x19a17cfff  libvminterpose.dylib arm64  <8b91407766be38159af29494baad3a71> /usr/lib/system/libvminterpose.dylib
0x19a180000 - 0x19a1a5fff  libxpc.dylib arm64  <a2e94a5dd00c34dfb8ad8479549a702a> /usr/lib/system/libxpc.dylib`
	cr, err := NewInitCrashReport(str)
	log.Println(err)
	assert.Nil(err)

	//assert.Equal(cr.UncaughtException, nil, "is not equal")
	assert.Equal(len(cr.Threads.Threads), 42, "is not equal")

	assert.Equal(cr.String(false), str, "is not equal!")

	uuidmaps := cr.UuidMaps()

	assert.Equal(uuidmaps["8b91407766be38159af29494baad3a71"], []int64(nil))
	assert.Equal(uuidmaps["5ffec8b142d93f8285276ccfaf3d7ae9"], []int64(nil))
	assert.Equal(len(uuidmaps["3f043d63e7503e9688fca40afa778931"]), 29)
	assert.Equal(cr.GetModuleName(), "KDaijiaDriver")

	assert.Nil(cr.ParseToString(2, "3f043d63e7503e9688fca40afa778930"))
	assert.False(cr.IsAllParsed)
	assert.False(cr.IsParsed)

	log.Println(cr.String(true))
	log.Println(cr.String(false))
}
