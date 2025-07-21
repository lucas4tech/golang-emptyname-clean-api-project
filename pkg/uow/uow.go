package uow

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type UnitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func New(db *gorm.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

func (u *UnitOfWork) Begin(ctx context.Context) error {
	if u.tx != nil {
		return errors.New("transaction already started")
	}
	tx := u.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	u.tx = tx
	return nil
}

func (u *UnitOfWork) Commit() error {
	if u.tx == nil {
		return errors.New("no transaction started")
	}
	err := u.tx.Commit().Error
	u.tx = nil
	return err
}

func (u *UnitOfWork) Rollback() error {
	if u.tx == nil {
		return errors.New("no transaction started")
	}
	err := u.tx.Rollback().Error
	u.tx = nil
	return err
}

func (u *UnitOfWork) GetTx() *gorm.DB {
	if u.tx != nil {
		return u.tx
	}
	return u.db
}
