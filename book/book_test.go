package book_test

import (
	. "kdh.com/book"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"sync"
)

var wg sync.WaitGroup

var _ = Describe("Asynchronous Testing", func() {
	It("should post to the channel, eventually", func(done Done) {
		wg.Add(1)
		c := make(chan string, 0)

		go DoSomething(c)

		Expect(<-c).To(ContainSubstring("Done!"))

		close(done)
		wg.Wait()
	})
})

func DoSomething(c chan string) {
	c <- "Done!"
	wg.Done()
}

var _ = Describe("Better Book Tests", func() {
	var (
		ing  Book
		err  error
		json string
	)

	BeforeEach(func() {
		json = `{"title":"Les Miserables","author":"Victor Hugo","pages":1488}`
	})

	JustBeforeEach(func() {
		ing, err = NewBookFromJSON(json)
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses successfully", func() {
			It("should populate the fields correctly", func() {
				Expect(ing.Title).To(Equal("Les Miserables"))
				Expect(ing.Author).To(Equal("Victor Hugo"))
				Expect(ing.Pages).To(Equal(1488))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fials to parse", func() {
			BeforeEach(func() {
				json = `{"title":"Les Miserables","author":"Victor Hugo","pages":1488oops}`
			})

			It("should return the zero-vale for the book", func() {
				Expect(ing).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Extracting the author's last name", func() {
		It("should correctly identify and return the last name", func() {
			Expect(ing.AuthorLastName()).To(Equal("Hugo"))
		})
	})
})

var _ = Describe("Messy Book Tests", func() {
	It("can be loaded from JSON", func() {
		book, err := NewBookFromJSON(`{"title":"Les Miserables","author":"Victor Hugo","pages":1488}`)

		Expect(err).NotTo(HaveOccurred())
		Expect(book.Title).To(Equal("Les Miserables"))
		Expect(book.Author).To(Equal("Victor Hugo"))
		Expect(book.Pages).To(Equal(1488))
	})

	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Suess",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("This is a failing test...", func() {
			By("inside Context...By uses GinkgoWriter meaning this message should only be seen on a test failure.")
			XIt("should fail", func() {
				Fail("This is an intentional failure to demonstrate the usage of the Fail() method.")
			})
		})

		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal(NOVEL))
			})
		})

		Context("With Fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal(SHORT_STORY))
			})
		})
	})
})
