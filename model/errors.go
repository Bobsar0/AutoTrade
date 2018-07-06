package model

// General errors.
const (
	ErrNilSessionStruct = Error("Nil Session Struct")
	ErrUnauthorized           = Error("unauthorized")
	ErrInternal               = Error("internal error")
	ErrUserNotFound           = Error("user not found")
	ErrTransactionNotFound    = Error("transaction not found")
	ErrExchNotFound           = Error("Exchange not found")
	ErrUserExists             = Error("user already exists")
	ErrUserIDRequired         = Error("user id required")
	ErrUserNameRequired       = Error("user's username required")
	ErrExchIDRequired         = Error("Exchange id required")
	ErrInvalidJSON            = Error("invalid json")
	ErrUserRequired           = Error("user required")
	ErrExchRequired           = Error("Exchange required")
	ErrInvalidEntry           = Error("invalid Entry")
	ErrUserNullPointer        = Error("User value is nill or User is Empty")
	ErrUserNotCached          = Error("User is not or was unable to be saved in Cache or Session")
	ErrUserNameEmpty          = Error("Username is empty. Please enter a Username")
	ErrOperatorNameEmpty      = Error("Operator details required Username of the operator is Empty")
	ErrOperatorNotAdmin       = Error("Requires an Admin Operator")
	ErrUserPasswordEmpty      = Error("Password is empty. Please enter correct Password")
	ErrUsrDbUnreachable       = Error("Unable to get the UserDB into the Method")
	ErrExchDbUnreachable       = Error("Unable to get the ExchDB into the Method")
	ErrSessionCookieSaveError = Error("could not save cookie session please ensure cookie is enable on your browser")
	ErrIvalidRedirect         = Error("invalid redirect URL, Please try again")
	ErrSessionCookieError     = Error("could not create a cookie session please ensure cookie is enable on your browser")
)

// Error represents a User error.
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }
