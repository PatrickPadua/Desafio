package lib

import (
	"encoding/json"
	b64 "encoding/base64"
)

type ConsulKVResponse struct{
	LockIndex int64 `json:"LockIndex"`
    Key string `json:"Key"`
    Flags int64 `json:"Flags"`
    Value string `json:"Value"`
    CreateIndex int64 `json:"CreateIndex"`
    ModifyIndex int64 `json:"ModifyIndex"`
}


func ConsulRequest(endpoint, kv string) string{

	url := endpoint + kv

	response := Request("consul", url, "GET")

	var resp []ConsulKVResponse

	json.Unmarshal(response, &resp)

	byteHash, _ := b64.StdEncoding.DecodeString(resp[0].Value)

	result := string(byteHash)

	return result
}