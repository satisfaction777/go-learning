package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcBalance(t *testing.T) {
	tests := []struct {
		name string
		income float64
		food float64
		transport float64
		expected float64
	}{
		{"обычный случай", 10000, 3000, 2000, 5000},
		{"нулевые расходы", 10000, 0, 0, 10000},
		{"расходы больше дохода", 5000, 3000, 3000, -1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calcBalance(tt.income, tt.food, tt.transport)
			assert.Equal(t, tt.expected, result)
		})
	}

}


func TestCanAfford(t *testing.T) {
    tests := []struct {
        name     string
        balance  float64
        price    float64
        wantOk   bool
        wantDiff float64
    }{
        {"денег хватило сполна", 10000, 3000, true, 7000},
        {"денег хватило в край", 10000, 8000, true, 2000},
        {"денег не хватило",     5000,  6000, false, 1000},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ok, diff := canAfford(tt.balance, tt.price)
            assert.Equal(t, tt.wantOk, ok)
            assert.Equal(t, tt.wantDiff, diff)
        })
    }
}

func TestValidateAge(t *testing.T) {
assert.NoError(t, validateAge(25))
assert.Error(t, validateAge(-1))
assert.Error(t, validateAge(167))
}

