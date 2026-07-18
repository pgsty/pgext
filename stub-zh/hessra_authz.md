## 用法

来源：

- [已复核 commit 的 hessra_authz README](https://github.com/Hessra-Labs/hessra-sdk.rs/blob/eab4015d98ccc36e4fb3844f4850d3cbea8afdf6/hessra-pgrx/hessra_authz/README.md)
- [扩展 control 文件](https://github.com/Hessra-Labs/hessra-sdk.rs/blob/eab4015d98ccc36e4fb3844f4850d3cbea8afdf6/hessra-pgrx/hessra_authz/hessra_authz.control)
- [Cargo 软件包与 PostgreSQL feature 矩阵](https://github.com/Hessra-Labs/hessra-sdk.rs/blob/eab4015d98ccc36e4fb3844f4850d3cbea8afdf6/hessra-pgrx/hessra_authz/Cargo.toml)
- [当前 SQL API 实现](https://github.com/Hessra-Labs/hessra-sdk.rs/blob/eab4015d98ccc36e4fb3844f4850d3cbea8afdf6/hessra-pgrx/hessra_authz/src/lib.rs)

`hessra_authz` 在 PostgreSQL 内部对 Biscuit 授权令牌执行本地验证。它存储命名公钥和服务链定义，然后检查令牌的主体、资源与操作；验证时无需访问授权服务。

### 密钥管理与验证

创建扩展，注册真实 PEM 公钥，再验证真实 Biscuit 令牌。示例使用 psql 变量，避免把凭据写入 SQL 历史。

```sql
CREATE EXTENSION hessra_authz;

SELECT add_public_key('issuer', :'issuer_public_key_pem', true);

SELECT verify_token_with_stored_key(
  :'biscuit_token',
  'issuer',
  'alice',
  'invoice/42',
  'read'
);
```

公钥管理函数包括 `add_public_key`、`get_public_key`、`update_public_key` 和 `delete_public_key`。服务链管理使用 `add_service_chain`、`get_service_chain`、`update_service_chain` 和 `delete_service_chain`。直接验证函数接收 PEM 密钥，而便捷函数会从扩展管理的表中读取命名密钥或默认密钥。

### 注意事项

- 当前源码为 `verify_token` 和 `verify_token_with_stored_key` 定义了五个参数；已复核 README 的示例包含一个已经过时的第六个 `NULL` 参数。应以已安装构建生成的源码/API 为准。
- 验证成功时返回 SQL `void`，拒绝或输入格式错误时抛出错误；它并不返回布尔值。因此，RLS 策略中的 `IS NULL` 测试会中止被拒绝的语句，而不是求值为 false。
- control 文件把版本 `0.3.2` 标记为仅超级用户可安装、不受信任且不可重定位。Cargo 构建 feature 覆盖 PostgreSQL 13 至 17，而 README 声称支持 PostgreSQL 12 及以上版本。
- 密钥与服务链管理函数会修改共享扩展表，而上游没有定义应用角色策略。请检查所有权，并从不得修改授权配置的角色撤销 `EXECUTE` 或表权限。
