package failer_test

import (
	. "github.com/teddyking/tidyblob/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/teddyking/tidyblob/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestFailer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Failer Suite")
}
