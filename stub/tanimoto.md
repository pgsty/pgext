## Usage

Sources:

- [Official extension control file](https://github.com/ergo70/tanimoto/blob/df31a86da4098c94c684167a789e34008d051e81/tanimoto.control)

`tanimoto` — Fast C implementation of the Tanimoto similarity coefficient for PostgreSQL bit-varying fingerprints.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "tanimoto";
SELECT extversion
FROM pg_extension
WHERE extname = 'tanimoto';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
