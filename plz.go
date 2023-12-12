package plz

import (
	_ "embed"
)

type PLZ struct {
	Name string
	Code []int
}

//go:generate go run ./gen/ "https://worldpostalcode.com/germany/baden-wurttemberg"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/bayern"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/berlin"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/brandenburg"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/bremen"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/hamburg"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/hessen"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/mecklenburg-vorpommern"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/niedersachsen"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/nordrhein-westfalen"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/rheinland-pfalz"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/saarland"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/sachsen"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/sachsen-anhalt"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/schleswig-holstein"
//go:generate go run ./gen/ "https://worldpostalcode.com/germany/thuringen"
//go:generate go fmt
