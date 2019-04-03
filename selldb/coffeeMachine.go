package selldb

// update constant values as per db
const CoffeeMachineTypeLarge = "large"
const CoffeeMachineTypeSmall = "small"
const CoffeeMachineTypeEspresso = "espresso"

var (
	CoffeeMachineTypes = []string{
		CoffeeMachineTypeLarge,
		CoffeeMachineTypeSmall,
		CoffeeMachineTypeEspresso,
	}
)

// CoffeeMachine struct represents flattened coffee_machine tbl
type CoffeeMachine struct {
	ID                    int    `json:"coffee_machine_id" db:"coffee_machine_id"`
	IsWaterLineCompatible bool   `json:"is_water_line_compatible" db:"is_water_line_compatible"`
	MachineType           string `json:"machine_type" db:"machine_type"`
	Name                  string `json:"name" db:"name"`
}

type CoffeeMachines []*CoffeeMachine

// TODO create machine type table
