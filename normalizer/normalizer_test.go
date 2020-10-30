package normalizer_test

import (
	"testing"

	"github.com/cyrildever/go-utls/normalizer"
	"gotest.tools/assert"
)

// TestUniformize ...
func TestUniformize(t *testing.T) {
	data := " cAfé#@~eT*%chocolAT )"
	uniformized := normalizer.Uniformize(data)
	assert.Equal(t, uniformized, "CAFE ET CHOCOLAT")
}

// TestAny ...
func TestAny(t *testing.T) {
	normalized, _ := normalizer.Normalize("Any   string to normalize().", normalizer.Any)
	assert.Equal(t, normalized, "ANY STRING TO NORMALIZE")

	// An empty string doesn't raise an error
	normalized, err := normalizer.Normalize("", normalizer.Any)
	assert.Assert(t, err == nil)
	assert.Equal(t, normalized, "")
}

// TestAddressLine ...
func TestAddressLine(t *testing.T) {
	// French address line 2
	normalized, _ := normalizer.Normalize("c/o Mr et Mme Dupont", normalizer.AddressLine)
	assert.Equal(t, normalized, "C O MR MME DUPONT")

	// French address line 3
	normalized, _ = normalizer.Normalize("Bât. 4, escalier G", normalizer.AddressLine)
	assert.Equal(t, normalized, "BAT 4 ESC G")

	// French address line 4
	normalized, _ = normalizer.Normalize("128 rue du Faubourg Saint Honoré ", normalizer.AddressLine)
	assert.Equal(t, normalized, "128 RUE FBG ST HONORE")

	normalized, _ = normalizer.Normalize("*** rue du \nHenner", normalizer.AddressLine)
	assert.Equal(t, normalized, "RUE HENNER")

	_, err := normalizer.Normalize("$µ%*+^)@", normalizer.AddressLine)
	assert.Error(t, err, "unable to build a normalized string")

	// French address line 5
	normalized, _ = normalizer.Normalize("Lieu-dit du domaine Vert", normalizer.AddressLine)
	assert.Equal(t, normalized, "LIEU DIT DOM VERT")

	// French address line 6
	normalized, _ = normalizer.Normalize("$.75009        Paris", normalizer.AddressLine)
	assert.Equal(t, normalized, "75009 PARIS")

	normalized, _ = normalizer.Normalize("75948 Paris Cedex 19", normalizer.AddressLine)
	assert.Equal(t, normalized, "75948 PARIS CDX 19")
}

// TestCity ...
func TestCity(t *testing.T) {
	normalized, _ := normalizer.Normalize("Jouy-en-Josas ", normalizer.City)
	assert.Equal(t, normalized, "JOUY EN JOSAS")

	normalized, _ = normalizer.Normalize("Paris        Cedex", normalizer.City)
	assert.Equal(t, normalized, "PARIS")

	normalized, _ = normalizer.Normalize("CDX y", normalizer.City)
	assert.Equal(t, normalized, "Y")

	normalized, _ = normalizer.Normalize("Paris Cedex   20", normalizer.City)
	assert.Equal(t, normalized, "PARIS")
}

// TestCodePostalFrance ...
func TestCodePostalFrance(t *testing.T) {
	normalized, _ := normalizer.Normalize("12345 ", normalizer.CodePostalFrance)
	assert.Equal(t, normalized, "12345")

	normalized, err := normalizer.Normalize("aghfkhgk", normalizer.CodePostalFrance)
	assert.Error(t, err, "invalid code postal")
	assert.Equal(t, normalized, "")

	normalized, _ = normalizer.Normalize("2A165", normalizer.CodePostalFrance)
	assert.Equal(t, normalized, "20165")
}

// TestDateOfBirth ...
func TestDateOfBirth(t *testing.T) {
	normalized, _ := normalizer.Normalize("1564/04/23", normalizer.DateOfBirth, "yyyy/MM/dd")
	assert.Equal(t, normalized, "23/04/1564")

	normalized, _ = normalizer.Normalize("95/04/23", normalizer.DateOfBirth, "yy/MM/dd")
	assert.Equal(t, normalized, "23/04/1995")

	normalized, _ = normalizer.Normalize("1472720661", normalizer.DateOfBirth, normalizer.TIMESTAMP)
	assert.Equal(t, normalized, "01/09/2016")

	normalized, _ = normalizer.Normalize("1472720661276", normalizer.DateOfBirth, normalizer.TIMESTAMP_MILLIS)
	assert.Equal(t, normalized, "01/09/2016")

	normalized, _ = normalizer.Normalize("10 MAY 1970", normalizer.DateOfBirth, "DD MMM YYYY", normalizer.ISO_FORMAT)
	assert.Equal(t, normalized, "19700510")

	// Hours, minutes and seconds are not supported yet, so strip them beforehand!
	_, err := normalizer.Normalize("24/04/2010 12:00:00", normalizer.DateOfBirth, "dd/MM/yyyy hh:mm:ss")
	assert.Error(t, err, "parsing time \"24/04/2010 12:00:00\" as \"02/01/2006 HH:01:SS\": cannot parse \"12:00:00\" as \" HH:\"")

	_, err = normalizer.Normalize("1969", normalizer.DateOfBirth, normalizer.FRENCH_FORMAT)
	assert.Error(t, err, "parsing time \"1969\" as \"02/01/2006\": cannot parse \"69\" as \"/\"")

	_, err = normalizer.Normalize("not-a-date", normalizer.DateOfBirth)
	assert.Error(t, err, "parsing time \"not-a-date\" as \"20060102\": cannot parse \"not-a-date\" as \"2006\"")

	_, err = normalizer.Normalize("not-a-timestamp", normalizer.DateOfBirth, normalizer.TIMESTAMP)
	assert.Error(t, err, "strconv.ParseInt: parsing \"not-a-timestamp\": invalid syntax")
}

// TestDepartementFrance ...
func TestDepartementFrance(t *testing.T) {
	normalized, _ := normalizer.Normalize("20167", normalizer.DepartementFrance)
	assert.Equal(t, normalized, "2A")

	normalized, _ = normalizer.Normalize(" 75009 ", normalizer.DepartementFrance)
	assert.Equal(t, normalized, "75")
}

// TestEmail ...
func TestEmail(t *testing.T) {
	ref := "cdever@edgewhere.fr"

	data := " cdever@EDGEWHERE.fr"
	normalized, err := normalizer.Normalize(data, normalizer.Email)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, ref, normalized)

	normalized, err = normalizer.Normalize("gregoire.albizzati@edge@where.fr", normalizer.Email)
	assert.Error(t, err, "invalid email")
	assert.Equal(t, normalized, "")

	wrongData := "pretty-long string that's not an email@at_all.com"
	_, err = normalizer.Normalize(wrongData, normalizer.Email)
	assert.Error(t, err, "invalid email")

	tooShort := "t@t.t"
	_, err = normalizer.Normalize(tooShort, normalizer.Email)
	assert.Error(t, err, "invalid email")

	// A valid email shouldn't be changed
	normalized, err = normalizer.Normalize(ref, normalizer.Email)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, ref, normalized)
}

// TestFirstName ...
func TestFirstName(t *testing.T) {
	normalized, _ := normalizer.Normalize("Cyril", normalizer.FirstName)
	assert.Equal(t, normalized, "CYRIL")

	normalized, _ = normalizer.Normalize("J Louis", normalizer.FirstName)
	assert.Equal(t, normalized, "JEAN LOUIS")

	normalized, _ = normalizer.Normalize("Unknown", normalizer.FirstName)
	assert.Equal(t, normalized, "UNKNOWN")

	// A weird string will raise an error and return an empty string anyway
	normalized, err := normalizer.Normalize("#@~*%", normalizer.FirstName)
	assert.Error(t, err, "unable to build a normalized string")
	assert.Equal(t, normalized, "")

	_, err = normalizer.Normalize("", normalizer.FirstName)
	assert.Error(t, err, "invalid empty string")
}

// TestMobile ...
func TestMobile(t *testing.T) {
	normalized, _ := normalizer.Normalize("0623456789", normalizer.Mobile)
	assert.Equal(t, normalized, "+33 (0) 623 456 789")

	normalized, _ = normalizer.Normalize("07-23-45-67-89", normalizer.Mobile)
	assert.Equal(t, normalized, "+33 (0) 723 456 789")

	// Valid landline phone numbers aren't mobile phones
	normalized, err := normalizer.Normalize("0123456789", normalizer.Mobile)
	assert.Error(t, err, "not a mobile phone")
	assert.Equal(t, normalized, "")
}

// TestPhoneNumber ...
func TestPhoneNumber(t *testing.T) {
	normalized, _ := normalizer.Normalize("0123456789", normalizer.PhoneNumber)
	assert.Equal(t, normalized, "+33 (0) 123 456 789")

	normalized, _ = normalizer.Normalize("01-23-45-67-89", normalizer.PhoneNumber)
	assert.Equal(t, normalized, "+33 (0) 123 456 789")

	normalized, err := normalizer.Normalize("01-23-45-67-89-10", normalizer.PhoneNumber)
	assert.Error(t, err, "invalid phone number")
	assert.Equal(t, normalized, "")

	normalized, _ = normalizer.Normalize("0033 0 1 23 45 67 89", normalizer.PhoneNumber)
	assert.Equal(t, normalized, "+33 (0) 123 456 789")

	// Mobile phones are phone numbers by definition
	normalized, _ = normalizer.Normalize("06.23.45.67.89", normalizer.PhoneNumber)
	assert.Equal(t, normalized, "+33 (0) 623 456 789")
}

// TestStreetNumber ...
func TestStreetNumber(t *testing.T) {
	normalized, _ := normalizer.Normalize("8bis", normalizer.StreetNumber)
	assert.Equal(t, normalized, "8B")

	normalized, _ = normalizer.Normalize("11SEXIES", normalizer.StreetNumber)
	assert.Equal(t, normalized, "11S")

	normalized, _ = normalizer.Normalize("221 bis", normalizer.StreetNumber)
	assert.Equal(t, normalized, "221B")

	normalized, _ = normalizer.Normalize("1 bis C", normalizer.StreetNumber)
	assert.Equal(t, normalized, "1B C")
}

// TestTitle ...
func TestTitle(t *testing.T) {
	normalized, _ := normalizer.Normalize("Mademoiselle", normalizer.Title)
	assert.Equal(t, normalized, "2")

	normalized, _ = normalizer.Normalize("Docteur", normalizer.Title)
	assert.Equal(t, normalized, "0")

	normalized, err := normalizer.Normalize(" ", normalizer.Title)
	assert.Error(t, err, "unable to build a normalized string")
	assert.Equal(t, normalized, "")

	normalized, _ = normalizer.Normalize("1", normalizer.Title)
	assert.Equal(t, normalized, "1")

	normalized, _ = normalizer.Normalize("Monsieur et Madame", normalizer.Title)
	assert.Equal(t, normalized, "0")
}
