## Usage

Sources:

- [Official extension control file](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/postgres_ical.control)
- [Official upstream documentation](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/README.md)
- [Official Rust package manifest](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/Cargo.toml)

`postgres_ical` — parse RFC 5545 iCalendar text and URLs inside PostgreSQL

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "postgres_ical";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgres_ical';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
