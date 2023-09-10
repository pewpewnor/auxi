# auxi
GO (Golang) library for HTTP method-specific handler, middleware group chaining, colored log message, better HTTP response generator, and handler utils

## Features

This library is a minimal abstraction on top of vanilla libraries to include features such as:
- HTTP method-specific handler (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS)
- Middleware chain group creator to apply multiple middlewares to any handler
- Colored log message generator,
- Complete JSON response function and JSON HTTP response generator
- Middleware utils such as for grabbing bearer token from auth header
- Bind function to bind Query String from URL to a struct
- Template CORS middleware

## Examples

- HTTP method-specific handler example:

```go
// Declare API handlers
func getHandlerForUser(w http.ResponseWriter, r *http.Request) {...}
func postHandlerForUser(w http.ResponseWriter, r *http.Request) {...}
func putHandlerForUser(w http.ResponseWriter, r *http.Request) {...}
func deleteHandlerForUser(w http.ResponseWriter, r *http.Request) {...}

func main() {
	mux := auxi.NewServeMux() // Create a new router called mux

	// Specify which API handler responds to a particular HTTP method 
	mux.HandleMethods("/user", auxi.MethodHandlers{
		GET: getHandlerForUser,
		POST: postHandlerForUser,
		PUT: putHandlerForUser,
		DELETE: deleteHandlerForUser,
	})

	server := &http.Server{
		Addr:    ":8080", // Change to your desired port
		Handler: mux,     // Use ausi ServeMux
	}
	
	err := server.ListenAndServe() // start server
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```

- Middleware chain group handler example:

```go
func getHandlerForUser(w http.ResponseWriter, r *http.Request) {...}
func postHandlerForUser(w http.ResponseWriter, r *http.Request) {...}

// Declare Middlewares
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {...}
func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {...}

func main() {
	mux := auxi.NewServeMux()

	// Create a middleware chain which applies CORSMiddleware and then AuthMiddleware
	group := auxi.NewChain(CORSMiddleware, AuthMiddleware)

	v1 := auxi.NewServeMux() // Create a new router called v1

	v1.HandleMethods("/user", auxi.MethodHandlers{
		GET: group.Apply(getHandlerForUser), // Applies middleware chain 'group' to the GET handler
		POST: CORSMiddleware(postHandlerForUser), // POST handler only uses CORS middleware
	})

	mux.Handle("/v1/", http.StripPrefix("/v1", v1))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```
