# Product API

## Implements:
-- Full testing
-- Modeling
-- Gorilla/Mux

## Clone the repo
```bash
$ git clone github.com/DLzer/go-product-api
```

## Set your default PostreSQL variables
```bash
export APP_DB_USERNAME=postgres
export APP_DB_PASSWORD=
export APP_DB_NAME=postgres
```

## Running the tests
Run the test suite with the following commands
```bash
$ go test -v
```
The expected output should be:
```bash
=== RUN   TestEmptyTable
--- PASS: TestEmptyTable (0.01s)
=== RUN   TestGetNonExistentProduct
--- PASS: TestGetNonExistentProduct (0.00s)
=== RUN   TestCreateProduct
--- PASS: TestCreateProduct (0.01s)
=== RUN   TestGetProduct
--- PASS: TestGetProduct (0.01s)
=== RUN   TestUpdateProduct
--- PASS: TestUpdateProduct (0.01s)
=== RUN   TestDeleteProduct
--- PASS: TestDeleteProduct (0.01s)
PASS
ok      github.com/DLzer/go-product-api       0.071s
```

## Running an instance of PostgreSQL
```bash
$ docker run -it -p 5432:5432 -d postgres
```

