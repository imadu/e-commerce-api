package main

import (
	"log"
	"net/http"
	"os"

	"github.com/imadu/e-commerce-api/auth"
	"github.com/imadu/e-commerce-api/db"
	"github.com/imadu/e-commerce-api/orders"
	"github.com/imadu/e-commerce-api/products"
	"github.com/imadu/e-commerce-api/users"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/subosito/gotenv"
)

func main() {
	err := gotenv.Load()
	if err != nil {
		log.Fatalf("Could not load env file: %+v", err)
	}

	db.InitDB()

	mux := echo.New()
	mux.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	mux.Use(middleware.Logger())
	mux.Use(middleware.Recover())

	mux.POST("/login", auth.Login)

	mux.GET("/products/:id", products.GetProductID)
	mux.GET("/products/", products.GetAllProducts)

	router := mux.Group("/products")
	router.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	router.POST("/products/", products.CreateNewProduct)
	router.PATCH("/products/:id", products.UpdateProductDetails)
	router.DELETE("/products/:id", products.DeleteSingleProduct)

	mux.POST("/users/", users.CreateNewUser)

	userRouter := mux.Group("/users")
	userRouter.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	userRouter.GET("/users/:id", users.GetUserID)
	userRouter.GET("/users/", users.GetAllUsers)
	userRouter.PATCH("/users/:id", users.UpdateUserDetails)
	userRouter.DELETE("/users/:id", users.DeleteUserAccount)

	mux.GET("/orders/:id", orders.GetOrderID)
	mux.GET("/orders/", orders.GetAllOrders)
	mux.POST("/orders/", orders.CreateNewOrder)
	mux.PATCH("/orders/:id", orders.UpdateOrderDetails)
	mux.DELETE("/orders/:id", orders.CancelOrder)

	mux.Logger.Fatal(mux.Start(":1323"))

}
