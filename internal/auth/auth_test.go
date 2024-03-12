package auth

import "testing"

func TestGetAPIKey(t *testing.T) {

	var headers = make(map[string][]string)
	headers["Authorization"] = []string{""}

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Errorf("Expected err, got nil")
	}
}
