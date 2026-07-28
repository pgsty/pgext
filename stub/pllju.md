## Usage

Sources:

- [Official upstream README](https://github.com/eugwne/pllj/blob/d5d63a265b60200a7606b5cb5c92824837fd95cf/README.md)
- [Official extension control file (pllju.control)](https://github.com/eugwne/pllj/blob/d5d63a265b60200a7606b5cb5c92824837fd95cf/pllju.control)
- [Official extension SQL (pllju--0.1.sql)](https://github.com/eugwne/pllj/blob/d5d63a265b60200a7606b5cb5c92824837fd95cf/pllju--0.1.sql)

`pllju` — LuaJIT FFI PostgreSQL language extension. Use it when database code must run in or interoperate with this procedural language. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pllju;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pllj_call_handler_u()` is an extension function and returns `language_handler`.
- `pllj_inline_handler_u(internal)` is an extension function and returns `VOID`.
- `pllj_validator_u(oid)` is an extension function and returns `VOID`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
