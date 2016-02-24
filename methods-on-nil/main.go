package main

import "fmt"

type I interface {
	// returns a string
	SameString() string

	// returns a string that depends on the particular implementation of I
	MyString() string
}

type S struct {
	mystring string
}

func (s *S) SameString() string {
	return "same string"
}

func (s *S) MyString() string {
	return s.mystring
}

func test(i I) {
	defer catchPanic()

	fmt.Println("  same string:", i.SameString())
	fmt.Println("  my string:  ", i.MyString())
}

func main() {
	var i I
	s1 := S{"this is s1"}
	s2 := &S{"this is s2"}

	fmt.Println("i == nil?", i == nil) // i is unassigned so it's nil
	test(i)                            // as expected we can pass nil into I and it will fail
	fmt.Println()

	fmt.Println("s1:", s1)
	// *S is an implementation of I and has a concrete value. test works as expected
	test(&s1)
	fmt.Println()

	fmt.Println("s2:", s2)
	// Similarly s2 is a pointer to S already so we see the same behavior
	test(s2)
	fmt.Println()

	// the question then is what do you expect when you pass nil in as your I
	// implementation when it's bound from a *S?
	s2 = nil
	fmt.Println("s2:", s2)
	test(s2)
	fmt.Println()

	// ...it almost behaves as expected; we get a panic but it also partially
	// works. So why?
	//
	// Look at the test function which worked and observe that it does not
	// depend on any insternal state of the (nil) struct. Now recall how structs
	// are maintained basically as a <type, value> pair (cf. nil-interfaces
	// gotcha).
	//
	// Taken together it's possible to call a method on a nil *value* as long
	// as the *type* of nil is known. As an example:
	s2 = nil // (yes, I know it's already nil; just making a point)
	fmt.Printf("calling a function on %s: %s\n", s2, s2.SameString())
	fmt.Println()

	// now that we've established it's legit to call methods on nil structs
	// we need to watch out for cases where we get called on nil or
	s2.MyString() // we panic...

	// I think in practice this isn't a huge deal but it does mean it's
	// potentially dangerous to assume that your host struct will have any
	// meaning when writing some method on a pointer to the struct. I'm
	// unclear what best practices are here.
}

func catchPanic() {
	if e := recover(); e != nil {
		fmt.Println("    caught", e)
	}
}
