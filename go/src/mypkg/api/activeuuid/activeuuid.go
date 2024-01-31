package activeuuid

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerActiveUuid struct{}

type ActiveResBodyStruct struct {
	Activity      string `json:"activity"`
	Type          string `json:"type"`
	Participants  int    `json:"participants"`
	Price         int    `json:"price"`
	Link          string `json:"link"`
	Key           string `json:"key"`
	Accessibility int    `json:"accessibility"`
}

type UuidResBodyStruct []string

func (h *HandlerActiveUuid) Handle(c *gin.Context, outboundcon []*http.Response) (int, interface{}) {

	// それぞれの応答データを取得
	activeresbody, _ := io.ReadAll(outboundcon[0].Body)
	uuidresbody, _ := io.ReadAll(outboundcon[1].Body)

	// jsonを構造体にキャスト
	var activeresbodystruct ActiveResBodyStruct
	if err := json.Unmarshal(activeresbody, &activeresbodystruct); err != nil {
		log.Fatal(err)
	}

	// jsonを構造体にキャスト
	var uuidresbodystruct UuidResBodyStruct
	if err := json.Unmarshal(uuidresbody, &uuidresbodystruct); err != nil {
		log.Fatal(err)
	}

	// 取得したそれぞれの構造体から必要なデータのみレスポンス
	return http.StatusOK, gin.H{
		"activity": activeresbodystruct.Activity,
		"uuid":     uuidresbodystruct[0],
	}
}
