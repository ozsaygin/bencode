// Package bencode implements encoding and decoding of bencode
// format as defined in bittorrent spesification.
//
// See bencoding section in https://www.bittorrent.org/beps/bep_0003.html
// for the format specification.
package bencode

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Decoder is a type which store decoding state.
type Decoder struct {
	data []byte
	v    interface{}
}

// Encoder is a type which store encoding state.
type Encoder struct {
}

// NewDecoder generates and returns a new Decoder instance.
func NewDecoder(data []byte, v interface{}) *Decoder {
	return &Decoder{data, v}
}

// NewEncoder generates and returns a new Encoder instance.
func NewEncoder() *Encoder {
	return &Encoder{}
}

func (e *Encoder) marshal(v interface{}) []byte {

	switch v := reflect.ValueOf(v); v.Kind() {
	case reflect.String:
		word := v.Interface().(string)
		length := strconv.Itoa(len(word))
		encodedData := fmt.Sprintf("%s:%s", length, word)
		return []byte(encodedData)

	case reflect.Int:

		num := v.Interface().(int)
		encodedData := fmt.Sprintf("i%de", num)
		return []byte(encodedData)

	case reflect.Slice:

		if v.IsNil() {
			return []byte{}
		}

		data := []byte{'l'}
		for i := 0; i < v.Len(); i++ {
			elm := v.Index(i).Interface()
			data = append(data, e.marshal(elm)...)
		}

		data = append(data, 'e')
		return data

	case reflect.Map:

		if v.IsNil() {
			return []byte{}
		}

		data := []byte{'d'}
		iter := v.MapRange()
		for iter.Next() {
			data = append(data, e.marshal(iter.Key().Interface())...)
			data = append(data, e.marshal(iter.Value().Interface())...)
		}
		data = append(data, 'e')
		return data

	default:
		return []byte{}
	}
}

// Marshal encodes input to bencoded format.
// Only specific types (Map, Array, Integer, String) can be converted to bencoding format.
// It takes any object types mentioned above as parameter and returns bencoded data as byte array.
func Marshal(v interface{}) ([]byte, error) {
	e := NewEncoder()
	data := e.marshal(v)
	if len(data) == 0 {
		err := fmt.Sprintf("Illegal type for bencoding conversion: %s", reflect.TypeOf(v).Name())
		return nil, errors.New(err)
	}

	return data, nil
}

// // Unmarshal decodes bencoded string into Go object
// func Unmarshal(data []byte, v interface{}) error {

// 	d := NewDecoder(data, v)
// 	d.unmarshall()

// 	return nil
// }

// func (ds *Decoder) unmarshall() interface{} {

// 	prefix := ds.data[0]

// 	// parse data until nothing left to parse
// 	for len(ds.data) > 0 {

// 		switch prefix {

// 		case 'i':
// 			re := regexp.MustCompile(`i(\-?\d+)e`)
// 			matches := re.FindSubmatch(ds.data)
// 			value, err := strconv.Atoi(matches[1])

// 			if err != nil {
// 				fmt.Errorf("Corrupted bencode format")
// 				return nil
// 			}

// 			ds.data = bytes.TrimPrefix(ds.data, matches[0])
// 			return value

// 		case 'l':
// 			re := regexp.MustCompile(`l(.*)e`)
// 			matches := re.FindSubmatch(ds.data)
// 			elems := matches[1]
// 			lst := []interface{}{}

// 			localDecoder := &Decoder{data: elems}
// 			for localDecoder.data != "" {
// 				elm := localDecoder.unmarshall()
// 				lst = append(lst, elm)
// 			}

// 			ds.data = strings.TrimPrefix(ds.data, matches[0])
// 			return lst

// 		case 'd':
// 			// parse dictionary content
// 			re := regexp.MustCompile(`d(.*)e`)
// 			matches := re.FindSubmatch(ds.data)
// 			inner := matches[1]
// 			dict := make(map[string]interface{})

// 			// rest of data can be any format
// 			// recursively calling unmarshall handles the pairs
// 			localDecoder := &Decoder{data: inner}
// 			for localDecoder.data != "" {
// 				key := localDecoder.unmarshall().(string)
// 				value := localDecoder.unmarshall()
// 				dict[key] = value
// 			}

// 			// update data by trimming parsed part
// 			ds.data = bytes.TrimPrefix(ds.data, matches[0])
// 			return dict

// 		default:
// 			// default, assume data type is string
// 			re := regexp.MustCompile(`(\d+)\:(.+)`)
// 			matches := re.FindSubmatch(ds.data)

// 			ss := matches[1]
// 			length, _ := strconv.Atoi(ss)
// 			word := matches[2]
// 			word = word[:length]

// 			encodedStr := strconv.Itoa(length) + ":" + word
// 			ds.data = strings.TrimPrefix(ds.data, encodedStr)

// 			return word
// 		}
// 	}
// 	return nil
// }
