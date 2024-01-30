package repository

import "github.com/ankur-toko/quick-links/core/models"

type QuickLinkRepo interface {
	Save(models.QuickLink) error
	Get(string) *models.QuickLink
}
