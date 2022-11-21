package Login

import "testing"

type credentials struct {
	userName string
	password string
}

func TestLogin(t *testing.T) {
	var tests = []struct {
		userName string
		Password string
		expected bool
	}{
		{"TestUser", "TestPassword", true},
		{"FakeUser", "FakePassword", false},
		{"testuser", "testpassword", false},
		{"PeterParker", "Manhattan", true},
		{"PETERpaRKer", "Manhattan", false},
		{"PeterParker", "MANHATTAn", false}}
}
