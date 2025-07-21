package sqlite

import (
	"app-challenge/internal/infra/persistence/gorm/model"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("failed to connect database: %v", err)
	}
	err = db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.OrderItem{})
	if err != nil {
		logrus.Fatalf("failed to migrate database: %v", err)
	}
	return db
}

func Seed(db *gorm.DB) {
	users := []model.User{
		{Name: "John Doe", Email: "johndoe@example.com"},
		{Name: "Mary Doe", Email: "marydoe@example.com"},
	}
	for _, u := range users {
		db.FirstOrCreate(&u, model.User{Email: u.Email})
	}

	products := []model.Product{
		{Name: "Notebook", Price: 3500.00, Stock: 10},
		{Name: "Mouse", Price: 100.00, Stock: 50},
		{Name: "Keyboard", Price: 200.00, Stock: 30},
	}
	for _, p := range products {
		db.FirstOrCreate(&p, model.Product{Name: p.Name})
	}
}
