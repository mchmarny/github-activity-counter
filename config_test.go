package fn

import (
	"os"
	"testing"
)

func TestDefaultConfigInitializer(t *testing.T) {
	err := defaultConfigInitializer()
	if err != nil {
		t.Error(err)
	}
}

func TestDefaultConfigInitializerError(t *testing.T) {
	hookSecret := os.Getenv(hookSecretEnvVarName)
	os.Setenv(hookSecretEnvVarName, "")
	err := defaultConfigInitializer()
	if err == nil {
		t.Error(err)
	}
	os.Setenv(hookSecretEnvVarName, hookSecret)
}
