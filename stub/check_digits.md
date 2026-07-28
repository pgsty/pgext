## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/check_digits/check_digits-0.3.3/README.md)
- [Official extension control file (check_digits.control)](https://api.pgxn.org/src/check_digits/check_digits-0.3.3/check_digits.control)
- [Official extension SQL (check_digits--0.3.2.sql)](https://api.pgxn.org/src/check_digits/check_digits-0.3.3/sql/check_digits--0.3.2.sql)

`check_digits` — At last to create the extension type in psql. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION check_digits;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `check_digits_inn(inn text)` is an extension function and returns `boolean`.
- `check_digits_isbn(isbn text)` is an extension function and returns `boolean`.
- `check_digits_ogrn(ogrn text)` is an extension function and returns `boolean`.
- `check_digits_okpo(okpo text)` is an extension function and returns `boolean`.
- `check_digits_snils(snils text)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `0.3.3`.
- The control file marks the extension as relocatable.
- The former GitHub repository URL returned 404 during the 2026-07-28 review; treat the pinned PGXN distribution above as the available source boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
