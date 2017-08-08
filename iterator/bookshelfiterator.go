package iterator

type BookShelfIterator struct {
	Iterator
	BookShelf 	*BookShelf
	Index 		int
}

func (this *BookShelfIterator) BookShelfIterator(bookshelf *BookShelf) {
	this.BookShelf = bookshelf
	this.Index = 0
}

func (this *BookShelfIterator) HasNext() bool {
	if this.Index < this.BookShelf.GetLength() {
		return true
	} else {
		return false
	}
}

func (this *BookShelfIterator) Next() interface{} {
	book := this.BookShelf.GetBookAt(this.Index)
	this.Index +=1
	//println(this.Index)
	return book
}
