## 📁 File structure

```
.
└── project_root/
    ├── internal/
    │   ├── database/
    │   │   └── connection.go (mongo driver)
    │   │   └── models.go (data models)    
    │   └── server/
    │       ├── http.go (router/cors middleware)
    │       ├── routes.go (routes/paths)
    │       └── views.go (view functions)
    ├── main.go (starts server)
    ├── go.mod 
    ├── go.sum
    ├── .gitignore
    ├── README.md
    └── LICENSE
```
## Architechture

A Caddy Proxy has been set up to 
    
- Map the go backend API to /api
- Map the nextjs frontend to /
- (Optionally) Load balance api with goroutine workers

## Contributing

This project is open to contributions, fork this repository and send a pull request incase you wish
to contribute.

## License

This project is licensed under the Apache License, Version 2.0