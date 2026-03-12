

## Usage

> [pltclu: PL/Tcl untrusted procedural language](https://www.postgresql.org/docs/current/pltcl.html)

PL/Tcl Untrusted provides full Tcl capabilities including filesystem access and external program execution. Only superusers can create functions in this language.

```sql
CREATE EXTENSION pltclu;

-- Read a file from the server filesystem
CREATE FUNCTION read_file(filename text) RETURNS text
LANGUAGE pltclu AS $$
  set fd [open $1 r]
  set content [read $fd]
  close $fd
  return $content
$$;

-- Execute an external command
CREATE FUNCTION run_command(cmd text) RETURNS text
LANGUAGE pltclu AS $$
  return [exec {*}$1]
$$;

-- Access environment variables
CREATE FUNCTION get_env(varname text) RETURNS text
LANGUAGE pltclu AS $$
  if {[info exists ::env($1)]} {
    return $::env($1)
  }
  return ""
$$;

SELECT get_env('HOME');
```
