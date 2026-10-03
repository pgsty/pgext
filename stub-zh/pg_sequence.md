## 用法

来源：

- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_sequence/pg_sequence.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_sequence/pg_sequence.control)
- [extensions/pg_sequence/src/definition.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_sequence/src/definition.rs)
- [extensions/pg_sequence/src/generate.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_sequence/src/generate.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)

`pg_sequence` 0.3.0 在 `pgsequence` 中保存文档编号规则，支持前后缀、格式、作用域及按年／月重置。

### 核心用法

```sql
CREATE EXTENSION pg_sequence;
SELECT pgsequence.create_sequence('invoice');
SELECT pgsequence.next_val('invoice', 'customer-a');
SELECT pgsequence.next_formatted('invoice', 'customer-a');
```

### 运行边界

控制文件不可重定位，不限定超级用户安装，无需预加载。`create_sequence`、`alter_sequence` 和 `drop_sequence` 管理定义，`next_val`、`next_formatted`、`current_val` 与 `preview` 提供计数接口。返回的 `seq_id` 类型支持比较和索引。应限制重置与规则修改权限，并使作用域键符合业务唯一性要求。这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
