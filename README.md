# testing-mongodb-with-gomock
This is a demonstration of mocking MongoDB with [gomock](https://github.com/golang/mock).

## Get the code and run the test :)

**Get it:** `go get -u github.com/golovers/testing-mongodb-with-gomock`

**Run it**: `go test -v pkg/users/*`
**Result**:

```
=== RUN   TestMongoDB
=== PAUSE TestMongoDB
=== RUN   TestAddUser
=== PAUSE TestAddUser
=== RUN   TestGetUserEmptyResponse
=== PAUSE TestGetUserEmptyResponse
=== RUN   TestGetUserNonEmptyResponse
=== PAUSE TestGetUserNonEmptyResponse
=== CONT  TestMongoDB
=== CONT  TestGetUserNonEmptyResponse
=== CONT  TestGetUserEmptyResponse
--- PASS: TestGetUserNonEmptyResponse (0.00s)
--- PASS: TestMongoDB (0.00s)
=== CONT  TestAddUser
--- PASS: TestAddUser (0.00s)
--- PASS: TestGetUserEmptyResponse (0.00s)
PASS
ok      command-line-arguments  (cached)
```
