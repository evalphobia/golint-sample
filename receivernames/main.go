package receivernames

type myStruct struct {
	field1 string
}

func (_ *myStruct) underScore() {}

func (me *myStruct) me()     {}
func (this *myStruct) this() {}
func (self *myStruct) self() {}

func (m *myStruct) method1() {}
func (s *myStruct) method2() {}
