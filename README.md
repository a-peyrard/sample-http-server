# Sample HTTP server

This is a sample HTTP server written in Golang. With single endpoint `/health` which returns a dummy JSON response.

## Run

Launch the http server with the following command:
```bash
just server
```
No configuration is provided, this is a 2 minutes project, port is hardcoded to 8333.

Then you can test the endpoint with:
```bash
curl localhost:8333/health
```

Finally one can do a bit of load injection testing:
```bash
just inject
```