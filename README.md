# Ory Kratos OAuth2 Redirector

Redirects incoming get requests to the OAuth2 provider configured inside Kratos.

## Usage

Use the Docker image pushed to the Github Registry:

```
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
| `csrf_token` | CSRF token of the flow | - |
| `flow` | ID of the flow | - |
| `traits` | When traits are present, the `REGISTRATION_ENDPOINT` is used `LOGIN_ENDPOINT` otherwise. | `{"lang":"de","tac":"timestamp","dp":"timestamp"}` |
