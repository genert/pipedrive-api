package pipedrive

import "testing"

func TestActivityTypesService_Add(t *testing.T) {
	result, _, err := client.ActivityTypes.Add(&ActivityTypesAddOptions{
		Name:    "test",
		IconKey: "email",
	})

	if err != nil {
		t.Errorf("Could not create activity type: %v", err)
	}

	if result.Success != true {
		t.Error("Could not create activity type successfully")
	}
}

func TestActivityTypesService_Delete(t *testing.T) {

}

func TestActivityTypesService_DeleteMultiple(t *testing.T) {

}

func TestActivityTypesService_Edit(t *testing.T) {

}

func TestActivityTypesService_List(t *testing.T) {

}
