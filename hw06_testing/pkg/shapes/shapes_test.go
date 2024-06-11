package shapes

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name    string
		shape   Shape
		want    float64
		wantErr bool
	}{
		{"Круг", Circle{Radius: 5}, 78.53981633974483, false},
		{"Круг с отрицательным значением радиуса", Circle{Radius: -5}, 0, true},
		{"Прямоугольник", Rectangle{Width: 10, Height: 5}, 50, false},
		{"Прямоугольник с отрицательным значением ширины", Rectangle{Width: -10, Height: 5}, 0, true},
		{"Треугольник", Triangle{Base: 8, Height: 6}, 24, false},
		{"Треугольник с отрицательным значением основания", Triangle{Base: -8, Height: 6}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateArea(tt.shape)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateArea() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateArea() got = %v, want %v", got, tt.want)
			}
		})
	}
}
