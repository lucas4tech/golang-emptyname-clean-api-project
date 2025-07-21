package main

import (
	"net/http"

	"app-challenge/internal/infra/api/graphql/resolver"
	"app-challenge/internal/infra/api/graphql/schema"
	"app-challenge/internal/infra/database/sqlite"
	"app-challenge/internal/infra/persistence/gorm/repository"
	"app-challenge/internal/usecase"
	"app-challenge/pkg/uow"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	dbInstance := sqlite.NewSQLiteDB("sqlite.db")
	uowInstance := uow.New(dbInstance)
	db, err := dbInstance.DB()
	if err != nil {
		logrus.Fatalf("Could not get database connection: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(dbInstance)
	productRepo := repository.NewProductRepository(dbInstance)
	orderRepo := repository.NewOrderRepository(dbInstance)

	createUserUC := &usecase.CreateUserUseCase{UserRepo: userRepo, Uow: uowInstance}
	createProductUC := &usecase.CreateProductUseCase{ProductRepo: productRepo, Uow: uowInstance}
	createOrderUC := &usecase.CreateOrderUseCase{OrderRepo: orderRepo, ProductRepo: productRepo, UserRepo: userRepo, Uow: uowInstance}
	listUsersUC := &usecase.ListUsersWithOrdersUseCase{UserRepo: userRepo}
	listProductsUC := &usecase.ListProductsUseCase{ProductRepo: productRepo}
	listOrdersUC := &usecase.ListOrdersUseCase{OrderRepo: orderRepo, UserRepo: userRepo, ProductRepo: productRepo}

	resolver := &resolver.Resolver{
		CreateUserUC:    createUserUC,
		CreateProductUC: createProductUC,
		CreateOrderUC:   createOrderUC,
		ListUsersUC:     listUsersUC,
		ListProductsUC:  listProductsUC,
		ListOrdersUC:    listOrdersUC,
	}

	r := mux.NewRouter()
	r.Handle("/graphql", schema.NewGraphQLHandler(resolver))

	logrus.Println("Server running at :8080")
	logrus.Fatal(http.ListenAndServe(":8080", r))
}
