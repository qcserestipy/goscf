package main

// Implement periodic system of elements (PSE) for first two rows

// Element represents a chemical element
type Element struct {
	Name    string
	Symbol  string
	Number  int
	Row     int
	Col     int
	Weight  float64
	Group   string
	Valence int
}

// PSE is the periodic system of elements
var PSE = []Element{
	{"Hydrogen", "H", 1, 1, 1, 1.008, "Nonmetal", 1},
	{"Helium", "He", 2, 1, 18, 4.0026, "Noble gas", 2},
	{"Lithium", "Li", 3, 2, 1, 6.94, "Alkali metal", 1},
	{"Beryllium", "Be", 4, 2, 2, 9.0122, "Alkaline earth metal", 2},
	{"Boron", "B", 5, 2, 13, 10.81, "Metalloid", 3},
	{"Carbon", "C", 6, 2, 14, 12.011, "Nonmetal", 4},
	{"Nitrogen", "N", 7, 2, 15, 14.007, "Nonmetal", 5},
	{"Oxygen", "O", 8, 2, 16, 15.999, "Nonmetal", 6},
	{"Fluorine", "F", 9, 2, 17, 18.998, "Halogen", 7},
	{"Neon", "Ne", 10, 2, 18, 20.180, "Noble gas", 8},
	{"Sodium", "Na", 11, 3, 1, 22.990, "Alkali metal", 1},
	{"Magnesium", "Mg", 12, 3, 2, 24.305, "Alkaline earth metal", 2},
	{"Aluminum", "Al", 13, 3, 13, 26.982, "Post-transition metal", 3},
	{"Silicon", "Si", 14, 3, 14, 28.085, "Metalloid", 4},
	{"Phosphorus", "P", 15, 3, 15, 30.974, "Nonmetal", 5},
	{"Sulfur", "S", 16, 3, 16, 32.06, "Nonmetal", 6},
	{"Chlorine", "Cl", 17, 3, 17, 35.45, "Halogen", 7},
	{"Argon", "Ar", 18, 3, 18, 39.948, "Noble gas", 8},
}