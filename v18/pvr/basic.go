package pvr

/*
#include "../kodi/xbmc_pvr_dll.h"
*/
import "C"

//export OnSystemSleep
func OnSystemSleep() {}

//export OnSystemWake
func OnSystemWake() {}

//export OnPowerSavingActivated
func OnPowerSavingActivated() {}

//export OnPowerSavingDeactivated
func OnPowerSavingDeactivated() {}

//export GetAddonCapabilities
func GetAddonCapabilities(cap *C.PVR_ADDON_CAPABILITIES) C.PVR_ERROR {
	cap.bSupportsEPG = true
	cap.bSupportsTV = true
	cap.bSupportsRadio = true
	cap.bSupportsChannelGroups = true
	cap.bSupportsRecordings = true
	cap.bSupportsRecordingsUndelete = true
	cap.bSupportsTimers = true
	cap.bSupportsRecordingsRename = false
	cap.bSupportsRecordingsLifetimeChange = false
	cap.bSupportsDescrambleInfo = false
	return C.PVR_ERROR_NO_ERROR
}

//export GetBackendName
func GetBackendName() *C.cchar_t {
	return C.CString("JioTV PVR (botallen)")
}

//export GetBackendVersion
func GetBackendVersion() *C.cchar_t {
	return C.CString("0.0.1")
}

//export GetConnectionString
func GetConnectionString() *C.cchar_t {
	return C.CString("connected")
}

//export GetBackendHostname
func GetBackendHostname() *C.cchar_t {
	return C.CString("")
}

//export GetDriveSpace
func GetDriveSpace(total *C.longlong, used *C.longlong) C.PVR_ERROR {
	*total = 0
	*used = 0
	return C.PVR_ERROR_NO_ERROR
}

//export SignalStatus
func SignalStatus(signalStatus *C.PVR_SIGNAL_STATUS) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export CallMenuHook
func CallMenuHook(menuhook *C.cPVR_MENUHOOK_t, _ *C.cPVR_MENUHOOK_DATA_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

func StrToCCharArr(str string) [1024]C.char {
	dest := [1024]C.char{}
	for i, c := range []byte(str) {
		dest[i] = C.char(c)
	}
	// This is C, we need to terminate the string!
	dest[len(str)] = 0
	return dest
}

func StrToUint8Arr(str string) [1024]byte {
	dest := [1024]byte{}
	copy(dest[:], str)
	return dest
}