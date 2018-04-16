package angel

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/containerum/chkit/pkg/configdir"
	"github.com/containerum/chkit/pkg/configuration"
	"github.com/containerum/chkit/pkg/context"
	"github.com/skratchdot/open-golang/open"
)

var reportPreambula = fmt.Sprintf(`[REPORT]
chkit fatal report
version: %s
os: %s %s
`, context.GlobalContext.Version, runtime.GOOS, runtime.GOARCH)

func Angel(sin interface{}) {
	report := reportPreambula
	switch recoverData := sin.(type) {
	case nil:
		return
	case error:
		report += fmt.Sprintf("[FATAL] %v", recoverData)
	default:
		report += fmt.Sprintf("[FATAL] %v\n%s", recoverData, string(debug.Stack()))
	}
	logDir := configdir.LogDir()
	logFileName := configuration.LogFileName()
	logFilePath := path.Join(logDir, logFileName)
	reportFile := path.Join(logDir, "report.txt")

	err := ioutil.WriteFile(reportFile, []byte(report), os.ModePerm)
	if err != nil {
		fmt.Printf("[FATAL] something completely wrong.\n")
		fmt.Printf("Please, send report and log file from %q to support@exonlab.omnidesk.ru", logFilePath)
		return
	}
	/*
		logTail, err := readLogTail(logFilePath)
		if err != nil {
			fmt.Printf("[FATAL] something completely wrong.\n")
			fmt.Printf("Please, send report and log file from %q to support@exonlab.omnidesk.ru", configDir)
			return
		}

		report = report + logTail
		if err := openSupportPageWithReport(report); err != nil {
	*/
	fmt.Printf("Fatal error: %v\n", sin)
	fmt.Printf("Please, send report and log file from %q to support@exonlab.omnidesk.ru", logDir)
	//	}
}

func appendOrCreate(filepath string, data string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return file.Close()
}

func readLogTail(logPath string) (string, error) {
	tailLen := int64(2048 - len(reportPreambula))
	stat, err := os.Stat(logPath)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	} else if os.IsNotExist(err) {
		return "", nil
	}
	size := stat.Size()
	if size < tailLen {
		logData, err := ioutil.ReadFile(logPath)
		return string(logData), err
	}
	logFile, err := os.Open(logPath)
	if err != nil {
		return "", err
	}
	defer logFile.Close()

	buf := make([]byte, tailLen)
	n, err := logFile.ReadAt(buf, size-tailLen)
	if err != nil {
		return "", err
	}
	buf = buf[:n]
	lines := strings.SplitN(string(buf), "\n", 2)
	if len(lines) == 2 {
		return lines[1], nil
	}
	return string(buf), nil
}

func openSupportPageWithReport(report string) error {
	supportURL, err := url.Parse("https://web.containerum.io/support")
	if err != nil {
		return err
	}
	query := supportURL.Query()
	query.Set("report", report)
	supportURL.RawQuery = query.Encode()
	return open.Run(supportURL.String())
}
