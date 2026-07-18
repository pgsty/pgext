## 用法

来源：

- [固定提交的上游 README](https://github.com/fabriziomello/session_variables/blob/1022a217962bea72e9313478d1e0c27c1b07a3a4/README.md)
- [固定提交的安装 SQL](https://github.com/fabriziomello/session_variables/blob/1022a217962bea72e9313478d1e0c27c1b07a3a4/sql/session_variables.sql)
- [固定提交的发行元数据](https://github.com/fabriziomello/session_variables/blob/1022a217962bea72e9313478d1e0c27c1b07a3a4/META.json)
- [正式 PGXN 文档](https://pgxn.org/dist/session_variables/0.0.4/doc/session_variables.html)

session_variables 0.0.4 版是在 PostgreSQL 自定义设置之上实现的两个 SQL 封装函数。set_value 在 session_variables 命名空间下写入文本值，get_value 将它读回。

### 示例

```sql
CREATE EXTENSION session_variables;

SELECT set_value('request_id', 'req-123');
SELECT get_value('request_id');

BEGIN;
SELECT set_value('request_id', 'req-456');
ROLLBACK;

SELECT get_value('request_id');
```

最后一次读取仍返回 req-456：set_value 调用 set_config 时关闭了事务局部模式，所以设置会持续整个 PostgreSQL 会话，不会随事务回滚。值只有 text 类型。读取从未设置过的名称会报错，因为 get_value 没有启用 missing-ok。

### 连接池与安全限制

状态属于后端连接，而不是应用用户或逻辑请求。使用会话池或事务池时，如果每次取用连接时没有重置，后一个客户端可能继承前一个客户端留下的值。扩展没有列举、清除、类型、过期、所有权、权限或隔离模型；若不修改权限，其函数可由 PUBLIC 执行。

不要在 session_variables 中存放认证决策、租户身份、密钥或授权上下文。应优先使用事务局部的 set_config 并严格重置，或采用针对连接池模式设计的原生应用上下文。如果遗留代码依赖它，应撤销 PUBLIC 的执行权限、只开放经过验证的封装函数、每次取用连接时初始化全部变量，并在复用后端前清理状态。

仓库最后更新于 2012 年，文档描述的是很老的 PostgreSQL 安装路径，并被归类为已废弃。SQL 虽然简单，兼容性与运行语义仍需针对目标连接池和 PostgreSQL 版本测试。
