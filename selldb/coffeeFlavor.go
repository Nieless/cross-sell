package selldb

// CoffeeFlavor struct represents flattened coffee_flavor tbl
type CoffeeFlavor struct {
	ID                    int    `json:"coffee_flavor_id" db:"coffee_flavor_id"`
	Name                  string `json:"name" db:"name"`
}
