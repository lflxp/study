package adapter

import "fmt"

type Banner struct {
	string  string
}

func (this *Banner) Banner(st string) {
	this.string = st
}

func (this *Banner) ShowWithParen() {
	fmt.Println(fmt.Sprintf("(%s)",this.string))
}

func (this *Banner) ShowWithAster() {
	fmt.Println(fmt.Sprintf("*%s*",this.string))
}