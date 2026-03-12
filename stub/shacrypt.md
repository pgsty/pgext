


## Usage

> [shacrypt: SHA-crypt password hashing for PostgreSQL](https://github.com/dverite/postgres-shacrypt)

Generate SHA256-CRYPT and SHA512-CRYPT password hashes per the [SHA-crypt specification](https://www.akkadia.org/drepper/SHA-crypt.txt).

### Functions

#### `sha256_crypt(key text, salt text) RETURNS text`

```sql
SELECT sha256_crypt('clearpassword', 'somesalt');
-- $5$somesalt$l3SlbI688JBlRS9RWFC1EwZLNJqfQKcrF3yhcbc7ffA
```

With custom rounds:

```sql
SELECT sha256_crypt('clearpassword', '$5$rounds=10000$somesalt');
-- $5$rounds=10000$somesalt$OekH6Tu7EOJIAvxKJ4Ko4bG0DxgO83gZODJLTTjXJi5
```

#### `sha512_crypt(key text, salt text) RETURNS text`

```sql
SELECT sha512_crypt('clearpassword', 'somesalt');
-- $6$somesalt$dDcgWMHOtvHI6qT/Khi3uaaxXN6v4N9bnOeWFl/Y6K3pzxi/...
```

### Salt Format

- Simple salt: `'somesalt'`
- With algorithm prefix: `'$5$somesalt'` (SHA-256) or `'$6$somesalt'` (SHA-512)
- With custom rounds: `'$5$rounds=10000$somesalt'`
