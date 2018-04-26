# poc-go-filehippo

PoC reference for interacting with the [FileHippo App Manager](https://filehippo.com/download_app_manager/) API.

You probably shouldn't use this, it may violate the terms of service.

## Usage

```
go run *.go
```

## Known Issues

* The [official client](https://filehippo.com/download_app_manager/) appears to change it's `accessToken` with every request
  * Without understanding how that is generated in the client (.NET), this code is probably useless.. 
