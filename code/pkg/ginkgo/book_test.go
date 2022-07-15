package ginkgo

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Book", func() {
	// Get Catalog
	ginkgo.Describe("get model catalog", func() {
		var foxInSocks, lesMis *Book
		ginkgo.BeforeEach(func() {
			lesMis = &Book{
				Title:  "Les Miserables",
				Author: "Victor Hugo",
				Pages:  2783,
			}
			foxInSocks = &Book{
				Title:  "Fox In Socks",
				Author: "Dr. Seuss",
				Pages:  24,
			}
		})

		ginkgo.Context("pages count <= 300", func() {
			ginkgo.It("be a short story", func() {
				gomega.Expect(foxInSocks.Catalog()).To(gomega.Equal(CategoryShortStory))
			})
		})
		ginkgo.Context("pages count > 300", func() {
			ginkgo.It("be a novel", func() {
				gomega.Expect(lesMis.Catalog()).To(gomega.Equal(CategoryNovel))
			})
		})
	})
})
