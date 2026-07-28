## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/totp/readme.md)
- [Official extension control file (launchql-totp.control)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/totp/launchql-totp.control)
- [Official extension SQL (launchql-totp--0.4.5.sql)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/totp/sql/launchql-totp--0.4.5.sql)

`launchql-totp` — TOTP implementation in pure PostgreSQL plpgsql. This extension provides the HMAC Time-Based One-Time Password Algorithm (TOTP) as specified in RFC 6238 as pure plpgsql functions. Use it when implementing the corresponding security, audit, or access-control workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "launchql-totp";

SELECT totp.generate('mysecret');

-- you can also specify totp_interval, and totp_length
SELECT totp.generate('mysecret', 30, 6);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `totp.base32_to_hex(input text)` is an extension function and returns `text`.
- `totp.generate(secret text, period int DEFAULT 30, digits int DEFAULT 6, time_from timestamptz DEFAULT now(), hash text DEFAULT 'sha1', encoding text DEFAULT 'base32', clock_offset int DEFAULT 0)` is an extension function and returns `text`.
- `totp.generate_secret(hash text DEFAULT 'sha1')` is an extension function and returns `bytea`.
- `totp.hotp(key bytea, c int, digits int DEFAULT 6, hash text DEFAULT 'sha1')` is an extension function and returns `text`.
- `totp.pad_secret(input bytea, len int)` is an extension function and returns `bytea`.
- `totp.random_base32(_length int DEFAULT 20)` is an extension function and returns `text`.
- `totp.url(email text, totp_secret text, totp_interval int, totp_issuer text)` is an extension function and returns `text`.
- `totp.urlencode(in_str text)` is an extension function and returns `text`.
- `totp.verify(secret text, check_totp text, period int DEFAULT 30, digits int DEFAULT 6, time_from timestamptz DEFAULT now(), hash text DEFAULT 'sha1', encoding text DEFAULT 'base32', clock_offset int DEFAULT 0)` is an extension function and returns `boolean`.
- `totp` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.4.5`.
- Install the confirmed extension dependencies first: `pgcrypto`, `plpgsql`, `launchql-base32`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
