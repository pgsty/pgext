## Usage

Sources:

- [Official upstream README](https://github.com/jerrysievert/pllisp/blob/219df653a6b736626e45898e819852949612af93/README.md)
- [Official extension control file (pllisp.control)](https://github.com/jerrysievert/pllisp/blob/219df653a6b736626e45898e819852949612af93/pllisp.control)
- [Official extension SQL (pllisp--0.1.0.sql)](https://github.com/jerrysievert/pllisp/blob/219df653a6b736626e45898e819852949612af93/pllisp--0.1.0.sql)

`pllisp` — This project started as what was supposed to be a humorous quip for a presentation on building extensions to Postgres, and then I went ahead and started it. Use it when database code must run in or interoperate with this procedural language. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pllisp;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pllisp_call_handler()` is an extension function and returns `language_handler`.
- `pllisp_inline_handler(internal)` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
