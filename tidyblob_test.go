package tidyblob_test

import (
	. "github.com/teddyking/tidyblob"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tidyblob", func() {
	var (
		blobs_yml_file_path             = "assets/blobs.yml"
		invalid_blobs_yml_file_path     = "assets/blobs.txt"
		nonexistent_blobs_yml_file_path = "assets/nope.yml"
		bosh_packages_dir_path          = "assets/packages"
		invalid_bosh_packages_dir_path  = "assets/invalid_packages"
		empty_bosh_packages_dir_path    = "assets"
	)

	Describe("Blobs", func() {
		It("returns a []string of all blob names in the provided blobs.yml file", func() {
			blobs, err := Blobs(blobs_yml_file_path)

			Expect(err).ToNot(HaveOccurred())
			Expect(len(blobs)).To(Equal(5))
			Expect(blobs).To(ConsistOf(
				"cf-cli/cf-linux-amd64.tgz",
				"docker/docker-1.6.2",
				"golang/go1.4.3.linux-amd64.tar.gz",
				"lifecycles/windows_app_lifecycle-95ad2f6.tgz",
				"rootfs/cflinuxfs2-1.11.0.tar.gz",
			))
		})

		Context("when the blobs.yml file doesn't exist", func() {
			It("returns an error", func() {
				blobs, err := Blobs(nonexistent_blobs_yml_file_path)

				Expect(err).To(HaveOccurred())
				Expect(blobs).To(BeNil())
			})
		})

		Context("when the blobs.yml file contains invalid YAML", func() {
			It("returns an error", func() {
				blobs, err := Blobs(invalid_blobs_yml_file_path)

				Expect(err).To(HaveOccurred())
				Expect(blobs).To(BeNil())
			})
		})
	})

	Describe("RequiredBlobs", func() {
		It("returns a []string of blob names that are required for BOSH packages in the provided packages dir", func() {
			blobs, err := RequiredBlobs(bosh_packages_dir_path)

			Expect(err).ToNot(HaveOccurred())
			Expect(len(blobs)).To(Equal(3))
			Expect(blobs).To(ConsistOf(
				"cf-cli/cf-linux-amd64.tgz",
				"docker/docker-1.6.2",
				"golang/go1.4.3.linux-amd64.tar.gz",
			))
		})

		Context("when no <package>/spec files exist under the provided packages dir", func() {
			It("returns an empty []string", func() {
				blobs, err := RequiredBlobs(empty_bosh_packages_dir_path)

				Expect(err).ToNot(HaveOccurred())
				Expect(len(blobs)).To(Equal(0))
			})
		})

		Context("when a matched package's spec file contains invalid YAML", func() {
			It("returns an error", func() {
				blobs, err := RequiredBlobs(invalid_bosh_packages_dir_path)

				Expect(err).To(HaveOccurred())
				Expect(blobs).To(BeNil())
			})
		})
	})

	Describe("StaleBlobs", func() {
		It("returns a []string of blob names that are not required by a BOSH release", func() {
			blobs, err := StaleBlobs(blobs_yml_file_path, bosh_packages_dir_path)

			Expect(err).ToNot(HaveOccurred())
			Expect(len(blobs)).To(Equal(2))
			Expect(blobs).To(ConsistOf(
				"lifecycles/windows_app_lifecycle-95ad2f6.tgz",
				"rootfs/cflinuxfs2-1.11.0.tar.gz",
			))
		})
	})
})
