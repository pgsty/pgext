## 用法

来源：

- [固定提交的上游 README](https://github.com/supabase-community/pg_headerkit/blob/265d4402221c5d38c1206eb6a7d64ccb02584e21/README.md)
- [固定提交的 1.0 版安装 SQL](https://github.com/supabase-community/pg_headerkit/blob/265d4402221c5d38c1206eb6a7d64ccb02584e21/pg_headerkit--1.0.sql)
- [固定提交的扩展控制文件](https://github.com/supabase-community/pg_headerkit/blob/265d4402221c5d38c1206eb6a7d64ccb02584e21/pg_headerkit.control)

pg_headerkit 1.0 版是纯 SQL 的 PostgREST 请求头辅助扩展。它创建固定的 hdr 模式、允许列表/拒绝列表，以及读取 request.headers 设置、提取代理/客户端字段、检查 IP 列表和做简单移动端用户代理识别的函数。

### 修正后的预期用法

下面展示预期 API，但固定提交中的 1.0 安装脚本必须先修正，这段流程才能成功执行：

```sql
CREATE EXTENSION pg_headerkit;

INSERT INTO hdr.allow_list (ip)
VALUES ('203.0.113.10');

SELECT hdr.header('user-agent');
SELECT hdr.ip();
SELECT hdr.in_allow_list('203.0.113.10'::inet);
```

这些函数只暴露由 PostgREST 或其他可信网关提供的请求上下文。它们本身不执行速率限制、路由、请求过滤或授权；应用必须在策略或函数中调用它们并实现这些决策。

### 当前缺陷与安全

在审阅的提交中，hdr.projectref 与 hdr.ref 的函数体都包含两个相邻且没有分隔符的 SELECT 语句。因此发行的 1.0 SQL 存在语法错误，CREATE EXTENSION pg_headerkit 会整体失败。打包前必须修补并审阅安装 SQL，或等待上游修正。

若干读取请求设置或允许/拒绝表的函数被错误声明为不可变。PostgreSQL 可能预计算或缓存这类结果，因此在将其用于安全决策前，必须改成正确的易变性。IP 辅助函数信任 x-forwarded-for 的第一个值；只有受控反向代理会覆盖该请求头时才可采用。还必须明确设计模式、表和函数权限，因为上游 SQL 没有提供完整的角色模型，函数也使用调用者权限运行。

README 中的速率限制和访问控制是可以基于这些辅助函数构建的用途，并非扩展已经安装的强制机制。应审计修正后的 SQL、撤销不需要的 PUBLIC 执行权限，并通过实际的 PostgREST 角色和连接池模式测试每条策略。
