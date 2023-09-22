package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type createOTPRequest struct {
	UserID string `json:"user_id"`
}

type createOTPSuccessResponse struct {
	UserID string `json:"user_id"`
	Otp    string `json:"otp"`
}

func (h *Handler) CreateOTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req, err := parseCreateOTPRequest(r)
	if err != nil {
		resp := NewErrorResponse(ErrInvalidRequestBody, "failed to parse request", "invalid request body")
		resp.Write(w, http.StatusBadRequest)
		return
	}

	if err := validateCreateOTPRequest(req); err != nil {
		resp := NewErrorResponse(ErrInvalidRequestBody, "no userID provided", "invalid request body")
		resp.Write(w, http.StatusBadRequest)
		return
	}

	// call to service layer
	otpModel, err := h.service.CreateOTP(ctx, req.UserID)
	if err != nil {
		resp := NewErrorResponse(ErrInternalServerError, err.Error(), "internal server error")
		resp.Write(w, http.StatusInternalServerError)
		return
	}

	data := createOTPSuccessResponse{
		Otp:    otpModel.Otp,
		UserID: otpModel.UserID,
	}

	resp := NewSuccessResponse(data)
	resp.Write(w, http.StatusOK)
	return
}

func parseCreateOTPRequest(r *http.Request) (*createOTPRequest, error) {
	var req *createOTPRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse createOTP request")
	}

	return req, nil
}

func validateCreateOTPRequest(req *createOTPRequest) error {
	if len(req.UserID) == 0 {
		return errors.Errorf("no userID provided")
	}

	return nil
}
