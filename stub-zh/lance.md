## 用法

来源：

- [所审计修订版的官方 README](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/README.md)
- [扩展控制文件](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/lance.control)
- [SQL 可见的 Rust 入口](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/src/lib.rs)
- [Arrow/Lance 类型映射](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/src/fdw/type_mapping.rs)

`lance` 是 `pglance` Rust crate 生成的 PostgreSQL 扩展。版本 `0.0.0` 通过 `lance_fdw` 将本地或远程 Lance 数据集暴露为只读外部表，并提供模式发现和命名空间协调辅助函数。

### 核心流程

创建扩展和 FDW 服务器，然后按 URI 导入数据集：

```sql
CREATE EXTENSION lance;
CREATE SERVER lance_srv FOREIGN DATA WRAPPER lance_fdw;

SELECT lance_import(
  'lance_srv',
  'public',
  'events_lance',
  '/srv/lance/events.lance',
  batch_size => 2048
);

SELECT * FROM public.events_lance LIMIT 10;
```

`lance_import` 检查 Lance 模式、创建外部表，并在嵌套 `struct` 字段需要时创建 PostgreSQL 复合类型。

对于 Lance 命名空间，应先查看计划再应用 DDL：

```sql
SELECT *
FROM lance_attach_namespace(
  'lance_srv',
  schema_prefix => 'lance',
  dry_run => true
);

SELECT *
FROM lance_sync_namespace(
  'lance_srv',
  schema_prefix => 'lance',
  drop_missing => false,
  recreate_changed => false,
  dry_run => true
);
```

### 对象与选项

- `lance_fdw` 实现外部扫描；所审计处理器没有插入、更新或删除回调。
- `lance_import(server_name, local_schema, table_name, uri, batch_size)` 根据数据集 URI 创建一个外部表。
- `lance_attach_namespace(...)` 为命名空间树创建本地模式、外部表和嵌套复合类型。
- `lance_sync_namespace(...)` 报告或协调漂移；`drop_missing` 和 `recreate_changed` 会启用破坏性 DDL。
- `uri` 可设置在服务器或外部表上。`batch_size` 默认为 1024，表选项优先于服务器选项。
- 命名空间表用 `ns.table_id` 保存字符串 JSON 数组；服务器上的 `ns.*` 选项配置命名空间实现。

原生映射包括：Boolean 到 `boolean`，有符号/无符号整数到最接近的 PostgreSQL 整数类型，字符串到 `text`，二进制值到 `bytea`，日期与时间戳到 PostgreSQL 时间类型，列表到数组，结构体到自动生成的复合类型，映射或不支持的类型到 `jsonb`。

### 运维边界

控制文件将 `lance` 标记为仅超级用户可安装的不受信任扩展。所审计 crate 声明了 PostgreSQL 13–17 构建特性，但版本 `0.0.0` 仍是早期接口，必须针对实际服务器和 Lance 文件测试。

外部扫描使用数据库服务器的文件系统、网络和凭据上下文读取文件或远程对象存储，因此应限制 URI 和存储选项。命名空间附加或同步前始终检查 `dry_run` 输出；启用 `drop_missing` 或 `recreate_changed` 可能删除或替换本地外部表及其依赖对象。
