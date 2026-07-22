## Usage

Sources:

- [Official README](https://api.pgxn.org/src/otp/otp-1.0.0/README.md)
- [Official extension SQL](https://api.pgxn.org/src/otp/otp-1.0.0/sql/otp.sql)
- [Official extension control file](https://api.pgxn.org/src/otp/otp-1.0.0/otp.control)

`otp` version `1.0.0` implements a small SHA-1 TOTP workflow in untrusted server-side Perl. It can generate provisioning URLs and current-window codes, but its secret generator is not cryptographically secure and the extension should be treated as legacy reference code rather than a production authentication system.

### Core Workflow

Installation creates `plperlu`, so it requires superuser privileges, server PL/PerlU support, and the Perl module `Digest::HMAC_SHA1`:

```sql
CREATE EXTENSION otp;

SELECT provisioning_url(
  'alice@example.com',
  'jbswy3dpehpk3pxp',
  30,
  'Example'
);

SELECT generate_totp('jbswy3dpehpk3pxp', 30, 6);
SELECT verify_totp('jbswy3dpehpk3pxp', 30, '123456');
```

`generate_totp(text, integer, integer)` produces a code for the current time step. `verify_totp(text, integer, text)` checks only the current step and always generates a six-digit candidate; it does not implement a clock-drift window or replay prevention.

### Helpers and Security Caveats

`random_base32(integer)` emits a lowercase Base32-like secret, while `urlencode(text)`, `provisioning_url(text, text, integer, text)`, `pack(text)`, `unpack(text)`, and `perl_hmac(text, text)` support the workflow.

`random_base32(integer)` uses PostgreSQL's non-cryptographic `random()` and must not generate authentication secrets. Create secrets with a cryptographically secure external generator, encrypt them at rest, restrict SQL and log exposure, and implement rate limiting, replay protection, recovery, and clock-skew policy outside this extension. `urlencode(text)` handles only limited UTF-8 sequences, so non-ASCII labels can fail or encode incorrectly.
