package main

import "fmt"

// 支付策略接口
type PaymentStrategy interface {
	Pay(amount float64) string
}

// 具体支付策略：信用卡支付
type CreditCardStrategy struct {
	name string
}

func (c *CreditCardStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card %s", amount, c.name)
}

// 具体支付策略：支付宝支付
type AlipayStrategy struct {
	account string
}

func (a *AlipayStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Alipay Account %s", amount, a.account)
}

// 具体支付策略：微信支付
type WechatPayStrategy struct {
	account string
}

func (w *WechatPayStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Wechat Pay Account %s", amount, w.account)
}

// 上下文，用于选择和执行支付策略
type ShoppingCart struct {
	paymentStrategy PaymentStrategy
	total           float64
}

func (c *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	c.paymentStrategy = strategy
}

func (c *ShoppingCart) CalculateTotal() {
	// 假设这里有一些计算总价的逻辑
	c.total = 100.0
}

func (c *ShoppingCart) Checkout() string {
	if c.paymentStrategy == nil {
		return "No payment strategy set"
	}
	return c.paymentStrategy.Pay(c.total)
}

func main() {
	// 创建购物车
	cart := &ShoppingCart{}
	// 计算总价
	cart.CalculateTotal()

	// 设置信用卡支付策略
	creditCard := &CreditCardStrategy{name: "1234-5678-9012-3456"}
	cart.SetPaymentStrategy(creditCard)
	fmt.Println(cart.Checkout())

	// 设置支付宝支付策略
	alipay := &AlipayStrategy{account: "alipay@example.com"}
	cart.SetPaymentStrategy(alipay)
	fmt.Println(cart.Checkout())

	// 设置微信支付策略
	wechatPay := &WechatPayStrategy{account: "wechatpay@example.com"}
	cart.SetPaymentStrategy(wechatPay)
	fmt.Println(cart.Checkout())
}
