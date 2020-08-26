package pvr

/*
#include <stdlib.h>
#include "workaround.h"
#include "../kodi/libKODI_guilib.h"
*/
import "C"
import (
	"errors"
	"sync"
	"unsafe"
)

type GoHelper_libKODI_guilib struct {
	mutex sync.Mutex
}

// type guiProgress struct {
// 	window   C.CGoHelper_CAddonGUIWindow
// 	progress C.CGoHelper_CAddonGUIProgressControl
// 	gui      *GoHelper_libKODI_guilib
// }

func NewlibKODI_guilib() GoHelper_libKODI_guilib {
	return GoHelper_libKODI_guilib{}
}

func (gui GoHelper_libKODI_guilib) RegisterMe(hdl unsafe.Pointer) C.bool {
	return C.CGoHelper_libKODI_guilibRegisterMe(hdl)
}

func (gui GoHelper_libKODI_guilib) DialogYesNo(heading, body, noLabel, yesLabel string) (bool, error) {
	cHeading := C.CString(heading)
	cBody := C.CString(body)
	cNoLabel := C.CString(noLabel)
	cYesLabel := C.CString(yesLabel)
	cCanceled := C.bool(false)
	defer C.free(unsafe.Pointer(cHeading))
	defer C.free(unsafe.Pointer(cBody))
	defer C.free(unsafe.Pointer(cNoLabel))
	defer C.free(unsafe.Pointer(cYesLabel))
	r := bool(C.CGoHelper_libKODI_guilibDialog_YesNo_ShowAndGetInput(cHeading, cBody, &cCanceled, cNoLabel, cYesLabel))
	if cCanceled {
		return false, errors.New("Dialog canceled by user")
	}
	return r, nil
}

func (gui GoHelper_libKODI_guilib) DialogOk(heading, body string) {
	cHeading := C.CString(heading)
	cBody := C.CString(body)
	defer C.free(unsafe.Pointer(cHeading))
	defer C.free(unsafe.Pointer(cBody))
	C.CGoHelper_libKODI_guilibDialog_OK_ShowAndGetInput(cHeading, cBody)
}

func (gui GoHelper_libKODI_guilib) DialogKeyboard(heading string, hiddenInput bool) string {
	cHeading := C.CString(heading)
	cResp := (*C.char)(C.malloc(C.sizeof_char * 256))
	r := C.CGoHelper_libKODI_guilibDialog_Keyboard_ShowAndGetInput(cResp, C.uint(256), cHeading, C.bool(true), C.bool(hiddenInput))
	if !r {
		return ""
	}
	return C.GoString(cResp)
}

// func (gui GoHelper_libKODI_guilib) CreateProgress() *guiProgress {
// 	cXMLFile := C.CString("DialogConfirm.xml")
// 	cDefaultSkin := C.CString("")
// 	window := C.CGoHelper_libKODI_guilibWindow_create(cXMLFile, cDefaultSkin, C.bool(false), C.bool(true))
// 	progress := C.CGoHelper_libKODI_guilibControl_getProgress(gui.addon, unsafe.Pointer(window), C.int(48996))
// 	return &guiProgress{window, progress, &gui}
// }

// func (p guiProgress) Update(percentage float64) {
// 	C.CGoHelper_libKODI_guilibSetProgressPercentage(p.progress, C.float(percentage))
// }

// func (p guiProgress) Close() {
// 	C.CGoHelper_libKODI_guilibControl_releaseProgress(p.gui.addon, unsafe.Pointer(p.progress))
// 	C.CGoHelper_libKODI_guilibWindow_destroy(p.gui.addon, unsafe.Pointer(p.window))
// }

func (gui GoHelper_libKODI_guilib) Free() {
	C.CGoHelper_libKODI_guilibFree()
}
