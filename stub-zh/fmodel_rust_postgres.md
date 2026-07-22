## 用法

来源：

- [官方 README](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/README.md)
- [扩展控制文件](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/fmodel_rust_postgres.control)
- [Rust 扩展入口](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/src/lib.rs)
- [事件溯源模式](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/sql/event_sourcing.sql)

`fmodel_rust_postgres` 在 PostgreSQL 中嵌入了一个具体的 Rust 事件溯源演示。所核对的源码使用 fmodel-rust 建模餐厅与订单命令；它是项目模板和领域示例，并非可通过 SQL 配置的通用事件溯源框架。

### 核心流程

```sql
CREATE EXTENSION fmodel_rust_postgres;

SELECT *
FROM handle(
  '{"CreateRestaurant":{"name":"Ada Cafe","menu":[]}}'::Command
);

SELECT *
FROM handle_all(ARRAY[
  '{"CreateRestaurant":{"name":"North Cafe","menu":[]}}'::Command,
  '{"CreateRestaurant":{"name":"South Cafe","menu":[]}}'::Command
]);
```

`handle(Command)` 为单个命令返回一组 `Event` 值。`handle_all(Command[])` 在一个 PostgreSQL 事务中处理命令批次，后续命令能看到之前的结果，任一错误都会回滚整个批次。安装对象还包括 `deciders`、`events`、`restaurants`、`orders`，以及投影餐厅和订单事件并检查前序事件链的触发器。

### 领域与安全边界

`Command`、`Event` 类型及其 JSON 表示是编译进 Rust 的领域类型。新增领域、修改载荷或改变投影行为都需要修改 Rust 源码并重建扩展；任意应用不能通过 SQL 注册决策器。初始化脚本会创建名称通用且未限定模式的对象，因此启用前应规划专用数据库或模式，并检查命名冲突。

事件表和函数会继承普通的 PostgreSQL 默认权限，除非管理员主动撤销。直接写入或对表执行 DDL 可以绕过预期命令路径，扩展也没有加入行级安全、租户隔离或授权策略。应限制表写入和函数执行，暴露应用专用 API，并在生产使用前设计备份、重放、保留与模式演进流程。固定构建使用 pgrx 0.16 及 PostgreSQL 13–16 特性；二进制必须匹配服务器确切大版本，并需要测试升级。
