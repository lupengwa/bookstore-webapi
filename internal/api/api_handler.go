package api

import (
	"bookstore-webapi/internal/api/restutils"
	"net/http"
)

type Handler interface {
	GetRestUriToHandlerConfig() map[restutils.RestApiUriKey]http.HandlerFunc
}
