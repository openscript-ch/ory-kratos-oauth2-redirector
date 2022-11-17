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