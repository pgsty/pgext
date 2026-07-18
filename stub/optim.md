## Usage

Sources:

- [Official extension control file](https://github.com/ChristoferBerruz/hyper_db/blob/07c3b41e90ef434f29b3a97ac701bfce62e73363/optim/optim.control)

`optim` — Experimental pgrx extension defining optimization-direction and hyperparameter trial types.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "optim";
SELECT extversion
FROM pg_extension
WHERE extname = 'optim';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
