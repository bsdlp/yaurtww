package yaurtww

import flag "github.com/docker/docker/pkg/mflag"

type Manifest struct {
	Version string
	Assets  []ManifestAsset
}

type ManifestAsset struct {
	MD5Sum   string
	FileName string
}

const CDN_URL = "http://cdn.urbanterror.info/urt/%s/%s/q3ut4/%s"

var (
	ManifestPath *string
	DownloadPath = flag.String([]string{"d", "-dest"}, "./", "Path to destination directory")
)
