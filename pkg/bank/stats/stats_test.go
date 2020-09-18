package stats
import (
	"github.com/iamgafurov/alif_go_academy/pkg/bank/types"
	"fmt"
)
<<<<<<< HEAD
func AvgTest(){
=======
func statsTest(){
>>>>>>> master
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
<<<<<<< HEAD

func TotalInCategoryTest(){
	payments := []types.Payment{
	types.Payment{
		ID: 1,
		Amount: 1232,
		Category:"b",
	},
	types.Payment{
		ID: 2,
		Amount: 1232,
	},
	types.Payment{
		ID: 3,
		Category:"b",
	},
	types.Payment{
		ID: 4,
		Amount: 1232,
	},
	}
fmt.Println(TotalInCategory(payments,"b"))
//Output: 2464
}
=======
//c
>>>>>>> master
