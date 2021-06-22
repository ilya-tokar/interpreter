# Лабараторные работы по СПО

Студент ИКБО-05-18 Токар И.И.

- [x] Лексер.
- [x] Парсер.
- [x] Интерпретатор.
- [x] Типы данных.

Creating variables.
```
let a = 10;
let b = 15
let name = "Monkey";
let result = 10 * (20 / 2);
let  result = ((2 + 3 * a) - 15) / 2
```
Creating arrays and hases.
```
let myArray = [1, 2, 3, 4, 5];
let thorsten = {"name": "Thorsten", "age": 28};
```
Accessing the elements in arrays and hashes is done with index expressions:
```
myArray[0]
// => 1

thorsten["name"] 
// => "Thorsten"
````
The let statements can also be used to bind functions to names. Here’s a small function that adds two numbers:
```
let add = fn(a, b) { return a + b; };
```
Calling a function:
```
add(1, 2);
```
A more complex function, such as a fibonacci function that returns the Nth Fibonacci number,
might look like this:
```
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};
```
