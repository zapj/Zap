package tests

import (
	"testing"

	"github.com/zapj/zap/core/utils/zaputil"
)

func TestPwdHash(t *testing.T) {
	zaputil.CheckPasswordHash("123456", "$2a$14$xb322kIESsiqobBYkfNDeellO7u2tGRPxdJ9xo4aUe2IRMvNUiVs2")

}
