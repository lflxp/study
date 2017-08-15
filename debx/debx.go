package debx

import (
	"math"
)

type Debx struct {
	Total 	float64 	// 贷款额度
	Rates 	float64 //月利率
	Months  float64   //还款月数
	PayMonths float64 //每月还款额
	RateTotal float64 //还款利息总和
	PerMonth  float64  //偿还本息和
}

//获得每个月偿还的钱
//偿还本息和
//每月还款额计算公式如下： [贷款本金×月利率×(1+月利率)^还款月数]÷[(1+月利率)^还款月数-1]
func (this *Debx) GetPerMonthPay() {
	this.PerMonth = this.Total*this.Rates*math.Pow(1.0+this.Rates,this.Months)/(math.Pow(1.0+this.Rates,this.Months)-1)
}

//每月还款利息
//第n月还款利息为：＝（a×i－b）×（1＋i）^（n－1）＋b
func (this *Debx) GetRatePay(n float64) float64 {
	return (this.Total*this.Rates-this.PerMonth)*math.Pow(1.0+this.Rates,n-1.0)+this.PerMonth
}

//每月偿还本金
//每月偿还本金=偿还本息和-当月还款利息
func (this *Debx) GetBenJin(n float64) float64 {
	return this.PerMonth-this.GetRatePay(n)
}

//未偿还本金
func (this *Debx) GetUnPay(n int)  float64 {
	tmp := this.Total
	for i:=1;i<n;i++ {
		tmp = tmp - this.GetBenJin(float64(i))
	}
	return tmp
}

//还款利息总和
//求以上和为：Y＝（a×i－b）×〔（1＋i）^n－1〕÷i＋n×b
func (this *Debx) SumRates() float64 {
	return (this.Total*this.Rates-this.PerMonth)*(math.Pow(1.0+this.Rates,this.Months)-1.0)/this.Rates+this.Months*this.PerMonth
}

//还款总额
func (this *Debx) TotalPay() float64 {
	return this.Total+this.SumRates()
}
