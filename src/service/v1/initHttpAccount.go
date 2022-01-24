// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import "net/http"

func (s *Service) initAccountApi() {
	// ********************
	// Account API Endpoint
	// ********************

	// Enable API service:
	// curl -X POST  -H "auth: KEY"  http://localhost/api/on
	http.HandleFunc("/api/on", s.switchAPIOn)

	// Disable API service:
	// curl -X POST  -H "auth: KEY"  http://localhost/api/off
	http.HandleFunc("/api/off", s.switchAPIOff)

	// Create API
	// curl -X POST  -H "auth: KEY" -H 'Content-Type: application/json' -d "{\"id\": \"\", \"accountName\": \"accountName\", \"key\": \"key\", \"secret\": \"secret\"}" http://localhost/api/create
	http.HandleFunc("/api/create", s.createApiHandler)

	// Set API leverage
	// curl -H "auth: KEY"  -H 'Content-Type: application/json' -d "{\"api_id\": \"API_ID\", \"market\": \"ETHUSD\", \"leverage\": 3}" http://localhost/api/setleverage
	http.HandleFunc("/api/setleverage", s.setApiLeverageHandler)

	// Reset API leverage
	// curl -H "auth: KEY"  -H 'Content-Type: application/json' -d "{\"api_id\": \"API_ID\", \"market\": \"ETHUSD\", \"leverage\": 0}" http://localhost/api/resetleverage
	http.HandleFunc("/api/resetleverage", s.resetApiLeverageHandler)

	// Delete API
	// curl -H "auth: KEY"  http://localhost/api/delete?id=API-ID
	http.HandleFunc("/api/delete", s.deleteApiHandler)

}
