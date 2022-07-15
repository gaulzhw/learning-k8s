package ginkgo

type Catalog int8

const (
	CategoryNovel Catalog = iota
	CategoryShortStory
)

const MaxShortStoryPages = 300

type Book struct {
	Title  string
	Author string
	Pages  int32
}

func (b *Book) Catalog() Catalog {
	if b.Pages < MaxShortStoryPages {
		return CategoryShortStory
	}
	return CategoryNovel
}
