package main

// API for latest vaccination statistics
const VaccinationSummarySource = "https://keralastats.coronasafe.live/vaccination_summary.json"

// Population of Kerala in 2021 as projected by the report of
// National Commission on Population
const KeralaPopulation = 3_54_89_000

// VaccineStat represents vaccination stats
type VaccineStat struct {
	TotalDose int `json:"tot_person_vaccinations"`
	// FirstDose value is faultily present in "second_dose" json key
	// due to incorrect scraping at API end.
	FirstDose int `json:"second_dose"`
	// SecondDose is calculated after fetching the data
	SecondDose int `json:""`
}

// VaccineSummary represents present total and additions withint a day
// in VaccineStat
type VaccineSummary struct {
	Summary VaccineStat
	Delta   VaccineStat
}

// CalcSecondDose assigns the SecondDose value to fields in VaccineSummary
func (v *VaccineSummary) CalcSecondDose() {
	v.Summary.SecondDose = v.Summary.TotalDose - v.Summary.FirstDose
	v.Delta.SecondDose = v.Delta.TotalDose - v.Delta.FirstDose
}

// VaccinatedPercent returns the population percentage of Kerala
// who received first dose of vaccination
func (v *VaccineSummary) VaccinatedPercent() float64 {
	return float64(v.Summary.FirstDose) / KeralaPopulation * 100
}
