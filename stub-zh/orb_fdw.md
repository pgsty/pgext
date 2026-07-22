## 用法

来源：

- [上游 README](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/README.md)
- [FDW 实现](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/src/lib.rs)
- [Rust 包清单](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/Cargo.toml)

`orb_fdw` 版本 `0.13.3` 是一个只读外部数据包装器，用于通过 Orb API 查询指定的 Orb 计费对象。

### 核心流程

```sql
CREATE EXTENSION orb_fdw;
CREATE FOREIGN DATA WRAPPER orb_wrapper
  HANDLER orb_fdw_handler VALIDATOR orb_fdw_validator;
CREATE SERVER my_orb_server FOREIGN DATA WRAPPER orb_wrapper
  OPTIONS (api_key '<orb secret key>');

CREATE FOREIGN TABLE orb_customers (
  user_id text,
  organization_id text,
  email text,
  created_at timestamp,
  attrs jsonb
) SERVER my_orb_server OPTIONS (object 'customers');

SELECT user_id, email FROM orb_customers;
```

服务器选项接受 `api_key`；若省略，服务器进程会读取 `ORB_API_KEY`。外部表必须把 `object` 设为 `customers`、`subscriptions` 或 `invoices` 之一。只应定义文档列出的列名与兼容的 PostgreSQL 类型；`attrs` 保留完整响应对象。

### 运行边界

该实现只提供扫描回调，因此应视为只读。它会分页读取远端 API，而筛选、排序和限制由 PostgreSQL 在本地完成；看似选择性很高的查询仍可能产生大量网络请求。

存入服务器选项的 API 密钥可被拥有足够权限的系统目录读取者看到；运行条件允许时优先使用进程环境，并限制服务器与外部表的所有权。需要为 HTTPS 故障、Orb API 模式变化、配额、速率限制和延迟做好处理。上游 README 描述的是较早的 `0.13.0` 前置条件，而本次核对的 Rust 清单为 `0.13.3`；应测试实际打包构建，不能直接按 README 推断兼容性。
