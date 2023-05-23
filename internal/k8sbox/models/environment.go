package model

import (
	"fmt"
	"os"

	"github.com/twelvee/k8sbox/internal/k8sbox"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
)

func RunEnvironment(tomlFile string) error {
	environment, runDirectory := lookForEnvironmentStep(tomlFile)
	isSaved := checkIfEnvironmentIsSavedStep(environment)
	validateEnvironmentStep(environment)
	validateBoxesStep(&environment, runDirectory)
	if isSaved {
		checkIfEnvironmentHasSameBoxesStep(&environment)
	}
	createTempDeployDirectoryStep(&environment, runDirectory, isSaved)
	deployEnvironmentStep(&environment, isSaved)

	fmt.Println("Aight we're done here!")
	return nil
}

func lookForEnvironmentStep(tomlFile string) (structs.Environment, string) {
	fmt.Print("Looking for environment...")
	environment, runDirectory, err := k8sbox.GetTomlFormatter().GetEnvironmentFromToml(tomlFile)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
	return environment, runDirectory
}

func checkIfEnvironmentIsSavedStep(environment structs.Environment) bool {
	fmt.Print("Matching with already saved environments...")
	saved, err := utils.IsEnvironmentSaved(environment.Id)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	if saved {
		fmt.Println(" OK - SAVED")
		return true
	}
	fmt.Println(" OK - NEW")
	return false
}

func checkIfEnvironmentHasSameBoxesStep(environment *structs.Environment) {
	fmt.Print("Matching boxes on founded environment...")
	savedEnvironment, err := utils.GetEnvironment(environment.Id)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
	if len(savedEnvironment.Boxes) > 0 {
		fmt.Printf("Found %d legacy boxes. Removing...", len(savedEnvironment.Boxes))
	}
	for _, savedBox := range savedEnvironment.Boxes {
		_, err := k8sbox.GetBoxService().UninstallBox(&savedBox, environment.Id)
		if err != nil {
			fmt.Println(" FAIL :(")
			fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
			os.Exit(1)
		}
	}
	fmt.Println(" OK")
}

func validateEnvironmentStep(environment structs.Environment) {
	fmt.Print("Validating environment...")
	err := k8sbox.GetEnvironmentService().ValidateEnvironment(&environment)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}

func validateBoxesStep(environment *structs.Environment, runDirectory string) {
	fmt.Print("Validating boxes...")
	err := k8sbox.GetBoxService().ValidateBoxes(environment.Boxes, runDirectory)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}

	for i, _ := range environment.Boxes {
		err = k8sbox.GetBoxService().FillEmptyFields(&environment.Boxes[i], environment.Namespace)
		if err != nil {
			fmt.Println(" FAIL :(")
			fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
			os.Exit(1)
		}
	}
	fmt.Println(" OK")
}

func createTempDeployDirectoryStep(environment *structs.Environment, runDirectory string, isSaved bool) {
	fmt.Print("Moving files to a temporary directory...")
	var err error
	environment.TempDirectory, err = k8sbox.GetEnvironmentService().CreateTempDeployDirectory(environment, runDirectory, isSaved)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}

func deployEnvironmentStep(environment *structs.Environment, isSaved bool) {
	fmt.Print("Deploying...")
	err := k8sbox.GetEnvironmentService().DeployEnvironment(environment, isSaved)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "Reasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}
