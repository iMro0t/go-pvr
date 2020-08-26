#include "../kodi/libXBMC_addon.h"
#include "../kodi/libKODI_guilib.h"
#include "../kodi/xbmc_pvr_dll.h"
#include "../kodi/libXBMC_pvr.h"

#ifdef __cplusplus
extern "C"
{
#endif
    bool CGoHelper_libXBMC_addonRegisterMe(void *handle);
    void CGoHelper_libXBMC_addonLog(int loglevel, char *format);
    bool CGoHelper_libXBMC_addonGetSetting(cchar_t *settingName, void *settingValue);
    void CGoHelper_libXBMC_addonFree();

    bool CGoHelper_libXBMC_pvrRegisterMe(void *handle);
    void CGoHelper_libXBMC_pvrTransferChannelEntry(ADDON_HANDLE handle, cPVR_CHANNEL_t *entry);
    void CGoHelper_libXBMC_pvrTransferChannelGroupEntry(ADDON_HANDLE handle, cPVR_CHANNEL_GROUP_t *entry);
    void CGoHelper_libXBMC_pvrTransferChannelGroupMember(ADDON_HANDLE handle, cPVR_CHANNEL_GROUP_MEMBER_t *entry);
    void CGoHelper_libXBMC_pvrTransferEpgEntry(ADDON_HANDLE handle, EPG_TAG entry);
    void CGoHelper_libXBMC_pvrSetProperty(PVR_NAMED_VALUE *properties, int index, char *key, char *value);
    void CGoHelper_libXBMC_pvrFree();

    typedef void *CGoHelper_libKODI_guilib;
    typedef void *CGoHelper_CAddonGUIWindow;
    typedef void *CGoHelper_CAddonGUIProgressControl;
    bool CGoHelper_libKODI_guilibRegisterMe(void *handle);
    void CGoHelper_libKODI_guilibLock();
    void CGoHelper_libKODI_guilibUnlock();
    bool CGoHelper_libKODI_guilibDialog_YesNo_ShowAndGetInput(cchar_t *heading, cchar_t *text, bool *bCanceled, cchar_t *noLabel, cchar_t *yesLabel);
    void CGoHelper_libKODI_guilibDialog_OK_ShowAndGetInput(cchar_t *heading, cchar_t *text);
    bool CGoHelper_libKODI_guilibDialog_Keyboard_ShowAndGetInput(char *strText, unsigned int iMaxStringSize, cchar_t *strHeading, bool allowEmptyResult, bool hiddenInput);
    // CGoHelper_CAddonGUIWindow CGoHelper_libKODI_guilibWindow_create(cchar_t *xmlFilename, cchar_t *defaultSkin, bool forceFallback, bool asDialog);
    // void CGoHelper_libKODI_guilibWindow_destroy(void *p);
    // CGoHelper_CAddonGUIProgressControl CGoHelper_libKODI_guilibControl_getProgress(void *window, int controlId);
    // void CGoHelper_libKODI_guilibSetProgressPercentage(CGoHelper_CAddonGUIProgressControl, float fPercent);
    // void CGoHelper_libKODI_guilibControl_releaseProgress(CGoHelper_libKODI_guilib GUI, void *p);
    void CGoHelper_libKODI_guilibFree();
    // void get_addon(void *ptr);

#ifdef __cplusplus
}
#endif