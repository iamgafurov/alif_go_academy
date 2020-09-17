package card
import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleDeposit_negativeBalance() {
	result := Deposit(&types.Card{Balance: -20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: -1000000
	}
func ExampleDeposit_positive() {
	result := Deposit(&types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 3000000
	}
func Example3() {
	result := Deposit(&types.Card{Balance: 20_000_00, Active: true}, 55_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
	}