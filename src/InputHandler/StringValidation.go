package InputHandler

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"regexp"
	"strconv"
)

func IsNameValid(name *string) bool {
	nameLength := len(*name)
	if nameLength > 0 {
		if nameLength <= 50 {
			matched, _ := regexp.MatchString("^[a-zA-Z0-9\\s]+$", *name)
			return matched
		}
	}
	return false
}

func IsTeamNameValid(name *string) bool {
	teamNameLength := len(*name)
	if teamNameLength > 0 {
		if teamNameLength <= 60 {
			matched, _ := regexp.MatchString("^[a-zA-Z0-9-_\\s]+$", *name)
			return matched
		}
	}
	return false
}

func IsAValidHex(hexString *string, targetLength int) bool {
	matched, _ := regexp.MatchString("[a-zA-Z0-9]{"+strconv.Itoa(targetLength)+"}", *hexString)
	return matched
}

func IsUsernameHashValid(usernameHash *string) bool {
	return IsAValidHex(usernameHash, 64)
}

func ToStringSlice(original []interface{}) []string {
	originalLength := len(original)
	result := make([]string, originalLength)
	for index, value := range original {
		result[index] = value.(string)
	}
	return result
}

func IsLanguageValid(languageName *string) bool {
	for _, language := range Neo4jDatabase.Languages {
		if language == *languageName {
			return true
		}
	}
	return false
}

func IsGenderValid(genderName *string) bool {
	for _, gender := range Neo4jDatabase.Genders {
		if *genderName == gender {
			return true
		}
	}
	return false
}

func IsLocationValid(locationName *string) bool {
	for _, country := range Neo4jDatabase.Countries {
		if *locationName == country {
			return true
		}
	}
	return false
}

func IsNationalityValid(nationalityName *string) bool {
	for _, nationality := range Neo4jDatabase.Nationalities {
		if *nationalityName == nationality {
			return true
		}
	}
	return false
}

func IsEmailValid(email *string) bool {
	if IsStringValid(email) {
		// Regex extracted from: "https://www.golangprograms.com/regular-expression-to-validate-email-address.html"
		matched, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", *email)
		return matched
	}
	return false
}

func IsCountryCodeValid(countryCode *string) bool {
	for _, dialCode := range SQLDatabase.ValidDialCodes {
		if *countryCode == dialCode {
			return true
		}
	}
	return false
}

// Requires implementation
func IsPhoneNumberValid(phoneNumber *string) bool {
	matched, _ := regexp.MatchString("^\\d{10}$", *phoneNumber)
	return matched
}

func IsUsernameValid(username *string) bool {
	usernameLength := len(*username)
	if usernameLength >= 4 {
		if usernameLength <= 20 {
			if matched, _ := regexp.MatchString("\\s", *username); !matched {
				if IsStringValid(username) {
					return true
				}
			}
		}
	}
	return false
}

func IsPasswordValid(password *string) bool {
	passwordLength := len(*password)
	if passwordLength >= 8 {
		if passwordLength <= 32 {
			if matched, _ := regexp.MatchString("[[:punct:]]", *password); matched {
				if matched, _ := regexp.MatchString("[[:lower:]]", *password); matched {
					if matched, _ := regexp.MatchString("[[:upper:]]", *password); matched {
						if matched, _ := regexp.MatchString("[[:digit:]]", *password); matched {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func IsStringValid(challengeString *string) bool {
	if len(*challengeString) > 0 {
		if matched, _ := regexp.MatchString("^(\\s|[A-Za-z0-9!\"#$%&'()*+,\\-./:;<=>?@[\\\\\\]^_`{|}~])+$", *challengeString); matched {
			if matched, _ := regexp.MatchString("\\s{2,}", *challengeString); !matched {
				return true
			}
		}
	}
	return false
}

func IsDescriptionValid(description *string) bool {
	descriptionLength := len(*description)
	if descriptionLength > 0 {
		if descriptionLength <= 350 {
			matched, _ := regexp.MatchString("^[a-zA-Z0-9\\s]+$", *description)
			return matched
		}
	}
	return false
}
