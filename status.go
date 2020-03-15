package vlcctrl

import (
	"errors"
	"strconv"
)

// Status : Returns information of the instances' status. Among them:
// Get VLC status information, current item info and meta
// Get VLC version, and http api version
// Displays the equalizer band gains:
//     Band 0: 60 Hz, 1: 170 Hz, 2: 310 Hz, 3: 600 Hz, 4: 1 kHz,
//     5: 3 kHz, 6: 6 kHz, 7: 12 kHz , 8: 14 kHz , 9: 16 kHz
// Displays the list of presets available for the equalizer
func (instance *vlc) Status() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format)
	return
}

// Play playlist item with given id. If id is omitted, play last active item
func (instance *vlc) Play(itemID ...int) (response string, statusCode int, err error) {
	// Check variadic arguments and form urlSegment
	if len(itemID) > 1 {
		err = errors.New("please provide only up to one ID")
		return
	}
	urlSegment := "/requests/status." + instance.Format + "?command=pl_play"
	if len(itemID) == 1 {
		urlSegment = urlSegment + "&id=" + strconv.Itoa(itemID[0])
	}
	response, _, statusCode, err = instance.RequestMaker(urlSegment)
	return
}

// Toggle Pause: If current state was 'stop', play item with given id, if no id specified, play current item.
// If no current item, play 1st item in the playlist.
func (instance *vlc) Pause(itemID ...int) (response string, statusCode int, err error) {
	// Check variadic arguments and form urlSegment
	if len(itemID) > 1 {
		err = errors.New("please provide only up to one ID")
		return
	}
	urlSegment := "/requests/status." + instance.Format + "?command=pl_pause"
	if len(itemID) == 1 {
		urlSegment = urlSegment + "&id=" + strconv.Itoa(itemID[0])
	}
	response, _, statusCode, err = instance.RequestMaker(urlSegment)
	return
}

// Stop
func (instance *vlc) Stop() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_stop")
	return
}

// Next
func (instance *vlc) Next() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_next")
	return
}

// Previous
func (instance *vlc) Previous() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_previous")
	return
}

// Empty Playlist
func (instance *vlc) EmptyPlaylist() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_empty")
	return
}

// Random Playback Toggle
func (instance *vlc) ToggleLoop() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_random")
	return
}

// Playback Looping Toggle
func (instance *vlc) ToggleRepeat() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_loop")
	return
}

// Playback Repeat Toggle
func (instance *vlc) ToggleRandom() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_repeat")
	return
}

// Playback Repeat Toggle
func (instance *vlc) ToggleFullscreen() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=fullscreen")
	return
}

// Add URI to playlist and start playback. the option field is optional, and can have the values: noaudio, novideo
// URI string is expected to be percent-encoded like URLs
func (instance *vlc) AddStart(uri string, option ...string) (response string, statusCode int, err error) {
	// Check variadic arguments and form urlSegment
	if len(option) > 1 {
		err = errors.New("please provide only one option")
		return
	}
	urlSegment := "/requests/status." + instance.Format + "?command=in_play&input=" + uri
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

// Add URI to playlist. URI string is expected to be percent-encoded like URLs
func (instance *vlc) Add(uri string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=in_enqueue&input=" + uri)
	return
}

// Add subtitle to currently playing file
func (instance *vlc) AddSubtitle(uri string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=addsubtitle&val=" + uri)
	return
}

// Resume playback if paused, else do nothing
func (instance *vlc) Resume() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_forceresume")
	return
}

// Pause playback, do nothing if already paused
func (instance *vlc) ForcePause() (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_forcepause")
	return
}

// Delete item with given id from playlist
func (instance *vlc) Delete(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_delete&id=" + strconv.Itoa(id))
	return
}

// Set Audio Delay in seconds
func (instance *vlc) AudioDelay(delay float64) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=audiodelay&val=" + strconv.FormatFloat(delay, 'f', -1, 64))
	return
}

// Set Subtitle Delay in seconds
func (instance *vlc) SubDelay(delay float64) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=subdelay&val=" + strconv.FormatFloat(delay, 'f', -1, 64))
	return
}

// Set Playback Rate. Must be > 0
func (instance *vlc) PlaybackRate(rate float64) (response string, statusCode int, err error) {
	if rate <= 0 {
		err = errors.New("rate must be greater than 0")
		return
	}
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=rate&val=" + strconv.FormatFloat(rate, 'f', -1, 64))
	return
}

// Set aspect ratio. Must be one of the following values. Any other value will reset aspect ratio to default
// Valid aspect ratio values: 1:1 , 4:3 , 5:4 , 16:9 , 16:10 , 221:100 , 235:100 , 239:100
func (instance *vlc) AspectRatio(ratio string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=aspectratio&val=" + ratio)
	return
}

// Sort playlist using sort mode <val> and order <id>
// If id=0 then items will be sorted in normal order, if id=1 they will be sorted in reverse order
// A non exhaustive list of sort modes: 0 Id, 1 Name, 3 Author, 5 Random, 7 Track number
func (instance *vlc) Sort(id int, val int) (response string, statusCode int, err error) {
	if (id != 0) && (id != 1) {
		err = errors.New("sorting order must be 0 or 1")
		return
	}
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_sort&id=" + strconv.Itoa(id) + "&val=" + strconv.Itoa(val))
	return
}

// Toggle enable service discovery module <val>
// Typical values are: sap shoutcast, podcast, hal
func (instance *vlc) ToggleSD(val string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=pl_sd&val=" + val)
	return
}

// Set Volume level <val> (can be absolute integer, or +/- relative value). Percentage isn't working at the moment.
// Allowed values are of the form: +<int>, -<int>, <int> or <int>%
func (instance *vlc) Volume(val string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=volume&val=" + val)
	return
}

// Seek to <val>
//   Allowed values are of the form:
//    [+ or -][<int><H or h>:][<int><M or m or '>:][<int><nothing or S or s or ">]
//    or [+ or -]<int>%
//    (value between [ ] are optional, value between < > are mandatory)
//  examples:
//    1000 -> seek to the 1000th second
//    +1H:2M -> seek 1 hour and 2 minutes forward
//    -10% -> seek 10% back
func (instance *vlc) Seek(val string) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=seek&val=" + val)
	return
}

// Sets the preamp gain value, must be >=-20 and <=20
func (instance *vlc) Preamp(gain int) (response string, statusCode int, err error) {
	if (gain < -20) || (gain > 20) {
		err = errors.New("preamp must be between -20 and 20")
		return
	}
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=preamp&val=" + strconv.Itoa(gain))
	return
}

// Set the gain for a specific Equalizer band
func (instance *vlc) SetEQ(band int, gain int) (response string, statusCode int, err error) {
	if (gain < -20) || (gain > 20) {
		err = errors.New("gain must be between -20 and 20")
		return
	}
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=equalizer&band=" + strconv.Itoa(band) + "&val=" + strconv.Itoa(gain))
	return
}

// Toggle the EQ (true to enable, false to disable)
func (instance *vlc) ToggleEQ(enable bool) (response string, statusCode int, err error) {
	enableStr := "0"
	if enable == true {
		enableStr = "1"
	}
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=enableeq&val=" + enableStr)
	return
}

// Set the equalizer preset as per the id specified
func (instance *vlc) SetEQPreset(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=setpreset&id=" + strconv.Itoa(id))
	return
}

// Select the title using the title number
func (instance *vlc) SelectTitle(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=title&val=" + strconv.Itoa(id))
	return
}

// Select chapter using the chapter number
func (instance *vlc) SelectChapter(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=chapter&val=" + strconv.Itoa(id))
	return
}

// Select the audio track (use the number from the stream)
func (instance *vlc) SelectAudioTrack(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=audio_track&val=" + strconv.Itoa(id))
	return
}

// Select the video track (use the number from the stream)
func (instance *vlc) SelectVideoTrack(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=video_track&val=" + strconv.Itoa(id))
	return
}

// Select the subtitle track (use the number from the stream)
func (instance *vlc) SelectSubtitleTrack(id int) (response string, statusCode int, err error) {
	response, _, statusCode, err = instance.RequestMaker("/requests/status." + instance.Format + "?command=subtitle_track&val=" + strconv.Itoa(id))
	return
}
