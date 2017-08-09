package adapter
//使用委托的示例程序
type PrintBanner2 struct {
	banner *Banner
}

func (this *PrintBanner2) PrintBanner(st string) {
	b := Banner{}
	b.Banner(st)
	this.banner = &b
}

func (this *PrintBanner2) PrintWeak() {
	this.banner.ShowWithParen()
}

func (this *PrintBanner2) PrintStrong() {
	this.banner.ShowWithAster()
}