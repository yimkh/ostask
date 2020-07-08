package module

//Block is the piece of memory
type Block struct {
	Name  string
	Start int
	End   int
	//0 is free
	State int
}
