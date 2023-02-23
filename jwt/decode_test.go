package jwt

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDecode(t *testing.T) {

	tests := []struct {
		name    string
		token   string
		want    [2]string
		wantErr bool
	}{
		{
			name:  "should return",
			token: "eyJraWQiOiJQZ1c5K3ZFWlRoZG0wN1pDZW5SSEIxaXh4MXA2dDBjeDZYdGdOSGtLbkV3PSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiI5NmFkMzRjNS0zZWM3LTRiNzMtOTg1YS0yMWE2OWZlMGEzOWEiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmFwLXNvdXRoZWFzdC0xLmFtYXpvbmF3cy5jb21cL2FwLXNvdXRoZWFzdC0xX0FYcUNacml4NCIsImNvZ25pdG86dXNlcm5hbWUiOiI5NmFkMzRjNS0zZWM3LTRiNzMtOTg1YS0yMWE2OWZlMGEzOWEiLCJnaXZlbl9uYW1lIjoiZGsiLCJvcmlnaW5fanRpIjoiYmMzYTkzYjMtYzljYS00NjMxLWIxNjYtOTYzOTRmOWMzZDlkIiwiYXVkIjoiMTRnNGJra2FsYm42bWFybnJsdGxwYWsxbmsiLCJldmVudF9pZCI6ImY0YjJmZWE0LWM5NDItNGQ4Mi04MDYxLWM2Y2M0YjRjNGY3YyIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNjc2OTY3MjQwLCJleHAiOjE2NzY5NzA4NDAsImlhdCI6MTY3Njk2NzI0MCwiZmFtaWx5X25hbWUiOiJsYW0gZ2EiLCJqdGkiOiJhYjFkOWE4ZC1mNjhlLTRlNDItOTQ5Ny1jYjQxY2E0ZWVmNWIiLCJlbWFpbCI6ImRrQGdtYWlsLmNvbSJ9.vRt3U1koxp8jHy4WmcpEKeF-ixFq3wi2k9JaJ6Zw2daTIGJcnjvmjlRrF40TKYcZ62YnI2iHtJBlM6maVSh10RlQzZgbeTe7iU68SndZQj5Ddd7eEREHb1dHpQyvTE7k6m0o9qYLAI8oHqrvKPIqZoRuId0P6ZKsBZkOhmB3XFp3OD5mW4pGK35tMcrIhurfu3FYhKCxbSWy2jbbbh7JFZvoXSJJWPZNQxpNlIlzRVSZOov3h0CsNJlzBwDXiHen82LP1_o-PwU8WAxjC1lcPg0mu4Ex4lKnxnxp8NnJoTKUm0xOm9VWdJlQuuOqYN3PoIIT1-SQRJ3SE2DTMgfV3A",
			want: [2]string{`{
				"alg": "RS256",
				"kid": "PgW9+vEZThdm07ZCenRHB1ixx1p6t0cx6XtgNHkKnEw="
			}`, `{
				"aud": "14g4bkkalbn6marnrltlpak1nk",
				"auth_time": 1676967240,
				"cognito:username": "96ad34c5-3ec7-4b73-985a-21a69fe0a39a",
				"email": "dk@gmail.com",
				"email_verified": true,
				"event_id": "f4b2fea4-c942-4d82-8061-c6cc4b4c4f7c",
				"exp": 1676970840,
				"family_name": "lam ga",
				"given_name": "dk",
				"iat": 1676967240,
				"iss": "https://cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_AXqCZrix4",
				"jti": "ab1d9a8d-f68e-4e42-9497-cb41ca4eef5b",
				"origin_jti": "bc3a93b3-c9ca-4631-b166-96394f9c3d9d",
				"sub": "96ad34c5-3ec7-4b73-985a-21a69fe0a39a",
				"token_use": "id"
			}`},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf(cmp.Diff(got, tt.want))
			}
		})
	}
}
