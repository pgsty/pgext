## Usage

Sources:

- [Official extension control file](https://github.com/vog/texcaller/blob/fc5fbc9862a6df3e907e2ee409e5d995aa175ef6/postgresql/texcaller.control)
- [Official upstream documentation](https://vog.github.io/texcaller/group__postgresql.html)

`texcaller` — PostgreSQL functions for converting TeX or LaTeX source to DVI or PDF.

The reviewed catalog snapshot records version `0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "texcaller";
SELECT extversion
FROM pg_extension
WHERE extname = 'texcaller';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
