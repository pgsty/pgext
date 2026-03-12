


## Usage

> [pg_html5_email_address: HTML5 email address validation for PostgreSQL](https://github.com/bigsmoke/pg_html5_email_address)

Provides email address validation consistent with the HTML5 `<input type="email">` specification.

### Domain Type: `html5_email`

A domain type that enforces HTML5 email validation rules with case-insensitive comparison:

```sql
SELECT 'user@example.com'::html5_email;

-- Case-insensitive equality:
SELECT 'User@Example.com'::html5_email = 'user@example.com'::html5_email;
-- t

-- Invalid emails raise check_violation:
SELECT 'user @example.com'::html5_email;
-- ERROR: check_violation
```

### Function: `html5_email_regexp()`

Returns a regex matching valid HTML5 email addresses:

```sql
-- Check if a string is a valid HTML5 email
SELECT 'user@example.com' ~ html5_email_regexp();
-- t

SELECT 'user @example.com' ~ html5_email_regexp();
-- f
```

With optional capturing groups for local and domain parts:

```sql
SELECT (regexp_matches('user@example.com', html5_email_regexp(true)))[1];
-- 'user'
SELECT (regexp_matches('user@example.com', html5_email_regexp(true)))[2];
-- 'example.com'
```

### Validation Rules

- Spaces are not allowed
- Non-ASCII characters are not allowed (neither in local nor domain part)
- There must be something after the `@`
- Special characters like `!#$%&'*+/=?^_`{|}~-` are allowed in the local part
