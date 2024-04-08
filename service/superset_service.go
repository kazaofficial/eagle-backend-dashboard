package service

import (
	"bytes"
	"eagle-backend-dashboard/dto"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type SupersetServiceImpl struct {
}

func NewSupersetService() SupersetService {
	return &SupersetServiceImpl{}
}

func (service *SupersetServiceImpl) AuthTokenRequest(request *dto.SupersetAuthRequest) (*dto.SupersetTokenResponse, error) {
	client := &http.Client{}
	auth, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	byteReq := bytes.NewReader(auth)
	req, err := http.NewRequest(http.MethodPost, os.Getenv("SUPERSET_HOST")+"/v1/security/login", byteReq)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response dto.SupersetTokenResponse
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (service *SupersetServiceImpl) GetCsrfToken(accessToken string) (string, string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, os.Getenv("SUPERSET_HOST")+"/v1/security/csrf_token", nil)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}
	var response dto.SupersetCsrfResponse
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return "", "", err
	}
	return response.Result, res.Header.Get("Set-Cookie"), nil
}

func (service *SupersetServiceImpl) GetGuestToken(accessToken string, csrfToken string, cookie string, request *dto.SupersetGuestTokenRequest) (string, error) {
	client := &http.Client{}
	reqJson, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	byteReq := bytes.NewReader(reqJson)
	req, err := http.NewRequest(http.MethodPost, os.Getenv("SUPERSET_HOST")+"/v1/security/guest_token", byteReq)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Csrftoken", csrfToken)
	req.Header.Set("Cookie", cookie)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var response dto.SupersetGuestTokenResponse
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return "", err
	}
	return response.Token, nil
}
