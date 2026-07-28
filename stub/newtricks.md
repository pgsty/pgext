## Usage

Sources:

- [Official upstream README](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/newtricks/README.md)
- [Official extension control file (newtricks.control)](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/newtricks/newtricks.control)
- [Official implementation source](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/newtricks/src/lib.rs)

`newtricks` — Examples of postgres extensions using rust. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION newtricks;

SELECT is_valid_fiscal_code('RSSMRA85T10A562S');

CREATE TABLE users (id INT, cf TEXT CHECK(is_valid_fiscal_code(cf)));
INSERT INTO users (id, cf) VALUES (1, 'asd');
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `emojify` is an extension function.
- `is_secure` is an extension function.
- `is_valid_fiscal_code` is an extension function.
- `list_emojis()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
