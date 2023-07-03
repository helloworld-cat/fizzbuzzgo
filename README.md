## Spec.

Write a REST API that prints the numbers from 1 to 200.
For multiples of seven print "Fizz" instead of the number and for the multiples of nine print "Buzz". For numbers which are multiples of both seven and nine print "FizzBuzz".

Requirements:
- Allow user to define the limit of numbers that he wants to display (ex: 100).
- Allow user to define the two multiples numbers he wants to override (ex: 3 and 5).
- Allow user to define the two words he wants to display (ex: fizz and buzz).

## Usages

### Development mode

```
$ go run cmds/fizzbuzz/main.go
$ curl ... # see routes examples
```

### Run tests

```
$ go test ./...
```

## Available routes

### List numbers

- `POST /numbers`

#### HTTP Request

```
$ curl -i -X POST -d '{"number_a": 7, "number_b": 9, "word_a": "Fizz", "word_b": "Buzz", "limit": 100}' "http://localhost:8080/numbers"

@limit: optional
@word_a: optional, warning: it is case sensitive
@word_b: optional, warning: it is case sensitive
```

#### Response

```
HTTP/1.1 200 OK

[
  1,
  2,
  // ....
  6,
  "Fizz",
  8,
  "Buzz",
  10,
  // ....
  13,
  "Fizz",
  // ....
  62,
  "FizzBuzz",
  // ...
  100
]%
```

### Fetch stats

- `POST /stats`

#### HTTP Request

```
$ curl -i -X POST -d '{"number_a": 7, "number_b": 9, "word_a": "Fizz", "word_b": "Buzz"}' "http://localhost:8080/stats"

@word_a: optional, warning: it is case sensitive
@word_b: optional, warning: it is case sensitive
```

#### Response

```
HTTP/1.1 200 OK
{
  // ... parameters sent
  "stats": 1 // number of calls about parameters
}%
```

### More documentation ?

See `*_test.go` files :)
