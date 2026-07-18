## Usage

Sources:

- [Upstream README](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/README)
- [Extension SQL](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/email--0.1.sql)
- [C implementation](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/email.c)
- [PGXN distribution](https://pgxn.org/dist/email/)

`email` adds a variable-length email-address type plus `getuser(email)` and `getdomain(email)` helpers. Values are compared byte-for-byte by the extension's equality operators:

```sql
CREATE EXTENSION email;

CREATE TABLE contacts (address email NOT NULL);
INSERT INTO contacts VALUES ('alice@example.com');

SELECT address, getuser(address), getdomain(address)
FROM contacts
WHERE address = 'alice@example.com'::email;
```

### Validation and indexing limits

Extension version `0.1` accepts only the implementation's narrow lowercase-ASCII pattern, including a two-to-four-letter final domain suffix; uppercase addresses and many modern valid address forms are rejected. The PGXN package is labeled `0.1.0`, while the control and install SQL expose `0.1`. Upstream warns that the hash support is provisional: `email_hash` is implemented in PL/pgSQL, and no ordering operators or B-tree operator class are defined. Test compatibility and validation rules before adopting this legacy type.
