package core

import (
	"fmt"

	"github.com/ankur-toko/quick-links/core/models"
	"github.com/ankur-toko/quick-links/core/repository"
	"github.com/ankur-toko/quick-links/core/validations"
)

type Core interface {
	SaveQuickLink(models.QuickLink) error
	GetQuickLink(key string) *models.QuickLink
}

type BaseCore struct {
	AppName     string
	DB          repository.QuickLinkRepo
	Validations validations.ValidationCheck
}

func CreateCoreClassObject() (Core, error) {
	bc := BaseCore{}
	bc.Validations = validations.DefaultValidatorObj()
	db, e := repository.GetFileDB(repository.GetMemoryDB())
	if e != nil {
		return nil, e
	}
	bc.DB = db
	return &bc, nil
}

var AppName string

func (c *BaseCore) SaveQuickLink(record models.QuickLink) error {
	e := c.Validations.Check(record)
	if e == nil {
		c.DB.Save(record)
		fmt.Printf("save key:%v and url:%v successful", record.Key, record.URL)
	} else {
		return e
	}
	return nil
}

func (c *BaseCore) GetQuickLink(key string) *models.QuickLink {
	if v := c.DB.Get(key); v != nil {
		return v
	} else {
		return nil
	}
}
