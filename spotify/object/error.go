package object

// AuthError represents AuthenticationErrorObject
// Link: https://developer.spotify.com/documentation/web-api/#response-schema
type AuthError struct {
	Error     string `json:"error"`
	ErrorDesc string `json:"error_description"`
}

// Error represents ErrorObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-errorobject
type Error struct {
	Error errorInner `json:"error"`
}

type errorInner struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
