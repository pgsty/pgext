

## Usage

> [hstore_plperl: Transform between hstore and PL/Perl](https://www.postgresql.org/docs/current/hstore.html)

Provides a transform for the `hstore` type for PL/Perl. When loaded, `hstore` values are automatically converted to Perl hashes and vice versa.

```sql
CREATE EXTENSION hstore_plperl;

CREATE FUNCTION hstore_to_text(val hstore) RETURNS text
LANGUAGE plperl TRANSFORM FOR TYPE hstore AS $$
  # val is now a Perl hash reference
  my $result = '';
  for my $key (sort keys %$val) {
    $result .= "$key => $val->{$key}\n";
  }
  return $result;
$$;

CREATE FUNCTION make_hstore(key text, value text) RETURNS hstore
LANGUAGE plperl TRANSFORM FOR TYPE hstore AS $$
  my ($k, $v) = @_;
  return { $k => $v };
$$;

SELECT hstore_to_text('a=>1, b=>2'::hstore);
SELECT make_hstore('color', 'blue');
```
