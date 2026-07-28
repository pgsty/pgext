## Usage

Sources:

- [Official upstream README](https://github.com/rvianello/ribollita/blob/b7b48c0b29d22166d7e42859e7b5b132a3e22b73/README.md)
- [Official extension control file (ribollita.control)](https://github.com/rvianello/ribollita/blob/b7b48c0b29d22166d7e42859e7b5b132a3e22b73/src/ribollita.control)
- [Official extension SQL (ribollita--0.0.sql)](https://github.com/rvianello/ribollita/blob/b7b48c0b29d22166d7e42859e7b5b132a3e22b73/src/ribollita--0.0.sql)

`ribollita` — Se non è zuppa, è pan bagnato. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION ribollita;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `boost_version()` is an extension function and returns `cstring`.
- `mol_add_hs(mol, explicit_only bool DEFAULT false, add_coords bool DEFAULT false, add_residue_info bool DEFAULT false)` is an extension function and returns `mol`.
- `mol_amw(mol)` is an extension function and returns `float8`.
- `mol_formal_charge(mol)` is an extension function and returns `integer`.
- `mol_fragments(mol, sanitize_fragments bool DEFAULT true, copy_conformers bool DEFAULT true)` is an extension function and returns `TABLE`.
- `mol_from_molblock(molblock cstring, sanitize bool DEFAULT true, remove_hs bool DEFAULT true, strict_parsing bool DEFAULT true)` is an extension function and returns `mol`.
- `mol_from_smarts(smiles cstring, allow_cxsmiles bool DEFAULT true, strict_cxsmiles bool DEFAULT true, parse_name bool DEFAULT false, merge_hs bool DEFAULT true)` is an extension function and returns `mol`.
- `mol_from_smiles(smiles cstring, sanitize bool DEFAULT true, allow_cxsmiles bool DEFAULT true, strict_cxsmiles bool DEFAULT true, parse_name bool DEFAULT false, remove_hs bool DEFAULT true, use_legacy_stereo bool DEFAULT true)` is an extension function and returns `mol`.
- `mol_in(cstring)` is an extension function and returns `mol`.
- `mol_kekulize(mol)` is an extension function and returns `mol`.
- `mol_out(mol)` is an extension function and returns `cstring`.
- `mol_recv(internal)` is an extension function and returns `mol`.
- `mol_remove_all_hs(mol, sanitize bool DEFAULT true)` is an extension function and returns `mol`.
- `mol_remove_hs(mol, implicit_only bool DEFAULT false, update_explicit_count bool DEFAULT false, sanitize bool DEFAULT true)` is an extension function and returns `mol`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
