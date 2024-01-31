package marvelbuzz

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerMarvelBuzz struct{}

type MarvelResBodyStruct struct {
	DaysUntil           int `json:"days_until"`
	FollowingProduction struct {
		DaysUntil   int    `json:"days_until"`
		ID          int    `json:"id"`
		Overview    string `json:"overview"`
		PosterURL   string `json:"poster_url"`
		ReleaseDate string `json:"release_date"`
		Title       string `json:"title"`
		Type        string `json:"type"`
	} `json:"following_production"`
	ID          int    `json:"id"`
	Overview    string `json:"overview"`
	PosterURL   string `json:"poster_url"`
	ReleaseDate string `json:"release_date"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type BuzzResBodyStruct struct {
	Phrase string `json:"phrase"`
}

func (h *HandlerMarvelBuzz) Handle(c *gin.Context, outboundcon []*http.Response) (int, interface{}) {

	// それぞれの応答データを取得
	marvelresbody, _ := io.ReadAll(outboundcon[0].Body)
	buzzresbody, _ := io.ReadAll(outboundcon[1].Body)

	// jsonを構造体にキャスト
	var marvelresbodystruct MarvelResBodyStruct
	if err := json.Unmarshal(marvelresbody, &marvelresbodystruct); err != nil {
		log.Fatal(err)
	}

	// jsonを構造体にキャスト
	var buzzresbodystruct BuzzResBodyStruct
	if err := json.Unmarshal(buzzresbody, &buzzresbodystruct); err != nil {
		log.Fatal(err)
	}

	// 取得したそれぞれの構造体から必要なデータのみレスポンス
	return http.StatusOK, gin.H{
		"marvelTitle": marvelresbodystruct.Title,
		"buzzPhrase":  buzzresbodystruct.Phrase,
	}
}
