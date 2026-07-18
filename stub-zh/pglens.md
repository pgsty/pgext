## 用法

来源：

- [上游 README 与当前状态说明](https://github.com/jamie-steele/pg-lens/blob/c824858bcd1d888c53dd5a246f019aec238b7ebe/README.md)
- [扩展 control 文件](https://github.com/jamie-steele/pg-lens/blob/c824858bcd1d888c53dd5a246f019aec238b7ebe/pglens.control)
- [SQL API](https://github.com/jamie-steele/pg-lens/blob/c824858bcd1d888c53dd5a246f019aec238b7ebe/sql/pglens--0.1.0.sql)
- [重计算实现](https://github.com/jamie-steele/pg-lens/blob/c824858bcd1d888c53dd5a246f019aec238b7ebe/src/recompute.c)
- [路线图](https://github.com/jamie-steele/pg-lens/blob/c824858bcd1d888c53dd5a246f019aec238b7ebe/docs/ROADMAP.md)

`pglens` 是实验性的投影/读模型扩展。一个 lens 会把普通 SQL 视图、物理目标表和稳定键关联起来。注册时会按视图当前结构创建带主键的空表；显式重计算则按一个键查询视图，并插入、更新或删除对应目标行。

### 注册并重计算 lens

```sql
CREATE EXTENSION pglens;

CREATE TABLE accounts (
  id bigint PRIMARY KEY,
  email text NOT NULL,
  status text NOT NULL
);

CREATE VIEW account_projection_view AS
SELECT id, email, status FROM accounts;

SELECT pglens.register_lens(
  'account_projection',
  'account_projection_view'::regclass,
  'public',
  'account_projection',
  'id'
);

INSERT INTO accounts VALUES (1, 'a@example.com', 'active');
SELECT pglens.recompute_lens('account_projection', '1');

TABLE account_projection;
```

尽管目标架构是异步处理，版本 `0.1.0` 尚未实现 WAL 解码、已提交变更捕获、受影响键发现、后台工作进程、检查点、重试或重启恢复。当前必须手工重计算。

它只支持普通视图；键最多只能标识一条投影行；目标表结构在注册时固定；后续视图变更不会自动传播；异常路径测试也很有限。上游明确将项目标为实验性且不适合生产。在一致性、崩溃恢复、权限边界、schema 迁移和运维可观测性得到实现与验证前，应把它视为设计原型。
