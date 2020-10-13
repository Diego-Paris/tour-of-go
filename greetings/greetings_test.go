package greetings

import (
	"testing"
	"regexp"
)

//* To test, run: go test
//? The go test command executes test functions 
//? (whose names begin with Test) in test files 
//? (whose names end with _test.go). You can add 
//? the -v flag to get verbose output that lists 
//? all of the tests and their results.


//? TestHelloName calls the Hello function, 
//? passing a name value with which the function 
//? should be able to return a valid response message. 
//? If the call returns an error or an unexpected 
//? response message (one that doesn't include the name 
//? you passed in), you use the t parameter's Fatalf 
//? method to print a message to the console and end 
//? execution.

// TestHelloName calles greetings.Hello with a name,
// checking for a valid return value.
func TestHelloName(t *testing.T) {
	// Get an arbitrary name
	name := "Gladys"

	//! \b is a word boundary, letting it know we expect
	//! a sentence which includes the word
	want := regexp.MustCompile(`\b`+name+`\b`)
	
	// Execute Hello function with a name
	msg, err := Hello("Gladys")

	//! matchString is making sure that the response
	//! includes the name we passed in
	// if error is not null it means that something
	// is in its place when there shouldn't be
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

//? TestHelloEmpty calls the Hello function with 
//? an empty string. This test is designed to confirm 
//? that your error handling works. If the call 
//? returns a non-empty string or no error, you 
//? use the t parameter's Fatalf method to print 
//? a message to the console and end execution.

// TestHelloEmpty calls greetings.Hello with an empty string
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)	
	}
}