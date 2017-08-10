package templateMethod

type CharDisplay struct {
	ch 	string
}

func (this *CharDisplay) CharDisplay(ch string) {
	this.ch = ch
}

func (this *CharDisplay) Open()  {
	print("<<")
}

func (this *CharDisplay) Print()  {
	print(this.ch)
}

func (this *CharDisplay) Close() {
	println(">>")
}

func (this *CharDisplay) Display() {
	this.Open()
	for i:=0;i<5;i++ {
		this.Print()
	}
	this.Close()
}
