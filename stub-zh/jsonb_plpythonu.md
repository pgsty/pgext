## 用法

来源：

- [上游 README](https://github.com/postgrespro/jsonb_plpython/blob/dff6c22394f97ed0618ba2c8e350ad2e1779ca64/README.md)
- [扩展 control 文件](https://github.com/postgrespro/jsonb_plpython/blob/dff6c22394f97ed0618ba2c8e350ad2e1779ca64/jsonb_plpythonu.control)
- [SQL 安装脚本](https://github.com/postgrespro/jsonb_plpython/blob/dff6c22394f97ed0618ba2c8e350ad2e1779ca64/jsonb_plpythonu--1.0.sql)
- [PL/Python 3 回归示例](https://github.com/postgrespro/jsonb_plpython/blob/dff6c22394f97ed0618ba2c8e350ad2e1779ca64/sql/jsonb_plpython3u.sql)

`jsonb_plpythonu` `1.0` 版定义 PostgreSQL `jsonb` 与 PL/Python 值之间的转换。显式指定 `TRANSFORM FOR TYPE jsonb` 的函数会收到 Python 字典、列表、字符串、数字、布尔值或 `None`，而不是默认文本表示，也可以把这些值作为 `jsonb` 返回。

### 使用转换

```sql
CREATE EXTENSION jsonb_plpythonu CASCADE;
CREATE FUNCTION jsonb_key_count(value jsonb) RETURNS integer
LANGUAGE plpythonu
TRANSFORM FOR TYPE jsonb
AS $$
return len(value)
$$;
SELECT jsonb_key_count('{"a":1,"b":2}'::jsonb);
```

这个具体扩展依赖旧式无版本 `plpythonu` 语言，其 control/模块元数据还引用 Python 2 时代的库名。现代 PostgreSQL 安装通常应改用同仓库的 `jsonb_plpython3u` 变体。PL/Python 是不可信语言，因此只有超级用户能创建其函数。仓库自 2017 年后没有变化；应验证 Python/PostgreSQL ABI 兼容性及数字/容器往返，不要仅为运行当前 Python 代码加载旧变体。
