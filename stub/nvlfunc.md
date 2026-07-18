## Usage

Sources:

- [Official extension control file](https://github.com/rongfengliang/nvl-pg-extension/blob/b2c153a2ff86e5fb2acb4b094448a8835700f031/nvlfunc.control)
- [Official upstream documentation](https://github.com/rongfengliang/nvl-pg-extension/blob/b2c153a2ff86e5fb2acb4b094448a8835700f031/README.md)

`nvlfunc` — Oracle-compatible NVL function

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "nvlfunc";
SELECT extversion
FROM pg_extension
WHERE extname = 'nvlfunc';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
