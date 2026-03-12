


## 用法

> [omni_web: 通用 Web 栈原语](https://docs.omnigres.org/omni_web/intro/)

`omni_web` 扩展提供 Web 栈实用函数，通常与 `omni_httpd` 配合使用。

### Cookies

将 `Cookie` 头部解析为键值对：

```sql
SELECT * FROM omni_web.cookies('PHPSESSID=298zf09hf012fh2; csrftoken=u32t4o3tb3gg43; _gat=1');
```

返回包含 `name` 和 `value` 列的表。

### 查询字符串

**`omni_web.parse_query_string(text)`** -- 将查询字符串解析为键值对数组：

```sql
SELECT omni_web.parse_query_string('key=value&a=1&a=2');
```

**`omni_web.param_get(parsed, key)`** -- 获取参数的第一个值：

```sql
SELECT omni_web.param_get(omni_web.parse_query_string('a=1&a=2'), 'a');  -- '1'
```

**`omni_web.param_get_all(parsed, key)`** -- 获取参数的所有值：

```sql
SELECT omni_web.param_get_all(omni_web.parse_query_string('a=1&a=2'), 'a');
-- 1
-- 2
```
