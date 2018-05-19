package main

//Profile is the main holder for stuff
type Profile struct {
	ProfileID     string `json:"profileid"`
	Name          string `json:"name"`
	Prescriptions []Prescription
}

//Prescription is the holder of all the info
type Prescription struct {
	PrescriptionID string  `json:"prescriptionid"`
	ProfileID      string  `json:"profileid"`
	MedicineName   string  `json:"medicinename"`
	MedicineID     string  `json:"medicineid"`
	Schedule       string  `json:"schedule"`
	Dosage         string  `json:"dosage"`
	MedicineAmount float32 `json:"medicineamount"`
	Take           string  `json:"take"`
}

//Medication is the holder for stuff
type Medication struct {
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
}

//PrescriptionStandardSchedule is the thing
type PrescriptionStandardSchedule struct {
	PrescriptionID string   `json:"prescriptionid"`
	Interval       string   `json:"standardinterval"`
	ScheduledTimes []string `json:"time"`
}

//PrescriptionFactoredSchedule is the thing
/* type PrescriptionFactoredSchedule struct {
	PrescriptionID string   `json:"prescriptionid"`
	ScheduledTimes []string `json:"time"`
} */

//PrescriptionTaken is the holder
type PrescriptionTaken struct {
	PrescriptionID string `json:"prescriptionid"`
	TakenTime      string `json:"takentime"`
}

//Prescriber reprents a prescriber
type Prescriber struct {
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
}

//Pharmacy is the holder for pharmacies
type Pharmacy struct {
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
}

//PrescriptionFactoredSchedule is the thing
type PrescriptionFactoredSchedule struct {
	PrescriptionID string                             `json:"prescriptionid"`
	ScheduleType   string                             `json:"scheduletype"`
	MonthBasis     []PrescriptionScheduleByDayofMonth `json:"monthbasis"`
	DayofWeekBasis []PrescriptionScheduleByDayofWeek  `json:"dayofweekbasis"`
}

/* //PrescriptionScheduleByMonth is the thing
type PrescriptionScheduleByMonth struct {
	DaysOfMonth []PrescriptionScheduleByDayofMonth
} */

//PrescriptionScheduleByDayofMonth is the thing
type PrescriptionScheduleByDayofMonth struct {
	DayOfMonth     int      `json:"dayofmonth"`
	ScheduledTimes []string `json:"time"`
}

/* //PrescriptionScheduleByWeek is the thing
type PrescriptionScheduleByWeek struct {
	Days []PrescriptionScheduleByDayofWeek
}
*/
//PrescriptionScheduleByDayofWeek is the thing
type PrescriptionScheduleByDayofWeek struct {
	DayOfWeek      string   `json:"dayofweek"`
	ScheduledTimes []string `json:"time"`
}
