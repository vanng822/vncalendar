package vncalendar

import "fmt"

func padd(digits int) string {
	if digits > 9 {
		return fmt.Sprintf("%d", digits)
	}
	return fmt.Sprintf("0%d", digits)
}
