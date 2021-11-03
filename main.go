package main

import (
	"fmt"

	module "github.com/patrickbyan/sysacademy_module_project/v2"
)

const (
	buyerName = "Patrick"
	bornDate  = 1635875790371
)

func main() {
	fmt.Println("Begin Transaction \n=====================")
	defer fmt.Println("=====================\nEnd Transaction")

	activeStatus, discrepancy, unit := module.GetDataBuyer(buyerName, bornDate)
	newBuyer := module.DataBuyer{
		Name:         buyerName,
		ActiveStatus: activeStatus,
		Discrepancy:  int64(discrepancy),
		UnitId:       unit,
	}

	if activeStatus {
		err, errMessage := newBuyer.UnitEligibilityCheck()

		if err {
			module.HandleError(err, 4, errMessage, newBuyer.UnitId, 0)
			// fmt.Printf("error: %v, message: %s \n", err, errMessage)
		} else {
			unitNames, totalPrice, metaData := newBuyer.PurchaseUnits()
			module.HandleError(false, 2, "Purchase success", newBuyer.UnitId, int8(metaData))
			fmt.Printf("purchase %v units (total: %v data) with a price of %v success \n", unitNames, metaData, totalPrice)
		}
	} else {
		module.HandleError(true, 4, "Status Inactive", newBuyer.UnitId, 0)
	}
}
