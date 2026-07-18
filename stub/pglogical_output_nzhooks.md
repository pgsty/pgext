## Usage

Sources:

- [Official extension control file](https://github.com/troels/pghooks/blob/f82688238e25bb220478404eeeb4d6a825f014a1/pglogical_output_nzhooks.control)

`pglogical_output_nzhooks` — Example pglogical output-plugin row-filter hook that keeps only changes for a hard-coded relation name

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pglogical`.

```sql
CREATE EXTENSION "pglogical_output_nzhooks";
SELECT extversion
FROM pg_extension
WHERE extname = 'pglogical_output_nzhooks';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
