package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ccClient struct {
	client *http.Client
	apiKey string
}

type PriceHistoryResponse struct {
	Response   string `json:"Response"`
	Message    string `json:"Message"`
	HasWarning bool   `json:"HasWarning"`
	Type       int    `json:"Type"`
	RateLimit  struct {
	} `json:"RateLimit"`
	Data struct {
		Aggregated bool `json:"Aggregated"`
		TimeFrom   int  `json:"TimeFrom"`
		TimeTo     int  `json:"TimeTo"`
		Data       []struct {
			Time             int     `json:"time"`
			High             float64 `json:"high"`
			Low              float64 `json:"low"`
			Open             float64 `json:"open"`
			Volumefrom       float64 `json:"volumefrom"`
			Volumeto         float64 `json:"volumeto"`
			Close            float64 `json:"close"`
			ConversionType   string  `json:"conversionType"`
			ConversionSymbol string  `json:"conversionSymbol"`
		} `json:"Data"`
	} `json:"Data"`
}

const BASE_URL = "https://min-api.cryptocompare.com/data/v2/histoday"
const TO_CURRENCY = "USD"

func (cc *ccClient) fetchPriceData(ticker string) (*[]PriceData, error) {
	url := fmt.Sprintf("%s?fsym=%s&tsym=%s&limit=400", BASE_URL, ticker, TO_CURRENCY)
	// Construct the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")

	// Execute the request
	resp, err := cc.client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Println("Error creating gzip reader:", err)
			return nil, err
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	// Read the response body
	body, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	// Print the response body
	// fmt.Println("Response:", string(body))
	var rawData PriceHistoryResponse
	err = json.Unmarshal(body, &rawData)
	if err != nil {
		return nil, err
	}
	var priceData []PriceData

	for _, row := range rawData.Data.Data {
		priceData = append(priceData, PriceData{
			Timestamp: int64(row.Time),
			PriceUSD:  fmt.Sprintf("%.10f", row.Close),
		})
	}

	return &priceData, nil
}

func NewCCClient(apiKey string) *ccClient {
	return &ccClient{
		client: &http.Client{},
		apiKey: apiKey,
	}
}
