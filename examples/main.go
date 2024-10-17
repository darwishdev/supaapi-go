package main

import ()

//	http://127.0.0.1:54321
//	    GraphQL URL: http://127.0.0.1:54321/graphql/v1
//	 S3 Storage URL: http://127.0.0.1:54321/storage/v1/s3
//	         DB URL: postgresql://postgres:postgres@127.0.0.1:54322/postgres
//	     Studio URL: http://127.0.0.1:54323
//	   Inbucket URL: http://127.0.0.1:54324
//	     JWT secret: super-secret-jwt-token-with-at-least-32-characters-long
//	       anon key: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZS1kZW1vIiwicm9sZSI6ImFub24iLCJleHAiOjE5ODM4MTI5OTZ9.CRXP1A7WOeoJeXxjNni43kdQwgnWNReilDMblYTn_I0
//
// service_role key: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZS1kZW1vIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImV4cCI6MTk4MzgxMjk5Nn0.EGIM96RAZx35lJzdJsyH-qQwv8Hdp7fsn3W0YpN81IU
//
//	S3 Access Key: 625729a08b95bf1b7ff351a663f3a23c
//	S3 Secret Key: 850181e4652dd023b7a98c58ae0d2d34bd487ee0cc3254aed6eda37307425907
//	    S3 Region: local
//
// http://localhost:5173/provider-login?state=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkwOTU1NTQsInNpdGVfdXJsIjoiaHR0cDovLzEyNy4wLjAuMTozMDAwIiwiaWQiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJmdW5jdGlvbl9ob29rcyI6bnVsbCwicHJvdmlkZXIiOiJnb29nbGUiLCJyZWZlcnJlciI6Imh0dHA6Ly8xMjcuMC4wLjE6MzAwMCIsImZsb3dfc3RhdGVfaWQiOiIifQ.FJ-Ce4kOL37eWvzzMWIdEtn47kM0fOwWglnybvKKh8Q&code=4%2F0AVG7fiR2TacnsnJWWbWkUEAUIp-v0YdQ0zfmvf5G3_bpvK3FQTRTjHmrTwt8kVnUSMcurQ&scope=email+profile+openid+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email&authuser=0&prompt=consent
func main() {
	// supapi := supaapigo.NewSupaapi(supaapigo.SupaapiConfig{
	// 	ProjectRef:            "ewooobaokfikpwibdfdh",
	// 	Port:                  54321,
	// 	OAuthRegisterCallback: "http://localhost:5173/provider-register",
	// 	OAuthLoginCallback:    "http://localhost:5173/provider-login",
	// 	Env:                   "PROD",
	// 	ServiceRoleKey:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImV3b29vYmFva2Zpa3B3aWJkZmRoIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTcxMTU0OTU1NCwiZXhwIjoyMDI3MTI1NTU0fQ.QMOznbjiS0hh4ZWMbLimKV0sTvW7zgxW4Eec4IfLfm4",
	// 	ApiKey:                "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImV3b29vYmFva2Zpa3B3aWJkZmRoIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTE1NDk1NTQsImV4cCI6MjAyNzEyNTU1NH0.4VM8tl_hLqFAyvWW4k1PHhHZ0q-jL18gNupgK1T6Y-0",
	// })

	// supapi := supaapigo.NewSupaapi(supaapigo.SupaapiConfig{
	// 	ProjectRef:            "ewooobaokfikpwibdfdh",
	// 	Port:                  54321,
	// 	GoogleClientID:        "1004945141938-7i4e5mokgordq00a6kc4v68vo2bitd0b.apps.googleusercontent.com",
	// 	GoogleClientSecret:    "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",
	// 	OAuthRegisterCallback: "http://localhost:5173/provider-register",
	// 	OAuthLoginCallback:    "http://localhost:5173/provider-login",
	// 	Env:                   "DEV",
	// 	ServiceRoleKey:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZS1kZW1vIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImV4cCI6MTk4MzgxMjk5Nn0.EGIM96RAZx35lJzdJsyH-qQwv8Hdp7fsn3W0YpN81IU",
	// 	ApiKey:                "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZS1kZW1vIiwicm9sZSI6ImFub24iLCJleHAiOjE5ODM4MTI5OTZ9.CRXP1A7WOeoJeXxjNni43kdQwgnWNReilDMblYTn_I0",
	// })
	// resp2, err2 := supapi.GoogleLogin()
	// resp, err := supapi.GoogleExchangeCode("4%2F0AVG7fiRbEGZDzAeqZTfDdhE7YhCjHWyVVT-YKKLJQ3Yf9VnIMu2DFDZOeHcPkoOwFU7Oxg")
	// fmt.Println(err)
	// fmt.Println(resp)
	// fmt.Println(err2)
	// fmt.Println(resp2)
	//
	// resetLink, err := supapi.AuthClient.Invite(types.InviteRequest{Email: "darwishdev.com@gmail.com"})
	// resp, err := supapi.AuthClient.WithToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwOi8vMTI3LjAuMC4xOjU0MzIxL2F1dGgvdjEiLCJzdWIiOiIwMzE5ZjFhZi05ODgyLTQzYjMtYjhmYS01YzRiNTlkMzdmNGEiLCJhdWQiOiJhdXRoZW50aWNhdGVkIiwiZXhwIjoxNzI5MTA2NzY4LCJpYXQiOjE3MjkxMDMxNjgsImVtYWlsIjoiYS5kYXJ3aXNoLmRldkBnbWFpbC5jb20iLCJwaG9uZSI6IiIsImFwcF9tZXRhZGF0YSI6eyJwcm92aWRlciI6ImVtYWlsIiwicHJvdmlkZXJzIjpbImVtYWlsIl19LCJ1c2VyX21ldGFkYXRhIjp7fSwicm9sZSI6ImF1dGhlbnRpY2F0ZWQiLCJhYWwiOiJhYWwxIiwiYW1yIjpbeyJtZXRob2QiOiJvdHAiLCJ0aW1lc3RhbXAiOjE3MjkxMDMxNjh9XSwic2Vzc2lvbl9pZCI6ImVmNjgwMTZlLTQ5N2MtNGQxMy1iYTFlLTY1MzNiNDJjNmI1MiIsImlzX2Fub255bW91cyI6ZmFsc2V9.7nWc2tIXRYcw1CpQt_InMDC8hUXa_4dqvx9-5BWYAEw").GetUser()
	//
	// err := supapi.AuthClient.Recover(types.RecoverRequest{Email: "a.darwish.dev@gmail.com"})
	// resp , err := supapi.AuthClient.
	// fmt.Println(resp)
	// fmt.Println(err)
}
