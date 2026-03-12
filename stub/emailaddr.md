

## Usage

> [emailaddr: email address data type for PostgreSQL](https://github.com/petere/pgemailaddr)

The `emailaddr` extension provides a data type for storing and validating email addresses conforming to the `addr-spec` format defined in RFC 5322.

```sql
CREATE EXTENSION emailaddr;

CREATE TABLE accounts (
    id    int PRIMARY KEY,
    name  text,
    email emailaddr
);

INSERT INTO accounts VALUES (1, 'Peter Eisentraut', 'peter@eisentraut.org');
```

### Data Type

The `emailaddr` type validates email addresses on input according to RFC 5322 `addr-spec` rules. Simple formats like `user@domain.com` are accepted. Display name syntax such as `"User Name" <user@domain.com>` is not supported.

### Operators

Standard comparison operators are supported: `=`, `<>`, `<`, `>`, `<=`, `>=`.

### Index Support

Btree indexes are available for efficient lookups and sorting on `emailaddr` columns.
