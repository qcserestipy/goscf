package system

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


var PSE = map[string]Element{
    "H":  {"Hydrogen", "H", 1, 1, 1, 1.008, "Nonmetal", 1},
    "He": {"Helium", "He", 2, 1, 18, 4.0026, "Noble gas", 2},
    "Li": {"Lithium", "Li", 3, 2, 1, 6.94, "Alkali metal", 1},
    "Be": {"Beryllium", "Be", 4, 2, 2, 9.0122, "Alkaline earth metal", 2},
    "B":  {"Boron", "B", 5, 2, 13, 10.81, "Metalloid", 3},
    "C":  {"Carbon", "C", 6, 2, 14, 12.011, "Nonmetal", 4},
    "N":  {"Nitrogen", "N", 7, 2, 15, 14.007, "Nonmetal", 5},
    "O":  {"Oxygen", "O", 8, 2, 16, 15.999, "Nonmetal", 6},
    "F":  {"Fluorine", "F", 9, 2, 17, 18.998, "Halogen", 7},
    "Ne": {"Neon", "Ne", 10, 2, 18, 20.180, "Noble gas", 8},
    "Na": {"Sodium", "Na", 11, 3, 1, 22.990, "Alkali metal", 1},
    "Mg": {"Magnesium", "Mg", 12, 3, 2, 24.305, "Alkaline earth metal", 2},
    "Al": {"Aluminum", "Al", 13, 3, 13, 26.982, "Post-transition metal", 3},
    "Si": {"Silicon", "Si", 14, 3, 14, 28.085, "Metalloid", 4},
    "P":  {"Phosphorus", "P", 15, 3, 15, 30.974, "Nonmetal", 5},
    "S":  {"Sulfur", "S", 16, 3, 16, 32.06, "Nonmetal", 6},
    "Cl": {"Chlorine", "Cl", 17, 3, 17, 35.45, "Halogen", 7},
    "Ar": {"Argon", "Ar", 18, 3, 18, 39.948, "Noble gas", 8},
}