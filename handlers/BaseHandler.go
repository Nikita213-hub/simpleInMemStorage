package handlers

import "net/http"

type BaseHandler struct {
	Path    string
	Handler http.Handler
}
