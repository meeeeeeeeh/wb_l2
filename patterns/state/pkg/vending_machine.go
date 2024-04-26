package pkg

type VendingMachine struct {
	hasItem State
	itemRequested State
	hasMoney State
	noIems State
	currentState State
	itemCount int
	itemPrice int
}

//инициализация вендинговой машины
func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}
	v.SetState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noIems = noItemState
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
    v.itemCount = v.itemCount + count
}