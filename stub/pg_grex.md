## Usage

Sources:

- [Official upstream README](https://github.com/rustedbytes/pg_grex/blob/f0c830c13d93fb039a6904312fb19356040b7aad/README.md)
- [Official extension control file (pg_grex.control)](https://github.com/rustedbytes/pg_grex/blob/f0c830c13d93fb039a6904312fb19356040b7aad/pg_grex.control)
- [Official implementation source](https://github.com/rustedbytes/pg_grex/blob/f0c830c13d93fb039a6904312fb19356040b7aad/src/lib.rs)

`pg_grex` — PostgreSQL extension for generating generalized regular expressions from a set of input strings. Implemented in Rust using the pgrx framework. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_grex;

CREATE TABLE product_skus (
    id  serial PRIMARY KEY,
    sku text
);

INSERT INTO product_skus (sku) VALUES
('PROD-100'), ('PROD-101'), ('PROD-102'),
('TEST-999'), ('BETA-001'), ('BETA-002');
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `grex_build` is an extension function.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
