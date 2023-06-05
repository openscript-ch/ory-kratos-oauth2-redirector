# Ory Kratos OAuth2 Redirector

Redirects incoming get requests to the OAuth2 provider configured inside Kratos.

## Usage

Use the Docker image pushed to the Github Registry:

```bash
docker pull ghcr.io/openscript-ch/ory-kratos-oauth2-redirector:latest
```

### Environment variables

| Name | Description | Default |
|---|---|---|
| `PORT` | Listning port of this service for HTTP requests | `3000` |
| `LOGIN_ENDPOINT` | Endpoint which initiates the OAuth2 login process | - |
| `REGISTRATION_ENDPOINT` | Endpoint which initiates the OAuth2 registration process | - |

### Query parameters

| Name | Description | Example |
|---|---|---|
| `provider` | OIDC provider | `google` |
| `csrf_token` | CSRF token of the flow | `32kljfa2lc03gov8sal...` |
| `flow` | ID of the flow | `32kljfa2lc03gov8sal...` |
| `traits` | When traits are present, the `REGISTRATION_ENDPOINT` is used, otherwise it's `LOGIN_ENDPOINT`. | `{"lang":"de","tac":"timestamp","dp":"timestamp"}` |
| `upstream_parameters` | https://www.ory.sh/docs/kratos/bring-your-own-ui/custom-ui-advanced-integration#upstream-provider-parameters | `{"prompt":"select_account"}` |

### Example requests

- **Login**: `http://localhost:3000?provider=google&csrf_token=S9qg1Y2EZZ2vghjFAnXtDV2VTosrGFAe/XdOmlyD9MQyB+K/2g9FrtpkMe40/yJ8lOCuFDQHXbVitbR5iIsRtw==&flow=56dfcc4a-95af-41c6-876f-93b9c75f9e29`
- **With traits**: `http://localhost:3000?oauth2?provider=google&csrf_token=92kMiSR3SQ/gkipWWeW6kXMwExTMwCUWIveiNwzD4JUSvyjk1WBzCgZTneaFRBTxXzS/jpQirg+uZnfri31NDw==&flow=497d988d-167a-4588-9821-38f4863addb1&traits=%7B%22lang%22%3A%22en%22%2C%22tac%22%3A%221681216465715%22%2C%22dp%22%3A%221681216465715%22%7D`
- **With traits and upstream_parameters**: `http://localhost:3000?oauth2?provider=google&csrf_token=92kMiSR3SQ/gkipWWeW6kXMwExTMwCUWIveiNwzD4JUSvyjk1WBzCgZTneaFRBTxXzS/jpQirg+uZnfri31NDw==&flow=497d988d-167a-4588-9821-38f4863addb1&traits=%7B%22lang%22%3A%22en%22%2C%22tac%22%3A%221681216465715%22%2C%22dp%22%3A%221681216465715%22%7D&upstream_parameters=%7B%22prompt%22%3A%22select_account%22%7D`

## Development

Run with `LOGIN_ENDPOINT=https://httpbin.org/anything REGISTRATION_ENDPOINT=https://httpbin.org/anything go run ory_kratos_oauth2_redirector.go`

### Upgrade dependencies

1. Update dependencies with `go get -u`
1. Clean up dependency tree with `go mod tidy`
