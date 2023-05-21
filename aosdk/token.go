package aosdk

import (
	"encoding/json"
	"time"
)

type TokenResult struct {
	ServiceId string    `json:"serviceId"`
	BoxRegKey string    `json:"boxRegKey"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type AuthResponse struct {
	BoxUUID      string         `json:"boxUUID"`
	TokenResults []TokenResult  `json:"tokenResults"`
	Error        *ErrorResponse `json:"error"`
}

type TokenInfo struct {
	BoxUUID    string   `json:"boxUUID"`
	ServiceIds []string `json:"serviceIds"`
	Sign       string   `json:"sign"`
}

func CreateTokens(ti TokenInfo, requestId string) (AuthResponse, error) {
	rawBytes, err := DoPost("platform/auth/box_reg_keys", requestId, "", nil, ti)
	if err != nil {
		return AuthResponse{}, err
	}
	var resp AuthResponse
	err = json.Unmarshal(rawBytes, &resp)
	if err != nil {
		return AuthResponse{}, err
	}
	if len(resp.TokenResults) == 0 {
		err = json.Unmarshal(rawBytes, &resp.Error)
		if err != nil {
			return AuthResponse{}, err
		}
	}
	return resp, nil
}
