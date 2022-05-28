package gotest

func ExampleSayHello() {
	SayHello()
	// Output: hello World
}

func ExampleSayGoodBye() {
	SayGoodBye()
	// Output:
	// hello
	// good bye
}

func ExampleSayName() {
	SayName()
	// Unordered output:
	// 2::Jim
	// 3::Kitty
	// 4::Erik
	// 1::Tom
}
