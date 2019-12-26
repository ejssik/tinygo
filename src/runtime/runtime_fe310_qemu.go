// +build fe310,qemu

package runtime

import (
	"runtime/volatile"
	"unsafe"
)

// Special memory-mapped device to exit tests, created by SiFive.
var testExit = (*volatile.Register32)(unsafe.Pointer(uintptr(0x100000)))

var timestamp timeUnit

func abort() {
	// Signal a successful exit.
	testExit.Set(0x5555)

	// Note: the SiFive test finisher does not seem to be implemented in QEMU
	// 4.2 (but is included in the SiFive build of QEMU). And while the above
	// write somehow results in exiting the process on Windows, that's not the
	// case on macOS. Therefore, make sure to halt the process with an endless
	// loop.
	for {
	}
}

func ticks() timeUnit {
	return timestamp
}

func sleepTicks(d timeUnit) {
	// Note: QEMU doesn't seem to support the RTC peripheral at the time of
	// writing so just simulate sleeping here.
	timestamp += d
}
