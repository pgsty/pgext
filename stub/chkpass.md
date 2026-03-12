

## Usage

> [chkpass: auto-encrypted password data type](https://github.com/lacanoid/chkpass)

The `chkpass` extension provides a data type for storing encrypted passwords. Originally bundled with PostgreSQL (removed in PG 11), this is a standalone version for modern PostgreSQL.

```sql
CREATE EXTENSION chkpass;
```

### Data Type

The `chkpass` type automatically encrypts passwords using Unix `crypt()` on input and stores only the encrypted form.

```sql
CREATE TABLE accounts (
    username text PRIMARY KEY,
    password chkpass
);

INSERT INTO accounts VALUES ('admin', 'mysecretpassword');
```

### Operators

The `=` operator checks a plaintext password against the stored encrypted value:

```sql
SELECT * FROM accounts WHERE password = 'mysecretpassword';
-- Returns the matching row if the password is correct
```

### Behavior

- Passwords are automatically encrypted on input -- the plaintext is never stored
- Output displays the encrypted hash, not the original password
- Comparison with `=` encrypts the right-hand operand and compares hashes
- Uses the standard Unix `crypt()` function for encryption
