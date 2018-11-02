package models

import (
    "encoding/json"
    "net/http"
)


// Errors

type Errors struct {
    Errors []*Error `json:"errors"`
}

type Error struct {
    Id     string `json:"id"`
    Status int    `json:"status"`
    Title  string `json:"title"`
    Detail string `json:"detail"`
}

var (
    ErrBadRequest = &Error{"bad_request", 400, "Bad Request", "Request body is not well-formed. It must be JSON."}
    ErrNotAcceptable = &Error{"not_acceptable", 406, "Not Acceptable", "Accept header must be set to 'application/vnd.api+json'."}
    ErrUnsupportedMediaType = &Error{"unsupported_media_type", 415, "Unsupported Media Type", "Content-Type header must be set to: 'application/vnd.api+json'."}
    ErrInternalServer = &Error{"internal_server_error", 500, "Internal Server Error", "Oops... something went wrong."}
    ErrUserAlreadyExists = &Error{"user_already_exists", 701, "User Already Exists", "A user already exists using this email address."}

    ErrUserNotFound = &Error{"not_found", 702, "Not Found", "Your email or password is incorrect. Please try again."}
    ErrUserTokenRejected = &Error{"token_rejected", 703, "Token Not Acceptable", "Token rejected. Please try again."}
    ErrUserMissingData = &Error{"missing_fdata", 704, "Missing data", "Required data is missing. Please try again."}

    ErrAccountDisabled = &Error{"account_disabled", 805, "Account disabled", "Account has been disabled."}
    ErrLexpExpired = &Error{"lexp_expired", 806, "lexp Expired", "Long term token expired."}
)

func WriteError(w http.ResponseWriter, err *Error) {
    w.Header().Set("Content-Type", "application/vnd.api+json")
    w.WriteHeader(err.Status)
    json.NewEncoder(w).Encode(Errors{[]*Error{err}})
}