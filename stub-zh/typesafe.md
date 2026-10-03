## 用法

来源：

- [README](https://github.com/giuliosmall/pg_typesafe/blob/v0.0.1/README.md)
- [Control file / 控制文件](https://github.com/giuliosmall/pg_typesafe/blob/v0.0.1/typesafe.control)
- [typesafe--0.0.1.sql](https://github.com/giuliosmall/pg_typesafe/blob/v0.0.1/typesafe--0.0.1.sql)

`typesafe` 在 SQL 中调用 TypeSafe AI，完成分类、概率式回答和评分。该客户端仍处于早期预览阶段，会把传入文本发送到外部 API，需要 API 凭据和 libcurl 7.61 或更新版本。

### 核心工作流

安装共享库并创建扩展。可在启动 PostgreSQL 前通过服务端环境变量 `TYPESAFE_API_KEY` 提供密钥，或通过 `TYPESAFE_API_KEY_FILE` 指定密钥文件，也可通过仅超级用户可设定的 `typesafe.api_key_file` 指定服务端可读的密钥文件。

```sql
CREATE EXTENSION typesafe;
SELECT typesafe_noul('My payouts have been failing for three days.',
                     'Does this convey urgency?');
SELECT * FROM typesafe_detect_many(
  ARRAY['Service is unavailable', 'Thank you for your help'], 'Is this urgent?');
```

### 接口与权限

`typesafe_noul` 返回概率式结果。检测、分类、评分和问答辅助函数有标量及批量形式，批量调用可减少重复 HTTP 请求。默认已撤销 `PUBLIC` 的 `EXECUTE` 权限，应仅为应用授予实际需要的函数签名。

v0.0.1 发布说明现声明支持 PostgreSQL 15–18。扩展可重定位，无需预加载。会话凭据配置 `typesafe.api_key` 可能出现在 SQL 日志中，应优先使用环境变量或受保护文件。网络延迟、服务限额和 API 失败都会影响查询，远程模型输出不能作为数据库完整性保证。
