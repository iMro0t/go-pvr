#include "workaround.h"
#include "../kodi/xbmc_pvr_dll.h"

ADDON::CHelper_libXBMC_addon *XBMC = new ADDON::CHelper_libXBMC_addon;
CHelper_libKODI_guilib *GUI = new CHelper_libKODI_guilib;
CHelper_libXBMC_pvr *PVR = new CHelper_libXBMC_pvr;

// XBMC

bool CGoHelper_libXBMC_addonRegisterMe(void *handle)
{
    return XBMC->RegisterMe(handle);
}

void CGoHelper_libXBMC_addonLog(int loglevel, char *format)
{
    ADDON::addon_log_t lvl = (ADDON::addon_log_t)loglevel;
    XBMC->Log(lvl, format);
}

bool CGoHelper_libXBMC_addonGetSetting(cchar_t *settingName, void *settingValue)
{
    return XBMC->GetSetting(settingName, settingValue);
}

void CGoHelper_libXBMC_addonFree()
{
    delete XBMC;
}

// PVR

bool CGoHelper_libXBMC_pvrRegisterMe(void *handle)
{

    return PVR->RegisterMe(handle);
}

void CGoHelper_libXBMC_pvrTransferChannelEntry(ADDON_HANDLE handle, cPVR_CHANNEL_t *entry)
{
    char *str = new char[1024];
    sprintf(str, "------Channel ID : %d -------", entry->iUniqueId);
    XBMC->Log(ADDON::LOG_DEBUG, str);
    sprintf(str, "Channel is Radio : %d", entry->bIsRadio);
    XBMC->Log(ADDON::LOG_DEBUG, str);
    sprintf(str, "Channel Number : %d", entry->iChannelNumber);
    XBMC->Log(ADDON::LOG_DEBUG, str);
    sprintf(str, "Sub Channel Number : %d", entry->iSubChannelNumber);
    XBMC->Log(ADDON::LOG_DEBUG, str);
    XBMC->Log(ADDON::LOG_DEBUG, entry->strChannelName);
    XBMC->Log(ADDON::LOG_DEBUG, entry->strInputFormat);
    sprintf(str, "%d", entry->iEncryptionSystem);
    XBMC->Log(ADDON::LOG_DEBUG, str);
    XBMC->Log(ADDON::LOG_DEBUG, entry->strIconPath);
    sprintf(str, "%d", entry->bIsHidden);
    XBMC->Log(ADDON::LOG_DEBUG, str);
    PVR->TransferChannelEntry(handle, entry);
}

void CGoHelper_libXBMC_pvrTransferChannelGroupEntry(ADDON_HANDLE handle, cPVR_CHANNEL_GROUP_t *entry)
{

    PVR->TransferChannelGroup(handle, entry);
}

void CGoHelper_libXBMC_pvrTransferChannelGroupMember(ADDON_HANDLE handle, cPVR_CHANNEL_GROUP_MEMBER_t *entry)
{

    PVR->TransferChannelGroupMember(handle, entry);
}

void CGoHelper_libXBMC_pvrTransferEpgEntry(ADDON_HANDLE handle, EPG_TAG entry)
{
    try
    {
        // EPG_TAG *tag;
        // tag = (struct EPG_TAG*) malloc(sizeof(struct EPG_TAG));
        // (*tag).iUniqueBroadcastId = (*entry).iUniqueBroadcastId;
        // (*tag).iUniqueChannelId = (*entry).iUniqueChannelId;
        // (*tag).strTitle = (*entry).strTitle;
        // (*tag).startTime = (*entry).startTime;
        // (*tag).endTime = (*entry).endTime;
        EPG_TAG tag;
        memset(&tag, 0, sizeof(EPG_TAG));
        tag.iUniqueBroadcastId = entry.iUniqueBroadcastId;
        tag.iUniqueChannelId = entry.iUniqueChannelId;
        tag.strTitle = entry.strTitle;
        tag.startTime = entry.startTime;
        tag.endTime = entry.endTime;
        tag.strPlotOutline = entry.strPlotOutline;
        tag.strPlot = entry.strPlot;
        tag.strOriginalTitle = entry.strOriginalTitle;
        tag.strCast = entry.strCast;
        tag.strDirector = entry.strDirector;
        tag.strWriter = entry.strWriter;
        tag.iYear = 2020;//entry.iYear;
        tag.strIMDBNumber = "tt8110330";//NULL;//entry.strIMDBNumber;
        tag.strIconPath = NULL;//"   https://peach.blender.org/wp-content/uploads/bbb-splash.png?x28130"; // Flag
        tag.iGenreType = 0;//entry.iGenreType;
        tag.iGenreSubType = 0;//entry.iGenreSubType;
        tag.strGenreDescription = NULL;//entry.strGenreDescription; // Flag
        tag.firstAired = time(0);//entry.firstAired;
        tag.iParentalRating = 0;//entry.iParentalRating;
        tag.iStarRating = 0;//entry.iStarRating;
        tag.bNotify = false;//entry.bNotify;
        tag.iSeriesNumber = 0;//entry.iSeriesNumber;
        tag.iEpisodeNumber = 0;//entry.iEpisodeNumber;
        tag.iEpisodePartNumber = 0;//entry.iEpisodePartNumber;
        tag.strEpisodeName = "Episode Name";//entry.strEpisodeName;
        tag.iFlags = EPG_TAG_FLAG_UNDEFINED;
        tag.strSeriesLink = NULL;//entry.strSeriesLink;
        char *str = new char[1024];
        sprintf(str, "-----Channel ID : %u -------", tag.iUniqueChannelId);
        XBMC->Log(ADDON::LOG_DEBUG, str);
        sprintf(str, "Broadcast ID : %u", tag.iUniqueBroadcastId);
        XBMC->Log(ADDON::LOG_DEBUG, str);
        XBMC->Log(ADDON::LOG_DEBUG, tag.strTitle);
        sprintf(str, "Start TIME : %ld", tag.startTime);
        XBMC->Log(ADDON::LOG_DEBUG, str);
        sprintf(str, "End TIME : %ld", tag.endTime);
        XBMC->Log(ADDON::LOG_DEBUG, tag.strPlot);
        XBMC->Log(ADDON::LOG_DEBUG, tag.strOriginalTitle);
        XBMC->Log(ADDON::LOG_DEBUG, tag.strCast);
        XBMC->Log(ADDON::LOG_DEBUG, tag.strDirector);
        XBMC->Log(ADDON::LOG_DEBUG, tag.strWriter);
        // XBMC->Log(ADDON::LOG_DEBUG, tag.strIMDBNumber);
        // XBMC->Log(ADDON::LOG_DEBUG, tag.strIconPath);
        // XBMC->Log(ADDON::LOG_DEBUG, tag.strGenreDescription);
        // XBMC->Log(ADDON::LOG_DEBUG, tag.strEpisodeName);
        // XBMC->Log(ADDON::LOG_DEBUG, tag.strSeriesLink);
        PVR->TransferEpgEntry(handle, &tag);
        // free(&tag);
    }
    catch (char *excp)
    {
        XBMC->Log(ADDON::LOG_ERROR, excp);
    }
    catch (...)
    {
        XBMC->Log(ADDON::LOG_ERROR, "Unknwon exception occured while transfering EPG");
    }
}

void CGoHelper_libXBMC_pvrSetProperty(PVR_NAMED_VALUE *properties, int index, char *key, char *value)
{
    strncpy(properties[index].strName, key, sizeof(properties[index].strName) - 1);
    strncpy(properties[index].strValue, value, sizeof(properties[index].strValue) - 1);
}

void CGoHelper_libXBMC_pvrFree()
{

    delete PVR;
}

// GUI

// CGoHelper_libKODI_guilib CGoHelper_libKODI_guilibInit()
// {
//     // CHelper_libKODI_guilib *GUI = new CHelper_libKODI_guilib;
//     return (void *)GUI;
// }

bool CGoHelper_libKODI_guilibRegisterMe(void *handle)
{
    return GUI->RegisterMe(handle);
}

void CGoHelper_libKODI_guilibLock()
{
    GUI->Lock();
}

void CGoHelper_libKODI_guilibUnlock()
{
    GUI->Unlock();
}

bool CGoHelper_libKODI_guilibDialog_YesNo_ShowAndGetInput(cchar_t *heading, cchar_t *text, bool *bCanceled, cchar_t *noLabel = "", cchar_t *yesLabel = "")
{
    return GUI->Dialog_YesNo_ShowAndGetInput(heading, text, *bCanceled, noLabel, yesLabel);
}

void CGoHelper_libKODI_guilibDialog_OK_ShowAndGetInput(cchar_t *heading, cchar_t *text)
{
    GUI->Dialog_OK_ShowAndGetInput(heading, text);
}

bool CGoHelper_libKODI_guilibDialog_Keyboard_ShowAndGetInput(char *strText, unsigned int iMaxStringSize, cchar_t *strHeading, bool allowEmptyResult, bool hiddenInput)
{
    return GUI->Dialog_Keyboard_ShowAndGetInput(*strText, iMaxStringSize, strHeading, allowEmptyResult, hiddenInput);
}

// CGoHelper_CAddonGUIWindow CGoHelper_libKODI_guilibWindow_create(CGoHelper_libKODI_guilib GUI, cchar_t *xmlFilename, cchar_t *defaultSkin, bool forceFallback, bool asDialog)
// {
//     CHelper_libKODI_guilib *gui = (CHelper_libKODI_guilib *)GUI;
//     return (void *)gui->Window_create(xmlFilename, defaultSkin, forceFallback, asDialog);
// }

// void CGoHelper_libKODI_guilibWindow_destroy(CGoHelper_libKODI_guilib GUI, void *p)
// {
//     CHelper_libKODI_guilib *gui = (CHelper_libKODI_guilib *)GUI;
//     CAddonGUIWindow *window = (CAddonGUIWindow *)p;
//     gui->Window_destroy(window);
// }

// CGoHelper_CAddonGUIProgressControl CGoHelper_libKODI_guilibControl_getProgress(CGoHelper_libKODI_guilib GUI, void *window, int controlId)
// {
//     CHelper_libKODI_guilib *gui = (CHelper_libKODI_guilib *)GUI;
//     CAddonGUIWindow *w = (CAddonGUIWindow *)window;
//     CAddonGUIProgressControl *progress = gui->Control_getProgress(w, controlId);
//     w->SetFocusId(controlId);
//     w->SetControlLabel(controlId, "Test Lable");
//     progress->SetPercentage(float(50.00));
//     w->Show();
//     return (void *)progress;
// }

// void CGoHelper_libKODI_guilibSetProgressPercentage(CGoHelper_CAddonGUIProgressControl p, float fPercent)
// {
//     CAddonGUIProgressControl *progress = (CAddonGUIProgressControl *)p;
//     progress->SetPercentage(fPercent);
// }

// void CGoHelper_libKODI_guilibControl_releaseProgress(CGoHelper_libKODI_guilib GUI, void *p)
// {
//     CHelper_libKODI_guilib *gui = (CHelper_libKODI_guilib *)GUI;
//     CAddonGUIProgressControl *progress = (CAddonGUIProgressControl *)p;
//     gui->Control_releaseProgress(progress);
// }

void CGoHelper_libKODI_guilibFree()
{
    delete GUI;
}