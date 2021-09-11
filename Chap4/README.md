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

*Slice with 3 indexes*

`slice[i:j:k]` the length will be `j-i` capacity will be `k-i`. `k` isn't allowed to be larger than underlying array max
index.

With 3 indexes we can create a slice that length and capacity are equal. Therefore, when do `append` because it exceeds
the underlying slice capacity a new array will be generated. So the new value won't change the original slice.

```
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice := source[2:3:3]
	slice = append(slice, "KiWi")
	fmt.Printf("source:%v\n", source)
	fmt.Printf("slice:%v\n", slice)
	/*
	source:[Apple Orange Plum Banana Grape]
	slice:[Plum KiWi]
	*/

```

*Using `...` in `append`*

By using `...` in `append` we can add two slices together.

```
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...))
	// [1 2 3 4]
```

*Iterate slice*

Using `for range` to iterate a slice.

```
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
	/*
	Index: 0 Value: 10
	Index: 1 Value: 20
	Index: 2 Value: 30
	Index: 3 Value: 40
	*/
```

`range` will return the index and copy value of the slice element at that index. Every iteration will assign the value
of the slice to that copy value.

```
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}
	/*
	Value: 10 Value-Addr: C000016178 ElemAddr: C00000E2A0
	Value: 20 Value-Addr: C000016178 ElemAddr: C00000E2A8
	Value: 30 Value-Addr: C000016178 ElemAddr: C00000E2B0
	Value: 40 Value-Addr: C000016178 ElemAddr: C00000E2B8
	*/
```

*Build in `len` and `cap`*

Both `len` and `cap` can be used on:

1. Array.
2. Slice.
3. Channel.

*Multi dimensional slice*

```
slice := [][]int{{10},{100,200}}
```

*Passing slice between function*

On 64bit arch machine , one slice takes up 24 bytes , 8 bytes for pointer, len and cap are 8 bytes respectively.

*Size comparison between array and slice*

```
	intArr := [4]int{10, 20, 30, 40}
	intSlice := []int{10, 20, 30, 40}
	strArr := [3]string{"a", "b", "c"}
	strSlice := []string{"a", "b", "c"}
	fmt.Printf("intArr size:%v\n", unsafe.Sizeof(intArr))
	fmt.Printf("intSlice size:%v\n", unsafe.Sizeof(intSlice))
	fmt.Printf("strArr size:%v\n", unsafe.Sizeof(strArr))
	fmt.Printf("strSlice size:%v\n", unsafe.Sizeof(strSlice))
	/*
	intArr size:32
	intSlice size:24
	strArr size:48
	strSlice size:24
	*/
```

*Map*

Map is a data structure , it can be used to store a series of unordered key value pair.

*How map implement?*

Use hash function to generate hash value of the key. Use the lower bit of that hash value to locate the bucket, inside
the bucket there're two arrays. One array is used to store higher bit of hash value it used to confirm that key/value
exist in the bucket. The second array is to store key and value.

*Map initialisation*

```
dict := make(map[string]int)
dict := map[string]string{"Red":"#da1337","Orange":"#e95a22"}
```

*Key of map*
The key could be built in type or struct. Slice, function and struct include slice can not.

```
	// ok
	_ = map[[]string]int{}
	// error
	_ = map[int][]string{}
```

*Assign map*

```
colors := map[string]string{}
colors["Red"] = "#da1337"
```

*nil map*

When assign a value to nil map, an error will occur.

```
	var colors map[string]string
	colors["Red"] = "#da1337"
	//panic: assignment to entry in nil map
```

*Retrieve value from map*

```
value, exists := colors["Blue]
if exists {
    fmt.Println(value)
}
```

```
value := colors["Blue"]
if value != ""{
    fmt.Println(value)
}
```

*Iterate map*

```
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
```

*Delete key value from map*

```
delete(colors,"Coral")
```

*Passing map between function*
Because map is reference type when pass map between function any modification to the map will be updated to all reference to this map.

