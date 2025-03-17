package util

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"

	"github.com/qcserestipy/goscf/pkg/system"
)


func ReadXYZFile(filename string) ([]system.Atom, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Read first line: number of atoms
    if !scanner.Scan() {
        return nil, fmt.Errorf("XYZ file is empty or missing the number of atoms line")
    }
    _, err = strconv.Atoi(scanner.Text()) // parse the number of atoms if you need it
    if err != nil {
        return nil, fmt.Errorf("invalid number of atoms in first line")
    }

    // Read second line (comment) - we ignore its contents
    if !scanner.Scan() {
        return nil, fmt.Errorf("XYZ file missing second line (comment)")
    }

    var atoms []system.Atom
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
		elementInfo := system.PSE[element]
		atom := system.NewAtom(element, x, y, z, elementInfo.Weight, elementInfo.Number)
        atoms = append(atoms, atom)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return atoms, nil
}