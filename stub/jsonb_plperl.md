

## Usage

> [jsonb_plperl: Transform between jsonb and PL/Perl](https://www.postgresql.org/docs/current/datatype-json.html)

Provides a transform for the `jsonb` type for PL/Perl. When loaded, `jsonb` values are automatically converted to Perl data structures (hashes, arrays, scalars) and vice versa.

```sql
CREATE EXTENSION jsonb_plperl;

CREATE FUNCTION jsonb_keys_sorted(val jsonb) RETURNS text
LANGUAGE plperl TRANSFORM FOR TYPE jsonb AS $$
  # val is now a Perl hash/array reference
  return join(', ', sort keys %$val);
$$;

CREATE FUNCTION build_jsonb(name text, age integer) RETURNS jsonb
LANGUAGE plperl TRANSFORM FOR TYPE jsonb AS $$
  my ($name, $age) = @_;
  return { name => $name, age => $age };
$$;

SELECT jsonb_keys_sorted('{"b": 2, "a": 1, "c": 3}'::jsonb);
SELECT build_jsonb('Alice', 30);
```
