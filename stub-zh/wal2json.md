## 用法

来源：

- [wal2json 2.7 发布](https://github.com/eulerto/wal2json/releases/tag/wal2json_2_7)
- [官方 README](https://github.com/eulerto/wal2json)

`wal2json` 2.7 是一个逻辑解码输出插件，从 PostgreSQL WAL 生成 JSON 格式的变更数据捕获。它支持当前 PostgreSQL 18 和 19 开发分支，并作为输出插件加载，因此不使用 `CREATE EXTENSION` 或 `shared_preload_libraries`。

对于受 CVE-2026-6471 影响的 PostgreSQL 18.6、17.11、16.15、15.19、14.24 及后续小版本，应按上游发布说明将 `wal2json` 加入 `output_plugin_libraries`。

### 配置

在 `postgresql.conf` 中：

```ini
wal_level = logical
max_replication_slots = 10
max_wal_senders = 10
```

### 使用流协议（pg_recvlogical）

```bash
# 创建复制槽
pg_recvlogical -d postgres --slot test_slot --create-slot -P wal2json

# 开始消费变更
pg_recvlogical -d postgres --slot test_slot --start -o pretty-print=1 -f -

# 完成后删除槽
pg_recvlogical -d postgres --slot test_slot --drop-slot
```

### 使用 SQL 函数

```sql
-- 创建逻辑复制槽
SELECT * FROM pg_create_logical_replication_slot('test_slot', 'wal2json');

-- 查看变更（不消费）
SELECT data FROM pg_logical_slot_peek_changes('test_slot', NULL, NULL);

-- 获取并消费变更
SELECT data FROM pg_logical_slot_get_changes('test_slot', NULL, NULL,
    'pretty-print', '1');

-- 删除槽
SELECT pg_drop_replication_slot('test_slot');
```

### 输出格式 v1（每事务 JSON）

```json
{
  "change": [
    {
      "kind": "insert",
      "schema": "public",
      "table": "my_table",
      "columnnames": ["a", "b"],
      "columntypes": ["integer", "text"],
      "columnvalues": [1, "hello"]
    },
    {
      "kind": "delete",
      "schema": "public",
      "table": "my_table",
      "oldkeys": {
        "keynames": ["a"],
        "keytypes": ["integer"],
        "keyvalues": [1]
      }
    }
  ]
}
```

### 输出格式 v2（每元组 JSON）

启用方式：`'format-version', '2'`

### 关键参数

- `include-xids` - 添加事务 ID（默认：false）
- `include-timestamp` - 添加时间戳（默认：false）
- `include-schemas` - 添加模式名（默认：true）
- `include-types` - 添加列类型（默认：true）
- `include-pk` - 添加主键信息（默认：false）
- `include-lsn` - 添加 WAL LSN（默认：false）
- `include-not-null` - 添加 NOT NULL 信息（默认：false）
- `include-default` - 添加默认表达式（默认：false）
- `pretty-print` - 格式化 JSON 输出（默认：false）
- `filter-tables` - 逗号分隔的要包含的表列表
- `add-tables` - 与 filter-tables 相同
- `filter-msg-prefixes` - 按前缀过滤逻辑消息
- `format-version` - 1（每事务）或 2（每元组）
- `actions` - 按操作类型过滤：insert、update、delete、truncate
