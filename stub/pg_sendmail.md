## Usage

Sources:

- [Official extension control file](https://github.com/sshutdownow/pg_sendmail/blob/f8c355d3ffcfc31302694d85c0740b2017c35f3f/pg_sendmail.control)

`pg_sendmail` — Sends email from SQL by invoking a sendmail-compatible executable at a fixed server path.

The reviewed catalog snapshot records version `0.0.2`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_sendmail";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sendmail';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
