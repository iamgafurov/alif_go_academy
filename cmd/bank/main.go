package main
import (
	"github.com/iamgafurov/alif_go_academy/pkg/bank/stats"
	"github.com/iamgafurov/alif_go_academy/pkg/bank/types"
	"fmt"
)
func main(){
	//result := card.AddBonus(&types.Card{Balance: 0, Active: true, MinBalance:10_000_00}, 3,30,365)
	//fmt.Println(result.Balance)
	payments := []types.Payment{
	types.Payment{
		ID: 1,
		Amount: 1232,
	},
	types.Payment{
		ID: 2,
		Amount: 1232,
	},
	types.Payment{
		ID: 3,
		Amount: 1232,
	},
	types.Payment{
		ID: 4,
		Amount: 1232,
	},
	}
fmt.Println(stats.Avg(payments))
}
