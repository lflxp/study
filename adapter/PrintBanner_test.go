package adapter

import "testing"

//使用接口类型Print 初始化PrintBanner类并赋值
func TestPrintBanner_PrintBanner(t *testing.T) {
	var tmp PrintBanner = PrintBanner{}
	tmp.PrintBanner("Hello")
	var p Print = &tmp
	//p := &PrintBanner{}
	//p.PrintBanner("Hello")
	p.PrintWeak()
	p.PrintStrong()
	t.Log("ok")
}

func TestPrintBanner2_PrintBanner(t *testing.T) {
	var tmp PrintBanner = PrintBanner{}
	tmp.PrintBanner("Lixueping")
	var p Print = &tmp
	//p := &PrintBanner2{}
	//p.PrintBanner("Lixueping")
	p.PrintWeak()
	p.PrintStrong()
	t.Log("ok2")
}