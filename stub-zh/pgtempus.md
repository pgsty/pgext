## 用法

来源：

- [官方 README](https://gitlab.com/tempus-projects/tempus_pg/-/blob/v2.1.0/README.md)
- [扩展控制文件](https://gitlab.com/tempus-projects/tempus_pg/-/blob/v2.1.0/pgext/pgtempus.control)
- [版本化安装 SQL](https://gitlab.com/tempus-projects/tempus_pg/-/blob/v2.1.0/pgext/pgtempus--2.1.sql)

`pgtempus` 把多模式交通图加载到 PostgreSQL backend，并按照时间、费用、换乘与交通方式切换成本计算 one-to-many 或 many-to-one 路径。

### 启用

安装 2.1 原生库与扩展文件后，在目标 schema 中创建可迁移扩展：

```sql
CREATE SCHEMA tempus;
CREATE EXTENSION pgtempus SCHEMA tempus;
```

控制文件不要求超级用户安装或预加载。上游没有发布当前 PostgreSQL 大版本支持矩阵，因此部署前必须用准确的服务器构件验证 2.1。

### 构建会话图

`tempus_build_multimodal_graph` 接收描述 transport mode、node 与 arc 的 SQL 查询。查询字符串在调用者的数据库上下文中执行。

```sql
SELECT tempus_build_multimodal_graph(
  'city',
  'SELECT id, name, category, traffic_rules FROM transport_modes',
  'SELECT id, parking_transport_modes, x, y, z FROM nodes',
  'SELECT id, node_from_id, node_to_id, is_pt, traffic_rules FROM arcs'
);

SELECT * FROM tempus_loaded_graphs();
DELETE FROM session_graph WHERE id = 'city';
```

图保存在 backend 内存中，并在会话结束时消失。连接池复用会话时必须显式设计，不能假定其他 backend 能看到已加载图。

### 成本与路由

使用 `tempus_set_static_road_section_costs`、`tempus_set_pt_section_timetable`、`tempus_set_arcs_sequence_costs` 及相关函数附加 generalized cost。`tempus_one_to_many_paths` 处理 departure-time 搜索，`tempus_many_to_one_paths` 处理 arrive-before 搜索。

API 接收调用者提供的 SQL 文本，也可能把大型图集读入 backend 内存。应只向可信角色开放执行权，校验传入查询，并按照图与结果基数规划会话资源。2.1 已停止上游活动；现代 PostgreSQL 兼容性与恢复行为必须作为部署测试，而不是默认假设。

