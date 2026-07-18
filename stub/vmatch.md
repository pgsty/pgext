## Usage

Sources:

- [Official upstream documentation](https://github.com/nuko-yokohama/vmatch/blob/6fd255384005a84a3bd8e05dc2e099eb60fa1bdf/readme_jp.txt)

`vmatch` — Text similarity functions and a fuzzy-match /= operator with typo-aware weighting.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "vmatch";
SELECT extversion
FROM pg_extension
WHERE extname = 'vmatch';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
