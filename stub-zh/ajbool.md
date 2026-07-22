## 用法

来源：

- [官方 README](https://github.com/adjust/ajbool/blob/b5a47572713f475006fa81e45eda87cc4e965315/README.md)
- [0.0.3 版本扩展 SQL](https://github.com/adjust/ajbool/blob/b5a47572713f475006fa81e45eda87cc4e965315/ajbool--0.0.3.sql)
- [C 实现](https://github.com/adjust/ajbool/blob/b5a47572713f475006fa81e45eda87cc4e965315/src/ajbool.c)

`ajbool` 增加了一种单字节三态布尔类型，其存储值为 `t`、`f` 和 `u`。与 SQL `NULL` 不同，显式的 `u` 值可以参与等值、排序、哈希索引、B-tree 索引、唯一约束和主键。`ajbool` 列仍可另外保存 SQL `NULL`。

### 核心流程

```sql
CREATE EXTENSION ajbool;

CREATE TABLE feature_state (
    feature text PRIMARY KEY,
    state ajbool NOT NULL
);

INSERT INTO feature_state VALUES
    ('search', 't'),
    ('billing', 'f'),
    ('recommendations', 'u');

SELECT feature, state, state::boolean
FROM feature_state
ORDER BY state, feature;
```

类型转换会保留普通布尔值，并在显式未知值与 SQL 空值之间转换：

```sql
SELECT true::ajbool, false::ajbool, NULL::boolean::ajbool;
SELECT 't'::ajbool::boolean, 'f'::ajbool::boolean, 'u'::ajbool::boolean;
```

第一个查询返回 `t`、`f` 和 `u`；第二个查询返回 `true`、`false` 与 SQL `NULL`。

### 对象索引

- `ajbool` 是标量类型，文本输出为 `t`、`f` 或 `u`。
- `boolean AS ajbool` 是仅赋值类型转换，并把 SQL `NULL` 映射为 `u`。
- `ajbool AS boolean` 是隐式类型转换，并把 `u` 映射为 SQL `NULL`。
- 运算符 `=`、`<>`、`<`、`<=`、`>` 和 `>=` 支持比较。
- `btree_ajbool_ops` 与 `hash_ajbool_ops` 分别提供默认 B-tree 和哈希运算符类。

### 运维说明

版本 `0.0.3` 可重定位，且未声明依赖或预加载要求。输入解析只检查第一个字符：`t` 表示真，`f` 表示假，其他任何输入都会变成 `u`。如果不能接受拼写错误被静默转换为未知值，应在类型转换前校验应用输入。

不要混淆存储值 `u` 与 SQL `NULL`：`u = u` 是普通比较，而 `NULL = NULL` 的结果仍为未知。把 `u` 转换为 `boolean` 后会重新进入 SQL 三值逻辑，因此转换后的谓词可能需要处理 `IS NULL`。
