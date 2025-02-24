package services

import (
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/json"
    "encoding/pem"
    "errors"
    "fmt"
    "github.com/go-resty/resty/v2"
    "strings"
)

type TeleBirrService struct {
    appID     string
    appKey    string
    publicKey string
    baseURL   string
}

func NewTeleBirrService(appID, appKey, publicKey, baseURL string) *TeleBirrService {
    return &TeleBirrService{
        appID:     appID,
        appKey:    appKey,
        publicKey: publicKey,
        baseURL:   baseURL,
    }
}

// InitiatePayment initiates a payment via TeleBirr
func (s *TeleBirrService) InitiatePayment(amount float64, phoneNumber, callbackURL string) (string, error) {
    client := resty.New()

    // Prepare the request payload
    payload := map[string]interface{}{
        "app_id":     s.appID,
        "app_key":    s.appKey,
        "amount":     amount,
        "phone":     phoneNumber,
        "callback":  callbackURL,
        "timestamp": fmt.Sprintf("%d", time.Now().Unix()),
    }

    // Encrypt the payload
    encryptedPayload, err := s.encryptPayload(payload)
    if err != nil {
        return "", err
    }

    // Send the request to TeleBirr
    resp, err := client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(map[string]string{"data": encryptedPayload}).
        Post(s.baseURL + "/api/payment/initiate")
    if err != nil {
        return "", err
    }

    // Decrypt the response
    var response map[string]interface{}
    if err := json.Unmarshal(resp.Body(), &response); err != nil {
        return "", err
    }

    if response["code"] != "200" {
        return "", errors.New(response["message"].(string))
    }

    return response["transaction_id"].(string), nil
}

// encryptPayload encrypts the payload using TeleBirr's public key
func (s *TeleBirrService) encryptPayload(payload map[string]interface{}) (string, error) {
    // Convert payload to JSON
    payloadJSON, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    // Decode the public key
    block, _ := pem.Decode([]byte(s.publicKey))
    if block == nil {
        return "", errors.New("failed to decode public key")
    }

    // Parse the public key
    pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return "", err
    }

    // Encrypt the payload using RSA
    encrypted, err := rsa.EncryptPKCS1v15(nil, pubKey.(*rsa.PublicKey), payloadJSON)
    if err != nil {
        return "", err
    }

    // Encode the encrypted payload in base64
    return base64.StdEncoding.EncodeToString(encrypted), nil
}