// package main

// import (
// 	"fmt"
// 	"log"
// 	"math"

// 	"gonum.org/v1/gonum/mat"
// )

// func main() {
// 	// Initialize dummy core Hamiltonian (Hcore) matrix
// 	HcoreData := []float64{
// 		1.0, 0.1,
// 		0.1, 0.5,
// 	}
// 	Hcore := mat.NewDense(2, 2, HcoreData)
// 	fmt.Printf("Core Hamiltonian (Hcore):\n%v\n\n", mat.Formatted(Hcore, mat.Prefix(""), mat.Excerpt(0)))

// 	// Initialize dummy overlap (S) matrix
// 	SData := []float64{
// 		1.0, 0.2,
// 		0.2, 1.0,
// 	}
// 	S := mat.NewDense(2, 2, SData)
// 	fmt.Printf("Overlap matrix (S):\n%v\n\n", mat.Formatted(S, mat.Prefix(""), mat.Excerpt(0)))

// 	// SCF loop parameters
// 	maxIterations := 50
// 	convergenceThreshold := 1e-6

// 	// Initialize dummy Fock matrix as a copy of Hcore
// 	F := mat.DenseCopyOf(Hcore)

// 	// Initialize placeholder density matrix (P)
// 	P := mat.NewDense(2, 2, nil)

// 	var energy, prevEnergy float64

// 	// Begin the SCF iterations
// 	for iter := 0; iter < maxIterations; iter++ {
// 		// In a real HF SCF, you would:
// 		// 1. Diagonalize the Fock matrix (with proper orthogonalization using S)
// 		// 2. Form the new density matrix (P) from the occupied orbitals
// 		// 3. Compute the electronic energy
// 		// 4. Build the new Fock matrix using the updated density matrix

// 		// For demonstration, we use a dummy energy evaluation function
// 		energy = dummyEnergy(F, P)
// 		fmt.Printf("Iteration %d: Energy = %.8f\n", iter+1, energy)

// 		// Check for convergence
// 		if iter > 0 && math.Abs(energy-prevEnergy) < convergenceThreshold {
// 			fmt.Println("SCF converged!")
// 			break
// 		}
// 		prevEnergy = energy

// 		// Dummy update for Fock matrix: in a real implementation,
// 		// update F based on the current density matrix and electron repulsion integrals.
// 		updateFockMatrix(F, P)
// 	}

// 	log.Println("goscf run completed.")
// }

// // dummyEnergy calculates a simple energy from the diagonal of F.
// // Replace this with your actual energy evaluation routine.
// func dummyEnergy(F, P *mat.Dense) float64 {
// 	var energy float64
// 	rows, _ := F.Dims()
// 	for i := 0; i < rows; i++ {
// 		energy += F.At(i, i)
// 	}
// 	return energy
// }

// // updateFockMatrix performs a dummy update on F.
// // Replace with the proper Fock matrix construction using electron repulsion integrals.
// func updateFockMatrix(F, P *mat.Dense) {
// 	rows, cols := F.Dims()
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			F.Set(i, j, F.At(i, j)+0.01)
// 		}
// 	}
// }



package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/qcserestipy/goscf/pkg/util"
	"github.com/qcserestipy/goscf/pkg/system"
)

func init(){
	logrus.SetFormatter(&logrus.JSONFormatter{

		TimestampFormat: time.RFC3339,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

// MoleculeRequest defines the structure for sending molecule information.
type MoleculeRequest struct {
	Atom  string `json:"atom"`
	Basis string `json:"basis"`
}

// IntegralResponse defines the expected structure from the microservice.
type IntegralResponse struct {
	OverlapIntegrals [][]float64 `json:"overlap_integrals"`
	KineticIntegrals [][]float64 `json:"kinetic_integrals"`
}

func main() {
	url := "http://localhost:5000/compute_integrals"

	var xyzfile string = "h2.xyz"
	atoms, err := util.ReadXYZFile(xyzfile)
	if err != nil {
		logrus.Fatalf("Error reading XYZ file: %v", err)
	}
	logrus.Infof("Read %d atoms from %s", len(atoms), xyzfile)
	var Geometry system.Geometry = system.Geometry{Atoms: atoms}
	logrus.Infof("Center of mass: %v", Geometry.CenterOfMass())
	logrus.Infof("Core-core repulsion: %v", Geometry.CoreCoreRepulsion())

	// Define the molecule parameters.
	reqData := MoleculeRequest{
		Atom:  "H 0 0 0; H 0 0 0.74", // Example: H₂ molecule
		Basis: "sto-3g",
	}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		logrus.Fatalf("Error marshalling request: %v", err)
	}

	// Send a POST request to the microservice.
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	// Read and decode the response.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("Error reading response: %v", err)
	}

	var res IntegralResponse
	if err := json.Unmarshal(body, &res); err != nil {
		logrus.Fatalf("Error unmarshalling response: %v", err)
	}

	// Print the received integrals.
	fmt.Printf("Overlap Integrals:\n%v\n", res.OverlapIntegrals)
	fmt.Printf("Kinetic Integrals:\n%v\n", res.KineticIntegrals)
}