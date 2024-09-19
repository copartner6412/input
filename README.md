# common: Overview
common is a Go project consisting of three primary packages: pseudorandom, random, and validate. These packages provide a range of utility functions for generating deterministic pseudo-random values, cryptographically-secure random values, and performing input validation.

__Contributions and suggestions are welcome and appreciated.__

## Packages

### 1. pseudorandom
The pseudorandom package generates deterministic pseudorandom data using seeded randomness. It is useful when reproducibility is required for __testing__ or simulation purposes.

### 2. random
The random package provides cryptographically-secure random data generation, ensuring unpredictable results for secure applications.

### 3. validate
The validate package provides utility functions to validate common input formats such as email addresses, domains, IP addresses, and more.

## Installation
To install the common project and its packages, you can run:

```bash
go get -u github.com/copartner6412/common
```
Then import the specific packages you need:

```go
import (
    "github.com/copartner6412/input/pseudorandom"
    "github.com/copartner6412/input/random"
    "github.com/copartner6412/input/validate"
)
```

## Usage Examples
### Example 1: Generate a Random Password
```go
password, err := random.Password(20, true, true, true, true)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Generated Password: %s\n", password)
```
### Example 2: Validate an Email Address
```go
err := validate.Email("example@domain.com")
if err != nil {
    return fmt.Printf("Invalid email: %v", err)
}
```
### Example 3: Generate a Pseudorandom Domain
```go
r := rand.New(rand.NewSource(42)) // Seeded for reproducibility
domain := pseudorandom.Domain(r)
fmt.Printf("Generated Domain: %s\n", domain)
```
## Contribution
Contributions are welcome! If you’d like to contribute to the project, please follow these steps:

- Fork the repository.
- Create a feature branch (`git checkout -b feature-name`).
- Commit your changes (`git commit -m "Description of feature"`).
- Push to the branch (`git push origin feature-name`).
- Create a pull request, and describe the changes in detail.
License
## License
This project is licensed under the MIT License - see the LICENSE file for details.

