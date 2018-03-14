package integration

import (
	"testing"
)

func TestFilesService_GetDownloadLinkByID(t *testing.T) {
	result, _, err := client.Files.GetDownloadLinkByID(1)

	if err != nil {
		t.Errorf("Could not get result: %v", err)
	}

	if result != "https://api.pipedrive.com/v1/files/1/download" {
		t.Error("Got invalid download link")
	}
}
