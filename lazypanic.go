package util

import (
	"log"
	"os"
	"strings"

	"github.com/go-errors/errors"
	errs "github.com/pkg/errors"
)

func getFileWithPrefix(filename, precontent string) (file *os.File) {
	if f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatalf("error in getFileForAppend - OpenFile: %v", err)
	} else {
		file = f
		file.WriteString(precontent)
	}
	return
}

// PanicHandle : simple calling for 'PanicHandleEx', p is recover()
func PanicHandle(p interface{}, logfile string) {
	PanicHandleEx(p, logfile, nil)
}

// PanicHandleEx hooks a function for dealing a panic,
// The 1st param of hook function is error-trace information.
func PanicHandleEx(p interface{}, logfile string, onPanic func(string, ...interface{}), params ...interface{}) {
	if p != nil {
		err := fSp(p)
		isNoFatal := TerOp(sCtn(err, NONFATALMARK), true, false).(bool)
		logfile = TerOp(logfile == "", defLog, logfile).(string)

		f := getFileWithPrefix(logfile, fSf("\n*** Panic Error *** Fatal : %t ***\n", isNoFatal))
		defer f.Close()
		log.SetOutput(f)
		log.Println(p)
		if onPanic != nil {
			onPanic(err, params...)
		}

		if !isNoFatal {
			f.Close()
			log.SetOutput(os.Stderr)
			log.Fatalln(p)
		}
	}
}

// LogOnError : write error (with stack track) to the log file when get the error
func LogOnError(logfile string, errs ...error) {
	for _, err := range errs {
		if err != nil {
			logfile = TerOp(logfile == "", defLog, logfile).(string)
			f := getFileWithPrefix(logfile, "\n*** Log Error ***\n")
			defer f.Close()
			log.SetOutput(f)
			errStackStr := errStack(err, 1, 2, 3, 4)
			log.Println(errStackStr)
		}
	}
}

// PanicOnError : launch a panic with error (stack track) when get the error
func PanicOnError(errs ...error) {
	for _, err := range errs {
		if err != nil {
			if emsg := errStack(err, 1, 2, 3, 4); len(emsg) != 0 {
				panic(emsg)
			}
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

// Must :
func Must(r interface{}, err error) interface{} {
	PanicOnError(err)
	return r
}

// Must2 :
func Must2(r1, r2 interface{}, err error) (_, _ interface{}) {
	PanicOnError(err)
	return r1, r2
}

// Must3 :
func Must3(r1, r2, r3 interface{}, err error) (_, _, _ interface{}) {
	PanicOnError(err)
	return r1, r2, r3
}

// Must4 :
func Must4(r1, r2, r3, r4 interface{}, err error) (_, _, _, _ interface{}) {
	PanicOnError(err)
	return r1, r2, r3, r4
}

// Must5 :
func Must5(r1, r2, r3, r4, r5 interface{}, err error) (_, _, _, _, _ interface{}) {
	PanicOnError(err)
	return r1, r2, r3, r4, r5
}

// Must6 :
func Must6(r1, r2, r3, r4, r5, r6 interface{}, err error) (_, _, _, _, _, _ interface{}) {
	PanicOnError(err)
	return r1, r2, r3, r4, r5, r6
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

	NEXT:
		for i, ln := range strings.FieldsFunc(e.ErrorStack(), func(c rune) bool {
			return c == '\n' || c == '\r'
		}) {
			for _, o := range omit {
				if i == o {
					continue NEXT
				}
			}
			lines = append(lines, ln)
		}
		return strings.Join(lines, "\n")
	}
	return ""
}
