package validator

import (
	"net"
	"net/url"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsAlpha checks if the string contains only letters (a-zA-Z). Empty string is valid.
func IsAlpha(str string) bool {
	if IsNull(str) {
		return true
	}

	return alphaRegexp.MatchString(str)
}

// IsAlphanumeric checks if the string contains only letters and numbers. Empty string is valid.
func IsAlphanumeric(str string) bool {
	if IsNull(str) {
		return true
	}

	return alphanumericRegexp.MatchString(str)
}

// IsEmail checks if the string is an email.
func IsEmail(str string) bool {
	return emailRegexp.MatchString(str)
}

// IsExistingEmail checks if the string is an email of existing domain
func IsExistingEmail(email string) bool {
	if len(email) < 6 || len(email) > 254 {
		return false
	}

	at := strings.LastIndex(email, "@")
	if at <= 0 || at > len(email)-3 {
		return false
	}

	user, host := email[:at], email[at+1:]
	if len(user) > 64 {
		return false
	}
	if host == "localhost" || host == "example.com" {
		return true
	}
	if userDotRegexp.MatchString(user) || !userRegexp.MatchString(user) || !hostRegexp.MatchString(host) {
		return false
	}

	if _, err := net.LookupMX(host); err != nil {
		if _, err := net.LookupIP(host); err != nil {
			return false
		}
	}

	return true
}

// IsNull checks if the string is null.
func IsNull(str string) bool {
	return len(str) == 0
}

// IsNotNull checks if the string is not null.
func IsNotNull(str string) bool {
	return !IsNull(str)
}

// IsNumeric checks if the string contains only numbers. Empty string is valid.
func IsNumeric(str string) bool {
	if IsNull(str) {
		return true
	}

	return numericRegexp.MatchString(str)
}

const maxURLRuneCount = 2083
const minURLRuneCount = 3

// IsURL checks if the string is an URL.
func IsURL(str string) bool {
	if str == "" || utf8.RuneCountInString(str) >= maxURLRuneCount || len(str) <= minURLRuneCount || strings.HasPrefix(str, ".") {
		return false
	}

	strTemp := str
	if strings.Contains(str, ":") && !strings.Contains(str, "://") {
		// support no indicated urlscheme but with colon for port number
		// http:// is appended so url.Parse will succeed, strTemp used so it does not impact rxURL.MatchString
		strTemp = "http://" + str
	}

	u, err := url.Parse(strTemp)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	return urlRegexp.MatchString(str)
}

// IsRequestURL checks if the string rawurl, assuming
// it was received in an HTTP request, is a valid
// URL confirm to RFC 3986
func IsRequestURL(rawurl string) bool {
	url, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false // couldn't even parse the rawurl
	}
	if len(url.Scheme) == 0 {
		return false // no Scheme found
	}
	return true
}

// IsRequestURI checks if the string rawurl, assuming
// it was received in an HTTP request, is an
// absolute URI or an absolute path.
func IsRequestURI(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	return err == nil
}

// IsUTFDigit checks if the string contains only unicode radix-10 decimal digits. Empty string is valid.
func IsUTFDigit(str string) bool {
	if IsNull(str) {
		return true
	}

	if strings.IndexAny(str, "+-") > 0 {
		return false
	}

	if len(str) > 1 {
		str = strings.TrimPrefix(str, "-")
		str = strings.TrimPrefix(str, "+")
	}
	for _, c := range str {
		if !unicode.IsDigit(c) { // digits && minus sign are ok
			return false
		}
	}

	return true
}

// IsUTFLetter checks if the string contains only unicode letter characters.
// Similar to IsAlpha but for all languages. Empty string is valid.
func IsUTFLetter(str string) bool {
	if IsNull(str) {
		return true
	}

	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}

	return true
}

// IsUTFLetterNumeric checks if the string contains only unicode letters and numbers. Empty string is valid.
func IsUTFLetterNumeric(str string) bool {
	if IsNull(str) {
		return true
	}

	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) { // letters && numbers are ok
			return false
		}
	}

	return true
}

// IsUTFNumeric checks if the string contains only unicode numbers of any kind.
// Numbers can be 0-9 but also Fractions ¾,Roman Ⅸ and Hangzhou 〩. Empty string is valid.
func IsUTFNumeric(str string) bool {
	if IsNull(str) {
		return true
	}

	if strings.IndexAny(str, "+-") > 0 {
		return false
	}

	if len(str) > 1 {
		str = strings.TrimPrefix(str, "-")
		str = strings.TrimPrefix(str, "+")
	}
	for _, c := range str {
		if !unicode.IsNumber(c) { // numbers && minus sign are ok
			return false
		}
	}

	return true
}
