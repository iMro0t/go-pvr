#include "workaround.h"
#include "../kodi/xbmc_pvr_dll.h"

ADDON::CHelper_libXBMC_addon *XBMC = NULL;

CGoHelper_libXBMC_addon CGoHelper_libXBMC_addonInit()
{
    XBMC = new ADDON::CHelper_libXBMC_addon;
    return (void *)XBMC;
}

bool CGoHelper_libXBMC_addonRegisterMe(CGoHelper_libXBMC_addon XBMC, void *handle)
{
    ADDON::CHelper_libXBMC_addon *xbmc = (ADDON::CHelper_libXBMC_addon *)XBMC;
    return xbmc->RegisterMe(handle);
}

void CGoHelper_libXBMC_addonLog(CGoHelper_libXBMC_addon XBMC, int loglevel, char *format)
{
    ADDON::CHelper_libXBMC_addon *xbmc = (ADDON::CHelper_libXBMC_addon *)XBMC;
    ADDON::addon_log_t lvl = (ADDON::addon_log_t)loglevel;
    xbmc->Log(lvl, format);
}

void CGoHelper_libXBMC_addonFree(CGoHelper_libXBMC_addon XBMC)
{
    ADDON::CHelper_libXBMC_addon *xbmc = (ADDON::CHelper_libXBMC_addon *)XBMC;
    delete xbmc;
}

CGoHelper_libXBMC_pvr CGoHelper_libXBMC_pvrInit()
{
    CHelper_libXBMC_pvr *PVR = new CHelper_libXBMC_pvr;
    return (void *)PVR;
}

bool CGoHelper_libXBMC_pvrRegisterMe(CGoHelper_libXBMC_pvr PVR, void *handle)
{
    CHelper_libXBMC_pvr *pvr = (CHelper_libXBMC_pvr *)PVR;
    return pvr->RegisterMe(handle);
}

void CGoHelper_libXBMC_pvrTransferChannelEntry(CGoHelper_libXBMC_pvr PVR, ADDON_HANDLE handle, cPVR_CHANNEL_t *entry)
{
    CHelper_libXBMC_pvr *pvr = (CHelper_libXBMC_pvr *)PVR;
    pvr->TransferChannelEntry(handle, entry);
}

void CGoHelper_libXBMC_pvrTransferChannelGroupEntry(CGoHelper_libXBMC_pvr PVR, ADDON_HANDLE handle, cPVR_CHANNEL_GROUP_t *entry)
{
    CHelper_libXBMC_pvr *pvr = (CHelper_libXBMC_pvr *)PVR;
    pvr->TransferChannelGroup(handle, entry);
}

void CGoHelper_libXBMC_pvrTransferChannelGroupMember(CGoHelper_libXBMC_pvr PVR, ADDON_HANDLE handle, cPVR_CHANNEL_GROUP_MEMBER_t *entry)
{
    CHelper_libXBMC_pvr *pvr = (CHelper_libXBMC_pvr *)PVR;
    pvr->TransferChannelGroupMember(handle, entry);
}

void CGoHelper_libXBMC_pvrTransferEpgEntry(CGoHelper_libXBMC_pvr PVR, ADDON_HANDLE handle, cEPG_TAG_t *entry)
{
    CHelper_libXBMC_pvr *pvr = (CHelper_libXBMC_pvr *)PVR;
    pvr->TransferEpgEntry(handle, entry);
}

void CGoHelper_libXBMC_pvrSetProperty(PVR_NAMED_VALUE *properties, int index, char *key, char *value)
{
    strncpy(properties[index].strName, key, sizeof(properties[index].strName) - 1);
    strncpy(properties[index].strValue, value, sizeof(properties[index].strValue) - 1);
}

void CGoHelper_libXBMC_pvrFree(CGoHelper_libXBMC_pvr PVR)
{
    CHelper_libXBMC_pvr *pvr = (CHelper_libXBMC_pvr *)PVR;
    delete pvr;
}