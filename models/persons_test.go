package models

import (
	"testing"
	"time"
)

func TestPerson_tableName(t *testing.T) {
	type fields struct {
		ID          uint
		Name        string
		LastName    string
		Age         int
		Dni         int
		CreatedAt   time.Time
		ProcessedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Table name",
			fields: fields{ID: 1, Name: "Test", LastName: "Test", Age: 10, Dni: 10, CreatedAt: time.Now(), ProcessedAt: time.Now()},
			want:   "person",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Person{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				LastName:    tt.fields.LastName,
				Age:         tt.fields.Age,
				Dni:         tt.fields.Dni,
				CreatedAt:   tt.fields.CreatedAt,
				ProcessedAt: tt.fields.ProcessedAt,
			}
			if got := p.tableName(); got != tt.want {
				t.Errorf("Person.tableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
