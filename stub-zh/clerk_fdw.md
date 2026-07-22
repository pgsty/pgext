## 用法

来源：

- [Pinned official README](https://github.com/Jayko001/clerk_fdw/blob/ff55a536f658b2e4461407b8e10b3e47a6dfb5d1/README.md)
- [Pinned FDW implementation](https://github.com/Jayko001/clerk_fdw/blob/ff55a536f658b2e4461407b8e10b3e47a6dfb5d1/src/lib.rs)

`clerk_fdw` 是面向 Clerk 用户管理 API 的只读外部数据包装器。它把用户、组织和组织成员关系暴露为外部表；扫描执行时才从 Clerk 获取行，而不是将其存入 PostgreSQL。

### 核心流程

0.3.3 版安装处理器与验证器函数。定义包装器、包含 Clerk 密钥的服务器，以及一个或多个外部表：

```sql
CREATE EXTENSION clerk_fdw;

CREATE FOREIGN DATA WRAPPER clerk_wrapper
  HANDLER clerk_fdw_handler
  VALIDATOR clerk_fdw_validator;

CREATE SERVER my_clerk_server
  FOREIGN DATA WRAPPER clerk_wrapper
  OPTIONS (api_key '<clerk-secret-key>');

CREATE FOREIGN TABLE clerk_users (
  user_id text,
  first_name text,
  last_name text,
  email text,
  gender text,
  created_at bigint,
  updated_at bigint,
  last_sign_in_at bigint,
  phone_numbers bigint,
  username text,
  attrs jsonb
)
SERVER my_clerk_server
OPTIONS (object 'users');

SELECT user_id, email, attrs FROM clerk_users;
```

如果没有 `api_key`，实现会从 PostgreSQL 服务进程环境读取 `CLERK_API_KEY`。服务器选项可能含有凭据，应保护外部服务器定义和系统目录的访问权限。

### 支持的外部对象

- `object 'users'`：用户标识、姓名、主邮箱、时间戳、用户名，以及 `attrs` 中的完整响应。
- `object 'organizations'`：组织标识、名称、slug、时间戳、创建者，以及 `attrs`。
- `object 'organization_memberships'`：`user_id`、`organization_id` 和 `role`。

列名和类型必须与 FDW 实现的映射一致。需要未扁平化的 API 对象时，应包含 `attrs jsonb` 列。

### 查询与性能边界

远程抓取时，FDW 会忽略 PostgreSQL 的过滤条件、排序键和行数限制。`WHERE organization_id = 'org_123'` 之类的谓词在扫描之后求值，不会下推到 Clerk。

成员关系扫描会先枚举全部组织，再逐个请求其成员。实现会对受到限流的成员请求进行指数退避重试，因此可能耗时很长；重复分析时应将结果物化到本地表。实现仅提供扫描回调，因此不要期望支持 `INSERT`、`UPDATE` 或 `DELETE`。

0.3.3 版面向 PostgreSQL 14–17 构建，不可迁移，创建时需要超级用户权限。它不要求 `shared_preload_libraries`。
