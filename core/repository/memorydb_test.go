package repository

import (
	"testing"

	"github.com/ankur-toko/quick-links/core/models"
	"github.com/gookit/goutil/testutil/assert"
)

func TestGetMemoryDB(t *testing.T) {

	r := GetMemoryDB()
	assert.NotNil(t, r)

	assert.True(t, r.(*MemoryDB).m != nil)
}

func TestMemoryDB_Save(t *testing.T) {
	type fields struct {
		m map[string]string
	}
	type args struct {
		r models.QuickLink
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Save Data Test 1",
			fields{
				map[string]string{},
			},
			args{models.QuickLink{"key", "value"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &MemoryDB{
				m: tt.fields.m,
			}
			if err := db.Save(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("MemoryDB.Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.True(t, len(db.m) > 0)
		})
	}
}

func TestMemoryDB_Get(t *testing.T) {
	db := GetMemoryDB()
	db.Save(models.QuickLink{Key: "key", URL: "url"})

	val := db.Get("key")

	assert.True(t, val.Key == "key")
	assert.True(t, val.URL == "url")

	val = db.Get("random")

	assert.True(t, val == nil)
}
