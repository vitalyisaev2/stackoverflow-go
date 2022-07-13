//go:build linux
// +build linux

package stackoverflow_go

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

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

func trackStats(name string) error {
	const perm = 0600

	filename := fmt.Sprintf("/tmp/stackoverflow-go/%s.csv", name)

	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_SYNC|os.O_TRUNC, perm)
	if err != nil {
		return errors.Wrap(err, "open file")
	}

	wr := csv.NewWriter(fd)

	defer func() {
		errClose := fd.Close()
		if errClose != nil {
			fmt.Println(errClose)
		}
	}()

	err = wr.Write([]string{"Step", "VmRSS", "VmStk"})
	if err != nil {
		return errors.Wrap(err, "write header")
	}

	wr.Flush()

	if wr.Error() != nil {
		return errors.Wrap(wr.Error(), "flush")
	}

	step := 0

	// ask Linux' procfs for stats and dump it to file
	ticker := time.NewTicker(10 * time.Microsecond)
	defer ticker.Stop()

	for range ticker.C {
		report, err := getMemoryReportOwn()
		if err != nil {
			return errors.Wrap(err, "get memory report own")
		}

		err = wr.Write([]string{
			fmt.Sprint(step),
			fmt.Sprint(report.vmRSS),
			fmt.Sprint(report.vmStk),
		})
		if err != nil {
			return errors.Wrap(err, "write row")
		}

		wr.Flush()
		if wr.Error() != nil {
			return errors.Wrap(wr.Error(), "flush")
		}

		step++
	}

	return nil
}
