// Package funcname provides utility functions for retrieving and manipulating function names.
// It offers various methods to obtain function names from different contexts, such as
// runtime program counters, interface values, and the current call stack.
package funcname

import (
	"reflect"
	"runtime"
	"strings"
)

const (
	// Unknown is the full name returned when a function name cannot be determined.
	// It is a combination of UnknownPackage and UnknownFunction.
	Unknown = UnknownPackage + "." + UnknownFunction

	// UnknownPackage is the package name returned when a package cannot be determined.
	UnknownPackage = "(unknown)"

	// UnknownFunction is the function name returned when a function cannot be determined.
	UnknownFunction = "(Unknown)"
)

// Of returns the full name of the function or method represented by v.
// It wraps ForPC(pc) where pc is taken from reflect.ValueOf(v).Pointer().
// The parameter v must be a function or a method. If v is not a function or method, it returns Unknown.
func Of(v interface{}) string {
	rv := reflect.ValueOf(v)
	return ForPC(rv.Pointer())
}

// SplitOf returns the package name and function name of the function or method represented by v.
// It wraps Split(Of(v)). If v is not a function or method, it returns (UnknownPackage, UnknownFunction).
func SplitOf(v interface{}) (pkgname, name string) {
	return Split(Of(v))
}

// This returns the name of the calling function.
// It is equivalent to Caller(1).
func This() string {
	return Caller(1)
}

// SplitThis returns the package name and function name of the calling function.
// It is equivalent to Split(Caller(1)).
func SplitThis() (pkgname, name string) {
	return Split(Caller(1))
}

// Caller returns the name of the calling function, skipping the specified number of stack frames.
// It wraps ForPC(pc) where pc is taken from runtime.Caller(skip + 1).
// If runtime.Caller fails to retrieve the program counter, it returns Unknown.
func Caller(skip int) string {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return Unknown
	}
	return ForPC(pc)
}

// SplitCaller returns the package name and function name of the calling function,
// skipping the specified number of stack frames.
// It is equivalent to Split(Caller(skip + 1)).
func SplitCaller(skip int) (pkgname, name string) {
	return Split(Caller(skip + 1))
}

// ForPC returns the name of the function corresponding to the given program counter.
// It wraps runtime.FuncForPC(pc).Name(). If the function cannot be found, it returns Unknown.
func ForPC(pc uintptr) string {
	f := runtime.FuncForPC(pc)
	if f == nil {
		return Unknown
	}
	return f.Name()
}

// SplitForPC returns the package name and function name corresponding to the given program counter.
// It is equivalent to Split(ForPC(pc)).
func SplitForPC(pc uintptr) (pkgname, name string) {
	return Split(ForPC(pc))
}

// Split separates a fully qualified function name into package name and function name.
// It handles three main cases:
// 1. Full path with package and function: "github.com/user/pkg.Func" -> ("github.com/user/pkg", "Func")
// 2. Package and function without full path: "pkg.Func" -> ("pkg", "Func")
// 3. Only package name: "pkg" -> ("pkg", "")
//
// The function works by finding the last '/' character (if any) to separate the path,
// then finding the first '.' character after that to separate the package and function names.
// If there's no '.' character, it assumes the entire string is the package name.
func Split(s string) (pkgname, name string) {
	i := strings.LastIndexByte(s, '/') + 1
	j := strings.IndexByte(s[i:], '.') + 1
	if j == 0 {
		return s, ""
	}
	k := i + j
	return s[:k-1], s[k:]
}
