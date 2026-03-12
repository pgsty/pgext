


## 用法

> [omni_datasets: 数据集配置](https://docs.omnigres.org/omni_datasets/northwind/)

`omni_datasets` 扩展提供对示例数据集的访问，目前包含 Northwind 数据库。

### Northwind 数据集

```sql
CREATE EXTENSION omni_datasets;
CREATE SCHEMA northwind;
SELECT omni_datasets.instantiate_northwind(schema => 'northwind');
```

Northwind 数据集是经典的示例数据库，最初来自 Microsoft Access/SQL Server，使用 Yugabyte 维护的 PostgreSQL 兼容版本。它提供了一个小而真实的商业数据集，包含客户、订单、产品、员工等表。
