## Usage

Sources:

- [Official documentation](https://github.com/dpage/plpy_venv/blob/0fc90895033be8a10351a6383a275bcf09f2e404/README.md)
- [Extension control file](https://github.com/dpage/plpy_venv/blob/0fc90895033be8a10351a6383a275bcf09f2e404/plpy_venv.control)
- [Official repository](https://github.com/dpage/plpy_venv)

`plpy_venv` Select and manage Python virtual environments for PL/Python functions.

### Enablement

Install the files for the intended server, then create `plpy_venv` in the target database:

```sql
CREATE EXTENSION plpy_venv CASCADE;
```

The reviewed control or official workflow requires `plpython3u`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
plpy=# SELECT plpy_venv.create_venv('myvenv');
ERROR:  plpy.Error: Virtual environment directory /path/to/postgresql/data/venvs/myvenv already exists.
CONTEXT:  Traceback (most recent call last):
  PL/Python function "create_venv", line 30, in <module>
    plpy.error('Virtual environment directory {} already exists.'.format(venv_dir))
PL/Python function "create_venv"
```

### Main Objects

The official sources define the extension surface dynamically or through provider tooling; inspect the installed version before granting access.

### Operations and Boundaries

- The extension fixes or creates schema objects under `plpy_venv`; include them in privilege and backup review.
- Catalog lifecycle is abandoned; test upgrades, dump/restore, and server compatibility before production use.
