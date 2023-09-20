package puppy

import (
	"github.com/dhirajkashyap/Dog"
)


func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!"
}

func Barkz() string {
	return "Woof! Woof! Woof!"
}


func BigBark() string {
	return dog.WhenGrownUp(Barkz()))
}