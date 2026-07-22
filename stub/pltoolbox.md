## Usage

Sources:

- [Official README](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/README.md)
- [Official control file](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/pltoolbox.control)
- [Official installation SQL](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/pltoolbox--1.0.3.sql)
- [Official regression tests](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/sql/pltoolbox.sql)

`pltoolbox` is a collection of C utility functions intended for PostgreSQL stored procedures. Version `1.0.3` creates the non-relocatable `pst` schema and groups formatting, string, date, record, diff, and bitmap-set helpers there.

### Core Workflow

Install the extension, then schema-qualify its functions to avoid ambiguity with PostgreSQL built-ins.

```sql
CREATE EXTENSION pltoolbox;

SELECT pst.format('Hello %s! Now it is %s', 'world', now());
SELECT pst.concat_ws('-', 'alpha', 42, NULL, 'omega');
SELECT pst.record_get_field(ROW(10, 'ten'), 'f2');

WITH s AS (
  SELECT pst.intarray_to_bms(ARRAY[1, 3, 5]) AS bits
)
SELECT pst.bitmapset_num_members(bits), pst.bms_to_intarray(bits)
FROM s;
```

`pst.format` and `pst.sprintf` implement percent-style formatting. `pst.concat`, `pst.concat_ws`, `pst.concat_js`, and `pst.concat_sql` assemble values with different quoting rules.

### Object Index

- String helpers: `pst.reverse`, `pst.left`, `pst.right`, and `pst.chars_to_array`.
- Date helpers: `pst.next_day` and `pst.last_day`.
- Diff helpers: `pst.diff_string` and `pst.lc_substring`.
- Record helpers: `pst.record_expand`, `pst.record_get_field`, `pst.record_get_fields`, and `pst.record_set_fields`.
- Bitmap helpers: the `pst.bitmapset` type, membership and set-operation functions, integer-array casts, `pst.bitmapset_collect`, and union operator `||`.
- Formatting operator: `%%` is installed in both `public` and `pst`.

### Compatibility and Maintenance

The upstream README says that its documentation remains incomplete and that some functions later appeared in PostgreSQL core with behavior that is not fully equal. In particular, the project's `left` and `right` semantics changed historically for negative lengths. Always schema-qualify calls when the same name exists in core, and test the edge cases used by application code.

This extension is not a general procedural language despite its name. The catalog marks it abandoned, and its C implementation depends on PostgreSQL internals such as bitmap sets and tuple layouts. Pin a tested server build; do not assume binary or behavioral compatibility across major PostgreSQL releases.
