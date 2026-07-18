## Usage

Sources:

- [Official project or provider page](https://stripe.com)

`stripe` — Experimental generated SQL/PLPython SDK for calling the Stripe REST API from PostgreSQL.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpython3u`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "stripe";
SELECT extversion
FROM pg_extension
WHERE extname = 'stripe';
```

The upstream project is associated with `Stainless`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
