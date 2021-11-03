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

	var (
		success    bool
		units      []string
		totalPrice int64
		metaData   int
	)

	if activeStatus {
		err, errMessage := newBuyer.UnitEligibilityCheck()

		if err {
			module.HandleError(err, 4, errMessage, newBuyer.UnitId, 0)
			// fmt.Printf("error: %v, message: %s \n", err, errMessage)
		} else {
			unitNames, price, meta := newBuyer.PurchaseUnits()
			module.HandleError(false, 2, "Purchase success", newBuyer.UnitId, int8(meta))

			success = true
			units = unitNames
			totalPrice = price
			metaData = meta
		}
	} else {
		module.HandleError(true, 4, "Status Inactive", newBuyer.UnitId, 0)
	}

	if success {
		module.PrintReceipt(int32(metaData), totalPrice, units...)
	}
}
