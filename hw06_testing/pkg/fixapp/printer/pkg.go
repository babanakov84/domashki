package printer

import (
	"errors"
	"fmt"

	"github.com/babanakov84/domashki/hw06_testing/pkg/fixapp/types"
)

func PrintStaff(staff []types.Employee) error {
	var err error
	for i := 0; i < len(staff) && err == nil; i++ {
		if staff[i].UserID == 0 || staff[i].Name == "" || staff[i].DepartmentID == 0 || staff[i].Age == 0 {
			err = errors.New("bad JSON data")
		} else {
			fmt.Println(staff[i])
		}
	}
	return err
}
