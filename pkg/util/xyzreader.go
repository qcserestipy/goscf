package main

func readxyzfile(filename string) ([]Atom, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer
	file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	var atoms []Atom
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 4 {
			return nil, fmt.Errorf("invalid line: %v", fields)
		}
		element := fields[0]
		x, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return nil, err
		}
		z, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return nil, err
		}
		atoms = append(atoms, Atom{
			Position: Point{X: x, Y: y, Z: z},
			Element:  element,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return atoms, nil
}