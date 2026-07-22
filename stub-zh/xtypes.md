## 用法

来源：

- [官方 README](https://github.com/mkindahl/pg_xtypes/blob/eee55dfe3fb8d02baa36926692d1ccbcee930df6/README.md)
- [官方扩展 SQL](https://github.com/mkindahl/pg_xtypes/blob/eee55dfe3fb8d02baa36926692d1ccbcee930df6/xtypes--0.1.sql)
- [官方 bytes 实现](https://github.com/mkindahl/pg_xtypes/blob/eee55dfe3fb8d02baa36926692d1ccbcee930df6/bytes.c)

`xtypes` 版本 `0.1` 当前提供一个 `bytes` 类型，以无符号 64 位值保存字节数量，按 `pg_size_pretty` 风格格式化，并执行有限的算术运算。当无符号范围与较小的操作符集合可以满足需求时，它适合保存人类可读的容量字段。

### 核心流程

创建扩展，使用显式单位保存容量值，并使用提供的算术运算：

```sql
CREATE EXTENSION xtypes;

CREATE TABLE storage_limits (
  service text PRIMARY KEY,
  quota bytes NOT NULL
);

INSERT INTO storage_limits VALUES ('archive', '1.5 TiB');

SELECT service, quota, quota + '512 GiB'::bytes
FROM storage_limits;
```

输出会自动选择二进制单位，例如 `1.5 TB` 或 `2 TB`。

### 类型与操作符

- `bytes` 使用无符号 64 位表示保存基础字节数。
- 输入接受 `B`、`bytes`、`kB`/`KB`/`KiB`，以及从 `MB`/`MiB` 到 `EB`/`EiB` 的对应后缀。即使短标签没有 `i`，也都按 1024 的幂缩放。
- 操作符包括 `+`、`-`、`<`、`>`、两个方向的 `int8` 乘法，以及除以 `int8`。
- `round(bytes)` 按自动选择的显示单位取整，并返回另一个 `bytes` 值。

### 运维注意事项

最大可表示值略低于 16 EiB。加法、减法与乘法会检查无符号溢出或下溢，但复核的输入路径和有符号 `int8` 算术路径仍需要针对负值与边界情况做显式测试。小数输入会转换为整数个字节，格式化输出使用 `KB` 等按二进制缩放的短标签，容易与十进制 SI 单位混淆。SQL 没有定义相等操作符、类型转换或 B-tree 操作符类，因此不能假设明确提供的操作符之外还具备普通唯一索引或排序行为。应将二进制固定到目标 PostgreSQL 大版本，并验证转储恢复与客户端驱动的文本处理。
