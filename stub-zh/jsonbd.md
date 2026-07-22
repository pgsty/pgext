## 用法

来源：

- [官方 jsonbd README](https://github.com/postgrespro/jsonbd/blob/6a8aefcca9999596d75212a0f68c28590b5c3570/README.md)
- [0.1 版扩展 SQL](https://github.com/postgrespro/jsonbd/blob/6a8aefcca9999596d75212a0f68c28590b5c3570/jsonbd--0.1.sql)
- [压缩实现与预加载设置](https://github.com/postgrespro/jsonbd/blob/6a8aefcca9999596d75212a0f68c28590b5c3570/jsonbd.c)
- [PostgreSQL 列压缩补丁提案](https://commitfest.postgresql.org/15/1294/)

`jsonbd` 是一个未完成的实验项目，通过共享字典压缩重复的 JSONB 对象键。它不能在未经修改的标准 PostgreSQL 上工作：扩展依赖一个已搁置的列压缩访问方法提案及其非核心 SQL 语法。

### 实验环境配置

使用所引用补丁构建兼容的 PostgreSQL 源码树后，预加载模块并重启服务器：

```conf
shared_preload_libraries = 'jsonbd'
jsonbd.workers_count = 1
jsonbd.queue_size = 1
```

然后安装扩展，并为 JSONB 列选择其压缩访问方法：

```sql
CREATE EXTENSION jsonbd;

CREATE TABLE jsonbd_demo (
    payload jsonb COMPRESSION jsonbd
);
```

扩展创建 `jsonbd` 压缩访问方法和 `jsonbd_dictionary` 表；后者以压缩对象 OID、字典 ID 和 JSON 键为核心字段。后台工作进程通过共享内存队列协调字典查询。此方案不压缩 JSON 标量值。

### 设置与限制

- `jsonbd.workers_count` 控制后台工作进程数；`0` 表示禁用，源码默认值为 `1`。
- `jsonbd.queue_size` 以 KiB 配置每个共享内存队列，范围为 `1` KiB 到 `1024` KiB；源码默认值为 `1` KiB。

共享内存大小和工作进程注册都在预加载期间完成，因此修改配置后必须重启。若未通过 `shared_preload_libraries` 加载该库，模块会报错。

上游明确声明项目尚未完成。其对象删除回调没有清理行为，字典生命周期和遗留行问题仍未解决。项目没有受支持的标准 PostgreSQL 兼容目标、稳定的磁盘格式承诺或生产迁移路径。`jsonbd` 只应在一次性集群中用于研究历史补丁；信任任何数据前都要测试重启、崩溃、备份、恢复和对象删除行为。
