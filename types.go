package ral

import (
	"image/color"
	"strings"
)

type Color struct {
	color.RGBA
	Name string
}

func FindColor(id string) *Color {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) == 5 {
		id = strings.ReplaceAll(id, "_", "")
	}
	if len(id) == 4 {
		id = "RAL" + id
	}

	if strings.HasPrefix(id, "RAL") && len(id) == 7 {
		c, ok := Classic[id]
		if !ok {
			return nil
		}

		return &c
	}

	id = strings.TrimPrefix(id, "RAL")

	if len(id) == 7 {
		h := id[0:3]
		l := id[3:5]
		c := id[5:7]

		id = "H" + h + "L" + l + "C" + c
	}

	c, ok := Design_system[id]
	if !ok {
		return nil
	}

	return &c
}
