## Usage

Sources:

- [Official upstream documentation](https://github.com/ibarwick/config_log/blob/db66842eec9398701bd045fad7e6a09f6b77eeb5/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/config_log/)

`config_log` — custom background worker process to log changes to postgresql.conf to a table

The reviewed catalog snapshot records version `0.1.7`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "config_log";
SELECT extversion
FROM pg_extension
WHERE extname = 'config_log';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
