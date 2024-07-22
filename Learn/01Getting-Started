## **Golang Basics**

### **Setting Up Your Environment**
1. **Install Go**: Download and install Go from [golang.org](https://golang.org/dl/).
2. **Install VS Code**: Download and install Visual Studio Code from [code.visualstudio.com](https://code.visualstudio.com/).
3. **Connect GitHub**: Set up Git and connect your project to a GitHub repository for backup.

### **Initializing a Go Module**
A Go module is a collection of related Go packages that are versioned together as a unit. To start a new module:
```bash
go mod init hello
```
This creates a `go.mod` file which tracks the module's dependencies.

### **Writing and Running Your First Go Program**
Create a file named `main.go` and write a simple "Hello, World!" program:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
Run the program using:
```bash
go run main.go
```

### **Basic Go Concepts**

#### **Variables and Types**
Go is statically typed, and variables need to be declared with their type. 

**Example:**
```go
package main

import "fmt"

func main() {
    var name string = "Mayank"
    var age int = 30
    var height float64 = 5.9
    var isStudent bool = false

    city := "Mumbai" // Short variable declaration with type inference

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)
    fmt.Println("City:", city)
}
```

#### **Functions**
Functions in Go are defined using the `func` keyword. 

**Example:**
```go
package main

import "fmt"

func concatenate(a string, b string) string {
    return a + " " + b
}

func main() {
    result := concatenate("Hello", "World")
    fmt.Println("Concatenated string:", result)
}
```

#### **Control Structures**
Go has standard control structures like `if`, `for`, and `switch`.

**Example:**
```go
package main

import "fmt"

func main() {
    number := 10

    // If-else statement
    if number%2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }

    // For loop
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // Switch statement
    day := "Tuesday"
    switch day {
    case "Monday":
        fmt.Println("Start of the work week")
    case "Friday":
        fmt.Println("End of the work week")
    default:
        fmt.Println("Midweek day")
    }
}
```

### **Exploring Go Commands**

1. **`go env`**: Prints Go environment information.
   - Usage:
     ```bash
     go env
     ```
   - Example:
     ```bash
     go env GOPATH
     ```

2. **`go mod`**: Manages Go modules.
   - Common subcommands:
     - `init`: Initializes a new module.
       ```bash
       go mod init <module-name>
       ```
     - `tidy`: Adds missing and removes unused modules.
       ```bash
       go mod tidy
       ```
     - `download`: Downloads modules to the cache.
       ```bash
       go mod download
       ```
     - `graph`: Prints the module dependency graph.
       ```bash
       go mod graph
       ```
     - `verify`: Verifies module dependencies.
       ```bash
       go mod verify
       ```

3. **`go build`**: Compiles packages and dependencies.
   - Usage:
     ```bash
     go build [packages]
     ```
   - Example:
     ```bash
     go build
     ```

4. **`go run`**: Compiles and runs Go programs.
   - Usage:
     ```bash
     go run <file.go>
     ```
   - Example:
     ```bash
     go run main.go
     ```

5. **`go fmt`**: Formats Go source code according to Go standards.
   - Usage:
     ```bash
     go fmt [files]
     ```
   - Example:
     ```bash
     go fmt main.go
     ```

6. **`go test`**: Runs tests for Go packages.
   - Usage:
     ```bash
     go test [packages]
     ```
   - Example:
     ```bash
     go test ./...
     ```

7. **`go get`**: Downloads and installs packages and dependencies.
   - Usage:
     ```bash
     go get [packages]
     ```
   - Example:
     ```bash
     go get github.com/gin-gonic/gin
     ```

8. **`go install`**: Compiles and installs packages and dependencies.
   - Usage:
     ```bash
     go install [packages]
     ```
   - Example:
     ```bash
     go install
     ```

9. **`go clean`**: Removes object files and cached files.
   - Usage:
     ```bash
     go clean [flags] [packages]
     ```
   - Example:
     ```bash
     go clean -i
     ```

10. **`go list`**: Lists packages or modules.
    - Usage:
      ```bash
      go list [flags] [packages]
      ```
    - Example:
      ```bash
      go list -m all
      ```

### **Practice Exercises**

1. **Initialize a New Module**:
   ```bash
   go mod init mymodule
   ```

2. **Create and Run a Simple Program**:
   ```go
   // main.go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, Go!")
   }
   ```
   ```bash
   go run main.go
   ```

3. **Format the Code**:
   ```bash
   go fmt main.go
   ```

4. **Build the Program**:
   ```bash
   go build
   ```

5. **Run Tests** (create a simple test file first):
   ```go
   // main_test.go
   package main

   import "testing"

   func TestMain(t *testing.T) {
       expected := "Hello, Go!"
       if expected != "Hello, Go!" {
           t.Errorf("Expected %s, but got %s", expected, "Hello, Go!")
       }
   }
   ```
   ```bash
   go test
   ```
