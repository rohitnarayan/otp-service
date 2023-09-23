package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rohitnarayan/otp-service/internal/config"
	"github.com/rohitnarayan/otp-service/internal/handler"
	"github.com/rohitnarayan/otp-service/internal/postgres"
	"github.com/rohitnarayan/otp-service/internal/service"
	"github.com/rohitnarayan/otp-service/internal/store"
)

func Server() {
	db, err := postgres.NewDB(config.App.Database.Postgres)
	if err != nil {
		log.Fatalf("failed to init db")
	}

	otpStore := store.NewStore(db, db)
	otpService := service.NewOTPService(otpStore)
	h := handler.NewHandler(otpService)

	r := router(h)
	http.Handle("/", r)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.App.Server.Port), r); err != nil {
		log.Fatalf("error occurred while starting the server")
	}
}
