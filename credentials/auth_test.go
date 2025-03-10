package credentials

import (
	"testing"
)

func TestGetAuthentication(t *testing.T) {
	tk := New("key", "secret")
	_, err := tk.GetAuthentication()
	if err != nil {
		t.Fatal(err)
	}
}
