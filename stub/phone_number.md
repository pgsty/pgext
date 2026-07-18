## Usage

Sources:

- [phone_number README at the reviewed commit](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/README.md)
- [phone_number Cargo configuration at the reviewed commit](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/Cargo.toml)
- [phone_number SQL-facing type and function at the reviewed commit](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/src/phone_number/mod.rs)
- [phone_number input/output implementation at the reviewed commit](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/src/phone_number/implementation.rs)

`phone_number` 0.0.0 is a minimal pgrx prototype that defines a structured phone-number value and one SQL function, `create_random_number`.

```sql
CREATE EXTENSION phone_number;

SELECT create_random_number();
```

The input implementation accepts exactly five space-delimited numeric groups such as `1 212 555 12 34`. The output formatter renders a value such as `+1 212 555-12-34`.

### Caveats

- Version 0.0.0 and the one-line README indicate prototype maturity. The control file requires superuser.
- Cargo enables PostgreSQL 16 only; no other PostgreSQL feature or compatibility claim is present.
- Input and output formats do not round-trip: the parser rejects the plus sign and hyphens produced by the formatter. Do not persist textual output for later recasting without a normalization layer.
- Randomly generated components are neither validated real telephone assignments nor cryptographically secure identifiers.
