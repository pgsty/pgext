## 用法

来源：

- [README.md](https://github.com/realZachi/pg-jev/blob/8d9598d87d5ff460998d91ec070226176024a841/README.md)
- [CHANGELOG.md](https://github.com/realZachi/pg-jev/blob/8d9598d87d5ff460998d91ec070226176024a841/CHANGELOG.md)
- [sql/jev--0.2.0--0.2.1.sql](https://github.com/realZachi/pg-jev/blob/8d9598d87d5ff460998d91ec070226176024a841/sql/jev--0.2.0--0.2.1.sql)
- [PGXN 0.2.1](https://pgxn.org/dist/jev/0.2.1/)

`jev` 使用配置的 Jev 兼容 API 对自然语言条件求值，实现行过滤、排序和分类。它依赖 `plpython3u`，需要超级用户安装；TypeSafe 托管端点还需要 API 密钥。上游文档支持 PostgreSQL 14–17，无需预加载。

### 查询数据行

```sql
CREATE EXTENSION jev CASCADE;
SET jev.api_key = 'your-key';
CREATE TABLE jev_demo (id integer, body text);
INSERT INTO jev_demo VALUES (1, 'The customer requests a refund');
SELECT id, jev_prob(jev_demo, 'the customer requests a refund')
FROM jev_demo;
SELECT jev_stats();
```

`jev()` 返回布尔谓词，`jev_prob()` 返回概率；`jev_choice()` 从选项中分类，`jev_score()` 对有序等级评分。缓存和连接池属于会话，`jev_cache_clear()` 清理会话状态。

### 服务与数据边界

扩展会将行内容发送至配置的 API。通过 `jev.api_url` 指定服务，使用 `jev.max_rows_per_statement` 和 `jev.max_chars_per_statement` 限制工作量。API 延迟和费用取决于服务与数据量，API 密钥须按凭据管理。

PL/Python 以数据库服务进程的操作系统权限运行；不提供超级用户或 PL/Python 的托管服务无法运行此扩展。本地包测试使用上游模拟 API，不验证远程模型判断的质量。

### 0.2.1 端点配置

0.2.1 仅对 `*.typesafe.ai` 主机要求 `jev.api_key`。本地运行或其他兼容端点在未设置密钥时不发送 Authorization 头，若已配置密钥仍会发送。将 `jev.api_url` 指向预期的可信服务。为 0.2.0 数据库安装匹配文件后执行 `ALTER EXTENSION jev UPDATE`。模型准确性与数据处理策略仍取决于所选服务。
