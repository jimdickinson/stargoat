package stargoat

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	astraDatabaseBaseURL, ok := os.LookupEnv("ASTRA_DB_BASE_URL")
	if !ok {
		t.Log("Missing ASTRA_DB_BASE_URL env needed for test run")
		t.FailNow()
	}

	type args struct {
		creds   CredsProvider
		authURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				creds: &EnvVarCredsProvider{
					UsernameEnv: "ASTRA_DB_USERNAME",
					PasswordEnv: "ASTRA_DB_PASSWORD",
				},
				authURL: astraDatabaseBaseURL + "/api/rest/v1/auth",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Authenticate(tt.args.creds, tt.args.authURL)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Len(t, got.AuthToken, 36)
		})
	}
}
