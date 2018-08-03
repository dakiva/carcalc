package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/dakiva/carcalc"
)

func main() {
	log.SetFlags(0)
	log.Println("Carcalc (v1.0)")
	leaseFile := flag.String("lease", "", "Lease json file path")
	financeFile := flag.String("finance", "", "Finance json file path")
	flag.Parse()

	if *leaseFile != "" {
		if body, err := ioutil.ReadFile(*leaseFile); err == nil {
			leaseInfo := &carcalc.LeaseInfo{}
			if jsonErr := json.Unmarshal(body, leaseInfo); jsonErr == nil && leaseInfo.Valid() {
				log.Println("Adjusted Capitalized Cost", carcalc.Round(leaseInfo.AdjustedCost(), 2))
				log.Println("Rent charge", carcalc.Round(leaseInfo.RentCharge(), 2))
				log.Println("Pretax payment", carcalc.Round(leaseInfo.PretaxPayment(), 2))
				log.Println("Final payment", carcalc.Round(leaseInfo.Payment(), 2))
			} else {
				log.Fatalf("lease JSON is invalid or incomplete: %v", jsonErr)
			}
		} else {
			log.Fatalf("Error reading from file: %v", err)
		}
	} else if *financeFile != "" {
		if body, err := ioutil.ReadFile(*financeFile); err == nil {
			financeInfo := &carcalc.FinanceInfo{}
			if jsonErr := json.Unmarshal(body, financeInfo); jsonErr == nil && financeInfo.Valid() {
				log.Println("Principal", carcalc.Round(financeInfo.Principal(), 2))
				log.Println("Final payment", carcalc.Round(financeInfo.Payment(), 2))
			} else {
				log.Fatalf("finance JSON is invalid or incomplete: %v", jsonErr)
			}
		} else {
			log.Fatalf("Error reading from file: %v", err)
		}
	} else {
		log.Fatalf("One of lease or finance is required.")
	}
}
