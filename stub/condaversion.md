## Usage

Sources:

- [Official extension control file](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_ext/condaversion.control)
- [Official upstream documentation](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_ext/README.md)

`condaversion` — Conda package version number data type implemented as a C++ PostgreSQL extension.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C++`.

```sql
CREATE EXTENSION "condaversion";
SELECT extversion
FROM pg_extension
WHERE extname = 'condaversion';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
