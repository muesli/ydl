package ydl

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestBoolOption(t *testing.T) {
	youtubeDl := NewYdl()
	youtubeDl.Options.ExtractAudio.Value = true
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "-x") {
		t.Fail()
	}
}

func TestTimeOption(t *testing.T) {
	youtubeDl := NewYdl()
	youtubeDl.Options.Date.Value = time.Date(2016, time.March, 19, 0, 0, 0, 0, time.UTC)
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "--date 20160319") {
		t.Fail()
	}
}

func TestFileSizeRateOption1(t *testing.T) {
	youtubeDl := NewYdl()
	youtubeDl.Options.BufferSize.Value = FileSizeRateFromString("5.5M")
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "--buffer-size 5.5M") {
		t.Fail()
	}
}

func TestFileSizeRateOption2(t *testing.T) {
	youtubeDl := NewYdl()
	youtubeDl.Options.BufferSize.Value = FileSizeRateFromValues(5.5, 'M')
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "--buffer-size 5.5M") {
		t.Fail()
	}
}

func TestIntOption(t *testing.T) {
	youtubeDl := NewYdl()
	youtubeDl.Options.SocketTimeout.Value = 5
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "--socket-timeout 5") {
		t.Fail()
	}
}

func TestIntOptionNegativeIsInfinite(t *testing.T) {
	youtubeDl := NewYdl()
	youtubeDl.Options.Retries.Value = -1
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "-R infinite") {
		t.Fail()
	}
}

func TestStringOption(t *testing.T) {
	youtubeDl := NewYdl()
	youtubeDl.Options.Username.Value = "testUser"
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "-u testUser") {
		t.Fail()
	}
}

func TestDownloadInfo(t *testing.T) {
	youtubeDl := NewYdl()
	info, err := youtubeDl.FetchInfo("https://www.youtube.com/watch?v=dQw4w9WgXcQ")

	if err != nil {
		t.Error(err)
	}

	exp := "Rick Astley - Never Gonna Give You Up (Official Music Video)"
	if info.Title != exp {
		t.Errorf("Expected '%s', got '%s'", exp, info.Title)
	}
}

func TestDownload(t *testing.T) {
	youtubeDl := NewYdl()

	ch, err := youtubeDl.Download("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	if err != nil {
		t.Error(err)
	}

	for p := range ch {
		fmt.Printf("%f%%, %d bytes, %d rate, %s eta\n", p.Percentage, p.Total, p.Rate, p.ETA)
	}
}
