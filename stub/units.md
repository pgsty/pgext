## Usage

Sources:

- [Official README](https://github.com/freevryheid/units/blob/eb52d6246e6fda6668e3e50cb5048e80453dc8d0/README.md)
- [Official extension SQL](https://github.com/freevryheid/units/blob/eb52d6246e6fda6668e3e50cb5048e80453dc8d0/units--0.0.1.sql)

`units` is a small PL/pgSQL learning extension that stores length-conversion factors in a table and multiplies an input value by the selected factor. It is useful as a transparent example or a locally curated conversion table, not as a comprehensive units system.

### Core Workflow

The implementation treats both unit arguments as PostgreSQL regular expressions. Anchor the names to request an exact pair.

```sql
CREATE EXTENSION units;

SELECT unitsConvert(1, '^mile$', '^meter$');
SELECT unitsConvert(12, '^inch$', '^foot$');

SELECT msr, uin, niu, mop
FROM units
WHERE uin = 'meter' AND niu = 'foot';
```

### Objects

- `units(msr, uin, niu, mop)` stores the measurement family, input unit, output unit, and multiplication factor. Its primary key is `(uin, niu)`.
- `unitsConvert(numeric, text, text) RETURNS numeric` looks up a factor with `uin ~ input_pattern` and `niu ~ output_pattern`, then returns the supplied value multiplied by that factor.

The bundled data covers length conversions among kilometer, meter, centimeter, millimeter, micrometer, nanometer, mile, yard, foot, inch, and nautical mile.

### Caveats

Because lookup uses regular expressions and a non-strict `SELECT INTO`, a broad pattern can match several rows and select an unspecified one. Use anchored patterns such as `^mile$`; a missing match returns `NULL`. Review conversion precision and edit the table deliberately if it becomes application data.

The upstream README describes the project as a learning experiment, and version `0.0.1` provides no broader dimensional analysis, unit aliases, compound units, validation across measurement families, or documented license. Qualify the `units` table and `unitsConvert` function if the extension is installed outside the default schema.
