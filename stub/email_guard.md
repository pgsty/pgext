## Usage

Sources:

- [Official version 0.35.0 README](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/email_guard/README.md)
- [Official version 0.35.0 control file](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/email_guard/email_guard.control)
- [Official version 0.35.0 SQL](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/email_guard/email_guard--0.35.0.sql)
- [Official package page](https://database.dev/mansueli/email_guard)

`email_guard` provides email normalization, disposable-domain checks, and a Supabase Auth before-user-created hook. Version `0.35.0` blocks disposable addresses and prevents duplicate Gmail identities that differ only by dots, plus-tags, or the `googlemail.com` alias.

### Core Workflow

Install the extension in a stable schema and call its pure helper functions directly:

```sql
CREATE EXTENSION email_guard SCHEMA extensions;

SELECT extensions.normalize_email('First.Last+news@gmail.com');
SELECT extensions.is_disposable_email('person@example.invalid');
```

For Supabase Auth, configure the project's before-user-created Postgres Function hook to call:

```text
extensions.hook_prevent_disposable_and_enforce_gmail_uniqueness
```

The hook accepts the Auth event as `jsonb`. It returns an empty JSON object for an allowed signup, an HTTP-style `403` error for a disposable domain, or `409` when the normalized Gmail identity already exists in `auth.users`. Events without an email, such as phone signup, pass through.

### API and Data

- `normalize_email(text)`: lowercases/trims email input and canonicalizes Gmail/Googlemail local parts.
- `is_disposable_email_domain(text)`: checks the seeded blocklist, including parent-domain matches.
- `is_disposable_email(text)`: extracts the domain and applies the blocklist check.
- `hook_prevent_disposable_and_enforce_gmail_uniqueness(jsonb)`: Supabase Auth hook.
- `disposable_email_domains`: the versioned domain blocklist table.

### Caveats

The hook queries `auth.users`, so it is specific to a Supabase Auth deployment and needs suitable privileges. Its Gmail uniqueness check covers users already present at hook execution time; validate concurrency and error handling in the real signup path. Keep the seeded domain data current by upgrading the extension rather than editing undocumented internals.

The catalog target documented here is `0.35.0`. Upstream `main` has since advanced to `0.36.0`, so pin the reviewed version when reproducing these exact functions and blocklist contents.
