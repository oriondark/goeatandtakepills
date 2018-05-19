package main

//PrescriptionProfile is the main holder for stuff
type PrescriptionProfile struct {
	ProfileID string                `json:"profileid"`
	Name      string                `json:"name"`
	Summary   []PrescriptionSummary `json:"prescriptions"`
}

//PrescriptionSummary sucks
type PrescriptionSummary struct {
	PrescriptionID string `json:"prescriptionid"`
	Medication     struct {
		MedicationID                  string `json:"medicationid"`
		MedicationName                string `json:"medicationame"`
		MedicationDosage              string `json:"medicationdosage"`
		MedicationDosageType          string `json:"medicationdosagetype"`
		MedicationMustTakeWithFood    bool   `json:"medicationmusttakewithfood"`
		MedicationMustTakeWithoutFood bool   `json:"medicationmusttakewithoutfood"`
		MedicationMeasuredAmount      string `json:"medicationmeasuredamount"`
		MedicationNDC                 string `json:"medicationndc"`
		MedicationClass               string `json:"medicationclass"`
		MedicationImage               string `json:"medicationimage"`
		Overview                      string `json:"medicationoverview"`
		SideEffects                   string `json:"medicationsideeffects"`
		FAQ                           string `json:"medicationfaq"`
		MedicationDosageFormName      string `json:"medicationdosageformname"`
		MedicationIngredientUnit      string `json:"medicationingredientunit"`
		MedicationStrength            string `json:"medicationstrength"`
		MedicationSubstances          string `json:"medicationsubstances"`
	} `json:"medication"`
	Schedule struct {
		PrescriptionID string                             `json:"prescriptionid"`
		ScheduleType   string                             `json:"scheduletype"`
		MonthBasis     []PrescriptionScheduleByDayofMonth `json:"monthbasis"`
		DayofWeekBasis []PrescriptionScheduleByDayofWeek  `json:"dayofweekbasis"`
	} `json:"schedule"`
	/* Schedule struct {
		ScheduledTimes []string `json:"time"`
	} `json:"schedule"` */
	TakenHistory struct {
		TakenTime string `json:"takentime"`
	} `json:"taken"`
	Prescriber struct {
		PrescriberID        string  `json:"prescriberid"`
		PrescriberName      string  `json:"prescribername"`
		PrescriberStreet    string  `json:"prescriberstreet"`
		PrescriberCity      string  `json:"prescribercity"`
		PrescriberState     string  `json:"prescriberstate"`
		PrescriberPostCode  string  `json:"prescriberpostcode"`
		PrescriberImage     string  `json:"prescriberimage"`
		PrescriberPhone     string  `json:"prescriberphone"`
		PrescriberLatitude  float64 `json:"prescriberlatitude"`
		PrescriberLongitude float64 `json:"prescriberlongitude"`
	} `json:"prescriber"`
	Pharmacy struct {
		PharmacyID        string  `json:"pharmacyid"`
		PharmacyName      string  `json:"pharmacyname"`
		PharmacyStreet    string  `json:"pharmacystreet"`
		PharmacyCity      string  `json:"pharmacycity"`
		PharmacyState     string  `json:"pharmacystate"`
		PharmacyPostCode  string  `json:"pharmacypostcode"`
		PharmacyImage     string  `json:"pharmacyimage"`
		PharmacyPhone     string  `json:"pharmacyphone"`
		PharmacyLatitude  float64 `json:"pharmacylatitude"`
		PharmacyLongitude float64 `json:"pharmacylongitude"`
	} `json:"pharmacy"`
}
