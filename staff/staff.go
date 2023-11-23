package staff

var OverpaidLimit = 75000
var underPaidLimit = 20000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee // all staff -весь персонал
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverpaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}
func (e *Office) Underpaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary < underPaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}

//func (e *Office) notVisible() {
//	log.Println("Hello, world")
