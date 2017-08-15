package debx

import (
	"testing"
	"fmt"
)

func TestDebx_GetPerMonthPay(t *testing.T) {
	a := Debx{Total:200000.0,Rates:0.0155,Months:36.0}
	a.GetPerMonthPay()
	t.Log(fmt.Sprintf("%.2f",a.PerMonth))
}

func TestDebx_GetRatePay(t *testing.T) {
	a := Debx{Total:200000.0,Rates:0.0155,Months:36.0}
	a.GetPerMonthPay()
	t.Log(fmt.Sprintf("%.2f",a.GetRatePay(3)))
}

func TestDebx_GetBenJin(t *testing.T) {
	a := Debx{Total:200000.0,Rates:0.0155,Months:36.0}
	a.GetPerMonthPay()
	t.Log(fmt.Sprintf("%.2f",a.GetBenJin(3)))
}

func TestDebx_GetUnPay(t *testing.T) {
	a := Debx{Total:200000.0,Rates:0.0155,Months:36.0}
	a.GetPerMonthPay()
	t.Log(fmt.Sprintf("%.2f",a.GetUnPay(3)))
}

func TestDebx_SumRates(t *testing.T) {
	a := Debx{Total:200000.0,Rates:0.0155,Months:36.0}
	a.GetPerMonthPay()
	t.Log(fmt.Sprintf("%.2f",a.SumRates()))
}

func TestDebx_TotalPay(t *testing.T) {
	a := Debx{Total:200000.0,Rates:0.0155,Months:36.0}
	a.GetPerMonthPay()
	t.Log(fmt.Sprintf("%.2f",a.TotalPay()))
}