package resolver

import (
	"app-challenge/internal/usecase"

	"github.com/graphql-go/graphql"
)

type Resolver struct {
	CreateUserUC    *usecase.CreateUserUseCase
	CreateProductUC *usecase.CreateProductUseCase
	CreateOrderUC   *usecase.CreateOrderUseCase
	ListUsersUC     *usecase.ListUsersWithOrdersUseCase
	ListProductsUC  *usecase.ListProductsUseCase
	ListOrdersUC    *usecase.ListOrdersUseCase
}

func (r *Resolver) Users(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.ListUsersUC.Execute(p.Context, usecase.ListUsersWithOrdersRequest{Limit: 100, Offset: 0})
	if err != nil {
		return nil, err
	}
	return resp.Users, nil
}

func (r *Resolver) Products(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.ListProductsUC.Execute(p.Context, usecase.ListProductsRequest{Limit: 100, Offset: 0})
	if err != nil {
		return nil, err
	}
	return resp.Products, nil
}

func (r *Resolver) Orders(p graphql.ResolveParams) (interface{}, error) {
	resp, err := r.ListOrdersUC.Execute(p.Context, usecase.ListOrdersRequest{Limit: 100, Offset: 0})
	if err != nil {
		return nil, err
	}
	return resp.Orders, nil
}

func (r *Resolver) CreateUser(p graphql.ResolveParams) (interface{}, error) {
	name, _ := p.Args["name"].(string)
	email, _ := p.Args["email"].(string)
	resp, err := r.CreateUserUC.Execute(p.Context, usecase.CreateUserRequest{Name: name, Email: email})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Resolver) CreateProduct(p graphql.ResolveParams) (interface{}, error) {
	name, _ := p.Args["name"].(string)
	price, _ := p.Args["price"].(float64)
	stock, _ := p.Args["stock"].(int)
	resp, err := r.CreateProductUC.Execute(p.Context, usecase.CreateProductRequest{Name: name, Price: price, Stock: stock})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Resolver) CreateOrder(params graphql.ResolveParams) (interface{}, error) {
	var items []usecase.CreateOrderItemRequest
	for _, item := range params.Args["items"].([]interface{}) {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		productId, _ := itemMap["productId"].(string)
		quantity, _ := itemMap["quantity"].(int)
		items = append(items, usecase.CreateOrderItemRequest{
			ProductID: productId,
			Quantity:  quantity,
		})
	}

	req := usecase.CreateOrderRequest{
		UserID: params.Args["userId"].(string),
		Items:  items,
	}

	resp, err := r.CreateOrderUC.Execute(params.Context, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Resolver) UserOrders(p graphql.ResolveParams) (interface{}, error) {
	user, ok := p.Source.(usecase.UserWithOrdersResponse)
	if !ok {
		return nil, nil
	}
	return user.Orders, nil
}

func (r *Resolver) OrderUser(p graphql.ResolveParams) (interface{}, error) {
	order, ok := p.Source.(usecase.OrderResponse)
	if !ok {
		return nil, nil
	}
	return order, nil
}

func (r *Resolver) OrderItems(p graphql.ResolveParams) (interface{}, error) {
	order, ok := p.Source.(usecase.OrderResponse)
	if !ok {
		return nil, nil
	}
	return order, nil
}

func (r *Resolver) OrderItemProduct(p graphql.ResolveParams) (interface{}, error) {
	item, ok := p.Source.(usecase.OrderItemResponse)
	if !ok {
		return nil, nil
	}
	return item, nil
}
