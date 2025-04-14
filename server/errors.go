package server

import "encoding/json"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var errorEnvNameNotSpecified Error
var errorEnvNameNotSpecifiedBytes []byte

var errorFeatureFlagNotFound Error
var errorFeatureFlagNotFoundBytes []byte

var errorEnvironmentNotFound Error
var errorEnvironmentNotFoundBytes []byte

func init() {
	// Create errors and pre-serialize them.

	var err error

	errorEnvNameNotSpecified = Error{
		Code:    "DF001",
		Message: "Environment name not specified",
	}
	errorEnvNameNotSpecifiedBytes, err = json.Marshal(errorEnvNameNotSpecified)
	if err != nil {
		// TODO
		panic("TODO")
	}

	errorFeatureFlagNotFound = Error{
		Code:    "DF002",
		Message: "Feature flag not found",
	}
	errorFeatureFlagNotFoundBytes, err = json.Marshal(errorFeatureFlagNotFound)
	if err != nil {
		// TODO
		panic("TODO")
	}

	errorEnvironmentNotFound = Error{
		Code:    "DF003",
		Message: "Environment not found",
	}
	errorEnvironmentNotFoundBytes, err = json.Marshal(errorEnvironmentNotFound)
	if err != nil {
		// TODO
		panic("TODO")
	}
}
