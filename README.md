
## Repo contains the examples and knowledge I gained from [W3Schools](https://www.w3schools.com/go/index.php) Go Tutorial


# What is Go?

- Go is a cross-platform, open-source programming language
- Go can be used to create high-performance applications
- Go is a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language
- Go was developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson in 2007
- Go's syntax is similar to C++

| Go | Python | C++ |
| --- | --- | --- |
| Statically typed | Dynamically typed | Statically typed |
| Fast run time | Slow run time | Fast run time |
| Compiled | Interpreted | Compiled |
| Fast compile time | Interpreted | Slow compile time |
| Supports concurrency through goroutines and channel | No built-in concurrency mechanism | Supports concurrency through threads |
| It has automatic garbage collection | It has automatic garbage collection | It does not have automatic garbage collection |
| It does not support classes and objects | Has classes and objects | Has classes and objects |
| Does not support inheritance | Supports inheritance | Supports inheritance |

# Variables

Go variable naming rules:

- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (`a-z, A-Z`, `0-9`, and `_` )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords

| var | := |
| --- | --- |
| Can be used inside and outside of functions | It can only be used inside functions |
| Variable declaration and value assignment can be done separately | Variable declaration and value assignment cannot be done separately (must be done in the same line) |

You can use the `const` keyword if a variable should have a fixed value that cannot be changed.

# Slices

Slices are similar to arrays, but are more powerful and flexible.

Like arrays, slices are also used to store multiple values of the same type in a single variable.

However, unlike arrays, the length of a slice can grow and shrink as you see fit.

In Go, there are several ways to create a slice:

- Using the []*datatype*{*values*} format
- Create a slice from an array
- Using the make() function

`len()` function - returns the length of the slice (the number of elements in the slice)

`cap()` function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

# Operators

1. `%=`: It represents the modulus assignment operator. It takes the value of the left-hand variable, calculates its modulus with the right-hand value, and assigns the result back to the left-hand variable. For example:

    ```go
    x %= 3 // Equivalent to: x = x % 3
    ```

2. `&=`: It represents the bitwise AND assignment operator. It performs a bitwise AND operation between the left-hand variable and the right-hand value, and assigns the result back to the left-hand variable. For example:

    ```go
    x &= 3 // Equivalent to: x = x & 3
    ```

3. `|=`: It represents the bitwise OR assignment operator. It performs a bitwise OR operation between the left-hand variable and the right-hand value and assigns the result back to the left-hand variable. For example:

    ```go
    x |= 3 // Equivalent to: x = x | 3
    ```

4. `^=`: It represents the bitwise XOR assignment operator. It performs a bitwise XOR (exclusive OR) operation between the left-hand variable and the right-hand value and assigns the result back to the left-hand variable. For example:

    ```go
    x ^= 3 // Equivalent to: x = x ^ 3
    ```

5. `>>=`: It represents the right shift assignment operator. It shifts the bits of the left-hand variable to the right by the number of positions specified by the right-hand value and assigns the result back to the left-hand variable. For example:

    ```go
    x >>= 3 // Equivalent to: x = x >> 3
    ```

6. `<<=`: It represents the left shift assignment operator. It shifts the bits of the left-hand variable to the left by the number of positions specified by the right-hand value and assigns the result back to the left-hand variable. For example:

    ```go
    x <<= 3 // Equivalent to: x = x << 3
    ```

# Loops

The `range` keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.

**Tip:** To only show the value or the index, you can omit the other output using an underscore (`_`).

# Functions

**Note:** When a **parameter** is passed to the function, it is called an **argument**. So, from the example above: `fname` is a **parameter**, while `Liam`, `Jenny` and `Anja` are **arguments**.
