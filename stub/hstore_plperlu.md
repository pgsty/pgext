

## Usage

> [hstore_plperlu: Transform between hstore and PL/Perl untrusted](https://www.postgresql.org/docs/current/hstore.html)

Provides a transform for the `hstore` type for PL/Perl Untrusted. When loaded, `hstore` values are automatically converted to Perl hashes and vice versa.

```sql
CREATE EXTENSION hstore_plperlu;

CREATE FUNCTION hstore_to_json_u(val hstore) RETURNS text
LANGUAGE plperlu TRANSFORM FOR TYPE hstore AS $$
  use JSON;
  # val is now a Perl hash reference
  return encode_json($val);
$$;

CREATE FUNCTION make_hstore_u(key text, value text) RETURNS hstore
LANGUAGE plperlu TRANSFORM FOR TYPE hstore AS $$
  my ($k, $v) = @_;
  return { $k => $v };
$$;

SELECT hstore_to_json_u('a=>1, b=>2'::hstore);
SELECT make_hstore_u('color', 'blue');
```
