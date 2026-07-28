## Usage

Sources:

- [Official upstream README](https://github.com/borkdude/plsci/blob/ac7d249b7182c65ec462da4541b6287819a726ff/README.md)
- [Official extension control file (plsci.control)](https://github.com/borkdude/plsci/blob/ac7d249b7182c65ec462da4541b6287819a726ff/plsci.control)
- [Official implementation source](https://github.com/borkdude/plsci/blob/ac7d249b7182c65ec462da4541b6287819a726ff/src/lib.rs)

`plsci` — PostgreSQL procedural language handler for Clojure via SCI. Use it when database code must run in or interoperate with this procedural language. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION plsci;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `plsci` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
