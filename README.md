### Dependencies:
    Go Version == 1.23.1
    Postgres >= 14
### Setup:
    - Modify below config.yaml params to your DB credentials
    ```
    db:
    host: localhost
    user : admin
    password: admin
    name: transfers
    port : 5432 
    ```
    - Run the application using `go run main.go`
