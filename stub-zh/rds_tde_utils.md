## 用法

来源：

- [阿里云官方 rds_tde_utils 指南](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-tde-utils-extension-to-encrypt-and-decrypt-multiple-data-records-at-a-time)

`rds_tde_utils` 是阿里云 ApsaraDB RDS for PostgreSQL 扩展，在启用透明数据加密后，可加密或解密单张表及其索引，或者数据库中的全部表与索引。版本 `1.1` 只由服务商提供，并会执行可能规模很大的同步重写。

### 前置条件与设置

服务商要求 PostgreSQL 10 或以上版本、`20221030` 或更高的 RDS 小版本、已启用 TDE，以及特权账户：

```sql
CREATE EXTENSION rds_tde_utils;
```

改变加密状态前，应确认确切的受支持引擎，并取得可恢复的备份。

### 核心流程

使用 lazy 方法降低影响，或者在非高峰维护窗口使用 full 方法：

```sql
SELECT rds_tde_lazy_encrypt_table('public.orders'::regclass);
SELECT rds_tde_encrypt_table('public.orders'::regclass);

SELECT rds_tde_lazy_decrypt_table('public.orders'::regclass);
SELECT rds_tde_decrypt_table('public.orders'::regclass);
```

数据库级变体会处理全部表和索引：

```sql
SELECT rds_tde_lazy_encrypt_database();
SELECT rds_tde_encrypt_database();

SELECT rds_tde_lazy_decrypt_database();
SELECT rds_tde_decrypt_database();
```

### 方法选择

- 四个 `rds_tde_lazy_*` 函数采用文档所称类似 lazy VACUUM 的重写行为。
- 四个 full 函数采用文档所称类似 `VACUUM FULL` 的行为；服务商警告不要在高峰期运行。
- 表级函数接收 `regclass` 并包含表的索引；数据库级函数覆盖当前数据库中的全部表和索引。

### 运维注意事项

八个调用全部同步执行，只在处理完成后返回。大型数据库可能耗时很长，并出现锁、I/O、额外存储需求、副本延迟和延迟尖峰。应盘点对象大小、避开流量高峰、监控可用空间与复制，并在预发布实例上测试取消和重启行为。TDE 启用、KMS/密钥策略、快照、备份、副本与灾难恢复属于这些辅助函数之外的托管服务事项。应通过阿里云支持的控制面验证加密状态，不能把 SQL 成功返回等同于端到端密钥保护已经得到证明。
