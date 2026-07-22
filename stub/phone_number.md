## Usage

Sources:

- [Official README](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/README.md)
- [Cargo version and PostgreSQL feature](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/Cargo.toml)
- [SQL-facing type and generator](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/src/phone_number/mod.rs)
- [Input and output implementation](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/src/phone_number/implementation.rs)

`phone_number` 0.0.0 is a minimal pgrx prototype that defines a five-part `PhoneNumber` type and `create_random_number()`. It does not implement E.164 parsing, regional validation, formatting policy, carrier metadata, or real number allocation.

### Prototype API

The input parser accepts exactly five space-separated digit groups with lengths 1, 3, 3, 2, and 2:

```sql
CREATE EXTENSION phone_number;

SELECT '1 212 555 12 34'::phone_number;
SELECT create_random_number();
```

Output is formatted like `+1 212 555-12-34`. The random function fills each numeric component from a pseudorandom range; it does not confirm that country, city, exchange, or subscriber components are assigned or callable.

### Critical Type-I/O Defect

Version 0.0.0's output is not accepted by its own input parser: the parser requires spaces and rejects the `+` and hyphens emitted by output. PostgreSQL types are normally expected to round-trip through text, so this defect can break casts, copy workflows, logical interchange, and dump/restore of stored values. Do not persist this type until the input/output contract is fixed and a dump/restore regression test passes.

### Additional Boundaries

Cargo enables only PostgreSQL 16, the control file requires superuser, and upstream provides only a one-line README. The parser checks group length and numeric syntax but not meaningful numbering-plan constraints; values such as zero-filled components can pass. Generated values use a general pseudorandom generator and are neither globally unique nor cryptographically secure.

Use this extension only as a pgrx type-development example. Production phone data should retain a canonical text representation and use a maintained numbering-plan library in a controlled validation layer.
