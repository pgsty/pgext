## 用法

来源：

- [官方 README](https://github.com/autodatadirect/uniqueidentifier/blob/85976b564e5cd3362c30ba388b95c3df19bc4b34/README.uniqueidentifier)
- [官方扩展 SQL](https://github.com/autodatadirect/uniqueidentifier/blob/85976b564e5cd3362c30ba388b95c3df19bc4b34/uniqueidentifier--1.0.sql)
- [官方 C 实现](https://github.com/autodatadirect/uniqueidentifier/blob/85976b564e5cd3362c30ba388b95c3df19bc4b34/uniqueidentifier.c)

`uniqueidentifier` 是一个旧式 16 字节类 UUID 类型，提供生成、文本转换、比较、连接以及 B-tree/hash 运算符类。它最初编写于 2001–2003 年，后来才封装为 PostgreSQL 扩展；除非必须兼容这个特定旧类型，新模式应优先使用 PostgreSQL 内置 `uuid` 类型。

### 核心流程

原生库依赖 e2fsprogs 的 `libuuid` 实现。安装匹配的二进制后，创建扩展，并用 `newid()` 生成默认值：

```sql
CREATE EXTENSION uniqueidentifier;

CREATE TABLE legacy_entities (
  id uniqueidentifier PRIMARY KEY DEFAULT newid(),
  payload text
);

INSERT INTO legacy_entities(payload) VALUES ('example');

SELECT id, text(id)
FROM legacy_entities;
```

`newid(text)` 用于显式解析文本。扩展还在两个方向上都声明了隐式转换：

```sql
SELECT newid('550e8400-e29b-41d4-a716-446655440000');
SELECT '550e8400-e29b-41d4-a716-446655440000'::uniqueidentifier;
```

### SQL 接口

- `newid()` 使用 `libuuid` 生成值，是易变函数。
- `newid(text)` 解析文本，`text(uniqueidentifier)` 格式化值。
- `<`、`<=`、`=`、`<>`、`>=` 和 `>` 支持比较。
- 默认 B-tree 与 hash 运算符类支持主键、等值查找、排序和索引。
- `||` 重载允许值与 `text` 按任意参数顺序连接。

### 迁移与兼容性

- 该类型不是 PostgreSQL 内置 `uuid`，即使二者都占 16 字节且文本格式相似。在它们之间迁移前，应测试语义、排序、二进制、复制、驱动、转储恢复和类型转换差异。
- 与 `text` 的双向隐式转换可能改变重载解析并掩盖不清晰的类型边界。新 SQL 应使用显式转换，并在升级后测试表达式。
- 当前审阅的安装 SQL 为连接运算符硬编码了 `public`，并把辅助函数所有者改为 `postgres` 角色。安装前应根据目标模式和角色模型审查脚本。
- 不应根据 README 中 PostgreSQL 9.1 的封装说明推断当前版本兼容性。应针对目标大版本构建并测试准确的源码或软件包。
- 如果不需要兼容，应使用核心 `uuid` 和服务器支持的 UUID 生成器，并通过明确且经过验证的文本转换进行迁移。
