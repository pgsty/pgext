


## 用法

> [omni_python: 一等 Python 支持](https://docs.omnigres.org/omni_python/intro/)

`omni_python` 扩展将 Python 代码作为存储过程集成到 PostgreSQL 中。

### 基本 Python 函数

创建带 `@pg` 装饰器的 Python 文件：

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

### 加载 Python 文件

```sql
CREATE EXTENSION omni_python CASCADE;
CREATE EXTENSION omni_schema CASCADE;
CREATE EXTENSION omni_vfs CASCADE;

CREATE OR REPLACE FUNCTION demo_fs() RETURNS omni_vfs.local_fs
    LANGUAGE sql AS $$ SELECT omni_vfs.local_fs('/python-files') $$;

SELECT omni_schema.load_from_fs(demo_fs());
```

### Flask 集成

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

配置 HTTP 处理器：

```sql
UPDATE omni_httpd.handlers SET query = $$select handle(request.*) from request$$;
```
