//go:build linux
// +build linux

package stackoverflow_go

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/c2h5oh/datasize"
	"github.com/pkg/errors"
)

// explanation: https://man7.org/linux/man-pages/man5/proc.5.html
type memoryReport struct {
	vmRSS uint64
	vmStk uint64
}

func getMemoryReportOwn() (*memoryReport, error) {
	pid := os.Getpid()

	report, err := getMemoryReportByPID(pid)
	if err != nil {
		return nil, errors.Wrap(err, "get memory report by PID")
	}

	return report, nil
}

func getMemoryReportByPID(pid int) (*memoryReport, error) {
	path := fmt.Sprintf("/proc/%d/status", pid)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}

	split := strings.Split(string(data), "\n")

	out := &memoryReport{}

	for _, line := range split {
		if strings.HasPrefix(line, "VmRSS") {
			out.vmRSS, err = extractValue(line)

			if err != nil {
				return nil, errors.Wrap(err, "extract value")
			}
		}

		if strings.HasPrefix(line, "VmStk") {
			out.vmStk, err = extractValue(line)

			if err != nil {
				return nil, errors.Wrap(err, "extract value")
			}
		}
	}

	return out, nil
}

func extractValue(line string) (uint64, error) {
	split := strings.Split(line, ":")
	src := strings.TrimLeft(split[1], " \t")

	var dst datasize.ByteSize

	err := dst.UnmarshalText([]byte(src))
	if err != nil {
		return 0, errors.Wrapf(err, "unmarshal text '%s'", src)
	}

	return dst.Bytes(), nil
}
