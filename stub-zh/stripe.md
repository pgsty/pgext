## 用法

来源：

- [v0.2.1 官方 README](https://github.com/stainless-commons/stripe-sql/blob/v0.2.1/README.md)
- [v0.2.1 官方 control 文件](https://github.com/stainless-commons/stripe-sql/blob/v0.2.1/stripe.control)
- [v0.2.1 官方变更日志](https://github.com/stainless-commons/stripe-sql/blob/v0.2.1/CHANGELOG.md)
- [v0.2.1 官方 coupon SQL](https://github.com/stainless-commons/stripe-sql/blob/v0.2.1/sql/stripe_coupons.sql)

`stripe` 是实验性的生成式 PL/Python 扩展，它把 Stripe REST API 资源暴露为 PostgreSQL 函数与复合类型。每次 API 调用都会离开数据库进程，并受 Stripe 凭据、网络行为、速率限制与外部副作用约束。

本次核验的上游发行版是 `v0.2.1`，但其 control 文件和生成的扩展 SQL 仍声明扩展版本 `0.0.1`。除了 SQL 扩展版本，还必须固定发行 commit；不能只根据 `extversion` 推断生成 API 接口。

### 前置条件与只读查询

上游要求 PostgreSQL 14+、Python 3.9+、`plpython3u`，并在 PostgreSQL 使用的解释器中安装相匹配的 `stainless_commons_stripe` Python 包。

```sql
CREATE EXTENSION IF NOT EXISTS plpython3u;
CREATE EXTENSION stripe;

SET stripe.secret_key = '<session-injected-secret>';

SELECT *
FROM stripe_coupons.list()
LIMIT 20;
```

API 按 `stripe_accounts`、`stripe_coupons`、`stripe_customers`、`stripe_invoices` 与 `stripe_payment_intents` 等 schema 分组。资源函数执行获取、列举、创建或其他端点操作；`make_*` 辅助函数用于构造嵌套复合参数。

### 分页与缓存

列表函数会自动获取后续页面。应像示例一样把 `LIMIT` 直接放在集合返回函数上层，使 PostgreSQL 能提前停止求值；隐藏在视图或其他查询层后的 limit 可能无法下推，导致请求并缓冲所有页面。

重复分析可物化只读拉取结果，并有计划地刷新。`pg_cron` 可调度刷新，但它是可选项，不会由 `stripe` 安装。

### 凭据与网络边界

客户端读取 `stripe.secret_key` 与 `stripe.base_url`。通过数据库/角色设置持久保存密钥，可能使其暴露给特权目录读取者和备份；应优先受控的逐会话注入、限制连接日志并轮换凭据。修改 `stripe.base_url` 也会形成外连目标/SSRF 边界，因此要限制设置权限，并在 PostgreSQL 外实施出站策略。

`plpython3u` 不受信任，其安装属于超级用户边界。应撤销网络与变更函数的宽泛 `EXECUTE` 权限，再只授予每个应用角色确切需要的读写端点。

### 事务与版本风险

Stripe 变更属于外部副作用：回滚 PostgreSQL 事务不会撤销已成功的远程创建、更新、退款或删除。应围绕这一分裂事务边界设计幂等键、重试规则、对账与 webhook 处理；不能假定普通 SQL 原子性覆盖 API。

项目明确处于实验阶段；`v0.2.0` 曾包含破坏性的生成名称变更，而 SQL 扩展版本仍是 `0.0.1`。应从固定 SQL 中核验确切函数签名，使用 Stripe 测试账户验证，并在生产使用前监控请求数、延迟、部分失败、分页与速率限制。
