package main

import (
	"fmt"
	. "github.com/lflxp/study/debx"
)

func main() {
	a := Debx{Total:200000.0,Rates:0.0155,Months:36.0}
	a.GetPerMonthPay()
	fmt.Println("期数 未偿还本金 偿还本金 偿还利息 偿还本息和")
	for i:=1;i<=int(a.Months);i++ {
		fmt.Println(fmt.Sprintf("%d %.2f %.2f %.2f %.2f",i,a.GetUnPay(i),a.GetBenJin(float64(i)),a.GetRatePay(float64(i)),a.PerMonth))
	}
	fmt.Println(fmt.Sprintf("%s %.2f %.2f %.2f %.2f","合计",0.00,a.Total,a.SumRates(),a.TotalPay()))
}