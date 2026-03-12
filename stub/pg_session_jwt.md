


## Usage

> [pg_session_jwt: JWT session management for PostgreSQL](https://github.com/neondatabase/pg_session_jwt)

`pg_session_jwt` handles authenticated sessions through JWTs. When configured with a JWK, it verifies JWT authenticity. Without a JWK, it falls back to PostgREST-compatible `request.jwt.claims`.

```sql
CREATE EXTENSION pg_session_jwt;
```

### Mode 1: JWK Validation

Set the JWK at connection time via libpq options:

```bash
export PGOPTIONS="-c pg_session_jwt.jwk=$MY_JWK"
```

Then within the session:

```sql
SELECT auth.init();                        -- Initialize with JWK
SELECT auth.jwt_session_init('eyJ...');    -- Set and validate the JWT
SELECT auth.user_id();                     -- Get the 'sub' claim
SELECT auth.session();                     -- Get full JWT payload as JSONB
```

### Mode 2: PostgREST-Compatible (No JWK)

Works out of the box with PostgREST. No initialization needed:

```sql
SELECT auth.user_id();   -- Returns 'sub' from request.jwt.claims
SELECT auth.session();   -- Returns full claims as JSONB
```

### Functions

| Function | Returns | Description |
|----------|---------|-------------|
| `auth.init()` | `void` | Initialize session using JWK |
| `auth.jwt_session_init(jwt text)` | `void` | Set and validate a JWT |
| `auth.session()` | `jsonb` | Get JWT payload or fallback claims |
| `auth.jwt()` | `jsonb` | Alias for `auth.session()` |
| `auth.user_id()` | `text` | Get the `sub` claim |
| `auth.uid()` | `uuid` | Get `sub` as UUID (or NULL) |

### Configuration

| Parameter | Description |
|-----------|-------------|
| `pg_session_jwt.jwk` | JWK for JWT validation (set at startup or connection) |
| `pg_session_jwt.audit_log` | Enable audit logging (`on`/`off`) |

### RLS Example

```sql
CREATE POLICY user_isolation ON my_table
    USING (user_id = auth.user_id());
```
