package validate_test

import (
	"testing"

	"github.com/copartner6412/input/validate"
)

func TestCountryNameSuccessfulForValidInput(t *testing.T) {
	t.Parallel()

	testCases := []string{
		"Afghanistan", "Albania", "Algeria", "Andorra", "Angola", "Antigua and Barbuda", "Argentina",
		"Armenia", "Australia", "Austria", "Azerbaijan", "Bahamas", "Bahrain", "Bangladesh", "Barbados",
		"Belarus", "Belgium", "Belize", "Benin", "Bhutan", "Bolivia", "Bosnia and Herzegovina", "Botswana",
		"Brazil", "Brunei", "Bulgaria", "Burkina Faso", "Burundi", "Cabo Verde", "Cambodia", "Cameroon",
		"Canada", "Central African Republic", "Chad", "Chile", "China", "Colombia", "Comoros",
		"Congo", "Costa Rica", "Croatia", "Cuba", "Cyprus",
		"Czechia", "Denmark", "Djibouti", "Dominica", "Dominican Republic", "Ecuador",
		"Egypt", "El Salvador", "Equatorial Guinea", "Eritrea", "Estonia", "Eswatini", "Ethiopia",
		"Fiji", "Finland", "France", "Gabon", "Gambia", "Georgia", "Germany", "Ghana", "Greece",
		"Grenada", "Guatemala", "Guinea", "Guinea-Bissau", "Guyana", "Haiti", "Honduras", "Hungary",
		"Iceland", "India", "Indonesia", "Iran", "Iraq", "Ireland", "Israel", "Italy",
		"Jamaica", "Japan", "Jordan", "Kazakhstan", "Kenya", "Kiribati", "Kuwait", "Kyrgyzstan",
		"Laos", "Latvia", "Lebanon", "Lesotho", "Liberia", "Libya", "Liechtenstein", "Lithuania",
		"Luxembourg", "Madagascar", "Malawi", "Malaysia", "Maldives", "Mali", "Malta", "Marshall Islands",
		"Mauritania", "Mauritius", "Mexico", "Micronesia", "Moldova", "Monaco", "Mongolia", "Montenegro",
		"Morocco", "Mozambique", "Myanmar", "Namibia", "Nauru", "Nepal", "Netherlands", "New Zealand",
		"Nicaragua", "Niger", "Nigeria", "North Macedonia", "Norway", "Oman",
		"Pakistan", "Palau", "Panama", "Papua New Guinea", "Paraguay", "Peru", "Philippines", "Poland",
		"Portugal", "Qatar", "Romania", "Russia", "Rwanda", "Saint Kitts and Nevis", "Saint Lucia",
		"Saint Vincent and the Grenadines", "Samoa", "San Marino", "Sao Tome and Principe", "Saudi Arabia",
		"Senegal", "Serbia", "Seychelles", "Sierra Leone", "Singapore", "Slovakia", "Slovenia",
		"Solomon Islands", "Somalia", "South Africa", "South Sudan", "Spain",
		"Sri Lanka", "Sudan", "Suriname", "Sweden", "Switzerland", "Syria", "Taiwan", "Tajikistan",
		"Tanzania", "Thailand", "Togo", "Tonga", "Trinidad and Tobago", "Tunisia", "Turkey",
		"Turkmenistan", "Tuvalu", "Uganda", "Ukraine", "United Arab Emirates", "United Kingdom",
		"United States of America", "Uruguay", "Uzbekistan", "Vanuatu", "Holy See", "Venezuela", "Viet Nam",
		"Yemen", "Zambia", "Zimbabwe", "North Korea", "South Korea",
	}

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			err := validate.CountryName(tc)
			if err != nil {
				t.Errorf("expected no error for valid input: %q, but got error: %v", tc, err)
			}
		})
	}
}

func TestCountryNameFailsForInvalidInput(t *testing.T) {
	t.Parallel()

	testCaseGroups := map[string][]string{
		"empty names": {
			"",    // Completely empty
			"   ", // Only whitespace
		},
		"non-alphabetical characters": {
			"12345",       // Only digits
			"!@#$%",       // Special characters
			"France123",   // Letters with numbers
			"Germany!",    // Letters with special characters
			"Brazil@Rio",  // Letters with special characters in middle
			"China*China", // Invalid use of special characters
		},
		"too short or too long": {
			"A",                                // One letter, too short
			"Fr",                               // Two letters, still too short
			"AntidisestablishmentarianismPlus", // Exceeds reasonable length for country names
			"ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ", // Extremely long name
		},
		"capitalization issues": {
			"france",  // All lowercase
			"GERMANY", // All uppercase
			"bRAZIL",  // Mixed improper casing
		},
		"misformatted names": {
			" UnitedStates",   // Leading whitespace
			"United States ",  // Trailing whitespace
			"   Canada  ",     // Extra spaces
			"New  Zealand",    // Double space in the middle
			"Kingdom, United", // Improper format, comma-separated
		},
		"duplicate or repetitive letters": {
			"AAA",  // Repeated characters, not a real country
			"ZZZZ", // Invalid letters with repetition
		},
		"non-existent countries": {
			"Atlantis",     // Fictional country
			"Elbonia",      // Fictional country
			"Wakanda",      // Fictional country
			"Narnia",       // Fictional country
			"Middle Earth", // Fictional country
		},
	}

	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCaseGroup {
				err := validate.CountryName(tc)
				if err == nil {
					t.Errorf("expected error for invalid input: %q", tc)
				}
			}
		})
	}
}

func TestCountryCode2SuccessfulForValidInput(t *testing.T) {
	t.Parallel()

	testCases := []string{
		"US", "GB", "FR", "JP", "CN",
		"AU", "CA", "DE", "IN", "BR",
		"RU", "MX", "ZA", "KR", "NG",
		"IT", "ES", "SE", "CH", "NL",
	}

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			err := validate.CountryCode2(tc)
			if err != nil {
				t.Errorf("expected no error for valid input: %q, but got error: %v", tc, err)
			}
		})
	}
}

func TestCountryCode2FailsForInvalidInput(t *testing.T) {
	t.Parallel()

	testCaseGroups := map[string][]string{
		"empty code":                    {""},                       // Empty input
		"1 letter code":                 {"a", "f", "B", "G"},       // 1-letter code, should fail
		"too long code":                 {"USA", "FRA", "CHN"},      // 3-letter code, should fail for alpha-2 check
		"invalid 2 letter code":         {"AB", "GK", "NQ"},         // Invalid country codes that do not exist
		"lowercase valid 2 letter code": {"us", "gb", "fr"},         // Lowercase valid codes, should fail because country codes are usually uppercase
		"numeric code":                  {"12", "A1", "1B"},         // Numbers mixed in or fully numeric should fail
		"special characters":            {"@", "&", "#$", "X#"},     // Special characters in the code
		"code with space":               {" U", "F ", " GB", "FR "}, // Codes with spaces, should fail
		"mixed case invalid code":       {"Us", "Gb", "Fr"},         // Mixed case valid codes should fail because ISO 3166-1 alpha-2 codes are uppercase
	}

	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCaseGroup {
				err := validate.CountryCode2(tc)
				if err == nil {
					t.Errorf("expected error for invalid input: %q", tc)
				}
			}
		})
	}
}

func TestCountryCode3SuccessfulForValidInput(t *testing.T) {
	t.Parallel()

	testCases := []string{
		"USA", "GBR", "FRA", "JPN", "CHN", // Common country codes
		"AUS", "CAN", "DEU", "IND", "BRA", // Large and influential countries
		"RUS", "MEX", "ZAF", "KOR", "NGA", // KR => KOR (South Korea)
		"ITA", "ESP", "SWE", "CHE", "NLD", // European countries
		"DNK", "NOR", "FIN", "ISL", "NZL", // Scandinavian and island countries
		"ARG", "COL", "PER", "CHL", "URY", // South American countries
		"SGP", "THA", "VNM", "PHL", "MYS", // Southeast Asian countries
		"ZWE", "KEN", "ETH", "TZA", "UGA", // African countries
		"ARE", "SAU", "ISR", "JOR", "QAT", // Middle Eastern countries
		"FJI", "TUV", "WSM", "PNG", "TON", // Pacific island countries
		"VAT", "MCO", "SMR", "LIE", "AND", // Small European nations
		"TLS", "MRT", "CPV", "BHS", "COM", // Lesser-known countries
		"MDG", "MLT", "BTN", "BRN", "LAO", // More small countries
		"KWT", "OMN", "BHR", "KGZ", "TJK", // Central Asian countries
	}

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			err := validate.CountryCode3(tc)
			if err != nil {
				t.Errorf("expected no error for valid input: %q, but got error: %v", tc, err)
			}
		})
	}
}

func TestCountryCode3FailsForInvalidInput(t *testing.T) {
	t.Parallel()

	testCaseGroups := map[string][]string{
		"empty string": {""},
		"too short":    {"A", "AB"},
		"too long":     {"ABCD", "ABCDE", "ABCDEFG"}, // Strings longer than 3 characters
		"invalid characters": {
			"AB1", "A2B", "123", // Numeric characters
			"abC", "aBc", "abc", // Lowercase characters
			"@BC", "A$C", "A C", // Special characters and spaces
		},
		"whitespace issues": {
			" USA", "USA ", " U S", // Leading, trailing, or embedded spaces
			"  ", "\t", "\n", // Only whitespace characters
		},
		"invalid formatted codes": {
			"XYZ", "ZZZ", "QQQ", "AAA", // Invalid but properly formatted
			"UNK", "XXX", // ISO reserved or non-assigned codes
		},
	}

	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCaseGroup {
				err := validate.CountryCode3(tc)
				if err == nil {
					t.Errorf("expected error for invalid input: %q", tc)
				}
			}
		})
	}
}
