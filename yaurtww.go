package yaurtww

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cheggaaa/pb"

	flag "github.com/docker/docker/pkg/mflag"
)

// Manifest is the manifest provided by the urban terror site, and describes
// what assets are required and states the version of the game.
type Manifest struct {
	Version string
	Assets  []ManifestAsset
}

// ManifestAsset is a single manifest, and contains a checksum and filename.
type ManifestAsset struct {
	MD5Sum   string
	FileName string
}

// const CDNURL = "http://cdn.urbanterror.info/urt/%s/%s/q3ut4/%s"
const CDNURL = "http://cdn.urbanterror.info/urt/"

var (
	// ManifestPath is the path to the manifest.
	ManifestPath string
	// DownloadPath points to the destination directory for download.
	DownloadPath = flag.String([]string{"d", "-dest"}, "./", "Path to destination directory")
	// Version prints the version of yaurtww
	Version = flag.Bool([]string{"v", "-version"}, false, "Print the name and version")
)

func init() {
	ManifestPath = flag.String([]string{"m", "-manifest"}, RequiredFlag("Manifest is required."), "Path to yaurtww manifest")
}

// RequiredFlag is a shim to print error messages for flags.
func RequiredFlag(ErrorMessage string) string {
	// shim to trick the compiler so that we can actually call os.Exit(1) and
	// print a helpful error message when a flag is required.
	log.Fatalln(ErrorMessage)
	return "requiredstring"
}

// ReadManifest reads the stored manifest and returns it as Manifest
func ReadManifest(path *string) (manifest *Manifest, err error) {
	file, err := ioutil.ReadFile(*path)
	return
}

// Download downloads the file.
func (asset ManifestAsset) Download(url string) (err error) {
	var sourceSize int64

	downloadPath := *DownloadPath + asset.FileName

	assetURL := url + asset.FileName
	resp, err := http.Get(assetURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error getting %s: HTTP status %v", assetURL, resp.Status)
	}

	contentlength, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	sourceSize = int64(contentlength)

	file, err := os.Create(downloadPath)
	if err != nil {
		return
	}
	defer file.Close()

	bar := pb.New(int(sourceSize)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.ShowSpeed = true
	bar.Start()

	writer := io.MultiWriter(file, bar)

	io.Copy(writer, resp.Body)
	bar.Finish()
	return
}
