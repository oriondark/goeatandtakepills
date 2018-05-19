package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	//CreateRandom()
	//fmt.Printf("%+v\n", randommedication())
	//med := randommedication()
	//GetRandomPharmacy()
	var med [5]*PrescriptionProfile
	med[0] = randomprescriptionprofile("1", "Donald Duck")
	med[1] = randomprescriptionprofile("2", "DaffyDuck")
	med[2] = randomprescriptionprofile("3", "Mickey Mouse")
	med[3] = randomprescriptionprofile("4", "Scrooge McDuck")
	med[4] = randomprescriptionprofile("5", "Daisy Duke")
	b, err := json.Marshal(med)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	/* med := randomprescriptionprofile("daffy")
	b, err := json.Marshal(med)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b)) */

	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, string(b))

}

func randommedication() Medication {
	var m Medication
	var item NDCItem
	item = GetRandomMedication()
	m.MedicationClass = item.Classes
	m.MedicationDosage = strconv.Itoa(randomdata.Number(1, 4))
	m.MedicationDosageType = item.DosageFormName
	m.MedicationID = item.ProductID
	m.MedicationMeasuredAmount = randomdata.SillyName()
	m.MedicationMustTakeWithFood = randomdata.Boolean()
	m.MedicationMustTakeWithoutFood = randomdata.Boolean()
	m.MedicationName = item.Proprietary
	m.MedicationNDC = item.NDC
	m.Overview = randomdata.Paragraph()
	m.SideEffects = randomdata.Paragraph()
	m.FAQ = randomdata.Paragraph()
	m.MedicationDosageFormName = item.DosageFormName
	m.MedicationIngredientUnit = item.IngredientUnit
	m.MedicationStrength = item.Strength
	m.MedicationSubstances = item.Substances

	return m
}

func randompharmacy() Pharmacy {
	var myrand PharmaItem
	myrand = GetRandomPharmacy()
	var p Pharmacy
	/* p.PharmacyID = randomdata.IpV4Address()
	p.PharmacyCity = randomdata.City()
	p.PharmacyLatitude = randomdata.Decimal(180)
	p.PharmacyLongitude = randomdata.Decimal(180)
	p.PharmacyName = randomdata.SillyName()
	p.PharmacyPhone = randomdata.PhoneNumber()
	p.PharmacyPostCode = randomdata.PostalCode("US")
	p.PharmacyState = randomdata.State(0)
	p.PharmacyStreet = strconv.Itoa(randomdata.Number(10, 800)) + " " + randomdata.Noun() */

	p.PharmacyID = myrand.PharmacyID
	p.PharmacyCity = myrand.City
	p.PharmacyLatitude = myrand.Latitude
	p.PharmacyLongitude = myrand.Longitude
	p.PharmacyName = myrand.Name
	p.PharmacyPhone = myrand.Phone
	p.PharmacyPostCode = strconv.Itoa(myrand.PostCode)
	p.PharmacyState = myrand.State
	p.PharmacyStreet = myrand.Street

	return p

}

func randompresciber() Prescriber {
	var ps Prescriber
	ps.PrescriberID = randomdata.IpV4Address()
	ps.PrescriberCity = randomdata.City()
	ps.PrescriberLatitude = randomdata.Decimal(180)
	ps.PrescriberLongitude = randomdata.Decimal(180)
	ps.PrescriberName = randomdata.SillyName()
	ps.PrescriberPhone = randomdata.PhoneNumber()
	ps.PrescriberPostCode = randomdata.PostalCode("US")
	ps.PrescriberState = randomdata.State(0)
	ps.PrescriberStreet = strconv.Itoa(randomdata.Number(10, 800)) + " " + randomdata.Noun()
	return ps

}

/* func randomrealschedule() PrescriptionFactoredSchedule {
	var rs PrescriptionFactoredSchedule
	rs.PrescriptionID = randomdata.IpV4Address()
	var t [4]string
	for i := 0; i < 4; i++ {
		t[i] = strconv.Itoa(randomdata.Number(23)) + ":" + strconv.Itoa(randomdata.Number(59))
	}
	rs.ScheduledTimes = t[0:4]
	//rs.ScheduledTimes =
	return rs
} */

func randomtakenhistory() PrescriptionTaken {

	var hist PrescriptionTaken
	hist.PrescriptionID = randomdata.IpV6Address()
	hist.TakenTime = randomdata.FullDate()
	return hist
}

func randomstandardschedule() PrescriptionStandardSchedule {

	var ss PrescriptionStandardSchedule
	ss.PrescriptionID = randomdata.IpV6Address()
	var t [4]string
	for i := 0; i < 4; i++ {
		t[i] = strconv.Itoa(randomdata.Number(23)) + ":" + strconv.Itoa(randomdata.Number(59))
	}
	ss.ScheduledTimes = t[0:4]
	ss.Interval = "daily"
	return ss
}
func randomprescriptionsummary() PrescriptionSummary {
	var pp PrescriptionSummary
	pp.PrescriptionID = randomdata.IpV4Address()
	//pp.ProfileID = randomdata.IpV4Address()
	//pp.Name = randomdata.SillyName()

	m := randommedication()
	pp.Medication.FAQ = m.FAQ
	pp.Medication.MedicationClass = m.MedicationClass
	pp.Medication.MedicationDosage = m.MedicationDosage
	pp.Medication.MedicationID = m.MedicationID
	pp.Medication.MedicationImage = m.MedicationImage
	pp.Medication.MedicationMeasuredAmount = m.MedicationMeasuredAmount
	pp.Medication.MedicationMustTakeWithFood = m.MedicationMustTakeWithFood
	pp.Medication.MedicationMustTakeWithoutFood = m.MedicationMustTakeWithoutFood
	pp.Medication.MedicationName = m.MedicationName
	pp.Medication.MedicationNDC = m.MedicationNDC
	pp.Medication.Overview = m.Overview
	pp.Medication.SideEffects = m.SideEffects
	pp.Medication.MedicationDosageType = m.MedicationDosageType
	pp.Medication.MedicationDosageFormName = m.MedicationDosageFormName
	pp.Medication.MedicationIngredientUnit = m.MedicationIngredientUnit
	pp.Medication.MedicationStrength = m.MedicationStrength
	pp.Medication.MedicationSubstances = m.MedicationSubstances

	p := randompharmacy()
	pp.Pharmacy.PharmacyID = p.PharmacyID
	pp.Pharmacy.PharmacyCity = p.PharmacyCity
	pp.Pharmacy.PharmacyImage = p.PharmacyImage
	pp.Pharmacy.PharmacyLatitude = p.PharmacyLatitude
	pp.Pharmacy.PharmacyLongitude = p.PharmacyLongitude
	pp.Pharmacy.PharmacyName = p.PharmacyName
	pp.Pharmacy.PharmacyPhone = p.PharmacyPhone
	pp.Pharmacy.PharmacyPostCode = p.PharmacyPostCode
	pp.Pharmacy.PharmacyState = p.PharmacyState
	pp.Pharmacy.PharmacyStreet = p.PharmacyStreet

	ps := randompresciber()
	pp.Prescriber.PrescriberID = ps.PrescriberID
	pp.Prescriber.PrescriberImage = ps.PrescriberImage
	pp.Prescriber.PrescriberCity = ps.PrescriberCity
	pp.Prescriber.PrescriberLatitude = ps.PrescriberLatitude
	pp.Prescriber.PrescriberLongitude = ps.PrescriberLongitude
	pp.Prescriber.PrescriberName = ps.PrescriberName
	pp.Prescriber.PrescriberPhone = ps.PrescriberPhone
	pp.Prescriber.PrescriberPostCode = ps.PrescriberPostCode
	pp.Prescriber.PrescriberState = ps.PrescriberState
	pp.Prescriber.PrescriberStreet = ps.PrescriberStreet

	rs := randomrealschedule()
	pp.Schedule.PrescriptionID = pp.PrescriptionID
	pp.Schedule.ScheduleType = rs.ScheduleType
	pp.Schedule.DayofWeekBasis = rs.DayofWeekBasis
	pp.Schedule.MonthBasis = rs.MonthBasis

	hist := randomtakenhistory()
	pp.TakenHistory.TakenTime = hist.TakenTime

	return pp

}

func randomprescriptionprofile(someid string, somename string) *PrescriptionProfile {

	profile := &PrescriptionProfile{}
	profile.ProfileID = someid // randomdata.IpV4Address()
	profile.Name = somename    //randomdata.SillyName()

	var ps [10]PrescriptionSummary
	for i := 0; i < 10; i++ {
		ps[i] = randomprescriptionsummary()
	}
	profile.Summary = ps[0:10]

	return profile

}

func randomrealschedule() PrescriptionFactoredSchedule {
	var rs PrescriptionFactoredSchedule
	rs.PrescriptionID = randomdata.IpV4Address()
	rs.ScheduleType = "DayOfWeek"
	//var ss PrescriptionScheduleByDayofWeek
	//var daysofweek = randdaysofweek()
	var randtimes = randstandtimes()
	//var ary = len(daysofweek)

	var dob [7]PrescriptionScheduleByDayofWeek

	for z := 0; z < 7; z++ {
		dob[z].DayOfWeek = strconv.Itoa(z)
		dob[z].ScheduledTimes = randtimes
	}
	rs.DayofWeekBasis = dob[:]

	return rs
}

func randdaysofweek() []int {
	return []int{1, 2, 3, 4, 5, 6, 7}
}
func randstandtimes() []string {
	values := [][]string{}
	qid := []string{"0900", "1300", "1700", "2100"}
	tid := []string{"0900", "1400", "2100"}
	bid := []string{"0900", "2100"}
	tidpsych := []string{"0800", "1200", "1700"}
	daily := []string{"0900"}
	bedtime := []string{"2100"}
	withmeals := []string{"0800", "1200", "1700"}
	values = append(values, qid)
	values = append(values, tid)
	values = append(values, bid)
	values = append(values, tidpsych)
	values = append(values, daily)
	values = append(values, bedtime)
	values = append(values, withmeals)
	rnd := randomdata.Number(0, len(values))
	return values[rnd]

}
