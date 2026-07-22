## 用法

来源：

- [官方 README](https://github.com/citusdata/pg_intpair/blob/master/README.md)
- [扩展控制文件](https://github.com/citusdata/pg_intpair/blob/master/intpair.control)
- [0.2 版扩展 SQL](https://github.com/citusdata/pg_intpair/blob/master/intpair--0.2.sql)
- [回归测试 SQL 示例](https://github.com/citusdata/pg_intpair/blob/master/sql/intpair.sql)

intpair 安装一种定长的紧凑基础类型，用于保存一对有符号 64 位整数。SQL 扩展名是 intpair，而上游项目与软件包名为 pg_intpair；当且仅当需要存储两个整数，并看重紧凑存储、比较和索引时，可用它替代复合类型或数组。

### 核心流程

可以用双参数函数或文本形式构造值，并用从零开始的下标访问两个分量。

```sql
CREATE EXTENSION intpair;

CREATE TABLE edges (
    endpoints int64pair NOT NULL
);

INSERT INTO edges VALUES
    (int64pair(10, 20)),
    ('(10,30)'::int64pair),
    ('(-1,5)'::int64pair);

SELECT endpoints,
       endpoints[0] AS first_value,
       endpoints[1] AS second_value
FROM edges
ORDER BY endpoints;
```

### 运算符与索引

0.2 版定义了相等比较，并先按第一个分量、再按第二个分量进行字典序排序。它提供默认的 B-tree 与哈希运算符类，因此普通索引可支持相等、范围谓词和排序。

```sql
CREATE INDEX edges_endpoints_idx ON edges (endpoints);

SELECT *
FROM edges
WHERE endpoints >= int64pair(10, 0)
  AND endpoints <  int64pair(11, 0);
```

### 迁移与注意事项

上游 README 演示了通过文本输入输出，从包含两个大整数的复合类型迁移。应先在副本上测试，因为输入必须符合扩展的整数对语法，而且隐式类型转换会改变后续查询的类型解析。

```sql
ALTER TABLE legacy_edges
ALTER COLUMN endpoints TYPE int64pair
USING endpoints::text::int64pair;
```

该类型始终保存一个 16 字节的整数对，不能表示变长列表。它可重定位，也没有声明预加载或重启要求。应用应使用示例所示的准确扩展名 intpair 和类型名，而不要根据 pg_intpair 软件包名推导它们。
