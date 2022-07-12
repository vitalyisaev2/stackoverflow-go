package stackoverflow_go

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func trackStats(t *testing.T) {
	ticker := time.NewTicker(10 * time.Microsecond)
	defer ticker.Stop()

	for range ticker.C {
		report, err := getMemoryReportOwn()
		require.NoError(t, err)

		fmt.Println(report)
	}
}

func recursive() {
	recursive()
}

func TestStackOverflow(t *testing.T) {
	go trackStats(t)

	recursive()
}
