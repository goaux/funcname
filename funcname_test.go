package funcname_test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/goaux/funcname"
)

func TestOf(t *testing.T) {
	tests := []struct {
		name     string
		fn       interface{}
		expected string
	}{
		{"TestOf", TestOf, "github.com/goaux/funcname_test.TestOf"},
		{"fmt.Println", fmt.Println, "fmt.Println"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := funcname.Of(tt.fn)
			if result != tt.expected {
				t.Errorf("Of(%v) = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

func TestSplitOf(t *testing.T) {
	tests := []struct {
		name         string
		fn           interface{}
		expectedPkg  string
		expectedFunc string
	}{
		{"TestSplitOf", TestSplitOf, "github.com/goaux/funcname_test", "TestSplitOf"},
		{"fmt.Println", fmt.Println, "fmt", "Println"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg, fn := funcname.SplitOf(tt.fn)
			if pkg != tt.expectedPkg || fn != tt.expectedFunc {
				t.Errorf("SplitOf(%v) = (%v, %v), want (%v, %v)", tt.name, pkg, fn, tt.expectedPkg, tt.expectedFunc)
			}
		})
	}
}

func TestThis(t *testing.T) {
	result := funcname.This()
	expected := "github.com/goaux/funcname_test.TestThis"
	if result != expected {
		t.Errorf("This() = %v, want %v", result, expected)
	}
}

func TestSplitThis(t *testing.T) {
	pkg, fn := funcname.SplitThis()
	expectedPkg := "github.com/goaux/funcname_test"
	expectedFunc := "TestSplitThis"
	if pkg != expectedPkg || fn != expectedFunc {
		t.Errorf("SplitThis() = (%v, %v), want (%v, %v)", pkg, fn, expectedPkg, expectedFunc)
	}
}

func TestCaller(t *testing.T) {
	result := funcname.Caller(0)
	expected := "github.com/goaux/funcname_test.TestCaller"
	if result != expected {
		t.Errorf("Caller(0) = %v, want %v", result, expected)
	}
}

func TestSplitCaller(t *testing.T) {
	pkg, fn := funcname.SplitCaller(0)
	expectedPkg := "github.com/goaux/funcname_test"
	expectedFunc := "TestSplitCaller"
	if pkg != expectedPkg || fn != expectedFunc {
		t.Errorf("SplitCaller(0) = (%v, %v), want (%v, %v)", pkg, fn, expectedPkg, expectedFunc)
	}
}

func TestForPC(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	result := funcname.ForPC(pc)
	expected := "github.com/goaux/funcname_test.TestForPC"
	if result != expected {
		t.Errorf("ForPC() = %v, want %v", result, expected)
	}
}

func TestSplitForPC(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	pkg, fn := funcname.SplitForPC(pc)
	expectedPkg := "github.com/goaux/funcname_test"
	expectedFunc := "TestSplitForPC"
	if pkg != expectedPkg || fn != expectedFunc {
		t.Errorf("SplitForPC() = (%v, %v), want (%v, %v)", pkg, fn, expectedPkg, expectedFunc)
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input        string
		expectedPkg  string
		expectedFunc string
	}{
		{"github.com/user/pkg.Func", "github.com/user/pkg", "Func"},
		{"pkg.Func", "pkg", "Func"},
		{"pkg", "pkg", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			pkg, fn := funcname.Split(tt.input)
			if pkg != tt.expectedPkg || fn != tt.expectedFunc {
				t.Errorf("Split(%v) = (%v, %v), want (%v, %v)", tt.input, pkg, fn, tt.expectedPkg, tt.expectedFunc)
			}
		})
	}
}

func ExampleOf() {
	fmt.Println(funcname.Of(strings.ToUpper))
	// Output: strings.ToUpper
}

func ExampleSplit() {
	pkg, fn := funcname.Split("github.com/user/pkg.Func")
	fmt.Printf("Package: %s, Function: %s\n", pkg, fn)
	// Output: Package: github.com/user/pkg, Function: Func
}
