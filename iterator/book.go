package iterator

type Book struct {
	Name 	string
}

func (this *Book) Book(name string) {
	this.Name = name
}

func (this *Book) String() string {
	return this.Name
}

func (this *Book) GetName() string {
	return this.Name
}
