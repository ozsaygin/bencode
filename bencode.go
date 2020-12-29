package bencode

import (
	"fmt"
	"regexp"
	"strconv"
)

// Strings are length-prefixed base ten followed by a colon and the string. For example 4:spam corresponds to 'spam'.

// Integers are represented by an 'i' followed by the number in base 10 followed by an 'e'. For example i3e corresponds to 3 and i-3e corresponds to -3. Integers have no size limitation. i-0e is invalid. All encodings with a leading zero, such as i03e, are invalid, other than i0e, which of course corresponds to 0.

// Lists are encoded as an 'l' followed by their elements (also bencoded) followed by an 'e'. For example l4:spam4:eggse corresponds to ['spam', 'eggs'].

// Dictionaries are encoded as a 'd' followed by a list of alternating keys and their corresponding values followed by an 'e'. For example, d3:cow3:moo4:spam4:eggse corresponds to {'cow': 'moo', 'spam': 'eggs'} and d4:spaml1:a1:bee corresponds to {'spam': ['a', 'b']}. Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics).

type DecodeState struct {
	data string
	v    interface{}
}

// Decode unmarshalls bencoded string to map object
func Decode(data string) interface{} {

	prefix := data[0]

	for data != "" {
		switch prefix {

		// case 'i':
		// 	re := regexp.MustCompile(`i(\-?\d+)e`)
		// 	matches := re.FindStringSubmatch(data)
		// 	value, _ := strconv.Atoi(matches[0])
		// 	return value

		// case 'l':
		// 	re := regexp.MustCompile(`l(.*)e`)
		// matches := re.FindStringSubmatch(data)
		// 	value, _ := strconv.Atoi(matches[0])
		// 	return value

		case 'd':
			re := regexp.MustCompile(`d(.*)e`)
			matches := re.FindStringSubmatch(data)
			inner := matches[1]
			dict := make(map[string]interface{})

			key := Decode(inner).(string)
			value := Decode(inner)

			dict[key] = value
			return

		default:
			re := regexp.MustCompile(`(\d+)\:([a-zA-Z]+)`)
			matches := re.FindStringSubmatch(data)
			word := matches[2]
			return word
		}
	}
}

// Encode marshalls any input type into bencoded format
func Encode(v interface{}) string {

	switch v.(type) {

	case string:

		word := v.(string)
		length := strconv.Itoa(len(word))
		encoded := fmt.Sprintf("%s:%s", length, word)
		return encoded

	case int:

		number := v.(int)
		encoded := fmt.Sprintf("i%de", number)
		return encoded

	case []interface{}:

		content := ""
		for elm := range v.([]interface{}) {
			content += Encode(elm)
		}

		encoded := fmt.Sprintf("l%se", content)
		return encoded

	case map[string]interface{}:

		content := ""
		for k, v := range v.(map[string]interface{}) {
			content += fmt.Sprintf("%s%s", Encode(k), Encode(v))
		}

		encoded := fmt.Sprintf("d%se", content)
		return encoded

	default:
		return ""

	}
}

// func main() {

// 	mapExample := "d3:cow3:moo4:spam4:eggse"
// 	Decode(mapExample)
// }
