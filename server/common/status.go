package common

const (
	// BusinessSuccessCode success
	BusinessSuccessCode = 200000
	// BusinessErrorCode internal exception
	BusinessErrorCode = 200101
	// JSONFormatErrorCode json marshl or unmarshl exception
	JSONFormatErrorCode = 200102
	// UnknownErrorCode Common exception
	UnknownErrorCode = 200201
	// DatabaseQueryErrorCode Fetch data from database exception
	DatabaseQueryErrorCode = 200202
	// JWTAuthFailedCode token mismatch
	JWTAuthFailedCode = 200401
	// TokenExpiredCode token expired
	TokenExpiredCode = 200402
	// ServerErrorCode Other exception
	ServerErrorCode = 500
)

var statusText = map[int]string{
	BusinessErrorCode:      "Mt internal system error ...",
	UnknownErrorCode:       "Mt Unknown Error",
	ServerErrorCode:        "Mt Server Error",
	BusinessSuccessCode:    "success",
	JSONFormatErrorCode:    "Mt json data format error",
	JWTAuthFailedCode:      "User authentication failure",
	TokenExpiredCode:       "Login expired, login again.",
	DatabaseQueryErrorCode: "Fetch data from database failed",
}

// StatusText returns a text for the business status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
