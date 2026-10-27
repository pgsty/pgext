## 用法

来源：

- [README](https://github.com/apjapj/pg_limit_runway/blob/001157bde784bd3f9beb002b7a392e740c480f69/README.md)
- [Control](https://github.com/apjapj/pg_limit_runway/blob/001157bde784bd3f9beb002b7a392e740c480f69/pg_limit_runway.control)
- [Install and grants](https://github.com/apjapj/pg_limit_runway/blob/001157bde784bd3f9beb002b7a392e740c480f69/INSTALL.md)
- [Install SQL](https://github.com/apjapj/pg_limit_runway/blob/001157bde784bd3f9beb002b7a392e740c480f69/pg_limit_runway--0.7.0.sql)

`pg_limit_runway` 跟踪数据库硬限制，并根据采样速率预测耗尽时间。发行标签 v0.7.1 所带的控制文件与扩展 SQL 版本为 0.7.0。它是纯 SQL 扩展，测试覆盖 PostgreSQL 14–19rc1，需要超级用户安装，不要求预加载或重启。

### 查看余量

```sql
CREATE EXTENSION pg_limit_runway;
SELECT * FROM limit_runway_findings();
SELECT * FROM limit_runway_unwatched();
```

目录读取器覆盖序列、identity 范围、较窄的外键列、事务与 MultiXact 年龄、分区余量、复制槽保留、连接余量和角色密码到期。无法监测的对象与缺失速率会显式列出；没有诊断结果不证明所有资源限制都安全。按速率预测日期需要重复调度采样，不同时间窗口可能给出不同结果。

### 权限与调度

应按文档授予对象、函数权限、统计信息访问和序列权限。采样、修剪和告警另需写权限及外部调度。操作辅助函数只返回供审阅的命令，不自动修复。生成的脚本安装路径适用于无法安装扩展文件的环境，不应与同库 CREATE EXTENSION 混用。升级应遵循 SQL 0.7.0 路径，不要假设存在 0.7.1 SQL 版本。
