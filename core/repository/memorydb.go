package repository

import "github.com/ankur-toko/quick-links/core/models"

type MemoryDB struct {
	m map[string]string
}

func GetMemoryDB() QuickLinkRepo {
	return &MemoryDB{map[string]string{}}
}

func (db *MemoryDB) Save(r models.QuickLink) error {
	db.m[r.Key] = r.URL
	return nil
}

func (db *MemoryDB) Get(key string) *models.QuickLink {
	v, ok := db.m[key]
	if ok {
		return &models.QuickLink{Key: key, URL: v}
	} else {
		return nil
	}
}
