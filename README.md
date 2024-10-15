# supaapigo

`supaapigo` is a simple Go package that wraps Supabase's authentication and storage functionalities using the `auth-go` and `storage-go` libraries. It provides an easy-to-use interface to interact with Supabase's authentication and storage APIs.

## Features

- **Authentication**: Integrate Supabase's authentication features into your Go projects.
- **Storage**: Easily interact with Supabase storage for file uploads, downloads, and management.

## Installation

To use `supaapigo` in your Go projects, you can install it with:

```bash
go get github.com/yourusername/supaapigo
```

Make sure to replace `yourusername` with the actual username or organization where your repository is hosted.

## Usage

### Initialize the Usecase

To use the package across your projects, start by initializing the `SupaapiUsecase`. You'll need your Supabase project reference, storage URL, and API key to authenticate.

```go
package main

import (
	"github.com/yourusername/supaapigo"
)

func main() {
	// Initialize the SupaapiUsecase
	projectReference := "your-supabase-project-reference"
	storageURL := "https://your-supabase-storage-url"
	apiKey := "your-supabase-api-key"

	supaapi := supaapigo.NewSupaapiUsecase(projectReference, storageURL, apiKey)

	// Use supaapi for authentication and storage operations
	// Example: Uploading files, managing authentication, etc.
}
```

### Example: Authentication

The `SupaapiUsecase` uses the `auth-go` package to handle authentication-related operations. You can access the `authClient` from the usecase to sign up users, sign in, and manage sessions.

Example for signing up a user:

```go
import "github.com/supabase-community/auth-go"

func SignUpUser(supaapi supaapigo.SupaapiUsecaseInterface, email, password string) error {
    user, err := supaapi.AuthClient.SignUp(auth.SignUpUserInput{
        Email:    email,
        Password: password,
    })
    if err != nil {
        return err
    }
    fmt.Printf("User created: %v\n", user)
    return nil
}
```

### Example: Storage

The `SupaapiUsecase` also integrates the `storage-go` package, allowing you to easily manage files in Supabase Storage.

Example for uploading a file:

```go
import (
    "fmt"
    "os"
)

func UploadFile(supaapi supaapigo.SupaapiUsecaseInterface, bucketName, filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    res, err := supaapi.StorageClient.UploadFile(bucketName, filePath, file)
    if err != nil {
        return err
    }
    fmt.Printf("File uploaded: %v\n", res)
    return nil
}
```

### Interfaces for Flexibility

By using the `SupaapiUsecaseInterface`, you can easily mock this interface and write unit tests for your project without depending directly on Supabase’s services.

## Project Structure

- **authClient**: A reference to the Supabase authentication client, which allows you to perform user-related operations such as sign-ups, logins, etc.
- **storageClient**: A reference to the Supabase storage client, used to upload, download, and manage files in the Supabase storage.

## Contributing

If you find a bug or have a feature request, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.

---

This README provides a clear explanation of the package's purpose, how to set it up, and how to use the key functionalities across multiple projects. Adjust the package URLs based on where your project is hosted.
