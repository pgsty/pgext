## 用法

来源：

- [pgfr_analyze v2.29.2 说明文档](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_analyze/README.md)
- [pgfr_analyze 控制文件](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_analyze/extension.control)
- [pg_flight_recorder v2.29.2 参考手册](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/REFERENCE.md)
- [v2.29.2 发行说明](https://github.com/dventimisupabase/pg_flight_recorder/releases/tag/v2.29.2)

pgfr_analyze 是 PostgreSQL Flight Recorder 的读取分析部分。它解释由 pgfr_record 捕获的历史数据，用于比较时间段、汇总等待事件、组装事故时间线、检测退步并估算影响范围。在安装 pgfr_record 之后，请使用它进行诊断而不是收集数据。

### 安装分析层

    CREATE EXTENSION pg_cron;
    CREATE EXTENSION pgfr_record;
    CREATE EXTENSION pgfr_analyze;
    SELECT pgfr_record.enable();

pgfr_analyze 依赖于记录器对象，在没有收集样本之前没有任何有用的历史数据。它不需要自己的后台工作进程。

### 开始事故调查

在已知时间戳周围总结发生了什么：

    SELECT *
    FROM pgfr_analyze.what_happened_at(
      now() - interval '15 minutes'
    );

在一个时间段上构建有序的事故视图：

    SELECT *
    FROM pgfr_analyze.incident_timeline(
      now() - interval '30 minutes',
      now()
    );

使用 compare 对比可疑时间段与基线，然后通过 wait_summary 或 anomaly_report 窄化结果。始终使用 psql 或 pg_get_function_arguments 检查安装的函数签名，因为可选参数可能在不同版本之间有所变化。

### 分析索引

- compare：对比基准窗口和事故窗口中的指标。
- wait_summary：聚合采样等待事件。
- report 和 anomaly_report：生成广泛或侧重异常的总结。
- what_happened_at：围绕一个时间戳检查观察结果。
- incident_timeline：在一个范围内按重要事件排序。
- detect_regressions 和 detect_query_storms：标记退步行为或查询突发情况。
- blast_radius：识别受影响的会话或工作负载。
- capacity_summary 和 capacity_report：总结与容量相关的信号。
- 配置分析：将相关设置变化关联到该时间段。

### 解释流程

1. 确认 pgfr_record.health_check() 及可用样本间隔。
2. 选择具有可比工作负载的明确基准窗口和事故窗口。
3. 使用 compare 和 wait_summary 定位主要的变化。
4. 关联活动、锁、复制、真空和配置证据。
5. 使用 PostgreSQL 日志、查询计划、应用程序遥测数据及主机指标验证假设。

### 注意事项

- 结果是观察和启发式信息，而不是因果关系的证明。稀疏采样可能会错过短暂事件。
- 计数器重置、扩展更新、重启、保留间隔以及工作负载季节性都可能扭曲比较结果。
- 一些发现需要在启用 pg_stat_statements 并且足够大时才会更丰富。
- 分析输出的访问权限可能会暴露查询文本和操作元数据；相应地限制权限。
- v2.29.2 主要改进了 pgfr_record 中托管服务调度行为。它并不取代验证实际运行收集任务的需求。
