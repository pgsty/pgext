

## Usage

> [bool_plperlu: Transform between bool and PL/Perl untrusted](https://www.postgresql.org/docs/current/plperl.html)

Provides a transform for the `bool` type for PL/Perl Untrusted. When loaded, PostgreSQL `boolean` values are automatically converted to Perl native boolean representations and vice versa, instead of being passed as strings.

```sql
CREATE EXTENSION bool_plperlu;

CREATE FUNCTION check_flag_u(val boolean) RETURNS text
LANGUAGE plperlu TRANSFORM FOR TYPE boolean AS $$
  # val is a native Perl boolean (1 or undef), not a string
  if ($_[0]) {
    return "flag is set";
  }
  return "flag is not set";
$$;

SELECT check_flag_u(true);   -- flag is set
SELECT check_flag_u(false);  -- flag is not set
```
