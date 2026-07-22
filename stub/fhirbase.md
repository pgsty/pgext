## Usage

Sources:

- [Official repository README](https://github.com/niquola/fhirbase/blob/a0a90ac9501e915daa22b72e76d97cec47e8be9d/README.md)
- [Version 1.0 extension script](https://github.com/niquola/fhirbase/blob/a0a90ac9501e915daa22b72e76d97cec47e8be9d/fhirbase--1.0.sql)
- [Prototype development SQL](https://github.com/niquola/fhirbase/tree/a0a90ac9501e915daa22b72e76d97cec47e8be9d/dev)

`fhirbase` is an unfinished 2014 prototype for representing and searching FHIR resources in PostgreSQL. The published version 1.0 extension script contains no schema objects or callable SQL API, while the repository's design notes and development files explore generated resource tables and search indexes. Do not treat those experiments as an installed extension contract.

### What Installation Provides

If the historical files are available to PostgreSQL, extension creation records the extension but the versioned SQL script does not create the patient tables, search functions, or indexes discussed in the repository:

```sql
CREATE EXTENSION fhirbase;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'fhirbase';

SELECT classid::regclass, objid, deptype
FROM pg_depend
WHERE refobjid = (
  SELECT oid FROM pg_extension WHERE extname = 'fhirbase'
);
```

The last query is an audit step: on the upstream 1.0 script there are no user-facing member objects to use. The control file names a module library, but the repository Makefile does not identify compiled object files that would expose a C API.

### Prototype Material

The `dev` directory sketches FHIR resource generation, JSON-backed datatypes, insert/update helpers, and string/token search indexes. The README discusses case- and accent-insensitive FHIR search parameters, but it is a design document headed as a different project and leaves its Usage section empty. These files are not installed by `fhirbase--1.0.sql`.

For an application that needs FHIR storage, derive a current schema from the intended FHIR release and validate resource versioning, references, terminology, search parameters, access control, and migrations independently. Copying the prototype SQL into production would create an unsupported fork rather than enabling version 1.0.

### Compatibility and Maintenance

The repository's last commit is from 2014 and the catalog has no documented active lifecycle. There is no authoritative compatibility matrix, upgrade path, preload requirement, or supported FHIR release in the installed SQL. The control file is relocatable, but that alone does not make the prototype operational.

Use this stub to identify the catalog entry and its absence of a usable upstream SQL surface. If a different project or packaged artifact also uses the name `fhirbase`, inspect its own control and extension SQL rather than importing APIs from this repository by name similarity.
