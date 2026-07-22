## Usage

Sources:

- [Extension SQL](https://github.com/quentusrex/pg_sip/blob/master/pg_sip--0.0.1.sql)
- [Version implementation](https://github.com/quentusrex/pg_sip/blob/master/src/pg_sip_version.c)
- [SIP header prototype](https://github.com/quentusrex/pg_sip/blob/master/src/pg_sip_header.c)
- [Official repository](https://github.com/quentusrex/pg_sip)

`pg_sip` 0.0.1 is an incomplete prototype for exposing SIP message parsing in PostgreSQL. The extension's installed SQL surface contains only a version function. A header-extraction C function exists in the source tree but is not declared by the extension install script.

### Installed Workflow

```sql
CREATE EXTENSION pg_sip;

SELECT pg_sip_version();
-- 0.0.1

SELECT p.oid::regprocedure
FROM pg_proc AS p
JOIN pg_depend AS d
  ON d.classid = 'pg_proc'::regclass
 AND d.objid = p.oid
JOIN pg_extension AS e
  ON e.oid = d.refobjid
WHERE e.extname = 'pg_sip';
```

### API Boundary

- `pg_sip_version()` is immutable and strict and returns the extension version string.
- The source function for reading the first matching SIP header is exercised only by a regression script that creates a declaration manually. It is not part of CREATE EXTENSION and therefore is not a supported installed API.
- Roadmap items such as validation and broader SIP accessors are not implemented interfaces.

### Build and Safety Notes

- The parser prototype links against the bundled libre library through a repository submodule. A source archive without initialized submodules is incomplete.
- SIP messages are untrusted network input. Do not expose undeclared C symbols manually in production without fuzzing, memory-safety review, length limits, and target-major testing.
- Check the packaged extension SQL after every build; the presence of a compiled symbol does not make it an extension-owned database object.
