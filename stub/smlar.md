

## Usage

> [smlar](https://github.com/jirutka/smlar): Effective similarity search for PostgreSQL arrays.
> Source: [README](https://github.com/jirutka/smlar/blob/master/README)

The `smlar` extension provides effective similarity search on PostgreSQL arrays using configurable similarity formulas, GiST and GIN index support, and TF/IDF weighting.


--------

## Functions

```
float4 smlar(anyarray, anyarray)
```
Computes similarity of two arrays. Arrays should be the same type.

```
float4 smlar(anyarray, anyarray, bool useIntersect)
```
Computes similarity of two arrays of composite types. Composite type looks like:

```sql
CREATE TYPE type_name AS (element_name anytype, weight_name FLOAT4);
```

The `useIntersect` option points to use only intersected elements in the denominator.

```
float4 smlar(anyarray a, anyarray b, text formula)
```
Computes similarity of two arrays by a given formula. Predefined variables in formula:

- `N.i` -- number of common elements in both arrays (intersection)
- `N.a` -- number of unique elements in first array
- `N.b` -- number of unique elements in second array

Example:

```sql
SELECT smlar('{1,4,6}'::int[], '{5,4,6}');
SELECT smlar('{1,4,6}'::int[], '{5,4,6}', 'N.i / sqrt(N.a * N.b)');
-- These two calls are equivalent.
```

```
anyarray % anyarray
```
Returns true if similarity of the arrays is greater than the threshold limit.

```
text[] tsvector2textarray(tsvector)
```
Transforms tsvector type to text array.

```
anyarray array_unique(anyarray)
```
Sort and unique array.

```
float4 inarray(anyarray, anyelement)
```
Returns zero if second argument does not present in the first one and 1.0 in opposite case.

```
float4 inarray(anyarray, anyelement, float4, float4)
```
Returns fourth argument if second argument does not present in the first one and third argument in opposite case.


--------

## GUC Configuration Variables

```
smlar.threshold  FLOAT
```
Arrays with similarity lower than threshold are not similar by `%` operation.

```
smlar.persistent_cache  BOOL
```
Cache of global stat is stored in transaction-independent memory.

```
smlar.type  STRING
```
Type of similarity formula: `cosine` (default), `tfidf`, `overlap`.

```
smlar.stattable  STRING
```
Name of table storing set-wide statistic. Table should be defined as:

```sql
CREATE TABLE table_name (
    value   data_type UNIQUE,
    ndoc    int4 (or bigint)  NOT NULL CHECK (ndoc > 0)
);
```

A row with null value means total number of documents. Used only for `smlar.type = 'tfidf'`.

```
smlar.tf_method  STRING
```
Calculation method for term frequency. Values:
- `"n"` -- simple counting of entries (default)
- `"log"` -- 1 + log(n)
- `"const"` -- TF is equal to 1

Used only for `smlar.type = 'tfidf'`.

```
smlar.idf_plus_one  BOOL
```
If false (default), calculate idf as `log(d/df)`. If true, as `log(1+d/df)`. Used only for `smlar.type = 'tfidf'`.

It is highly recommended to add to `postgresql.conf`:

```
smlar.threshold = 0.6  # or any other value > 0 and < 1
```


--------

## GiST/GIN Index Support

The `%` and `&&` operations are supported with GiST and GIN indexes for many array types:

| Array Type | GIN operator class | GiST operator class |
|---|---|---|
| `bit[]` | `_bit_sml_ops` | |
| `bytea[]` | `_bytea_sml_ops` | `_bytea_sml_ops` |
| `char[]` | `_char_sml_ops` | `_char_sml_ops` |
| `cidr[]` | `_cidr_sml_ops` | `_cidr_sml_ops` |
| `date[]` | `_date_sml_ops` | `_date_sml_ops` |
| `float4[]` | `_float4_sml_ops` | `_float4_sml_ops` |
| `float8[]` | `_float8_sml_ops` | `_float8_sml_ops` |
| `inet[]` | `_inet_sml_ops` | `_inet_sml_ops` |
| `int2[]` | `_int2_sml_ops` | `_int2_sml_ops` |
| `int4[]` | `_int4_sml_ops` | `_int4_sml_ops` |
| `int8[]` | `_int8_sml_ops` | `_int8_sml_ops` |
| `interval[]` | `_interval_sml_ops` | `_interval_sml_ops` |
| `macaddr[]` | `_macaddr_sml_ops` | `_macaddr_sml_ops` |
| `money[]` | `_money_sml_ops` | |
| `numeric[]` | `_numeric_sml_ops` | `_numeric_sml_ops` |
| `oid[]` | `_oid_sml_ops` | `_oid_sml_ops` |
| `text[]` | `_text_sml_ops` | `_text_sml_ops` |
| `time[]` | `_time_sml_ops` | `_time_sml_ops` |
| `timestamp[]` | `_timestamp_sml_ops` | `_timestamp_sml_ops` |
| `timestamptz[]` | `_timestamptz_sml_ops` | `_timestamptz_sml_ops` |
| `timetz[]` | `_timetz_sml_ops` | `_timetz_sml_ops` |
| `varbit[]` | `_varbit_sml_ops` | |
| `varchar[]` | `_varchar_sml_ops` | `_varchar_sml_ops` |
