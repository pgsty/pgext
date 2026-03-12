

## Usage

> [jsonb_plperlu: Transform between jsonb and PL/Perl untrusted](https://www.postgresql.org/docs/current/datatype-json.html)

Provides a transform for the `jsonb` type for PL/Perl Untrusted. When loaded, `jsonb` values are automatically converted to Perl data structures (hashes, arrays, scalars) and vice versa.

```sql
CREATE EXTENSION jsonb_plperlu;

CREATE FUNCTION process_json_u(val jsonb) RETURNS text
LANGUAGE plperlu TRANSFORM FOR TYPE jsonb AS $$
  use JSON;
  # val is already a Perl data structure, re-encode with sorting
  return encode_json($val);
$$;

CREATE FUNCTION build_jsonb_u(name text, age integer) RETURNS jsonb
LANGUAGE plperlu TRANSFORM FOR TYPE jsonb AS $$
  my ($name, $age) = @_;
  return { name => $name, age => $age };
$$;

SELECT process_json_u('{"b": 2, "a": 1}'::jsonb);
SELECT build_jsonb_u('Alice', 30);
```
