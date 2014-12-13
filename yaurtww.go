package yaurtww

import (
	"io/ioutil"
	"log"

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
