package internal

import (
	"fmt"
)

type Device interface {
	Gelis()
	İcerik()
}

type Phone struct {
	Markasi string
	Modeli  string
}

func (p Phone) Gelis() string {
	return "Trrrrrrr...\n"
}

func (p Phone) İcerik() string {
	return fmt.Sprintf("Telefonun Markası:%s\nTelefonun Modeli:%s\n", p.Markasi, p.Modeli)
}
