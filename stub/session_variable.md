

## Usage

> [session_variable: Registration and manipulation of session variables and constants](https://github.com/splendiddata/session_variable)

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

-- Set variable value (returns previous value)
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
