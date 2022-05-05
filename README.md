# type-safety
Lets demonstrate what [type safety](https://en.wikipedia.org/wiki/Type_safety) is.

This is a [Python code](./python/test.py). Not type safe. Python will try its best to run.

The Python interpreter will run the code until it hits an error. 

The first 3 calls to the `add_two_things` function will be executed.

```python
def add_two_things(a, b):
    c = a + b
    print(c)
    print(type(c))

add_two_things(1, 2) # int + int: expect int 3
add_two_things(1, 2.3) # int + float: expect float 3.3
add_two_things("a", "b") # str + str: expect string 'ab'

add_two_things(1, "2") # expect TypeError: unsupported operand type(s) for +: 'int' and 'str'
```

This is a [Go code](./go/main.go). Type safe. 

The Go compiler won't compile the code.

None of the calls to the `addTwoThings` function will be executed.

```Go
package main

import "fmt"

// addTwoThings expects 2 int type arguments, and add them.
func addTwoThings(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	addTwoThings(1, 2)

	// following calls will prevent program to be compiled.
	addTwoThings(1, 2.3)
	addTwoThings("a", "b")
	addTwoThings(1, "2")
}
```
