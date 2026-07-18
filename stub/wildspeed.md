## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/wildspeed/blob/a1b4e02f56e81596993848cd6692b0cd5b2e38d5/wildspeed.control)
- [Official upstream documentation](https://github.com/postgrespro/wildspeed/blob/a1b4e02f56e81596993848cd6692b0cd5b2e38d5/README.md)

`wildspeed` — GIN operator class for fast wildcard LIKE searches using permuterm indexing

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "wildspeed";
SELECT extversion
FROM pg_extension
WHERE extname = 'wildspeed';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
