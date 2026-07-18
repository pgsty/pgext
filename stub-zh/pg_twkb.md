## 用法

来源：

- [已复核 commit 的 README](https://github.com/nicklasaven/pg_twkb/blob/345cd979c9dafa804556d2aabdc872ca194d44f8/README.md)
- [扩展 control 文件](https://github.com/nicklasaven/pg_twkb/blob/345cd979c9dafa804556d2aabdc872ca194d44f8/pg_twkb.control)
- [SQL API 与索引 tile 实现](https://github.com/nicklasaven/pg_twkb/blob/345cd979c9dafa804556d2aabdc872ca194d44f8/pg_twkb--0.1.sql)
- [TWKB C 入口](https://github.com/nicklasaven/pg_twkb/blob/345cd979c9dafa804556d2aabdc872ca194d44f8/pg_twkb.c)
- [服务器端文件写入器](https://github.com/nicklasaven/pg_twkb/blob/345cd979c9dafa804556d2aabdc872ca194d44f8/lwout_twkb.c)

`pg_twkb` 是一个尚未完成、依赖 PostGIS 的索引 Tiny Well-Known Binary 数据集工具箱。`TWKB_IndexedTiles` 会把几何表递归划分为 tile，并把 tile ID 与 TWKB payload 组合起来；`TWKB_Collect` 用于组合已经编码的 payload 数组。

### 构建索引 Payload

对于可信、包含整数标识符和 PostGIS geometry 列的表：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_twkb;

SELECT TWKB_IndexedTiles(
  'public.features',
  'geom',
  'id',
  5,
  4326,
  1000
);
```

其他安装函数包括 `TWKB_MakeSquarebox`、`TWKB_DivideBox`、`TWKB_getTileId`、重载的 `TWKB_Write2File` 以及 `TWKB_Write2SQLite`。

### 注意事项

- 上游表示该仓库只是用于个人版本管理，相关功能还需要文档化后才能使用。版本 `0.1` 是陈旧实验代码，不是受支持的交换格式或生产工具。
- `TWKB_Write2File` 和 `TWKB_Write2SQLite` 会写入 PostgreSQL 服务器操作系统账号可见的路径。安装 SQL 没有提供特别保护；应从不可信角色撤销 `EXECUTE`，且绝不能接受调用方控制的路径。
- 两个 `TWKB_Write2File` 重载都有外部文件副作用，却被错误声明为 `IMMUTABLE`。不要在 PostgreSQL 可能折叠、重复或重排调用的表达式中使用它们。
- 索引 tile 函数根据表名和列名文本构造动态 SQL。只能提供可信标识符，并检查被引用关系上的权限。
- control 文件要求 `postgis`，而实现使用历史版本的 PostGIS 与 PostgreSQL C API，且没有当前兼容矩阵。请测试确切构建，并针对生成的自定义容器验证消费者。
