package card
import (
	"github.com/iamgafurov/alif_go_academy/pkg/bank/types"
)
func Deposit(card *types.Card, amount types.Money) types.Card{
	if(amount <= 50_000_00){
		if ((*card).Active){
			if(amount > 0){
			(*card).Balance = (*card).Balance + amount
			}
		}
	}
	return *card
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int)types.Card{
	bonus:= 0
	if((*card).Balance >= 0 ){
		if ((*card).Active){
			bonus = int((*card).MinBalance) * percent * daysInMonth / daysInYear /100
			if(bonus <= 5_000_00){
				(*card).Balance = ((*card).Balance) + types.Money(bonus)
			}
			
		}
	}
	return *card
}

func Total(cards []types.Card) types.Money{
	sum := types.Money(0)
	for _, item := range cards{
		sum = sum + item.Balance
		
	}
	return sum
 }
 func PaymentSources(cards []types.Card) []types.PaymentSource {
	paymentSources:= []types.PaymentSource{}
	for _,card := range cards{
		if(card.Active && card.Balance >0){
			paymentSources = append(paymentSources,
				types.PaymentSource{
					Type:"card",
					Number: string(card.PAN),
					Balance: card.Balance})
		}
	}
	return paymentSources
   }
