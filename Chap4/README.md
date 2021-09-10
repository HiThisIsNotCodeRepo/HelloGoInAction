# Chap 4

*3 ways to manage multiple data in go*

1. Array
2. Slice
3. Map

*Array*

Array has fixed length, and each element in array is of the same type.

The access speed of array is very fast.

When declare array you have to define element quantity and type. Once declared an array the array type and length won't
change. If you need to store more element , you need to create a longer array.

*Literal value declaration*

```
array := [5]int{10,20,30,40,50}
```

If you use `...` replace array length. go will datetime the array length base on the quantity of initialised array
element.

```
array := [...]int{10,20,30,40,50}
```

In below initialisation the specific index element will be initialised by the value while other element will be zero
value.

```
array := [5]int{1:10,2:20}
```

*Use array*

Using `[]` operator to access the element, for example the index 2 element has been modified to 35

```
array[2] = 35
```

*Pointer array*

The below snippet use `int` pointer to init index 0 and 1 element.

```
array := [5]*int{0:new(int), 1:new(int)}
```

The below snippet assign value to index 0 and 1 element.

```
*array[0] = 10
*array[1] = 20
```

*Assign two same type array*

```
var array1 [5]string
arrasy2 := [5]string{"Red", "Blue", "green", "Yellow", "Pink"}
array1 = array2

```

If the array length is different the compiler will stop you assign value

```
var array1 [4]string
array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

// It won't work
array1 = array2

``` 

*Multi dimensional array*

```
var array [4][2]int

// Use literal value to init 2 dimensional array
array := [4][2]{{10,11},{20,21},{30,31},{40,41}}

// Declare index 1 an 3 element
array := [4][2]int{1: [20,21], 3:[40,41]}

// Delicare single element of multi dimension array
array := [4][2]int{1:{0:20}, 3:{1:41}}
```

The multidimensional array can be assigned to multidimensional array of same type

```
var array1 [2][2]int
var array2 [2][2]int

array2[0][0] = 10
array2[0][1] = 20
array2[1][0] = 30
array2[1][1] = 40
array1 = array2
```

Also, tha specific index, no matter 1 dimension or 2 dimension can be assigned.

```
var array3[2] = array1[1]

var value int = array1[1][0]
```

*Pass array between function*

When pass array by value it will consume much more space, because by default the value is passed by copy, therefore
recommend passing array by pointer (the pointer is consuming less space)

```
// Not recommend

var array [1e6]int

foo(array)
func foo(array [1e6int]){
...
}

// Recommend

var array [1e6]int

foo(&array)

func foo(array *[1e6]int){
...
}
```

*Slice*

Slice is another type of data structure. It's basically a dynamic array, and it will increase or decrease automatically
base on need.

The slice has:

1. Pointer to underlying array.
2. The length of slice.
3. The capacity of slice.

*Create a slice*

1. Using `make`, when use it, needs to pass an argument to indicate the slice length

```
// The length and capacity are the same
slice := make([]string,5)
// The length and capacity are different
slice := make([]int,3,5)
``` 

When specify length and capacity respectively, after initialisation could not access all capacity element.

*Use literal value to init a slice*

In the below snippet the length and capacity of a slice will be determined by the actual value.

```
slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
```

If you need to specify the length of slice recommend use index. In below snippet the slice's length and capacity are
both 100.

```
slice := []string{99:""}
```

*nil slice*

Use below syntax to create `nil` slice.

```
var slice []int
```

In a `nil` slice the len and cap are both 0. When a function expected to return a slice and a panic occur it will return
a `nil` slice.

*Empty slice*

Use below syntax to create an empty slice

```
slice := make([]int, 0)
slice := []int{}
```

For example when a database return 0 enquiry result can use empty slice.

*Use slice*

By index.

```
slice := []int{10,20,30,40,50}
slice[1]=25
```

Create a slice from a slice

```
slice := []int{10,20,30,40,50}
newSlice := slice[1:3]
```

Both `slice` and `newSlice` share the same underlying array, like below example:

```
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice[0] = 25
	fmt.Printf("slice:%v\n", slice)
	fmt.Printf("newSlice:%v\n", newSlice)
	/*
	slice:[10 25 30 40 50]
	newSlice:[25 30]
	*/
```

*Calculate length and cap of a slice*

If the length of underlying array is k , a `slice[i:j]`:

- Length: `j-i`
- Cap: `k-i`

*Use `append` to increase slice*

Compare with array ,slice's length is dynamic, use `append` can increase its length, it depends on the actual length of
the slice. Also, it will change the underlying array, any other slice share the same underlying array will see the
effect.

```
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Printf("slice:%v\n", slice)
	fmt.Printf("newSlice:%v\n", newSlice)
	/*
	slice:[10 20 30 60 50]
	newSlice:[20 30 60]
	*/
```

If the slice's underlying array's length is not sufficient the `append` will increase its length by creating a new array
and then append new value.

When slice's capacity is below 1000 every time when `append` apply will multiply its capacity by 2. Once capacity is
above 1000 it will be multiplied by 1.25.