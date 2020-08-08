package pvr

/*
#include "../kodi/xbmc_pvr_dll.h"
*/
import "C"

var (
	EPGGenre = struct {
		Unknown    int
		MovieDrama int
		News       int
		Show       int
		Sports     int
		Child      int
		Music      int
		Arts       int
		Social     int
		Science    int
		Hobby      int
		Speical    int
		Other      int
		Custom     int
	}{
		Unknown:    0,
		MovieDrama: 16,
		News:       32,
		Show:       48,
		Sports:     64,
		Child:      80,
		Music:      96,
		Arts:       112,
		Social:     128,
		Science:    144,
		Hobby:      160,
		Speical:    176,
		Other:      192,
		Custom:     256,
	}
	epgs []EPG
)

type EPG struct {
	BroadcastID       int
	ChannelID         int
	Title             string
	StartTime         int64
	EndTime           int64
	PlotOutline       string
	Plot              string
	OriginalTitle     string
	Cast              []string
	Director          []string
	Writer            []string
	Year              int
	IMDBNumber        string
	IconPath          string
	GenreType         int
	GenreSubType      int
	GenreDescription  []string
	FirstAired        int64
	ParentalRating    int
	StarRating        int
	Notify            bool
	SeriesNumber      int
	EpisodeNumber     int
	EpisodePartNumber int
	EpisodeName       string
	Flags             uint
	SeriesLink        string
	Catchup           Stream
}

func AddEPG(epg EPG) {
	epgs = append(epgs, epg)
}

func GetEPG(tag *C.cEPG_TAG_t) *EPG {
	for _, e := range epgs {
		if e.ChannelID == (int)(tag.iUniqueChannelId) && e.StartTime == int64(tag.startTime) {
			return &e
		}
	}
	return nil
}

//export GetEPGForChannel
func GetEPGForChannel(handle C.ADDON_HANDLE, channel *C.cPVR_CHANNEL_t, start C.time_t, end C.time_t) C.PVR_ERROR {
	for _, epg := range epgs {
		if int(channel.iUniqueId) == epg.ChannelID { //&& epg.EndTime > int(start) && epg.EndTime < int(end) {
			PVR.TransferEpgEntry(handle, epg)
		}
	}
	return C.PVR_ERROR_NO_ERROR
}

//export IsEPGTagPlayable
func IsEPGTagPlayable(tag *C.cEPG_TAG_t, isPlayable *C.bool) C.PVR_ERROR {
	epg := GetEPG(tag)
	if epg != nil && epg.Catchup.URL == "" {
		*isPlayable = false
	} else {
		*isPlayable = true
	}
	return C.PVR_ERROR_NO_ERROR
}

//export GetEPGTagStreamProperties
func GetEPGTagStreamProperties(tag *C.cEPG_TAG_t, props *C.PVR_NAMED_VALUE, propsCount *C.uint) C.PVR_ERROR {
	if tag == nil || props == nil || propsCount == nil || len(epgs) == 0 {
		return C.PVR_ERROR_SERVER_ERROR
	}

	if (int)(*propsCount) < 2 {
		return C.PVR_ERROR_INVALID_PARAMETERS
	}
	if e := GetEPG(tag); e != nil {
		count := PVR.SetProperties(props, e.Catchup)
		*propsCount = C.uint(count)
		return C.PVR_ERROR_NO_ERROR
	}
	return C.PVR_ERROR_SERVER_ERROR
}

//export IsEPGTagRecordable
func IsEPGTagRecordable(*C.cEPG_TAG_t, *C.bool) C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }

//export GetEPGTagEdl
func GetEPGTagEdl(epgTag *C.cEPG_TAG_t, edl *C.PVR_EDL_ENTRY, size *C.int) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export SetEPGTimeFrame
func SetEPGTimeFrame(C.int) C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }
