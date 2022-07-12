package stackoverflow_go

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func trackStats() error {
	const perm = 0600

	fd, err := os.OpenFile("/tmp/stackoverflow-go/stats.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_SYNC|os.O_TRUNC, perm)
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
	ticker := time.NewTicker(time.Millisecond)
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

func recursive() {
	recursive()
}

func TestStackOverflow(t *testing.T) {
	go func() {
		if err := trackStats(); err != nil {
			fmt.Println(err)
			t.Fatal(err)
		}
	}()

	time.Sleep(time.Second)
	recursive()
}
