## 用法

来源：

- [Official README and WIP warning](https://github.com/brahmlower/pgzan/blob/b2e45f17059639863dae8fdadff805fe51ca55ac/readme.md)
- [Rust implementation](https://github.com/brahmlower/pgzan/blob/b2e45f17059639863dae8fdadff805fe51ca55ac/src/lib.rs)
- [Extension control file](https://github.com/brahmlower/pgzan/blob/b2e45f17059639863dae8fdadff805fe51ca55ac/pgzan.control)

pgzan 是一个从 PostgreSQL 调用 Oso/Polar 授权策略的概念验证，长期目标是支持类似 Zanzibar 的行级授权。上游明确警告它仍在开发，绝不能靠近生产数据库。所核验代码只有一条硬编码策略，并未实现 README 中设想的可配置 ACL。

### 核心流程

只能在一次性数据库中使用实际存在的单参数函数。它解析包含 ID 和角色的 JSON 用户，然后检查固定 Organizations 资源上的固定 `update` 权限：

```sql
CREATE EXTENSION pgzan;

SELECT pgzan_check(
  '{"id":"07b30b3a-8da9-465e-96ef-4054f870cd8a", "role":"manager"}'
);

SELECT pgzan_check(
  '{"id":"07b30b3a-8da9-465e-96ef-4054f870cd8a", "role":"readonly"}'
);
```

在编译策略下，第一次调用返回 true，第二次返回 false。这只能演示求值；它没有连接会话身份、表行、请求操作或外部关系存储。

### 重要对象

- `pgzan_check` 解析一个 JSON 文本参数并计算硬编码授权决策。
- `hello_pgzan` 返回开发测试所用的固定问候语。
- 编译策略定义 readonly 和 manager 角色，以及固定资源名和权限集合。

### 安全与兼容性说明

README 的示例 RLS 调用了与源码不同的函数签名，因此并非已经实现的流程。调用方能够控制 JSON 角色值，所以如果没有可信边界构造输入，直接把布尔结果用作 RLS 决策，就会让调用方自称 manager。解析错误可以中止语句，而且每次调用都会新建 Oso 求值器。不要用 pgzan 保护数据；应评估仍受维护的 pgauthz 扩展或外部授权服务。
