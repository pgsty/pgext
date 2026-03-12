

## 用法

> [plpython3u: PL/Python3 不受信过程语言](https://www.postgresql.org/docs/current/plpython.html)

PL/Python3U 允许使用 Python 3 编写 PostgreSQL 函数。作为不受信语言，它可以完全访问 Python 生态系统。只有超级用户可以创建函数。

```sql
CREATE EXTENSION plpython3u;

-- 简单函数
CREATE FUNCTION py_hello(name text) RETURNS text
LANGUAGE plpython3u AS $$
  return f"Hello, {name}!"
$$;

SELECT py_hello('world');

-- 使用 Python 标准库
CREATE FUNCTION py_sha256(data text) RETURNS text
LANGUAGE plpython3u AS $$
  import hashlib
  return hashlib.sha256(data.encode()).hexdigest()
$$;

-- 返回复合类型
CREATE TYPE address AS (street text, city text, zip text);

CREATE FUNCTION parse_address(raw text) RETURNS address
LANGUAGE plpython3u AS $$
  import re
  m = re.match(r'(.+),\s*(.+)\s+(\d{5})', raw)
  if m:
    return (m.group(1), m.group(2), m.group(3))
  return None
$$;

-- 集合返回函数
CREATE FUNCTION py_generate_dates(start text, days int) RETURNS SETOF date
LANGUAGE plpython3u AS $$
  from datetime import datetime, timedelta
  d = datetime.strptime(start, '%Y-%m-%d')
  for i in range(days):
    yield (d + timedelta(days=i)).strftime('%Y-%m-%d')
$$;

-- 通过 plpy 访问数据库
CREATE FUNCTION py_row_count(table_name text) RETURNS bigint
LANGUAGE plpython3u AS $$
  result = plpy.execute(f"SELECT count(*) AS cnt FROM {table_name}")
  return result[0]['cnt']
$$;

-- 使用外部包（必须在服务器上安装）
CREATE FUNCTION py_parse_json(url text) RETURNS jsonb
LANGUAGE plpython3u AS $$
  import json, urllib.request
  response = urllib.request.urlopen(url)
  data = json.loads(response.read())
  return json.dumps(data)
$$;
```
