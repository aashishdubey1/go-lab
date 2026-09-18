package encapsulationcomposition

import "fmt"

// Multiple Embedding & Name Collision
// Define GPS{} with method Locate() string and Radio{} with method Locate() string (different meaning).
// Embed both into Ship.
// Show the compile error you get calling ship.Locate() directly,
// then fix it by qualifying with ship.GPS.Locate().

type GPS struct {
	Location string
}

func (gps *GPS) Locate() string {
	return fmt.Sprintf("Gps location is %s", gps.Location)
}

type Radio struct {
	Location string
}

func (r *Radio) Locate() string {
	return fmt.Sprintf("Radio Location is %s", r.Location)
}

type Ship struct {
	GPS
	Radio
}

func RunEx2() {
	gps := GPS{Location: "Patna"}
	radio := Radio{Location: "patna"}
	ship := Ship{GPS: gps, Radio: radio}

	// -> this will give error
	// fmt.Println(ship.Locate())

	fmt.Println(ship.GPS.Locate())
	fmt.Println(ship.Radio.Locate())
}
