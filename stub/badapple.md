## Usage

Sources:

- [Official extension control file](https://github.com/higuoxing/pg_badapple/blob/92ddbeea6fee794bfaad70f1b1d9c4bfab068c91/badapple.control)

`badapple` — Play "Bad Apple!!" in the psql client.

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "badapple";
SELECT extversion
FROM pg_extension
WHERE extname = 'badapple';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
