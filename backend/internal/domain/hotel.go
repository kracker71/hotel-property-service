package domain

type Hotel struct {
	ID       string
	Name     string
	Address  string
	IsActive bool
	Facility []Facility
}
