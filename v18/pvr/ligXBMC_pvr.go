package pvr

import (
	"bytes"
	"encoding/binary"
	"log"
	"strings"
	"unsafe"
)

/*
#include "workaround.h"
*/
import "C"

type GoHelper_libXBMC_pvr struct {
	addon C.CGoHelper_libXBMC_pvr
}

type GoPackChannel struct {
	iUniqueId         uint32
	bIsRadio          bool
	iChannelNumber    uint32
	iSubChannelNumber uint32
	strChannelName    [1024]uint8
	strInputFormat    [32]uint8
	iEncryptionSystem uint32
	strIconPath       [1024]uint8
	bIsHidden         bool
}

type GoPackChannelGroup struct {
	strGroupName [1024]uint8
	bIsRadio     bool
	iPosition    uint32
}

type GoPackChannelGroupMember struct {
	strGroupName      [1024]uint8
	iChannelUniqueId  uint32
	iChannelNumber    uint32
	iSubChannelNumber uint32
}

type GoPackEPGTag struct {
	iUniqueBroadcastId  uint32
	iUniqueChannelId    uint32
	strTitle            int64
	startTime           uint64
	endTime             uint64
	strPlotOutline      int64
	strPlot             int64
	strOriginalTitle    int64
	strCast             int64
	strDirector         int64
	strWriter           int64
	iYear               int32
	strIMDBNumber       int64
	strIconPath         int64
	iGenreType          int32
	iGenreSubType       int32
	strGenreDescription int64
	firstAired          uint64
	iParentalRating     int32
	iStarRating         int32
	bNotify             bool
	iSeriesNumber       int32
	iEpisodeNumber      int32
	iEpisodePartNumber  int32
	strEpisodeName      int64
	iFlags              uint32
	strSeriesLink       int64
}

func NewlibXBMC_pvr() GoHelper_libXBMC_pvr {
	return GoHelper_libXBMC_pvr{addon: C.CGoHelper_libXBMC_pvrInit()}
}

func (pvr GoHelper_libXBMC_pvr) RegisterMe(hdl unsafe.Pointer) C.bool {
	return C.CGoHelper_libXBMC_pvrRegisterMe(pvr.addon, hdl)
}

func (pvr GoHelper_libXBMC_pvr) TransferChannelEntry(handle C.ADDON_HANDLE, channel Channel) {
	// XBMC.Log(XBMC.DEBUG, "Transfering channel ", channel.Name)
	f := [32]byte{}
	copy(f[:], "")

	packed := GoPackChannel{
		iUniqueId:         uint32(channel.ID),
		bIsRadio:          channel.IsRadio,
		iChannelNumber:    uint32(channel.Number),
		iSubChannelNumber: uint32(channel.SubNumber),
		strChannelName:    StrToUint8Arr(channel.Name),
		strInputFormat:    f,
		iEncryptionSystem: uint32(0),
		strIconPath:       StrToUint8Arr(channel.IconPath),
		bIsHidden:         channel.IsHidden,
	}

	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, packed)

	entry := (*C.cPVR_CHANNEL_t)(C.CBytes(buf.Bytes()))
	C.CGoHelper_libXBMC_pvrTransferChannelEntry(pvr.addon, handle, entry)
}

func (pvr GoHelper_libXBMC_pvr) TransferChannelGroupEntry(handle C.ADDON_HANDLE, group ChannelGroup) {
	log.Println("Transfering channel group", group.Name)
	packed := GoPackChannelGroup{
		strGroupName: StrToUint8Arr(group.Name),
		bIsRadio:     group.IsRadio,
		iPosition:    uint32(group.Position),
	}

	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, packed.strGroupName)
	binary.Write(buf, binary.LittleEndian, packed.bIsRadio)
	binary.Write(buf, binary.LittleEndian, packed.iPosition)

	entry := (*C.cPVR_CHANNEL_GROUP_t)(C.CBytes(buf.Bytes()))
	C.CGoHelper_libXBMC_pvrTransferChannelGroupEntry(pvr.addon, handle, entry)
}

func (pvr GoHelper_libXBMC_pvr) TransferChannelGroupMember(handle C.ADDON_HANDLE, group ChannelGroup, channel Channel) {

	packed := GoPackChannelGroupMember{
		strGroupName:      StrToUint8Arr(group.Name),
		iChannelUniqueId:  uint32(channel.ID),
		iChannelNumber:    uint32(channel.Number),
		iSubChannelNumber: uint32(channel.SubNumber),
	}

	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, packed)

	entry := (*C.cPVR_CHANNEL_GROUP_MEMBER_t)(C.CBytes(buf.Bytes()))
	C.CGoHelper_libXBMC_pvrTransferChannelGroupMember(pvr.addon, handle, entry)
}

func (pvr GoHelper_libXBMC_pvr) TransferEpgEntry(handle C.ADDON_HANDLE, epg EPG) {

	packed := GoPackEPGTag{
		iUniqueBroadcastId:  uint32(epg.BroadcastID),
		iUniqueChannelId:    uint32(epg.ChannelID),
		strTitle:            int64(uintptr(unsafe.Pointer(C.CString(epg.Title)))),
		startTime:           uint64(epg.StartTime),
		endTime:             uint64(epg.EndTime),
		strPlotOutline:      int64(uintptr(unsafe.Pointer(C.CString(epg.PlotOutline)))),
		strPlot:             int64(uintptr(unsafe.Pointer(C.CString(epg.Plot)))),
		strOriginalTitle:    int64(uintptr(unsafe.Pointer(C.CString(epg.OriginalTitle)))),
		strCast:             int64(uintptr(unsafe.Pointer(C.CString(strings.Join(epg.Cast, C.EPG_STRING_TOKEN_SEPARATOR))))),
		strDirector:         int64(uintptr(unsafe.Pointer(C.CString(strings.Join(epg.Director, C.EPG_STRING_TOKEN_SEPARATOR))))),
		strWriter:           int64(uintptr(unsafe.Pointer(C.CString(strings.Join(epg.Writer, C.EPG_STRING_TOKEN_SEPARATOR))))),
		iYear:               int32(epg.Year),
		strIMDBNumber:       int64(uintptr(unsafe.Pointer(C.CString(epg.IMDBNumber)))),
		strIconPath:         int64(uintptr(unsafe.Pointer(C.CString(epg.IconPath)))),
		iGenreType:          int32(epg.GenreType),
		iGenreSubType:       int32(epg.GenreSubType),
		strGenreDescription: int64(uintptr(unsafe.Pointer(C.CString(strings.Join(epg.GenreDescription, C.EPG_STRING_TOKEN_SEPARATOR))))),
		firstAired:          uint64(epg.FirstAired),
		iParentalRating:     int32(epg.ParentalRating),
		iStarRating:         int32(epg.StarRating),
		bNotify:             epg.Notify,
		iSeriesNumber:       int32(epg.SeriesNumber),
		iEpisodeNumber:      int32(epg.EpisodeNumber),
		iEpisodePartNumber:  int32(epg.EpisodePartNumber),
		strEpisodeName:      int64(uintptr(unsafe.Pointer(C.CString(epg.EpisodeName)))),
		iFlags:              uint32(epg.Flags),
		strSeriesLink:       int64(uintptr(unsafe.Pointer(C.CString(epg.SeriesLink)))),
	}
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, packed)
	entry := (*C.cEPG_TAG_t)(C.CBytes(buf.Bytes()))
	C.CGoHelper_libXBMC_pvrTransferEpgEntry(pvr.addon, handle, entry)
}

func (pvr GoHelper_libXBMC_pvr) SetProperties(properties *C.struct_PVR_NAMED_VALUE, stream Stream) int {
	C.CGoHelper_libXBMC_pvrSetProperty(properties, C.int(0), C.CString(C.PVR_STREAM_PROPERTY_STREAMURL), C.CString(stream.URL))
	index := 1
	for k, v := range stream.Properties {
		ck := C.CString(k)
		cv := C.CString(v)
		C.CGoHelper_libXBMC_pvrSetProperty(properties, C.int(index), ck, cv)
		C.free(unsafe.Pointer(ck))
		C.free(unsafe.Pointer(cv))
		index++
	}
	return index
}

func (pvr GoHelper_libXBMC_pvr) Free() {
	C.CGoHelper_libXBMC_pvrFree(pvr.addon)
}
