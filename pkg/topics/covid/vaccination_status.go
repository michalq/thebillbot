package covid

type VaccinationStatus struct {
}

func NewVaccinationStatus() *VaccinationStatus {
	return &VaccinationStatus{}
}

func (s *VaccinationStatus) Answer(message string) []string {
	return make([]string, 0)
}
