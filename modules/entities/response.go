package entities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadfarhankt/nft-marketplace/pkg/nftlogger"
)

type IResponse interface {
	Success(code int, data any) IResponse
	Error(code int, traceId, msg string) IResponse
	Res() error
}

type Response struct {
	StatusCode int
	Data       any
	ErrorRes   *ErrorResponse
	Context    *fiber.Ctx
	IsError    bool
}

type ErrorResponse struct {
	TraceId string `json:"trace_id"`
	Msg     string `json:"message"`
}

func NewResponse(ctx *fiber.Ctx) IResponse {
	return &Response{
		Context: ctx,
	}
}

func (r *Response) Success(code int, data any) IResponse {
	r.StatusCode = code
	r.Data = data
	nftlogger.InitNftLogger(r.Context, &r.Data).Print().Save()
	return r
}

func (r *Response) Error(code int, traceId, msg string) IResponse {
	r.StatusCode = code
	r.IsError = true
	r.ErrorRes = &ErrorResponse{
		TraceId: traceId,
		Msg:     msg,
	}
	nftlogger.InitNftLogger(r.Context, &r.ErrorRes).Print().Save()
	return r
}

func (r *Response) Res() error {
	return r.Context.Status(r.StatusCode).JSON(func() any {
		if r.IsError {
			return &r.ErrorRes
		}
		return r.Data
	}())
}
