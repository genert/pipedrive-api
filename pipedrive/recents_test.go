package pipedrive

import (
	"testing"
	"time"
)

func TestRecentsService_List(t *testing.T) {
	if *apiIntegration {
		sinceTime := time.Date(2017, time.September, 10, 10, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")

		opt := &RecentsListOptions{
			SinceTimestamp: sinceTime,
			Start:          0,
		}

		recents, _, err := client.Recents.List(opt)

		if err != nil {
			t.Errorf("Could not get recents: %v", err)
		}

		if recents.Success != true {
			t.Error("Could not get successful recents response")
		}
	}
}
