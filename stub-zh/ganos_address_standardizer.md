## 用法

来源：

- [ApsaraDB RDS for PostgreSQL 扩展支持矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_address_standardizer` 是阿里云基于 PAGC 规则的 GanosBase 地址标准化扩展。它仅由服务商提供，可用版本取决于 ApsaraDB RDS PostgreSQL 大版本、实例版本和引擎小版本。当前 Standard Edition 矩阵为 PostgreSQL 17 提供 7.0，为 PostgreSQL 12 至 16 提供 6.9，为 PostgreSQL 10 和 11 提供 6.3，而 PostgreSQL 18 不受支持。

### 安装前验证实例

具体实例应以服务目录为准。创建扩展前同时检查基础扩展和可选的美国数据包：

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name IN (
  'ganos_address_standardizer',
  'ganos_address_standardizer_data_us'
)
ORDER BY name;

CREATE EXTENSION ganos_address_standardizer;
CREATE EXTENSION ganos_address_standardizer_data_us;
```

只有规则和参考数据与地址群体匹配时才安装美国数据扩展。应在实例上清点已安装函数及签名，不要假定社区 PostGIS 示例与 Ganos 构建完全一致。

### 服务商与数据质量边界

目录中的 7.0 并不是适用于所有 PostgreSQL 大版本的统一版本。服务安全限制还可能禁止在某些引擎小版本中新建扩展，同时不影响已有安装。恢复、创建副本、大版本升级或跨区域迁移前必须重新确认可用性。

地址标准化属于解析和归一化，不能证明地址真实存在或归属于某个人。结果取决于规则、参考数据、语言、缩写和区域习惯。应保留原始输入，记录扩展和数据包版本，测试有代表性的国际地址与错误值，并把歧义结果交由复核。备份与迁移仍绑定于阿里云扩展生命周期和受支持的引擎组合。
