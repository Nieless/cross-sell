package sellapi

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"cross-sell/selldb"
	"net/http"
)

//GetMachineSKUsByType will return all the skus for machine by passed machine_type
//
//Request Type: GET
//
//URL: /machines/{machine_type}/skus
//
func GetMachineSKUsByType(w http.ResponseWriter, r *http.Request) {
	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the variables from the request
	vars := mux.Vars(r)
	machineType := vars["machine_type"]

	// validate machine type
	if len(machineType) == 0 {
		err := fmt.Errorf("machine type query param mandatory")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var isMachineTypeValid bool
	for _, v := range selldb.CoffeeMachineTypes {
		if machineType != v {
			continue
		}
		isMachineTypeValid = true
	}

	if !isMachineTypeValid {
		err := fmt.Errorf("passed machine type is invalid")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	skus, err := selldb.GetSkusByMachineType(machineType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Encode array returned into JSON and return
	if err := json.NewEncoder(w).Encode(skus); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding JSON response")
	}
}

//GetPodSKUsByPodType will return all the SKUs for pods by passed pod_type
//
//Request Type: GET
//
//URL: /machines/{pod_type}/skus
//
func GetPodSKUsByPodType(w http.ResponseWriter, r *http.Request) {
	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the variables from the request
	vars := mux.Vars(r)
	podType := vars["pod_type"]

	// validate pod type
	if len(podType) == 0 {
		err := fmt.Errorf("pod type query param mandatory")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var isPodTypeValid bool
	for _, v := range selldb.CoffeePodTypes {
		if podType != v {
			continue
		}
		isPodTypeValid = true
	}

	if !isPodTypeValid {
		err := fmt.Errorf("passed pod type is invalid")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pods, err := selldb.GetSkusByPodType(podType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Encode array returned into JSON and return
	if err := json.NewEncoder(w).Encode(pods); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding JSON response")
	}
}

//GetCrossSellSKUsByQueryParam will return all the SKUs from db by passed query params
//
//Request Type: GET
//
//URL: /crosssell/skus
//
//Query parameters:
//     sell_for
//         description: machine or pod
//         type: string
//     sell_type
//         description: large or small or espresso
//         type: string
//     sell_per
//         description: per flavor or per product type
//         type: string
//     sell_order
//         description: smallest or largest
//         type: string
//
func GetCrossSellSKUsByQueryParam(w http.ResponseWriter, r *http.Request) {
	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// FixME : validate query params or allow ids from db
	// get params
	qp := &selldb.CrossSellQueryParams{
		CrossSellFor:          r.FormValue("sell_for"),   // if machine then return pod else if pod then return machine
		CrossSellForType:      r.FormValue("sell_type"),  // large or small or espresso
		CrossSellBasedOn:      r.FormValue("sell_per"),   // per flavor, per product type
		CrossSellBasedOnOrder: r.FormValue("sell_order"), // smallest or largest
	}

	skus, err := selldb.GetCrossSellSkus(qp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Encode array returned into JSON and return
	if err := json.NewEncoder(w).Encode(skus); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding JSON response")
	}
}
