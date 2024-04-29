// структурный паттерн который представляет из себя простой интерфейс к сложной системе

//пример
//есть интернет магазин и платежная система
// это два интерфейса
//для покупки нужно сделать интерфейс для покупки - этто будет фасад над общим поведением платежной системой

// плюсы
//изолирует клиентов от сложного поведения

//минусы
//слишком большая привязка к интерфейсу фасада

package main

import (
	"facade/pkg"
	"fmt"
)

var (
	bank = pkg.Bank{
		Name:  "БАНК",
		Cards: []pkg.Card{},
	}
	card1 = pkg.Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = pkg.Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}
	user1 = pkg.User{
		Name: "Покупатель-1",
		Card: &card1,
	}
	user2 = pkg.User{
		Name: "Покупатель-2",
		Card: &card2,
	}
	prod = pkg.Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = pkg.Shop{
		Name: "SHOP",
		Products: []pkg.Product{
			prod,
		},
	}
)

// функция sell - фасад над бизнес логикой поведения по безналичному рассчету
func main() {
	fmt.Println("[Банк] Выпуск карты")
	bank.Cards = append(bank.Cards, card1, card2)

	fmt.Printf("[%s]\n", user1.Name)
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
