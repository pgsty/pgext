## Usage

Sources:

- [Version 0.0.1 SQL implementation](https://github.com/pindlebot/pg_gen_uid/blob/b5352b72551c56e9a83fcd2981c2c0c82168759b/pg_gen_uid--0.0.1.sql)
- [Extension control file](https://github.com/pindlebot/pg_gen_uid/blob/b5352b72551c56e9a83fcd2981c2c0c82168759b/pg_gen_uid.control)

`pg_gen_uid` installs `public.gen_uid(size)`, a PL/pgSQL helper that returns a requested number of lowercase ASCII letters and digits. It uses `pgcrypto` to obtain random bytes.

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_gen_uid;
SELECT gen_uid(16);
```

The extension has no README, tests, release notes, or declared license in the reviewed repository. Its function is fixed in `public`, and each byte is reduced modulo the 36-character alphabet, which introduces a small distribution bias. Generated values are not a substitute for a database uniqueness constraint; use a unique index and retry collisions when identifiers must be unique.

Validate the requested length before exposing the function to untrusted callers, because it allocates random bytes and loops once per output character.
