package templateMethod

type StringDisplay struct {
	st  	string
	width 	int
}

func (this *StringDisplay) StringDisplay(ch string)  {
	this.st = ch
	this.width = len(this.st)
}

func (this *StringDisplay) printLine() {
	print("+")
	for i:=0;i<this.width;i++ {
		print("-")
	}
	println("+")
}

func (this *StringDisplay) Open()  {
	this.printLine()
}

func (this *StringDisplay) Print()  {
	println("|"+this.st+"|")
}

func (this *StringDisplay) Close() {
	this.printLine()
}

func (this *StringDisplay) Display() {
	this.Open()
	for i:=0;i<5;i++ {
		this.Print()
	}
	this.Close()
}
