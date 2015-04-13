package unexportedreturn

type privateStruct struct{}

// Get returns the unexported struct
func Get() privateStruct {
	return privateStruct{}
}
