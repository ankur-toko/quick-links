package models

import "testing"

func TestQuickLink_ToJSON(t *testing.T) {
	type fields struct {
		Key string
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"To JSON Test",
			fields{Key: "key-a", URL: "url-a"},
			"{\"key\":\"key-a\",\"url\":\"url-a\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QuickLink{
				Key: tt.fields.Key,
				URL: tt.fields.URL,
			}
			if got := q.ToJSON(); got != tt.want {
				t.Errorf("QuickLink.ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
