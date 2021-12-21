package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	// panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
	// panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	val, ok := units[unit]
	if ok {
		bill[item] += val
		return ok
	} else {
		return false
	}
	// panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	val1, ok1 := units[unit]
	val2, ok2 := bill[item]
	if ok1 && ok2 {
		remainder := val2 - val1
		if remainder < 0 {
			return false
		} else if remainder > 0 {
			bill[item] = remainder
			return true
		} else {
			delete(bill, item)
			return true
		}
	} else {
		return false
	}
	// panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, ok := bill[item]
	if ok {
		return val, ok
	} else {
		return 0, ok
	}
	// panic("Please implement the GetItem() function")
}
