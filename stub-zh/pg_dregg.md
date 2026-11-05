## 用法

来源：

- [上游来源 1](https://github.com/emberian/dregg/blob/4caff0718e0a34b9a76281008ebf908e965f856f/pg-dregg/README.md)
- [上游来源 2](https://github.com/emberian/dregg/blob/4caff0718e0a34b9a76281008ebf908e965f856f/pg-dregg/pg_dregg.control)

`pg_dregg` 0.1.0 是用于对象能力授权和行级安全的实验性 PostgreSQL 适配器。

### 核心流程

DBA 配置签发者信任密钥，应用通过 `dregg.token` 提交已签发的能力令牌。RLS 条件调用 `dregg_admits`，检查操作与行标识符。`dregg_cap_explain` 提供拒绝原因，`dregg_revoke` 支持凭据撤销。应在受保护的表上启用 RLS，并按需要强制执行。

```sql
CREATE EXTENSION pg_dregg;
SELECT dregg_issuer_status();
```

### 运行边界

文档中的实时适配器面向 PostgreSQL 18，安装需要超级用户。信任配置缺失或格式错误时会拒绝授权。令牌签发与权限收窄不在此最小 SQL 示例中，生产环境不能使用示意密钥或令牌。缺少验证器时，可选证明认证会封闭失败。上游的形式化验证声明不能替代对实际部署的独立安全审计。
