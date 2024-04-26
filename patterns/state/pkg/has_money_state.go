package pkg

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("iItem dispense in progress")
}

func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

//когда деньги внесены и было сравнение цены и внесенной суммы
//товар мжно только выдать
func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispending item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noIems)
	} else {
		i.vendingMachine.SetState(i.vendingMachine.hasItem)
	}
	return nil
}