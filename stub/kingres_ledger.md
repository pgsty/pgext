## Usage

Sources:

- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_ledger/README.md)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_ledger/kingres_ledger.control)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_ledger/kingres_ledger--0.3.sql)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/kingres/README.md)

`kingres_ledger` 0.3 records changes in hash chains and Merkle blocks for append-only or updatable ledger tables. It detects tampering only when compared with a previously trusted digest kept outside the database; an internally consistent rewritten chain is not proof of integrity.

### Enable a ledger

The reviewed KingresDB source supports PostgreSQL 18 and 19 branches. Preload and restart; the extension owner must remain a superuser. Creating SQL objects without preload is insufficient for enabling a ledger. The following example permanently enables append-only mode.

```ini
shared_preload_libraries = 'kingres_ledger'
```

```sql
CREATE EXTENSION kingres_ledger;
CREATE TABLE public.events (id bigint, payload text);
SELECT kingres.ledger_enable('public.events'::regclass);
INSERT INTO public.events VALUES (1, 'record');
SELECT kingres.ledger_generate_digest();
SELECT * FROM kingres.ledger_verify();
```

### Verification and privileges

`kingres.ledger_enable(relation, 'updatable')` instead creates history for changed/deleted row versions. Enabling is irreversible and the mode cannot later be switched. Protected tables cannot be dropped or truncated through ordinary administration. Explicit execution grants are needed for administrative APIs and the caller must own the table. History access is separate from the base table’s RLS and column privileges.

`kingres.ledger_inclusion_proof` and `kingres.ledger_verify_proof` verify row membership. Export the result of `kingres.ledger_generate_digest()` to independent trusted storage: the copy recorded inside the database is not a trust anchor. New `kingres.ledger_head`, `kingres.ledger_blocks_entries` and `kingres.ledger_block_preimage` functions expose publication inputs; full external signing and anchoring are separate components.

### Maintenance boundary

There is an unprotected interval until committed rows enter a block and its digest leaves the database. Privileged catalog/file edits and disabled guards remain outside the protection boundary. Upgrade from 0.2 by deploying the matching library, restarting and applying `ALTER EXTENSION kingres_ledger UPDATE TO '0.3'`; the supplied migration preserves existing blocks and proofs.
