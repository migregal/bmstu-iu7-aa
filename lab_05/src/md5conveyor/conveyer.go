package md5conveyor

type Conveyor struct {
	numDigesters int
}

func NewConveyor(numDigesters int) Conveyor {
	return Conveyor{numDigesters}
}
