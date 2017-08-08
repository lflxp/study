package iterator

type BookShelf struct {
	books  []Book
	last 	int
}

func (this *BookShelf) BookShelf(maxsize int) {
	//this.books = make([]Book,maxsize)
	this.books = []Book{}
	this.last = 0
}
func (this *BookShelf) GetBookAt(index int) *Book {
	return &this.books[index]
}

func (this *BookShelf) AppendBook(book Book) {
	this.books = append(this.books,book)
	this.last += 1
}

func (this *BookShelf) GetLength() int {
	return this.last
}

func (this *BookShelf) Iterator() *BookShelfIterator {
	book := BookShelfIterator{}
	book.BookShelfIterator(this)
	return &book
}