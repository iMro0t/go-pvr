package pvr

/*
#include "../kodi/xbmc_pvr_dll.h"
*/
import "C"
import (
	"log"
)

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
	epgs           = []EPG{}
	epgHasCallback = false
	epgCallback    = func(channelID int, _, _ int64) ([]EPG, error) {
		resp := []EPG{}
		for _, epg := range epgs {
			if epg.ChannelID == channelID {
				resp = append(resp, epg)
			}
		}
		return resp, nil
	}
	epgStreamPropsCallback = func(channelID int, start int64) *Stream {
		if e := GetEPG(channelID, start); e != nil && e.Catchup.URL != "" {
			return &e.Catchup
		}
		return nil
	}
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
	mutex.Lock()
	epgs = append(epgs, epg)
	mutex.Unlock()
}

func GetEPG(channelID int, start int64) *EPG {
	for _, e := range epgs {
		if e.ChannelID == channelID && e.StartTime == start {
			return &e
		}
	}
	return nil
}

func SetEPGCallback(f func(channelID int, start int64, end int64) ([]EPG, error)) {
	epgHasCallback = true
	epgCallback = f
}

func SetEPGCatchupCallback(f func(channelID int, start int64) *Stream) {
	epgStreamPropsCallback = f
}

//export GetEPGForChannel
func GetEPGForChannel(handle C.ADDON_HANDLE, channel *C.cPVR_CHANNEL_t, start C.time_t, end C.time_t) C.PVR_ERROR {
	var rEPGs []EPG
	var err error
	rEPGs, err = epgCallback(int(channel.iUniqueId), int64(start), int64(end))
	if err != nil {
		log.Println(err)
		return C.PVR_ERROR_SERVER_ERROR
	}
	for _, e := range rEPGs {
		if epgHasCallback {
			go AddEPG(e)
		}
		if int(channel.iUniqueId) == e.ChannelID { //&& e.EndTime > int(start) && e.EndTime < int(end) {
			PVR.TransferEpgEntry(handle, e)
		}
	}
	return C.PVR_ERROR_NO_ERROR
}

//export IsEPGTagPlayable
func IsEPGTagPlayable(tag *C.cEPG_TAG_t, isPlayable *C.bool) C.PVR_ERROR {
	if e := GetEPG((int)(tag.iUniqueChannelId), int64(tag.startTime)); e != nil && e.Catchup.URL != "" {
		*isPlayable = true
		return C.PVR_ERROR_NO_ERROR
	}
	*isPlayable = false
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
	if catchup := epgStreamPropsCallback((int)(tag.iUniqueChannelId), int64(tag.startTime)); catchup != nil {
		count := PVR.SetProperties(props, *catchup)
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
