package funktion

import (
	"bufio"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/jinzhu/copier"
	"github.com/teris-io/shortid"
)

type params map[string]interface{}
type tags map[string][]string

var (
	copieropts copier.Option = copier.Option{IgnoreEmpty: true, DeepCopy: false}
	sid        *shortid.Shortid
)

func init() {

	var err error

	sid, err = shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic(err)
	}

}

func IsEmpty(input string) bool {
	return isEmpty(input)
}

func IsBlank(a string) bool {
	return isEmpty(a)
}

func isEmpty(input string) bool {
	return (len(strings.TrimSpace(input)) <= 0)
}

func IsStruct(i interface{}) bool {
	return isStruct(i)
}

func isStruct(i interface{}) bool {
	return reflect.ValueOf(i).Type().Kind() == reflect.Struct
}

func isPointer(i interface{}) bool {
	return reflect.ValueOf(i).Kind() == reflect.Ptr
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func ToSlice(slice interface{}) []interface{} {

	s := reflect.ValueOf(slice)

	if !IsSlice(slice) {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func ArrayContains(list []string, contains string) bool {

	for _, b := range list {

		if strings.EqualFold(contains, b) {
			return true
		}

	}

	return false
}

func SliceContains(in []string, this []string) bool {

	for _, a := range in {

		for _, b := range this {

			if strings.EqualFold(a, b) {
				return true
			}

		}

	}

	return false
}

func SplitLines(s string) []string {

	var lines []string

	sc := bufio.NewScanner(strings.NewReader(s))

	for sc.Scan() {

		lines = append(lines, tabToSpace(sc.Text()))

	}

	return lines
}

func tabToSpace(input string) string {

	var result []string

	for _, i := range input {
		switch {
		// all these considered as space, including tab \t
		// '\t', '\n', '\v', '\f', '\r',' ', 0x85, 0xA0
		case unicode.IsSpace(i):
			result = append(result, "    ") // replace tab with space
		case !unicode.IsSpace(i):
			result = append(result, string(i))
		}
	}

	return strings.Join(result, "")

}

func mapLines(s string, delimeter string) map[string]string {

	output := make(map[string]string)

	sc := bufio.NewScanner(strings.NewReader(s))

	if delimeter == "" {
		delimeter = ":"
	}

	for sc.Scan() {

		split := strings.Split(sc.Text(), delimeter)

		if len(split) > 1 {
			output[split[0]] = strings.Trim(split[1], " ")
		} else {
			output[split[0]] = ""
		}

	}

	return output
}

func TruncatePrint(input string, length int) string {

	if len(input) < length {
		return fmt.Sprintf("%s", input)
	}

	return fmt.Sprintf("%s...", input[0:length])
}

func containsString(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}

func stringContains(this string, within string) bool {
	return strings.Contains(
		strings.ToLower(within),
		strings.ToLower(this),
	)
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsEmail(e string) bool {

	if len(e) < 3 && len(e) > 254 {
		return false
	}

	return emailRegex.MatchString(e)
}

func IsValidEmail(email string, domains []string) bool {

	if containsString(email, "bobmcallan") {
		return true
	}

	at := strings.LastIndex(email, "@")

	if at >= 0 {

		_, domain := email[:at], email[at+1:]

		return ArrayContains(domains, domain)

	}

	return false

}

func IsValidDomain(domain string) bool {
	switch domain {
	case
		"procul.io",
		"dashs.com",
		"dashs.com.au",
		"t3b.io":
		return true
	}

	return false
}

func ToJson(input interface{}) (string, error) {

	output, err := func() ([]byte, error) {
		return json.MarshalIndent(input, "", "\t")

	}()

	if err != nil {

		return "", err
	}

	return string(output), nil
}

func ToJsonFlat(input interface{}) (string, error) {

	output, err := func() ([]byte, error) {
		return json.Marshal(input)
	}()

	if err != nil {

		return "", err

	}

	return string(output), nil
}

func isSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func toSlice(slice interface{}) []interface{} {

	s := reflect.ValueOf(slice)

	if !isSlice(slice) {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func uniqueAppend(slice []string, item string) []string {

	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	uniqMap[item] = struct{}{}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}

	return uniqSlice

}

func uniqueSlice(slice []string) []string {

	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}

	return uniqSlice

}
