package stats
import (
	"bank/pkg/bank/types"
)
func Avg(payments []types.Payment) types.Money{
	sum:= types.Money(0)
	for _,payment := range payments{
		sum= sum + payment.Amount
	}
	return sum / types.Money(len(payments))
}