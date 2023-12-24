
# Go (Golang)

Go is a statically typed, compiled, open source programming language designed by Google and first
appeared in 2009.

Go is very popular for server-side projects, but it is also very popular for writing Command
Line Interfaces (CLIs) because the binaries can be statically linked.

Compiled - code is compiled ahead of execution.

Statically typed - value types are defined in advance.

Go is especially popular for building scalable web servers but also great for automation scripts and command-line programs.
Not good at games and desktop apps.

Go tries to combine the best of C++, Java and Python.

Here are some projects that come to mind, which are entirely written in Go:

- Docker - a container engine written by Docker Inc
- Kubernetes - a clustering system for containers built by Google and open source contributors
- Terraform, Consul - tools for managing infrastructure built by HashiCorp
- Traefik - a load-balancer built by Traefik Labs
- Caddy - a high performance HTTP server written by Matt Holt
- OpenFaaS - serverless framework for hosting your own functions, built by OpenFaaS Ltd

 Go has an impressive set of packages available within its standard libraries or `stdlib`. There are
many useful packages and utilities that would need to be installed separately in a programming language like JavaScript,
Ruby or Java.

- `crypto` - libraries for X.509 certificates and cryptography
- `compress` - work with zip files and archives
- `http` - a very powerful and simple HTTP client and server package including things like reverse proxies
- `net` - work directly with sockets, URLS and DNS
- `encoding/json` - work directly with JSON files
- `text/template` - a powerful templating engine for replacing tokens in files, and generating text and HTML
- `os` - work with low level OS primitives

## Data Types

Go (or Golang) is a statically typed language, meaning that variables must be declared with a specific type and that type cannot change during the execution of the program. Here are some of the basic data types in Go:

1. Numeric Types:

    - `int`: Signed integer types (int8, int16, int32, int64)
    - `uint`: Unsigned integer types (uint8, uint16, uint32, uint64)
    - `float32` and `float64`: Floating-point types
 
2. String Type:

    - `string`: A sequence of characters

3. Boolean Type:

    - `bool`: Boolean values (true or false)

4. Composite Types:

    - `array`: A fixed-size sequence of elements of the same type
    - `slice`: A dynamically-sized, flexible view into the elements of an array
    - `map`: A collection of key-value pairs
    - `struct`: A composite data type that groups together variables under a single name

5. Pointer Types:

    - `pointer`: A variable that stores the memory address of another variable

6. Function Types:

    - `func`: A function type denotes the set of all functions with the same parameter and result types.

7. Interface Types:

    - `interface`: A set of method signatures. Interfaces allow you to define a contract for types that implement the methods.

8. Channel Types:

    - `chan`: A communication channel used for goroutine communication and synchronization

---

Books: Everyday Golang - The Fast Track
