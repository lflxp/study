package templateMethod

import "testing"

func TestStringDisplay_Open(t *testing.T) {
	var d1 AbstractDisplay
	var d2 AbstractDisplay
	var d3 AbstractDisplay

	h := &CharDisplay{}
	h.CharDisplay("H")
	d1 = h

	hello := &StringDisplay{}
	hello.StringDisplay("Hello,world.")
	d2 = hello

	nihao := &StringDisplay{}
	nihao.StringDisplay("你好，世界。")
	d3 = nihao

	d1.Display()
	d2.Display()
	d3.Display()
}
