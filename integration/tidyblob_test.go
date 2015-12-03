package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os/exec"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Tidyblob", func() {
	var (
		command *exec.Cmd
		session *gexec.Session

		bosh_release_dir_path     = "../assets/test-boshrelease"
		non_bosh_release_dir_path = "../assets"
	)

	BeforeEach(func() {
		command = exec.Command(pathToTidyblobCLI)
	})

	JustBeforeEach(func() {
		var err error
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("when run inside of a BOSH release directory", func() {
		BeforeEach(func() {
			command.Dir = bosh_release_dir_path
		})

		It("exits with a status code 0", func() {
			Eventually(session).Should(gexec.Exit(0))
		})
	})

	Context("when run outside of a BOSH release directory", func() {
		BeforeEach(func() {
			command.Dir = non_bosh_release_dir_path
		})

		It("exits with a status code 1", func() {
			Eventually(session).Should(gexec.Exit(1))
		})

		It("writes a useful error message to stderr", func() {
			Eventually(session.Err).Should(gbytes.Say("Sorry, your current directory doesn't look like release directory"))
		})
	})
})
