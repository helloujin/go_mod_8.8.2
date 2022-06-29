package auto

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		if u.T == CM {
			value /= 2.54
		} else {
			value *= 2.54
		}
	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type dimensions struct {
	unitType UnitType
	length   float64
	width    float64
	height   float64
}

type auto struct {
	dimensions  dimensions
	brand       string
	model       string
	maxSpeed    int
	enginePower int
}

func NewAuto(unitType UnitType, brand, model string, length, width, height float64, maxSpeed, enginePower int) *auto {
	return &auto{
		dimensions{
			unitType,
			length,
			width,
			height,
		},
		brand, model, maxSpeed, enginePower,
	}
}

func (a dimensions) Length() Unit {
	return Unit{
		Value: a.length,
		T:     a.unitType,
	}
}

func (a dimensions) Width() Unit {
	return Unit{
		Value: a.width,
		T:     a.unitType,
	}
}

func (a dimensions) Height() Unit {
	return Unit{
		Value: a.height,
		T:     a.unitType,
	}
}

func (a auto) Brand() string {
	return a.brand
}

func (a auto) Model() string {
	return a.model
}

func (a auto) Dimensions() Dimensions {
	return a.dimensions
}

func (a auto) MaxSpeed() int {
	return a.maxSpeed
}

func (a auto) EnginePower() int {
	return a.enginePower
}
