package server

import (
	"github.com/gorilla/mux"
	"github.com/rohitnarayan/otp-service/internal/handler"
)

func router(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/internal/otp", h.CreateOTP)
	return r
}
