package main

import "encoding/json"

// API for latest vaccination statistics
const VaccinationSummarySource = "https://keralastats.coronasafe.live/vaccination_summary.json"

// Population of Kerala in 2021 as projected by the report of
// National Commission on Population
const KeralaPopulation = 3_54_89_000

// VaccineStat represents vaccination stats
type VaccineStat struct {
	FirstDose  int `json:"tot_person_vaccinations"`
	SecondDose int `json:"second_dose"`
}

// VaccineSummary represents present total and additions withint a day
// in VaccineStat
type VaccineSummary struct {
	Summary VaccineStat
	Delta   VaccineStat
}

// Unmarshal unmarshals jsonData to v
func (v *VaccineSummary) Unmarshal(jsonData []byte) error {
	return json.Unmarshal(jsonData, v)
}
