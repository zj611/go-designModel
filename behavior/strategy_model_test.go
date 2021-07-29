package behavior

import (
	"fmt"
	"testing"
)

//策略模式
//定义了算法家族，在调用算法家族的时候不感知算法的变化，客户也不会受到影响。

//例：超市中经常进行促销活动，
//促销活动的促销方法就是一个个策略，
//如“满一百减20”，“打八折”等。
//现在实现策略模式，用CashContext生产策略，并完成策略的调用。

type CashSuper interface {
	AcceptMoney(money float32) float32
}

//普通模式，没有折扣
type CashNormal struct {
}

func NewCashNormal() CashNormal {
	return CashNormal{}
}

func (c CashNormal) AcceptMoney(money float32) float32 {
	return money
}

//打折模式，传入折扣，输出money * 折扣
type CashDiscount struct {
	Discount float32
}

func NewCashDiscount(discount float32) CashDiscount {
	return CashDiscount{
		Discount: discount,
	}
}

func (c CashDiscount) AcceptMoney(money float32) float32 {
	return c.Discount * money
}

//直接返利，如满100返20
type CashReturn struct {
	MoneyCondition float32
	MoneyReturn    float32
}

func newCashReturn(moneyCondition, moneyReturn float32) CashReturn {
	return CashReturn{
		MoneyCondition: moneyCondition,
		MoneyReturn:    moneyReturn,
	}
}

func (c CashReturn) AcceptMoney(money float32) float32 {
	if money >= c.MoneyCondition {
		moneyMinus := int(money / c.MoneyCondition)
		return money - float32(moneyMinus)*c.MoneyReturn
	}
	return money
}

//定义CashContext结构，用来做策略筛选
type CashContext struct {
	strategy CashSuper
}

func NewCashContext(cashType string) CashContext {
	c := CashContext{}

	switch cashType {
	case "discount0.8":
		c.strategy = NewCashDiscount(0.8)
	case "100return20":
		c.strategy = newCashReturn(100, 20)
	default:
		c.strategy = NewCashNormal()
	}
	return c
}

//在策略生产成功后，我们就可以直接调用策略的函数。
func (c CashContext) GetMoney(money float32) float32 {
	return c.strategy.AcceptMoney(money)
}

func TestStrategyModel(t *testing.T) {
	money := 100.0
	cc := NewCashContext("discount0.8")
	fmt.Println("money: ", cc.GetMoney(float32(money)))

	money = 190.0
	cc = NewCashContext("100return20")
	fmt.Println("money: ", cc.GetMoney(float32(money)))

	money = 190.0
	cc = NewCashContext("")
	fmt.Println("money: ", cc.GetMoney(float32(money)))

}
