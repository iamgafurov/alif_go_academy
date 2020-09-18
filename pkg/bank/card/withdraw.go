package card

import(
	"github.com/iamgafurov/alif_go_academy/pkg/bank/types"
)

func Withdraw( card types.Card, amount types.Money) types.Card{
	const limit = 20_000_00
	if (card.Balance > 0){
		if(card.Balance >= amount){
			if(card.Active == true){
				if(amount <= limit){
					card.Balance = card.Balance - amount
					return card
				}else{
					return card
				}
			}else{
				return card
			}
		}else{
			return card
		}
	}else{
		return card
	}
}