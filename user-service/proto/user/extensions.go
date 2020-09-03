package user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// BeforeCreate -
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("Id", uuid.String())
}