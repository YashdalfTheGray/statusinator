package env

import (
	"os"
	"strings"
)

var envVarMap = make(map[string]string, len(allEnvKeys))

// Get returns the value of an environment variable. If it doesn't
// exist, it asks os.LookupEnv and caches the value.
func Get(envKey string) string {
	if envVarMap[envKey] == "" {
		envValue, _ := os.LookupEnv(envKey)
		envVarMap[envKey] = envValue
	}
	return envVarMap[envKey]
}

// IsDevelopment returns whether the runtime environment
// is development. Specifically, it checks the value
// of the ENV environment variable.
func IsDevelopment() bool {
	return strings.ToLower(Get(RuntimeEnv)) == RuntimeDev
}
