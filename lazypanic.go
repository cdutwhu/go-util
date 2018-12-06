package util

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/go-errors/errors"
	errs "github.com/pkg/errors"
)

//
var defLog = path.Join(os.TempDir(), "/GoErrorLog.txt")

func getFileWithPrefix(filename, precontent string) (file *os.File) {
	if f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatalf("error in getFileForAppend - OpenFile: %v", err)
	} else {
		file = f
		file.WriteString(precontent)
	}
	return
}

// PanicHandle : simple calling for 'PanicHandleEx'
func PanicHandle(p interface{}, logfile string, isFatal bool) {
	PanicHandleEx(p, logfile, isFatal, nil, "")
}

// PanicHandleEx hooks a function for dealing a panic,
// The 1st param of hook function is error-trace information.
func PanicHandleEx(p interface{}, logfile string, isFatal bool, onPanic func(string, ...interface{}), params ...interface{}) {
	if p != nil {
		f := getFileWithPrefix(Str(logfile).DefValue(defLog), fmt.Sprintf("\n*** Panic Error *** Fatal : %t ***\n", isFatal))
		defer f.Close()
		log.SetOutput(f)
		log.Println(p)
		if onPanic != nil {
			onPanic(fmt.Sprint(p), params...)
		}
		if isFatal {
			f.Close()
			log.SetOutput(os.Stderr)
			log.Fatalln(p)
		}
	}
}

// LogOnError : write error (with stack track) to the log file when get the error
func LogOnError(err error, logfile string) {
	if err != nil {
		f := getFileWithPrefix(Str(logfile).DefValue(defLog), "\n*** Log Error ***\n")
		defer f.Close()
		log.SetOutput(f)
		errStackStr := errStack(err, 1, 2, 3, 4)
		log.Println(errStackStr)
	}
}

// PanicOnError : launch a panic with error (stack track) when get the error
func PanicOnError(err error) {
	if err != nil {
		if emsg := errStack(err, 1, 2, 3, 4); len(emsg) != 0 {
			panic(emsg)
		}
	}
}

// PanicOnError1 : launch a panic with error (stack track) and one description when get the error
func PanicOnError1(err error, estr string) {
	if err != nil {
		if emsg := errStack(errs.Wrap(err, estr), 1, 2, 3, 4); len(emsg) != 0 {
			panic(emsg)
		}
	}
}

// PanicOnCondition : launch a panic when the error condition comes true, input the error condition's error
func PanicOnCondition(errCondition bool, err error) {
	if errCondition {
		if emsg := errStack(err, 1, 2, 3, 4); len(emsg) != 0 {
			panic(emsg)
		}
	}
}

// errStack : get the error track stack
func errStack(err error, omit ...int) string {
	if err != nil {
		e := errors.New(err)
		lines := []string{}
		for i, l := range strings.FieldsFunc(e.ErrorStack(), func(c rune) bool {
			return C32(c).InArr('\n', '\r')
		}) {
			if !I32(i).InArr(omit...) {
				lines = append(lines, l)
			}
		}
		return strings.Join(lines, "\n")
	}
	return ""
}
