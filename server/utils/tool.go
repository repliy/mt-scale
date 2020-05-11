package utils

import (
	"bytes"
	"encoding/gob"
	"log"
	"mt-scale/exception"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// ValidateStructParams Validate params of struct
func ValidateStructParams(obj interface{}) {
	if validate == nil {
		validate = validator.New()
	}
	if err := validate.Struct(obj); err != nil {
		exception.ThrowBusinessErrorMsg(err.Error())
	}
}

// ValidateVarParams Validate var param
func ValidateVarParams(obj interface{}, reg string) {
	if validate == nil {
		validate = validator.New()
	}
	if err := validate.Var(obj, reg); err != nil {
		exception.ThrowBusinessErrorMsg(err.Error())
	}
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// ByteEncoder Encode to byte
func ByteEncoder(s interface{}) []byte {
	var encResult bytes.Buffer
	enc := gob.NewEncoder(&encResult)
	if err := enc.Encode(s); err != nil {
		log.Fatal("encode error:", err)
	}

	return encResult.Bytes()
}
