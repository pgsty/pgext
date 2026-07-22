## 用法

来源：

- [官方 README](https://github.com/postgrespro/jsonb_schema/blob/master/README.md)
- [扩展控制文件](https://github.com/postgrespro/jsonb_schema/blob/master/jsonb_schema.control)
- [1.0 版扩展 SQL](https://github.com/postgrespro/jsonb_schema/blob/master/jsonb_schema--1.0.sql)
- [上游回归测试示例](https://github.com/postgrespro/jsonb_schema/blob/master/sql/test.sql)

jsonb_schema 把 JSONB 值的结构模式与标量数据分离，并在数据库表中只保存一次每种不同模式。当大量文档具有相同形状时，它可能减少存储；但打包表示并非自包含，也不能像原始文档一样直接查询。

### 核心流程

存储前打包文档，需要原始 JSONB 值时再解包。

```sql
CREATE EXTENSION jsonb_schema;

CREATE TABLE compact_events (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    packed jsonb NOT NULL
);

INSERT INTO compact_events (packed)
VALUES
    (jsonb_pack('{"kind":"click","user":10,"tags":["a","b"]}'::jsonb)),
    (jsonb_pack('{"kind":"view","user":11,"tags":["c"]}'::jsonb));

SELECT id, jsonb_unpack(packed) AS document
FROM compact_events;
```

打包值是一个双元素 JSONB 数组，包含模式标识符和编码数据。共享的模式字节保存在扩展拥有的字典表中。

### 已安装对象

1.0 版创建模式字典表、模式字节与标识符的唯一索引、两个 C 层模式/数据转换函数，以及打包与解包包装函数。第一次使用某种模式时会插入新行，冲突时则复用已有标识符。

### 存储与生命周期边界

采用前应使用代表性数据测量：形状重复的文档可能受益，而形状高度多样的数据会承担字典和编码开销。针对原始 JSON 字段的索引或表达式必须先解包，因此普通 JSONB 查询与索引行为不会透明地作用于打包列。

打包行依赖字典表中完全一致的标识符映射。转储、恢复、复制或迁移时，必须把该表与全部打包值一起处理；只把 JSONB 列复制到另一个数据库，可能按错误模式解码或得不到值。不能删除仍被打包文档引用的字典行。扩展不提供引用计数或模式垃圾回收。

虽然打包以函数形式调用，但它会写入共享元数据，因此在事务、权限和并发规划中应把它视为写操作。该扩展可重定位，也没有声明预加载或重启要求。
