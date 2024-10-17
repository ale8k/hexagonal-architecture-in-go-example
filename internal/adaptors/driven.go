package adaptors

import "fmt"

type dbAdaptor struct{}

func NewDbAdaptor() *dbAdaptor {
	return &dbAdaptor{}
}

func (dba *dbAdaptor) Save(data any) {
	fmt.Println("Saved data: ", data)
}
