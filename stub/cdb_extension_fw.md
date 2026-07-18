## Usage

Sources:

- [Project README](https://github.com/rafatower/cdb_extension_fw/blob/a728599d718fce887ec53c788121bb27191cd7ef/README.md)
- [Extension control file](https://github.com/rafatower/cdb_extension_fw/blob/a728599d718fce887ec53c788121bb27191cd7ef/cdb_extension_fw.control)
- [Example SQL sources](https://github.com/rafatower/cdb_extension_fw/tree/a728599d718fce887ec53c788121bb27191cd7ef/src/sql/src)

`cdb_extension_fw` is an experimental CartoDB framework and demonstration extension rather than an application feature. The reviewed `current` build installs a fixed `cdb_extension_fw` schema with two PL/pgSQL functions that emit notices.

### Exercise the demonstration API

```sql
CREATE EXTENSION cdb_extension_fw;

SELECT cdb_extension_fw.foo();
SELECT cdb_extension_fw.bar();
```

Both calls return void and emit a notice containing their function name. The upstream build's installation check runs these same calls after creating the extension.

### Scope and maintenance

The framework separates SQL and Python sources from generated extension release files, but this repository contains only the two demonstration functions. It does not provide a general runtime API for packaging arbitrary data or configuration.

The control file uses the nonnumeric version `current`, fixes objects in the `cdb_extension_fw` schema, and marks the extension non-relocatable. The repository publishes no release and its own limitations make upgrade and downgrade correctness the extension developer's responsibility. Treat it as historical build-system source, inspect generated release artifacts before installation, and do not depend on the demonstration functions in production schemas.
