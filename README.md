
### 1. **Overview**
The `supaapigo` package provides a convenient API wrapper for interacting with Supabase authentication and storage services. It simplifies user management (creating and updating users), token authorization, and integration with third-party OAuth providers such as Google.

### 2. **Installation**
To install the `supaapigo` package, run the following command:
```bash
go get github.com/yourusername/supaapigo
```

### 3. **Configuration**
Before using the package, you need to create a `SupaapiConfig` structure to hold your Supabase project configuration:
```go
type SupaapiConfig struct {
    ProjectRef        string // Your Supabase project reference
    ApiKey            string // API key for the Supabase project
    ServiceRoleKey    string // Service role key for authentication
    OAuthLoginCallback string // OAuth callback URL for login
    OAuthRegisterCallback string // OAuth callback URL for registration
    Env               Env    // Environment (DEV or PROD)
    Port              int    // Port number for development
}
```
### 4. **Usage**
#### 4.1 **Creating a New Supaapi Instance**
To create a new instance of `Supaapi`, use the `NewSupaapi` function:
```go
package main

import "github.com/yourusername/supaapigo"

func main() {
    config := supaapigo.SupaapiConfig{
        ProjectRef:         "your_project_ref",
        ApiKey:             "your_api_key",
        ServiceRoleKey:     "your_service_role_key",
        OAuthLoginCallback: "https://yourdomain.com/auth/callback",
        OAuthRegisterCallback: "https://yourdomain.com/auth/register/callback",
        Env:                supaapigo.DEV, // or supaapigo.PROD
        Port:               3000, // Development port
    }
    
    supaAPI := supaapigo.NewSupaapi(config)
}
```

#### 4.2 **Creating or Updating a User**
To create a new user or update an existing one, use the `UserCreateUpdate` method:
```go
createUpdateReq := types.AdminUpdateUserRequest{
    UserID:       uuid.New(), // or existing user ID
    Aud:          "authenticated",
    Role:         "user",
    Email:        "user@example.com",
    Phone:        "1234567890",
    Password:     "password123",
    EmailConfirm: true,
    PhoneConfirm: true,
    UserMetadata: map[string]interface{}{"key": "value"},
    AppMetadata:  map[string]interface{}{"key": "value"},
}

user, err := supaAPI.UserCreateUpdate(createUpdateReq)
if err != nil {
    // Handle error
}
```

#### 4.3 **Authorizing a Token**
To authorize a token and retrieve user details:
```go
token := "your_jwt_token"
user, err := supaAPI.AuthorizeToken(token)
if err != nil {
    // Handle error
}
```

#### 4.4 **Google Login**
To initiate Google login:
```go
authorizeResp, err := supaAPI.GoogleLogin()
if err != nil {
    // Handle error
}
```

### 5. **Error Handling**
The package returns errors for various operations. Common errors include:
- Invalid API key
- User already exists
- Token expiration

Check the error messages for details.

### 6. **Advanced Features**
The package allows integration with OAuth providers. You can extend the functionality by adding other providers as needed.

### 7. **Examples**
Include specific examples for user creation, token authorization, and OAuth login as shown in the Usage section.

### 8. **Testing**
For testing, consider using mock services to simulate Supabase responses. Ensure that all public methods have corresponding unit tests.

### 9. **Contributing**
To contribute to this package, please follow these guidelines:
- Fork the repository
- Create a new branch for your feature or bug fix
- Submit a pull request with a description of your changes

### 10. **License**
Specify the license under which this package is distributed (e.g., MIT License).

---

### Conclusion
The `supaapigo` package streamlines interactions with Supabase's authentication and storage services. It simplifies user management and OAuth integrations, making it easier to build applications that leverage these powerful features.

---

Feel free to modify the above documentation to suit your style or add any additional details you think are necessary. If you need further assistance with specific sections or examples, let me know!
