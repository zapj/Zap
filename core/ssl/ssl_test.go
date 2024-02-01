package ssl_test

import (
	"testing"

	"github.com/zapj/zap/core/ssl"
)

func TestGenerateKey(t *testing.T) {
	t.Log(ssl.GeneratePrivateKey("RSA", 2048))
}
