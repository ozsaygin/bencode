package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DecodeState struct {
	data string
	v    interface{}
}

// Decode unmarshalls bencoded string to map object
func (ds *DecodeState) Decode() interface{} {

	prefix := ds.data[0]

	for ds.data != "" {
		switch prefix {

		case 'i':
			re := regexp.MustCompile(`i(\-?\d+)e`)
			matches := re.FindStringSubmatch(ds.data)
			value, _ := strconv.Atoi(matches[1])
			ds.data = strings.TrimPrefix(ds.data, matches[0])
			return value

		case 'l':
			re := regexp.MustCompile(`l(.*)e`)
			matches := re.FindStringSubmatch(ds.data)
			elems := matches[1]
			lst := []interface{}{}

			ids := &DecodeState{data: elems}
			for ids.data != "" {
				elm := ids.Decode()
				lst = append(lst, elm)
			}
			ds.data = strings.TrimPrefix(ds.data, matches[0])
			return lst

		case 'd':
			re := regexp.MustCompile(`d(.*)e`)
			matches := re.FindStringSubmatch(ds.data)
			inner := matches[1]
			dict := make(map[string]interface{})

			ids := &DecodeState{data: inner}
			for ids.data != "" {
				key := ids.Decode().(string)
				value := ids.Decode()

				dict[key] = value
			}

			ds.data = strings.TrimPrefix(ds.data, matches[0])
			return dict

		default:
			re := regexp.MustCompile(`(\d+)\:(.+)`)
			matches := re.FindStringSubmatch(ds.data)

			ss := matches[1]
			length, _ := strconv.Atoi(ss)
			word := matches[2]
			word = word[:length]

			encodedStr := strconv.Itoa(length) + ":" + word

			ds.data = strings.TrimPrefix(ds.data, encodedStr)

			return word
		}
	}
	return nil
}

func main() {

	// data := "4:spam"
	// data := "d3:cat4:meow3:cow3:mooe"
	// data := "d4:spaml1:a1:bee"
	data := `d1:0d4:body50:quia e rerum est autem sunt rem eveniet architecto2:idi1e5:title74:sunt aut facere repellat provident occaecati excepturi optio reprehenderit6:userIdi1eee`
	ds := &DecodeState{data: data}
	fmt.Println(ds.Decode())

}
