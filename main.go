package main

import (
	"fmt"

	module_project "github.com/patrickbyan/sysacademy_module_project"
)

const (
	buyerName = "Belinda"
	bornDate  = 1635875790371
)

func main() {
	fmt.Println("Begin Transaction \n=====================")
	defer fmt.Println("=====================\nEnd Transaction")

	activeStatus, discrepancy, unit := module_project.GetDataBuyer(buyerName, bornDate)

	if activeStatus {
		err, errMessage := module_project.UnitEligibilityCheck(unit...)

		if err {
			fmt.Printf("error: %v, message: %s \n", err, errMessage)
		} else {
			newBuyer := module_project.DataBuyer{
				Name:         buyerName,
				ActiveStatus: activeStatus,
				Discrepancy:  int64(discrepancy),
			}

			unitNames, totalPrice, metaData := newBuyer.PurchaseUnits(unit...)
			fmt.Printf("error: false, message: purchase %v units (total: %v data) with a price of %v success \n", unitNames, metaData, totalPrice)
		}
	} else {
		fmt.Println("error: true, message: You are born literally today")
	}
}
