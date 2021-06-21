package main

import "fmt"

//饮料接口
type Beverage interface {
	getDescription() string
	cost() int
}

//实现咖啡的过程
type Coffee struct {
	description string
}

func (c Coffee) getDescription() string {
	return c.description
}

func (c Coffee) cost() int {
	return 1
}

type Mocha struct {
	beverage Beverage
	description string
}

func (m Mocha) getDescription() string {
	return fmt.Sprintf("%s, %s", m.beverage.getDescription(), m.description)
}

func (m Mocha) cost() int {
	return m.beverage.cost() + 1
}

type Whip struct {
	beverage Beverage
	description string
}

func (w Whip) getDescription() string {
	return fmt.Sprintf("%s, %s", w.beverage.getDescription(), w.description)
}

func (w Whip) cost() int {
	return w.beverage.cost() + 1
}

//reference： https://learnku.com/articles/26820
func main() {
	var beverage Beverage
	//买了一杯咖啡
	beverage = Coffee{description: "houseBlend"}
	//给咖啡加上Mocha
	beverage = Mocha{beverage: beverage, description: "Mocha"}
	//给咖啡加上Whip
	beverage = Whip{beverage: beverage, description: "whip"}
	//最后计算Coffee的价格
	fmt.Println(beverage.getDescription(), ", cost is", beverage.cost())
}
