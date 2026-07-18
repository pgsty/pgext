## Usage

Sources:

- [Official extension control file](https://github.com/knizhnik/online_advisor/blob/e24652f824ffac872141634394ea76320fbae4f0/online_advisor.control)
- [Official upstream documentation](https://github.com/knizhnik/online_advisor/blob/e24652f824ffac872141634394ea76320fbae4f0/README.md)

`online_advisor` — Online index advisor for PostgreSQL workload-based index, statistics, and prepared statement recommendations.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "online_advisor";
SELECT extversion
FROM pg_extension
WHERE extname = 'online_advisor';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
