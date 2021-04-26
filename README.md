# Sezzle Backend Challenge

This code challenge is an opportunity to demonstrate proficiency in Sezzle's backend technologies and the problem solving skills we expect you to use.

This project is scoped to take less than 1 hour to complete. We respect your time, and ask that you not spend any more than 3 hours on this project. If you find yourself pressed for time, please document what you would do in your pull request description.

## How to Submit the Exercise

Send a pull request to the Github repository provided, with your changes. Please send us back a link to your open pull request to review your changes.

## Problem

We'd like you to add a calculator API to this project that enables a user to POST calculations, and GET a list of historical calculations. To do this, you will need to implement both GET and POST implementations for a `/calculations` endpoint.

Assuming this API is running on `http://localhost:8000`, the following should work:

```
$ curl -XPOST -H "content-type: application/json" -d'{"a": 1, "b": 2, "op": "+"}' http://localhost:8000/calculations
{"result": 3}
$ curl -XGET http://localhost:8000/calculations
[{"result": 3}]
```

## Technologies Used

* [Go programming language](https://golang.org/)
* [Gin web framework](https://pkg.go.dev/github.com/gin-gonic/gin?utm_source=godoc)

## Review Criteria

Your pull request will be reviewed by at least 2 engineers. They will consider the following critera when reviewing:

* **Correctness**: Is the implementation correct? Does the provided test suite pass?
* **Code Quality**: Is the code easy to understand and maintain?
