


## Usage

> [omni_python: First-class Python support](https://docs.omnigres.org/omni_python/intro/)

The `omni_python` extension integrates Python code as stored procedures in PostgreSQL.

### Basic Python Function

Create a Python file with the `@pg` decorator:

```python
# hello.py
from omni_python import pg

@pg
def hello() -> str:
    return "Hey there!"

@pg
def add(a: int, b: int) -> int:
    return a + b
```

### Loading Python Files

```sql
CREATE EXTENSION omni_python CASCADE;
CREATE EXTENSION omni_schema CASCADE;
CREATE EXTENSION omni_vfs CASCADE;

CREATE OR REPLACE FUNCTION demo_fs() RETURNS omni_vfs.local_fs
    LANGUAGE sql AS $$ SELECT omni_vfs.local_fs('/python-files') $$;

SELECT omni_schema.load_from_fs(demo_fs());
```

### Flask Integration

```python
from omni_python import pg
from omni_http import omni_httpd
from omni_http.omni_httpd import flask
from flask import Flask, jsonify, request

app = Flask('myapp')

@app.route('/employees', methods=['GET'])
def get_employees():
    employees = plpy.execute(plpy.prepare("select * from employees"))
    return json.dumps([dict(e) for e in employees])

handle = pg(flask.Adapter(app))
```

Configure the HTTP handler:

```sql
UPDATE omni_httpd.handlers SET query = $$select handle(request.*) from request$$;
```
