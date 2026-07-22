## 用法

来源：

- [腾讯云 PostgreSQL：使用 starrocks_fdw](https://cloud.tencent.com/document/product/409/111197)

`starrocks_fdw` 是腾讯云 PostgreSQL 的云厂商扩展，把 StarRocks 表暴露为 PostgreSQL 外部表，用于冷热数据分离。其用法模型兼容 `mysql_fdw`，但包含腾讯云连接路由参数；这是托管云厂商能力，并非可移植的上游扩展工作流。

### 核心流程

StarRocks 服务必须已经能够从腾讯云 PostgreSQL 实例访问。依次创建扩展、外部服务器、用户映射，以及指定 StarRocks 数据库和表的外部表。

```sql
CREATE EXTENSION starrocks_fdw;

CREATE SERVER server_sr
  FOREIGN DATA WRAPPER starrocks_fdw
  OPTIONS (
    host '10.0.0.10', port '9030',
    instanceid 'ins-example', access_type '2',
    region 'ap-guangzhou', uin '100000000000',
    own_uin '100000000000', vpcid 'vpc-example',
    subnetid 'subnet-example'
  );

CREATE USER MAPPING FOR dbadmin
  SERVER server_sr
  OPTIONS (username 'root', password 'replace-me');

CREATE FOREIGN TABLE sr_m (
  id integer NOT NULL,
  str integer
)
SERVER server_sr
OPTIONS (dbname 'my_database', table_name 'sr_m1');

SELECT * FROM sr_m;
```

腾讯云示例还把外部表挂接为本地分区表的一个分区，使本地分区与 StarRocks 后端分区能够通过同一个父表查询。

### 云厂商与安全边界

官方文档声明只支持 `v13.14_r1.17` 或更高版本的腾讯云 PostgreSQL 13 构建。可用性、服务器选项值和网络路由由腾讯云服务控制；应针对具体实例确认，不能假设它兼容社区 PostgreSQL。

凭据保存在 PostgreSQL 用户映射中，服务器选项还包含云账号与网络标识。应限制所有权和目录可见性，使用最小权限的专用 StarRocks 账号，并保持 PostgreSQL 与 StarRocks 列类型一致。远程读取同步依赖 StarRocks 与 VPC 网络路径，因此查询规划和监控必须考虑远程延迟与故障。
