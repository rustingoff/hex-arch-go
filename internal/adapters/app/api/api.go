package api

import "github.com/rustingoff/hex-arch-go/internal/ports"

type Adapter struct {
	db    ports.DBPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{
		db:    db,
		arith: arith,
	}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	result, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(result, "addition")
	if err != nil {
		return 0, err
	}

	return result, nil
}
func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	return apia.arith.Subtraction(a, b)
}
func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	return apia.arith.Multiplication(a, b)
}
func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	return apia.arith.Division(a, b)
}
