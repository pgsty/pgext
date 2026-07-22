## Usage

Sources:

- [Official README](https://github.com/shereshevsky/pg_grants_manager/blob/d77b782f71bb1f7417aff3424680e3de7c54e760/README.md)
- [Extension control file](https://github.com/shereshevsky/pg_grants_manager/blob/d77b782f71bb1f7417aff3424680e3de7c54e760/grants_manager.control)
- [Permission alignment implementation](https://github.com/shereshevsky/pg_grants_manager/blob/d77b782f71bb1f7417aff3424680e3de7c54e760/gm_align_permissions.sql)

`grants_manager` is a small PL/pgSQL toolkit for snapshotting table and sequence privileges into an editable matrix, reporting drift, and optionally applying the declared grants. Use it only under controlled administration: applying the matrix executes `REVOKE ALL` and new `GRANT` statements.

### Core Workflow

```sql
CREATE EXTENSION grants_manager;

CREATE ROLE app_reader;
CREATE TABLE public.orders(id bigint PRIMARY KEY);
GRANT SELECT ON public.orders TO app_reader;

SELECT gm_generate_current();
SELECT * FROM public.grants_manager;

UPDATE public.grants_manager
SET app_reader = ARRAY['r', 'w']
WHERE object_name = 'orders';

SELECT gm_align_permissions(p_execute := false);
SELECT gm_align_permissions(p_execute := true);
```

`gm_generate_current()` rebuilds `public.grants_manager` from current table and sequence ACLs, renaming the previous matrix to `grants_manager_old`. Each role becomes an array column. Edit those arrays, run `gm_align_permissions(false)` for notices only, review the proposed changes, and use `gm_align_permissions(true)` only when ready to change privileges.

### Main Objects

- `gm_generate_current()` creates the editable privilege snapshot.
- `gm_get_status()` reports current grants as `schema_name`, `object_name`, `grantee`, and `grants`.
- `gm_align_permissions(boolean)` reports or applies differences.
- `gm_translate(text)` maps PostgreSQL privilege names to ACL letters and back.

The ACL letters used by the implementation include `r` for SELECT, `w` for UPDATE, `a` for INSERT, `d` for DELETE, `D` for TRUNCATE, `x` for REFERENCES, `t` for TRIGGER, and `U` for USAGE.

### Caveats

Version `0.0.1` handles ordinary tables and sequences and does not claim complete coverage of every grantable object. The upstream README lists group-role alignment as unfinished. Snapshot generation replaces the working matrix, while execution revokes all privileges before re-granting the declared set. Test with representative ownership, quoted identifiers, schemas, role inheritance, and concurrent changes before using it in production.
