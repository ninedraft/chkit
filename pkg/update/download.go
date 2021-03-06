package update

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/blang/semver"
	"github.com/containerum/chkit/pkg/chkitErrors"
	"github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
)

type LatestChecker interface {
	LatestVersion() (semver.Version, error)
}

type Downloader interface {
	Download(version semver.Version) (rc io.ReadCloser, size int64, err error)
}

type LatestCheckerDownloader interface {
	LatestChecker
	Downloader
}

const (
	ErrUpdateCheck    = chkitErrors.Err("unable to check latest version")
	ErrUpdateDownload = chkitErrors.Err("unable to download latest version")
)

const (
	MaxFileSize = 50 * (1 << 20) // 50 megabytes
)

func DownloadFileName(version semver.Version) string {
	extension := "tar.gz"
	if runtime.GOOS == "windows" {
		extension = "zip"
	}
	return fmt.Sprintf("chkit_%s_%s_v%s.%s", runtime.GOOS, runtime.GOARCH, version, extension)
}

type GithubLatestCheckerDownloader struct {
	client      *resty.Client
	downloadUrl string
}

func NewGithubLatestCheckerDownloader(owner, repo string, debugRequests bool) *GithubLatestCheckerDownloader {
	re := resty.New().
		SetHostURL(
			fmt.Sprintf("https://api.github.com/repos/%s/%s/releases",
				owner,
				repo))
	if debugRequests {
		re.SetLogger(logrus.StandardLogger().
			WriterLevel(logrus.DebugLevel)).
			SetDebug(true)
	}
	return &GithubLatestCheckerDownloader{
		client:      re,
		downloadUrl: fmt.Sprintf("https://github.com/%s/%s/releases/download", owner, repo),
	}
}

func (gh *GithubLatestCheckerDownloader) LatestVersion() (semver.Version, error) {
	logrus.Debug("get latest version from github")

	var latestVersionResp struct {
		LatestVersion string `json:"tag_name"`
	}

	_, err := gh.client.R().SetResult(&latestVersionResp).Get("/latest")
	if err != nil {
		logrus.WithError(err).Errorf("error while getting latest version from github")
		return semver.MustParse("0.0.1-alpha"), chkitErrors.Wrap(ErrUpdateCheck, err)
	}

	vers, err := semver.ParseTolerant(latestVersionResp.LatestVersion)
	if err != nil {
		logrus.WithError(err).Errorf("error while parsing latest version tag")
		return semver.MustParse("0.0.1-alpha"), chkitErrors.Wrap(ErrUpdateCheck, err)
	}
	return vers, nil
}

func (gh *GithubLatestCheckerDownloader) Download(version semver.Version) (io.ReadCloser, int64, error) {
	logrus.Debug("download update")

	url := fmt.Sprintf("%s/v%s/%s", gh.downloadUrl, version, DownloadFileName(version))
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, chkitErrors.Wrap(ErrUpdateDownload, err)
	}

	return resp.Body, resp.ContentLength, nil
}

type FileSystemLatestCheckerDownloader struct {
	baseDir string
}

func NewFileSystemLatestCheckerDownloader(baseDir string) *FileSystemLatestCheckerDownloader {
	return &FileSystemLatestCheckerDownloader{
		baseDir: baseDir,
	}
}

func (fs *FileSystemLatestCheckerDownloader) LatestVersion() (semver.Version, error) {
	logrus.Debug("get latest version from filesystem")

	ver, err := ioutil.ReadFile(path.Join(fs.baseDir, "version"))
	if err != nil {
		return semver.MustParse("0.0.1-alpha"), chkitErrors.Wrap(ErrUpdateCheck, err)
	}

	sv, err := semver.ParseTolerant(string(ver))
	if err != nil {
		return semver.MustParse("0.0.1-alpha"), chkitErrors.Wrap(ErrUpdateCheck, err)
	}

	return sv, nil
}

func (fs *FileSystemLatestCheckerDownloader) Download(version semver.Version) (io.ReadCloser, int64, error) {
	logrus.Debug("get latest version package from filesystem")

	pkg, err := os.Open(path.Join(fs.baseDir, DownloadFileName(version)))
	if err != nil {
		return nil, 0, chkitErrors.Wrap(ErrUpdateDownload, err)
	}

	stat, err := pkg.Stat()
	if err != nil {
		return nil, 0, chkitErrors.Wrap(ErrUpdateDownload, err)
	}

	return pkg, stat.Size(), nil
}
