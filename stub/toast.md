## Usage

Sources:

- [Official upstream README](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/README.md)
- [Official extension control file (toast.control)](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/7_toast/toast.control)
- [Official extension SQL (toast--1.0.sql)](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/7_toast/sql/toast--1.0.sql)

`toast` — Tutorial 7 TOAST. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION toast;

SELECT hello();
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `compare_storage(text, text)` is an extension function and returns `void`.
- `force_detoast(text)` is an extension function and returns `text`.
- `text_size_info(text)` is an extension function and returns `int`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
