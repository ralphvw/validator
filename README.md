# Validator Package

The `validator` package offers a utility function to validate required fields within a given struct. This function checks for the existence and emptiness of specified fields.

## Installation

This package is written in Go and can be installed using `go get`:

```bash
go get github.com/ralphvw/validator
```

## Usage

### Function Signature

```go
func Validate(obj interface{}, fields []string) (bool, []string)
```

- `obj`: An instance of a struct or any type with fields to be validated.
- `fields`: A list of field names to be checked within the provided `obj`.

### Example Usage

```go
import (
	"github.com/your-username/validator"
	"encoding/json"
	"net/http"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ProjectId int `json:"projectId"`
		UserId    int `json:"userId"`
	}

	// Decode the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Define the required fields for validation
	requiredFields := []string{"ProjectId", "UserId"}

	// Validate the request payload
	fieldsExist, missingFields := validator.Validate(requestPayload, requiredFields)

	if !fieldsExist {
		http.Error(w, fmt.Sprintf("Missing fields: %v\n", missingFields), http.StatusBadRequest)
		return
	}
	// Proceed with your logic as required
}
```

### Behavior

- The `Validate` function checks if the provided fields exist and are non-zero values within the given object or struct.
- It returns a boolean indicating whether all fields exist and are non-zero, along with a list of missing fields.

## Contribution

Feel free to contribute by forking this repository, making changes, and creating a pull request. Bug reports, suggestions, and enhancements are always welcome!
