

## 用法

> [dbt2: OSDL-DBT-2 测试套件](https://github.com/osdldbt/dbt2)

`dbt2` 是一个面向 PostgreSQL 的 TPC-C 基准测试实现。该扩展提供了实现五种标准 TPC-C 事务类型的存储过程。

```sql
CREATE EXTENSION dbt2;
```

### TPC-C 事务类型

该扩展提供了五种标准 TPC-C 事务的存储过程：

- **新订单（New Order）**：创建包含多个行项目的新订单，并更新库存水平
- **支付（Payment）**：处理客户付款，更新仓库和区域余额
- **订单状态（Order Status）**：查询客户最近一笔订单的状态
- **发货（Delivery）**：处理所有区域中待发货的订单
- **库存水平（Stock Level）**：检查近期售出且库存偏低的商品数量

### 基准测试工作流

`dbt2` 系统由以下部分组成：

1. **数据库扩展**（`dbt2`）：TPC-C 事务的存储过程
2. **数据加载器**：用 TPC-C 数据填充基准测试表
3. **驱动程序**：生成模拟终端用户的事务负载
4. **客户端**：管理驱动程序与数据库之间的连接

### 运行基准测试

基准测试通常使用 `dbt2` 命令行工具运行（与扩展本身分开）：

```bash
# 构建基准测试数据库
dbt2 build --dbms pgsql --warehouses 10

# 运行基准测试
dbt2 run --dbms pgsql --warehouses 10 --duration 300 --connections 10

# 生成报告
dbt2 report --dbms pgsql
```

### TPC-C 模式

基准测试使用以下标准表：`warehouse`、`district`、`customer`、`history`、`new_order`、`orders`、`order_line`、`item` 和 `stock`。

有关详细的配置和调优选项，请参阅仓库中的 `doc/` 目录。
