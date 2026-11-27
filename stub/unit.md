## Usage

Sources:

- [Official source](https://github.com/df7cb/postgresql-unit/blob/7.11/README.md)
- [Official source](https://github.com/df7cb/postgresql-unit/blob/7.11/unit.control)
- [Official source](https://github.com/df7cb/postgresql-unit/blob/7.11/Makefile)
- [Official source](https://github.com/df7cb/postgresql-unit/releases/tag/7.11)

`unit` 7.11 provides physical quantities with dimensional arithmetic and conversion. The package release keeps SQL/control version 7 and depends on `plpgsql`. It uses compiled code and installs unit-definition tables; install with administrative privileges. Shared preload is not required.

### Quantities and Conversion

```sql
CREATE EXTENSION unit;
SELECT '9.81 m/s^2'::unit;
SELECT '120 km/h'::unit @ 'm/s';
SELECT '1 km'::unit @@ 'm';
SELECT '5 kg'::unit * '9.81 m/s^2'::unit;
```

`@` returns text expressing the quantity in the requested unit; `@@` returns a double-precision magnitude. Addition and subtraction require compatible dimensions. Multiplication, division and integer powers combine dimensions. The eight base dimensions include the seven SI quantities plus bytes. Quantities use floating-point storage, not arbitrary-precision decimal arithmetic.

### Parsing and Output

Input supports fractions, time notation, prefixes and Unicode exponents. Version 7.11 additionally accepts underscores in numeric literals and updates the bundled GNU unit definitions to 3.26. Validate conversions where exact constants or temperature offsets matter.

`unit.byte_output_iec` selects binary byte prefixes. `unit.output_base_units`, `unit.time_output_custom` and `unit.output_superscript` control display. The extension supplies mathematical functions, aggregates and `unitrange`; use dimensionally compatible values.

### Maintenance

The `unit_prefixes` and `unit_units` tables may be created in a chosen installation schema, but the extension is not relocatable afterwards. Preserve locally customized definitions when upgrading. Install matching library and data files, reconnect sessions using the old library, and check the installed SQL version before choosing an available upgrade path; release 7.11 itself keeps version 7.
