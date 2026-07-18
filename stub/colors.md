## Usage

Sources:

- [Official upstream documentation](https://github.com/zejn/pg-colors/blob/f963a30f408bf8cbbb07782478a1ba0dc93ab190/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/colors/)

`colors` — C functions for perceptual color-distance calculations in the CIE L*a*b* color space

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "colors";
SELECT extversion
FROM pg_extension
WHERE extname = 'colors';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
