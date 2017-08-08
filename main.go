package main

import (
	. "github.com/lflxp/study/iterator"
	"fmt"
)

func main() {
	bookshelf := &BookShelf{}
	bookshelf.BookShelf(4)
	bookshelf.AppendBook(Book{Name:"Around the world in 80 days"})
	bookshelf.AppendBook(Book{Name:"Bible"})
	bookshelf.AppendBook(Book{Name:"Cinderella"})
	bookshelf.AppendBook(Book{Name:"Daddy-Long-Legs"})

	/*fmt.Println(bookshelf.GetLength())
	for i:=0;i<bookshelf.GetLength();i++ {
		fmt.Println(bookshelf.GetBookAt(i).GetName())
	}*/
	it := bookshelf.Iterator()
	for {
		if it.HasNext() {
			data := it.Next().(*Book)
			fmt.Println(data)
		} else {
			fmt.Println("over")
			break
		}
	}
}