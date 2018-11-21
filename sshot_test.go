package sshot

import (
	"fmt"
	"os"
	"testing"
)

func TestGetSShot(t *testing.T) {
	shotsDir := "shots"
	os.Mkdir(shotsDir, 0777)
	params := Parameters{
		Sizes:     "1024x768",
		Crop:      true,
		Scale:     "0.9",
		Timeout:   "30",
		Filename:  fmt.Sprintf(`%s/<%%= url %%>`, shotsDir),
		UserAgent: "",
		Format:    "jpg",
	}
	urls := []string{
		"http://google.com",
		"https://dbconvert.com",
		"http://something-that-doesnot-exists.com",
	}
	GetShots(urls, params)
	DeleteZeroLengthFiles(shotsDir)
	return
}
