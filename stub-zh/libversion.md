## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/repology/postgresql-libversion/blob/215a2d54573deb560b5165eb681a766085603829/README.md)
- [扩展 control 文件](https://github.com/repology/postgresql-libversion/blob/215a2d54573deb560b5165eb681a766085603829/libversion.control)
- [C 实现](https://github.com/repology/postgresql-libversion/blob/215a2d54573deb560b5165eb681a766085603829/libversion.c)
- [libversion 依赖](https://github.com/repology/libversion)

`libversion` 通过外部 libversion 库比较软件包风格的版本字符串。它提供比较函数和 `versiontext` 类型；该类型的等值与排序使用版本语义，而不是文本字典序。

```sql
CREATE EXTENSION libversion;
SELECT version_compare2('1.10', '1.2');
SELECT '1.10'::versiontext > '1.2'::versiontext;
SELECT '1.0'::versiontext = '1.0.0'::versiontext;
```

`version_compare4` 为两个操作数分别接收标志；辅助函数提供 `VERSIONFLAG_P_IS_PATCH`、`VERSIONFLAG_LOWER_BOUND` 和 `VERSIONFLAG_UPPER_BOUND` 等 libversion 标志。

扩展需要匹配的 libversion 开发与运行时库。目录中记录的兼容范围止于 PostgreSQL 16，因此在更新大版本上必须验证源码和二进制构建。`versiontext` 索引依赖该比较行为：主库、备库、转储恢复目标与逻辑订阅端应保持库版本一致，以免排序结果分歧。
