package thebookstore

type Customer struct {
	ID          string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber int32
	Adresses    []Address
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	ZipCode      int16
}
