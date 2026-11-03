# Go Path — From first line to full stack

A self-paced curriculum. Read, experiment, test, and build. Guided reading estimates exclude independent project work.

Recommended order: Your first Go program → Work with real data → Design with types → Build useful Go programs → Make correctness visible → Concurrency with a clear lifetime → Build a Go web API → Add durable storage → Connect the whole experience → From a project to a service.

Complete each exercise, pass its checkpoint, and review the module project before moving on. Revisit earlier modules whenever a concept feels uncertain.

# 1. Your first Go program

A friendly start: tools, values, and decisions

Level: Beginner. Prerequisite: None.

## Set up and say hello

25 minutes of guided study, plus practice.

Learning objectives:

- Install Go and inspect its version
- Explain packages, imports, and main
- Run and build a module

### Start with a small feedback loop

Install the current stable Go release from go.dev/dl and confirm that go version works in a fresh terminal. Your editor is a place to write text; the Go toolchain turns that text into a running program. Create a folder, enter it, and run go mod init example.com/hello. A module is a collection of packages with a shared dependency definition in go.mod. The example.com path is a learning placeholder, not a requirement to own a website.

### Read the program from the outside in

Every Go file starts with its package declaration. An executable uses package main and a main function with no arguments or return values. import makes another package available; fmt contains text-formatting tools. Braces enclose a function body. Save the sample as main.go, then use go run . from the module folder. You should see Hello, Go! once.

### Run, format, build

Use gofmt -w main.go to apply standard formatting. go run . compiles and runs during development; go build . creates an executable you can distribute for the target operating system and architecture. Compiler errors are useful directions: the filename and line number tell you where to investigate. Change one thing, rerun, and compare your prediction with what happened.

### Example

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

### Exercise

Change the program to print your name and one reason you want to learn Go on separate lines. Build it and run the executable.

Hint: Call fmt.Println twice inside main. Keep strings inside double quotes.

Review criteria:

- [ ] go run . prints exactly two lines
- [ ] gofmt -l . reports no unformatted Go files
- [ ] go build . succeeds

### Worked solution

```
package main

import "fmt"

func main() {
    fmt.Println("Hi, I am Alex.")
    fmt.Println("I want to build useful web services.")
}
```

### Checkpoint

What does go mod init create?

1. A module definition in go.mod
2. A running web server
3. An editor configuration

Answer: 1. go mod init creates the module definition. It does not write your application or start a server.

Sources:

- [Install Go](https://go.dev/doc/install)
- [Getting started](https://go.dev/doc/tutorial/getting-started)

## Values, variables, and zero values

30 minutes of guided study, plus practice.

Learning objectives:

- Choose basic numeric, string, and boolean types
- Distinguish declaration from assignment
- Predict zero values and explicit conversions

### Give information a name

A variable is a named storage location with a type. var count int declares an integer initialized to zero. Inside a function, count := 3 declares a variable with an inferred type; count = 4 updates it. The short form needs at least one new variable on its left side in the same scope. Go rejects unused local variables so that forgotten work does not silently accumulate.

### Make types visible in your thinking

Use int for ordinary counts, float64 for approximate decimal calculations, bool for yes-or-no decisions, and string for UTF-8 text. A string is an immutable sequence of bytes: len counts bytes, not necessarily user-perceived characters. Numeric conversions are explicit. float64(3) / 2 gives 1.5, while 3 / 2 is integer division and gives 1. Use integer cents rather than floating-point dollars for simple money calculations.

### Trust initialization, watch scope

Without an explicit initializer, numbers start at zero, booleans at false, and strings at the empty string. This makes many types usable immediately. A variable declared inside a block belongs to that block. Reusing a name with := in a nested block may shadow an outer variable instead of updating it. Prefer clear names and small scopes so that each value is easy to follow.

### Example

```
package main

import "fmt"

func main() {
    const priceCents = 1299
    quantity := 2
    totalCents := priceCents * quantity
    var paid bool
    fmt.Printf("total: %d cents; paid: %t\n", totalCents, paid)
}
```

### Exercise

Calculate the average of 7 and 10 as a float64. Also print the zero value of a string between quotation marks.

Hint: Convert before dividing; use the %q formatting verb to make an empty string visible.

Review criteria:

- [ ] Average prints 8.5
- [ ] Empty string prints as ""
- [ ] Explain why float64((7+10)/2) gives a different result

### Worked solution

```
package main
import "fmt"
func main() {
    average := float64(7 + 10) / 2
    var name string
    fmt.Printf("average: %.1f; name: %q\n", average, name)
}
```

### Checkpoint

What is the value of var ready bool?

1. true
2. false
3. nil

Answer: 2. The zero value of bool is false. nil is for types such as pointers, slices, maps, channels, interfaces, and functions.

Sources:

- [Go language specification](https://go.dev/ref/spec)
- [fmt package](https://pkg.go.dev/fmt)

## Decisions, loops, and readable flow

35 minutes of guided study, plus practice.

Learning objectives:

- Use if and switch for branching
- Write counted and conditional for loops
- Trace a loop without guessing

### Turn rules into branches

Use if for a condition and else for the alternative. Go conditions must be booleans; an integer cannot stand in for true. A switch compares one value with several cases. Cases do not fall through by default. When several answers are possible, put the most specific condition before a broader one so the broader rule does not hide it.

### One loop, several shapes

Go uses for for counted loops, condition-only loops, and indefinite loops. A counted loop has an initialization, condition, and post statement. A range loop walks a collection and is introduced next. break exits the nearest loop or switch; continue advances to the next iteration. Before running a loop, write down its first two iterations and its stopping condition. That habit catches off-by-one mistakes.

### Prefer a visible happy path

Nested branches become hard to read when every line is indented. In functions that return a result, reject invalid inputs early and let the normal path continue. Do not compress several rules into one clever expression. A future teammate should be able to connect each condition to a business rule. Practice with boundary values: zero, one, the threshold, and one above it.

### Example

```
package main
import "fmt"
func main() {
    for n := 1; n <= 5; n++ {
        if n%2 == 0 {
            fmt.Println(n, "even")
        } else {
            fmt.Println(n, "odd")
        }
    }
}
```

### Exercise

Print numbers 1 through 20. Replace multiples of 3 with Fizz, multiples of 5 with Buzz, and multiples of both with FizzBuzz.

Hint: Check divisibility by 15 before checking 3 or 5.

Review criteria:

- [ ] Exactly 20 lines print
- [ ] Line 15 is FizzBuzz
- [ ] Lines 1, 3, and 5 are 1, Fizz, and Buzz

### Worked solution

```
package main
import "fmt"
func main() {
    for n := 1; n <= 20; n++ {
        switch {
        case n%15 == 0: fmt.Println("FizzBuzz")
        case n%3 == 0: fmt.Println("Fizz")
        case n%5 == 0: fmt.Println("Buzz")
        default: fmt.Println(n)
        }
    }
}
```

### Checkpoint

Do Go switch cases fall through automatically?

1. Yes, until break
2. No, a matched case normally ends the switch
3. Only for integers

Answer: 2. Go switch cases do not fall through automatically. Explicit fallthrough exists, but most application code does not need it.

Sources:

- [Go language specification](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)

## Project: Build a command-line receipt

Create a receipt calculator that prints item names, quantities, subtotal, a fixed shipping cost, and the final total in cents. Start with hard-coded items; input parsing comes later.

Requirements:

- Use named constants for prices
- Apply free shipping at an explicit threshold
- Format money from integer cents
- Document three sample orders

Self-review rubric:

- [ ] Correct totals including zero items (40%)
- [ ] Readable names and small blocks (30%)
- [ ] Boundary examples and formatting (30%)

# 2. Work with real data

Functions, collections, and text

Level: Beginner. Prerequisite: Your first Go program.

## Functions with useful contracts

35 minutes of guided study, plus practice.

Learning objectives:

- Define parameters and return values
- Return more than one value
- Separate calculation from printing

### A function makes a promise

A function signature says what inputs it takes and what outputs it returns. Names should describe a job: subtotal is clearer than doStuff. Parameters are passed by value. A function that adds two integers receives copies of those integers. Keep calculations separate from printing so you can reuse them in a terminal program, a test, and eventually a web handler.

### Say what failure looks like

Go functions can return multiple values. The usual pattern for recoverable failure is a useful result plus an error; we will build that pattern carefully in the types module. For now, a value and boolean can express a successful lookup or a failed precondition. Make every path return values of the declared types. Avoid using a magic number such as -999 to signal failure.

### Test a contract at its edges

A function should handle the inputs its documentation promises. Is quantity zero allowed? Are negative amounts rejected? Write a sentence answering each question before the implementation. Then choose examples that force those decisions. Small functions are useful because they make those rules easy to verify, not because every function must have an arbitrary line limit.

### Example

```
package main
import "fmt"
func subtotal(priceCents, quantity int) (int, bool) {
    if priceCents < 0 || quantity < 0 { return 0, false }
    return priceCents * quantity, true
}
func main() {
    total, ok := subtotal(350, 3)
    fmt.Println(total, ok)
}
```

### Exercise

Write clamp(n, low, high int) (int, bool). Reject low > high; otherwise return n limited to the inclusive range.

Hint: Validate the range first, then handle values below low and above high.

Review criteria:

- [ ] clamp(5,0,10) returns 5,true
- [ ] clamp(-2,0,10) returns 0,true
- [ ] clamp(5,10,0) returns 0,false

### Worked solution

```
func clamp(n, low, high int) (int, bool) {
    if low > high { return 0, false }
    if n < low { return low, true }
    if n > high { return high, true }
    return n, true
}
// Contextual function: add package main and a main to run.
```

### Checkpoint

Why separate calculation from fmt.Println?

1. It always makes code faster
2. The calculation can be reused and tested without capturing output
3. Go forbids printing inside functions

Answer: 2. Separating the result from presentation makes the same behavior available to tests, command-line tools, and web APIs.

Sources:

- [Go language specification](https://go.dev/ref/spec)
- [Create a module](https://go.dev/doc/tutorial/create-module)

## Slices and maps without surprises

40 minutes of guided study, plus practice.

Learning objectives:

- Build and iterate slices
- Look up map values safely
- Understand slice sharing and map initialization

### A slice is a view

Arrays have a fixed length in their type; slices are flexible views over an underlying array. A slice carries a pointer, length, and capacity. append may allocate a new backing array, so assign its returned slice. Two slices can refer to the same elements. If you need an independent copy, allocate a new slice and copy the data. This matters when one function should not alter another function’s view.

### Maps answer keyed questions

A map associates keys with values. Initialize it with make or a map literal before writing. Reading an absent key returns the value type’s zero value. Use value, ok := m[key] to distinguish a missing entry from a stored zero. delete removes a key. Map iteration order is unspecified: collect and sort keys when you need deterministic display or test output.

### Know the empty cases

A nil slice has length zero and can be appended to. A nil map can be read but cannot receive assignments. With range on a slice, the first value is an index and the second is an element value. Mutating that element variable does not replace the slice element; use its index. Test empty inputs before building an algorithm around the first element.

### Example

```
package main
import "fmt"
func main() {
    names := []string{"Ada", "Lin", "Ada"}
    counts := make(map[string]int)
    for _, name := range names { counts[name]++ }
    n, found := counts["Ada"]
    fmt.Println(n, found, len(names))
}
```

### Exercise

Write unique([]string) []string that keeps the first occurrence of each name in input order.

Hint: Use a map[string]bool to remember names and a slice to preserve order.

Review criteria:

- [ ] Ada,Lin,Ada becomes Ada,Lin
- [ ] Empty input returns zero elements
- [ ] Input order is preserved without iterating the map

### Worked solution

```
func unique(names []string) []string {
    seen := make(map[string]bool)
    out := make([]string, 0, len(names))
    for _, name := range names {
        if !seen[name] {
            seen[name] = true
            out = append(out, name)
        }
    }
    return out
}
// Contextual function.
```

### Checkpoint

What happens when you assign m["x"] = 1 to a nil map?

1. The map initializes itself
2. The assignment panics
3. The key is ignored

Answer: 2. A nil map can be read but writing to it panics. Initialize writable maps with make or a literal.

Sources:

- [Go language specification](https://go.dev/ref/spec)
- [Go slices: usage and internals](https://go.dev/blog/slices-intro)

## Text, parsing, and useful validation

35 minutes of guided study, plus practice.

Learning objectives:

- Distinguish bytes from runes
- Normalize user input deliberately
- Parse numbers and handle errors

### Text has layers

A Go string stores bytes, commonly UTF-8. Indexing a string gives a byte; ranging over it decodes Unicode code points called runes. Neither a byte count nor a rune count always equals the number of visible characters because combining marks and emoji sequences exist. Use the representation your task needs. For a network byte limit, len is appropriate. For many code-point operations, range is appropriate.

### Parsing can fail

User input is text, even when it looks numeric. strconv.Atoi returns an integer and an error. Check the error before using the result. Trimming surrounding whitespace can help with human input, but silently changing passwords or identifiers may change their meaning. Validate business rules separately from syntax: parsing -3 succeeds, yet a quantity may still have to be positive.

### Build a boundary around input

At an input boundary, state allowed values, size limits, and normalization rules. Keep raw input out of SQL construction and HTML construction. Those risks are solved with the database and template APIs in later modules, not by inventing a universal sanitization function. A good error tells the user what to correct without exposing internal implementation details.

### Example

```
package main
import (
    "fmt"
    "strconv"
    "strings"
)
func main() {
    raw := " 12 "
    quantity, err := strconv.Atoi(strings.TrimSpace(raw))
    if err != nil || quantity < 1 {
        fmt.Println("Enter a positive whole number")
        return
    }
    fmt.Println("Quantity:", quantity)
}
```

### Exercise

Implement parseQuantity(string) (int,error) that trims spaces, accepts 1 through 100, and rejects other inputs with a helpful error.

Hint: Use strconv.Atoi first, then check the inclusive bounds.

Review criteria:

- [ ] " 12 " returns 12,nil
- [ ] "many", "0", and "101" return errors
- [ ] The parser never reports success after an Atoi error

### Worked solution

```
// Contextual: imports fmt, strconv, strings.
func parseQuantity(raw string) (int, error) {
    n, err := strconv.Atoi(strings.TrimSpace(raw))
    if err != nil { return 0, fmt.Errorf("quantity must be a whole number") }
    if n < 1 || n > 100 { return 0, fmt.Errorf("quantity must be 1 to 100") }
    return n, nil
}
```

### Checkpoint

What does len("Go✓") count?

1. Visible characters
2. UTF-8 bytes
3. Words

Answer: 2. len counts bytes in a string. The check mark uses more than one UTF-8 byte.

Sources:

- [strconv package](https://pkg.go.dev/strconv)
- [strings package](https://pkg.go.dev/strings)
- [Strings, bytes, runes and characters](https://go.dev/blog/strings)

## Project: Build a word-frequency tool

Turn a paragraph into a frequency report. Define whether punctuation counts as part of a word, document the choice, and print the most common words in deterministic order.

Requirements:

- Split and normalize text deliberately
- Use a map for counts
- Sort output; break ties alphabetically
- Handle empty text and repeated whitespace

Self-review rubric:

- [ ] Correct counting (40%)
- [ ] Explicit text rules (20%)
- [ ] Deterministic output and edge cases (40%)

# 3. Design with types

Structs, interfaces, errors, and generics

Level: Core Go. Prerequisite: Work with real data.

## Structs, methods, and pointers

40 minutes of guided study, plus practice.

Learning objectives:

- Model a domain with struct fields
- Choose value or pointer receivers
- Explain mutation through a pointer

### Give related values a home

A struct groups named fields into a new value. A Task can contain an ID, title, and completion flag. Values of the same struct type share a shape, which prevents accidental confusion between unrelated strings. Exported names start with an uppercase letter. Export only what callers need; a small public surface is easier to maintain.

### Methods attach behavior

A method is a function with a receiver. A value receiver works with a copy; a pointer receiver can change the original value. Use pointer receivers for intentional mutation and be consistent across methods when that improves clarity. A pointer is an address, not a separate owner of the data. Nil pointers must not be dereferenced. You do not manually free ordinary Go allocations; the runtime manages memory.

### Keep invariants close

A type should protect rules that must always hold. For example, refuse a blank title at construction instead of trusting every caller to remember. A constructor is a convention, usually NewTask, not a language feature. Avoid adding getters and setters mechanically: use methods when they protect a rule or express a domain action. MarkDone communicates more than SetDone(true).

### Example

```
package main
import "fmt"
type Task struct { Title string; Done bool }
func (t *Task) Complete() { t.Done = true }
func main() {
    t := Task{Title: "Learn methods"}
    t.Complete()
    fmt.Println(t.Title, t.Done)
}
```

### Exercise

Add Rename(title string) error to Task. Trim spaces and refuse an empty title without changing the existing one.

Hint: Validate before assignment. The receiver must be *Task to update the caller’s value.

Review criteria:

- [ ] Renaming to " New " stores "New"
- [ ] Blank rename returns an error
- [ ] Blank rename preserves the old title

### Worked solution

```
// Contextual: imports errors and strings; uses Task above.
func (t *Task) Rename(title string) error {
    title = strings.TrimSpace(title)
    if title == "" { return errors.New("title is required") }
    t.Title = title
    return nil
}
```

### Checkpoint

Which receiver lets Complete change the caller’s Task?

1. Task
2. *Task
3. string

Answer: 2. A pointer receiver refers to the original Task, so changing its field changes the caller’s value.

Sources:

- [Go language specification](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)

## Errors and small interfaces

45 minutes of guided study, plus practice.

Learning objectives:

- Return and wrap errors with context
- Match errors with errors.Is
- Define an interface at the point of use

### Failure is part of the contract

An error is a value describing why an operation could not complete. Return expected failures and let the caller decide how to present or recover from them. panic is not a replacement for ordinary validation. When a lower-level operation fails, fmt.Errorf with %w adds context while preserving the underlying error for inspection. Avoid logging and returning the same error at every layer; log once where the failure is handled.

### Depend on behavior

An interface lists methods. A type satisfies it implicitly by having those methods; there is no implements keyword. Define the smallest interface the consuming code needs. A notification function may need only Send(string) error rather than every feature of a mail client. This makes dependencies clearer and tests easier because a tiny fake can supply the required behavior.

### Inspect errors without reading their prose

Use errors.Is for sentinel errors that remain meaningful through wrapping, and errors.As for an error type with structured details. Comparing message strings breaks when context changes. Watch typed nil values stored in interfaces: an interface can be non-nil even if the concrete pointer it contains is nil. Return a plain nil error on success.

### Example

```
package main
import (
    "errors"
    "fmt"
)
var ErrMissing = errors.New("task not found")
type Finder interface { Find(int) (string, error) }
type Memory map[int]string
func (m Memory) Find(id int) (string, error) {
    title, ok := m[id]
    if !ok { return "", fmt.Errorf("find %d: %w", id, ErrMissing) }
    return title, nil
}
func main() {
    var f Finder = Memory{1: "Read"}
    _, err := f.Find(9)
    fmt.Println(errors.Is(err, ErrMissing))
}
```

### Exercise

Add a loadTitle(f Finder,id int) function that returns "untitled",nil only for ErrMissing and wraps any other error.

Hint: Use errors.Is rather than comparing err.Error().

Review criteria:

- [ ] Wrapped ErrMissing produces untitled,nil
- [ ] Other errors remain discoverable with errors.Is
- [ ] A successful title is unchanged

### Worked solution

```
// Contextual: uses Finder, ErrMissing, errors, and fmt above.
func loadTitle(f Finder, id int) (string, error) {
    title, err := f.Find(id)
    if errors.Is(err, ErrMissing) { return "untitled", nil }
    if err != nil { return "", fmt.Errorf("load title: %w", err) }
    return title, nil
}
```

### Checkpoint

How should you detect ErrMissing after %w wrapping?

1. Compare error strings
2. Use errors.Is(err, ErrMissing)
3. Ignore the wrapper

Answer: 2. errors.Is follows the wrapping chain and checks whether the error matches the target.

Sources:

- [errors package](https://pkg.go.dev/errors)
- [Working with Errors in Go](https://go.dev/blog/go1.13-errors)

## Generics when the algorithm stays the same

40 minutes of guided study, plus practice.

Learning objectives:

- Read type parameters and constraints
- Write a small generic utility
- Choose interfaces or generics deliberately

### Separate shape from algorithm

Sometimes the same algorithm works for several element types without needing their specific behavior. A generic function declares type parameters in square brackets. A constraint tells the compiler what operations are available. comparable permits equality checks and therefore supports a generic map-based set. any permits any type, but gives you no arbitrary arithmetic or method operations.

### Keep the useful distinction

Interfaces are a good fit when callers provide behavior, such as reading bytes or finding tasks. Generics are a good fit when an algorithm operates on values with a shared constraint while preserving their concrete type. Neither tool makes code automatically better. If a concrete function is short and clear and has only one use, start there instead of designing a universal abstraction.

### Read the result type

A function returning []T keeps the element type of its input. Type inference often lets callers omit explicit type arguments. Be careful with comparable: it is about equality, not sorting. Ordering requires a different constraint or a comparison function. Focus first on the contract of one small utility, and verify it with both strings and integers.

### Example

```
package main
import "fmt"
func Contains[T comparable](items []T, target T) bool {
    for _, item := range items {
        if item == target { return true }
    }
    return false
}
func main() {
    fmt.Println(Contains([]int{2, 4, 6}, 4))
    fmt.Println(Contains([]string{"go", "sql"}, "web"))
}
```

### Exercise

Implement Unique[T comparable] preserving first-occurrence order. Try it with []int and []string.

Hint: Adapt your earlier map-based unique function; the map becomes map[T]bool.

Review criteria:

- [ ] []int{2,1,2} becomes {2,1}
- [ ] []string{"a","a"} becomes {"a"}
- [ ] Explain why []int cannot itself be a comparable map key

### Worked solution

```
func Unique[T comparable](items []T) []T {
    seen := make(map[T]bool)
    out := make([]T, 0, len(items))
    for _, item := range items {
        if !seen[item] {
            seen[item] = true
            out = append(out, item)
        }
    }
    return out
}
// Contextual function.
```

### Checkpoint

What does comparable allow in a generic function?

1. All arithmetic
2. Equality and inequality comparisons
3. Calling any method

Answer: 2. comparable allows == and != and use as a map key. It does not imply arithmetic, ordering, or arbitrary methods.

Sources:

- [Generics tutorial](https://go.dev/doc/tutorial/generics)
- [Go language specification](https://go.dev/ref/spec)

## Project: Design an in-memory task library

Create a task package with IDs, titles, completion, creation, lookup, and rename operations. Keep storage in memory and return meaningful errors.

Requirements:

- Reject blank titles
- Define ErrNotFound and preserve it when wrapping
- Expose a small consumer-defined repository interface
- Demonstrate the package from a separate main

Self-review rubric:

- [ ] Domain rules and errors (40%)
- [ ] Useful boundaries and receiver choices (30%)
- [ ] Runnable demonstration with failure cases (30%)

# 4. Build useful Go programs

Modules, files, streams, and JSON

Level: Core Go. Prerequisite: Design with types.

## Packages and dependency hygiene

35 minutes of guided study, plus practice.

Learning objectives:

- Organize packages around responsibilities
- Use go.mod and go.sum correctly
- Inspect dependencies and tool help

### A package is a collaboration boundary

A package is made from Go files in one directory with a shared package name. A module contains one or more packages and declares its import path in go.mod. Start small: main can wire together an internal/task package. internal limits who may import that directory’s packages. Avoid a giant utils package; name a package after what it does for its callers.

### Dependencies are versioned inputs

Use go get to add or change a dependency, and go mod tidy to reconcile imports and module requirements. Commit both go.mod and go.sum. The sum file checks downloaded module content; it is not a substitute for reviewing a dependency’s behavior or maintenance. Avoid copying random installation commands into a production project without understanding the added module.

### Use tools as documentation

go help and go doc are available without a browser. go list -m all shows the selected module graph. A go.work file can coordinate local work across modules, but it is unnecessary for a one-module beginner project. Add complexity only when a real boundary requires it. Before releasing a library, read module versioning and compatibility guidance rather than treating version numbers as decoration.

### Example

```
# Terminal commands; run in a new folder.
go mod init example.com/taskboard
go help modules
go doc fmt.Printf
go list -m all
go mod tidy
go test ./...
```

### Exercise

Move your Task type and behavior into internal/task, then import it from cmd/taskboard/main.go. Explain which names must be exported.

Hint: The import path combines the module path with /internal/task; the package name can be task.

Review criteria:

- [ ] go run ./cmd/taskboard succeeds
- [ ] go test ./... discovers all packages
- [ ] Explain why another unrelated module cannot import this internal package

### Worked solution

```
// internal/task/task.go
package task
type Task struct { Title string; Done bool }
func (t *Task) Complete() { t.Done = true }

// cmd/taskboard/main.go (separate file)
package main
import (
    "fmt"
    "example.com/taskboard/internal/task"
)
func main() {
    t := task.Task{Title: "Organize code"}
    t.Complete()
    fmt.Println(t)
}
```

### Checkpoint

What should normally be committed for dependencies?

1. Only go.sum
2. go.mod and go.sum
3. The entire module download cache

Answer: 2. go.mod declares module requirements and go.sum records checksums used to verify downloads. Both belong in source control.

Sources:

- [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
- [Organizing a module](https://go.dev/doc/modules/layout)

## Files, readers, and resource ownership

40 minutes of guided study, plus practice.

Learning objectives:

- Read through io.Reader
- Close resources at the right lifetime
- Handle scanner and write errors

### Program to a stream

io.Reader describes something from which bytes can be read. Files, strings, and HTTP bodies can all implement it. A function taking an io.Reader is easier to reuse and test than one that insists on opening a specific filename. Let one layer choose the source and another process the content. This keeps file permissions and path choices outside your parsing logic.

### Acquire, check, release

After successfully opening a file, arrange its close with defer. Never defer a close before checking the open error. A defer runs when its surrounding function returns, so repeatedly opening files in a long loop can keep too many open; use a small helper per file. Write and close errors can matter for output durability, so do not assume a successful open means the whole write succeeded.

### Bound your assumptions

bufio.Scanner is convenient for lines and words, but it has a maximum token size unless configured otherwise. Always inspect scanner.Err after scanning. For untrusted input, set a deliberate byte or token limit and return a useful error when exceeded. A command-line toy may read a small file in one operation; a service must account for many simultaneous inputs.

### Example

```
package main
import (
    "bufio"
    "fmt"
    "io"
    "strings"
)
func countLines(r io.Reader) (int, error) {
    s := bufio.NewScanner(r)
    n := 0
    for s.Scan() { n++ }
    return n, s.Err()
}
func main() {
    n, err := countLines(strings.NewReader("Go\nSQL\n"))
    fmt.Println(n, err)
}
```

### Exercise

Write countWords(io.Reader) (int,error) with Scanner and ScanWords. Test empty input and text with repeated whitespace.

Hint: Call scanner.Split(bufio.ScanWords) before the scanning loop.

Review criteria:

- [ ] Empty reader returns 0,nil
- [ ] "Go  is\nfun" returns 3,nil
- [ ] A scanning failure is returned instead of a partial success

### Worked solution

```
// Contextual: imports bufio and io.
func countWords(r io.Reader) (int, error) {
    s := bufio.NewScanner(r)
    s.Split(bufio.ScanWords)
    n := 0
    for s.Scan() { n++ }
    if err := s.Err(); err != nil { return 0, err }
    return n, nil
}
```

### Checkpoint

When does a deferred file.Close run?

1. At the end of the current loop iteration
2. When the surrounding function returns
3. Immediately when defer is encountered

Answer: 2. defer schedules the call for the return of its surrounding function. Use a helper when each loop iteration needs its own resource lifetime.

Sources:

- [io package](https://pkg.go.dev/io)
- [bufio package](https://pkg.go.dev/bufio)
- [os package](https://pkg.go.dev/os)

## JSON as an explicit data contract

40 minutes of guided study, plus practice.

Learning objectives:

- Encode and decode struct data
- Use JSON tags and exported fields
- Separate transport input from domain data

### Define the wire shape

JSON is a language-independent format for exchanging data. encoding/json maps object keys to exported Go struct fields. Tags such as json:"title" make your public API naming intentional. Unexported fields do not become public just because they have a tag. A missing boolean and false both produce false in a plain bool field; use a pointer when an update must distinguish absence from an explicit value.

### Decoding is only the first check

A successful Decode means the input fits the destination representation, not that it obeys your business rules. Validate title length, allowed states, and ownership separately. At HTTP boundaries, limit request-body size and consider DisallowUnknownFields for strict request objects. Reject extra JSON documents when your endpoint promises exactly one. Never decode a user-supplied owner ID directly into a trusted ownership decision.

### Protect compatibility

Use separate request and response types when internal data contains secrets or fields clients must not set. Omitting a password field is safer than relying on every handler to remember not to serialize it. Write examples of valid and invalid requests. When changing a public API, think about existing clients before renaming keys or changing their types.

### Example

```
package main
import (
    "encoding/json"
    "fmt"
)
type Task struct {
    Title string `json:"title"`
    Done bool `json:"done"`
}
func main() {
    data, err := json.Marshal(Task{Title: "Learn JSON"})
    if err != nil { panic(err) } // no expected failure for this fixed value
    fmt.Println(string(data))
}
```

### Exercise

Define UpdateTask with a *bool done field. Decode {} and {"done":false}; explain how the results differ.

Hint: A missing field leaves the pointer nil. An explicit false creates a non-nil pointer to false.

Review criteria:

- [ ] {} means no change
- [ ] {"done":false} means explicitly mark incomplete
- [ ] Malformed JSON returns an error

### Worked solution

```
// Contextual: imports encoding/json and fmt.
type UpdateTask struct { Done *bool `json:"done"` }
func inspect(raw string) error {
    var in UpdateTask
    if err := json.Unmarshal([]byte(raw), &in); err != nil { return err }
    if in.Done == nil { fmt.Println("no change") } else { fmt.Println(*in.Done) }
    return nil
}
```

### Checkpoint

Can a JSON tag expose an unexported struct field?

1. Yes
2. No
3. Only with omitempty

Answer: 2. encoding/json operates on exported fields. Tags customize their representation but do not bypass visibility.

Sources:

- [encoding/json package](https://pkg.go.dev/encoding/json)
- [JSON and Go](https://go.dev/blog/json)

## Project: Save a task list to JSON

Build a CLI that loads tasks from a JSON file, modifies them in memory, and saves a new file. Keep parsing functions separate from filesystem operations.

Requirements:

- Use explicit JSON tags
- Reject malformed data and invalid titles
- Report file and encoding errors
- Document that this single-process file store is not a concurrent database

Self-review rubric:

- [ ] Data round trip and validation (40%)
- [ ] Resource and error handling (35%)
- [ ] Clear CLI usage and limitations (25%)

# 5. Make correctness visible

Tests, HTTP checks, fuzzing, and benchmarks

Level: Intermediate. Prerequisite: Build useful Go programs.

## Table-driven tests that catch mistakes

40 minutes of guided study, plus practice.

Learning objectives:

- Write Test functions in _test.go files
- Choose boundary cases from contracts
- Use subtests and useful failure messages

### Tests explain behavior

A test is an executable claim about a contract. Put tests in files ending _test.go and use functions named TestSomething with *testing.T. Run go test ./... from the module root. Good tests tell a reader which behavior matters. A test that repeats the implementation line by line can pass while both copies contain the same mistake. Choose expected results from the rules instead.

### A table makes cases easy to compare

Table-driven tests keep inputs and expected outputs together. Use t.Run to give each case a readable name. Include ordinary inputs, boundaries, and invalid inputs. Failure messages should say what input was used and how actual and expected results differ. Do not use sleeps to make deterministic behavior appear stable. Isolate time or external dependencies behind small boundaries when needed.

### Read coverage as a question

go test -cover reports executed statements, not proven correctness. A suite can reach every line while checking the wrong outcome. Use coverage to find unexamined paths, then add meaningful assertions. For error cases, assert error identity or a stable property instead of a fragile full message unless the message itself is the public contract.

### Example

```
// clamp_test.go; package must match your clamp implementation.
package main
import "testing"
func TestClamp(t *testing.T) {
    cases := []struct { name string; n, low, high, want int; ok bool }{
        {"inside", 5, 0, 10, 5, true},
        {"below", -1, 0, 10, 0, true},
        {"bad range", 5, 10, 0, 0, false},
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            got, ok := clamp(tc.n, tc.low, tc.high)
            if got != tc.want || ok != tc.ok {
                t.Fatalf("clamp(%d,%d,%d)=(%d,%t), want (%d,%t)",
                    tc.n, tc.low, tc.high, got, ok, tc.want, tc.ok)
            }
        })
    }
}
```

### Exercise

Add tests for an input equal to each boundary, an input above the high boundary, and a range whose low equals high.

Hint: A one-value range is valid. Derive expected results before running the tests.

Review criteria:

- [ ] go test ./... passes
- [ ] A deliberately broken comparison causes a relevant test to fail
- [ ] Failure output identifies the case

### Worked solution

```
// Add these rows to the test table:
{"at low", 0, 0, 10, 0, true},
{"at high", 10, 0, 10, 10, true},
{"above", 11, 0, 10, 10, true},
{"single value", 99, 4, 4, 4, true},
```

### Checkpoint

Does 100% statement coverage prove correctness?

1. Yes
2. No; it only shows those statements executed
3. Only in small packages

Answer: 2. Coverage measures execution. Assertions, input choices, and the intended contract determine what the tests actually verify.

Sources:

- [testing package](https://pkg.go.dev/testing)
- [Coverage for Go applications](https://go.dev/doc/build-cover)

## Test behavior at an HTTP boundary

40 minutes of guided study, plus practice.

Learning objectives:

- Use httptest without binding a port
- Assert status, headers, and body
- Test invalid requests and dependencies

### Exercise the real boundary

An HTTP handler has observable behavior: status code, headers, and response body. net/http/httptest creates requests and records responses without opening a network listener. Test the handler through the same routing and middleware stack clients will use when those layers affect behavior. Calling only a helper function can miss a route mismatch or a missing authorization check.

### Assert what clients depend on

Check Content-Type and the response status before decoding JSON. For JSON objects, compare decoded fields instead of byte-for-byte formatting unless whitespace is part of the contract. Include unsupported methods, malformed input, missing resources, and unauthorized ownership. A successful create request should not be your only test of a write endpoint.

### Keep fakes honest

A fake repository can force not-found and dependency-failure paths quickly. Integration tests against a real disposable database still matter because a fake cannot validate SQL syntax, constraints, or transaction behavior. Keep unit tests fast and focused, then add a smaller set of integration tests for the boundaries that can fail differently in production.

### Example

```
// health_test.go; self-contained test example.
package main
import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)
func TestHealth(t *testing.T) {
    h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        _, _ = w.Write([]byte("{\"ok\":true}\n"))
    })
    rec := httptest.NewRecorder()
    h.ServeHTTP(rec, httptest.NewRequest("GET", "/health", nil))
    if rec.Code != 200 || !strings.Contains(rec.Body.String(), "true") {
        t.Fatalf("unexpected response: %d %s", rec.Code, rec.Body.String())
    }
}
```

### Exercise

Improve the sample test to assert Content-Type and decode the body into a struct with an OK bool field.

Hint: Use json.NewDecoder(rec.Body).Decode(&body), and fail immediately if decoding fails.

Review criteria:

- [ ] Status, header, and decoded field are asserted
- [ ] Invalid JSON fails the test
- [ ] Changing ok to false fails the test

### Worked solution

```
// Contextual additions: import encoding/json.
if got := rec.Header().Get("Content-Type"); got != "application/json" {
    t.Fatalf("Content-Type = %q", got)
}
var body struct { OK bool `json:"ok"` }
if err := json.NewDecoder(rec.Body).Decode(&body); err != nil { t.Fatal(err) }
if !body.OK { t.Fatal("expected ok=true") }
```

### Checkpoint

What can a fake repository not prove?

1. Handler not-found behavior
2. The real SQL query obeys the database schema
3. A dependency error becomes a 500

Answer: 2. A fake can exercise handler branches, but only database-backed integration checks validate actual SQL, schema constraints, and transactions.

Sources:

- [net/http/httptest package](https://pkg.go.dev/net/http/httptest)
- [encoding/json package](https://pkg.go.dev/encoding/json)

## Fuzz invariants and measure before optimizing

45 minutes of guided study, plus practice.

Learning objectives:

- Choose a useful fuzz property
- Run the race detector and static checks
- Write a basic benchmark

### Ask a property instead of listing everything

A fuzz test generates many inputs for a stable property. A parser might promise never to panic, or a round trip might preserve valid data. Seed the corpus with meaningful cases, then run go test -fuzz=FuzzName -fuzztime=10s in the package. Keep discovered failures as regression cases. A property must be true by design: reversing arbitrary UTF-8 bytes twice is different from reversing user-visible characters.

### Use several kinds of evidence

go vet ./... finds suspicious constructs; go test -race ./... detects data races in executed paths on supported platforms. Neither tool proves that unexecuted paths are safe. Fuzzing and race detection complement ordinary examples. Treat a failure as a small reproducible input to understand, not an obstacle to suppress.

### Benchmark a real question

Benchmarks use functions named BenchmarkSomething with *testing.B. The traditional b.N loop works across many Go versions. Run go test -bench=. -benchmem and compare several runs before drawing conclusions. Keep setup outside the measured loop when it is not part of the operation. Do not trade away clear error handling because a tiny isolated benchmark looked faster. Profile an actual bottleneck first.

### Example

```
// properties_test.go; uses a simple identity-preserving codec property.
package main
import (
    "encoding/hex"
    "testing"
)
func FuzzHexRoundTrip(f *testing.F) {
    f.Add([]byte("Go"))
    f.Add([]byte{})
    f.Fuzz(func(t *testing.T, in []byte) {
        out, err := hex.DecodeString(hex.EncodeToString(in))
        if err != nil || string(out) != string(in) {
            t.Fatalf("round trip failed: %v", err)
        }
    })
}
func BenchmarkHex(b *testing.B) {
    data := []byte("learn Go by building")
    for i := 0; i < b.N; i++ { _ = hex.EncodeToString(data) }
}
```

### Exercise

Add a fuzz property for parseQuantity: every successful result must be between 1 and 100. Seed with valid, invalid, and boundary inputs.

Hint: Do not require every input to parse. Assert the range only when err is nil.

Review criteria:

- [ ] Fuzz seeds pass in ordinary go test
- [ ] A bounded fuzz run completes
- [ ] A deliberately relaxed upper bound is detected

### Worked solution

```
// Contextual test: same package as parseQuantity; import testing.
func FuzzQuantity(f *testing.F) {
    for _, s := range []string{"1", "100", "0", "101", "many", " 12 "} { f.Add(s) }
    f.Fuzz(func(t *testing.T, raw string) {
        n, err := parseQuantity(raw)
        if err == nil && (n < 1 || n > 100) {
            t.Fatalf("successful quantity out of range: %d", n)
        }
    })
}
```

### Checkpoint

Which is a useful parser fuzz invariant?

1. Every random string must parse successfully
2. Successful results always obey the documented range
3. The parser must finish in exactly one nanosecond

Answer: 2. Many strings should fail to parse. The useful invariant checks that accepted inputs satisfy the parser’s promised result constraints.

Sources:

- [Fuzzing tutorial](https://go.dev/doc/tutorial/fuzz)
- [Race detector](https://go.dev/doc/articles/race_detector)
- [testing package](https://pkg.go.dev/testing)

## Project: Give your task library a safety net

Test the task package and JSON boundary. Add at least one meaningful property test and a documented command for running the suite.

Requirements:

- Table tests for valid and invalid titles
- Errors.Is assertions for missing tasks
- JSON round-trip or validation property
- One deliberate mutation of the implementation must fail a test

Self-review rubric:

- [ ] Behavioral coverage (40%)
- [ ] Failure clarity and deterministic tests (30%)
- [ ] Useful property and reproducibility (30%)

# 6. Concurrency with a clear lifetime

Goroutines, channels, synchronization, and cancellation

Level: Intermediate. Prerequisite: Make correctness visible.

## Goroutines and channels

45 minutes of guided study, plus practice.

Learning objectives:

- Start concurrent work and collect results
- Explain blocking channel operations
- Give every goroutine a stopping condition

### Concurrency is a coordination problem

A goroutine runs a function concurrently with other goroutines. Starting one is easy; knowing when it finishes is the real design work. The main goroutine does not automatically wait for the others. Give work an owner, a way to report a result, and a lifetime. Do not reach for concurrency until independent work or waiting makes it useful.

### Channels communicate values

An unbuffered channel send waits for a receiver and a receive waits for a sender. A buffered channel can hold a bounded number of values before the sender blocks. Buffering changes coordination and resource use; it does not fix an unclear lifetime. Only the sender that knows no more values will be sent should normally close the channel. Receiving from a closed channel yields buffered values, then zero values with ok=false.

### Close the loop

A range over a channel ends when the channel closes. Sending on a closed channel panics. Leaving a receiver waiting forever leaks a goroutine. Sketch the sequence of sends, receives, and closure before running code. Use a channel for data transfer and synchronization when that makes the flow clear, and use a mutex for shared mutable state when that is simpler.

### Example

```
package main
import "fmt"
func squares(nums []int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums { out <- n*n }
    }()
    return out
}
func main() {
    for n := range squares([]int{1, 2, 3}) { fmt.Println(n) }
}
```

### Exercise

Write a producer that sends the strings Go, SQL, Web and closes its output. Have main collect all three without using time.Sleep.

Hint: The producing goroutine owns close; main ranges until closure.

Review criteria:

- [ ] Exactly three values arrive
- [ ] The receiving loop terminates
- [ ] No sleeps are used to guess completion

### Worked solution

```
package main
import "fmt"
func main() {
    out := make(chan string)
    go func() {
        defer close(out)
        for _, s := range []string{"Go", "SQL", "Web"} { out <- s }
    }()
    for s := range out { fmt.Println(s) }
}
```

### Checkpoint

Who should normally close a channel?

1. Any receiver that is done reading
2. The sender that knows no more values will be sent
3. Every goroutine using it

Answer: 2. The sender owns knowledge of completion. Closing from an arbitrary receiver can race with a send and cause a panic.

Sources:

- [Go concurrency patterns: pipelines](https://go.dev/blog/pipelines)
- [Go language specification](https://go.dev/ref/spec)

## Shared state and bounded workers

45 minutes of guided study, plus practice.

Learning objectives:

- Protect compound updates with a mutex
- Wait for workers explicitly
- Bound work instead of spawning without limits

### A race is about unsynchronized access

When goroutines access the same memory concurrently and at least one writes without appropriate synchronization, the program has a data race. Even counter++ is a read-modify-write sequence. A mutex protects the whole invariant, not just a single line chosen at random. Keep locked sections short and avoid calling unknown external code while holding a lock.

### Wait for the work you own

sync.WaitGroup tracks a set of tasks. Add the count before starting goroutines, call Done when each exits, and Wait at the owner. A WaitGroup coordinates completion; it does not protect shared data. Do not copy a mutex or WaitGroup after use. Put concurrency primitives in stable objects and access them through pointers.

### Bound throughput intentionally

Starting one goroutine per item can create unbounded pressure on memory, downstream APIs, and database connections. A worker pool uses a fixed number of workers and a bounded queue. Decide what happens when the queue is full: block, reject, or drop according to the application contract. Test shutdown and cancellation, not just the fast successful path.

### Example

```
package main
import (
    "fmt"
    "sync"
)
func main() {
    var mu sync.Mutex
    var wg sync.WaitGroup
    count := 0
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            count++
            mu.Unlock()
        }()
    }
    wg.Wait()
    fmt.Println(count)
}
```

### Exercise

Extend the program so each of ten goroutines increments 100 times. Confirm the result is 1000 under the race detector.

Hint: Protect each increment, or accumulate locally and add once under the lock.

Review criteria:

- [ ] Final count is always 1000
- [ ] go run -race . reports no races on a supported platform
- [ ] Removing synchronization demonstrates why the shared update needs protection

### Worked solution

```
// Replace the goroutine body; keep wg.Add before go.
go func() {
    defer wg.Done()
    local := 0
    for j := 0; j < 100; j++ { local++ }
    mu.Lock()
    count += local
    mu.Unlock()
}()
```

### Checkpoint

Does a WaitGroup protect count++ from data races?

1. Yes
2. No; it coordinates completion only
3. Only with one CPU

Answer: 2. WaitGroup handles completion. A mutex or another synchronization design is still needed for concurrent writes to shared state.

Sources:

- [sync package](https://pkg.go.dev/sync)
- [Memory model](https://go.dev/ref/mem)
- [Race detector](https://go.dev/doc/articles/race_detector)

## Cancellation and deadlines

45 minutes of guided study, plus practice.

Learning objectives:

- Pass context through operation boundaries
- Use select to stop blocking work
- Release cancellation resources

### Cancellation is cooperative

context.Context carries cancellation signals, deadlines, and request-scoped values. It does not forcibly kill a goroutine. Each operation must observe ctx.Done or pass the context to an API that does. Accept context as the first parameter and propagate it instead of replacing it with Background halfway through a request. Do not put optional application settings in context values.

### Unblock every waiting point

A worker can be stuck waiting to receive work, sending a result, or calling a dependency. Every blocking point needs a completion or cancellation path. Use select with ctx.Done around channel operations. A buffered result channel may help a one-result worker finish after its caller leaves, but it does not stop the underlying expensive work.

### Own cleanup

When you create a cancelable context, call its cancel function, usually with defer, to release resources. Timeouts should reflect an operation budget and leave time for the caller to respond. In a web request use r.Context so client disconnection can cancel database queries. Return ctx.Err when cancellation wins so the caller can distinguish a deadline from other failures.

### Example

```
package main
import (
    "context"
    "fmt"
    "time"
)
func wait(ctx context.Context, delay time.Duration) error {
    timer := time.NewTimer(delay)
    defer timer.Stop()
    select {
    case <-timer.C: return nil
    case <-ctx.Done(): return ctx.Err()
    }
}
func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
    defer cancel()
    fmt.Println(wait(ctx, time.Second))
}
```

### Exercise

Write send(ctx,out,value) error that sends an integer unless the context is canceled. Test it with an already-canceled context and no receiver.

Hint: Use select, not a context check followed by an unconditional blocking send.

Review criteria:

- [ ] An available receiver gets the value
- [ ] Canceled context with no receiver returns promptly
- [ ] No goroutine remains blocked after cancellation

### Worked solution

```
// Contextual: import context.
func send(ctx context.Context, out chan<- int, value int) error {
    select {
    case out <- value: return nil
    case <-ctx.Done(): return ctx.Err()
    }
}
```

### Checkpoint

Does cancel() forcibly terminate a goroutine?

1. Yes
2. No; the goroutine must observe cancellation
3. Only if it uses a channel

Answer: 2. Cancellation is a signal. Code must observe it or use cancellation-aware dependencies to stop its work.

Sources:

- [context package](https://pkg.go.dev/context)
- [Go concurrency patterns: context](https://go.dev/blog/context)

## Project: Build a bounded concurrent checker

Check a list of URLs with at most four concurrent requests. Report status or error for each URL and support a global timeout.

Requirements:

- Reuse an http.Client with a timeout
- Attach request contexts
- Bound worker count and queue size
- Close result channels only after senders finish
- Test canceled and failed requests with httptest

Self-review rubric:

- [ ] Correct completion and cancellation (40%)
- [ ] Bounded resource use (30%)
- [ ] Race-free behavior and deterministic tests (30%)

# 7. Build a Go web API

Routing, JSON boundaries, middleware, and contracts

Level: Backend. Prerequisite: Concurrency with a clear lifetime.

## HTTP handlers and routes

40 minutes of guided study, plus practice.

Learning objectives:

- Read request and response responsibilities
- Use method-aware ServeMux patterns
- Configure a server deliberately

### HTTP is a request-response conversation

A request has a method, path, headers, and possibly a body. A response has a status, headers, and a body. A Go handler implements ServeHTTP; http.HandlerFunc adapts an ordinary function. Set response headers before writing a status or body. The first body write implicitly sends status 200 if no status has been written.

### Make the route contract explicit

Modern ServeMux supports method patterns such as GET /api/tasks/{id}; this requires Go 1.22 or later. Read captured parameters with r.PathValue. A GET should retrieve data without changing application state. POST commonly creates a resource. Decide how clients identify resources, which errors they can expect, and which fields they can supply before filling handlers with database calls.

### Use a real server configuration

http.ListenAndServe is convenient for the first example. For a deployed service, configure http.Server with timeouts and a graceful shutdown path. Restrict body size at write endpoints and use a client timeout for outgoing HTTP. The server is a resource boundary, not merely a function that happens to listen on a port.

### Example

```
package main
import (
    "encoding/json"
    "log"
    "net/http"
    "time"
)
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /api/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        _ = json.NewEncoder(w).Encode(map[string]string{"hello": r.PathValue("name")})
    })
    server := &http.Server{Addr: ":8080", Handler: mux,
        ReadHeaderTimeout: 5*time.Second, WriteTimeout: 10*time.Second,
        IdleTimeout: 60*time.Second}
    log.Fatal(server.ListenAndServe())
}
```

### Exercise

Add GET /health returning {"status":"ok"}. Test its content type and status with httptest, then call it in a browser.

Hint: Register another method-specific route before starting the server.

Review criteria:

- [ ] GET /health returns 200 and JSON
- [ ] POST /health returns method-not-allowed
- [ ] Header is set before writing the body

### Worked solution

```
// Contextual: inside main after mux is created.
mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
})
```

### Checkpoint

When should Content-Type be set?

1. After writing the body
2. Before writing the status or body
3. Only in the browser

Answer: 2. Response headers must be set before the response is committed by WriteHeader or the first body write.

Sources:

- [net/http package](https://pkg.go.dev/net/http)
- [Go 1.22 routing changes](https://go.dev/blog/routing-enhancements)

## Validate JSON at the edge

45 minutes of guided study, plus practice.

Learning objectives:

- Limit body size and reject malformed requests
- Separate request types from trusted domain fields
- Map failures to meaningful statuses

### Treat the body as untrusted input

A create-task request should accept only a title and perhaps a due date. Derive the owner from an authenticated session rather than a JSON owner_id supplied by the client. Use MaxBytesReader before decoding so a tiny endpoint cannot consume arbitrarily large bodies. DisallowUnknownFields helps catch misspelled keys when the API promises a strict shape.

### Require one document and valid meaning

Decode the request object, then decode once more and require io.EOF to reject trailing JSON documents. Validate a trimmed, non-empty title and a clear size limit. A decoding error is a client error; an unavailable database is a server error. Return a stable machine-readable error shape and keep detailed internal failures in server logs.

### Commit the response once

Avoid writing a success body and then trying to change the status when later work fails. Perform fallible validation and persistence first, then send the final response. Creation generally returns 201 with a representation or location; missing resources return 404; authorization decisions must not be skipped on updates or deletes. Test all of these outcomes.

### Example

```
// Contextual handler fragment; imports encoding/json, io, net/http, strings.
type createTaskInput struct { Title string `json:"title"` }
func createTask(w http.ResponseWriter, r *http.Request) {
    r.Body = http.MaxBytesReader(w, r.Body, 16<<10)
    defer r.Body.Close()
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()
    var in createTaskInput
    if err := dec.Decode(&in); err != nil { http.Error(w,"invalid JSON",400); return }
    var extra any
    if err := dec.Decode(&extra); err != io.EOF { http.Error(w,"one JSON object required",400); return }
    in.Title = strings.TrimSpace(in.Title)
    if in.Title == "" || len(in.Title) > 200 { http.Error(w,"title must be 1-200 bytes",400); return }
    // Persist here before claiming success. This lesson only echoes validated input.
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(in)
}
```

### Exercise

Add tests that reject an unknown key, an empty title, and two consecutive JSON objects. Explain why this echo handler returns 200 rather than 201.

Hint: The sample validates and echoes; it has not created a stored resource.

Review criteria:

- [ ] All three invalid cases are rejected
- [ ] Whitespace normalization is documented as title-specific
- [ ] The handler never claims creation without persistence

### Worked solution

```
// Test request bodies and expected status:
// {"title":"Go","owner_id":7} -> 400 (unknown key)
// {"title":"   "} -> 400 (invalid title)
// {"title":"Go"}{"title":"SQL"} -> 400 (extra document)
// {"title":" Go "} -> 200 and title "Go"
// Use httptest.NewRequest("POST", "/api/tasks", strings.NewReader(body)).
```

### Checkpoint

Should a trusted task owner come from request JSON?

1. Yes, because JSON is typed
2. No, derive it from the authenticated identity
3. Only if the field is named user_id

Answer: 2. JSON validation does not make an identity trustworthy. Resource ownership must come from authenticated server-side identity and authorization checks.

Sources:

- [net/http package](https://pkg.go.dev/net/http)
- [encoding/json package](https://pkg.go.dev/encoding/json)
- [io package](https://pkg.go.dev/io)

## Middleware and a stable API contract

40 minutes of guided study, plus practice.

Learning objectives:

- Compose handlers without hiding control flow
- Use stable error shapes and request IDs
- Separate handlers, domain logic, and storage

### Middleware wraps a handler

Middleware takes a handler and returns a handler. Common responsibilities include request IDs, logging, authentication, and security headers. Order matters: a request ID must exist before a log line needs it. Middleware should either call the next handler or send a complete response and return. Accidentally doing both can leak protected data or corrupt a response.

### Make the contract reviewable

Write a small API document listing method, path, request body, response, status codes, and ownership rules. For a task list, use GET and POST /api/tasks plus GET, PATCH, and DELETE /api/tasks/{id}. Define pagination before a list can grow without bound. Keep error envelopes consistent so the frontend does not need a new parsing rule for every endpoint.

### Let each layer have a job

A handler translates HTTP to application inputs. Domain code enforces task rules. A repository reads and writes storage. main wires dependencies together. Do not create layers that simply rename every function, but keep network and database details out of a pure rule such as title validation. This separation makes failures easier to test and future changes smaller.

### Example

```
// Contextual middleware; import net/http.
func securityHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
        next.ServeHTTP(w, r)
    })
}
// In main: server.Handler = securityHeaders(mux)
// Add a Content-Security-Policy suited to your actual frontend assets.
```

### Exercise

Create middleware that rejects requests missing X-Demo-Key before calling next. Use this only to practice middleware, not as a production user-auth system.

Hint: Return immediately after writing 401; use a counter handler in a test to prove it was not called.

Review criteria:

- [ ] Missing key returns 401 and next is never called
- [ ] Expected key calls next once
- [ ] README clearly says hard-coded keys are not production authentication

### Worked solution

```
// Contextual practice middleware; import net/http.
func requireDemoKey(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("X-Demo-Key") != "practice-only" {
            http.Error(w, "missing demo key", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

### Checkpoint

After middleware writes an unauthorized response, what should it do?

1. Call next anyway
2. Return without calling next
3. Write a success response

Answer: 2. A rejecting middleware owns that response and must return. Calling the protected handler afterward defeats the check.

Sources:

- [net/http package](https://pkg.go.dev/net/http)
- [Writing web applications](https://go.dev/doc/articles/wiki/)

## Project: Ship a tested task API

Implement an in-memory HTTP task API with create, list, get, update, and delete. Write its contract before connecting a frontend.

Requirements:

- Method-specific routes and valid status codes
- Strict bounded JSON input
- Consistent JSON errors
- Pagination or an explicit maximum list size
- Tests for invalid inputs and missing resources

Self-review rubric:

- [ ] HTTP contract correctness (40%)
- [ ] Boundary validation and failure handling (35%)
- [ ] Readable dependency structure and tests (25%)

# 8. Add durable storage

PostgreSQL, queries, migrations, and transactions

Level: Backend. Prerequisite: Build a Go web API.

## Model data and connect safely

45 minutes of guided study, plus practice.

Learning objectives:

- Design constraints and foreign keys
- Understand sql.DB as a connection pool
- Keep credentials out of code

### The database protects shared truth

An in-memory map disappears when a process stops. PostgreSQL stores records durably and coordinates concurrent access. Design a schema before writing queries: tasks need an ID, owner, title, completion state, and timestamps. NOT NULL, CHECK, UNIQUE, and foreign-key constraints protect data even when a future script or second application writes it. Application validation gives friendly errors; constraints protect the final boundary.

### A handle is a pool

database/sql provides a common interface; a driver connects it to PostgreSQL. sql.Open creates a handle that manages a pool rather than proving a live connection immediately. Use PingContext during startup when availability is required. Choose bounded pool settings using deployment and database capacity, not arbitrary large numbers. Close the pool on application shutdown, not after every query.

### Make schema change reproducible

Store numbered migration files in source control and apply them deliberately. A migration changes the schema; it should not silently run conflicting changes in every app replica. Keep database URLs in environment configuration, require appropriate TLS for remote connections, and use a least-privilege database account. Do not print credentials in connection errors or logs.

### Example

```
-- PostgreSQL migration; SQL, not a standalone Go program.
CREATE TABLE app_users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email TEXT NOT NULL UNIQUE
);
CREATE TABLE tasks (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    owner_id BIGINT NOT NULL REFERENCES app_users(id),
    title TEXT NOT NULL CHECK (char_length(title) BETWEEN 1 AND 200),
    done BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX tasks_owner_id_id_idx ON tasks(owner_id, id);
```

### Exercise

Write a migration adding an optional due_at timestamp to tasks. Describe how old application code behaves before and after this change.

Hint: A nullable added column lets old insert statements that name their columns continue to work.

Review criteria:

- [ ] Migration applies to a populated development database
- [ ] Existing rows have NULL due_at
- [ ] Explain that dropping the column loses stored due dates

### Worked solution

```
-- Up migration:
ALTER TABLE tasks ADD COLUMN due_at TIMESTAMPTZ;
-- A destructive down migration, for controlled development rollback only:
-- ALTER TABLE tasks DROP COLUMN due_at;
-- Existing rows get NULL and existing explicit-column inserts remain valid.
```

### Checkpoint

What is *sql.DB?

1. One permanent SQL connection
2. A handle managing a pool of connections
3. An in-memory replacement for PostgreSQL

Answer: 2. sql.DB is a concurrency-safe pool handle. Reuse it across requests instead of opening and closing a pool for each query.

Sources:

- [Opening a database handle](https://go.dev/doc/database/open-handle)
- [PostgreSQL constraints](https://www.postgresql.org/docs/current/ddl-constraints.html)

## Queries, ownership, and SQL injection

50 minutes of guided study, plus practice.

Learning objectives:

- Parameterize SQL values
- Handle no rows and iteration errors
- Scope every resource query to its owner

### Values are parameters, not SQL text

For PostgreSQL use placeholders such as $1 and pass values as separate arguments. The driver sends values safely; concatenating a title into SQL allows its text to become syntax. Placeholder conventions depend on the driver. Parameters are for values, not table or column names. If you need user-selectable sorting, map a small allowed set of sort choices to fixed SQL fragments.

### Read the lifecycle of a query

QueryRowContext followed by Scan reads one row; sql.ErrNoRows is the normal missing-record signal. QueryContext returns rows that must be closed, iterated, scanned, and checked with rows.Err afterward. Check every Scan error. Use a request context so cancellation can stop a query that is no longer useful. Never return a partial list as a silent success after iteration fails.

### Ownership belongs in the query

An authenticated user must not retrieve another person’s task by guessing its ID. Include owner_id in reads, updates, and deletes. For updates and deletes, inspect RowsAffected or use RETURNING to distinguish a missing owned row. Using 404 for an inaccessible record can avoid revealing whether another user’s record exists. Test with two different users, not only one happy-path account.

### Example

```
// Contextual repository function; imports context, database/sql, fmt.
type StoredTask struct { ID int64; Title string; Done bool }
func findTask(ctx context.Context, db *sql.DB, ownerID, id int64) (StoredTask, error) {
    var t StoredTask
    err := db.QueryRowContext(ctx,
        "SELECT id, title, done FROM tasks WHERE owner_id=$1 AND id=$2",
        ownerID, id).Scan(&t.ID, &t.Title, &t.Done)
    if err != nil { return StoredTask{}, fmt.Errorf("find task: %w", err) }
    return t, nil
}
// A PostgreSQL driver must be installed and registered by the application.
```

### Exercise

Write a parameterized rename query scoped to owner and task ID. Return sql.ErrNoRows when no owned row is updated.

Hint: Use UPDATE ... RETURNING id with QueryRowContext to keep the outcome explicit.

Review criteria:

- [ ] A title containing a quote stores as data
- [ ] Wrong owner cannot rename the task
- [ ] Missing and inaccessible IDs produce a no-rows outcome

### Worked solution

```
// Contextual: imports context and database/sql.
func renameTask(ctx context.Context, db *sql.DB, ownerID, id int64, title string) error {
    var updatedID int64
    return db.QueryRowContext(ctx,
        "UPDATE tasks SET title=$1 WHERE owner_id=$2 AND id=$3 RETURNING id",
        title, ownerID, id).Scan(&updatedID)
}
// Validate title before calling; the database constraint remains a final guard.
```

### Checkpoint

Can $1 safely parameterize a column name?

1. Yes
2. No; parameters represent values, so use an allowlist for SQL identifiers
3. Only in SELECT

Answer: 2. SQL parameters bind data values, not query structure. Choose identifiers and sort expressions from fixed, reviewed alternatives.

Sources:

- [Avoiding SQL injection](https://go.dev/doc/database/sql-injection)
- [Querying for data](https://go.dev/doc/database/querying)
- [database/sql package](https://pkg.go.dev/database/sql)

## Transactions and consistent updates

50 minutes of guided study, plus practice.

Learning objectives:

- Use sql.Tx for all transactional work
- Commit only after every step succeeds
- Understand isolation and concurrent changes

### A transaction is one logical change

When two writes must either both happen or neither happen, use a transaction. For example, completing a task and appending an audit event should not leave a completed task with no event if the second write fails. BeginTx returns a transaction object. Use its query and execution methods for every operation in that unit; calling db.Exec halfway through uses a separate connection outside the transaction.

### Rollback is the safe default

Defer a rollback immediately after a successful BeginTx. On any error, return and let rollback undo the work. Commit only after all steps succeed, and check the commit error. A rollback after a successful commit is harmless to attempt. Avoid long external network calls while holding an open transaction because they extend lock lifetimes and consume pooled connections.

### Concurrent writers change the story

A transaction alone does not make every read-modify-write rule safe. Understand your isolation level, use atomic SQL updates or row locking when needed, and design retry behavior for serialization failures or deadlocks. Keep retries bounded and make externally visible operations idempotent where possible. Verify failure and concurrency behavior against a real test database, because mocks cannot reproduce database locking.

### Example

```
// Contextual: imports context and database/sql; tasks and task_events exist.
func completeWithEvent(ctx context.Context, db *sql.DB, ownerID, taskID int64) error {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil { return err }
    defer tx.Rollback()
    var id int64
    err = tx.QueryRowContext(ctx,
        "UPDATE tasks SET done=true WHERE owner_id=$1 AND id=$2 RETURNING id",
        ownerID, taskID).Scan(&id)
    if err != nil { return err }
    _, err = tx.ExecContext(ctx,
        "INSERT INTO task_events(task_id, kind) VALUES ($1,$2)", id, "completed")
    if err != nil { return err }
    return tx.Commit()
}
```

### Exercise

Create task_events with a foreign key, then prove that forcing the event insert to fail leaves the task unchanged. Run this in a disposable test database.

Hint: Seed an incomplete task, trigger a deliberate event constraint failure inside the transaction, and query the task afterward.

Review criteria:

- [ ] Successful transaction changes task and adds one event
- [ ] Failed second write leaves both tables unchanged
- [ ] All operations use tx, and commit errors are returned

### Worked solution

```
-- Supporting migration:
CREATE TABLE task_events (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    task_id BIGINT NOT NULL REFERENCES tasks(id),
    kind TEXT NOT NULL CHECK (kind IN ('completed')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- Test strategy: temporarily call the function with an invalid event kind
-- through an injected test helper; expect an error, then assert done=false.
-- Keep failure injection out of the production API.
```

### Checkpoint

Which call belongs inside a transaction?

1. db.ExecContext for the second statement
2. tx.ExecContext for every statement in the unit
3. A new sql.Open for each statement

Answer: 2. All statements in the atomic unit must run through the same sql.Tx. Calling the pool directly runs outside that transaction.

Sources:

- [Executing transactions](https://go.dev/doc/database/execute-transactions)
- [PostgreSQL transaction isolation](https://www.postgresql.org/docs/current/transaction-iso.html)

## Project: Replace memory with PostgreSQL

Implement the repository interface with PostgreSQL while keeping the HTTP contract stable. Use a driver such as pgx stdlib and pin its version in your module.

Requirements:

- Versioned schema migrations
- Parameterized owner-scoped CRUD
- Context-aware queries and checked iteration errors
- Integration tests using a disposable database
- Pool configuration and environment-based credentials

Self-review rubric:

- [ ] Persistence and contract compatibility (35%)
- [ ] Ownership and SQL safety (35%)
- [ ] Migration and failure tests (30%)

# 9. Connect the whole experience

Frontend integration, sessions, and application security

Level: Full stack. Prerequisite: Add durable storage.

## A frontend that understands the API

50 minutes of guided study, plus practice.

Learning objectives:

- Call JSON endpoints with fetch
- Handle loading, empty, success, and failure states
- Render user data safely

### The browser is another client

Start with semantic HTML, CSS, and a small JavaScript module. A framework can help later, but it cannot replace an understood HTTP contract. A form submits a title to your Go API; a list renders returned tasks. Prefer serving the frontend and API from the same origin for this learning project, which avoids unnecessary CORS and cross-site cookie complexity.

### Network success is not HTTP success

fetch rejects for some network failures but normally resolves for HTTP 400 or 500 responses. Check response.ok before treating the result as success. Show an accessible loading message, disable duplicate submits while a write is pending, and restore controls in finally. Handle an empty list as a normal state and preserve the user’s input if a save fails.

### Treat displayed text as text

Use textContent for user-authored task titles instead of injecting them with innerHTML. A title containing HTML syntax must display as text, not execute. Use labels, keyboard-operable buttons, and live status messages. A well-connected full-stack app includes the slow and failed paths, not just a screenshot of a populated list.

### Example

```
// Browser JavaScript module; expects <ul id="tasks"> and <p id="status">.
async function loadTasks() {
    const status = document.querySelector('#status');
    const list = document.querySelector('#tasks');
    status.textContent = 'Loading tasks…';
    try {
        const response = await fetch('/api/tasks');
        if (!response.ok) throw new Error('Could not load tasks');
        const tasks = await response.json(); // contract: array of task objects
        list.replaceChildren();
        for (const task of tasks) {
            const li = document.createElement('li');
            li.textContent = task.title;
            list.append(li);
        }
        status.textContent = tasks.length ? '' : 'No tasks yet. Add your first one.';
    } catch (error) {
        status.textContent = error.message;
    }
}
```

### Exercise

Add a labeled form that POSTs {title} to /api/tasks, checks response.ok, clears the input only on success, and refreshes the list.

Hint: Prevent default form submission; disable the submit button in try and re-enable it in finally.

Review criteria:

- [ ] 400 and 500 responses show failure
- [ ] Failed saves preserve the title
- [ ] A title containing <img> displays literally
- [ ] Form works with a keyboard

### Worked solution

```
// Contextual browser JS: form has #title input and a submit button.
form.addEventListener('submit', async event => {
    event.preventDefault();
    const button = form.querySelector('button[type="submit"]');
    const input = form.querySelector('#title');
    button.disabled = true;
    try {
        const response = await fetch('/api/tasks', {
            method: 'POST', headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({title: input.value})
        });
        if (!response.ok) throw new Error('Save failed. Please try again.');
        input.value = '';
        await loadTasks();
    } catch (error) { document.querySelector('#status').textContent = error.message; }
    finally { button.disabled = false; }
});
// Add the CSRF protection described in the next lesson before session-authenticated deployment.
```

### Checkpoint

Does fetch reject its promise for every HTTP 500 response?

1. Yes
2. No; check response.ok
3. Only for JSON

Answer: 2. fetch normally resolves with a Response for HTTP errors. Check ok or status before using the response as a successful result.

Sources:

- [MDN: Using Fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch)
- [MDN: textContent](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent)

## Identity, sessions, and authorization

55 minutes of guided study, plus practice.

Learning objectives:

- Distinguish authentication from authorization
- Design server-side sessions with safe cookies
- Protect state changes from CSRF

### Identity is more than a login page

Authentication establishes who a user is. Authorization decides which actions that user may perform. A valid session for user A does not allow access to user B’s tasks. For a first production deployment, prefer a maintained identity provider rather than inventing password-reset, verification, and account-recovery systems. If you store passwords, use an appropriate password-hashing library and never plaintext or a fast unsalted hash.

### A session is a revocable server-side record

Generate unpredictable session tokens with crypto/rand, store a server-side record with user ID and expiration, and send the token in a cookie configured HttpOnly, Secure, Path=/, and an appropriate SameSite policy. Rotate the session at login, expire it on logout, and enforce expiration on the server. HttpOnly reduces token access through script but does not make an XSS vulnerability harmless.

### Protect writes and keep ownership server-side

Cookie authentication sends credentials automatically, so state-changing requests need CSRF defenses. Use a proven token pattern and origin checking suited to your deployment; SameSite is useful defense in depth, not a blanket replacement for a reviewed design. CORS is not authentication. Derive owner IDs from the validated session, use them in every repository query, and test one user attempting another user’s IDs.

### Example

```
// Contextual token generation; imports crypto/rand and encoding/base64.
func newSessionToken() (string, error) {
    raw := make([]byte, 32)
    if _, err := rand.Read(raw); err != nil { return "", err }
    return base64.RawURLEncoding.EncodeToString(raw), nil
}
// Contextual HTTPS cookie setup; import net/http and time.
// token must already have a server-side session record with an expiration.
http.SetCookie(w, &http.Cookie{
    Name: "__Host-session", Value: token, Path: "/",
    HttpOnly: true, Secure: true, SameSite: http.SameSiteLaxMode,
    Expires: time.Now().Add(8*time.Hour),
})
// This cookie fragment belongs inside a handler, not at package scope.
```

### Exercise

Write an authorization test with users A and B: create a task as A, then attempt GET, PATCH, and DELETE as B. Require every operation to fail without modifying A’s task.

Hint: A successful login is not the test’s end. The repository queries must include the authenticated owner ID.

Review criteria:

- [ ] Cross-user reads and writes are rejected
- [ ] No-session requests are rejected
- [ ] Logout invalidates the server-side session
- [ ] CSRF-negative tests cover state changes

### Worked solution

```
// Integration test plan:
// 1. Create two users and isolated sessions through test fixtures.
// 2. POST /api/tasks as A; record the returned task ID.
// 3. GET, PATCH, DELETE that ID as B; expect 404 (or your documented 403 policy).
// 4. GET as A; confirm the title and existence are unchanged.
// 5. Repeat with no session; expect 401.
// 6. Try a state-changing request without a valid CSRF token; require rejection.
```

### Checkpoint

What is authorization?

1. Knowing the user’s email
2. Checking whether that user may perform this action on this resource
3. Encrypting a cookie

Answer: 2. Authorization checks permission for the specific action and resource. It remains necessary after successful authentication.

Sources:

- [crypto/rand package](https://pkg.go.dev/crypto/rand)
- [net/http package](https://pkg.go.dev/net/http)
- [Go security best practices](https://go.dev/doc/security/best-practices)
- [MDN: Set-Cookie](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Set-Cookie)

## Security as everyday engineering

50 minutes of guided study, plus practice.

Learning objectives:

- Threat-model inputs and privileges
- Prevent XSS, injection, and resource abuse
- Scan dependencies and protect secrets

### Name what you are protecting

List assets such as account data, session tokens, database credentials, and service availability. Then list entry points: forms, JSON bodies, URL parameters, file uploads, webhooks, and outbound URLs. For each boundary ask who controls the input and which privilege the server has. This simple threat model guides concrete controls instead of adding security headers without understanding their purpose.

### Use the right control for each boundary

Parameterized SQL separates data from query syntax. html/template escapes data for HTML contexts; do not mark untrusted text as template.HTML. Limit request bodies, apply timeouts, and rate-limit expensive or authentication endpoints. If a service fetches arbitrary URLs, SSRF becomes a concern: avoid arbitrary fetch features or apply a reviewed destination policy and network restrictions. Do not assume URL parsing alone makes a destination safe.

### Keep the supply chain observable

Run govulncheck against the code you actually build and keep Go and dependencies updated through a reviewed process. A scanner is useful evidence, not a proof of safety. Keep secrets out of source control and browser bundles; rotate exposed credentials. Log request IDs and useful events without passwords, session tokens, or full sensitive payloads. Review error messages from the perspective of both a user and an attacker.

### Example

```
# Terminal checks from your Go module; review the tool version before pinning it in CI.
go test ./...
go vet ./...
go test -race ./...
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...

# Security regression inputs for your app:
# title containing '<script>alert(1)</script>' must render as text
# title containing a single quote must store as data
# oversized JSON must be rejected
# another user's task ID must not disclose or mutate their data
```

### Exercise

Write a one-page threat model for the task app. For each of five threats, record an entry point, impact, control, and verification test.

Hint: Include cross-user access, stored XSS, SQL injection, CSRF, and oversized requests.

Review criteria:

- [ ] Five concrete threats have testable controls
- [ ] Each control maps to actual code or configuration
- [ ] No secrets or session tokens appear in logs
- [ ] Vulnerability findings are triaged rather than blindly ignored

### Worked solution

```
Threat: cross-user task modification
Entry: PATCH /api/tasks/{id}
Impact: unauthorized data changes
Control: authenticated owner ID included in UPDATE predicate
Test: user B cannot modify user A's task

Repeat the same four fields for XSS (textContent/templates), injection
(parameters), CSRF (token/origin checks), and resource abuse (limits/timeouts).
```

### Checkpoint

Which is safe for rendering an untrusted task title?

1. Assign it to innerHTML
2. Set textContent or pass it as data to html/template
3. Cast it to template.HTML

Answer: 2. textContent treats the value as text. html/template provides contextual escaping when data stays data; bypassing it with template.HTML removes that protection.

Sources:

- [Security best practices](https://go.dev/doc/security/best-practices)
- [html/template package](https://pkg.go.dev/html/template)
- [govulncheck](https://go.dev/doc/security/vuln/)

## Project: Build the complete task-board experience

Connect a responsive browser UI to your Go/PostgreSQL API. Add a maintained authentication integration or a carefully scoped server-session implementation with documented limitations.

Requirements:

- Loading, empty, error, and success UI states
- Accessible forms and safe text rendering
- Session expiration, logout, and CSRF protection
- Ownership checks on every resource operation
- End-to-end two-user isolation test and a threat model

Self-review rubric:

- [ ] Working end-to-end behavior (35%)
- [ ] Identity and security controls (40%)
- [ ] Accessibility and failure recovery (25%)

# 10. From a project to a service

Deployment, observability, and your capstone

Level: Full stack. Prerequisite: Connect the whole experience.

## Build, configure, and shut down cleanly

50 minutes of guided study, plus practice.

Learning objectives:

- Separate build artifacts from runtime configuration
- Handle termination with graceful shutdown
- Plan migrations, rollback, and backups

### Build once, configure deliberately

go build produces a binary for a target platform; dependencies such as CGO affect packaging decisions. Keep runtime configuration such as listen address and database URL outside the binary. Validate required settings at startup and fail with a useful, non-secret error. Run as a non-root user where supported, expose only required ports, and terminate HTTPS at a trusted edge or configure it explicitly.

### Shutdown is part of correctness

On termination, stop accepting new work and give active requests a bounded time to finish. signal.NotifyContext can translate an operating-system signal into cancellation. http.Server.Shutdown waits for ordinary active connections within the supplied context, but long-lived or hijacked connections need their own handling. Close the database pool after dependent requests finish. Treat http.ErrServerClosed as expected on shutdown.

### Recover before you need recovery

A deploy plan includes schema compatibility, health checks, rollback, and data backup restoration. Test restoration, not just backup creation. Avoid a destructive migration in the same step that deploys code requiring the old schema. The learning app can start with a single instance and a managed PostgreSQL database; multi-region systems introduce new failure modes and are a later design choice.

### Example

```
// Contextual graceful-shutdown pattern inside main.
// Imports context, errors, log, net/http, os/signal, syscall, time.
ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
defer stop()
server := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 5*time.Second}
errCh := make(chan error, 1)
go func() { errCh <- server.ListenAndServe() }()
select {
case err := <-errCh:
    if !errors.Is(err, http.ErrServerClosed) { log.Printf("server failed: %v", err) }
case <-ctx.Done():
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := server.Shutdown(shutdownCtx); err != nil {
        log.Printf("shutdown: %v", err)
        _ = server.Close()
    }
}
// Close the database pool after this point.
```

### Exercise

Write a runbook with startup requirements, a health-check URL, deployment steps, rollback criteria, and a tested database restore procedure.

Hint: State the exact signals and commands for your chosen host; do not assume every platform has the same process model.

Review criteria:

- [ ] A fresh machine can follow the runbook
- [ ] Termination completes within the configured budget
- [ ] Restore has been exercised against a separate database
- [ ] Secrets do not appear in artifacts or logs

### Worked solution

```
Runbook outline:
1. Required environment variables and secret provisioning (names only).
2. Build and test the exact commit; record artifact identity.
3. Back up data and apply compatible migrations deliberately.
4. Deploy, then verify health and one authenticated CRUD flow.
5. Roll back on defined error-rate or data-integrity failures.
6. Restore a backup into a separate database and verify record counts.
7. Record owners, timestamps, and known recovery limitations.
```

### Checkpoint

When should the database pool close during shutdown?

1. Before active requests finish
2. After dependent requests have finished or been stopped
3. After every SQL statement

Answer: 2. Requests still using the database need the pool. Stop and drain request work before closing shared dependencies.

Sources:

- [net/http package](https://pkg.go.dev/net/http)
- [os/signal package](https://pkg.go.dev/os/signal)
- [Go command](https://pkg.go.dev/cmd/go)

## Observe, diagnose, and improve

45 minutes of guided study, plus practice.

Learning objectives:

- Use logs and metrics to explain failures
- Profile before changing performance-sensitive code
- Set useful health and readiness checks

### Observe behavior users care about

Track request count, latency, error rate, and dependency failures. A request ID joins a client-facing error to a log event without exposing internal details. Use structured logging such as log/slog and avoid high-cardinality fields such as arbitrary URLs or user IDs as metric labels. A dashboard should help answer whether users can complete their tasks, not merely whether a process exists.

### Separate health from readiness

A liveness check asks whether the process is alive; a readiness check asks whether it should receive traffic. A database-dependent readiness check may be appropriate, but avoid overloading dependencies with aggressive health probes. Do not expose internal diagnostics publicly. Protect pprof endpoints with network and access restrictions because profiles and runtime details can reveal sensitive information.

### Measure a specific bottleneck

Use CPU and heap profiles to find expensive work in a realistic workload. Inspect allocations, database query latency, connection-pool waits, and slow external calls before adding caches or goroutines. The garbage collector guide and diagnostics documentation are advanced references for a reason: tune from evidence and rerun correctness checks after changes. A faster wrong answer is still a bug.

### Example

```
// Contextual handler wrapper; imports log/slog, net/http, time.
func logDuration(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        slog.Info("request complete", "method", r.Method,
            "duration_ms", time.Since(start).Milliseconds())
    })
}
// This minimal wrapper measures duration only. To log status and bytes,
// use a tested response recorder preserving required optional interfaces.
// Do not log authorization headers, cookies, or raw sensitive bodies.
```

### Exercise

Create a small load experiment: record baseline latency and errors, identify one bottleneck, change one thing, and compare under the same workload.

Hint: Use a fixed dataset, concurrency, warm-up, and measurement period. Record the environment so the result can be interpreted.

Review criteria:

- [ ] Before and after measurements are comparable
- [ ] Correctness and security tests still pass
- [ ] The result includes limitations, not only a favorable number

### Worked solution

```
Experiment record:
Hypothesis: task-list latency comes from fetching unbounded rows.
Baseline: fixed dataset, fixed request concurrency, p50/p95/error rate.
Change: add owner-scoped pagination with a maximum page size.
Correctness: contract and ownership tests still pass.
Comparison: same workload and environment; repeat several runs.
Decision: keep the change only if behavior and measured evidence support it.
```

### Checkpoint

What should happen before adding a performance optimization?

1. Guess which line looks slow
2. Measure a realistic workload and identify a bottleneck
3. Disable error checking

Answer: 2. Profiles and workload measurements help identify actual bottlenecks. Optimizing by appearance can add complexity without improving user-visible behavior.

Sources:

- [Diagnostics](https://go.dev/doc/diagnostics)
- [Garbage collector guide](https://go.dev/doc/gc-guide)
- [log/slog package](https://pkg.go.dev/log/slog)

## Capstone: a trustworthy full-stack task board

90 minutes of guided study, plus practice.

Learning objectives:

- Integrate Go, PostgreSQL, and a browser UI
- Demonstrate tested security and recovery behavior
- Explain architecture and tradeoffs clearly

### Build a vertical slice first

Your capstone is a multi-user task board with authenticated task creation, listing, completion, rename, and deletion. Use a Go net/http API, PostgreSQL via database/sql with a maintained driver, and a same-origin semantic HTML/CSS/JavaScript frontend. Deliver one end-to-end create-and-list flow before expanding features. This stack keeps the boundaries visible while teaching skills transferable to React, Vue, and other clients.

### Prove the hard paths

A complete demo includes malformed JSON, an unavailable database, an expired session, an unauthorized resource ID, a failed frontend save, and a graceful shutdown. Add unit tests for domain rules, HTTP tests for contract behavior, database integration tests for constraints and transactions, and an end-to-end browser scenario with two users. Run race detection and vulnerability scanning where supported, and document any remaining limitations honestly.

### Present the evidence

Your README should let another person run the app from a clean checkout. Include the architecture, schema and migration instructions, configuration names, test commands, threat model, API examples, deployment runbook, backup restoration evidence, and a short design discussion. Explain why you chose sessions, how ownership is enforced, and what happens during partial failure. Shipping means another person can operate and assess your work.

### Example

```
# Suggested source layout, not a generated application.
cmd/server/main.go          # configuration and dependency wiring
internal/task/              # domain rules and consumer-facing contracts
internal/postgres/          # SQL repository implementation
internal/httpapi/           # routes, middleware, JSON contracts
migrations/                # versioned PostgreSQL schema changes
web/                       # semantic HTML, CSS, browser JavaScript
docs/api.md                # request/response examples and error policy
docs/threat-model.md        # threats, controls, and verification
docs/runbook.md             # deploy, rollback, restore

# Reproducible checks:
go test ./...
go vet ./...
go test -race ./...
govulncheck ./...
```

### Exercise

Complete a self-review using the capstone rubric below. Record one concrete piece of evidence for each criterion and fix the weakest area before presenting.

Hint: Screenshots show appearance; tests, API traces, restore logs, and a clean-checkout run show behavior.

Review criteria:

- [ ] Every criterion has linked evidence
- [ ] No hard-coded production secrets or demo-only authentication remain
- [ ] A second person can run and assess the app
- [ ] You can explain one rejected design alternative and why

### Worked solution

```
Evidence checklist:
[ ] Fresh-checkout run follows README without undocumented steps.
[ ] CRUD survives a service restart.
[ ] Two-user tests prove isolation on reads and every write.
[ ] Invalid/oversized input, CSRF, and session expiry are handled.
[ ] UI covers keyboard, loading, empty, success, and failed saves.
[ ] Unit, HTTP, database, and end-to-end checks cover distinct risks.
[ ] Shutdown and backup restoration are demonstrated.
[ ] Design tradeoffs and remaining limitations are explicit.
```

### Checkpoint

Which is the strongest evidence that ownership is enforced?

1. The UI hides other users’ tasks
2. Two-user API tests attempt every read and write using another user’s ID
3. Task IDs are difficult to guess

Answer: 2. Authorization must hold at the server boundary even when clients bypass the UI or know a resource ID.

Sources:

- [Go documentation hub](https://go.dev/doc/)
- [Database documentation](https://go.dev/doc/database/)
- [Security best practices](https://go.dev/doc/security/best-practices)

## Project: Capstone: Taskboard, from source to service

Deliver an operable multi-user task board with a Go API, PostgreSQL persistence, same-origin frontend, and a documented security and deployment story. Allocate several days beyond the reading time to build and revise it.

Requirements:

- All task CRUD persists across restart
- Authenticated owner-scoped access with CSRF defenses
- Accessible frontend and complete failure states
- Domain, HTTP, database, and two-user end-to-end tests
- Versioned migrations, vulnerability review, graceful shutdown
- Clean-checkout README, threat model, API contract, and restore-tested runbook

Self-review rubric:

- [ ] Functionality and persistence (25%)
- [ ] Security and isolation (25%)
- [ ] Testing and failure recovery (20%)
- [ ] Frontend accessibility and clarity (15%)
- [ ] Operability and explanation of tradeoffs (15%)

## Coverage and learning records

The core path teaches the language and a practical web stack; specialist standard-library packages remain further reference tracks. See sources.json and devdocs-topics.json for the audit. The app saves progress on your device. Export a backup to move or preserve learning records. Code exercises require local execution or the official Go Playground; project rubrics are self-assessments.
