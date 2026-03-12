


## Usage

> [auth_delay: Pause before reporting authentication failure](https://www.postgresql.org/docs/current/auth-delay.html)

`auth_delay` pauses the server briefly before reporting authentication failures, making brute-force password attacks more difficult.

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'auth_delay'
auth_delay.milliseconds = '500'
```

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `auth_delay.milliseconds` | `0` | Milliseconds to wait before reporting auth failure |

### Notes

- Must be loaded via `shared_preload_libraries`
- Does not prevent denial-of-service attacks (waiting processes still hold connection slots)
- No `CREATE EXTENSION` is required -- this is a shared library module only
