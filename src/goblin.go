package projet

import (
	"fmt"
)


func Initgoblin()  {
	monster := &monster{
		name: "Gobelin",
		pv:   50,
		pvmax: 50,
		experience: 10,
	}
	return monster
}