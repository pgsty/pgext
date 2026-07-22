## Usage

Sources:

- [Official README](https://github.com/pgoracle/pgora-osql/blob/abb37e20238dc80d29a5440b62d453625e7e766b/README.md)
- [Version 9.4 extension SQL](https://github.com/pgoracle/pgora-osql/blob/abb37e20238dc80d29a5440b62d453625e7e766b/sql/pgosql--9.4.sql)
- [Official Makefile](https://github.com/pgoracle/pgora-osql/blob/abb37e20238dc80d29a5440b62d453625e7e766b/Makefile)

`pgosql` is an old compatibility experiment whose README describes a PL/SQL-like procedural language. Its extension SQL also creates a `sys` schema with catalog-like views and compatibility functions. The frozen upstream revision is internally inconsistent, so treat it as source for evaluation or migration archaeology rather than a production language runtime.

### Intended Installation and Objects

The control file requires superuser installation. On a build where the upstream library-name mismatch has been corrected, the intended entry point is:

```sql
CREATE EXTENSION pgosql;

SELECT sys.getdate(),
       sys.sysutcdatetime(),
       sys.isnull(NULL::integer, 42);

SELECT name, object_id, type_desc
FROM sys.objects
ORDER BY name;

SELECT lanname
FROM pg_language
WHERE lanname IN ('plosql', 'pltsql');
```

The SQL script defines `sys.tables`, `sys.views`, and `sys.objects`; time helpers `sys.sysdatetime()`, `sys.sysdatetimeoffset()`, `sys.sysutcdatetime()`, `sys.getdate()`, and `sys.getutcdate()`; overloaded `sys.isnull` functions; and an intended procedural language named `plosql`.

### Upstream Consistency Warning

Version `9.4` does not provide a reliable unmodified installation contract:

- The Makefile builds a shared module named `pgosql`, but the extension SQL registers language handlers from a module named `pgtsql`.
- The extension SQL creates `plosql`, while the repository regression tests exercise `LANGUAGE pltsql` and T-SQL-style constructs such as `@` variables and `PRINT`.
- The README calls the language PL/SQL, while the shipped compatibility objects and tests span different SQL dialect conventions.

As a result, `CREATE EXTENSION pgosql` can fail unless a downstream package patches these names or supplies the referenced module. Do not assume that a package carrying the same catalog version implements Oracle PL/SQL compatibility.

### Adoption Guidance

Inspect the exact downstream source and package patches, then test extension creation, handler-library resolution, language name, parser behavior, `sys` object semantics, privilege boundaries, and dump/restore. Migrate only code paths covered by regression tests. No preload is declared, but installation requires superuser privileges and the reviewed source should not be exposed to untrusted procedural code without a dedicated security review.
