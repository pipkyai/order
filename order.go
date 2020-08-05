package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/leizongmin/huobiapi"
)

var wg = sync.WaitGroup{}

// Orders - структура с ордерами
type Orders struct {
	Symbol  string
	Ask     float64
	AskSize float64
	Bid     float64
	BidSize float64
}

// Price - структура с ценами для каждой пары
type Price struct {
	eb float64 //последняя котировка eth/btc
	eu float64 //последняя котировка eth/usdt
	bu float64 //последняя котировка btc/usdt
}

func main() {
	p := &Price{1., 1., 1.}
	k := make(chan float64, 3) //канал передающий текущую цену
	l := make(chan string, 3)  //канал передающий название пары для переданной выше цены
	runtime.GOMAXPROCS(10)
	wg.Add(3)
	//включаем websocket для торговых пар
	go btcusdt(k, l)
	go xbtc(k, l)
	go xusdt(k, l)

	//при изменении цены или размера ордера пересчитываем коэффициент
	for msg := range k { //msg - принимает цену через канал для пары
		msg1 := <-l                  //msg1- принимает название пары
		x := calculate(msg1, msg, p) //х - принимает значение коэффициента
		fmt.Println(x)
	}
	fmt.Scanln()
}

func btcusdt(k chan float64, l chan string) {
	// создать экземпляр клиента

	market, err := huobiapi.NewMarket()
	if err != nil {
		panic(err)
	}
	// подписаться на темы
	market.Subscribe("market.btcusdt.bbo", func(topic string, json *huobiapi.JSON) {
		topic = ""
		// записываем все в структуру
		bu := Orders{}

		s, err := json.Get("tick").Get("symbol").String()
		if err != nil {
			panic(err)
		}

		a, err := json.Get("tick").Get("ask").Float64()
		if err != nil {
			panic(err)
		}
		b, err := json.Get("tick").Get("askSize").Float64()
		if err != nil {
			panic(err)
		}
		c, err := json.Get("tick").Get("bid").Float64()
		if err != nil {
			panic(err)
		}
		d, err := json.Get("tick").Get("bidSize").Float64()
		if err != nil {
			panic(err)
		}
		bu.Symbol = s
		bu.Ask = a
		bu.AskSize = b
		bu.Bid = c
		bu.BidSize = d
		fmt.Println(bu)
		l <- bu.Symbol
		k <- bu.Ask

	})
	// запрос данных
	json, err := market.Request("market.btcusdt.bbo")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(json)
	}
}

func xbtc(k chan float64, l chan string) {
	// создать экземпляр клиента
	market, err := huobiapi.NewMarket()
	if err != nil {
		panic(err)
	}
	// подписаться на темы
	market.Subscribe("market.htbtc.bbo", func(topic string, json *huobiapi.JSON) {
		topic = ""
		// записываем все в структуру
		eb := Orders{}
		s, err := json.Get("tick").Get("symbol").String()
		if err != nil {
			panic(err)
		}

		a, err := json.Get("tick").Get("ask").Float64()
		if err != nil {
			panic(err)
		}
		b, err := json.Get("tick").Get("askSize").Float64()
		if err != nil {
			panic(err)
		}
		c, err := json.Get("tick").Get("bid").Float64()
		if err != nil {
			panic(err)
		}
		d, err := json.Get("tick").Get("bidSize").Float64()
		if err != nil {
			panic(err)
		}
		eb.Symbol = s
		eb.Ask = a
		eb.AskSize = b
		eb.Bid = c
		eb.BidSize = d
		fmt.Println(eb)
		l <- eb.Symbol
		k <- eb.Ask
	})
	// запрос данных
	json, err := market.Request("market.htbtc.bbo")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(json)
	}
}

func xusdt(k chan float64, l chan string) {
	// создать экземпляр клиента
	market, err := huobiapi.NewMarket()
	if err != nil {
		panic(err)
	}
	// подписаться на темы
	market.Subscribe("market.htusdt.bbo", func(topic string, json *huobiapi.JSON) {
		topic = ""
		// записываем все в структуру
		eu := Orders{}
		s, err := json.Get("tick").Get("symbol").String()
		if err != nil {
			panic(err)
		}

		a, err := json.Get("tick").Get("ask").Float64()
		if err != nil {
			panic(err)
		}
		b, err := json.Get("tick").Get("askSize").Float64()
		if err != nil {
			panic(err)
		}
		c, err := json.Get("tick").Get("bid").Float64()
		if err != nil {
			panic(err)
		}
		d, err := json.Get("tick").Get("bidSize").Float64()
		if err != nil {
			panic(err)
		}
		eu.Symbol = s
		eu.Ask = a
		eu.AskSize = b
		eu.Bid = c
		eu.BidSize = d
		fmt.Println(eu)
		l <- eu.Symbol
		k <- eu.Bid
	})
	// запрос данных
	json, err := market.Request("market.htusdt.bbo")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(json)
	}
}

//функция получает название пары и ее текущую цену, и пересчитывает коэффициент
func calculate(s string, x float64, p *Price) float64 {
	switch s {
	case "btcusdt":
		p.bu = x
	case "htbtc":
		p.eb = x
	case "htusdt":
		p.eu = x
	}
	k := 1 / p.eb * p.eu / p.bu
	return k
}
