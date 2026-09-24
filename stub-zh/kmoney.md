## 用法

来源：

- [README.md](https://github.com/pt-immer/kamu-money-pg/blob/a83ceeaeb6d5c97cb6071c23c47bec172306842a/README.md)
- [DESIGN.md](https://github.com/pt-immer/kamu-money-pg/blob/a83ceeaeb6d5c97cb6071c23c47bec172306842a/DESIGN.md)
- [kamu-money-pg/kmoney.control](https://github.com/pt-immer/kamu-money-pg/blob/a83ceeaeb6d5c97cb6071c23c47bec172306842a/kamu-money-pg/kmoney.control)
- [kamu-money-pg/Cargo.toml](https://github.com/pt-immer/kamu-money-pg/blob/a83ceeaeb6d5c97cb6071c23c47bec172306842a/kamu-money-pg/Cargo.toml)

`kmoney` 提供定长、18 位小数的货币类型；每个 SQL 类型编码一个 ISO 4217 币种，并另有用于交换的混合币种载体。

### 启用

0.2.1 版本为 PostgreSQL 15–18 提供大版本 feature。安装匹配的 pgrx 构件后，以超级用户创建不可迁移、非 trusted 的扩展：

```sql
CREATE EXTENSION kmoney;
```

无需预加载或重启。上游通过带有 YugabyteDB 适配 feature 的 pgrx fork 构建；部署前必须验证准确的服务器/内核构件。

### 币种专用类型

每个币种都有独立的 16 字节类型，例如 `kmoney_usd` 或 `kmoney_idr`。解析器接受无标签金额或匹配的币种标签；跨币种操作符不存在。

```sql
SELECT '10.50'::kmoney_usd;
SELECT ('1.25'::kmoney_usd + '2.75'::kmoney_usd)::text;
SELECT '1.00'::kmoney_usd + '1.00'::kmoney_idr;
```

`kmoney_mixed` 把币种代码与金额一起保存，用于异构值。它有意不提供 `sum` 聚合；应对币种专用列执行聚合。

### 除法、分配与索引

除法同时返回商与余数，使舍入损失显式可见。分配函数按照整数权重切分且保持总额不变。

```sql
SELECT quotient, residue
FROM kmoney_usd_div('10.00'::kmoney_usd, 3, 'half_even');

SELECT unnest(kmoney_usd_allocate('10.00'::kmoney_usd, ARRAY[1,1,1]));
```

0.2.1 有意不提供 B-tree 或 hash 操作符类，因此这些自定义类型不能直接支撑普通 B-tree/hash 索引。取值范围限制为 `numeric(36,18)` 的数量级，并会拒绝过高精度，而不是静默舍入。

扩展现已迁移到独立的 kamu-money-pg 仓库。0.2.1 工作空间通过上游记录的 YugabyteDB 分支使用 pgrx 0.19.2，PostgreSQL 15–18 的类型和算术约定保持不变。
