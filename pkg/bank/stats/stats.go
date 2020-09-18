package stats
import (
	"github.com/iamgafurov/alif_go_academy/pkg/bank/types"
)
func Avg(payments []types.Payment) types.Money{
	sum:= types.Money(0)
	for _,payment := range payments{
		sum= sum + payment.Amount
	}
	return sum / types.Money(len(payments))
}