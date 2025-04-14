package server

import (
	"net/http"

	"github.com/denisdefreyne/dflag/stores"
)

func getHealth(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ok"))
}

var enabledTrueResponseBytes []byte
var enabledFalseResponseBytes []byte

func init() {
	enabledTrueResponseBytes = []byte(`{"enabled":true}`)
	enabledFalseResponseBytes = []byte(`{"enabled":false}`)
}

func makeFeatureFlagHandlerFunc(featureFlags *stores.FeatureFlagsStore) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		featureFlagName := req.PathValue("name")

		environmentName := req.FormValue("env")
		if environmentName == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorEnvNameNotSpecifiedBytes)
			return
		}

		// Find feature flag
		lookup, ok := featureFlags.Lookup[featureFlagName]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write(errorFeatureFlagNotFoundBytes)
			return
		}

		// Find environment
		lookup2, ok := lookup[environmentName]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write(errorEnvironmentNotFoundBytes)
			return
		}

		// Write response
		if lookup2 {
			w.Write(enabledTrueResponseBytes)
		} else {
			w.Write(enabledFalseResponseBytes)
		}
	}
}
