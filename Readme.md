# performance-tests-with-go

This repo is to run performance tests for any api.

To start test follow steps:
1. Add api details in main.go by changing host, api path, request payload if any
2. Update TPS (Transaction Per Second) in main.go to the desired performance rate.
3. Run with `go run main.go`