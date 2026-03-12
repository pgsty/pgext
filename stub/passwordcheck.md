


## Usage

> [passwordcheck: Check password strength on CREATE/ALTER ROLE](https://www.postgresql.org/docs/current/passwordcheck.html)

`passwordcheck` validates password strength whenever passwords are set using `CREATE ROLE` or `ALTER ROLE`. Weak passwords are rejected with an error.

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'passwordcheck'
```

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `passwordcheck.min_password_length` | `8` | Minimum password length in bytes (superuser only) |

### How It Works

The module checks passwords set via `CREATE ROLE` or `ALTER ROLE`:

```sql
-- Rejected if password is too short or too weak
CREATE ROLE myuser WITH LOGIN PASSWORD 'abc';
-- ERROR: password is too short

-- Accepted with a strong enough password
CREATE ROLE myuser WITH LOGIN PASSWORD 'Str0ng_P@ssword!';
```

### Default Checks

Without CrackLib, the module enforces:
- Minimum password length (configurable via `passwordcheck.min_password_length`)
- Password must not be the username
- Basic complexity requirements

### Limitations

- Pre-encrypted passwords sent by client programs cannot be fully validated
- The module can only guess the actual password from encrypted submissions
- For stronger security, consider external authentication methods (e.g., GSSAPI)
- No `CREATE EXTENSION` is required -- this is a shared library module only
