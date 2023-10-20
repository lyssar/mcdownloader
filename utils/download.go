package utils

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/kyokomi/emoji/v2"
	"github.com/lyssar/msdcli/logger"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, outputName string, path string) {
	fileToWrite := fmt.Sprintf("%s/%s", path, outputName)
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	f, _ := os.OpenFile(fileToWrite, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		color.LightCyan.Render(emoji.Emojize(":compass:"), "Downloading"),
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)

	logger.Success("Download successful.")
}
