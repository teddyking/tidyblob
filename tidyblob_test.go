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
	)

	Describe("Blobs", func() {
		It("returns a []Blob of all blobs in the provided blobs.yml file", func() {
			blobs, err := Blobs(blobs_yml_file_path)

			Expect(err).ToNot(HaveOccurred())
			Expect(len(blobs)).To(Equal(5))
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
})
