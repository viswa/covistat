package main

// API for latest vaccination statistics
const VaccinationSummarySource = "https://keralastats.coronasafe.live/vaccination_summary.json"

// Population of Kerala in 2021 as projected by the report of
// National Commission on Population
const KeralaPopulation = 3_54_89_000

type VaccineStat struct {
	FirstDose  int `json:"tot_person_vaccinations"`
	SecondDose int `json:"second_dose"`
}

type VaccineSummary struct {
	Summary VaccineStat
	Delta   VaccineStat
}
