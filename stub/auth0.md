## Usage

Sources:

- [Official README](https://github.com/brunoenten/pg_auth0/blob/c6831e8eb821384fd5ef9950e31bfe35450913e7/README.md)
- [Version 0.1 SQL implementation](https://github.com/brunoenten/pg_auth0/blob/c6831e8eb821384fd5ef9950e31bfe35450913e7/auth0--0.1.sql)
- [Extension control file](https://github.com/brunoenten/pg_auth0/blob/c6831e8eb821384fd5ef9950e31bfe35450913e7/auth0.control)

`auth0` is a SQL client for selected Auth0 Authentication and Management API operations. It sends HTTPS requests from PostgreSQL through the required `http` extension, so a query can create, read, or update Auth0 users and trigger a password-change message.

### Core Workflow

Create the dependency and extension, then store the tenant settings for the current database role. The SQL implementation uses `set_config`; the differently named setters shown in the README are not part of version `0.1`.

```sql
CREATE EXTENSION http;
CREATE EXTENSION auth0;

SELECT auth0.set_config('domain', 'example.auth0.com');
SELECT auth0.set_config('client_id', 'client-id');
SELECT auth0.set_config('client_secret', 'client-secret');
SELECT auth0.set_config('connection', 'Username-Password-Authentication');

SELECT auth0.get_user_by_email('alice@example.com', 'user_id,email,name');
```

`get_api_token()` performs the client-credentials exchange when the cached token is empty or expired. `get_user()`, `get_user_by_email()`, `create_user()`, and `update_user()` call the Management API; `change_password_prompt()` calls the database-connections endpoint.

### Operational Boundaries

Configuration and cached access tokens are written as per-role PostgreSQL settings by `ALTER ROLE`, not kept in an external secret store. Restrict access to the role and system catalogs, and plan explicit secret rotation. The functions perform synchronous outbound network calls inside a database backend; API latency, rate limits, TLS/DNS failures, and Auth0 permissions therefore affect SQL execution.

The implementation constructs some request paths from caller text and does not expose a general-purpose Auth0 API. Grant execution only to trusted roles, validate identifiers and JSON payloads in the application, and test HTTP status/error behavior before relying on it in a transaction.
