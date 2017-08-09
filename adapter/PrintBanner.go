package adapter
//使用继承的适配器
type PrintBanner struct {
	Banner
}

func (this *PrintBanner) PrintBanner(st string) {
	this.string = st
}

func (this *PrintBanner) PrintWeak() {
	this.ShowWithParen()
}

func (this *PrintBanner) PrintStrong() {
	this.ShowWithAster()
}