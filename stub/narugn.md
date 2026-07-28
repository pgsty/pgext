## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/narugn/narugn-0.3.0/README.md)
- [Official extension control file (narugn.control)](https://api.pgxn.org/src/narugn/narugn-0.3.0/narugn.control)
- [Official extension SQL (narugn--0.3.0.sql)](https://api.pgxn.org/src/narugn/narugn-0.3.0/sql/narugn--0.3.0.sql)

`narugn` — Narugn is a lightweight distributed computer, composed by one or more cells that are connected locally. It requires PostgreSQL with the PL/Proxy extension. Use it for the corresponding analytical or storage workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION narugn;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `are_adjacent(s1 local_server , s2 local_server)` is an extension function and returns `boolean`.
- `cell_logic(payload IN text[] , walked IN cdt[] , origin_tick IN bigint , rpfp IN boolean DEFAULT false , z OUT bigint , t OUT timestamp with time zone , output OUT text)` is an extension function and returns `SETOF RECORD`.
- `cell_new_server(payload IN text[] , walked IN cdt[] , origin_tick IN bigint , rpfp IN boolean DEFAULT false , z OUT bigint , t OUT timestamp with time zone , output OUT text)` is an extension function and returns `SETOF RECORD`.
- `cell_ping(payload IN text[] , walked IN cdt[] , origin_tick IN bigint , rpfp IN boolean DEFAULT false , z OUT bigint , t OUT timestamp with time zone , output OUT text)` is an extension function and returns `SETOF RECORD`.
- `cell_rescan(payload IN text[] , walked IN cdt[] , origin_tick IN bigint , rpfp IN boolean DEFAULT false , z OUT bigint , t OUT timestamp with time zone , output OUT text)` is an extension function and returns `SETOF RECORD`.
- `cell_version(payload IN text[] , walked IN cdt[] , origin_tick IN bigint , rpfp IN boolean DEFAULT false , z OUT bigint , t OUT timestamp with time zone , output OUT text)` is an extension function and returns `SETOF RECORD`.
- `code_version()` is an extension function and returns `text`.
- `configure_cell(cell IN cds , local_connstr IN text)` is an extension function and returns `text`.
- `configure_cell(short_name IN text , full_name IN text , polygon IN polygon , connstr IN text , local_connstr IN text)` is an extension function and returns `text`.
- `display_cct(this_cell cds , origin_cell cds , t bigint)` is an extension function and returns `text`.
- `execute_sync(cell_function IN text , payload VARIADIC text[] DEFAULT '{}' , c OUT cds , z OUT bigint , dt OUT interval , output OUT text)` is an extension function and returns `SETOF RECORD`.
- `execute_sync_abs(cell_function IN text , payload VARIADIC text[] DEFAULT '{}' , c OUT cds , z OUT bigint , t OUT timestamp with time zone , output OUT text)` is an extension function and returns `SETOF RECORD`.
- `execute_sync_raw(cell_function IN text , payload IN text[] DEFAULT '{}' , ts IN timestamp with time zone DEFAULT clock_timestamp() , origin_cell IN cds DEFAULT NULL , origin_tick IN bigint DEFAULT NULL , walked IN cdt[] DEFAULT NULL , max_delay IN float DEFAULT 0.2 , c OUT cds…)` is an extension function and returns `SETOF RECORD`.
- `global2dbname(i_c IN cds)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.3.0`.
- Install the confirmed extension dependencies first: `plproxy`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
