## 用法

来源：

- [官方 1.0 版 SQL](https://github.com/overplumbum/postgresql-zlib/blob/74f0e5126266105fea05dcd248400e73bc7bacc2/zlib--1.0.sql)
- [官方 C 实现](https://github.com/overplumbum/postgresql-zlib/blob/74f0e5126266105fea05dcd248400e73bc7bacc2/zlib.c)
- [官方 control 文件](https://github.com/overplumbum/postgresql-zlib/blob/74f0e5126266105fea05dcd248400e73bc7bacc2/zlib.control)
- [官方代码仓库](https://github.com/overplumbum/postgresql-zlib)

`zlib` 版本 `1.0` 提供一个函数，把采用 zlib 包装的 `bytea` 解压为 PostgreSQL `text`。它适合读取应用已经用 zlib 压缩的可信载荷；不负责压缩，也不处理 gzip/归档格式。

### 核心流程

以 `bytea` 传入完整 zlib 流：

```sql
CREATE EXTENSION zlib;

SELECT zlib_decompress(
  decode('789ccb48cdc9c90700062c0215', 'hex')
);
```

示例返回 `hello`。在应用表中，应把压缩字段保留为 `bytea`，只在确实需要文本输出时调用函数。

### API

`zlib_decompress(bytea) RETURNS text` 调用 zlib 的 `uncompress`。它声明为 `IMMUTABLE STRICT`：空输入返回空值，无效流抛出包含 zlib 诊断的外部例程错误，成功输出作为数据库文本返回。

### 注意事项

解压大小没有上限。实现最初分配压缩大小六倍的输出缓冲，并在 `Z_BUF_ERROR` 时反复翻倍；高压缩率或恶意载荷可能分配大量后端内存并造成拒绝服务。应只对可信数据开放函数，或在外部强制限制压缩/输出大小。

解压字节必须符合数据库文本编码。函数接受 zlib 流，不接受 gzip 头、ZIP 文件或任意 deflate 片段。源码发布于 2013 年，没有现代 PostgreSQL 兼容矩阵，因此部署前应测试确切的软件包/服务器构建与畸形输入。
