## Usage

Sources:

- [Version 0.0.1 SQL implementation](https://github.com/pindlebot/pg_gen_uid/blob/b5352b72551c56e9a83fcd2981c2c0c82168759b/pg_gen_uid--0.0.1.sql)
- [Extension control file](https://github.com/pindlebot/pg_gen_uid/blob/b5352b72551c56e9a83fcd2981c2c0c82168759b/pg_gen_uid.control)

`pg_gen_uid` installs one PL/pgSQL helper, `public.gen_uid(size integer)`, which returns lowercase ASCII letters and decimal digits. It depends on `pgcrypto` for random bytes. The result is a random label, not a UUID, ordered identifier, or guarantee of uniqueness.

### Core Workflow

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_gen_uid;

CREATE TABLE invitation (
  code text PRIMARY KEY DEFAULT gen_uid(20),
  created_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO invitation DEFAULT VALUES
RETURNING code;
```

`gen_uid(20)` returns 20 characters drawn from `abcdefghijklmnopqrstuvwxyz0123456789`. Keep a unique constraint and retry a collision if the identifier must be unique. The extension hard-codes the function in `public`, so account for that schema when setting privileges and `search_path`.

### Randomness and Input Boundaries

Each random byte is reduced modulo 36. Because 256 is not divisible by 36, four alphabet characters occur slightly more often than the others; do not use the output where uniform tokens or quantified cryptographic entropy are required.

Validate that `size` is non-NULL, positive, and reasonably bounded before exposing the helper. The function requests all random bytes at once and loops once per output character, while `gen_random_bytes` itself has a per-call size limit. The reviewed repository has no README, tests, release notes, or declared license, so independently test behavior and deployment policy.
