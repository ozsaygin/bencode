package main

import (
	"fmt"
	"regexp"
)

// Strings are length-prefixed base ten followed by a colon and the string. For example 4:spam corresponds to 'spam'.

// Integers are represented by an 'i' followed by the number in base 10 followed by an 'e'. For example i3e corresponds to 3 and i-3e corresponds to -3. Integers have no size limitation. i-0e is invalid. All encodings with a leading zero, such as i03e, are invalid, other than i0e, which of course corresponds to 0.

// Lists are encoded as an 'l' followed by their elements (also bencoded) followed by an 'e'. For example l4:spam4:eggse corresponds to ['spam', 'eggs'].

// Dictionaries are encoded as a 'd' followed by a list of alternating keys and their corresponding values followed by an 'e'. For example, d3:cow3:moo4:spam4:eggse corresponds to {'cow': 'moo', 'spam': 'eggs'} and d4:spaml1:a1:bee corresponds to {'spam': ['a', 'b']}. Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics).

// Decode unmarshalls bencoded string to map object
func Decode(data string) map[string]interface{} {

	m := make(map[string]interface{})

	// process dictionary
	if data[0] == 'd' {

		for len(data) > 1 {
			fmt.Println(len(data))
			fmt.Println("data is: " + data)
			// first element must be string
			// hence, it must be in form of digit:word
			dataLen := len(data)
			data = data[1 : dataLen-1]

			// first key comes and then value
			// repeat the same process
			re := regexp.MustCompile(`(\d+)\:([a-zA-Z]+)`)
			matches := re.FindStringSubmatch(data)
			var key string
			if len(matches) > 1 {
				key = matches[2]
			}
			data = data[len(matches[0]):]
			// Process the value
			// For now value can be only string
			re = regexp.MustCompile(`(\d+)\:([a-zA-Z]+)`)
			matches = re.FindStringSubmatch(data)
			if len(matches) > 1 {
				value := matches[2]
				m[key] = value
				fmt.Println(m)
			}
			data = data[len(matches[0]):]
		}
	}
	return nil
}

func main() {

	mapExample := "d3:cow3:moo4:spam4:eggse"
	Decode(mapExample)
}
