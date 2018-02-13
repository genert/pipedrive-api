# pipedrive-api tests

Install required dependencies first:

```shell
go get -t ../...
```

integration
-----------

This will exercise the library against the live Pipedrive API and verify actual behaviour of the API.

Run tests using:

    PIPEDRIVE_API_TOKEN=XXXXXX go test -v -tags=integration ./integration
