## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/README.md)
- [SQL function definition](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/sql/rpg_base36.sql)
- [Rust implementation](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/src/lib.rs)
- [Extension control file](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/rpg_base36.control)

`rpg_base36` is a toy Rust extension that converts a PostgreSQL integer to uppercase base 36. Negative inputs retain a leading minus sign.

```sql
CREATE EXTENSION rpg_base36;

SELECT base36_encode(123456);
SELECT base36_encode(-35);
```

### Caveats

The extension control version is 0.0.1, while the Cargo package reports 0.1.0; no release establishes a newer extension version. This project predates `pgrx` and uses the old `rpgffi` binding plus C glue, so it should be treated as a historical demonstration rather than a maintained production package. Its SQL file also contains a stale `CREATE EXTENSION rust_base36` guard message even though the actual extension name is `rpg_base36`.
