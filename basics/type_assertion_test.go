package basics

import (
	logUtil "algo/main/common_ds"
	"fmt"
	"strconv"
	"testing"
)

func TestAssertion(t *testing.T) {
	var a interface{} = 1
	// we do type assertion to check the actual type of a value
	// val.(type) => there is two form
	// two value (comma-ok) form: doesn't result in panic, ok can be used to check if the assertion is successful
	// single value form will result in panic for type mismatch

	//two value form: ok idiom
	str, ok := a.(string)
	if ok {
		logUtil.Logf("str: %s", str)
	} else {
		logUtil.Logf("int %d", a.(int))
	}

	// single value form: panic idiom
	str = a.(string)             // panic as a is an int, not string
	logUtil.Logf("str: %s", str) // panic

}

func TestConversion(t *testing.T) {
	anInt := 24

	// float conversion
	aFloat := float64(anInt)
	fmt.Println(aFloat)

	// string conversion
	aStr := strconv.Itoa(anInt)
	logUtil.Logf("aStr: %s", aStr)

	// string to int
	_, error := strconv.Atoi("23A")
	if error != nil {
		logUtil.Logf("error: %s", error.Error())
	}

	// In go, type conversion doesn't work like other language
	// for example: nu
}
