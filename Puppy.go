package puppy

import (
	dog "github.com/dhirajkashyap/Dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!"
}

func Barkz() string {
	return " Woof! Woof! Woof!"
}

func BigBark() string {
	//return Dog.WhenGrownUp(Barkz())
	return dog.WhenGrownUp(Barkz())
}

func BabyBark() string {
	return dog.WhenStillBaby()
}
