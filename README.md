# Task README

## Running & Building application

**Download dependencies with:**

```shell
$ go mod tidy && go mod vendor
```

**Run application with:**

```shell
$ go run main.go
```

**Build application with:**

```shell
$ go build -o xquestions main.go
```

**Access the GraphiQL interface:**
Open browser to `http://localhost:{PORT}/graphiql`

![Screenshot 2021-03-28 at 15 50 48](https://user-images.githubusercontent.com/9336187/112757015-adb91900-8fdf-11eb-95cf-9954ec6b693e.png)

## Running Tests

In the root of the application, run `go test ./...` to run all test suites

## GRAPHQL Definitions

### CalculatePrice Query

**Definitions**

- _type:string_
- _margin:string_

**Query**

```gql
query CalculatePrice {
  calculatePrice(type: "buy", margin: 0.2, exchangeRate: 476)
}
```

**Response**

```json
{
  "data": {
    "calculatePrice": 26467774.368
  }
}
```
