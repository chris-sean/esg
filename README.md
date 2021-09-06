# esg
Error Struct Generator

### Usage
`esg output_dir pkg_name error_code formatted_message [name_of_arguments..]`

name_of_argument must not be ErrorCode.
### Example
`esg /src/myproj errors InvalidPhone "%s is not valid phone number." Phone`

The generated file is at `/src/myproj/errors/InvalidPhone.go`

The source code is like below.
```go
package errors

import "fmt"

type InvalidPhone struct {
        ErrorCode string
        Phone interface{}
}

func (e InvalidPhone)Error() string {
        return fmt.Sprintf("%s is not valid phone number.", e.Phone)
}
```