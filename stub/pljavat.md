## Usage

Sources:

- [gpdb/installation/pljavat.control](https://github.com/cloudberry-contrib/pljava/blob/d1a1c03507c1af6d8fce0c35a92169545e3a2a74/gpdb/installation/pljavat.control)
- [README.md](https://github.com/cloudberry-contrib/pljava/blob/d1a1c03507c1af6d8fce0c35a92169545e3a2a74/README.md)
- [gpdb/installation/pljavat--1.5.0.sql](https://github.com/cloudberry-contrib/pljava/blob/d1a1c03507c1af6d8fce0c35a92169545e3a2a74/gpdb/installation/pljavat--1.5.0.sql)
- [Makefile](https://github.com/cloudberry-contrib/pljava/blob/d1a1c03507c1af6d8fce0c35a92169545e3a2a74/Makefile)
- [COPYRIGHT](https://github.com/cloudberry-contrib/pljava/blob/d1a1c03507c1af6d8fce0c35a92169545e3a2a74/COPYRIGHT)

`pljavat` is the trusted-only PL/Java extension shipped by the Cloudberry fork. Its minimal SQL registers the Java call handler and the trusted `java` language. The control version remains 1.5.0 even though the separate full PL/Java control in this repository has a newer version.

### Core Workflow

```sql
CREATE EXTENSION pljavat;
SELECT lanname, lanpltrusted FROM pg_language WHERE lanname = 'java';
```

### Operational Boundaries

This entry targets the Cloudberry/Greenplum distribution and does not assert stock PostgreSQL support. Install the matching PL/Java native library and JVM configuration, then enable with superuser privileges. The language being trusted does not make extension installation unprivileged. The full PL/Java extension creates overlapping handler/language objects, so choose the appropriate installation path rather than combining both. This minimal SQL does not supply the full SQL/J repository-management schema.
