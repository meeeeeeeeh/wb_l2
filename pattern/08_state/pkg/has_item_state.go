package pkg

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

//нужно выдать товар
func (i *HasItemState) RequestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noIems)
		return fmt.Errorf("no item pesent")
	}
	fmt.Println("Item requested")
	i.vendingMachine.SetState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.IncrementItemCount(count)
	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}