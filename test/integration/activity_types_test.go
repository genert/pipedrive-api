package integration

import (
	"context"
	"testing"

	"github.com/genert/pipedrive-api/pipedrive"
)

func TestActivityTypesService_Create(t *testing.T) {
	result, _, err := client.ActivityTypes.Create(context.Background(), &pipedrive.ActivityTypesAddOptions{
		Name:    RandomString(13),
		IconKey: "email",
	})

	if err != nil {
		t.Errorf("Could not create activity type: %v", err)
	}

	if result.Success != true {
		t.Error("Could not create activity type successfully")
	}
}
