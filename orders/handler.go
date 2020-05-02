package orders

import (
	"log"
	"strconv"

	"github.com/imadu/e-commerce-api/util"
	"github.com/labstack/echo"
)

//GetOrderID returns the product by id
func GetOrderID(context echo.Context) error {
	id := context.Param("id")

	order, err := GetOrder(id)
	if err != nil {
		log.Fatal(err)
		return util.SendError(context, "failed", "something went wrong", "400")
	}
	return util.SendData(context, order)

}

// GetAllOrders returns all the Products
func GetAllOrders(context echo.Context) error {
	limit, _ := strconv.ParseInt(context.QueryParam("limit"), 10, 64)
	size, _ := strconv.ParseInt(context.QueryParam("size"), 10, 64)

	orders, err := GetOrders(limit, size)
	if err != nil {
		log.Fatal(err)
		return util.SendError(context, "failed", "something went wrong", "500")
	}
	return util.SendData(context, orders)
}

// CreateNewOrder creates a new product
func CreateNewOrder(context echo.Context) error {
	order := new(Order)
	if err := context.Bind(order); err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	NewOrder, err := Create(*order)

	if err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	return util.SendData(context, NewOrder)

}

//UpdateOrderDetails update a user based on id
func UpdateOrderDetails(context echo.Context) error {
	id := context.Param("id")
	order := new(Order)
	if err := context.Bind(order); err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	UpdatedOrder, err := UpdateOrder(id, *order)

	if err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	return util.SendData(context, UpdatedOrder)

}

// CancelOrder deletes an order
func CancelOrder(context echo.Context) error {
	id := context.Param("id")

	DeletedID, err := DeleteOrder(id)
	if err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	return util.SendSuccess(context, DeletedID)

}
