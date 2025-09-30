package ral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindColor(t *testing.T) {
	testCases := []struct {
		ral   string
		color Color
	}{
		{
			ral:   "RAL 340 80 20",
			color: Design_system["H340L80C20"],
		},
		{
			ral:   "010 70 35",
			color: Design_system["H010L70C35"],
		},
		{
			ral:   "RAL4003",
			color: Classic["RAL4003"],
		},
		{
			ral:   "RAL 4003",
			color: Classic["RAL4003"],
		},
		{
			ral:   "4003",
			color: Classic["RAL4003"],
		},
	}
	for _, tC := range testCases {
		t.Run(tC.ral, func(t *testing.T) {
			color := FindColor(tC.ral)
			assert.NotNil(t, color)
			assert.Equal(t, &tC.color, color)
		})
	}
}
