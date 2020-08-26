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

type GoHelper_libXBMC_pvr struct{}

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
	strTitle            uint64
	startTime           C.time_t
	endTime             C.time_t
	strPlotOutline      uint64
	strPlot             uint64
	strOriginalTitle    uint64
	strCast             uint64
	strDirector         uint64
	strWriter           uint64
	iYear               int32
	strIMDBNumber       uint64
	strIconPath         uint64
	iGenreType          int32
	iGenreSubType       int32
	strGenreDescription uint64
	firstAired          C.time_t
	iParentalRating     int32
	iStarRating         int32
	bNotify             bool
	iSeriesNumber       int32
	iEpisodeNumber      int32
	iEpisodePartNumber  int32
	strEpisodeName      uint64
	iFlags              uint32
	strSeriesLink       uint64
}

func NewlibXBMC_pvr() GoHelper_libXBMC_pvr {
	return GoHelper_libXBMC_pvr{}
}

func (pvr GoHelper_libXBMC_pvr) RegisterMe(hdl unsafe.Pointer) C.bool {
	return C.CGoHelper_libXBMC_pvrRegisterMe(hdl)
}

func (pvr GoHelper_libXBMC_pvr) TransferChannelEntry(handle C.ADDON_HANDLE, channel Channel) {
	f := [C.PVR_ADDON_INPUT_FORMAT_STRING_LENGTH]byte{}
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
	cBytes := C.CBytes(buf.Bytes())
	defer C.free(cBytes)
	C.CGoHelper_libXBMC_pvrTransferChannelEntry(handle, (*C.cPVR_CHANNEL_t)(cBytes))
}

func (pvr GoHelper_libXBMC_pvr) TransferChannelGroupEntry(handle C.ADDON_HANDLE, group ChannelGroup) {
	packed := GoPackChannelGroup{
		strGroupName: StrToUint8Arr(group.Name),
		bIsRadio:     group.IsRadio,
		iPosition:    uint32(group.Position),
	}

	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, packed)

	entry := (*C.cPVR_CHANNEL_GROUP_t)(C.CBytes(buf.Bytes()))
	C.CGoHelper_libXBMC_pvrTransferChannelGroupEntry(handle, entry)
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
	C.CGoHelper_libXBMC_pvrTransferChannelGroupMember(handle, entry)
}

func (pvr GoHelper_libXBMC_pvr) TransferEpgEntry(handle C.ADDON_HANDLE, epg EPG) {
	packed := GoPackEPGTag{
		iUniqueBroadcastId:  uint32(epg.BroadcastID),
		iUniqueChannelId:    uint32(epg.ChannelID),
		strTitle:            StrToPtr(epg.Title),
		startTime:           C.time_t(epg.StartTime),
		endTime:             C.time_t(epg.EndTime),
		strPlotOutline:      StrToPtr("strPlotOutline"),   //StrToPtr(epg.PlotOutline),
		strPlot:             StrToPtr("strPlot"),          //StrToPtr(epg.Plot),
		strOriginalTitle:    StrToPtr("strOriginalTitle"), //StrToPtr(epg.OriginalTitle),
		strCast:             StrToPtr("strCast"),          //StrToPtr(strings.Join(epg.Cast, C.EPG_STRING_TOKEN_SEPARATOR)),
		strDirector:         StrToPtr("strDirector"),      //StrToPtr(strings.Join(epg.Director, C.EPG_STRING_TOKEN_SEPARATOR)),
		strWriter:           StrToPtr("strWriter"),        //StrToPtr(strings.Join(epg.Writer, C.EPG_STRING_TOKEN_SEPARATOR)),
		iYear:               int32(epg.Year),
		strIMDBNumber:       StrToPtr(epg.IMDBNumber),
		strIconPath:         StrToPtr(epg.IconPath),
		iGenreType:          int32(epg.GenreType),
		iGenreSubType:       int32(epg.GenreSubType),
		strGenreDescription: StrToPtr(strings.Join(epg.GenreDescription, C.EPG_STRING_TOKEN_SEPARATOR)),
		firstAired:          C.time_t(epg.FirstAired),
		iParentalRating:     int32(epg.ParentalRating),
		iStarRating:         int32(epg.StarRating),
		bNotify:             epg.Notify,
		iSeriesNumber:       int32(epg.SeriesNumber),
		iEpisodeNumber:      int32(epg.EpisodeNumber),
		iEpisodePartNumber:  int32(epg.EpisodePartNumber),
		strEpisodeName:      StrToPtr(epg.EpisodeName),
		iFlags:              uint32(epg.Flags),
		strSeriesLink:       StrToPtr(epg.SeriesLink),
	}
	buf := &bytes.Buffer{}
	log.Println("Size of buffer before writing is", buf.Len())
	binary.Write(buf, binary.LittleEndian, packed)
	x := C.CBytes(buf.Bytes())
	log.Printf("Go Type:%T Size:%d", packed.iUniqueBroadcastId, unsafe.Sizeof(packed.iUniqueBroadcastId))
	log.Printf("C Type:%T Size:%d", C.uint(1), unsafe.Sizeof(C.uint(1)))

	log.Printf("Go Type:%T Size:%d", packed.strEpisodeName, unsafe.Sizeof(packed.strEpisodeName))
	log.Printf("C Type:%T Size:%d", C.CString(""), unsafe.Sizeof(C.CString("")))

	log.Printf("Go Type:%T Size:%d", packed.startTime, unsafe.Sizeof(packed.startTime))
	log.Printf("C Type:%T Size:%d", C.time_t(0), unsafe.Sizeof(C.time_t(0)))

	log.Printf("Go Type:%T Size:%d", packed.iYear, unsafe.Sizeof(packed.iYear))
	log.Printf("C Type:%T Size:%d", C.int(0), unsafe.Sizeof(C.int(0)))

	entry := *(*C.EPG_TAG)(x)
	// log.Println("Size of entry is", C.sizeof_cEPG_TAG_t)
	// log.Println("Pointer of entry is", int64(uintptr(unsafe.Pointer(entry))))
	C.CGoHelper_libXBMC_pvrTransferEpgEntry(handle, entry)
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
	C.CGoHelper_libXBMC_pvrFree()
}
