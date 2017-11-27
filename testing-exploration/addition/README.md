## Some commands used to launch the tests in this package

`go test addition`

`go test addition -v`

`go test addition -coverprofile=cover.out`
- Generates the output file we will use for the next commands

`go tool cover -func=cover.out`

`go tool cover -html=cover.out`
- Opens up the browser and shows exactly which lines are not covered

