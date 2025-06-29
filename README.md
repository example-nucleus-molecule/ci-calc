# ci-calc

A simple command line tool to calculate the return on an investment with interest compounding monthly/quarterly/annual/at maturity.

## Dependencies

Go `v1.24.4`. It can be downloaded [here](https://go.dev/dl/). This will likely work fine on older Go versions, but hasn't been tested.

## Usage

Just clone this repo and in the root directory run `go run main.go`. If you wish to compile the application into a binary run `go build` which will create an executable named `ci-calc` (or `ci-calc.exe` if using Windows).

Note that Go dependencies are included in the `vendor` directory, the `go` command should detect these and use them without having to install anything separately. If not try running `go mod download`.

## Tests

Run all tests from the project root with `go test ./...`.

## Notes on the Code

The basic structure I have kept very simple, chucking the calculator logic in one package and creating another for the CLI code (conceivably other interfaces could be added that call the underlying calculator, a REST API or whatever). In my experience idomatic Go emphasises simplicity even at the sake of verbosity, so I have avoided searching for any abstractions at this point. If more features are added (e.g., accounting for regular deposits, extracting the interest payments only, etc.) perhaps some sensible abstractions would make themselves apparent, but in general I prefer to avoid premature optimisations.

Idiomatic Go also tends to heavily emphasise the standard library, although in this case I have imported a simple-ish lib to help with the CLI input. It makes the experience better for the user and provides a nice interface for validation of string input, and I don't think it would have been worth the time investment to re-invent this with the standard lib.

When working with floats, Go defaults everything to double precision (64-bit) and follows [IEEE 754](https://en.wikipedia.org/wiki/IEEE_754), which I have just kind of assumed will work the way I expect. If this were a 'real' app I'd spend a lot more time confirming that assumption, and testing edge cases, but it does seem to work fine in my simple 'happy path' tests.

The input validation is pretty rudimentary. It could be extended to disallow negatives or extremely large numbers (possibly the input lib I used already handles this okay, I didn't actually test it...), or to make the experience nicer by allowing input like "$1,250" without erroring. Also, providing constants for the limited 'interest frequency' options helps keep things simple with validation, although it would be nice if Go had a proper `enum` type for this to provide additional safety.

The tests are also very basic and could use a lot more edge cases. And normally I would import a package to make the assertions look a bit nicer (like `testify/assert`) but it didn't seem necessary here. I also ran out of time to add any tests for the CLI package!

Regarding error handling, a common bugbear of Go users, again I've kept it super simple by relying on the type system to avoid errors in the first place, and hence the `ci-calc` package can't actually return an error. If it needed to be expanded and errors _were_ possible, I'd look to add custom error types so that callers can have more control over how they handle errors and a stack trace can be provided. The only place errors _can_ be returned at present is from the 3rd party `promptui` package, although in the case of errors from the `Run` function these are unexpected and so will just kill the app, and from the validate function the prompt will be retried until valid input is received. So, again, nothing fancy required here.
