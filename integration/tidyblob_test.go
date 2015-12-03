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

		bosh_release_dir_path                = "../assets/test-boshrelease"
		bosh_release_dir_no_stale_blobs_path = "../assets/test-boshrelease-no-stale-blobs"
		non_bosh_release_dir_path            = "../assets"
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
		Context("when there are stale blobs", func() {
			BeforeEach(func() {
				command.Dir = bosh_release_dir_path
			})

			It("prints a list of stale blobs", func() {
				Eventually(session).Should(gbytes.Say("Found stale blobs:"))
				Eventually(session).Should(gbytes.Say("lifecycles/windows_app_lifecycle-95ad2f6.tgz"))
				Eventually(session).Should(gbytes.Say("rootfs/cflinuxfs2-1.11.0.tar.gz"))
			})

			It("doesn't print a list of required blobs", func() {
				Eventually(session).ShouldNot(gbytes.Say("cf-cli/cf-linux-amd64.tgz"))
				Eventually(session).ShouldNot(gbytes.Say("docker/docker-1.6.2"))
				Eventually(session).ShouldNot(gbytes.Say("golang/go1.4.3.linux-amd64.tar.gz"))
			})

			It("exits with a status code 0", func() {
				Eventually(session).Should(gexec.Exit(0))
			})
		})

		Context("when there aren't stale blobs", func() {
			BeforeEach(func() {
				command.Dir = bosh_release_dir_no_stale_blobs_path
			})

			It("prints a useful message to stdout", func() {
				Eventually(session).Should(gbytes.Say("Didn't detect any stale blobs"))
			})

			It("doesn't print a list of blobs", func() {
				Eventually(session).ShouldNot(gbytes.Say("cf-cli/cf-linux-amd64.tgz"))
				Eventually(session).ShouldNot(gbytes.Say("docker/docker-1.6.2"))
				Eventually(session).ShouldNot(gbytes.Say("golang/go1.4.3.linux-amd64.tar.gz"))
				Eventually(session).ShouldNot(gbytes.Say("lifecycles/windows_app_lifecycle-95ad2f6.tgz"))
				Eventually(session).ShouldNot(gbytes.Say("rootfs/cflinuxfs2-1.11.0.tar.gz"))
			})

			It("exits with a status code 0", func() {
				Eventually(session).Should(gexec.Exit(0))
			})
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
