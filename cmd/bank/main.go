package main
import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)
func main(){
	//result := card.AddBonus(&types.Card{Balance: 0, Active: true, MinBalance:10_000_00}, 3,30,365)
	//fmt.Println(result.Balance)
	cards := []types.Card{
	types.Card{
		PAN: "12312312312312312",
		Active: true,
		Balance: 2000,
	},
	types.Card{
		PAN: "66345345few5435345435",
		Active: true,
		Balance: 2000,
	},
	types.Card{
		PAN: "66345345few5435345435",
		Active: true,
		Balance: 2000,
	},
	types.Card{
		PAN: "66345345few5435345435",
		Active: true,
		Balance: 2000,
	},
	}
fmt.Println(card.PaymentSources(cards))
}
