package covid

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/michalq/thebillbot/internal/messenger"
	"github.com/michalq/thebillbot/pkg/szczepimysie"
)

const PolandPopulation = 37970000

type VaccinationStatus struct {
	szczepimysieClient *szczepimysie.Client
}

func NewVaccinationStatus(szczepimysieClient *szczepimysie.Client) *VaccinationStatus {
	return &VaccinationStatus{szczepimysieClient}
}

func (v *VaccinationStatus) Answer(ctx context.Context, message string) []messenger.Message {

	pattern := regexp.MustCompile(`vaccination`)
	if !pattern.MatchString(strings.ToLower(message)) {
		return []messenger.Message{}
	}
	status, err := v.szczepimysieClient.LatestStatus(ctx)
	if err != nil {
		return []messenger.Message{{Topic: "Error", Content: err.Error()}}
	}
	return []messenger.Message{
		{Content: fmt.Sprintf("💉 First dose of vaccination get %d people in Poland. It is %.2f%% of whole population.", status.Dawka1Suma, float64(status.Dawka1Suma)/float64(PolandPopulation)*100.0)},
		{Content: fmt.Sprintf("💉 Second dose of vaccination get %d people in Poland. It is %.2f%% of whole population.", status.Dawka2Suma, float64(status.Dawka2Suma)/float64(PolandPopulation)*100.0)},
	}
}

func (v *VaccinationStatus) Promote() string {
	return `Example "vaccination"`
}

func (v *VaccinationStatus) Name() string {
	return "Vaccination status"
}
