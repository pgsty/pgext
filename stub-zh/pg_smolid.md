## 用法

来源：

- [Official extension control file](https://github.com/dotvezz/pg_smolid/blob/b9840b6764545c4e59353506cfa4bcf2cfb2bf4f/pg_smolid.control)
- [Official Rust package manifest](https://github.com/dotvezz/pg_smolid/blob/b9840b6764545c4e59353506cfa4bcf2cfb2bf4f/Cargo.toml)
- [Official Rust implementation](https://github.com/dotvezz/pg_smolid/blob/b9840b6764545c4e59353506cfa4bcf2cfb2bf4f/src/lib.rs)

`pg_smolid` 0.1.0 增加由 Rust `smolid` crate 支持的 `smolid` 数据类型与生成器。其值具有紧凑文本表示，按底层无符号 64 位整数比较，并支持 B-tree 与哈希索引。

### 核心流程

```sql
CREATE EXTENSION pg_smolid;

CREATE TABLE events (
  id smolid PRIMARY KEY DEFAULT smolid(),
  payload jsonb NOT NULL
);

INSERT INTO events(payload) VALUES ('{"kind":"signup"}');
SELECT id, smolid_to_text(id), smolid_version(id) FROM events;
```

用 `smolid()` 生成新值，或用 `smolid_with_type(integer)` 请求编码类型值。`smolid_type(smolid)` 返回这个可选类型；`smolid_version(smolid)` 返回格式版本。`smolid_to_bigint(smolid)` 和 `smolid_to_text(smolid)` 暴露数值与文本形式。

### 类型行为与注意事项

扩展定义比较操作符、默认 B-tree 和哈希操作符类，以及从 `smolid` 到 `bigint` 和 `text` 的隐式转换。排序依据编码整数，而不是本地化文本顺序。控制文件把安装标记为仅超级用户且不可信。应在各节点固定扩展和 Rust crate 版本，并在把该类型用于持久键或逻辑复制前验证序列化与排序行为。
