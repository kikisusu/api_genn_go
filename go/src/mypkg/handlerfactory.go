package mypkg

import (
	"go/src/mypkg/api/activeuuid"
	"go/src/mypkg/api/marvelbuzz"
	"net/http"
)

type HandlerFactory struct{}

func (f *HandlerFactory) CreateHandler(method string, path string, outboundcon []*http.Response) ResponseHandler {
	switch method {
	case "GET":
		switch path {
		case "/marvelbuzz":
			return &marvelbuzz.HandlerMarvelBuzz{}
		case "/activeuuid":
			return &activeuuid.HandlerActiveUuid{}
		default:
			return nil
		}
	default:
		return nil
	}
}
