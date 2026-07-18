## Usage

Sources:

- [Official extension control file](https://github.com/satya10x/advcopy/blob/d7dfcb6fafca0ea0c4ad7b67fa06df662fec6fbc/advcopy.control)
- [Official upstream documentation](https://github.com/satya10x/advcopy/blob/d7dfcb6fafca0ea0c4ad7b67fa06df662fec6fbc/README.md)

`advcopy` — PL/Python PostgreSQL extension for copying data to and from S3.

The reviewed catalog snapshot records version `0.0.2`, kind `puresql`, and implementation language `Python`.
Install and validate the declared extension dependencies first: `plpython3u`.

```sql
CREATE EXTENSION "advcopy";
SELECT extversion
FROM pg_extension
WHERE extname = 'advcopy';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
