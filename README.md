# msgraph4go

msgraph4go provides a Go interface for the Microsoft Graph API.

**This is still a work in progress, but does have some working examples for OneDrive and OneNote**

In order to use this package, you must
[Register an application with the Microsoft identity platform](https://docs.microsoft.com/en-us/graph/auth-register-app-v2)

1. Sign in to the [Azure portal](https://portal.azure.com/) using either a work or school account
or a personal Microsoft account.

2. If your account gives you access to more than one tenant,
select your account in the top right corner,
and set your portal session to the Azure AD tenant that you want.

3. Select the **Azure Active Directory** service,
and then select **App registrations > New registration**.

4. When the **Register an application page** appears, enter your application's registration information:

    - **Name** - enter a meaningful name
    - **Supported account types** - select one of the options based on your planned usage
    - **Redirect URI (optional)** - select **Public client (mobile & desktop)**, and then enter `https://login.microsoftonline.com/common/oauth2/nativeclient`

In order to run the examples, you need to set the MSCLIENTID environmental variable to the **Application (client) ID** provided.


The current approach assumes the client runs on a host without a browser. The user is instructed to vist a URL to login and authorize the client. Once the login is successful, the user must copy the response URL and provide to the client program. For example, on the first run without a token file:
```
~/go/src/github.com/bnixon67/msgraph4go/examples$ go run GetMyProfile.go 
Vist the following URL in a browser to authenticate this application
After authentication, copy the response URL from the browser
https://login.microsoftonline.com/common/oauth2/v2.0/authorize?access_type=offline&client_id={client_id}&redirect_uri=https%3A%2F%2Flogin.microsoftonline.com%2Fcommon%2Foauth2%2Fnativeclient&response_type=code&scope=User.Read&state={state}
```
Copy and paste the URL into a browser with javascript enabled to login to your Microsoft account. Once logged in, then copy the URL from the browser into the program:
```
Enter the response URL:
https://login.microsoftonline.com/common/oauth2/nativeclient?code={code}&state={state}
{
  "displayName": "Bill Nixon",
  "givenName": "Bill",
  "id": "16be860d241e39e5",
  "surname": "Nixon",
  "userPrincipalName": "bnixon67@gmail.com"
}
```

The token is requested for offline access, which should include a refresh token to allow access for a long period of time.  The token is saved in the file provided to ```msgraph4go.New(...)```.

A simple example, which returns a JSON result:
```go
// Get Microsoft Application (client) ID
// The ID is not in the source code to avoid someone reusing the ID
clientID, present := os.LookupEnv("MSCLIENTID")
if !present {
	log.Fatal("Must set MSCLIENTID")
}

msGraphClient := msgraph4go.New(".token.json", clientID, []string{"User.Read"})

resp, err := msGraphClient.Get("/me/drive/root", nil)
if err != nil {
	log.Fatal(err)
}
```

Another example that returns a custom User type:
```go
// Get Microsoft Application (client) ID
// The ID is not in the source code to avoid someone reusing the ID
clientID, present := os.LookupEnv("MSCLIENTID")
if !present {
	log.Fatal("Must set MSCLIENTID")
}

msGraphClient := msgraph4go.New(".token.json", clientID, []string{"User.Read"})

resp, err := msGraphClient.GetMyProfile(nil)
if err != nil {
	log.Fatal(err)
}
```
