## Usage

Sources:

- [Official extension control file](https://github.com/tvondra/randomstring/blob/ae5259706484d66eb4dac63aaaea6ba745f7ffcd/randomstring.control)

`randomstring` — randomstring: generate pseudo-random text and bytea values

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "randomstring";
SELECT extversion
FROM pg_extension
WHERE extname = 'randomstring';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
