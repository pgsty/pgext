## 用法

来源：

- [已复核 commit 的 pgfga README](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/README.md)
- [已复核 commit 的 pgfga.control](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/pgfga.control)
- [Cargo 软件包元数据](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/Cargo.toml)
- [SQL API 与表定义](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/src/lib.rs)
- [授权求值器](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/src/check.rs)
- [元组存储实现](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/src/storage.rs)

`pgfga` 是一个受 Zanzibar 启发的实验性关系访问控制引擎。它存储 JSON 授权模式和关系元组，再通过 `pgfga.check` 求值直接关系与权限重写。管理函数包括 `pgfga.create_schema`、`pgfga.read_schemas`、`pgfga.create_tuple`、`pgfga.read_tuples` 和 `pgfga.delete_tuple`。

### 最小关系检查

下面的示例使用 psql 的 `\gset` 保存生成的模式 UUID：

```sql
CREATE EXTENSION pgfga;

SELECT pgfga.create_schema(
  '{"namespaces":{"document":{"relations":{"viewer":[{"namespace":"user"}]}},"user":{}}}'::json
) AS schema_id \gset

SELECT pgfga.create_tuple(
  :'schema_id'::uuid, 'document', 'doc-1', 'viewer', 'user', 'anya', ''
);
SELECT pgfga.check(
  :'schema_id'::uuid, 'document', 'doc-1', 'viewer', 'user', 'anya', ''
);
```

### 注意事项

- 仓库已归档，上游也把版本 `0.1.0` 称为实验性的开发中项目。不要把它用作生产授权边界。
- 元组写入只检查模式 UUID 是否存在，不会根据 JSON 模式验证命名空间、关系或主体限制。元组表也没有模式外键和二级查找索引。
- 已复核的交集与排除求值器会压制部分嵌套检查错误，并可能继续得到 true 结果。对于授权决策，这是不安全的 fail-open 行为。
- 读取路径会把匹配元组物化为 Rust vector，递归检查还会反复发出 SPI 查询。只有修正安全语义后，才应测试深度和扇出的性能。
