package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	// tells go test runner this fn isa test helper, when errored, the Go test runner
	// will report the filename and line number of the code that called this fn
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}
