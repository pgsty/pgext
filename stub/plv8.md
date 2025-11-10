
## Usage

```sql
CREATE EXTENSION plv8;

SELECT plv8_version();
SELECT plv8_info();

DO $$ plv8.elog(NOTICE, plv8.version); $$ LANGUAGE plv8;
```

Example:

```sql
CREATE FUNCTION plv8_test(keys TEXT[], vals TEXT[]) RETURNS JSON AS $$
    var o = {};
    for(var i=0; i<keys.length; i++){
        o[keys[i]] = vals[i];
    }
    return o;
$$ LANGUAGE plv8 IMMUTABLE STRICT;


SELECT plv8_test(ARRAY['name', 'age'], ARRAY['Tom', '29']);
```


## Build

Plv8 build breaks on EL10 (x86/arm) with the following problems:

- find g++ problem
- g++ 14 include `<algorithm>` problem
- lto problem, g++14 link time optimization issue



