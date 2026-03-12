

## 用法

> [pg_fact_loader: 使用 Postgres 构建事实表](https://github.com/enova/pg_fact_loader)

使用基于队列的变更数据捕获来构建和维护事实表。通过处理审计/变更日志表来增量更新事实表。

### 启用

```sql
CREATE EXTENSION pg_fact_loader;
```

可选地与 pglogical 配合用于基于副本的设置：

```sql
CREATE EXTENSION pglogical;
CREATE EXTENSION pglogical_ticker;
CREATE EXTENSION pg_fact_loader;
```

### 工作流程

1. **复制源表** 到报表数据库（通过 pglogical 或其他方式）
2. **创建审计/变更日志表** 在 OLTP 系统中为源表创建
3. **创建事实表** 结构用于聚合数据
4. **创建合并函数** 接受一个键 ID 并返回事实表的一行
5. **配置 pg_fact_loader** 将队列表关联到事实表
6. **回填** 初始化事实表数据
7. **调度** 工作进程持续处理变更

### 配置表

```sql
-- 注册事实表
INSERT INTO fact_loader.fact_tables (fact_table_relid, fact_table_agg_proid, ...)
VALUES ('public.customers_fact'::regclass, 'customers_fact_merge'::regproc, ...);

-- 注册队列（审计）表
INSERT INTO fact_loader.queue_tables (queue_table_relid, queue_of_base_table_relid, ...)
VALUES ('audit.customers_audit'::regclass, 'public.customers'::regclass, ...);

-- 将队列表与事实表通过合并函数关联
INSERT INTO fact_loader.queue_table_deps
    (fact_table_id, queue_table_id, insert_merge_proid, update_merge_proid, delete_merge_proid)
VALUES (1, 1, 'customers_fact_merge'::regproc, 'customers_fact_merge'::regproc, 'customers_fact_merge'::regproc);

-- 定义如何从队列条目中获取键
INSERT INTO fact_loader.key_retrieval_sequences
    (queue_table_dep_id, return_columns, is_fact_key)
VALUES (1, '{customer_id}', true);
```

### 运行工作进程

```sql
-- 处理待处理的变更
SELECT fact_loader.worker();

-- 定期调度此操作（例如通过 pg_cron 每几秒运行一次）
```

### 初始回填

```sql
-- 对每条现有记录运行合并函数
SELECT customers_fact_merge(customer_id) FROM customers;
```

### 添加批次 ID 字段

```sql
SELECT fact_loader.add_batch_id_fields();
```

### 主要功能

- 基于队列的增量事实表更新
- 支持插入、更新和删除事件
- 支持多级键检索（通过多个表的联接）
- 事实表依赖链（父事实更新后再更新子事实）
- 处理前检查复制延迟（与 pglogical 配合使用时）
