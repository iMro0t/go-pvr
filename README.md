# Kodi PVR API Implemented in Go

Demo Add-on [pvr.demo.go](https://github.com/iMro0t/pvr.demo.go)

Supported Versions : _Leia_

## Quick Start

`$ go get github.com/iMro0t/go-pvr`

```go
package main

import "github.com/iMro0t/go-pvr/v18/pvr"

func init() { pvr.Call = main }

func main() {
    // Add Channels, Groups, EPG ...
}
```

## Functions

### Add Channels

```go
pvr.AddChannel(pvr.Channel{
    ID:        1,
    Number:    1,
    Name:      "Big Buck Bunny",
    IconPath:  "https://peach.blender.org/wp-content/uploads/poster_bunny_small.jpg",
    Live: pvr.Stream{
        URL: "http://distribution.bbb3d.renderfarming.net/video/mp4/bbb_sunflower_1080p_30fps_normal.mp4",
    },
})
```

### Add Channel Group

```go
pvr.AddChannelGroup(pvr.ChannelGroup{
    IsRadio:  false,
    Name:     "Kids",
    Position: 1,
    Members:  []int{1},
})
```

### Add EPG

```go
pvr.AddEPG(pvr.EPG{
    BroadcastID:      100,
    ChannelID:        1,
    Title:            "Big Buck Bunny",
    StartTime:        time.Now().Unix(),
    EndTime:          time.Now().Add(time.Hour).Unix(),
})
```

## Build Instructions

### Linux

1. `go build -buildmode=c-shared -o pvr.demo.go/pvr.demo.go.so.0.0.1`
2. `zip -rq pvr.demo.go-0.0.1.zip pvr.demo.go`

##### Useful links

- [Kodi's PVR user support](http://forum.kodi.tv/forumdisplay.php?fid=167)
- [Kodi's PVR development support](http://forum.kodi.tv/forumdisplay.php?fid=136)
