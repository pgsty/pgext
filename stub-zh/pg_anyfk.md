## 用法

来源：

- [README 0.2.0](https://github.com/Vinnu124/pg_anyFK/blob/v0.2.0/README.md)
- [Control](https://github.com/Vinnu124/pg_anyFK/blob/v0.2.0/ext/pg_anyfk.control)
- [Cargo](https://github.com/Vinnu124/pg_anyFK/blob/v0.2.0/ext/Cargo.toml)
- [Release with generated SQL](https://github.com/Vinnu124/pg_anyFK/releases/tag/v0.2.0)

`pg_anyfk` 0.2.0 使用触发器约束数组中每个元素对父表键的引用，主要面向已经使用数组存储关系的模式；关联表仍是常规选择。

### 基本流程

上游支持 PostgreSQL 15–18。该扩展不受信任且不可重定位，安装需要超级用户，无需预加载或重启。

```sql
CREATE EXTENSION pg_anyfk;
CREATE TABLE tags (id bigint PRIMARY KEY);
CREATE TABLE posts (
    id bigserial PRIMARY KEY,
    tag_ids bigint[] NOT NULL DEFAULT '{}'
);
CREATE INDEX ON posts USING gin (tag_ids);
SELECT anyfk.add('posts(tag_ids)', 'tags(id)');
SELECT * FROM anyfk.constraints;
```

调用者必须拥有两张表。数组元素类型须与父列类型一致，父列须有单列唯一索引。GIN 索引可以避免每次删除父行时扫描全部子行。

### 校验与删除

`anyfk.add` 默认检查已有数据。启用前可以先查询违规引用：

```sql
SELECT * FROM anyfk.violations('posts(tag_ids)', 'tags(id)');
```

选项包括 `validate`、`is_deferrable`、`conname` 和 `on_delete`。默认的 `restrict` 拒绝删除仍被引用的父行。若希望删除时移除子数组中的对应元素，可选择 `scrub`；执行删除的角色还需要对子表具有 UPDATE 权限。`scrub` 不会允许修改仍被引用的父键。

`is_deferrable => true` 将校验延迟至事务完成。移除已登记约束时，从 `anyfk.constraints` 取得名称，再调用 `anyfk.drop_constraint`。

### 边界

允许空数组值、空数组、空元素和重复元素。存在非空引用时拒绝截断父表，但支持同时截断父表与子表。检查期间使用键级行锁防止并发父表变更破坏引用。

这些约束通过触发器实现，不是 `pg_constraint` 中的原生外键；禁用触发器也会禁用约束。扩展跟踪表标识，并使用事件触发器处理重命名和依赖对象清理。
