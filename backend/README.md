

```
/techvit/gin
  /cmd
    /main.go          // Application entry point
  /internal
    /domain
      /product        // Product domain
        /model.go     // Domain models
        /service.go   // Domain services
        /repository.go// Repository interfaces
    /application
      /service.go     // Application services, orchestrating domain services
    /infrastructure
      /repository     // Implementations of repository interfaces
        /product.go   // Product repository implementation
    /ui
      /handler        // Web handlers, REST API
        /product.go   // Product handlers
  /pkg
    /config           // Configuration management
    /logger           // Logging functionality
```

