## 用法

来源：

- [packages/errors/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/errors/README.md)
- [packages/errors/sql/errors--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/errors/sql/errors--0.47.0.sql)
- [packages/errors/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/errors/Makefile)
- [packages/errors/errors.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/errors/errors.control)

`errors` 提供 `errors.raise_error` 辅助函数，以 SQLSTATE P0001 抛出异常，MESSAGE 保存错误码，DETAIL 是包含 code、context 和 class 的 JSON 对象。

### 核心用法

```sql
CREATE EXTENSION "errors" CASCADE;
SELECT errors.raise_error('DEMO_ERROR', '{"field":"name"}'::jsonb, 'public');
```

### 运行边界

示例会主动抛出异常。class 是应用元数据，不是访问控制机制；调用者和日志仍可能获得传入的上下文，不应在错误载荷中放入凭据。

0.47.0 是 SQL/PLpgSQL 扩展，没有自己的共享库，也不要求预加载。须先安装配套版本的依赖：`plpgsql`。 控制文件允许非超级用户安装，但没有标记为 trusted；依赖、模式创建和角色授权权限仍须满足。上游未声明当前 PostgreSQL 主版本矩阵。
