## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/harukat/pg_pageinspect_plus/blob/5e7970fe57b80a909b0d6d65b1de52bf4eb5156f/README.md)
- [Version 1.0 SQL objects](https://github.com/harukat/pg_pageinspect_plus/blob/5e7970fe57b80a909b0d6d65b1de52bf4eb5156f/pg_pageinspect_plus--1.0.sql)
- [C implementation](https://github.com/harukat/pg_pageinspect_plus/blob/5e7970fe57b80a909b0d6d65b1de52bf4eb5156f/pg_pageinspect_plus.c)
- [Extension control file](https://github.com/harukat/pg_pageinspect_plus/blob/5e7970fe57b80a909b0d6d65b1de52bf4eb5156f/pg_pageinspect_plus.control)

`pg_pageinspect_plus` augments `pageinspect` with helpers that decode raw tuple attributes into common PostgreSQL types. It installs into `pg_catalog` and supports PostgreSQL 10–18 in current upstream documentation.

```sql
CREATE EXTENSION pageinspect;
CREATE EXTENSION pg_pageinspect_plus;
SELECT tuple_data_parse(
         'public.sample'::regclass,
         tuple_data_split('public.sample'::regclass,
                          t_data, t_infomask, t_infomask2, t_bits)
       )
FROM heap_page_items(get_raw_page('public.sample', 0));
```

Additional helpers convert raw `bytea` to timestamp, timestamptz, interval, integer, bigint, boolean, text, float4, and float8. They expose physical representation, not a portable serialization format.

Page inspection bypasses normal row visibility and type access paths and may reveal deleted or sensitive values. Restrict execution to trusted administrators. Results depend on the exact table descriptor, PostgreSQL build, endianness, alignment, and page state; use on a copy when investigating corruption, and never write raw pages based on decoded output.
