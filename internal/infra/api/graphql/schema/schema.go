package schema

import (
	"app-challenge/internal/infra/api/graphql/resolver"
	"app-challenge/internal/usecase"
	"fmt"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/sirupsen/logrus"
)

func NewSchema(resolver *resolver.Resolver) *graphql.Schema {
	var userType *graphql.Object
	var productType *graphql.Object
	var orderType *graphql.Object
	var orderItemType *graphql.Object

	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:   "User",
		Fields: graphql.Fields{},
	})
	productType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.String},
			"name":  &graphql.Field{Type: graphql.String},
			"price": &graphql.Field{Type: graphql.Float},
			"stock": &graphql.Field{Type: graphql.Int},
		},
	})
	orderItemType = graphql.NewObject(graphql.ObjectConfig{
		Name:   "OrderItem",
		Fields: graphql.Fields{},
	})
	orderType = graphql.NewObject(graphql.ObjectConfig{
		Name:   "Order",
		Fields: graphql.Fields{},
	})

	userType.AddFieldConfig("id", &graphql.Field{Type: graphql.String})
	userType.AddFieldConfig("name", &graphql.Field{Type: graphql.String})
	userType.AddFieldConfig("email", &graphql.Field{Type: graphql.String})
	userType.AddFieldConfig("createdAt", &graphql.Field{Type: graphql.String})

	productType.AddFieldConfig("id", &graphql.Field{Type: graphql.String})
	productType.AddFieldConfig("name", &graphql.Field{Type: graphql.String})
	productType.AddFieldConfig("price", &graphql.Field{Type: graphql.Float})
	productType.AddFieldConfig("stock", &graphql.Field{Type: graphql.Int})
	productType.AddFieldConfig("createdAt", &graphql.Field{Type: graphql.String})

	orderItemType.AddFieldConfig("id", &graphql.Field{Type: graphql.String})
	orderItemType.AddFieldConfig("product", &graphql.Field{
		Type: productType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if item, ok := p.Source.(usecase.ListOrderItemResponse); ok {
				return item.Product, nil
			}
			if item, ok := p.Source.(*usecase.ListOrderItemResponse); ok {
				return item.Product, nil
			}
			if item, ok := p.Source.(usecase.CreateOrderItemResponse); ok {
				return item.Product, nil
			}
			if item, ok := p.Source.(*usecase.CreateOrderItemResponse); ok {
				return item.Product, nil
			}
			if item, ok := p.Source.(usecase.OrderItemResponse); ok {
				return item.Product, nil
			}
			if item, ok := p.Source.(*usecase.OrderItemResponse); ok {
				return item.Product, nil
			}
			return nil, nil
		},
	})
	orderItemType.AddFieldConfig("quantity", &graphql.Field{Type: graphql.Int})
	orderItemType.AddFieldConfig("price", &graphql.Field{Type: graphql.Float})
	orderItemType.AddFieldConfig("createdAt", &graphql.Field{Type: graphql.String})

	orderType.AddFieldConfig("id", &graphql.Field{Type: graphql.String})
	orderType.AddFieldConfig("user", &graphql.Field{
		Type: userType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if order, ok := p.Source.(usecase.ListOrderResponse); ok {
				return order.User, nil
			}
			if order, ok := p.Source.(*usecase.ListOrderResponse); ok {
				return order.User, nil
			}
			if order, ok := p.Source.(usecase.CreateOrderResponse); ok {
				return order.User, nil
			}
			if order, ok := p.Source.(*usecase.CreateOrderResponse); ok {
				return order.User, nil
			}
			return nil, nil
		},
	})
	orderType.AddFieldConfig("items", &graphql.Field{
		Type: graphql.NewList(orderItemType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if order, ok := p.Source.(usecase.ListOrderResponse); ok {
				fmt.Println()
				return order.Items, nil
			}
			if order, ok := p.Source.(*usecase.ListOrderResponse); ok {
				return order.Items, nil
			}
			if order, ok := p.Source.(usecase.CreateOrderResponse); ok {
				return order.Items, nil
			}
			if order, ok := p.Source.(*usecase.CreateOrderResponse); ok {
				return order.Items, nil
			}
			return nil, nil
		},
	})
	orderType.AddFieldConfig("total", &graphql.Field{
		Type: graphql.Float,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if order, ok := p.Source.(usecase.ListOrderResponse); ok {
				return order.Total, nil
			}
			if order, ok := p.Source.(*usecase.ListOrderResponse); ok {
				return order.Total, nil
			}
			if order, ok := p.Source.(usecase.CreateOrderResponse); ok {
				return order.Total, nil
			}
			if order, ok := p.Source.(*usecase.CreateOrderResponse); ok {
				return order.Total, nil
			}
			return nil, nil
		},
	})
	orderType.AddFieldConfig("createdAt", &graphql.Field{Type: graphql.String})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:    graphql.NewList(userType),
				Resolve: resolver.Users,
			},
			"products": &graphql.Field{
				Type:    graphql.NewList(productType),
				Resolve: resolver.Products,
			},
			"orders": &graphql.Field{
				Type:    graphql.NewList(orderType),
				Resolve: resolver.Orders,
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: resolver.CreateUser,
			},
			"createProduct": &graphql.Field{
				Type: productType,
				Args: graphql.FieldConfigArgument{
					"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"price": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
					"stock": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: resolver.CreateProduct,
			},
			"createOrder": &graphql.Field{
				Type: orderType,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"items": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.NewInputObject(graphql.InputObjectConfig{
							Name: "OrderItemInput",
							Fields: graphql.InputObjectConfigFieldMap{
								"productId": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
								"quantity":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
							},
						})),
					},
				},
				Resolve: resolver.CreateOrder,
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	if err != nil {
		logrus.Fatal("Failed to create GraphQL schema", "error", err)
		os.Exit(1)
	}

	return &schema
}

func NewGraphQLHandler(resolver *resolver.Resolver) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   NewSchema(resolver),
		Pretty:   true,
		GraphiQL: true,
	})
}
