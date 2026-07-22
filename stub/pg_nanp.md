## Usage

Sources:

- [Official extension control file](https://github.com/rlichtenwalter/pg_nanp/blob/8cfb5213faaeef92a19997978f0b6801f0f61aa0/pg_nanp.control)
- [Official README](https://github.com/rlichtenwalter/pg_nanp/blob/8cfb5213faaeef92a19997978f0b6801f0f61aa0/README.md)
- [Official extension SQL](https://github.com/rlichtenwalter/pg_nanp/blob/8cfb5213faaeef92a19997978f0b6801f0f61aa0/pg_nanp--1.0.sql)

`pg_nanp` 1.0 adds the `nanp` base type for North American Numbering Plan telephone numbers. Its input function accepts common separators and an optional leading country code, validates a ten-digit NANP shape, and always displays the stored value as `AAA-EEE-NNNN`.

### Core Workflow

```sql
CREATE EXTENSION pg_nanp;

SELECT '2345678910'::nanp;
SELECT '1 (234) 567-8910'::nanp;

CREATE TABLE contacts (
    contact_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    phone nanp NOT NULL
);

INSERT INTO contacts (phone) VALUES ('234.567.8910');
```

Accepted punctuation is limited to optional spaces, hyphens, dots, and parentheses in the positions implemented by the input regular expression. The optional country prefix must be `1`. Both the area code and central-office code must begin with 2–9, and N11 central-office codes such as 211 or 911 are rejected.

### Type Behavior

`nanp` is stored as an eight-byte pass-by-value integer and has only text input and output functions. The extension does not define comparison or equality operators, casts to or from `text` or integer types, a B-tree operator class, formatting variants, extensions, or region/assignment lookups.

Because no comparison operators are provided, ordinary `PRIMARY KEY`, `UNIQUE`, `ORDER BY`, equality joins, and B-tree indexes on `nanp` are not available without additional operator definitions. Use the type primarily for input validation and display, or store a canonical text or numeric companion column when lookup and indexing are required.

### Validation Boundary

The regular expression validates only structural NANP rules encoded in this 1.0 release. It does not determine whether an area code or exchange is assigned, whether a number is reachable, whether it is mobile, or whether ownership has changed. The upstream implementation is a small historical C extension, so test it against the exact PostgreSQL version and locale before adopting it.
