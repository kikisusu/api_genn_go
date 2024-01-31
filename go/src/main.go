package main

import (
	"errors"
	"go/src/mypkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 全リクエストを処理
	router.Any("/*path", func(c *gin.Context) {

		// Httpメソッドの取得
		method := c.Request.Method
		// リクエストパスの取得
		path := c.Param("path")

		// API定義を取得
		apiOrcStruct, err := mypkg.Jsonconv()

		// API定義が取得できなかった場合はシステムエラーを応答
		if err != nil {
			c.JSON(500, gin.H{
				"message": "system error",
			})
			return
		}

		// API定義にマッチする定義がない場合は400エラーを応答
		outboundList, err := apimatch(apiOrcStruct, method, path)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "想定外のリクエストです。",
			})
			return
		}

		// Outbound応答の格納域
		var outboundcon []*http.Response

		// outboundに定義されているAPIをコール
		for _, outbound := range outboundList {
			obMethod := outbound.OBmethod
			obEndpoint := outbound.OBendpoint
			obResp, err := callAPI(obMethod, obEndpoint)
			if err != nil {
				c.JSON(400, gin.H{
					"message": "Outbound実行で失敗しました。",
				})
			}
			outboundcon = append(outboundcon, obResp)
		}

		factory := &mypkg.HandlerFactory{}
		handler := factory.CreateHandler(method, path, outboundcon)
		status, body := handler.Handle(c, outboundcon)

		c.JSON(status, body)
	})

	router.Run(":3001")
}

func apimatch(apiList mypkg.APIList, method, path string) ([]mypkg.Outbound, error) {

	var outboundList []mypkg.Outbound
	var err error

	// apiListの要素数分ループしてmethodとendpointを取得する
	for _, api := range apiList.APIList {

		if method == api.Method && path == api.Endpoint {
			outboundList = api.Outbound
		}
	}

	if outboundList == nil {
		err = errors.New("一致する定義がありません。")
	}

	return outboundList, err
}

func callAPI(method, endpoint string) (*http.Response, error) {
	var resp *http.Response
	var err error

	client := &http.Client{}

	req, _ := http.NewRequest(method, endpoint, strings.NewReader(""))
	resp, err = client.Do(req)

	return resp, err
}
