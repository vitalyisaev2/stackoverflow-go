//go:build linux
// +build linux

package stackoverflow_go

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type memoryReport struct {
	vmRss   uint64
	vmStack uint64
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
	path := fmt.Sprintf("/proc/%d/stat", pid)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}

	split := strings.Split(string(data), "\n")

	fmt.Println(split)

	return nil, nil
}
