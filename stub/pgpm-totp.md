## Usage

Sources:

- [packages/totp/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/totp/README.md)
- [packages/totp/sql/pgpm-totp--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/totp/sql/pgpm-totp--0.47.0.sql)
- [packages/totp/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/totp/Makefile)
- [packages/totp/pgpm-totp.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/totp/pgpm-totp.control)

`pgpm-totp` supplies `totp.generate` and `totp.verify`, backed by pgcrypto HMAC and PGPM Base32 helpers. Defaults are a 30-second period, six digits and SHA-1; the SQL also accepts an explicit time, hash, encoding and clock offset.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-totp";
WITH demo AS (SELECT totp.random_base32(20) AS secret)
SELECT totp.verify(secret, totp.generate(secret, 30, 6), 30, 6)
FROM demo;
```

### Operational Boundaries

`totp.random_base32` creates secret text from random bytes, and `totp.url` builds an enrollment URI that contains the secret. Protect both the secret and enrollment URI. Verification checks one selected time step; it does not track replay, enforce rate limits or search an acceptance window. Those policies belong to the application. The URI encoder rejects characters requiring three or more UTF-8 bytes. Dependencies are `pgcrypto`, `plpgsql`, `pgpm-base32`, `pgpm-verify`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
