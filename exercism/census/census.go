// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{Name: name, Age: age, Address: address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	return r.Name != "" && len(r.Address) > 0 && r.Address["street"] != ""
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	//panic("Please implement Delete.")
	r.Name = ""
	r.Age = 0
	//str := make(map[string]string)
	r.Address = map[string]string(nil) //make(map[string]string)
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	//panic("Please implement Count.")
	var count = 0
	for _, el := range residents {
		if el.HasRequiredInfo() {
			count++
		}
	}
	return count
}
