# API Documentation

## Folder Structure

- `/api/v<version-number>`
  - `/routes` - All API routes are defined here
  - `/handlers` - For validating requests and calling services
  - `/services` - For business logic, database calls and other services
  - `/middlewares` - For authentication, logging, rate limiting etc.
  -
- `/cmd` - Initializes the GO HTTP app and basic middlewares configuration
- `/config` - For handling configuration/env variables
- `/internal` - For holding internal functions to/for applicatio
- `/utils` - For utility functions
- `/server` - Acts as serving point of the api application
-
- `version.go` - For Version of current app
- `main.go` - Entrypoint of the app

## Adding a New Route

To add a new route to the API engine service, follow these steps:

1. Create a new file in the `routes/` directory with the name of your route, e.g., `someroutename.go`.
2. Define the route handler function in the new file, e.g.:

   ```go
   package routes

   import (
   	"go-api-boilerplate/api/v1/handlers"

   	"github.com/gin-gonic/gin"
   )

   // SomeNameRoute is a function that returns a value for route
   func SomeNameRoute(r *gin.RouterGroup) {
   	r.GET("/someroute", handlers.someroutenamegethandler)
   	r.POST("/someroute", handlers.someroutenameposthandler)
   }
   ```

3. Once done add the new route to `routes/base.go` for respective version files if any under `api\` folder, e.g.:

   ```go
   package routes

   import (
   	"github.com/gin-gonic/gin"
   )

   // BaseRoute is the base route for the API
   func SetupVer1Route(r *gin.RouterGroup) {
   	v1 := r.Group("/v1")

   	// Version route
   	VersionRoute(v1) // this is already in boiler plate code

   	// Add more routes here
       SomeNameRoute(v1)
   }
   ```

4. Similarly if want to add new version of apis it can be added as new folder under `api\v<versionnumber>`, e.g.: `api\v2`, `api\v3`. This should follow similar sub folder structure as in sample boilerplate v1
5. Import new version to existing routes so it can be served. Add new file under new version folder `routes\base.go` e.g.:

   ```go
   package routes

   import (
   	"github.com/gin-gonic/gin"
   )

   // BaseRoute is the base route for the API as v2
   func SetupVer2Route(r *gin.RouterGroup) {
   	v2 := r.Group("/v2")

   	// Version route
   	VersionRoute(v2)
   	// Add more routes here
   }
   ```
