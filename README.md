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

A simple example, which returns a JSON result:
```
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
```
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
