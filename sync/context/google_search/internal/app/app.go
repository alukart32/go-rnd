package app

import (
	"net/http"

	"alukart32.com/usage/context/config"
	"alukart32.com/usage/context/pkg/httpx"
)

func Run(cfg config.Config) error {
	return http.ListenAndServe(cfg.Http.GetUrl(), httpx.New())
}
