## Usage

Sources:

- [Official extension control file](https://github.com/ttfkam/pg_geekspeak/blob/8d176e395f7f8a98a3c4dd2d48a1ee1e359b4e47/geekspeak.control)
- [Official upstream documentation](https://github.com/ttfkam/pg_geekspeak/blob/8d176e395f7f8a98a3c4dd2d48a1ee1e359b4e47/README.md)

`geekspeak` — Database schema extension for the GeekSpeak podcast website.

The reviewed catalog snapshot records version `1.0.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `btree_gist`, `isn`, `multicorn`, `pgcrypto`, `plpgsql`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "geekspeak";
SELECT extversion
FROM pg_extension
WHERE extname = 'geekspeak';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
