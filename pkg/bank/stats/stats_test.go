package stats
import (
	"github.com/iamgafurov/alif_go_academy/pkg/bank/types"
	"fmt"
)
func statsTest(){
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
fmt.Println(Avg(payments))
//Output: 1232
}
//c