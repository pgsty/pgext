## 用法

来源：

- [README](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/README.md)
- [Control](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/pg_crdt.control)
- [Install SQL](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/pg_crdt--0.0.1.sql)
- [Installation and backup](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/docs/install.md)
- [Untrusted writers](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/docs/guides/security.md)
- [Collaborative text](https://github.com/kiwicopple/pg-extensions/blob/16ab54441c60a538bb9a73ffb6bf8c901480e743/pg_crdt/docs/types/collaborative-text.md)

`pg_crdt` 0.0.1 用 SQL/PL/pgSQL 提供可合并计数器和协作文本。支持 PostgreSQL 16–18，不依赖其他扩展或共享预加载，安装到固定的 `crdt` 模式。实验性的文本列表示在每次编辑时都会重写整个文档。

### 安装

服务器安装控制文件和版本化 SQL 后，拥有数据库 `CREATE` 权限的角色即可创建扩展，无需超级用户。不要指定其他模式。安装会拒绝使用其他角色可以写入的已有模式。官方指南也提供 pg_tle 注册和直接执行 SQL 的方式；这些是替代安装路径，不是额外依赖。

```sql
CREATE EXTENSION pg_crdt;
```

### 计数器流程

`crdt.pn` 保存各副本的正负累计值。`crdt.pn_increment`、`crdt.pn_decrement` 与 `crdt.pn_value` 用于修改或读取计数器；`crdt.pn_merge` 以幂等方式合并状态。应在一条 SQL 中更新存储值，让行锁保护并发修改；不要先读取，再在客户端合并后覆盖。

```sql
CREATE TABLE posts (
  id bigint PRIMARY KEY,
  likes crdt.pn NOT NULL DEFAULT crdt.pn_empty()
);
INSERT INTO posts (id) VALUES (1);
UPDATE posts SET likes = crdt.pn_increment(likes, 'alice', 3) WHERE id = 1;
UPDATE posts SET likes = crdt.pn_decrement(likes, 'alice', 1) WHERE id = 1;
SELECT crdt.pn_value(likes) FROM posts WHERE id = 1;
```

### 协作文本

`crdt.yseq` 将文档存于基于 JSONB 的列中；`crdt.yseq_insert`、`crdt.yseq_merge` 和 `crdt.yseq_value` 分别用于编辑、合并和读取。`yseq_` 函数中传输已删除字符的操作要求 UTF8。列状态保留删除标记，并随编辑历史增长。`ydoc_` 接口则把文档保存在扩展表中，应用包装函数需要相应的表和序列授权；这些表上的编辑不会触发引用它们的应用行的策略或变更流。

### 信任边界与备份

不要允许不可信客户端直接修改 CRDT 原始状态，否则它们可伪造其他副本的累计值或编辑。应围绕 `crdt.pn_apply_delta` 与文本 apply 接口创建受限应用包装函数，绑定可信副本身份、检查行访问权，并固定包装函数的搜索路径。客户端可控制的设置本身不能认证身份。扩展公开授予函数执行权限，但不授予文档表访问权限。扩展配置转储会保存文档表及序列位置，备份应包含 `crdt` 模式。删除扩展会删除其文档表，使用级联删除还会删除使用扩展类型的应用列。卸载前应先导出数据，并逐项处理依赖。
