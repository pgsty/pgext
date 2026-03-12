


## Usage

> [pgjwt: JSON Web Tokens for PostgreSQL](https://github.com/michelp/pgjwt)

Requires the `pgcrypto` extension.

### Sign a Token

```sql
SELECT sign(
    '{"sub":"1234567890","name":"John Doe","admin":true}',
    'secret'
);
```

Returns a JWT string like: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOi...`

### Verify a Token

```sql
SELECT * FROM verify(
    'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ',
    'secret'
);
```

Returns:

| header | payload | valid |
|--------|---------|-------|
| `{"alg":"HS256","typ":"JWT"}` | `{"sub":"1234567890","name":"John Doe","admin":true}` | `t` |

### Algorithm Selection

`sign()` and `verify()` accept an optional third argument for the algorithm: `'HS256'` (default), `'HS384'`, or `'HS512'`.

```sql
SELECT sign('{"data":"value"}', 'secret', 'HS384');
```
