## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/README)
- [Version 0.1 SQL objects](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/email--0.1.sql)
- [Type implementation](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/email.c)
- [Extension control file](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/email.control)
- [Official PGXN distribution](https://pgxn.org/dist/email/)

`email` is a legacy variable-length email-address type with equality, inequality, a hash operator class, and helpers that split the local and domain portions. Its validation rule recognizes only a narrow subset of addresses, and the reviewed C binary I/O path is unsafe. Do not adopt it as a standards-compliant validator or production data type.

### Core Workflow

Use text input only in an isolated compatibility test:

```sql
CREATE EXTENSION email;

CREATE TABLE contacts (
  address email NOT NULL
);

INSERT INTO contacts VALUES ('alice@example.com');

SELECT address, getuser(address), getdomain(address)
FROM contacts
WHERE address = 'alice@example.com'::email;
```

The C helpers are `getuser(email)` and `getdomain(email)`. The SQL script also installs PL/pgSQL variants named `get_user(email)` and `get_domain(email)`. Equality is byte-for-byte and case-sensitive. The default hash operator class is `email_equal_ops`; there is no ordering operator or B-tree operator class.

### Validation and Safety Boundaries

Version 0.1 accepts only lowercase ASCII characters matching the implementation's regular expression, including a final alphabetic suffix of two through four characters. It rejects uppercase input, internationalized addresses, and many valid modern domains, while syntactic acceptance does not prove that a mailbox or domain exists.

The hash support function is a provisional PL/pgSQL calculation marked volatile by the install script. More critically, the C receive function copies into an uninitialized pointer and the send function reads the varlena datum as a single character. Binary protocol input/output and binary copy can therefore corrupt memory or crash the backend. Avoid binary I/O, dumps that depend on the binary representation, and real workloads; migrate data to `text` plus application-appropriate validation.
