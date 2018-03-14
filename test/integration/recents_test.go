package integration

import (
	"context"
	"testing"
	"time"

	"github.com/genert/pipedrive-api/pipedrive"
)

func TestRecents_List(t *testing.T) {
	sinceTime := time.Date(2017, time.September, 10, 10, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05")

	result, _, err := client.Recents.List(context.Background(), &pipedrive.RecentsListOptions{
		SinceTimestamp: sinceTime,
		Start:          0,
	})

	if err != nil {
		t.Errorf("Could not get recents: %v", err)
	}

	if result.Success != true {
		t.Error("Could not get successful recents response")
	}
}
