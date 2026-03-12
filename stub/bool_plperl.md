

## Usage

> [bool_plperl: Transform between bool and PL/Perl](https://www.postgresql.org/docs/current/plperl.html)

Provides a transform for the `bool` type for PL/Perl. When loaded, PostgreSQL `boolean` values are automatically converted to Perl native boolean representations and vice versa, instead of being passed as strings.

```sql
CREATE EXTENSION bool_plperl;

CREATE FUNCTION check_flag(val boolean) RETURNS text
LANGUAGE plperl TRANSFORM FOR TYPE boolean AS $$
  # val is a native Perl boolean (1 or undef), not a string
  if ($_[0]) {
    return "flag is set";
  }
  return "flag is not set";
$$;

SELECT check_flag(true);   -- flag is set
SELECT check_flag(false);  -- flag is not set
```
