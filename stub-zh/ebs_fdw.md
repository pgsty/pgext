## 用法

来源：

- [锁定提交的上游示例](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/ebs_fdw/README.md)
- [锁定提交的 FDW 实现](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/ebs_fdw/src/lib.rs)
- [锁定提交的 Cargo 清单](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/ebs_fdw/Cargo.toml)
- [锁定提交的扩展控制文件](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/ebs_fdw/ebs_fdw.control)

ebs_fdw 是一个只读实验性外部数据包装器，它调用 AWS EC2 DescribeVolumes API，并暴露五个 EBS 卷属性。目录和 Cargo 清单将版本标识为 0.0.0，而包装器宏宣称 0.1.0；应核对实际软件包产物，不要只依赖某一个版本字符串。

### 外部服务器与表

上游构建不会创建外部数据包装器目录对象，所以在 CREATE EXTENSION 之后还需显式连接处理器和校验器：

```sql
CREATE EXTENSION ebs_fdw;

CREATE FOREIGN DATA WRAPPER ebs
  HANDLER ebs_fdw_handler
  VALIDATOR ebs_fdw_validator;

CREATE SERVER aws_ebs
  FOREIGN DATA WRAPPER ebs
  OPTIONS (region 'eu-west-3');

CREATE FOREIGN TABLE ebs_volumes (
  id text,
  name text,
  type text,
  size integer,
  encrypted boolean
) SERVER aws_ebs;

SELECT id, name, type, size, encrypted
FROM ebs_volumes;
```

PostgreSQL 服务器进程从 AWS SDK 默认凭据提供链获取 AWS 凭据。复核的代码中没有用户映射凭据选项。服务器的 region 选项必填，缺失时会引发 Rust panic。

### AWS、规划器与兼容性限制

每次扫描都会在不带过滤器的情况下调用 DescribeVolumes。源码虽然包含未使用的等值过滤辅助函数，begin_scan 仍将过滤器设为 none，并忽略 PostgreSQL 条件、排序键和 LIMIT。因此，即使查询条件很精确，也会分页读取 AWS 身份可见的每个 EBS 卷。只会映射 id、Name 标签得到的 name、卷类型、大小和加密状态；未识别的外部列会返回 NULL。

包装器会创建单线程 Tokio 运行时，并在 AWS 请求和分页期间阻塞 PostgreSQL 后端。SDK 错误会被转换为 panic。应设置语句超时，测量 API 限流和延迟，并避免在扫描周围持有应用锁。重复报表应将结果缓存到本地。

可以使用该服务器的数据库角色，能以数据库主机的操作系统身份发起调用。应只向专用资产盘点角色授予 USAGE，将 IAM 权限限制为 DescribeVolumes，限制出站网络，并在 AWS 侧审计活动，不应把 PostgreSQL 查询日志当作权威 API 记录。

锁定源码使用旧版 AWS SDK、pgrx 0.9.7 和 supabase-wrappers 0.1.15，且功能标志只覆盖 PostgreSQL 11 至 15。仓库是一个文档稀少的多扩展实验工作区。应针对目标大版本和 libc 精确编译，并在非实验环境使用前测试凭据刷新、分页、空账户、限流、权限拒绝、取消和后端恢复。
