package domain

type Project struct {
	Id                    int
	Name                  string
	Price                 float64
	PlasticType           string
	PlasticPriceByCoil    float64
	PlasticPriceByKg      float64
	OperatorSalaryByMonth float64
	OperatorSalaryByHour  float64
	ElectricityPrice      float64
	RentPrice             float64
	MinimalPrice          float64
	Discount              float64
}
