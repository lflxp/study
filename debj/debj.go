package debj

import (
	"math"
)

type Debj struct {
	Total 	float64 	// 贷款额度
	Rates 	float64 //月利率
	Months  float64   //还款月数
	PayMonths float64 //每月还款额
	RateTotal float64 //还款利息总和
	PerMonth  float64  //偿还本金和
}

//获得每个月偿还的钱
//偿还本息和
//每月还款额计算公式如下： [贷款本金×月利率×(1+月利率)^还款月数]÷[(1+月利率)^还款月数-1]
func (this *Debj) GetPerMonthPay() {
	this.PerMonth = this.Total/this.Months
}

//每月还款利息
//第n月还款利息为：＝（a×i－b）×（1＋i）^（n－1）＋b
func (this *Debj) GetRatePay(n float64) float64 {
	return (this.Total-(n-1.0)*(this.Total/this.Months))*this.Rates
}

//每月偿还本金
//每月偿还本金=贷款额度/还款月数
func (this *Debj) GetBenJin() float64 {
	return this.Total/this.Months
}

//未偿还本金
func (this *Debj) GetUnPay(n float64)  float64 {
	return this.Total - this.GetBenJin()*n
}

//还款利息总和
//求以上和为：Y＝（a×i－b）×〔（1＋i）^n－1〕÷i＋n×b
func (this *Debj) SumRates() float64 {
	result := 0
	for i:=1;i<=int(this.Months)+1;i++ {
		result = result + this.GetRatePay(float64(i))
	}
	return result
}

//还款总额
func (this *Debj) TotalPay() float64 {
	return this.Total+this.SumRates()
}

