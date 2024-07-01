package patients

type Patient struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	InsuranceId string `json:"insuranceId"`
	InsurerId   string `json:"insurerId"`
}
