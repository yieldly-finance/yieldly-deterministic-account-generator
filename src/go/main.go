package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	utils "../go/utils"
)


func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("\n****************************************************")
    fmt.Println("Welcome to Yieldlys Deterministic Account Generator!")
    fmt.Println("****************************************************\n")
    fmt.Println("Please choose carefully:")
    fmt.Println("1) Mnemonic-to-derivation key")
    fmt.Println("2) Derivation key-to-Mnemonic key")
    fmt.Println("3) Generate Accounts")
    fmt.Println("4) Get Account Mnemonic")
	fmt.Print("\nChoose your option: ")
	x, _ := reader.ReadString('\n')

	// TODO - Slice the strings up to the new line
	switch strings.TrimSuffix(x, "\n") {
	case "1":
		fmt.Print("\nEnter your Mnemonic: ")
		key, _ := reader.ReadString('\n')
		fmt.Println("\n" + utils.MasterToDerivation(strings.TrimSuffix(key, "\n")) + "\n")
	case "2":
		fmt.Print("\nEnter your derivation key: ")
		key, _ := reader.ReadString('\n')
		fmt.Println("\n" + utils.DerivationToMaster(strings.TrimSuffix(key, "\n")) + "\n")
	case "3":
		fmt.Print("\nEnter your seed (mnemonic or derivation key): ")
		seed, _ := reader.ReadString('\n')
		fmt.Print("\nEnter the index of the account list (max 2^63-1): ")
		index, _ := reader.ReadString('\n')
		newIndex, err:= strconv.Atoi(strings.TrimSuffix(index, "\n"))
		if err != nil {
			fmt.Printf("Cannot be convered to an Int")
		}
		
		fmt.Print("\nEnter the amount of addresses you would like after the index (max 2^63-1 - index): ")
		amount, _ := reader.ReadString('\n')
		newAmount, err:= strconv.Atoi(strings.TrimSuffix(amount, "\n"))
		if err != nil {
			fmt.Printf("Cannot be convered to an Int")
		}

		fmt.Println("\nYour accounts: ")

		fmt.Println(strings.Join(utils.GenerateAccountList(strings.TrimSuffix(seed, "\n"), newIndex, newAmount)[:], "\n\n") )
	case "4":
		fmt.Print("\nEnter your seed (mnemonic or derivation key): ")
		seed, _ := reader.ReadString('\n')
		fmt.Print("\nEnter the index of the account list (max 2^63-1): ")
		index, _ := reader.ReadString('\n')
		newIndex, err:= strconv.Atoi(strings.TrimSuffix(index, "\n"))
		if err != nil {
			fmt.Printf("Cannot be convered to an Int")
		}
		
		fmt.Println("\nYour accounts mnemonic: ")

		fmt.Println(strings.Join(utils.GenerateAccountMnemonic(strings.TrimSuffix(seed, "\n"), newIndex)[:], "\n\n") )
	default:
		fmt.Println("Invalid option selected")
	}
}