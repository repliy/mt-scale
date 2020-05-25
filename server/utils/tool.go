package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"log"
	"mt-scale/exception"
	"mt-scale/syslog"
	"os"
	"path/filepath"
	"strings"

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
		syslog.Error(err)
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

// MatchPath match url
func MatchPath(paths []string, url string) int {
	for i, path := range paths {
		if flag, _ := filepath.Match(path, url); flag {
			return i
		}
	}
	return -1
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

//GetFileSize get file size by path(B)
func GetFileSize(path string) int64 {
	if !exists(path) {
		return 0
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

//exists Whether the path exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Md5 md5 encrypt
func Md5(inputStr string) string {
	h := md5.New()
	h.Write([]byte(inputStr))
	str := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return str
}
