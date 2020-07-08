package device

//Type is the type of device
type Type struct {
	Name    string
	Count   int
	Remain  int
	Address int
}

//Device is the infomation of each device
type Device struct {
	Number  int
	Status  int
	Remain  int
	Jobname string
	Inumber int
}
