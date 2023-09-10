# auxi
GO (Golang) library for HTTP method-specific handler, middleware group chaining, colored log message, better HTTP response generator, and handler utils

This library is a minimal abstraction on top of vanilla libraries to include features such as:
- HTTP method-specific handler (GET, POST, PUT, DELETE, OPTIONS, etc)
- Middleware chain group creator to apply multiple middlewares to any handler
- Colored log message generator,
- Complete JSON response function and JSON HTTP response generator
- Middleware utils such as for grabbing bearer token from auth header
- Bind function to bind Query String from URL to a struct
- Template CORS middleware
