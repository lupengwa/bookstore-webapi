package restutils

type RestApiUriKey struct {
	HttpMethod string
	Path       string
}

// BookStoreErrorResponse is the error when the bookStore encounters a user input error.
type BookStoreErrorResponse struct {
	httpStatusCode int
	// The error message
	Msg string `json:"message"`
}
