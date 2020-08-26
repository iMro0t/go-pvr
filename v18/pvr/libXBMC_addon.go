package pvr

/*
#include <stdlib.h>
#include "workaround.h"
#include "../kodi/libXBMC_addon.h"
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

type GoHelper_libXBMC_addon struct {
	DEBUG  C.int
	INFO   C.int
	NOTICE C.int
	ERROR  C.int
}

func NewlibXBMC_addon() GoHelper_libXBMC_addon {
	return GoHelper_libXBMC_addon{
		DEBUG:  C.LOG_DEBUG,
		INFO:   C.LOG_INFO,
		NOTICE: C.LOG_NOTICE,
		ERROR:  C.LOG_ERROR,
	}
}

func (xbmc GoHelper_libXBMC_addon) RegisterMe(hdl unsafe.Pointer) C.bool {
	return C.CGoHelper_libXBMC_addonRegisterMe(hdl)
}

func (xbmc GoHelper_libXBMC_addon) Log(loglevel C.int, format ...interface{}) {
	log.Print(fmt.Sprint(format...))
	C.CGoHelper_libXBMC_addonLog(loglevel, C.CString(fmt.Sprint(format...)))
}

func (xbmc GoHelper_libXBMC_addon) GetSettingString(name string) string {
	cName := C.CString(name)
	cValue := C.CString("")
	C.CGoHelper_libXBMC_addonGetSetting(cName, unsafe.Pointer(cValue))
	// C.free(unsafe.Pointer(cName))
	// defer C.free(unsafe.Pointer(cValue))
	return C.GoString(cValue)
}

func (xbmc GoHelper_libXBMC_addon) GetSettingBool(name string) bool {
	cName := C.CString(name)
	cValue := C.bool(false)
	C.CGoHelper_libXBMC_addonGetSetting(cName, unsafe.Pointer(&cValue))
	// C.free(unsafe.Pointer(cName))
	// defer C.free(unsafe.Pointer(&cValue))
	return bool(cValue)
}

func (xbmc GoHelper_libXBMC_addon) Free() {
	C.CGoHelper_libXBMC_addonFree()
}
