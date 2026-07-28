## Usage

Sources:

- [Official upstream README](https://github.com/greenbea/pg_phonenumber/blob/bb0ed97ddb393445b6bed288fb99f7c113118ddc/README.md)
- [Official extension control file (pg_phonenumber.control)](https://github.com/greenbea/pg_phonenumber/blob/bb0ed97ddb393445b6bed288fb99f7c113118ddc/pg_phonenumber.control)
- [Official extension SQL (pg_phonenumber--1.0.sql)](https://github.com/greenbea/pg_phonenumber/blob/bb0ed97ddb393445b6bed288fb99f7c113118ddc/pg_phonenumber--1.0.sql)

`pg_phonenumber` — C++ libphonenumber-backed phone type, equality operators, and supported-region helpers. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_phonenumber;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `get_supported_calling_codes()` is an extension function and returns `INT[]`.
- `get_supported_regions()` is an extension function and returns `TEXT[]`.
- `phone_eq(phone, phone)` is an extension function.
- `phone_in(cstring, oid, integer)` is an extension function and returns `phone`.
- `phone_ne(phone, phone)` is an extension function.
- `phone_out(phone)` is an extension function and returns `cstring`.
- `phone` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
