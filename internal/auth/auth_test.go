package auth

import "testing"

func TestGetAPIKey(t *testing.T) {

	var headers map[string][]string
	headers = make(map[string][]string)
	headers["Authorization"] = []string{""}

	_, error := GetAPIKey(headers)
	if error == nil {
		t.Errorf("Expected error, got nil")
	}
}
