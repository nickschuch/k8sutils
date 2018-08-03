package configmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Backend struct {
	Host   string `k8s_configmap:"backend.host"`
	Port   int64  `k8s_configmap:"backend.port"`
	Secure bool   `k8s_configmap:"backend.secure"`
}

func TestUnmarshal(t *testing.T) {
	data := map[string]string{
		"backend.host":   "1.1.1.1",
		"backend.port":   "443",
		"backend.secure": "true",
	}

	var backend Backend

	err := unmarshal(data, &backend)
	assert.Nil(t, err)

	assert.Equal(t, "1.1.1.1", backend.Host)
	assert.Equal(t, int64(443), backend.Port)
	assert.True(t, backend.Secure)
}
