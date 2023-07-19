// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Student student
//
// swagger:model Student
type Student struct {

	// created
	// Format: datetime
	Created strfmt.DateTime `json:"created,omitempty"`

	// credentials
	Credentials *Credentials `json:"credentials,omitempty"`

	// dob
	// Pattern: (?:[0-9]{1,2})/([0-9]{1,2})/([0-9]{4})
	Dob *string `json:"dob,omitempty"`

	// ell status
	// Enum: [Y N ]
	EllStatus *string `json:"ell_status,omitempty"`

	// enrollments
	Enrollments []*SchoolEnrollment `json:"enrollments"`

	// ext
	Ext interface{} `json:"ext,omitempty"`

	// gender
	// Enum: [M F X ]
	Gender *string `json:"gender,omitempty"`

	// grade
	// Enum: [InfantToddler Preschool PreKindergarten TransitionalKindergarten Kindergarten 1 2 3 4 5 6 7 8 9 10 11 12 13 PostGraduate Ungraded Other ]
	Grade *string `json:"grade,omitempty"`

	// graduation year
	GraduationYear *string `json:"graduation_year,omitempty"`

	// hispanic ethnicity
	// Enum: [Y N ]
	HispanicEthnicity *string `json:"hispanic_ethnicity,omitempty"`

	// home language
	// Enum: [English Albanian Amharic Arabic Bengali Bosnian Burmese Cantonese Chinese Dutch Farsi French German Hebrew Hindi Hmong Ilocano Japanese Javanese Karen Khmer Korean Laotian Latvian Malay Mandarin Nepali Oromo Polish Portuguese Punjabi Romanian Russian Samoan Serbian Somali Spanish Swahili Tagalog Tamil Telugu Thai Tigrinya Turkish Ukrainian Urdu Vietnamese ]
	HomeLanguage *string `json:"home_language,omitempty"`

	// iep status
	IepStatus *string `json:"iep_status,omitempty"`

	// last modified
	// Format: datetime
	LastModified strfmt.DateTime `json:"last_modified,omitempty"`

	// legacy id
	LegacyID string `json:"legacy_id,omitempty"`

	// location
	Location *Location `json:"location,omitempty"`

	// race
	// Enum: [Caucasian Asian Black or African American American Indian Hawaiian or Other Pacific Islander Two or More Races Unknown ]
	Race *string `json:"race,omitempty"`

	// school
	School string `json:"school,omitempty"`

	// schools
	Schools []string `json:"schools"`

	// sis id
	SisID string `json:"sis_id,omitempty"`

	// state id
	StateID *string `json:"state_id,omitempty"`

	// student number
	StudentNumber *string `json:"student_number,omitempty"`

	// unweighted gpa
	UnweightedGpa *string `json:"unweighted_gpa,omitempty"`

	// weighted gpa
	WeightedGpa *string `json:"weighted_gpa,omitempty"`
}

// Validate validates this student
func (m *Student) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEllStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnrollments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrade(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHispanicEthnicity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomeLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Student) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "datetime", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Student) validateCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Student) validateDob(formats strfmt.Registry) error {
	if swag.IsZero(m.Dob) { // not required
		return nil
	}

	if err := validate.Pattern("dob", "body", *m.Dob, `(?:[0-9]{1,2})/([0-9]{1,2})/([0-9]{4})`); err != nil {
		return err
	}

	return nil
}

var studentTypeEllStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Y","N",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		studentTypeEllStatusPropEnum = append(studentTypeEllStatusPropEnum, v)
	}
}

const (

	// StudentEllStatusY captures enum value "Y"
	StudentEllStatusY string = "Y"

	// StudentEllStatusN captures enum value "N"
	StudentEllStatusN string = "N"

	// StudentEllStatusEmpty captures enum value ""
	StudentEllStatusEmpty string = ""
)

// prop value enum
func (m *Student) validateEllStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, studentTypeEllStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Student) validateEllStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EllStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateEllStatusEnum("ell_status", "body", *m.EllStatus); err != nil {
		return err
	}

	return nil
}

func (m *Student) validateEnrollments(formats strfmt.Registry) error {
	if swag.IsZero(m.Enrollments) { // not required
		return nil
	}

	for i := 0; i < len(m.Enrollments); i++ {
		if swag.IsZero(m.Enrollments[i]) { // not required
			continue
		}

		if m.Enrollments[i] != nil {
			if err := m.Enrollments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("enrollments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var studentTypeGenderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["M","F","X",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		studentTypeGenderPropEnum = append(studentTypeGenderPropEnum, v)
	}
}

const (

	// StudentGenderM captures enum value "M"
	StudentGenderM string = "M"

	// StudentGenderF captures enum value "F"
	StudentGenderF string = "F"

	// StudentGenderX captures enum value "X"
	StudentGenderX string = "X"

	// StudentGenderEmpty captures enum value ""
	StudentGenderEmpty string = ""
)

// prop value enum
func (m *Student) validateGenderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, studentTypeGenderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Student) validateGender(formats strfmt.Registry) error {
	if swag.IsZero(m.Gender) { // not required
		return nil
	}

	// value enum
	if err := m.validateGenderEnum("gender", "body", *m.Gender); err != nil {
		return err
	}

	return nil
}

var studentTypeGradePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InfantToddler","Preschool","PreKindergarten","TransitionalKindergarten","Kindergarten","1","2","3","4","5","6","7","8","9","10","11","12","13","PostGraduate","Ungraded","Other",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		studentTypeGradePropEnum = append(studentTypeGradePropEnum, v)
	}
}

const (

	// StudentGradeInfantToddler captures enum value "InfantToddler"
	StudentGradeInfantToddler string = "InfantToddler"

	// StudentGradePreschool captures enum value "Preschool"
	StudentGradePreschool string = "Preschool"

	// StudentGradePreKindergarten captures enum value "PreKindergarten"
	StudentGradePreKindergarten string = "PreKindergarten"

	// StudentGradeTransitionalKindergarten captures enum value "TransitionalKindergarten"
	StudentGradeTransitionalKindergarten string = "TransitionalKindergarten"

	// StudentGradeKindergarten captures enum value "Kindergarten"
	StudentGradeKindergarten string = "Kindergarten"

	// StudentGradeNr1 captures enum value "1"
	StudentGradeNr1 string = "1"

	// StudentGradeNr2 captures enum value "2"
	StudentGradeNr2 string = "2"

	// StudentGradeNr3 captures enum value "3"
	StudentGradeNr3 string = "3"

	// StudentGradeNr4 captures enum value "4"
	StudentGradeNr4 string = "4"

	// StudentGradeNr5 captures enum value "5"
	StudentGradeNr5 string = "5"

	// StudentGradeNr6 captures enum value "6"
	StudentGradeNr6 string = "6"

	// StudentGradeNr7 captures enum value "7"
	StudentGradeNr7 string = "7"

	// StudentGradeNr8 captures enum value "8"
	StudentGradeNr8 string = "8"

	// StudentGradeNr9 captures enum value "9"
	StudentGradeNr9 string = "9"

	// StudentGradeNr10 captures enum value "10"
	StudentGradeNr10 string = "10"

	// StudentGradeNr11 captures enum value "11"
	StudentGradeNr11 string = "11"

	// StudentGradeNr12 captures enum value "12"
	StudentGradeNr12 string = "12"

	// StudentGradeNr13 captures enum value "13"
	StudentGradeNr13 string = "13"

	// StudentGradePostGraduate captures enum value "PostGraduate"
	StudentGradePostGraduate string = "PostGraduate"

	// StudentGradeUngraded captures enum value "Ungraded"
	StudentGradeUngraded string = "Ungraded"

	// StudentGradeOther captures enum value "Other"
	StudentGradeOther string = "Other"

	// StudentGradeEmpty captures enum value ""
	StudentGradeEmpty string = ""
)

// prop value enum
func (m *Student) validateGradeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, studentTypeGradePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Student) validateGrade(formats strfmt.Registry) error {
	if swag.IsZero(m.Grade) { // not required
		return nil
	}

	// value enum
	if err := m.validateGradeEnum("grade", "body", *m.Grade); err != nil {
		return err
	}

	return nil
}

var studentTypeHispanicEthnicityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Y","N",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		studentTypeHispanicEthnicityPropEnum = append(studentTypeHispanicEthnicityPropEnum, v)
	}
}

const (

	// StudentHispanicEthnicityY captures enum value "Y"
	StudentHispanicEthnicityY string = "Y"

	// StudentHispanicEthnicityN captures enum value "N"
	StudentHispanicEthnicityN string = "N"

	// StudentHispanicEthnicityEmpty captures enum value ""
	StudentHispanicEthnicityEmpty string = ""
)

// prop value enum
func (m *Student) validateHispanicEthnicityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, studentTypeHispanicEthnicityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Student) validateHispanicEthnicity(formats strfmt.Registry) error {
	if swag.IsZero(m.HispanicEthnicity) { // not required
		return nil
	}

	// value enum
	if err := m.validateHispanicEthnicityEnum("hispanic_ethnicity", "body", *m.HispanicEthnicity); err != nil {
		return err
	}

	return nil
}

var studentTypeHomeLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["English","Albanian","Amharic","Arabic","Bengali","Bosnian","Burmese","Cantonese","Chinese","Dutch","Farsi","French","German","Hebrew","Hindi","Hmong","Ilocano","Japanese","Javanese","Karen","Khmer","Korean","Laotian","Latvian","Malay","Mandarin","Nepali","Oromo","Polish","Portuguese","Punjabi","Romanian","Russian","Samoan","Serbian","Somali","Spanish","Swahili","Tagalog","Tamil","Telugu","Thai","Tigrinya","Turkish","Ukrainian","Urdu","Vietnamese",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		studentTypeHomeLanguagePropEnum = append(studentTypeHomeLanguagePropEnum, v)
	}
}

const (

	// StudentHomeLanguageEnglish captures enum value "English"
	StudentHomeLanguageEnglish string = "English"

	// StudentHomeLanguageAlbanian captures enum value "Albanian"
	StudentHomeLanguageAlbanian string = "Albanian"

	// StudentHomeLanguageAmharic captures enum value "Amharic"
	StudentHomeLanguageAmharic string = "Amharic"

	// StudentHomeLanguageArabic captures enum value "Arabic"
	StudentHomeLanguageArabic string = "Arabic"

	// StudentHomeLanguageBengali captures enum value "Bengali"
	StudentHomeLanguageBengali string = "Bengali"

	// StudentHomeLanguageBosnian captures enum value "Bosnian"
	StudentHomeLanguageBosnian string = "Bosnian"

	// StudentHomeLanguageBurmese captures enum value "Burmese"
	StudentHomeLanguageBurmese string = "Burmese"

	// StudentHomeLanguageCantonese captures enum value "Cantonese"
	StudentHomeLanguageCantonese string = "Cantonese"

	// StudentHomeLanguageChinese captures enum value "Chinese"
	StudentHomeLanguageChinese string = "Chinese"

	// StudentHomeLanguageDutch captures enum value "Dutch"
	StudentHomeLanguageDutch string = "Dutch"

	// StudentHomeLanguageFarsi captures enum value "Farsi"
	StudentHomeLanguageFarsi string = "Farsi"

	// StudentHomeLanguageFrench captures enum value "French"
	StudentHomeLanguageFrench string = "French"

	// StudentHomeLanguageGerman captures enum value "German"
	StudentHomeLanguageGerman string = "German"

	// StudentHomeLanguageHebrew captures enum value "Hebrew"
	StudentHomeLanguageHebrew string = "Hebrew"

	// StudentHomeLanguageHindi captures enum value "Hindi"
	StudentHomeLanguageHindi string = "Hindi"

	// StudentHomeLanguageHmong captures enum value "Hmong"
	StudentHomeLanguageHmong string = "Hmong"

	// StudentHomeLanguageIlocano captures enum value "Ilocano"
	StudentHomeLanguageIlocano string = "Ilocano"

	// StudentHomeLanguageJapanese captures enum value "Japanese"
	StudentHomeLanguageJapanese string = "Japanese"

	// StudentHomeLanguageJavanese captures enum value "Javanese"
	StudentHomeLanguageJavanese string = "Javanese"

	// StudentHomeLanguageKaren captures enum value "Karen"
	StudentHomeLanguageKaren string = "Karen"

	// StudentHomeLanguageKhmer captures enum value "Khmer"
	StudentHomeLanguageKhmer string = "Khmer"

	// StudentHomeLanguageKorean captures enum value "Korean"
	StudentHomeLanguageKorean string = "Korean"

	// StudentHomeLanguageLaotian captures enum value "Laotian"
	StudentHomeLanguageLaotian string = "Laotian"

	// StudentHomeLanguageLatvian captures enum value "Latvian"
	StudentHomeLanguageLatvian string = "Latvian"

	// StudentHomeLanguageMalay captures enum value "Malay"
	StudentHomeLanguageMalay string = "Malay"

	// StudentHomeLanguageMandarin captures enum value "Mandarin"
	StudentHomeLanguageMandarin string = "Mandarin"

	// StudentHomeLanguageNepali captures enum value "Nepali"
	StudentHomeLanguageNepali string = "Nepali"

	// StudentHomeLanguageOromo captures enum value "Oromo"
	StudentHomeLanguageOromo string = "Oromo"

	// StudentHomeLanguagePolish captures enum value "Polish"
	StudentHomeLanguagePolish string = "Polish"

	// StudentHomeLanguagePortuguese captures enum value "Portuguese"
	StudentHomeLanguagePortuguese string = "Portuguese"

	// StudentHomeLanguagePunjabi captures enum value "Punjabi"
	StudentHomeLanguagePunjabi string = "Punjabi"

	// StudentHomeLanguageRomanian captures enum value "Romanian"
	StudentHomeLanguageRomanian string = "Romanian"

	// StudentHomeLanguageRussian captures enum value "Russian"
	StudentHomeLanguageRussian string = "Russian"

	// StudentHomeLanguageSamoan captures enum value "Samoan"
	StudentHomeLanguageSamoan string = "Samoan"

	// StudentHomeLanguageSerbian captures enum value "Serbian"
	StudentHomeLanguageSerbian string = "Serbian"

	// StudentHomeLanguageSomali captures enum value "Somali"
	StudentHomeLanguageSomali string = "Somali"

	// StudentHomeLanguageSpanish captures enum value "Spanish"
	StudentHomeLanguageSpanish string = "Spanish"

	// StudentHomeLanguageSwahili captures enum value "Swahili"
	StudentHomeLanguageSwahili string = "Swahili"

	// StudentHomeLanguageTagalog captures enum value "Tagalog"
	StudentHomeLanguageTagalog string = "Tagalog"

	// StudentHomeLanguageTamil captures enum value "Tamil"
	StudentHomeLanguageTamil string = "Tamil"

	// StudentHomeLanguageTelugu captures enum value "Telugu"
	StudentHomeLanguageTelugu string = "Telugu"

	// StudentHomeLanguageThai captures enum value "Thai"
	StudentHomeLanguageThai string = "Thai"

	// StudentHomeLanguageTigrinya captures enum value "Tigrinya"
	StudentHomeLanguageTigrinya string = "Tigrinya"

	// StudentHomeLanguageTurkish captures enum value "Turkish"
	StudentHomeLanguageTurkish string = "Turkish"

	// StudentHomeLanguageUkrainian captures enum value "Ukrainian"
	StudentHomeLanguageUkrainian string = "Ukrainian"

	// StudentHomeLanguageUrdu captures enum value "Urdu"
	StudentHomeLanguageUrdu string = "Urdu"

	// StudentHomeLanguageVietnamese captures enum value "Vietnamese"
	StudentHomeLanguageVietnamese string = "Vietnamese"

	// StudentHomeLanguageEmpty captures enum value ""
	StudentHomeLanguageEmpty string = ""
)

// prop value enum
func (m *Student) validateHomeLanguageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, studentTypeHomeLanguagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Student) validateHomeLanguage(formats strfmt.Registry) error {
	if swag.IsZero(m.HomeLanguage) { // not required
		return nil
	}

	// value enum
	if err := m.validateHomeLanguageEnum("home_language", "body", *m.HomeLanguage); err != nil {
		return err
	}

	return nil
}

func (m *Student) validateLastModified(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("last_modified", "body", "datetime", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Student) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

var studentTypeRacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Caucasian","Asian","Black or African American","American Indian","Hawaiian or Other Pacific Islander","Two or More Races","Unknown",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		studentTypeRacePropEnum = append(studentTypeRacePropEnum, v)
	}
}

const (

	// StudentRaceCaucasian captures enum value "Caucasian"
	StudentRaceCaucasian string = "Caucasian"

	// StudentRaceAsian captures enum value "Asian"
	StudentRaceAsian string = "Asian"

	// StudentRaceBlackOrAfricanAmerican captures enum value "Black or African American"
	StudentRaceBlackOrAfricanAmerican string = "Black or African American"

	// StudentRaceAmericanIndian captures enum value "American Indian"
	StudentRaceAmericanIndian string = "American Indian"

	// StudentRaceHawaiianOrOtherPacificIslander captures enum value "Hawaiian or Other Pacific Islander"
	StudentRaceHawaiianOrOtherPacificIslander string = "Hawaiian or Other Pacific Islander"

	// StudentRaceTwoOrMoreRaces captures enum value "Two or More Races"
	StudentRaceTwoOrMoreRaces string = "Two or More Races"

	// StudentRaceUnknown captures enum value "Unknown"
	StudentRaceUnknown string = "Unknown"

	// StudentRaceEmpty captures enum value ""
	StudentRaceEmpty string = ""
)

// prop value enum
func (m *Student) validateRaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, studentTypeRacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Student) validateRace(formats strfmt.Registry) error {
	if swag.IsZero(m.Race) { // not required
		return nil
	}

	// value enum
	if err := m.validateRaceEnum("race", "body", *m.Race); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this student based on the context it is used
func (m *Student) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnrollments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Student) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {
		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Student) contextValidateEnrollments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Enrollments); i++ {

		if m.Enrollments[i] != nil {
			if err := m.Enrollments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("enrollments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Student) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Student) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Student) UnmarshalBinary(b []byte) error {
	var res Student
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
