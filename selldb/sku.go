package selldb

const UOM = "dozon"
const UOMUnitValue = 12

// SKU struct represents flattened sku tbl
type SKU struct {
	ID              int    `json:"sku_id" db:"id"`
	SKUName         string `json:"sku_name" db:"sku_name"`
	CoffeeMachineID *int   `json:"coffee_machine_id" db:"coffee_machine_id"`
	CoffeePodID     *int   `json:"coffee_pod_id" db:"coffee_pod_id"`
	Size            int    `json:"size" db:"size"`
	CoffeeFlavorID  int    `json:"coffee_flavor_id" db:"coffee_flavor_id"`
}

type SKUs []*SKU
