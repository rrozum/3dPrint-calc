package domain

type Detail struct {
	Id                            int
	Name                          string
	MarginRatio                   float64
	PlasticConsumption            float64
	PrintTime                     string
	DataPreparationByOperatorTime string
	PostProcessingByOperatorTime  string
	ProjectId                     int
}
