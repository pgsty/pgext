## 用法

来源：

- [官方 README](https://github.com/pandrewhk/pgcurl/blob/97d1e7e7447adae345fd90695916031237e1b06b/README.md)
- [扩展 SQL](https://github.com/pandrewhk/pgcurl/blob/97d1e7e7447adae345fd90695916031237e1b06b/pgcurl--0.0.1.sql)
- [libcurl 实现](https://github.com/pandrewhk/pgcurl/blob/97d1e7e7447adae345fd90695916031237e1b06b/pgcurl.c)

`pgcurl` 0.0.1 只暴露一个 `curl(text) RETURNS text` 函数，它会在 PostgreSQL 后端中同步执行 libcurl GET。它可以在受控实验中读取小型文本资源，但这个已废弃实现缺少生产 HTTP 集成所需的控制能力。

### 基础调用

```sql
CREATE EXTENSION pgcurl;

SELECT curl('https://example.com/');
```

函数会跟随重定向，并要求 libcurl 解码 gzip 或 deflate 等受支持内容编码。它只把响应体作为 PostgreSQL 文本返回。

### 失败语义

传输错误会追加到返回文本中，而不是抛出 SQL 错误。函数不返回 HTTP 状态、响应头和最终 URL，也不区分 HTTP 错误状态与成功正文。因此，同一个返回值可能先含部分响应，再含错误消息。

SQL 接口没有方法、响应头、凭据、TLS 策略、代理、重定向上限、响应大小上限、连接超时或总超时选项。DNS、连接、重定向和响应传输期间，调用会一直阻塞数据库后端。绝不能传入空值：SQL 函数没有声明为严格函数，而 C 代码不处理空参数。

### 安全边界

任何可执行 `curl(text)` 的角色都能让数据库服务器请求攻击者选择的 URL，包括回环地址、私网、云元数据或重定向目标。大响应还会消耗后端内存。应撤销不可信角色的执行权，并在扩展外实施目标、DNS、出口、超时和响应大小控制。

项目已废弃，也没有维护中的 PostgreSQL 或 libcurl 兼容矩阵。生产环境应在应用或任务进程中执行 HTTP 工作，显式处理认证、可观测性、重试、熔断和独立于数据库事务的失败语义。
