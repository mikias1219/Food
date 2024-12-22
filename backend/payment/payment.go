package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ChapaPaymentRequest represents the data structure for a payment request
type ChapaPaymentRequest struct {
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	TxRef          string `json:"tx_ref"`
	CallbackURL    string `json:"callback_url"`
	Customizations struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Logo        string `json:"logo"`
	} `json:"customizations"`
}

// ChapaPaymentResponse represents the response from Chapa's API
type ChapaPaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}

// InitiatePayment sends a payment request to Chapa
func InitiatePayment(apiKey string, paymentRequest ChapaPaymentRequest) (string, error) {
	url := "https://api.chapa.co/v1/transaction/initialize"

	// Serialize request data to JSON
	requestBody, err := json.Marshal(paymentRequest)
	if err != nil {
		return "", fmt.Errorf("failed to serialize payment request: %w", err)
	}

	// Make the HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send payment request: %w", err)
	}
	defer resp.Body.Close()

	// Parse the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var chapaResponse ChapaPaymentResponse
	if err := json.Unmarshal(responseBody, &chapaResponse); err != nil {
		return "", fmt.Errorf("failed to parse Chapa response: %w", err)
	}

	if chapaResponse.Status != "success" {
		return "", fmt.Errorf("payment initialization failed: %s", chapaResponse.Message)
	}

	// Return the checkout URL
	return chapaResponse.Data.CheckoutURL, nil
}
