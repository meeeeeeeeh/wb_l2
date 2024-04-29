package pkg

import (
	"fmt"
	"time"
	"errors"
)

type Shop struct {
	Name string
	Products []Product
}

//фасад - функция sell 
//здесь происходит взаимодействие со всеми другим сервисами (банк, каты, товар, магазин)
//клиент приходит в магазин
//расплачивается картой через терминал
//с помощью карты запрашиваем инфу о балансе
//если он положительный то продаем товар 
// при этом сравнивается баланс на карте и стоимость товара

//это все скрывается под функцией sell - это фасад над всей бизнес логикой по покупке и оплате
func (shop Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю, для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)

	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}

	fmt.Printf("[Магазин] проверка - может ли [%s] пользователь купить товар\n", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостатточно средств для покупки товара")
		}
		fmt.Printf("[Магазин] Товар [%s] куплен\n", prod.Name)
	}
	return nil
}