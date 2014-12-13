package yaurtww

import (
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
	log.Fatalln(ErrorMessage)
	return "requiredstring"
}
