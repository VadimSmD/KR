package main

import (
	"context"
	"net/http"

	"github.com/VadimSmD/KR/gen/oas"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --config ../../ogen.yml --target ../../gen/oas -package oas --clean ../../docs/openapi.yaml

var _ oas.Handler = (*Handler)(nil)

type Handler struct {
	oas.UserHandler
}

func (h *Handler) NewError(_ context.Context, err error) *oas.DefaultStatusCode {
	return &oas.DefaultStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: oas.UnsuccessfulResponse{
			Msg: err.Error(),
		},
	}
}
