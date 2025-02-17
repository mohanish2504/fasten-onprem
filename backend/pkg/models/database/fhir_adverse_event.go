// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fasten-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirAdverseEvent struct {
	models.ResourceBase
	// actual | potential
	// https://hl7.org/fhir/r4/search.html#token
	Actuality datatypes.JSON `gorm:"column:actuality;type:text;serializer:json" json:"actuality,omitempty"`
	// product-problem | product-quality | product-use-error | wrong-dose | incorrect-prescribing-information | wrong-technique | wrong-route-of-administration | wrong-rate | wrong-duration | wrong-time | expired-drug | medical-device-use-error | problem-different-manufacturer | unsafe-physical-environment
	// https://hl7.org/fhir/r4/search.html#token
	Category datatypes.JSON `gorm:"column:category;type:text;serializer:json" json:"category,omitempty"`
	// When the event occurred
	// https://hl7.org/fhir/r4/search.html#date
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
	// Type of the event itself in relation to the subject
	// https://hl7.org/fhir/r4/search.html#token
	Event datatypes.JSON `gorm:"column:event;type:text;serializer:json" json:"event,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated *time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// Location where adverse event occurred
	// https://hl7.org/fhir/r4/search.html#reference
	Location datatypes.JSON `gorm:"column:location;type:text;serializer:json" json:"location,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Who recorded the adverse event
	// https://hl7.org/fhir/r4/search.html#reference
	Recorder datatypes.JSON `gorm:"column:recorder;type:text;serializer:json" json:"recorder,omitempty"`
	// Effect on the subject due to this event
	// https://hl7.org/fhir/r4/search.html#reference
	Resultingcondition datatypes.JSON `gorm:"column:resultingcondition;type:text;serializer:json" json:"resultingcondition,omitempty"`
	// Seriousness of the event
	// https://hl7.org/fhir/r4/search.html#token
	Seriousness datatypes.JSON `gorm:"column:seriousness;type:text;serializer:json" json:"seriousness,omitempty"`
	// mild | moderate | severe
	// https://hl7.org/fhir/r4/search.html#token
	Severity datatypes.JSON `gorm:"column:severity;type:text;serializer:json" json:"severity,omitempty"`
	// AdverseEvent.study
	// https://hl7.org/fhir/r4/search.html#reference
	Study datatypes.JSON `gorm:"column:study;type:text;serializer:json" json:"study,omitempty"`
	// Subject impacted by event
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Refers to the specific entity that caused the adverse event
	// https://hl7.org/fhir/r4/search.html#reference
	Substance datatypes.JSON `gorm:"column:substance;type:text;serializer:json" json:"substance,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirAdverseEvent) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"actuality":            "token",
		"category":             "token",
		"date":                 "date",
		"event":                "token",
		"id":                   "keyword",
		"language":             "token",
		"lastUpdated":          "date",
		"location":             "reference",
		"profile":              "reference",
		"recorder":             "reference",
		"resultingcondition":   "reference",
		"seriousness":          "token",
		"severity":             "token",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"study":                "reference",
		"subject":              "reference",
		"substance":            "reference",
		"tag":                  "token",
		"text":                 "string",
		"type":                 "special",
	}
	return searchParameters
}
func (s *FhirAdverseEvent) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
	s.ResourceRaw = datatypes.JSON(resourceRaw)
	// unmarshal the raw resource (bytes) into a map
	var resourceRawMap map[string]interface{}
	err := json.Unmarshal(resourceRaw, &resourceRawMap)
	if err != nil {
		return err
	}
	if len(fhirPathJs) == 0 {
		return fmt.Errorf("fhirPathJs script is empty")
	}
	vm := goja.New()
	// setup the global window object
	vm.Set("window", vm.NewObject())
	// set the global FHIR Resource object
	vm.Set("fhirResource", resourceRawMap)
	// compile the fhirpath library
	fhirPathJsProgram, err := goja.Compile("fhirpath.min.js", fhirPathJs, true)
	if err != nil {
		return err
	}
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Actuality
	actualityResult, err := vm.RunString(` 
							ActualityResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.actuality')
							ActualityProcessed = ActualityResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(ActualityProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ActualityProcessed)
							}
						 `)
	if err == nil && actualityResult.String() != "undefined" {
		s.Actuality = []byte(actualityResult.String())
	}
	// extracting Category
	categoryResult, err := vm.RunString(` 
							CategoryResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.category')
							CategoryProcessed = CategoryResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(CategoryProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(CategoryProcessed)
							}
						 `)
	if err == nil && categoryResult.String() != "undefined" {
		s.Category = []byte(categoryResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'AdverseEvent.date')[0]")
	if err == nil && dateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, dateResult.String())
		if err == nil {
			s.Date = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", dateResult.String())
			if err == nil {
				s.Date = &d
			}
		}
	}
	// extracting Event
	eventResult, err := vm.RunString(` 
							EventResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.event')
							EventProcessed = EventResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(EventProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(EventProcessed)
							}
						 `)
	if err == nil && eventResult.String() != "undefined" {
		s.Event = []byte(eventResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString(` 
							LanguageResult = window.fhirpath.evaluate(fhirResource, 'language')
							LanguageProcessed = LanguageResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(LanguageProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LanguageProcessed)
							}
						 `)
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", lastUpdatedResult.String())
			if err == nil {
				s.LastUpdated = &d
			}
		}
	}
	// extracting Location
	locationResult, err := vm.RunString(` 
							LocationResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.location')
						
							if(LocationResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LocationResult)
							}
						 `)
	if err == nil && locationResult.String() != "undefined" {
		s.Location = []byte(locationResult.String())
	}
	// extracting Profile
	profileResult, err := vm.RunString(` 
							ProfileResult = window.fhirpath.evaluate(fhirResource, 'meta.profile')
						
							if(ProfileResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ProfileResult)
							}
						 `)
	if err == nil && profileResult.String() != "undefined" {
		s.Profile = []byte(profileResult.String())
	}
	// extracting Recorder
	recorderResult, err := vm.RunString(` 
							RecorderResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.recorder')
						
							if(RecorderResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(RecorderResult)
							}
						 `)
	if err == nil && recorderResult.String() != "undefined" {
		s.Recorder = []byte(recorderResult.String())
	}
	// extracting Resultingcondition
	resultingconditionResult, err := vm.RunString(` 
							ResultingconditionResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.resultingCondition')
						
							if(ResultingconditionResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ResultingconditionResult)
							}
						 `)
	if err == nil && resultingconditionResult.String() != "undefined" {
		s.Resultingcondition = []byte(resultingconditionResult.String())
	}
	// extracting Seriousness
	seriousnessResult, err := vm.RunString(` 
							SeriousnessResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.seriousness')
							SeriousnessProcessed = SeriousnessResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(SeriousnessProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SeriousnessProcessed)
							}
						 `)
	if err == nil && seriousnessResult.String() != "undefined" {
		s.Seriousness = []byte(seriousnessResult.String())
	}
	// extracting Severity
	severityResult, err := vm.RunString(` 
							SeverityResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.severity')
							SeverityProcessed = SeverityResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(SeverityProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SeverityProcessed)
							}
						 `)
	if err == nil && severityResult.String() != "undefined" {
		s.Severity = []byte(severityResult.String())
	}
	// extracting Study
	studyResult, err := vm.RunString(` 
							StudyResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.study')
						
							if(StudyResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(StudyResult)
							}
						 `)
	if err == nil && studyResult.String() != "undefined" {
		s.Study = []byte(studyResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString(` 
							SubjectResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.subject')
						
							if(SubjectResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SubjectResult)
							}
						 `)
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	// extracting Substance
	substanceResult, err := vm.RunString(` 
							SubstanceResult = window.fhirpath.evaluate(fhirResource, 'AdverseEvent.suspectEntity.instance')
						
							if(SubstanceResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SubstanceResult)
							}
						 `)
	if err == nil && substanceResult.String() != "undefined" {
		s.Substance = []byte(substanceResult.String())
	}
	// extracting Tag
	tagResult, err := vm.RunString(` 
							TagResult = window.fhirpath.evaluate(fhirResource, 'meta.tag')
							TagProcessed = TagResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(TagProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(TagProcessed)
							}
						 `)
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirAdverseEvent) TableName() string {
	return "fhir_adverse_event"
}
