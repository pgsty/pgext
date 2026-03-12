


## 用法

> [omni_seq: 分布式整数序列](https://docs.omnigres.org/omni_seq/id/)

`omni_seq` 扩展提供分布式（带前缀的）标识类型，以解决 PostgreSQL 在逻辑复制中不复制序列数据的限制。

### 类型

类型遵循 `omni_seq.id_<PREFIX_TYPE>_<ID_TYPE>` 的模式。支持的基础类型：`int16`、`int32`、`int64`。当前缀类型和标识类型相同时，可使用简写 `omni_seq.id_int64`。

### 创建使用分布式 ID 的表

```sql
CREATE SEQUENCE seq;
CREATE TABLE t (
    id omni_seq.id_int64 PRIMARY KEY NOT NULL
       DEFAULT omni_seq.id_int64_nextval(omni_seq.system_identifier(), 'seq')
);
```

ID 显示为 `PREFIX:IDENTIFIER` 格式（例如 `7222168279780171472:1`）。

### 函数

- **`omni_seq.id_int64_nextval(node_id, sequence_name)`** -- 生成下一个分布式 ID
- **`omni_seq.system_identifier()`** -- 从 pg_control 返回唯一的 64 位系统标识符
- **`omni_seq.id_int64_int32_make(prefix, identifier)`** -- 从组件构建分布式 ID（迁移时有用）

### 迁移示例

```sql
BEGIN;
LOCK TABLE my_table;
ALTER TABLE my_table ALTER COLUMN id DROP IDENTITY IF EXISTS;
CREATE SEQUENCE my_table_id_seq;
ALTER TABLE my_table
    ALTER COLUMN id TYPE omni_seq.id_int64_int32
        USING omni_seq.id_int64_int32_make(0, id),
    ALTER COLUMN id SET DEFAULT omni_seq.id_int64_int32_nextval(
        omni_seq.system_identifier(), 'my_table_id_seq');
COMMIT;
```
