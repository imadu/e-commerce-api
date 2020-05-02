package products

import (
	"log"
	"strconv"

	"github.com/imadu/e-commerce-api/util"
	"github.com/labstack/echo"
)

//GetProductID returns the product by id
func GetProductID(context echo.Context) error {
	id := context.Param("id")

	product, err := GetProduct(id)
	if err != nil {
		log.Fatal(err)
		return util.SendError(context, "failed", "something went wrong", "400")
	}
	return util.SendData(context, product)

}

//ListCategories returns the products by categories
func ListCategories(context echo.Context) error {
	limit, _ := strconv.ParseInt(context.QueryParam("limit"), 10, 64)
	size, _ := strconv.ParseInt(context.QueryParam("size"), 10, 64)
	query := context.QueryParam("query")

	products, err := GetCategories(limit, size, query)
	if err != nil {
		log.Fatal(err)
		return util.SendError(context, "failed", "something went wrong", "500")
	}
	return util.SendData(context, products)

}

// GetAllProducts returns all the Products
func GetAllProducts(context echo.Context) error {
	limit, _ := strconv.ParseInt(context.QueryParam("limit"), 10, 64)
	size, _ := strconv.ParseInt(context.QueryParam("size"), 10, 64)

	products, err := GetProducts(limit, size)
	if err != nil {
		log.Fatal(err)
		return util.SendError(context, "failed", "something went wrong", "500")
	}
	return util.SendData(context, products)
}

// CreateNewProduct creates a new product
func CreateNewProduct(context echo.Context) error {
	product := new(Product)
	if err := context.Bind(product); err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	NewProduct, err := CreateProduct(*product)

	if err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	return util.SendData(context, NewProduct)

}

//UpdateProductDetails update a user based on id
func UpdateProductDetails(context echo.Context) error {
	id := context.Param("id")
	product := new(Product)
	if err := context.Bind(product); err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	UpdatedProduct, err := UpdateProduct(id, *product)

	if err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	return util.SendData(context, UpdatedProduct)

}

// DeleteSingleProduct deletes a user
func DeleteSingleProduct(context echo.Context) error {
	id := context.Param("id")

	DeletedID, err := DeleteProduct(id)
	if err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(context, "", "", "")
	}

	return util.SendSuccess(context, DeletedID)

}
