## Usage

Sources:

- [Official extension control file](https://github.com/dvarrazzo/replisome/blob/775fe3766b8e433470c1048ac79334ed04b8caa2/replisome.control)
- [Official upstream documentation](https://github.com/dvarrazzo/replisome/blob/775fe3766b8e433470c1048ac79334ed04b8caa2/README.rst)

`replisome` — Logical-decoding output plugin and Python consumer framework for streaming PostgreSQL row changes as JSON.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "replisome";
SELECT extversion
FROM pg_extension
WHERE extname = 'replisome';
```

The curated lifecycle is `preview`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
