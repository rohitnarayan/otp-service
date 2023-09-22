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

/**
createOTP system

createOTP API -> generate a OTP -> map to the userID provided -> insert into a otp table with createdAt timestamp
validateOTP API -> parse the req -> userID -> fetch the otp db -> if otp exists -> createdat + 2min < current time of the request

cache -> redis -> expiry for 120 seconds
*/
