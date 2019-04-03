package selldb

import "fmt"

type CrossSellQueryParams struct {
	CrossSellFor string // if machine then return pod else if pod then return machine
	CrossSellForType string	// large or small or espresso
	CrossSellBasedOn string	// per flavor, per product type
	CrossSellBasedOnOrder string // smallest or largest
}

func GetCrossSellSkus(csqp *CrossSellQueryParams) (SKUs, error) {

	q := fmt.Sprintf(`
		SELECT
			s.*
		FROM 
			%[1]s.sku s
			LEFT OUTER JOIN %[1]s.coffee_pod cp ON cp.id = s.coffee_pod_id
			LEFT OUTER JOIN %[1]s.coffee_machine cm ON cm.id = s.coffee_machine_id
		WHERE 1=1`,
		DBName)

	var skus SKUs

	// TODO add validate method

	parameters := make([]interface{}, 0)


	// for pod
	if csqp.CrossSellFor == "pod" {
		q = q + "AND s.coffee_machine_id is not null and s.coffee_pod_id is null "

		if csqp.CrossSellForType != "" {
			q = q + "AND cp.pod_type = ? "
			parameters = append(parameters, csqp.CrossSellForType)
		}
	}

	// for machine
	if csqp.CrossSellFor == "machine" {
		q = q + "AND s.coffee_pod_id is not null and s.coffee_machine_id is null "

		if csqp.CrossSellFor != "" {
			q = q + "AND cm.machine_type = ? "
			parameters = append(parameters, csqp.CrossSellFor)
		}
	}

	q = q + " GROUP BY sku_id "

	orderBy := "desc"
	if csqp.CrossSellBasedOnOrder == "smallest" {
		orderBy = "asc"
	}

	if csqp.CrossSellBasedOn == "flavor" {
		q = q + " order by coffee_flavor_id " + orderBy

	} else if csqp.CrossSellBasedOn == "product_type" {
		q = q + " order by coffee_flavor_id " + orderBy
	}

	err := db.Select(&skus, q, parameters...)
	if err != nil {
		return nil, err
	}

	return skus, nil
}

func GetSkusByMachineType(machineType string) (SKUs, error) {

	q := fmt.Sprintf(`
		SELECT
			s.*
		FROM 
			%[1]s.sku s
			LEFT OUTER JOIN %[1]s.coffee_machine cm ON cm.id = s.coffee_machine_id
		WHERE 
			cm.machine_type = '?' 
		GROUP BY s.sku_id`,
		DBName)

	var skus SKUs

	err := db.Select(&skus, q, machineType)
	if err != nil {
		return nil, err
	}

	return skus, nil
}

func GetSkusByPodType(podType string) (SKUs, error) {

	q := fmt.Sprintf(`
		SELECT
			s.*
		FROM 
			%[1]s.sku s
			LEFT OUTER JOIN %[1]s.coffee_pod cp ON cp.id = s.coffee_pod_id
		WHERE 
			cp.pod_type = '?' 
		GROUP BY s.sku_id`,
		DBName)

	var skus SKUs

	err := db.Select(&skus, q, podType)
	if err != nil {
		return nil, err
	}

	return skus, nil
}