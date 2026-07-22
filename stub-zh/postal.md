## 用法

来源：

- [Official README](https://github.com/pramsey/pgsql-postal/blob/105a6ca855dbf4d33b160ca50d00c3c76a7e2061/README.md)
- [Extension SQL](https://github.com/pramsey/pgsql-postal/blob/105a6ca855dbf4d33b160ca50d00c3c76a7e2061/postal--1.0.sql)
- [libpostal wrapper implementation](https://github.com/pramsey/pgsql-postal/blob/105a6ca855dbf4d33b160ca50d00c3c76a7e2061/postal.c)

postal 在 PostgreSQL 内暴露 libpostal 地址规范化及解析能力。它可以把地址展开为多个规范化候选项，或用 JSONB 标记地址组成部分。它适合搜索预处理和数据清洗，但规范化只是概率性预处理，不能证明地址真实存在或可以投递。

### 核心流程

服务器管理员安装 libpostal 及其模型数据后，创建扩展并调用两个函数：

```sql
CREATE EXTENSION postal;

SELECT unnest(postal_normalize('412 first ave, victoria, bc'));

SELECT postal_parse('412 first ave, victoria, bc');
```

应把原始输入与派生值一同保存。要按应用规则选择一个规范化结果，而不能假设第一个候选项就是权威形式。

### 重要对象

- `postal_normalize(text)` 返回包含 libpostal 规范化候选项的 `text[]`。
- `postal_parse(text)` 返回 `jsonb` 对象，其键标记门牌号、道路、城市或州等组成部分。

### 运维说明

libpostal 第一次调用有明显初始化延迟，并会在每个调用它的后端中载入大型模型数据；上游观察到每个活跃后端约占 1 GB。应限制并发，或把规范化隔离到受控工作进程池。结果依赖已安装的 libpostal 版本和数据模型，因此升级可能改变派生输出。所核验包装器面向 PostgreSQL 9.4 时代的 JSONB，且没有 Windows 安装路径；必须测试现代服务器兼容性及有代表性的国际地址。
