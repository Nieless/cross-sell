package selldb

// update constant values as per db
const CoffeePodTypeLarge = "large"
const CoffeePodTypeSmall = "small"
const CoffeePodTypeEspresso = "espresso"


var (
	CoffeePodTypes = []string{
		CoffeePodTypeLarge,
		CoffeePodTypeSmall,
		CoffeePodTypeEspresso,
	}
)

// CoffeePod struct represents flattened coffee_pod tbl
type CoffeePod struct {
	ID      int    `json:"coffee_pod_id" db:"coffee_pod_id"`
	PodType string `json:"pod_type" db:"pod_type"`
	Name    string `json:"name" db:"name"`
}

// TODO create pod type tbl with small, large and ESPRESSO values
