package main

import (
	"fmt"

	module "github.com/patrickbyan/sysacademy_module_project/v2"
)

const (
	buyerName = "Belinda"
	bornDate  = 1635875790371
)

func main() {
	fmt.Println("Begin Transaction \n=====================")
	defer fmt.Println("=====================\nEnd Transaction")

	activeStatus, discrepancy, unit := module.GetDataBuyer(buyerName, bornDate)

	if activeStatus {
		newBuyer := module.DataBuyer{
			Name:         buyerName,
			ActiveStatus: activeStatus,
			Discrepancy:  int64(discrepancy),
			UnitId:       unit,
		}

		err, errMessage := newBuyer.UnitEligibilityCheck()

		if err {
			fmt.Printf("error: %v, message: %s \n", err, errMessage)
		} else {
			unitNames, totalPrice, metaData := newBuyer.PurchaseUnits()
			fmt.Printf("error: false, message: purchase %v units (total: %v data) with a price of %v success \n", unitNames, metaData, totalPrice)
		}
	} else {
		fmt.Println("error: true, message: You are born literally today")
	}
}
