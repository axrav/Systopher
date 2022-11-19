package errors

import "fmt"

// Error is a custom error type

type Error struct {
	Code    string `json:"code"`
	Err     error  `json:"error"`
	Message string `json:"message,omitempty"`
}

// Error returns the error message
func (e *Error) Error() error {
	return e.Err
}

func (e *Error) Merror() Error {
	e.Message = e.Err.Error()
	return *e

}

// Code returns the error code
func (e *Error) ErrorCode() string {
	return e.Code
}

// Custom Errors

var (
	InvalidCred = Error{
		Code: "ERR-699",
		Err:  fmt.Errorf("invalid Credentials"),
	}

	InvalidData = Error{
		Code: "ERR-698",
		Err:  fmt.Errorf("invalid Data"),
	}

	InvalidToken = Error{
		Code: "ERR-697",
		Err:  fmt.Errorf("Unauthorized"),
	}

	InvalidUser = Error{
		Code: "ERR-696",
		Err:  fmt.Errorf("user doesn't exists"),
	}
	UsernameTaken = Error{
		Code: "ERR-695",
		Err:  fmt.Errorf("username already taken"),
	}
	InvalidEmail = Error{
		Code: "ERR-694",
		Err:  fmt.Errorf("invalid Email"),
	}

	EmailTaken = Error{
		Code: "ERR-693",
		Err:  fmt.Errorf("email already taken"),
	}

	InvalidUsername = Error{
		Code: "ERR-691",
		Err:  fmt.Errorf("invalid Username"),
	}

	InvalidServer = Error{
		Code: "ERR-690",
		Err:  fmt.Errorf("invalid Server"),
	}

	InternalServerError = Error{
		Code: "ERR-689",
		Err:  fmt.Errorf("internal Server Error"),
	}

	InvalidOtp = Error{
		Code: "ERR-688",
		Err:  fmt.Errorf("invalid Otp"),
	}

	NotFound = Error{
		Code: "ERR-687",
		Err:  fmt.Errorf("details Not Found"),
	}

	InvalidPassword = Error{
		Code: "ERR-686",
		Err:  fmt.Errorf("password not validated"),
	}

	AlreadyVerified = Error{
		Code: "ERR-685",
		Err:  fmt.Errorf("user already verified"),
	}
	AlreadyExists = Error{
		Code: "ERR-684",
		Err:  fmt.Errorf("already exists"),
	}

	NoResponse = Error{
		Code: "ERR-683",
		Err:  fmt.Errorf("no response from server"),
	}
)
