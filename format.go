// Licensed under the MIT license that can be found in the LICENSE file.

package xstr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	"unicode/utf8"
)

// ExpandTabs can expand tabs ('\t') rune in str to one or more spaces dpending on
// current column and tabSize.
// The column number is reset to zero after each newline ('\n') occurring in the str.
//
// ExpandTabs uses RuneWidth to decide rune's width.
// For example, CJK characters will be treated as two characters.
//
// If tabSize <= 0, ExpandTabs panics with error.
//
// Samples:
//     ExpandTabs("a\tbc\tdef\tghij\tk", 4) => "a   bc  def ghij    k"
//     ExpandTabs("abcdefg\thij\nk\tl", 4)  => "abcdefg hij\nk   l"
//     ExpandTabs("z中\t文\tw", 4)           => "z中 文  w"
func ExpandTabs(str string, tabSize int) string {
	if tabSize <= 0 {
		panic("tab size must be positive")
	}

	var r rune
	var i, size, column, expand int
	var output *bytes.Buffer

	orig := str

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)

		if r == '\t' {
			expand = tabSize - column%tabSize

			if output == nil {
				output = allocBuffer(orig, str)
			}

			for i = 0; i < expand; i++ {
				output.WriteByte(byte(' '))
			}

			column += expand
		} else {
			if r == '\n' {
				column = 0
			} else {
				column += RuneWidth(r)
			}

			if output != nil {
				output.WriteRune(r)
			}
		}

		str = str[size:]
	}

	if output == nil {
		return orig
	}

	return output.String()
}

// LeftJustify returns a string with pad string at right side if str's rune length is smaller than length.
// If str's rune length is larger than length, str itself will be returned.
//
// If pad is an empty string, str will be returned.
//
// Samples:
//     LeftJustify("hello", 4, " ")    => "hello"
//     LeftJustify("hello", 10, " ")   => "hello     "
//     LeftJustify("hello", 10, "123") => "hello12312"
func LeftJustify(str string, length int, pad string) string {
	l := Len(str)

	if l >= length || pad == "" {
		return str
	}

	remains := length - l
	padLen := Len(pad)

	output := &bytes.Buffer{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	output.WriteString(str)
	writePadString(output, pad, padLen, remains)
	return output.String()
}

// RightJustify returns a string with pad string at left side if str's rune length is smaller than length.
// If str's rune length is larger than length, str itself will be returned.
//
// If pad is an empty string, str will be returned.
//
// Samples:
//     RightJustify("hello", 4, " ")    => "hello"
//     RightJustify("hello", 10, " ")   => "     hello"
//     RightJustify("hello", 10, "123") => "12312hello"
func RightJustify(str string, length int, pad string) string {
	l := Len(str)

	if l >= length || pad == "" {
		return str
	}

	remains := length - l
	padLen := Len(pad)

	output := &bytes.Buffer{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	writePadString(output, pad, padLen, remains)
	output.WriteString(str)
	return output.String()
}

// Center returns a string with pad string at both side if str's rune length is smaller than length.
// If str's rune length is larger than length, str itself will be returned.
//
// If pad is an empty string, str will be returned.
//
// Samples:
//     Center("hello", 4, " ")    => "hello"
//     Center("hello", 10, " ")   => "  hello   "
//     Center("hello", 10, "123") => "12hello123"
func Center(str string, length int, pad string) string {
	l := Len(str)

	if l >= length || pad == "" {
		return str
	}

	remains := length - l
	padLen := Len(pad)

	output := &bytes.Buffer{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	writePadString(output, pad, padLen, remains/2)
	output.WriteString(str)
	writePadString(output, pad, padLen, (remains+1)/2)
	return output.String()
}

// PrettyJSON convert a json-str to a json-str that humans can read easily
// If str is not a valid json-str,will return original str and no-nil error
func PrettyJSON(str string, indent int) (string, error) {
	fmt.Println(indent)
	var dat map[string]interface{}
	var buf bytes.Buffer
	err := json.Unmarshal([]byte(str), &dat)
	if err != nil {
		return str, err
	}
	prettify(reflect.ValueOf(dat), 0, &buf)
	return buf.String(), nil
}

// MustPrettyJSON convert a json-str to a json-str that humans can read easily
// If str is not a valid json-str,will return original str only
func MustPrettyJSON(str string, indent int) string {
	str, _ = PrettyJSON(str, indent)
	return str
}

func writePadString(output *bytes.Buffer, pad string, padLen, remains int) {
	var r rune
	var size int

	repeats := remains / padLen

	for i := 0; i < repeats; i++ {
		output.WriteString(pad)
	}

	remains = remains % padLen

	if remains != 0 {
		for i := 0; i < remains; i++ {
			r, size = utf8.DecodeRuneInString(pad)
			output.WriteRune(r)
			pad = pad[size:]
		}
	}
}

// prettify will recursively walk value v to build a textual
// representation of the value.
func prettify(v reflect.Value, indent int, buf *bytes.Buffer) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		strtype := v.Type().String()
		if strtype == "time.Time" {
			fmt.Fprintf(buf, "%s", v.Interface())
			break
		} else if strings.HasPrefix(strtype, "io.") {
			buf.WriteString("<buffer>")
			break
		}

		buf.WriteString("{\n")

		names := []string{}
		for i := 0; i < v.Type().NumField(); i++ {
			name := v.Type().Field(i).Name
			f := v.Field(i)
			if name[0:1] == strings.ToLower(name[0:1]) {
				continue // ignore unexported fields
			}
			if (f.Kind() == reflect.Ptr || f.Kind() == reflect.Slice || f.Kind() == reflect.Map) && f.IsNil() {
				continue // ignore unset fields
			}
			names = append(names, name)
		}

		for i, n := range names {
			val := v.FieldByName(n)
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(n + ": ")
			prettify(val, indent+2, buf)

			if i < len(names)-1 {
				buf.WriteString(",\n")
			}
		}

		buf.WriteString("\n" + strings.Repeat(" ", indent) + "}")
	case reflect.Slice:
		strtype := v.Type().String()
		if strtype == "[]uint8" {
			fmt.Fprintf(buf, "<binary> len %d", v.Len())
			break
		}

		nl, id, id2 := "", "", ""
		if v.Len() > 3 {
			nl, id, id2 = "\n", strings.Repeat(" ", indent), strings.Repeat(" ", indent+2)
		}
		buf.WriteString("[" + nl)
		for i := 0; i < v.Len(); i++ {
			buf.WriteString(id2)
			prettify(v.Index(i), indent+2, buf)

			if i < v.Len()-1 {
				buf.WriteString("," + nl)
			}
		}

		buf.WriteString(nl + id + "]")
	case reflect.Map:
		buf.WriteString("{\n")

		for i, k := range v.MapKeys() {
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(k.String() + ": ")
			prettify(v.MapIndex(k), indent+2, buf)

			if i < v.Len()-1 {
				buf.WriteString(",\n")
			}
		}

		buf.WriteString("\n" + strings.Repeat(" ", indent) + "}")
	default:
		if !v.IsValid() {
			fmt.Fprint(buf, "<invalid value>")
			return
		}
		format := "%v"
		switch v.Interface().(type) {
		case string:
			format = "%q"
		case io.ReadSeeker, io.Reader:
			format = "buffer(%p)"
		}
		fmt.Fprintf(buf, format, v.Interface())
	}
}
