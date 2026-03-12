


## Usage

> [login_hook: Execute code on user login, comparable to Oracle's after logon trigger](https://github.com/splendiddata/login_hook)

`login_hook` allows executing custom PL/pgSQL code whenever a user logs into the database.

```sql
CREATE EXTENSION login_hook;
```

### Configuration

Add to `postgresql.conf`:

```ini
session_preload_libraries = 'login_hook'
```

### Creating the Login Function

Define a `login_hook.login()` function that will execute on every login:

```sql
CREATE OR REPLACE FUNCTION login_hook.login() RETURNS VOID LANGUAGE PLPGSQL AS $$
BEGIN
    IF NOT login_hook.is_executing_login_hook() THEN
        RAISE EXCEPTION 'Should only be invoked by login_hook';
    END IF;

    -- Your login logic here:
    RAISE NOTICE 'Hello %', current_user;

EXCEPTION
    WHEN OTHERS THEN
        RAISE LOG 'Error in login_hook.login(): %', SQLERRM;
END
$$;
GRANT EXECUTE ON FUNCTION login_hook.login() TO PUBLIC;
```

The `PUBLIC` grant is required because the function runs for every connecting user.

### Functions

| Function | Returns | Description |
|----------|---------|-------------|
| `login_hook.is_executing_login_hook()` | `boolean` | Returns true only when called during login hook execution |
| `login_hook.get_login_hook_version()` | `text` | Returns compiled version of login_hook |
| `login_hook.login()` | `void` | User-provided function executed at login |

### Important Notes

- The function is NOT invoked for background processes or during recovery mode
- Handle all exceptions within the function -- failures will prevent normal users from logging in
- Superusers get a warning but can still log in when the function fails
- For PostgreSQL 17+, consider using the native login event trigger instead
