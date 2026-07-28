## Usage

Sources:

- [Official extension control file (plocamlu.control)](https://github.com/higuoxing/plocaml/blob/6cbd404a94fa659785a3b6891a1cd1ea9c180594/plocamlu.control)
- [Official extension SQL (plocamlu--1.0.sql)](https://github.com/higuoxing/plocaml/blob/6cbd404a94fa659785a3b6891a1cd1ea9c180594/plocamlu--1.0.sql)

`plocamlu` — PL/OCaml Procedural Language Handler for PostgreSQL. Use it when database code must run in or interoperate with this procedural language. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION plocamlu;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `plocamlu_call_handler()` is an extension function and returns `language_handler`.
- `plocamlu_inline_handler(internal)` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
