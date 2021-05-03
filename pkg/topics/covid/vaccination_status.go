package covid

import (
	"context"
	"fmt"
	"regexp"

	"github.com/michalq/thebillbot/pkg/szczepimysie"
)

type VaccinationStatus struct {
	szczepimysieClient *szczepimysie.Client
}

func NewVaccinationStatus(szczepimysieClient *szczepimysie.Client) *VaccinationStatus {
	return &VaccinationStatus{szczepimysieClient}
}

func (v *VaccinationStatus) Answer(message string) []string {

	pattern := regexp.MustCompile(`vaccination`)
	if !pattern.MatchString(message) {
		return []string{}
	}
	status, err := v.szczepimysieClient.LatestStatus(context.TODO())
	if err != nil {
		return []string{err.Error()}
	}
	return []string{fmt.Sprintf("Dawka 1 suma %d", status.Dawka1Suma)}
}

func (v *VaccinationStatus) Promote() string {
	return `Example "vaccination"`
}
