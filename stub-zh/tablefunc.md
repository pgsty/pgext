

## 用法

> [tablefunc: 操作整个表的函数，包括交叉表(crosstab)](https://www.postgresql.org/docs/current/tablefunc.html)

`tablefunc` 模块提供了返回表（多行）的函数。
它是一个受信任的扩展，拥有 `CREATE` 权限的非超级用户即可使用。

```sql
CREATE EXTENSION tablefunc;
```

### normal_rand -- 生成随机值

生成一组正态分布的随机值（高斯分布）：

```sql
SELECT * FROM normal_rand(1000, 5, 3);
-- numvals: 值的数量, mean: 均值 5, stddev: 标准差 3
```

### crosstab(text) -- 单参数透视

将数据从长格式转换为宽格式。SQL 必须返回 `row_name`、`category` 和 `value` 列：

```sql
CREATE TABLE ct(id SERIAL, rowid TEXT, attribute TEXT, value TEXT);
INSERT INTO ct(rowid, attribute, value) VALUES
  ('test1','att1','val1'), ('test1','att2','val2'),
  ('test1','att3','val3'), ('test2','att1','val5'),
  ('test2','att2','val6'), ('test2','att3','val7');

SELECT *
FROM crosstab(
  'SELECT rowid, attribute, value FROM ct ORDER BY 1,2'
) AS ct(row_name text, category_1 text, category_2 text, category_3 text);

 row_name | category_1 | category_2 | category_3
----------+------------+------------+------------
 test1    | val1       | val2       | val3
 test2    | val5       | val6       | val7
```

输入查询应始终使用 `ORDER BY 1,2` 以确保正确分组。
超出可用值范围的输出列将填充 null。

### crosstab(text, text) -- 双参数透视（含类别列表）

处理某些分组可能缺少部分类别数据的情况：

```sql
CREATE TABLE sales(year int, month int, qty int);
INSERT INTO sales VALUES(2007,1,1000),(2007,2,1500),(2007,7,500),
  (2007,11,1500),(2007,12,2000),(2008,1,1000);

SELECT * FROM crosstab(
  'SELECT year, month, qty FROM sales ORDER BY 1',
  'SELECT m FROM generate_series(1,12) m'
) AS (
  year int, "Jan" int, "Feb" int, "Mar" int, "Apr" int,
  "May" int, "Jun" int, "Jul" int, "Aug" int,
  "Sep" int, "Oct" int, "Nov" int, "Dec" int
);

 year | Jan  | Feb  | Mar | Apr | May | Jun | Jul | Aug | Sep | Oct | Nov  | Dec
------+------+------+-----+-----+-----+-----+-----+-----+-----+-----+------+------
 2007 | 1000 | 1500 |     |     |     |     | 500 |     |     |     | 1500 | 2000
 2008 | 1000 |      |     |     |     |     |     |     |     |     |      |
```

源 SQL 可以在 `row_name` 与 `category`/`value` 之间包含"额外"列。

### crosstab2, crosstab3, crosstab4 -- 预定义封装函数

预构建的封装函数，无需编写 FROM 子句（仅支持文本输入/输出）：

```sql
SELECT * FROM crosstab3(
  'SELECT rowid, attribute, value FROM ct ORDER BY 1,2'
);
```

为其他类型创建自定义封装函数：

```sql
CREATE TYPE my_crosstab_float8_5_cols AS (
    my_row_name text, c1 float8, c2 float8, c3 float8, c4 float8, c5 float8
);

CREATE OR REPLACE FUNCTION crosstab_float8_5_cols(text)
    RETURNS setof my_crosstab_float8_5_cols
    AS '$libdir/tablefunc','crosstab' LANGUAGE C STABLE STRICT;
```

### connectby -- 层级树形展示

展示存储在表中的层级数据（含键字段和父键字段）：

```sql
CREATE TABLE connectby_tree(keyid text, parent_keyid text, pos int);
INSERT INTO connectby_tree VALUES
  ('row1',NULL,0), ('row2','row1',0), ('row3','row1',0),
  ('row4','row2',1), ('row5','row2',0), ('row6','row4',0),
  ('row7','row3',0), ('row8','row6',0), ('row9','row5',0);

-- 带分支显示和排序
SELECT * FROM connectby('connectby_tree', 'keyid', 'parent_keyid', 'pos', 'row2', 0, '~')
  AS t(keyid text, parent_keyid text, level int, branch text, pos int);

 keyid | parent_keyid | level |       branch        | pos
-------+--------------+-------+---------------------+-----
 row2  |              |     0 | row2                |   1
 row5  | row2         |     1 | row2~row5           |   2
 row9  | row5         |     2 | row2~row5~row9      |   3
 row4  | row2         |     1 | row2~row4           |   4
 row6  | row4         |     2 | row2~row4~row6      |   5
 row8  | row6         |     3 | row2~row4~row6~row8 |   6
```

参数格式：`connectby(relname, keyid_fld, parent_keyid_fld [, orderby_fld], start_with, max_depth [, branch_delim])`

- `start_with`：起始行的键值（文本类型）
- `max_depth`：最大下探深度（0 = 不限制）
- `branch_delim`：分支输出的分隔符（可选，默认 `~`）
- 建议对父键字段建立索引，以提升大表查询性能
