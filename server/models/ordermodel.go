package models

import (
	"fmt"
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"

	"go.mongodb.org/mongo-driver/bson"
)

// FetchOrder Query order list
func FetchOrder() []entitys.Order {
	fmt.Println("fetch order start...")

	clecInv, ctx := Collection("order")

	cur, err := clecInv.Aggregate(ctx, []bson.M{{
		"$lookup": bson.M{
			"from":         "orderdetail",
			"localField":   "_id",
			"foreignField": "orderid",
			"as":           "detail",
		},
	}})

	if err != nil {
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}

	// invList := make(list[string])
	var result []entitys.Order
	for cur.Next(ctx) {
		var row entitys.Order
		err := cur.Decode(&row)
		if err != nil {
			fmt.Println(err)
		}
		// row.Total_lbs, row.Total_boxes = sumOrder(row.Detail)
		result = append(result, row)
	}

	fmt.Println("fetch order end...")
	return result
}
