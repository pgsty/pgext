


## Usage

> [passwordcheck_cracklib: Strengthen PostgreSQL user password checks with cracklib](https://github.com/devrimgunduz/passwordcheck_cracklib)

`passwordcheck_cracklib` is like the regular PostgreSQL `passwordcheck` module, except it is built with cracklib for more strict password checks. It checks users' passwords whenever they are set with `CREATE ROLE` or `ALTER ROLE`. If a password is considered too weak, it will be rejected and the command will terminate with an error.

### Configuration

Add the library to `shared_preload_libraries` in `postgresql.conf`:

```ini
shared_preload_libraries = '$libdir/passwordcheck_cracklib'
```

Restart PostgreSQL to activate.

### How It Works

Once loaded, any `CREATE ROLE` or `ALTER ROLE` command that sets a password will have the password checked against cracklib's dictionary. Weak or easily guessable passwords will be rejected automatically.

```sql
-- This will be rejected if the password is too weak
CREATE ROLE myuser WITH LOGIN PASSWORD 'password123';
-- ERROR: password is easily cracked

-- A strong password will be accepted
CREATE ROLE myuser WITH LOGIN PASSWORD 'X9#kLm$vQ2!pR7';
```

No `CREATE EXTENSION` is required -- this is a shared library module only.
