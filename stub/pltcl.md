

## Usage

> [pltcl: PL/Tcl trusted procedural language](https://www.postgresql.org/docs/current/pltcl.html)

PL/Tcl allows writing PostgreSQL functions in the Tcl programming language. As a trusted language, it restricts access to the filesystem and other system resources.

```sql
CREATE EXTENSION pltcl;

-- Simple function returning a greeting
CREATE FUNCTION tcl_hello(name text) RETURNS text
LANGUAGE pltcl AS $$
  return "Hello, $1!"
$$;

SELECT tcl_hello('world');

-- Function with conditional logic
CREATE FUNCTION tcl_max(a integer, b integer) RETURNS integer
LANGUAGE pltcl AS $$
  if {$1 > $2} {return $1}
  return $2
$$;

-- Set-returning function
CREATE FUNCTION tcl_sequence(count integer) RETURNS SETOF integer
LANGUAGE pltcl AS $$
  for {set i 1} {$i <= $1} {incr i} {
    return_next $i
  }
$$;

-- Trigger function
CREATE FUNCTION tcl_audit() RETURNS trigger
LANGUAGE pltcl AS $$
  set NEW(modified_at) [clock format [clock seconds] -format "%Y-%m-%d %H:%M:%S"]
  return [array get NEW]
$$;
```

Database access from PL/Tcl uses the `spi_exec` command:

```sql
CREATE FUNCTION tcl_count_rows(tbl text) RETURNS integer
LANGUAGE pltcl AS $$
  spi_exec "SELECT count(*) AS cnt FROM $1"
  return $cnt
$$;
```
