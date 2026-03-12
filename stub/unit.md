

## Usage

> [unit: SI unit data type for PostgreSQL](https://github.com/df7cb/postgresql-unit)

The `unit` extension provides a data type for SI units, enabling dimensional analysis and unit conversion directly in SQL.

```sql
CREATE EXTENSION unit;

SELECT '9.81 m/s^2'::unit;
SELECT '120 km/h'::unit @ 'm/s' AS velocity;  -- 33.3333333333333 m/s
```

### Base Units

meter (m), kilogram (kg), second (s), ampere (A), kelvin (K), mole (mol), candela (cd), byte (B).

### Operators

| Operator | Description | Example |
|----------|-------------|---------|
| `+`, `-` | Add/subtract (same dimensions) | `'1 m'::unit + '50 cm'::unit` |
| `*`, `/` | Multiply/divide | `'5 kg'::unit * '9.81 m/s^2'::unit` |
| `^` | Exponentiation by integer | `'2 m'::unit ^ 3` |
| `@` | Convert to unit (returns unit) | `'2 MB/min'::unit @ 'GB/d'` |
| `@@` | Convert to unit (returns double precision) | `'1 km'::unit @@ 'm'` |

### Functions

Mathematical: `sqrt()`, `exp()`, `ln()`, `log2()`, `cbrt()`, `asin()`, `tan()`, etc.

Aggregates: `sum(unit)`, `avg(unit)`, `min(unit)`, `max(unit)`, `stddev()`, `variance()`.

### Input Formats

```sql
SELECT '3|4 m'::unit;            -- fractions: 0.75 m
SELECT '10:05:30 s'::unit;       -- time format: 36330 s
SELECT 'm⁻²'::unit;              -- Unicode superscripts
```

### Unit Conversion

```sql
SELECT '2 MB/min'::unit @ 'GB/d';       -- 2.88 GB/d
SELECT '1 hl'::unit @ '0.5 l';          -- 200 * 0.5 l
SELECT '100 degC'::unit @ 'degF';        -- Fahrenheit conversion
```

### Range Type

```sql
SELECT unitrange('earthradius_polar', 'earthradius_equatorial');
```

### Configuration

- `unit.byte_output_iec`: Binary prefixes (Ki, Mi, Gi)
- `unit.output_base_units`: Show only base units
- `unit.time_output_custom`: Format times using minutes/hours/days
- `unit.output_superscript`: Unicode superscript exponents
