package vlcctrl

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Status contains information related to the VLC instance status. Use parseStatus to parse the response from a
// status.go function.
type Status struct {
	// TODO: The Status structure is still a work in progress
	Fullscreen    bool              `json:"fullscreen"`
	Stats         Stats             `json:"stats"`
	AspectRatio   string            `json:"aspectratio"`
	AudioDelay    float64           `json:"audiodelay"`
	APIVersion    uint              `json:"apiversion"`
	CurrentPlID   uint              `json:"currentplid"`
	Time          uint              `json:"time"`
	Volume        uint              `json:"volume"`
	Length        uint              `json:"length"`
	Random        bool              `json:"random"`
	AudioFilters  map[string]string `json:"audiofilters"`
	Rate          float64           `json:"rate"`
	VideoEffects  VideoEffects      `json:"videoeffects"`
	State         string            `json:"state"`
	Loop          bool              `json:"loop"`
	Version       string            `json:"version"`
	Position      float64           `json:"position"`
	Information   Information       `json:"information"`
	Repeat        bool              `json:"repeat"`
	SubtitleDelay float64           `json:"subtitledelay"`
	Equalizer     Equalizer         `json:"equalizer"`
}

// Stats contains certain statistics of a VLC instance. A Stats variable is included in Status
type Stats struct {
	InputBitRate        float64 `json:"inputbitrate"`
	SentBytes           uint    `json:"sentbytes"`
	LosABuffers         uint    `json:"lostabuffers"`
	AveragedEMuxBitrate float64 `json:"averagedemuxbitrate"`
	ReadPackets         uint    `json:"readpackets"`
	DemuxReadPackets    uint    `json:"demuxreadpackets"`
	LostPictures        uint    `json:"lostpictures"`
	DisplayedPictures   uint    `json:"displayedpictures"`
	SentPackets         uint    `json:"sentpackets"`
	DemuxReadBytes      uint    `json:"demuxreadbytes"`
	DemuxBitRate        float64 `json:"demuxbitrate"`
	PlayedABuffers      uint    `json:"playedabuffers"`
	DemuxDiscontinuity  uint    `json:"demuxdiscontinuity"`
	DecodeAudio         uint    `json:"decodedaudio"`
	SendBitRate         float64 `json:"sendbitrate"`
	ReadBytes           uint    `json:"readbytes"`
	AverageInputBitRate float64 `json:"averageinputbitrate"`
	DemuxCorrupted      uint    `json:"demuxcorrupted"`
	DecodedVideo        uint    `json:"decodedvideo"`
}

// VideoEffects contains the current video effects configuration. A VideoEffects variable is included in Status
type VideoEffects struct {
	Hue        int `json:"hue"`
	Saturation int `json:"saturation"`
	Contrast   int `json:"contrast"`
	Brightness int `json:"brightness"`
	Gamma      int `json:"gamma"`
}

// VideoEffects contains information related to the item currently being played. An Information variable
// is included in Status
type Information struct {
	Chapter int `json:"chapter"`
	// TODO: Chapters definition might need to be changed
	Chapters []interface{} `json:"chapters"`
	Title    int           `json:"title"`
	// TODO: Category definition might need to be updated/modified
	Category map[string]struct {
		Filename      string `json:"filename"`
		Codec         string `json:"Codec"`
		Channels      string `json:"Channels"`
		BitsPerSample string `json:"Bits_per_sample"`
		Type          string `json:"Type"`
		SampleRate    string `json:"Sample_rate"`
	} `json:"category"`
	Titles []interface{} `json:"titles"`
}

// Equalizer contains information related to the equalizer configuration. An Equalizer variable is included in Status
type Equalizer struct {
	Presets map[string]string `json:"presets"`
	Bands   map[string]string `json:"bands"`
	Preamp  int               `json:"preamp"`
}

// ParseStatus parses GetStatus() responses to Status struct.
func ParseStatus(statusResponse string) (status Status, err error) {
	err = json.Unmarshal([]byte(statusResponse), &status)
	if err != nil {
		return
	}
	return status, nil
}

// GetStatus returns a Status object containing information of the instances' status
func (instance *VLC) GetStatus() (status Status, err error) {
	// Make request
	var response string
	response, _, _, err = instance.RequestMaker("/requests/status.json")
	// Error handling
	if err != nil {
		return
	}
	// Parse response to Status
	status, err = ParseStatus(response)
	return
}

// Play playlist item with given id. If id is omitted, play last active item
func (instance *VLC) Play(itemID ...int) (response string, statusCode int, err error) {
	// Check variadic arguments and form urlSegment
	if len(itemID) > 1 {
		err = errors.New("please provide only up to one ID")
		return
	}
	urlSegment := "/requests/status.json?command=pl_play"
	if len(itemID) == 1 {
		urlSegment = urlSegment + "&id=" + strconv.Itoa(itemID[0])
	}
	response, _, statusCode, err = instance.RequestMaker(urlSegment)
	return
}

// Toggle Pause: If current state was 'stop', play item with given id, if no id specified, play current item.
// If no current item, play 1st item in the playlist.
func (instance *VLC) Pause(itemID ...int) (response string, statusCode int, err error) {
	// Check variadic arguments and form urlSegment
	if len(itemID) > 1 {
		err = errors.New("please provide only up to one ID")
		return
	}
	urlSegment := "/requests/status.json?command=pl_pause"
	if len(itemID) == 1 {
		urlSegment = urlSegment + "&id=" + strconv.Itoa(itemID[0])
	}
	response, _, statusCode, err = instance.RequestMaker(urlSegment)
	return
}

// Stop Playback
func (instance *VLC) Stop() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_stop")
	return
}

// Next
func (instance *VLC) Next() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_next")
	return
}

// Previous
func (instance *VLC) Previous() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_previous")
	return
}

// Empty Playlist
func (instance *VLC) EmptyPlaylist() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_empty")
	return
}

// Random Playback Toggle
func (instance *VLC) ToggleLoop() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_random")
	return
}

// Playback Looping Toggle
func (instance *VLC) ToggleRepeat() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_loop")
	return
}

// Playback Repeat Toggle
func (instance *VLC) ToggleRandom() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_repeat")
	return
}

// Fullscreen Toggle
func (instance *VLC) ToggleFullscreen() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=fullscreen")
	return
}

// Add URI to playlist and start playback. The option field is optional, and can have the values: noaudio, novideo
func (instance *VLC) AddStart(uri string, option ...string) (response string, statusCode int, err error) {
	// Check variadic arguments and form urlSegment
	if len(option) > 1 {
		err = errors.New("please provide only one option")
		return
	}
	urlSegment := "/requests/status.json?command=in_play&input=" + uri
	if len(option) == 1 {
		if (option[0] != "noaudio") && (option[0] != "novideo") {
			err = errors.New("invalid option")
			return
		}
		urlSegment = urlSegment + "&option=" + option[0]
	}
	response, _, statusCode, err = instance.RequestMaker(urlSegment)
	return
}

// Add URI to playlist
func (instance *VLC) Add(uri string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=in_enqueue&input=" + uri)
	return
}

// Add subtitle from URI to currently playing file
func (instance *VLC) AddSubtitle(uri string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=addsubtitle&val=" + uri)
	return
}

// Resume playback if paused, else do nothing
func (instance *VLC) Resume() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_forceresume")
	return
}

// Pause playback, do nothing if already paused
func (instance *VLC) ForcePause() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_forcepause")
	return
}

// Delete item with given id from playlist
func (instance *VLC) Delete(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_delete&id=" + strconv.Itoa(id))
	return
}

// Set Audio Delay in seconds
func (instance *VLC) AudioDelay(delay float64) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=audiodelay&val=" + strconv.FormatFloat(delay, 'f', -1, 64))
	return
}

// Set Subtitle Delay in seconds
func (instance *VLC) SubDelay(delay float64) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=subdelay&val=" + strconv.FormatFloat(delay, 'f', -1, 64))
	return
}

// Set Playback Rate. Must be > 0
func (instance *VLC) PlaybackRate(rate float64) (response string, statusCode int, err error) {
	if rate <= 0 {
		err = errors.New("rate must be greater than 0")
		return
	}
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=rate&val=" + strconv.FormatFloat(rate, 'f', -1, 64))
	return
}

// Set aspect ratio. Must be one of the following values. Any other value will reset aspect ratio to default.
// Valid aspect ratio values: 1:1 , 4:3 , 5:4 , 16:9 , 16:10 , 221:100 , 235:100 , 239:100
func (instance *VLC) AspectRatio(ratio string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=aspectratio&val=" + ratio)
	return
}

// Sort playlist using sort mode <val> and order <id>.
// If id=0 then items will be sorted in normal order, if id=1 they will be sorted in reverse order.
// A non exhaustive list of sort modes: 0 Id, 1 Name, 3 Author, 5 Random, 7 Track number.
func (instance *VLC) Sort(id int, val int) (response string, statusCode int, err error) {
	if (id != 0) && (id != 1) {
		err = errors.New("sorting order must be 0 or 1")
		return
	}
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_sort&id=" + strconv.Itoa(id) + "&val=" + strconv.Itoa(val))
	return
}

// Toggle enable service discovery module <val>.
// Typical values are: sap shoutcast, podcast, hal
func (instance *VLC) ToggleSD(val string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=pl_sd&val=" + val)
	return
}

// Set Volume level <val> (can be absolute integer, or +/- relative value). Percentage isn't working at the moment.
// Allowed values are of the form: +<int>, -<int>, <int> or <int>%
func (instance *VLC) Volume(val string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=volume&val=" + val)
	return
}

// Seek to <val>
//  Allowed values are of the form:
//    [+ or -][<int><H or h>:][<int><M or m or '>:][<int><nothing or S or s or ">]
//    or [+ or -]<int>%
//    (value between [ ] are optional, value between < > are mandatory)
//  examples:
//    1000 -> seek to the 1000th second
//    +1H:2M -> seek 1 hour and 2 minutes forward
//    -10% -> seek 10% back
func (instance *VLC) Seek(val string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=seek&val=" + val)
	return
}

// Sets the preamp gain value, must be >=-20 and <=20
func (instance *VLC) Preamp(gain int) (response string, statusCode int, err error) {
	if (gain < -20) || (gain > 20) {
		err = errors.New("preamp must be between -20 and 20")
		return
	}
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=preamp&val=" + strconv.Itoa(gain))
	return
}

// Set the gain for a specific Equalizer band
func (instance *VLC) SetEQ(band int, gain int) (response string, statusCode int, err error) {
	if (gain < -20) || (gain > 20) {
		err = errors.New("gain must be between -20 and 20")
		return
	}
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=equalizer&band=" + strconv.Itoa(band) + "&val=" + strconv.Itoa(gain))
	return
}

// Toggle the EQ (true to enable, false to disable)
func (instance *VLC) ToggleEQ(enable bool) (response string, statusCode int, err error) {
	enableStr := "0"
	if enable == true {
		enableStr = "1"
	}
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=enableeq&val=" + enableStr)
	return
}

// Set the equalizer preset as per the id specified
func (instance *VLC) SetEQPreset(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=setpreset&id=" + strconv.Itoa(id))
	return
}

// Select the title using the title number
func (instance *VLC) SelectTitle(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=title&val=" + strconv.Itoa(id))
	return
}

// Select chapter using the chapter number
func (instance *VLC) SelectChapter(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=chapter&val=" + strconv.Itoa(id))
	return
}

// Select the audio track (use the number from the stream)
func (instance *VLC) SelectAudioTrack(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=audio_track&val=" + strconv.Itoa(id))
	return
}

// Select the video track (use the number from the stream)
func (instance *VLC) SelectVideoTrack(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=video_track&val=" + strconv.Itoa(id))
	return
}

// Select the subtitle track (use the number from the stream)
func (instance *VLC) SelectSubtitleTrack(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status.json?command=subtitle_track&val=" + strconv.Itoa(id))
	return
}
