## 用法

来源：

- [阿里云 RDS PostgreSQL 扩展矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`encdb` 是阿里云 RDS PostgreSQL 的服务商扩展，官方描述为提供机密数据库能力。它不是可移植的社区软件包：可用性、启用方式、升级时间、权限与受支持操作都由托管服务控制。

### 先检查可用性

当前标准版矩阵列出：PostgreSQL 11 至 17 提供 `encdb` `1.1.14`，PostgreSQL 10 提供 `1.1.13`，PostgreSQL 18 不提供。其他版本、地域、引擎小版本或服务更新可能不同，因此应查询目标实例，而不是依赖静态矩阵。

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'encdb';
```

只有服务商报告扩展可用、且已满足服务启用条件后，获授权的管理员才应安装：

```sql
CREATE EXTENSION encdb;

SELECT extversion
FROM pg_extension
WHERE extname = 'encdb';
```

### 托管服务边界

公开扩展矩阵没有说明 `encdb` 的可移植 SQL API、预加载设置、客户端要求、密钥管理流程或权限模型。使用前应取得实例对应的机密数据库文档与服务支持确认；不要从同名项目推断 API。需要测试备份恢复、复制与故障转移、客户端兼容性、加密密钥生命周期、版本升级，以及服务管理员仍可看到哪些元数据或查询结果。创建与移除扩展时应遵循阿里云限制。
