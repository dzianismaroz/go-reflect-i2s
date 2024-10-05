#### solution result:
```
go test -v
=== RUN   TestSimple
--- PASS: TestSimple (0.00s)
=== RUN   TestComplex
--- PASS: TestComplex (0.00s)
=== RUN   TestSlice
--- PASS: TestSlice (0.00s)
=== RUN   TestErrors
--- PASS: TestErrors (0.00s)
PASS
ok      stepik/8/99_hw/i2s      0.001s
```
interface2struct function

i2s - interface to struct. A function that fills the values ​​of the structure from map[string]interface{} and similar - what you get if you unpack json into interface{} (see example in json/dynamic.go)

Reflection task.

Despite some sophistication at first glance, reflection is used very often. Understanding how it works and how you can work with it will be very useful in the future.

Implementation takes 80-100 lines of code

Of the data types, it is enough to provide those that are in the test.

Run go test -v

Write the code in the i2s.go file

Tips:



* All the functions you need are in the reflect package - https://golang.org/pkg/reflect/ - read the documentation carefully
* json unpacks int into float. This is stated in the documentation, not a bug. In this case, it will be correct to result in int if we encounter a float
* Always check what you receive at the entrance. And look at what you pass to the function (yes, recursion shows itself well here) not reflect.Value, but the original data that you got to through the necessary reflect methods
* If you use some structure names in a function that appear in the test, this is not correct
