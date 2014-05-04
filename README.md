* assert

assertion library for go test

** Example Test

```go
package test

import (
	"github.com/Jxck/assert"
	"testing"
)

func TestInt(t *testing.T) {
	type I int
	var actual I = 1
	expected := 2
	assert.Equal(t, actual, expected)
}

func TestString(t *testing.T) {
	type S string
	var actual S = "aaa"
	expected := "bbb"
	assert.Equal(t, actual, expected)
}

func TestBool(t *testing.T) {
	type B bool
	var actual B = false
	expected := true
	assert.Equal(t, actual, expected)
}

func TestSlice(t *testing.T) {
	type S []int
	var actual S = []int{1, 2, 3}
	expected := []int{1, 2, 3, 4}
	assert.Equal(t, actual, expected)
}

func TestStruct(t *testing.T) {
	type Foo struct {
		name  string
		age   int
		veget bool
		lang  []string
	}

	actual := Foo{"john", 20, true, []string{"ja", "en"}}
	expected := Foo{"emily", 22, false, []string{"ja", "en", "ch"}}

	assert.Equal(t, actual, expected)
}

func TestNestedStruct(t *testing.T) {
	type Foo struct {
		name  string
		age   int
		veget bool
		lang  []string
	}

	john := Foo{"john", 20, true, []string{"ja", "en"}}
	emily := Foo{"emily", 22, false, []string{"ja", "en", "ch"}}

	type Bar struct {
		Foo
		class int
	}

	actual := Bar{john, 1}
	expected := Bar{emily, 1}

	assert.Equal(t, actual, expected)
}
```


** Example Out

```
--- FAIL: TestInt (0.00 seconds)
	assert.go:94: 
		assert_test.go:12
		[actual]  :1(test.I)
		[expected]:2(int)
		
--- FAIL: TestString (0.00 seconds)
	assert.go:94: 
		assert_test.go:19
		[actual]  :"aaa"(test.S)
		[expected]:"bbb"(string)
		
--- FAIL: TestBool (0.00 seconds)
	assert.go:94: 
		assert_test.go:26
		[actual]  :false(test.B)
		[expected]:true(bool)
		
--- FAIL: TestSlice (0.00 seconds)
	assert.go:94: 
		assert_test.go:33
		[actual]  :[1(int), 2(int), 3(int), ](test.S[3])
		[expected]:[1(int), 2(int), 3(int), 4(int), ]([]int[4])
		
--- FAIL: TestStruct (0.00 seconds)
	assert.go:94: 
		assert_test.go:47
		[actual]  :
		{
			name:	"john"(string)
			age:	20(int)
			veget:	true(bool)
			lang:	["ja"(string), "en"(string), ]([]string[2])
		}
		
		[expected]:
		{
			name:	"emily"(string)
			age:	22(int)
			veget:	false(bool)
			lang:	["ja"(string), "en"(string), "ch"(string), ]([]string[3])
		}
		
		
--- FAIL: TestNestedStruct (0.00 seconds)
	assert.go:94: 
		assert_test.go:69
		[actual]  :
		{
			Foo:	
			{
				name:	"john"(string)
				age:	20(int)
				veget:	true(bool)
				lang:	["ja"(string), "en"(string), ]([]string[2])
			}
			
			class:	1(int)
		}
		
		[expected]:
		{
			Foo:	
			{
				name:	"emily"(string)
				age:	22(int)
				veget:	false(bool)
				lang:	["ja"(string), "en"(string), "ch"(string), ]([]string[3])
			}
			
			class:	1(int)
		}
		
		
FAIL
exit status 1
FAIL	_/assert/test	0.020s
```
