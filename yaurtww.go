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

type Manifest struct {
	Version string
	Assets  []ManifestAsset
}

type ManifestAsset struct {
	MD5Sum   string
	FileName string
}

// const CDN_URL = "http://cdn.urbanterror.info/urt/%s/%s/q3ut4/%s"
const CDN_URL = "http://cdn.urbanterror.info/urt/"

var (
	ManifestPath *string
	DownloadPath = flag.String([]string{"d", "-dest"}, "./", "Path to destination directory")
	Version      = flag.Bool([]string{"v", "-version"}, false, "Print the name and version")
)

func init() {
	ManifestPath = flag.String([]string{"m", "-manifest"}, RequiredFlag("Manifest is required."), "Path to yaurtww manifest")
}

func RequiredFlag(ErrorMessage string) string {
	// shim to trick the compiler so that we can actually call os.Exit(1) and
	// print a helpful error message when a flag is required.
	log.Fatalln(ErrorMessage)
	return "requiredstring"
}

func ReadManifest(path *string) (*Manifest, error) {
	var manifest = Manifest{}
	file, err := ioutil.ReadFile(*path)
	if err != nil {
		return manifest, err
	}
	return manifest, nil
}

func (asset ManifestAsset) Download(url string) error {
	var source io.Reader
	var sourceSize int64

	downloadPath := *DownloadPath + asset.FileName

	assetURL := url + asset.FileName
	resp, err := http.Get(assetURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error getting %s: HTTP status %v", assetURL, resp.Status)
	}

	i, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	sourceSize = int64(i)

	file, err := os.Create(downloadPath)
	if err != nil {
		return err
	}
	defer file.close()

	bar := pb.New(int(sourceSize)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.ShowSpeed = true
	bar.Start()

	writer := io.MultiWriter(file, bar)

	io.Copy(writer, resp.Body)
	bar.Finish()
}
