# Golang Pipedrive client

[![Build Status](https://travis-ci.org/Genert/go-pipedrive.svg?branch=master)](https://travis-ci.org/Genert/go-pipedrive)
[![Coverage Status](https://coveralls.io/repos/github/Genert/go-pipedrive/badge.svg?branch=master)](https://coveralls.io/github/Genert/go-pipedrive?branch=master)

> Work in progress. Do not use in production.

# Supported resources

- [ ] Activities
- [ ] ActivityTypes
- [ ] Authorizations
- [x] Currencies
- [ ] Deals
- [ ] DealFields
- [ ] Files
- [ ] Filters
- [ ] Notes
- [x] NoteFields
- [ ] Organizations
- [ ] OrganizationFields
- [ ] Persons
- [ ] PersonFields
- [ ] Pipelines
- [ ] Products
- [ ] ProductFields
- [x] Recents
- [x] SearchResults
- [ ] Stages
- [ ] Users

## Usage

```go
import "github.com/genert/api-client/pipedrive"
```

Construct a new Pipedrive client, then use the various services on the client to
access different parts of the API. For example:

```go
    const apiKey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

    client := pipedrive.New(&pipedrive.Config{
        ApiKey: apiKey,
    })

    // Return list of all fields for note
    noteFields, _, _ := client.NoteFields.List()

    // You can then access data like this:
    fmt.Println("Success = ", noteFields.Success)
    fmt.Println("First note field: ", noteFields.Data[0].Name)
```

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.