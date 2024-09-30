# auth-service
The service is responsible for generation od JWT token.

This project implements [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) in order to assure both high elasticity and reusability of the code.
#### Layers

- `internal/api` - can import any other layers.
- `internal/infrastructure` - can import everything except `api`.
- `internal/use_cases` - should not import anything else than `internal/app` or standard library (or `utils`)

#### Middleware

In order to reuse common code such as authorization this project utilizes also middleware
pattern. Handlers are wrapped in pipelines where each step wraps and calls the next one.

## API usage:

**Create Token:**
```text
POST /token
    {
        "username": "test_user@email.com",
        "password": "xxxxx"
    }
```

**Validate token:**
```text
PATCH /token/validate
    {
        "token": "example_token_string"
    }
```