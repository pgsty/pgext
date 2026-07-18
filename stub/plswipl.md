## Usage

Sources:

- [Pinned upstream README and example](https://github.com/salva/plswipl/blob/769a10e0ab37ca0bfdd1da84c2632af9fc9fb365/README.md)
- [Pinned extension control file](https://github.com/salva/plswipl/blob/769a10e0ab37ca0bfdd1da84c2632af9fc9fb365/plswipl.control)

`plswipl` embeds SWI-Prolog as an untrusted PostgreSQL procedural language. A function body names a server-side Prolog source file, and the language invokes predicates from that file to produce PostgreSQL results.

Given a reviewed server-side file that defines predicate `break_chars/2`, the upstream example uses this shape:

```sql
CREATE EXTENSION plswipl;

CREATE FUNCTION break_chars(text)
RETURNS SETOF text
AS $pl$"/srv/plswipl/example.prolog"$pl$
LANGUAGE plswipl;

SELECT * FROM break_chars('abc');
```

The control file fixes extension objects in `pg_catalog`, declares the extension non-relocatable and superuser-only, and installs an untrusted language. The named file is read in the database server's security context; never allow untrusted users to choose or modify that path or its contents.

Upstream explicitly labels version `0.1` a barely working work in progress. Its README lists data conversion, arrays/JSON, `SPI`, database-stored modules, trusted execution, schema/session-user behavior, and exception handling as unfinished. Do not deploy it on a production server without a full native-code audit, backend-crash tests, strict filesystem containment, and remediation of the unfinished validator, trigger, and error paths.
