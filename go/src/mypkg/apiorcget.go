package mypkg

import (
	"encoding/json"
	"log"
	"os"
)

type Outbound struct {
	OBmethod   string `json:"obMethod"`
	OBendpoint string `json:"obEndpoint"`
}

type API struct {
	Apiid    string     `json:"appid"`
	Method   string     `json:"method"`
	Endpoint string     `json:"endpoint"`
	Outbound []Outbound `json:"outboud"`
}

type APIList struct {
	APIList []API `json:"apiList"`
}

func Jsonconv() (APIList, error) {
	// JSONファイルを読み込む
	bytes, err := os.ReadFile("/go/src/go/src/config/apiOrc.json")
	if err != nil {
		log.Fatal(err)
	}

	// JSONを構造体に変換する
	var apiList APIList
	if err := json.Unmarshal(bytes, &apiList); err != nil {
		log.Fatal(err)
	}

	return apiList, err
}
