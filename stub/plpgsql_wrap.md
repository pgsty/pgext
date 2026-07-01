


## Usage

Sources: [README](https://github.com/HexaCluster/plpgsql_wrap/blob/v1.0/README.md), [v1.0 release](https://github.com/HexaCluster/plpgsql_wrap/releases/tag/v1.0), [control file](https://github.com/HexaCluster/plpgsql_wrap/blob/v1.0/plpgsql_wrap.control)

`plpgsql_wrap` provides an Oracle WRAP-style procedural language for PostgreSQL. Functions written with `LANGUAGE plpgsql_wrap` are validated as PL/pgSQL and then stored encrypted in `pg_proc.prosrc` as `PLPGSQLWRAP:1:<hex>`.

### Install With A Key

Build the extension with a 32-byte AES-256-GCM key:

```bash
export WRAP_KEY_HEX=$(openssl rand -hex 32)
make WRAP_KEY_HEX=$WRAP_KEY_HEX
sudo make install
```

Back up the key. Wrapped functions can only be unwrapped or restored safely when the correct compiled key is available.

Install the extension in each database that needs the language:

```sql
CREATE EXTENSION plpgsql_wrap; -- requires plpgsql
```

### Create Wrapped Functions

Use normal PL/pgSQL syntax with a different language name:

```sql
CREATE OR REPLACE FUNCTION public.calculate_bonus(emp_id int, yr int)
RETURNS numeric
LANGUAGE plpgsql_wrap
AS $$
DECLARE
  v_salary numeric;
BEGIN
  SELECT salary INTO v_salary FROM employees WHERE id = emp_id;
  RETURN v_salary * 0.15;
END;
$$;
```

The stored body is opaque:

```sql
SELECT substring(prosrc, 1, 32) AS wrapped_code
FROM pg_proc
WHERE proname = 'calculate_bonus';
```

### Dump, Restore, And Unwrap

`pg_dump` emits the encrypted `PLPGSQLWRAP:1:` blob. A restore on a server with the same compiled key works normally. A different key leaves the blob stored, but calls fail when the validator/authentication path cannot authenticate it.

Superusers can permanently unwrap a function when they know the key:

```sql
SELECT plpgsql_wrap.unwrap_procedure(
  'myhexkey',
  'public',
  'calculate_bonus',
  'emp_id int, yr int'
);
```

### Caveats

- Version 1.0 supports PostgreSQL 14-18.
- The control file requires `plpgsql` and superuser installation.
- This protects casual source inspection and dumps, but the compiled key is a critical secret. Treat package artifacts and build logs accordingly.
- Syntax is validated before encryption, so ordinary PL/pgSQL syntax errors abort `CREATE FUNCTION` before encrypted storage is written.
