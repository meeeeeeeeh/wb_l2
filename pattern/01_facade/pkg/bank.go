package pkg

import (
	"fmt"
	"time"
	"errors"
)

type Bank struct {
	Name string
	Cards []Card
}

func (bank Bank) CheckBalance(cartNumber string) error {
	fmt.Printf("[Банк] Получение остатка по карте %s\n", cartNumber)
	time.Sleep(time.Millisecond * 300)

	for _, card := range bank.Cards {
		if card.Name != cartNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}

	fmt.Println("[Банк] Остаток положительный")
	return nil


}