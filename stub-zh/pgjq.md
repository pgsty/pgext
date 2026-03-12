

## 用法

> [pgjq: 在 PostgreSQL 中使用 jq JSON 处理语言](https://github.com/Florents-Tselai/pgJQ)

提供 `jqprog` 数据类型和 `jq()` 函数，用于对 `jsonb` 对象执行 jq 程序。

### 基本过滤

```sql
SELECT jq('[{"bar": "baz", "balance": 7.77}]'::jsonb, '.[0].bar');
-- "baz"
```

### 使用 `@@` 运算符

```sql
SELECT '[{"bar": "baz"}]' @@ '.[0].bar'::jqprog;
-- "baz"
```

### 复杂程序

```sql
SELECT jq('[true,false,[5,true,[true,[false]],false]]',
          '(..|select(type=="boolean")) |= if . then 1 else 0 end');
-- [1, 0, [5, 1, [1, [0]], 0]]

SELECT jq('[1,5,3,0,7]', '(.[] | select(. >= 2)) |= empty');
-- [1, 0]
```

### 传递参数

以 `jsonb` 对象传递动态参数，通过 `$var` 引用：

```sql
SELECT jq(
    '{"jobs": [{"id": 9, "ok": true}, {"id": 100, "ok": false}]}'::jsonb,
    '.jobs[] | select(.ok == $ok and .id == 100) | .',
    '{"ok": false}'
);
```

### 与 jsonpath 链式使用

```sql
SELECT jq('[{"cust":"baz","active":true,"trans":{"balance":100}}]',
          '(.[] | select(.active == true))') - '{trans}' @> '{"cust": "baz"}';
-- t
```

### 处理文件

```sql
SELECT jq(pg_read_file('/path/to/data.json'), '.[]');
```
