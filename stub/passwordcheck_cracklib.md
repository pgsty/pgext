## Usage

Sources:

- [3.2.1 README](https://github.com/devrimgunduz/passwordcheck_cracklib/blob/3.2.1/README.md)
- [3.2.1 password hook](https://github.com/devrimgunduz/passwordcheck_cracklib/blob/3.2.1/passwordcheck_cracklib.c)
- [PostgreSQL passwordcheck manual](https://www.postgresql.org/docs/18/passwordcheck.html)

`passwordcheck_cracklib` checks passwords supplied through `CREATE ROLE` and `ALTER ROLE` with CrackLib. It is a server hook library, with no SQL extension objects.

### Enable the Hook

Add it to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = '$libdir/passwordcheck_cracklib'
```

Do not run `CREATE EXTENSION passwordcheck_cracklib`. CrackLib's library and dictionary must be available to the PostgreSQL operating-system account.

### Password Checks

```sql
CREATE ROLE app_user LOGIN PASSWORD 'password123';
```

Weak plaintext passwords cause an error. The hook checks length, the relationship to the username, character composition, and the CrackLib dictionary. A password passing those checks is not a guarantee of resistance to every attack.

### Security Boundary

Dictionary checks require the plaintext password at password-change time. When a client supplies an already hashed password, the module cannot perform full strength checks; its remaining check is whether the password equals the username. Enforce the intended password-change path and protect the connection carrying plaintext passwords.

Existing passwords are not scanned retroactively. The hook also chains to a previously installed password-check hook; review other credential-policy libraries before loading them together.
