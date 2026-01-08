package internal_test

import (
	"os"
	"testing"

	"github.com/nerdsec/goaes/internal"
)

const totalIterationTests = 100

func TestNewDEK(t *testing.T) {
	for range totalIterationTests {
		dek, err := internal.NewDEK()
		if err != nil {
			t.Errorf("failed to create dek. error: %v", err)
		}

		if len(dek) < 32 {
			t.Errorf("dek too small")
		}
	}
}

func TestNewSalt(t *testing.T) {
	for range totalIterationTests {
		salt, err := internal.NewSalt()
		if err != nil {
			t.Errorf("failed to create salt. error: %v", err)
		}

		if len(salt) < 32 {
			t.Errorf("salt too small")
		}
	}
}

func TestNewKEKFromEnvB64(t *testing.T) {
	tests := []struct {
		name             string
		passphraseEnvVar string
		passphrase       string
		salt             internal.Salt
		wantErr          bool
	}{
		{
			name:             "Valid base64",
			passphraseEnvVar: "GOAES_PASSPHRASE",
			passphrase:       "dJyHOdMbG94EMvQGQrs6YZiXGiAGQgDYtx6+eqLufQg=",
			salt:             []byte("kD+tNSxjss1XchcyyrKJyZBGg2mdmhh/IO3I87WW2Ds="),
			wantErr:          false,
		},
		{
			name:             "Invalid passphrase base64",
			passphraseEnvVar: "GOAES_PASSPHRASE",
			passphrase:       "dJyHOdMbG94EMvQGQrs6YZiXGiAGQgDYtx6eqLufQg=",
			salt:             []byte("kD+tNSxjss1XchcyyrKJyZBGg2mdmhh/IO3I87WW2Ds="),
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.Setenv(tt.passphraseEnvVar, tt.passphrase)
			if err != nil {
				t.Fatal("failed to set env var")
			}

			_, gotErr := internal.NewKEKFromEnvB64(tt.passphraseEnvVar, tt.salt)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewKEKFromEnvB64() failed: %v", gotErr)
				}

				return
			}

			if tt.wantErr {
				t.Fatal("NewKEKFromEnvB64() succeeded unexpectedly")
			}
		})
	}
}
