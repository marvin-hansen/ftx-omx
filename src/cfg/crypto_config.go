// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package cfg

// Read Readme for generating master & API keys
const (
	// masterKey decrypts DB data. Base64 encoded. Will be auto-decoded internally before usage.
	masterKey = ""

	// restApiAuthKey  authenticates clients to the API-API.  Base64 encoded. Will be auto-decoded.
	// The API-API is switched off by default and must be enabled before usage.
	// Any request must use the restApiAuthKey as auth header.
	restApiAuthKey = ""

	// restOrderAuthKey authenticates clients to the order API. Base64 encoded. Will be auto-decoded.
	// The order API is enabled by default and expects the restOrderAuthKey as auth header with each request.
	restOrderAuthKey = ""

	// SeedRandomNumbGen  Seeds the pseudo number generator to ensure unique values underlying newly generated API keys.
	// When switched off during development, you get same, non-random, keys after each restart which simplifies testing.
	// Note that non-random keys also allow anyone else using this code to brute-force-guess your actual API keys therefore,
	// seed is switched on by default to ensure safe & sane production.
	SeedRandomNumbGen = true
)

func GetMasterKey() string {
	return masterKey
}

func GetApiAuthKey() string {
	return restApiAuthKey
}

func GetOrderAuthKey() string {
	return restOrderAuthKey
}
