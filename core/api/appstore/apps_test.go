package appstore_test

import (
	"slices"
	"testing"
)

func TestAppStoreSlice(t *testing.T) {
	a := []string{"1.0.0", "1.2.1"}

	t.Log(slices.Contains(a, "1.0.0"))
}
