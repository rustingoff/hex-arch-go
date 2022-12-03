package db

type Adapter struct {
	db map[string]interface{}
}

func NewAdapter() *Adapter {
	db := make(map[string]interface{})
	return &Adapter{
		db: db,
	}
}

func (dba Adapter) CloseDBConnection() {}
func (dba Adapter) AddToHistory(answer int32, operation string) error {
	return nil
}
