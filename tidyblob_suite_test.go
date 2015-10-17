package tidyblob_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTidyblob(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tidyblob Suite")
}
