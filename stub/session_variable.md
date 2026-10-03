## Usage

Sources:

- [3.6 README](https://github.com/splendiddata/session_variable/blob/3.6/README.md)
- [3.6 SQL API](https://github.com/splendiddata/session_variable/blob/3.6/session_variable--3.6.sql)
- [3.6 control file](https://github.com/splendiddata/session_variable/blob/3.6/session_variable.control)
- [3.6 initialization code](https://github.com/splendiddata/session_variable/blob/3.6/session_variable.c)

`session_variable` defines database-level variables and constants with a separate value in each session. Changes to session-local values do not affect other connections.

### Creating Variables and Constants

```sql
CREATE EXTENSION session_variable;

-- Create a variable with initial value
SELECT session_variable.create_variable('my_var', 'text'::regtype, 'initial text'::text);

-- Create a variable with NULL initial value
SELECT session_variable.create_variable('my_date_var', 'date'::regtype);

-- Create a constant (cannot be changed via set())
SELECT session_variable.create_constant('my_env', 'text'::regtype, 'Production'::text);
```

### Getting and Setting Values

```sql
-- Get variable value (second arg is type hint)
SELECT session_variable.get('my_var', null::text);

-- Set variable value (returns true on success)
SELECT session_variable.set('my_var', 'new text'::text);
```

### Using in PL/pgSQL

```sql
DO $$
DECLARE
    my_field text;
BEGIN
    my_field := session_variable.get('my_var', my_field);
    RAISE NOTICE 'Value: %', my_field;
END
$$ LANGUAGE plpgsql;
```

### Administration Functions

```sql
-- Alter the initial/constant value (affects new sessions)
SELECT session_variable.alter_value('my_env', 'Development'::text);

-- Reload all variables from database definitions
SELECT session_variable.init();

-- Drop a variable or constant
SELECT session_variable.drop('my_var');

-- Check if a variable exists
SELECT session_variable.exists('my_var');

-- Get the type of a variable
SELECT session_variable.type_of('my_var');
```

### Key Behaviors

- Variables are defined at the database level; each session gets a local copy
- `set()` only changes the session-local copy; other sessions are unaffected
- `alter_value()` changes the stored value; new sessions see it, existing sessions need `init()` to refresh
- Constants cannot be changed via `set()`, only via `alter_value()`
- Variable and constant names must be unique across both types

### Privileges, Reads and Version Boundaries

A superuser installs this non-relocatable extension in schema `session_variable`; no shared preload or restart is required for the SQL workflow. Grant `session_variable_administrator_role` to manage definitions and `session_variable_user_role` for ordinary access. The administrator role includes the user role.

`session_variable.set` and `session_variable.alter_value` return boolean true on success, not the previous value. A successful stored-definition change becomes visible in the caller immediately and in future sessions after commit; existing sessions keep their local state until `session_variable.init()` resets all values. Definitions are stored in `session_variable.variables` and included in logical backups.

`session_variable.get_stable` may cache a value during a statement; use the regular getter when a trigger changes that value within the same statement. `session_variable.get_constant` is IMMUTABLE and can return cached results after administrative constant changes; use the regular getter during such changes.

Release 3.6 removes obsolete version-1 initialization support. The preceding 3.5 rejects non-null initial values for `collection` and `icollection` because text initialization can crash a backend. For those types, create a null-initialized variable and follow the source's initialization hook `session_variable.variable_initialisation()`; the README's alternate hook name is inconsistent with the implementation. Upstream lists PostgreSQL 14-18 and provisional PostgreSQL 19 support.
