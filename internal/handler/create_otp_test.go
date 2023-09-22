package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/rohitnarayan/otp-service/internal/service"
	"github.com/rohitnarayan/otp-service/internal/store"
)

func TestCreateOTP(t *testing.T) {

	dummyOTPModel := &store.OTPModel{
		Otp:       "3451",
		UserID:    "Rohit",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Status:    "created",
	}

	tests := []struct {
		name        string
		body        io.Reader
		mockService func(ctrl *gomock.Controller) *service.MockOTPService
	}{
		{
			name: "success - create OTP",
			body: getSuccessRequestBody(),
			mockService: func(ctrl *gomock.Controller) *service.MockOTPService {
				s := service.NewMockOTPService(ctrl)
				s.EXPECT().CreateOTP(gomock.Any(), gomock.Any()).Return(dummyOTPModel, nil)
				return s
			},
		},
	}

	otpResp := &Response{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			handler := &Handler{service: tt.mockService(ctrl)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", tt.body)
			handler.CreateOTP(w, r)
			fmt.Println(w)

			body, err := io.ReadAll(w.Body)
			assert.NoError(t, err)

			err = json.Unmarshal(body, otpResp)
			assert.NoError(t, err)

			otpData := otpResp.Data.(map[string]interface{})

			assert.Equal(t, "3451", otpData["otp"])
			assert.Equal(t, "Rohit", otpData["user_id"])

		})
	}
}

func getSuccessRequestBody() io.Reader {
	request := `{"user_id":"Rohit"}`
	return strings.NewReader(request)
}
