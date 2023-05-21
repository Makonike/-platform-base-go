package aosdk

import "testing"

func TestCreateToken(t *testing.T) {
	InitConfig("http://localhost:8080", "v2")
	resp, err := CreateTokens(TokenInfo{
		BoxUUID:    "wwwww",
		ServiceIds: []string{"100w01"},
		Sign:       "",
	}, "e9993fc787d94b6c886cbaa340f9c0f4")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Error == nil {
		t.Errorf("error is nil, want not nil")
	}

	resp, err = CreateTokens(TokenInfo{
		BoxUUID:    "wwwww",
		ServiceIds: []string{"10001"},
		Sign:       "",
	}, "e9993fc787d94b6c886cbaa340f9c0f4")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Error != nil {
		t.Errorf("error is not nil(%v), want nil", resp.Error)
	}

}
