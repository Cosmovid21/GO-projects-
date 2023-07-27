package valid
import (
	"unicode"
	"testing"
)

func Password(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                int
	)

	 for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}
 
	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}
 
	return true
}

func TestPassword(t *testing.T) {
	tests := []struct {
		name  string
		pass  string
		valid bool
	}{
		{
			"NoCharacterAtAll",
			"",
			false,
		},
		{
			"JustEmptyStringAndWhitespace",
			" \n\t\r\v\f ",
			false,
		},
		{
			"MixtureOfEmptyStringAndWhitespace",
			"U u\n1\t?\r1\v2\f34",
			false,
		},
		{
			"MissingUpperCaseString",
			"uu1?1234",
			false,
		},
		{
			"MissingLowerCaseString",
			"UU1?1234",
			false,
		},
		{
			"MissingNumber",
			"Uua?aaaa",
			false,
		},
		{
			"MissingSymbol",
			"Uu101234",
			false,
		},
		{
			"LessThanRequiredMinimumLength",
			"Uu1?123",
			false,
		},
		{
			"ValidPassword",
			"Uu1?1234",
			true,
		},
	}
 
	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if c.valid != Password(c.pass) {
				t.Fatal("invalid password")
			}
		})
	}
}
