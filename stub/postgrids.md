## Usage

Sources:

- [Official extension control file](https://github.com/BIS-Brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/postgrids.control)
- [Official upstream documentation](https://github.com/BIS-Brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/README.md)
- [Official Rust package manifest](https://github.com/BIS-Brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/Cargo.toml)

`postgrids` — pgrx extension for parsing, formatting, and changing the precision of British OSGB and Irish OSI national grid references.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "postgrids";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgrids';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
