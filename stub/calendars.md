## Usage

Sources:

- [Official extension control file](https://github.com/macdice/calendars/blob/22d7900c371ad142505667c14c7d9605dccf077d/calendars.control)
- [Official upstream documentation](https://github.com/macdice/calendars/blob/22d7900c371ad142505667c14c7d9605dccf077d/README)

`calendars` — calendar conversion routines

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "calendars";
SELECT extversion
FROM pg_extension
WHERE extname = 'calendars';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
