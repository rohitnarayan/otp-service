package handler

import "github.com/rohitnarayan/otp-service/internal/service"

type Handler struct {
	service service.OTPService
}

func NewHandler(service service.OTPService) *Handler {
	return &Handler{
		service: service,
	}
}
