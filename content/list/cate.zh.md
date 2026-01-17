---
title: "按分类"
weight: 100
---

PostgreSQL 扩展（444 ext / 375 pkg）归属 16 个分类。



| {{< category "time" >}} | {{< category "gis" >}}  | {{< category "rag" >}}   | {{< category "fts" >}}  | {{< category "olap" >}} | {{< category "feat" >}} | {{< category "lang" >}} | {{< category "type" >}} |
|------------------------|-------------------------|--------------------------|-------------------------|-------------------------|-------------------------|-------------------------|-------------------------|
| {{< category "util" >}} | {{< category "func" >}} | {{< category "admin" >}} | {{< category "stat" >}} | {{< category "sec" >}}  | {{< category "fdw" >}}  | {{< category "sim" >}}  | {{< category "etl" >}}  |


## TIME

时间时态扩展：时序数据库 TimescaleDB，时态数据库，版本控制表，定时任务，异步后台任务调度扩展。

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 1000 | {{< alias "timescaledb" >}} | 2.24.0 | 时序数据库扩展插件 |
| 1010 | {{< alias "timescaledb_toolkit" >}} | 1.22.0 | 超表分析查询，时间序列流式处理，以及其他SQL工具 |
| 1020 | {{< alias "timeseries" "pg_timeseries" >}} | 0.2.0 | 时序数据API封装 |
| 1030 | {{< alias "periods" >}} | 1.2.3 | 为 PERIODs 和 SYSTEM VERSIONING 提供标准 SQL 功能 |
| 1040 | {{< alias "temporal_tables" >}} | 1.2.2 | 时态表功能支持 |
| 1050 | {{< alias "emaj" >}} | 4.7.1 | 让数据库的子集具有细粒度日志和时间旅行功能 |
| 1060 | {{< alias "table_version" >}} | 1.11.1 | PostgreSQL 版本控制表扩展 |
| 1070 | {{< alias "pg_cron" >}} | 1.6.7 | 定时任务调度器 |
| 1080 | {{< alias "pg_task" >}} | 1.0.0 | 在特定时间点在后台执行SQL命令 |
| 1090 | {{< alias "pg_later" >}} | 0.4.0 | 执行查询，并在稍后异步获取查询结果 |
| 1100 | {{< alias "pg_background" >}} | 1.5 | 在后台运行 SQL 查询 |

## GIS

地理空间扩展：PostGIS，地理空间数据类型、函数与索引，天空索引 Q3C，OGR FDW， 寻路算法，地理正/逆查询。

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 1500 | {{< alias "postgis" >}} | 3.6.1 | PostGIS 几何和地理空间扩展 |
| 1501 | {{< alias "postgis_topology" "postgis" >}} | 3.6.1 | PostGIS 拓扑空间类型和函数 |
| 1502 | {{< alias "postgis_raster" "postgis" >}} | 3.6.1 | PostGIS 光栅类型和函数 |
| 1503 | {{< alias "postgis_sfcgal" "postgis" >}} | 3.6.1 | PostGIS SFCGAL 函数 |
| 1504 | {{< alias "postgis_tiger_geocoder" "postgis" >}} | 3.6.1 | PostGIS tiger 地理编码器和反向地理编码器 |
| 1505 | {{< alias "address_standardizer" "postgis" >}} | 3.6.1 | 地址标准化函数。 |
| 1506 | {{< alias "address_standardizer_data_us" "postgis" >}} | 3.6.1 | 地址标准化函数：美国数据集示例 |
| 1510 | {{< alias "pgrouting" >}} | 3.8.0 | 提供寻路能力 |
| 1520 | {{< alias "pointcloud" >}} | 1.2.5 | 提供激光雷达点云数据类型支持 |
| 1521 | {{< alias "pointcloud_postgis" "pointcloud" >}} | 1.2.5 | 将激光雷达点云与PostGIS几何类型相集成 |
| 1530 | {{< alias "h3" "pg_h3" >}} | 4.2.3 | H3六边形层级索引支持 |
| 1531 | {{< alias "h3_postgis" "pg_h3" >}} | 4.2.3 | H3与PostGIS集成的扩展插件 |
| 1540 | {{< alias "q3c" >}} | 2.0.1 | Q3C天空索引插件 |
| 1550 | {{< alias "ogr_fdw" >}} | 1.1.7 | GIS 数据外部数据源包装器 |
| 1560 | {{< alias "geoip" >}} | 0.3.0 | IP 地理位置扩展（围绕 MaxMind GeoLite 数据集的包装器） |
| 1570 | {{< alias "pg_polyline" >}} | 0.0.1 | Google快速Polyline编码解码扩展 |
| 1590 | {{< alias "pg_geohash" >}} | 1.0 | 使用GeoHash处理空间坐标的函数包 |
| 1650 | {{< alias "mobilitydb" >}} | 1.3.0 | MobilityDB地理空间投影数据管理分析平台 |
| 1651 | {{< alias "mobilitydb_datagen" "mobilitydb" >}} | 1.3.0 | MobilityDB随机数据生成函数 |
| 1680 | {{< alias "tzf" "pg_tzf" >}} | 0.2.3 | 快速根据GPS经纬度坐标查找时区 |
| 1690 | {{< alias "earthdistance" >}} | 1.2 | 计算地球表面上的大圆距离 |

## RAG

AI与RAG扩展插件：向量数据库，DiskANN 向量索引，相似度度量函数集，库内机器学习与推理 pgml，等等。

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 1800 | {{< alias "vector" "pgvector" >}} | 0.8.1 | 向量数据类型和 ivfflat / hnsw 访问方法 |
| 1810 | {{< alias "vchord" >}} | 1.0.0 | 使用Rust重写的高性能向量扩展 |
| 1820 | {{< alias "vectorscale" "pgvectorscale" >}} | 0.9.0 | 使用DiskANN算法对向量进行高效索引 |
| 1830 | {{< alias "vectorize" "pg_vectorize" >}} | 0.26.0 | 在PostgreSQL中封装RAG向量检索服务 |
| 1840 | {{< alias "pg_similarity" >}} | 1.0 | 提供17种距离度量函数 |
| 1850 | {{< alias "smlar" >}} | 1.0 | 高效的相似度搜索函数 |
| 1860 | {{< alias "pg_summarize" >}} | 0.0.1 | 使用LLM对文本字段进行总结 |
| 1870 | {{< alias "pg_tiktoken" >}} | 0.0.1 | 在PostgreSQL中计算OpenAI使用的Token数 |
| 1880 | {{< alias "pg4ml" >}} | 2.0 | PG4ML是一个机器学习框架 |
| 1890 | {{< alias "pgml" >}} | 2.10.0 | PostgresML：用SQL运行机器学习算法并训练模型 |

## FTS

全文检索扩展：ES 替代 pg_search，BM25，中文分词，欧洲语言分词字典 hunspell，模糊检索，2-gram/3-gram 索引。

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 2100 | {{< alias "pg_search" >}} | 0.21.2 | ParadeDB BM25算法全文检索插件，ES全文检索 |
| 2110 | {{< alias "pgroonga" >}} | 4.0.4 | 使用Groonga，面向所有语言的高速全文检索平台 |
| 2111 | {{< alias "pgroonga_database" "pgroonga" >}} | 4.0.4 | PGGroonga 数据库管理模块 |
| 2120 | {{< alias "pg_bigm" >}} | 1.2 | 基于二字组的多语言全文检索扩展 |
| 2130 | {{< alias "zhparser" >}} | 2.3 | 中文分词，全文搜索解析器 |
| 2140 | {{< alias "pg_bestmatch" >}} | 0.0.2 | 在数据库内生成BM25稀疏向量 |
| 2150 | {{< alias "vchord_bm25" >}} | 0.3.0 | BM25排序算法 |
| 2160 | {{< alias "pg_tokenizer" >}} | 0.1.1 | 用于全文检索的分词器 |
| 2170 | {{< alias "biscuit" "pg_biscuit" >}} | 2.2.2 | 使用IAM的高性能文本模式匹配 |
| 2180 | {{< alias "pg_textsearch" >}} | 0.4.0 | 带有BM25排序的全文搜索扩展 |
| 2270 | {{< alias "hunspell_cs_cz" >}} | 1.0 | Hunspell捷克语全文检索词典 |
| 2271 | {{< alias "hunspell_de_de" >}} | 1.0 | Hunspell德语全文检索词典 |
| 2272 | {{< alias "hunspell_en_us" >}} | 1.0 | Hunspell英语全文检索词典 |
| 2273 | {{< alias "hunspell_fr" >}} | 1.0 | Hunspell法语全文检索词典 |
| 2274 | {{< alias "hunspell_ne_np" >}} | 1.0 | Hunspell尼泊尔语全文检索词典 |
| 2275 | {{< alias "hunspell_nl_nl" >}} | 1.0 | Hunspell荷兰语全文检索词典 |
| 2276 | {{< alias "hunspell_nn_no" >}} | 1.0 | Hunspell挪威语全文检索词典 |
| 2277 | {{< alias "hunspell_pt_pt" >}} | 1.0 | Hunspell葡萄牙语全文检索词典 |
| 2278 | {{< alias "hunspell_ru_ru" >}} | 1.0 | Hunspell俄语全文检索词典 |
| 2279 | {{< alias "hunspell_ru_ru_aot" >}} | 1.0 | Hunspell俄语全文检索词典（来自AOT.ru小组） |
| 2380 | {{< alias "fuzzystrmatch" >}} | 1.2 | 确定字符串之间的相似性和距离 |
| 2390 | {{< alias "pg_trgm" >}} | 1.6 | 文本相似度测量函数与模糊检索 |

## OLAP

分析能力扩展：列式存储，DuckDB与外部数据源包装器，Parquet S3，数据冷热分级存储，分布式计算，透明分片，GPU加速

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 2400 | {{< alias "citus" >}} | 14.0.0 | Citus 分布式数据库 |
| 2401 | {{< alias "citus_columnar" "citus" >}} | 14.0.0 | Citus 列式存储引擎 |
| 2410 | {{< alias "columnar" "hydra" >}} | 1.1.2 | 开源列式存储扩展 |
| 2420 | {{< alias "pg_analytics" >}} | 0.3.7 | 由 DuckDB 驱动的数据分析引擎 |
| 2430 | {{< alias "pg_duckdb" >}} | 1.1.1 | 在PostgreSQL中的嵌入式DuckDB扩展 |
| 2440 | {{< alias "pg_mooncake" >}} | 0.2.0 | PostgreSQL列式存储表 |
| 2460 | {{< alias "pg_clickhouse" >}} | 0.1.2 | 从PostgreSQL中查询ClickHouse的接口 |
| 2470 | {{< alias "duckdb_fdw" >}} | 1.1.2 | DuckDB 外部数据源包装器 |
| 2480 | {{< alias "pg_parquet" >}} | 0.5.1 | 在PostgreSQL与本地/S3中的Parquet文件复制数据 |
| 2500 | {{< alias "pg_fkpart" >}} | 1.7.0 | 按外键实用程序进行表分区的扩展 |
| 2510 | {{< alias "pg_partman" >}} | 5.4.0 | 用于按时间或 ID 管理分区表的扩展 |
| 2520 | {{< alias "plproxy" >}} | 2.11.0 | 作为过程语言实现的数据库分区 |
| 2530 | {{< alias "pg_strom" >}} | 6.0 | 使用GPU与NVMe加速大数据处理 |
| 2590 | {{< alias "tablefunc" >}} | 1.0 | 交叉表函数 |

## FEAT

功能特性扩展：图数据库，Hyperloglog，Rum索引，GraphQL，JsonSchema，Hint，虚拟索引，增量物化视图，消息队列等等

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 2730 | {{< alias "age" >}} | 1.6.0 | Apache AGE，图数据库扩展 （Deb可用） |
| 2740 | {{< alias "hll" >}} | 2.19 | hyperloglog 数据类型 |
| 2750 | {{< alias "rum" >}} | 1.3.15 | RUM 索引访问方法 |
| 2760 | {{< alias "pg_ai_query" >}} | 0.1.1 | AI驱动的 Postgres SQL 查询生成 |
| 2780 | {{< alias "pg_ttl_index" >}} | 2.0.0 | 基于TTL索引的自动数据过期清理 |
| 2790 | {{< alias "pg_graphql" >}} | 1.5.12 | PG内的GraphQL支持 |
| 2800 | {{< alias "pg_jsonschema" >}} | 0.3.3 | 提供JSON Schema校验能力 |
| 2810 | {{< alias "jsquery" >}} | 1.2 | 用于内省 JSONB 数据类型的查询类型 |
| 2820 | {{< alias "pg_hint_plan" >}} | 1.8.0 | 添加强制指定执行计划的能力 |
| 2830 | {{< alias "hypopg" >}} | 1.4.2 | 假设索引，用于创建一个虚拟索引检验执行计划 |
| 2840 | {{< alias "index_advisor" >}} | 0.2.0 | 查询索引建议器 |
| 2850 | {{< alias "plan_filter" "pg_plan_filter" >}} | 0.0.1 | 使用执行计划代价过滤阻止特定查询语句 |
| 2860 | {{< alias "imgsmlr" >}} | 1.0 | 使用Haar小波分析计算图片相似度 |
| 2870 | {{< alias "pg_ivm" >}} | 1.13 | 增量维护的物化视图 |
| 2880 | {{< alias "pg_incremental" >}} | 1.2.0 | 增量处理流式事件 |
| 2890 | {{< alias "pgmq" >}} | 1.8.1 | 基于Postgres实现类似AWS SQS/RSMQ的消息队列 |
| 2900 | {{< alias "pgq" >}} | 3.5.1 | 通用队列的PG实现 |
| 2910 | {{< alias "orioledb" >}} | 1.5 | OrioleDB，下一代事务处理引擎 |
| 2920 | {{< alias "pg_cardano" >}} | 1.1.1 | Cardano相关工具包：加密函数，地址编解码，区块链处理 |
| 2930 | {{< alias "rdkit" >}} | 202503.1 | 在PostgreSQL化学领域数据管理功能 |
| 2940 | {{< alias "omni" "omnigres" >}} | 0.2.14 | PostgreSQL即平台，Omnigres主扩展与加载器 |
| 2941 | {{< alias "omni_auth" "omnigres" >}} | 0.1.3 | Omnigres 基础会话认证管理模块 |
| 2942 | {{< alias "omni_aws" "omnigres" >}} | 0.1.2 | Omnigres AWS S3 API封装 |
| 2943 | {{< alias "omni_cloudevents" "omnigres" >}} | 0.1.0 | Omnigres CloudEvents 支持 |
| 2944 | {{< alias "omni_containers" "omnigres" >}} | 0.2.0 | Omnigres Docker容器管理模块 |
| 2945 | {{< alias "omni_credentials" "omnigres" >}} | 0.2.0 | Omnigres 应用密钥管理模块 |
| 2946 | {{< alias "omni_csv" >}} | 0.1.1 | Omnigres CSV 工具箱 |
| 2947 | {{< alias "omni_datasets" >}} | 0.1.0 | Omnigres 数据库置备工具 |
| 2948 | {{< alias "omni_email" "omnigres" >}} | 0.1.0 | Omnigres Email 框架 |
| 2949 | {{< alias "omni_http" "omnigres" >}} | 0.1.0 | Omnigres 基本HTTP类型 |
| 2950 | {{< alias "omni_httpc" "omnigres" >}} | 0.1.10 | Omnigres HTTP客户端 |
| 2951 | {{< alias "omni_httpd" "omnigres" >}} | 0.4.11 | Omnigres HTTP服务器 |
| 2952 | {{< alias "omni_id" "omnigres" >}} | 0.4.3 | Omnigres ID身份数据类型 |
| 2953 | {{< alias "omni_json" "omnigres" >}} | 0.1.1 | Omnigres JSON工具箱 |
| 2954 | {{< alias "omni_kube" "omnigres" >}} | 0.4.2 | Omnigres Kubernetes集成模块 |
| 2955 | {{< alias "omni_ledger" "omnigres" >}} | 0.1.3 | Omnigres 金融账本模块 |
| 2956 | {{< alias "omni_manifest" "omnigres" >}} | 0.1.2 | Omnigres 包管理清单模块 |
| 2957 | {{< alias "omni_mimetypes" "omnigres" >}} | 0.1.0 | Omnigres MIME数据类型 |
| 2958 | {{< alias "omni_os" "omnigres" >}} | 0.1.1 | Omnigres 操作系统集成模块 |
| 2959 | {{< alias "omni_polyfill" "omnigres" >}} | 0.2.2 | Omnigres Postgres多态API |
| 2960 | {{< alias "omni_python" "omnigres" >}} | 0.1.1 | Omnigres 第一类Python支持模块 |
| 2961 | {{< alias "omni_regex" "omnigres" >}} | 0.1.0 | Omnigres PCRE兼容正则表达式模块 |
| 2962 | {{< alias "omni_rest" "omnigres" >}} | 0.1.1 | Omnigres REST API 工具包 |
| 2963 | {{< alias "omni_schema" "omnigres" >}} | 0.3.0 | Omnigres 高级模式管理组件 |
| 2964 | {{< alias "omni_seq" "omnigres" >}} | 0.1.1 | Omnigres 分布式整型序列号 |
| 2965 | {{< alias "omni_service" "omnigres" >}} | 0.1.0 | Omnigres 服务管理器 |
| 2966 | {{< alias "omni_session" "omnigres" >}} | 0.2.0 | Omnigres 会话管理器 |
| 2967 | {{< alias "omni_shmem" >}} | 0.1.0 | Omnigres 共享内存管理 |
| 2968 | {{< alias "omni_sql" "omnigres" >}} | 0.5.3 | Omnigres SQL编程组件 |
| 2969 | {{< alias "omni_sqlite" "omnigres" >}} | 0.2.2 | Omnigres 嵌入的SQLite支持 |
| 2970 | {{< alias "omni_test" "omnigres" >}} | 0.4.0 | Omnigres 测试框架 |
| 2971 | {{< alias "omni_txn" "omnigres" >}} | 0.5.0 | Omnigres 事务管理器模块 |
| 2972 | {{< alias "omni_types" "omnigres" >}} | 0.3.6 | Omnigres 高级数据类型模块 |
| 2973 | {{< alias "omni_var" "omnigres" >}} | 0.3.0 | Omnigres 局部变量模块 |
| 2974 | {{< alias "omni_vfs" "omnigres" >}} | 0.2.2 | Omnigres 虚拟文件系统 |
| 2975 | {{< alias "omni_vfs_types_v1" "omnigres" >}} | 0.1.0 | Omnigres 虚拟文件系统（v1） |
| 2976 | {{< alias "omni_web" "omnigres" >}} | 0.3.0 | Omnigres Web工具箱 |
| 2977 | {{< alias "omni_worker" "omnigres" >}} | 0.2.1 | Omnigres 通用Worker池 |
| 2978 | {{< alias "omni_xml" "omnigres" >}} | 0.1.2 | Omnigres XML工具包 |
| 2979 | {{< alias "omni_yaml" "omnigres" >}} | 0.1.0 | Omnigres YAML工具包 |
| 2990 | {{< alias "bloom" >}} | 1.0 | bloom 索引-基于指纹的索引 |

## LANG

存储过程语言扩展：使用各种编程语言开发，调试，打包，分发，测试 PostgreSQL 存储过程：Java，Js，Lua，R，SH，PRQL…

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 3000 | {{< alias "pg_tle" >}} | 1.5.2 | AWS 可信语言扩展 |
| 3010 | {{< alias "plv8" >}} | 3.2.4 | PL/JavaScript (v8) 可信过程程序语言 |
| 3011 | {{< alias "pljs" >}} | 1.0.4 | PL/JS 可信过程程序语言 |
| 3020 | {{< alias "pllua" >}} | 2.0.12 | Lua 程序语言 |
| 3021 | {{< alias "hstore_pllua" "pllua" >}} | 2.0.12 | Lua 程序语言的Hstore适配扩展 |
| 3030 | {{< alias "plluau" "pllua" >}} | 2.0.12 | Lua 程序语言（不受信任的） |
| 3031 | {{< alias "hstore_plluau" "pllua" >}} | 2.0.12 | Lua 程序语言的Hstore适配扩展（不受信任的） |
| 3040 | {{< alias "plprql" >}} | 18.0.0 | 在PostgreSQL使用PRQL——管线式关系查询语言 |
| 3050 | {{< alias "pldbgapi" "pldebugger" >}} | 1.9 | 用于调试 PL/pgSQL 函数的服务器端支持 |
| 3060 | {{< alias "plpgsql_check" >}} | 2.8.5 | 对 plpgsql 函数进行扩展检查 |
| 3070 | {{< alias "plprofiler" >}} | 4.2.5 | 剖析 PL/pgSQL 函数 |
| 3080 | {{< alias "plsh" >}} | 1.20220917 | PL/sh 程序语言 |
| 3090 | {{< alias "pljava" >}} | 1.6.10 | Java 程序语言 |
| 3100 | {{< alias "plr" >}} | 8.4.8 | 从数据库中加载R语言解释器并执行R脚本 |
| 3110 | {{< alias "plxslt" >}} | 0.20140221 | XSLT 存储过程语言 |
| 3200 | {{< alias "pgtap" >}} | 1.3.4 | PostgreSQL单元测试框架 |
| 3210 | {{< alias "faker" >}} | 0.5.3 | 插入生成的测试伪造数据，Python库的包装 |
| 3220 | {{< alias "dbt2" >}} | 0.61.7 | OSDL-DBT-2 测试组件 |
| 3240 | {{< alias "pltcl" >}} | 1.0 | PL/TCL 存储过程语言 |
| 3250 | {{< alias "pltclu" "pltcl" >}} | 1.0 | PL/TCL 存储过程语言（未受信/高权限） |
| 3260 | {{< alias "plperl" >}} | 1.0 | PL/Perl 存储过程语言 |
| 3261 | {{< alias "bool_plperl" "plperl" >}} | 1.0 | 在 bool 和 plperl 之间转换 |
| 3262 | {{< alias "hstore_plperl" "plperl" >}} | 1.0 | 在 hstore 和 plperl 之间转换适配类型 |
| 3263 | {{< alias "jsonb_plperl" "plperl" >}} | 1.0 | 在 jsonb 和 plperl 之间转换 |
| 3270 | {{< alias "plperlu" >}} | 1.0 | PL/PerlU 存储过程语言（未受信/高权限） |
| 3271 | {{< alias "bool_plperlu" "plperlu" >}} | 1.0 | 在 bool 和 plperlu 之间转换 |
| 3272 | {{< alias "jsonb_plperlu" "plperlu" >}} | 1.0 | 在 jsonb 和 plperlu 之间转换 |
| 3273 | {{< alias "hstore_plperlu" "plperlu" >}} | 1.0 | 在 hstore 和 plperlu 之间转换适配类型 |
| 3280 | {{< alias "plpgsql" >}} | 1.0 | PL/pgSQL 程序设计语言 |
| 3290 | {{< alias "plpython3u" >}} | 1.0 | PL/Python3 存储过程语言（未受信/高权限） |
| 3291 | {{< alias "jsonb_plpython3u" "plpython3u" >}} | 1.0 | 在 jsonb 和 plpython3u 之间转换 |
| 3292 | {{< alias "ltree_plpython3u" "plpython3u" >}} | 1.0 | 在 ltree 和 plpython3u 之间转换 |
| 3293 | {{< alias "hstore_plpython3u" "plpython3u" >}} | 1.0 | 在 hstore 和 plpython3u 之间转换 |

## TYPE

自定义类型扩展：前缀树，语义版本号，SI单位，位图，无符号整型，高精度数值，有理数，哈希值，IP地址段，球面，RRULE等

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 3500 | {{< alias "prefix" "pg_prefix" >}} | 1.2.10 | 前缀树数据类型 |
| 3510 | {{< alias "semver" "pg_semver" >}} | 0.41.0 | 语义版本号数据类型 |
| 3520 | {{< alias "unit" "pgunit" >}} | 7.10 | SI 国标单位扩展 |
| 3530 | {{< alias "pgpdf" >}} | 0.1.0 | PDF数据类型，管理函数与全文检索 |
| 3540 | {{< alias "pglite_fusion" >}} | 0.0.6 | 在PG表中嵌入SQLite数据库作为数据类型 |
| 3550 | {{< alias "md5hash" >}} | 1.0.1 | 提供128位MD5的原生数据类型 |
| 3560 | {{< alias "asn1oid" >}} | 1.6 | ASN1OID数据类型支持 |
| 3570 | {{< alias "roaringbitmap" "pg_roaringbitmap" >}} | 1.1.0 | 支持RoaringBitmap数据类型 |
| 3580 | {{< alias "pgfaceting" >}} | 0.2.0 | 使用倒排索引的高速切面查询 |
| 3590 | {{< alias "pg_sphere" "pgsphere" >}} | 1.5.2 | 球面对象函数、运算符与索引支持 |
| 3600 | {{< alias "country" "pg_country" >}} | 0.0.3 | 国家代码数据类型，遵循ISO 3166-1标准 |
| 3610 | {{< alias "pg_xenophile" >}} | 0.8.3 | PostgreSQL i8n与l10n工具包 |
| 3611 | {{< alias "l10n_table_dependent_extension" "pg_xenophile" >}} | 0.8.3 | PostgreSQL l10n 工具包 |
| 3620 | {{< alias "currency" "pg_currency" >}} | 0.0.3 | 使用1字节表示的货币数据类型 |
| 3630 | {{< alias "collection" "pgcollection" >}} | 1.1.0 | 在PlPGSQL中使用的内存优化高性能集合数据结构 |
| 3700 | {{< alias "pgmp" >}} | 1.0.5 | 多精度算术扩展 |
| 3710 | {{< alias "numeral" >}} | 1.3 | 数值类型扩展 |
| 3720 | {{< alias "pg_rational" >}} | 0.0.2 | 使用BIGINT表示的有理数数据类型 |
| 3730 | {{< alias "uint" "pguint" >}} | 1.20250815 | 无符号整型数据类型 |
| 3740 | {{< alias "uint128" "pg_uint128" >}} | 1.1.1 | 原生128位无符号整型数据类型 |
| 3750 | {{< alias "hashtypes" >}} | 0.1.5 | 包括SHA1，MD5在内的多种哈希数据类型 |
| 3820 | {{< alias "ip4r" >}} | 2.4.2 | PostgreSQL 的 IPv4/v6 和 IPv4/v6 范围索引类型 |
| 3830 | {{< alias "pg_duration" >}} | 1.0.2 | 用于表示时间段的强化数据类型 |
| 3840 | {{< alias "uri" "pg_uri" >}} | 1.20151224 | URI数据类型 |
| 3850 | {{< alias "emailaddr" "pg_emailaddr" >}} | 0 | Email地址数据类型 |
| 3860 | {{< alias "acl" "pg_acl" >}} | 1.0.4 | ACL数据类型 |
| 3870 | {{< alias "debversion" >}} | 1.2.0 | Debian版本号数据类型 |
| 3880 | {{< alias "pg_rrule" >}} | 0.3.0 | 日历重复规则RRULE数据类型 |
| 3890 | {{< alias "timestamp9" >}} | 1.4.0 | 纳秒分辨率时间戳 |
| 3920 | {{< alias "chkpass" >}} | 1.0 | 数据类型：自动加密的密码 |
| 3930 | {{< alias "isn" >}} | 1.2 | 用于国际产品编号标准的数据类型 |
| 3940 | {{< alias "seg" >}} | 1.4 | 表示线段或浮点间隔的数据类型 |
| 3950 | {{< alias "cube" >}} | 1.5 | 用于存储多维立方体的数据类型 |
| 3960 | {{< alias "ltree" >}} | 1.3 | 用于表示分层树状结构的数据类型 |
| 3970 | {{< alias "hstore" >}} | 1.8 | 用于存储（键，值）对集合的数据类型 |
| 3980 | {{< alias "citext" >}} | 1.6 | 提供大小写不敏感的字符串类型 |
| 3990 | {{< alias "xml2" >}} | 1.1 | XPath 查询和 XSLT |

## UTIL

实用功能扩展：HTTP请求，GZIP压缩，JWT处理，邮件客户端，正则，字符编码，编码解码，加密解密等实用功能

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 4010 | {{< alias "gzip" "pg_gzip" >}} | 1.0.0 | 使用SQL执行Gzip压缩与解压缩 |
| 4020 | {{< alias "bzip" "pg_bzip" >}} | 1.0.0 | BZIP压缩解压缩函数包 |
| 4030 | {{< alias "zstd" "pg_zstd" >}} | 1.1.2 | ZSTD压缩解压缩函数包 |
| 4070 | {{< alias "http" "pg_http" >}} | 1.7.0 | HTTP客户端，允许在数据库内收发HTTP请求 (supabase) |
| 4080 | {{< alias "pg_net" >}} | 0.20.0 | 用 SQL 进行异步非阻塞HTTP/HTTPS 请求的扩展 (supabase) |
| 4090 | {{< alias "pg_curl" >}} | 2.4.5 | 封装CURL，执行各种用URL传输数据的操作 |
| 4100 | {{< alias "pg_retry" >}} | 1.0.0 | 在临时错误中使用指数退避重试语句 |
| 4150 | {{< alias "pgjq" >}} | 0.1.0 | 在Postgres中使用jq查询JSON |
| 4160 | {{< alias "pgjwt" >}} | 0.2.0 | JSON Web Token API 的PG实现 (supabase) |
| 4170 | {{< alias "pg_smtp_client" >}} | 0.2.1 | 使用SMTP从PostgreSQL内发送邮件的客户端扩展 |
| 4180 | {{< alias "pg_html5_email_address" >}} | 1.2.3 | 验证Email是否符合HTML5规范的扩展 |
| 4190 | {{< alias "url_encode" >}} | 1.2.5 | 提供URL编码解码函数 |
| 4200 | {{< alias "pgsql_tweaks" >}} | 1.0.2 | 一些日常会用到的便利函数与视图 |
| 4220 | {{< alias "pg_extra_time" >}} | 2.0.0 | 一些关于日期与时间的扩展函数 |
| 4230 | {{< alias "pgpcre" >}} | 0.20190509 | PCRE/Perl风格的正则表达式支持 |
| 4240 | {{< alias "icu_ext" >}} | 1.10.0 | 访问ICU库提供的函数 |
| 4250 | {{< alias "pgqr" >}} | 1.0 | 从数据库中直接生成QR二维码 |
| 4260 | {{< alias "pg_protobuf" >}} | 1.0 | 提供Protobuf函数支持 |
| 4270 | {{< alias "envvar" "pg_envvar" >}} | 1.0.1 | 获取环境变量的函数 |
| 4280 | {{< alias "floatfile" >}} | 1.3.1 | 将浮点数组存储到文件中而不是堆表中 |
| 4290 | {{< alias "pg_render" >}} | 0.1.3 | 使用SQL渲染HTML页面 |
| 4300 | {{< alias "pg_readme" >}} | 0.7.0 | 为模式与扩展生成Markdown文档 |
| 4301 | {{< alias "pg_readme_test_extension" "pg_readme" >}} | 0.7.0 | 为模式与扩展生成Markdown文档 |
| 4310 | {{< alias "ddl_historization" >}} | 0.0.7 | 用SQL将所有DDL变更写入到数据库表中 |
| 4320 | {{< alias "data_historization" >}} | 1.1.0 | 用SQL将数据变更历史保存到分区表中 |
| 4330 | {{< alias "schedoc" "pg_schedoc" >}} | 0.0.1 | 在Django与DBT之间通过注释文档交换元数据 |
| 4400 | {{< alias "hashlib" "pg_hashlib" >}} | 1.1 | 稳定哈希函数包 |
| 4430 | {{< alias "xxhash" "pg_xxhash" >}} | 0.0.1 | xxhash哈希函数包 |
| 4440 | {{< alias "shacrypt" >}} | 1.1 | 实现SHA256-CRYPT与SHA512-CRYPT密钥加密算法 |
| 4450 | {{< alias "cryptint" >}} | 1.0.0 | 加密INT与BIGINT类型 |
| 4460 | {{< alias "pguecc" "pg_ecdsa" >}} | 1.0 | PostgreSQL的uECC绑定，椭圆曲线加解密函数包 |
| 4470 | {{< alias "sparql" "pgsparql" >}} | 1.0 | 使用SQL查询SPARQL数据源 |

## FUNC

标识聚合函数：ID生成器，各类聚合函数，摘要函数，数组处理函数，数学函数，统计量，伪随机，等等

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 4500 | {{< alias "pg_idkit" >}} | 0.4.0 | 生成各式各样的唯一标识符：UUIDv6, ULID, KSUID |
| 4510 | {{< alias "pgx_ulid" >}} | 0.2.2 | ULID数据类型与函数 |
| 4540 | {{< alias "pg_uuidv7" >}} | 1.7.0 | UUIDv7 支持 |
| 4550 | {{< alias "permuteseq" >}} | 1.2.2 | 伪随机数ID置换生成器 |
| 4560 | {{< alias "pg_hashids" >}} | 1.3 | 加盐将整型ID转为短字符串ID |
| 4570 | {{< alias "sequential_uuids" >}} | 1.0.3 | 生成连续生成的UUID |
| 4580 | {{< alias "typeid" "pg_typeid" >}} | 0.3.0 | PG原生TypeID类型与函数 |
| 4600 | {{< alias "topn" >}} | 2.7.0 | top-n JSONB 的类型 |
| 4610 | {{< alias "quantile" >}} | 1.1.8 | Quantile聚合函数 |
| 4620 | {{< alias "lower_quantile" >}} | 1.0.3 | Lower Quantile 聚合函数 |
| 4630 | {{< alias "count_distinct" >}} | 3.0.2 | COUNT(DISTINCT …) 聚合的替代方案 |
| 4640 | {{< alias "omnisketch" >}} | 1.0.2 | 实现OmniSketch数据结构，实现近似摘要聚合 |
| 4650 | {{< alias "ddsketch" >}} | 1.0.1 | 实现DDSketch数据结构，实现在线的Quantile聚合 |
| 4660 | {{< alias "vasco" >}} | 0.1.0 | 使用MIC发现数据中隐含的关联 |
| 4670 | {{< alias "xicor" "pgxicor" >}} | 0.1.0 | 在PG中计算XI相关系数 |
| 4680 | {{< alias "weighted_statistics" "pg_weighted_statistics" >}} | 1.0.0 | 针对稀疏数据的高性能加权统计量计算 |
| 4700 | {{< alias "tdigest" >}} | 1.4.3 | tdigest 聚合函数 |
| 4710 | {{< alias "first_last_agg" >}} | 0.1.4 | first() 与 last() 聚合函数 |
| 4720 | {{< alias "extra_window_functions" >}} | 1.0 | 额外的窗口函数 |
| 4730 | {{< alias "floatvec" >}} | 1.1.1 | 数组类型数学运算扩展 |
| 4740 | {{< alias "aggs_for_vecs" >}} | 1.4.0 | 针对数组类型的聚合函数集合扩展 |
| 4750 | {{< alias "aggs_for_arrays" >}} | 1.3.3 | 计算数组聚合统计值的函数包 |
| 4760 | {{< alias "pg_csv" >}} | 1.0.1 | 灵活的CSV聚合处理函数 |
| 4770 | {{< alias "arraymath" "pg_arraymath" >}} | 1.1 | 数组逐元素数学运算符包 |
| 4780 | {{< alias "pg_math" >}} | 1.0 | 使用GSL库的数学统计函数 |
| 4790 | {{< alias "random" "pg_random" >}} | 2.0.0 | 随机数生成器 |
| 4800 | {{< alias "base36" "pg_base36" >}} | 1.0.0 | Base36编码解码扩展 |
| 4810 | {{< alias "base62" "pg_base62" >}} | 0.0.1 | Base62编码解码扩展 |
| 4830 | {{< alias "pg_base58" >}} | 0.0.1 | Base58 编码/解码函数 |
| 4840 | {{< alias "financial" "pg_financial" >}} | 1.0.1 | 金融领域聚合函数 |
| 4850 | {{< alias "convert" "pg_convert" >}} | 0.1.0 | 用于空间里程等的公英制转换函数 |
| 4880 | {{< alias "refint" >}} | 1.0 | 实现引用完整性的函数 |
| 4881 | {{< alias "autoinc" >}} | 1.0 | 用于自动递增字段的函数 |
| 4882 | {{< alias "insert_username" >}} | 1.0 | 用于跟踪谁更改了表的函数 |
| 4883 | {{< alias "moddatetime" >}} | 1.0 | 跟踪最后修改时间 |
| 4890 | {{< alias "tsm_system_time" >}} | 1.0 | 接受毫秒数限制的 TABLESAMPLE 方法 |
| 4900 | {{< alias "dict_xsyn" >}} | 1.0 | 用于扩展同义词处理的文本搜索字典模板 |
| 4910 | {{< alias "tsm_system_rows" >}} | 1.0 | 接受行数限制的 TABLESAMPLE 方法 |
| 4920 | {{< alias "tcn" >}} | 1.0 | 用触发器通知变更 |
| 4930 | {{< alias "uuid-ossp" >}} | 1.1 | 生成通用唯一标识符（UUIDs） |
| 4940 | {{< alias "btree_gist" >}} | 1.7 | 用GiST索引常见数据类型 |
| 4950 | {{< alias "btree_gin" >}} | 1.3 | 用GIN索引常见数据类型 |
| 4960 | {{< alias "intarray" >}} | 1.5 | 1维整数数组的额外函数、运算符和索引支持 |
| 4970 | {{< alias "intagg" >}} | 1.1 | 整数聚合器和枚举器（过时） |
| 4980 | {{< alias "dict_int" >}} | 1.0 | 用于整数的文本搜索字典模板 |
| 4990 | {{< alias "unaccent" >}} | 1.1 | 删除重音的文本搜索字典 |

## ADMIN

管理工具扩展：膨胀治理，脏读，检视缓冲区，数据目录，校验和，腐败检查，优先级管理，权限管理，语句准备，限制批量更新等

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 5010 | {{< alias "pg_repack" >}} | 1.5.3 | 在线垃圾清理与表膨胀治理 |
| 5020 | {{< alias "pg_rewrite" >}} | 2.0.0 | 在线重写整表，不阻塞读写 |
| 5040 | {{< alias "pg_squeeze" >}} | 1.9.1 | 从关系中删除未使用空间 |
| 5050 | {{< alias "pg_dirtyread" >}} | 2.7 | 从表中读取尚未垃圾回收的行 |
| 5060 | {{< alias "pgfincore" >}} | 1.3.1 | 检查和管理操作系统缓冲区缓存 |
| 5070 | {{< alias "pg_cooldown" >}} | 0.1 | 从缓冲区中移除特定关系的页面 |
| 5080 | {{< alias "ddlx" "pg_ddlx" >}} | 0.30 | 提取数据库对象的DDL |
| 5090 | {{< alias "pglinter" >}} | 1.0.1 | PG数据库规则检查插件 |
| 5100 | {{< alias "prioritize" "pg_prioritize" >}} | 1.0.4 | 获取和设置 PostgreSQL 后端的优先级 |
| 5110 | {{< alias "pg_checksums" >}} | 1.3 | 在离线模式下激活/启用/禁用数据库集群的校验和功能 |
| 5120 | {{< alias "pg_readonly" >}} | 1.0.3 | 将集群设置为只读 |
| 5130 | {{< alias "pgdd" >}} | 0.6.1 | 提供通过标准SQL查询数据库目录集簇的能力 |
| 5140 | {{< alias "pg_permissions" >}} | 1.4 | 查看对象权限并将其与期望状态进行比较 |
| 5150 | {{< alias "pgautofailover" >}} | 2.2 | PG 自动故障迁移 |
| 5160 | {{< alias "pg_catcheck" >}} | 1.6.0 | 用于诊断系统目录是否损坏的工具 |
| 5170 | {{< alias "pre_prepare" "preprepare" >}} | 0.9 | 在服务端预先准备好PreparedStatement备用 |
| 5180 | {{< alias "pg_upless" >}} | 0.0.3 | 检测表上的无用UPDATE |
| 5190 | {{< alias "pgcozy" >}} | 1.0 | 根据先前的pg_buffercache快照预热内存缓冲区 |
| 5200 | {{< alias "pg_orphaned" >}} | 1.0 | 处理孤儿文件的扩展插件 |
| 5210 | {{< alias "pg_crash" >}} | 1.0 | 向数据库进程随机发送信号模拟故障 |
| 5220 | {{< alias "pg_cheat_funcs" >}} | 1.0 | 一些超级实用的作弊函数 |
| 5230 | {{< alias "fio" "pg_fio" >}} | 1.0 | PostgreSQL文件IO函数包 |
| 5810 | {{< alias "pg_savior" >}} | 0.0.1 | 阻止不带条件的全表更新以避免意外事故 |
| 5820 | {{< alias "safeupdate" >}} | 1.5 | 强制在 UPDATE 和 DELETE 时提供 Where 条件 |
| 5830 | {{< alias "pg_drop_events" >}} | 0.1.0 | 记录删表删列删视图的事务号，辅助PITR确定时间点 |
| 5840 | {{< alias "table_log" >}} | 0.6.4 | 记录某张表的修改日志并做表/行级时间点恢复 |
| 5880 | {{< alias "pgagent" >}} | 4.2.3 | PostgreSQL任务调度工具，与PGADMIN配合使用 |
| 5890 | {{< alias "pg_prewarm" >}} | 1.2 | 预热关系数据 |
| 5900 | {{< alias "pgpool_adm" "pgpool" >}} | 4.7.0 | PGPool 管理函数 |
| 5910 | {{< alias "pgpool_recovery" "pgpool" >}} | 4.7.0 | PGPool辅助扩展，从v4.3提供的恢复函数 |
| 5920 | {{< alias "pgpool_regclass" "pgpool" >}} | 4.7.0 | PGPool辅助扩展，RegClass替代 |
| 5930 | {{< alias "lo" >}} | 1.1 | 大对象维护 |
| 5940 | {{< alias "basic_archive" >}} | - | 归档模块样例 |
| 5950 | {{< alias "basebackup_to_shell" >}} | - | 添加一种备份到Shell终端到基础备份方式 |
| 5960 | {{< alias "old_snapshot" >}} | 1.0 | 支持 old_snapshot_threshold 的实用程序 |
| 5970 | {{< alias "adminpack" >}} | 2.1 | PostgreSQL 管理函数集合 |
| 5980 | {{< alias "amcheck" >}} | 1.4 | 校验关系完整性 |
| 5990 | {{< alias "pg_surgery" >}} | 1.0 | 对损坏的关系进行手术 |

## STAT

监控统计扩展：AWR报告，可观测性指标，显示执行计划，查询统计信息，内存使用，配置变更，等待事件采样，慢查询日志，等等

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 6000 | {{< alias "pg_profile" >}} | 4.11 | PostgreSQL 数据库负载记录与AWR报表工具 |
| 6010 | {{< alias "pg_tracing" >}} | 0.1.3 | PostgreSQL分布式Tracing |
| 6210 | {{< alias "pg_show_plans" >}} | 2.1.7 | 打印所有当前正在运行查询的执行计划 |
| 6220 | {{< alias "pg_stat_kcache" >}} | 2.3.1 | 内核统计信息收集 |
| 6230 | {{< alias "pg_stat_monitor" >}} | 2.3.1 | 提供查询聚合统计、客户端信息、执行计划详细信息和直方图 |
| 6240 | {{< alias "pg_qualstats" >}} | 2.1.3 | 收集有关 quals 的统计信息的扩展 |
| 6250 | {{< alias "pg_store_plans" >}} | 1.9 | 跟踪所有执行的 SQL 语句的计划统计信息 |
| 6260 | {{< alias "pg_track_settings" >}} | 2.1.2 | 跟踪设置更改 |
| 6270 | {{< alias "pg_wait_sampling" >}} | 1.1.9 | 基于采样的等待事件统计 |
| 6280 | {{< alias "pgsentinel" >}} | 1.3.1 | 活跃会话历史 |
| 6290 | {{< alias "system_stats" >}} | 3.2 | PostgreSQL 的系统统计函数 |
| 6300 | {{< alias "meta" "pg_meta" >}} | 0.4.0 | 标准化，更友好的PostgreSQL系统目录视图 |
| 6310 | {{< alias "pgnodemx" >}} | 1.7 | 使用SQL查询获取操作系统指标 |
| 6320 | {{< alias "pg_proctab" "pgnodemx" >}} | 1.7 | 通过SQL接口访问操作系统进程表 |
| 6330 | {{< alias "pg_sqlog" >}} | 1.6 | 提供访问PostgreSQL日志的SQL接口 |
| 6340 | {{< alias "bgw_replstatus" >}} | 1.0.8 | 用于汇报本机主从状态的后台工作进程 |
| 6350 | {{< alias "pgmeminfo" >}} | 1.0.0 | 显示内存使用情况 |
| 6360 | {{< alias "toastinfo" >}} | 1.5 | 显示TOAST字段的详细信息 |
| 6370 | {{< alias "explain_ui" "pg_explain_ui" >}} | 0.0.2 | 快速跳转至PEV查阅可视化执行计划 |
| 6380 | {{< alias "pg_relusage" >}} | 0.0.1 | 打印查询引用的表与列 |
| 6800 | {{< alias "pagevis" >}} | 0.1 | 使用ASCII字符可视化数据库物理页面布局 |
| 6810 | {{< alias "powa" >}} | 5.1.1 | PostgreSQL 工作负载分析器-核心 |
| 6880 | {{< alias "pg_overexplain" >}} | 1.0 | 允许 EXPLAIN 转储更多详细 |
| 6890 | {{< alias "pg_logicalinspect" >}} | 1.0 | 检视逻辑解码组件详情 |
| 6900 | {{< alias "pageinspect" >}} | 1.12 | 检查数据库页面二进制内容 |
| 6910 | {{< alias "pgrowlocks" >}} | 1.2 | 显示行级锁信息 |
| 6920 | {{< alias "sslinfo" >}} | 1.2 | 关于 SSL 证书的信息 |
| 6930 | {{< alias "pg_buffercache" >}} | 1.5 | 检查共享缓冲区缓存 |
| 6940 | {{< alias "pg_walinspect" >}} | 1.1 | 用于检查 PostgreSQL WAL 日志内容的函数 |
| 6950 | {{< alias "pg_freespacemap" >}} | 1.2 | 检查自由空间映射的内容（FSM） |
| 6960 | {{< alias "pg_visibility" >}} | 1.2 | 检查可见性图（VM）和页面级可见性信息 |
| 6970 | {{< alias "pgstattuple" >}} | 1.5 | 显示元组级统计信息 |
| 6980 | {{< alias "auto_explain" >}} | - | 提供一种自动记录执行计划的手段 |
| 6990 | {{< alias "pg_stat_statements" >}} | 1.11 | 跟踪所有执行的 SQL 语句的计划和执行统计信息 |

## SEC

安全功能扩展：强制密码强度，阉割超级用户，密钥管理，商密算法，PII匿名处理，扩展白名单，审计日志，变更追溯，反病毒等等

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 7000 | {{< alias "passwordcheck_cracklib" >}} | 3.1.0 | 使用cracklib加固PG用户密码 |
| 7010 | {{< alias "supautils" >}} | 3.0.2 | 用于在云环境中确保数据库集群的安全 |
| 7020 | {{< alias "pgsodium" >}} | 3.1.9 | 表数据加密存储 TDE |
| 7030 | {{< alias "supabase_vault" "pg_vault" >}} | 0.3.1 | 在 Vault 中存储加密凭证的扩展 (supabase) |
| 7040 | {{< alias "pg_session_jwt" >}} | 0.4.0 | 使用JWT进行会话认证 |
| 7050 | {{< alias "anon" "pg_anon" >}} | 2.5.1 | 数据匿名化处理工具 |
| 7060 | {{< alias "pgsmcrypto" >}} | 0.1.1 | 为PostgreSQL提供商密算法支持：SM2,SM3,SM4 |
| 7070 | {{< alias "pg_enigma" >}} | 0.5.0 | PostgreSQL 加密数据类型 |
| 7100 | {{< alias "pgaudit" >}} | 18.0 | 提供审计功能 |
| 7120 | {{< alias "pgauditlogtofile" >}} | 1.7.6 | pgAudit 子扩展，将审计日志写入单独的文件中 |
| 7130 | {{< alias "pg_auditor" >}} | 0.2 | 审计数据变更并提供闪回能力 |
| 7140 | {{< alias "logerrors" >}} | 2.1.5 | 用于收集日志文件中消息统计信息的函数 |
| 7150 | {{< alias "pg_auth_mon" >}} | 3.0 | 监控每个用户的连接尝试 |
| 7160 | {{< alias "pg_jobmon" >}} | 1.4.1 | 记录和监控函数 |
| 7310 | {{< alias "credcheck" >}} | 4.4 | 明文凭证检查器 |
| 7320 | {{< alias "pgcryptokey" >}} | 0.85 | PG密钥管理 |
| 7360 | {{< alias "login_hook" >}} | 1.7 | 在用户登陆时执行login_hook.login()函数 |
| 7370 | {{< alias "set_user" >}} | 4.2.0 | 增加了日志记录的 SET ROLE |
| 7380 | {{< alias "pg_snakeoil" >}} | 1.4 | PostgreSQL动态链接库反病毒功能 |
| 7390 | {{< alias "pgextwlist" >}} | 1.19 | PostgreSQL扩展白名单功能 |
| 7410 | {{< alias "sslutils" >}} | 1.4 | 使用SQL管理SSL证书 |
| 7420 | {{< alias "noset" "pg_noset" >}} | 0.3.0 | 阻止非超级用户使用SET/RESET设置变量 |
| 7500 | {{< alias "pg_tde" >}} | 1.0 | Percona加密存储引擎 |
| 7960 | {{< alias "sepgsql" >}} | - | 基于SELinux标签的强制访问控制 |
| 7970 | {{< alias "auth_delay" >}} | - | 在返回认证失败前暂停一会，避免爆破 |
| 7980 | {{< alias "pgcrypto" >}} | 1.3 | 实用加解密函数 |
| 7990 | {{< alias "passwordcheck" >}} | - | 用于强制拒绝修改弱密码的扩展 |

## FDW

外部数据源包装器：FDW开发框架 Wrappers，Multicorn，访问外部的 Mongo，MySQL，SQLite，HDFS，MSSQL，Oracle，DB2，……

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 8500 | {{< alias "wrappers" >}} | 0.5.7 | Supabase提供的外部数据源包装器捆绑包 |
| 8510 | {{< alias "multicorn" >}} | 3.2 | 用Python编写自定义的外部数据源包装器 |
| 8520 | {{< alias "odbc_fdw" >}} | 0.5.1 | 访问ODBC可访问的任何外部数据源 |
| 8530 | {{< alias "jdbc_fdw" >}} | 0.4.0 | 访问JDBC可访问的任何外部数据源 |
| 8540 | {{< alias "pgspider_ext" >}} | 1.3.0 | 使用多种FDW访问远程数据库服务器 |
| 8600 | {{< alias "mysql_fdw" >}} | 2.9.3 | MySQL外部数据包装器 |
| 8610 | {{< alias "oracle_fdw" >}} | 2.8.0 | 提供对Oracle的外部数据源包装器 |
| 8620 | {{< alias "tds_fdw" >}} | 2.0.5 | TDS 数据库（Sybase/SQL Server）外部数据包装器 |
| 8630 | {{< alias "db2_fdw" >}} | 18.0.1 | 提供对DB2的外部数据源包装器 |
| 8640 | {{< alias "sqlite_fdw" >}} | 2.5.0 | SQLite 外部数据包装器 |
| 8650 | {{< alias "pgbouncer_fdw" >}} | 1.4.0 | 用SQL查询pgbouncer统计信息，并执行pgbouncer命令 |
| 8660 | {{< alias "etcd_fdw" >}} | 0.0.0 | etcd分布式键值存储外部数据包装器 |
| 8700 | {{< alias "mongo_fdw" >}} | 5.5.3 | MongoDB 外部数据包装器 |
| 8710 | {{< alias "redis_fdw" >}} | 1.0 | 查询外部Redis数据源 |
| 8720 | {{< alias "redis" "pg_redis_pubsub" >}} | 0.0.1 | 从PG向Redis发送Pub/Sub消息 |
| 8730 | {{< alias "kafka_fdw" >}} | 0.0.3 | Kafka外部数据源包装器 |
| 8740 | {{< alias "hdfs_fdw" >}} | 2.3.3 | hdfs 外部数据包装器 |
| 8750 | {{< alias "firebird_fdw" >}} | 1.4.1 | Firebird外部数据源包装器 |
| 8800 | {{< alias "aws_s3" >}} | 0.0.1 | 从S3导入导出数据的外部数据源包装器 |
| 8810 | {{< alias "log_fdw" >}} | 1.4 | 访问PostgreSQL日志文件的FDW |
| 8970 | {{< alias "dblink" >}} | 1.2 | 从数据库内连接到其他 PostgreSQL 数据库 |
| 8980 | {{< alias "file_fdw" >}} | 1.0 | 访问外部文件的外部数据包装器 |
| 8990 | {{< alias "postgres_fdw" >}} | 1.1 | 用于远程 PostgreSQL 服务器的外部数据包装器 |

## SIM

数据库兼容扩展：仿真其他 DBMS 的行为：MySQL，Memcache，Mongo，Oracle，Babelfish for Microsoft SQL Server……

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 9000 | {{< alias "documentdb" >}} | 0.109 | 微软DocumentDB的API层 |
| 9010 | {{< alias "documentdb_core" "documentdb" >}} | 0.109 | 微软DocumentDB的核心API层实现 |
| 9020 | {{< alias "documentdb_distributed" "documentdb" >}} | 0.109 | DocumentDB多节点模式的API层 |
| 9030 | {{< alias "documentdb_extended_rum" "documentdb" >}} | 0.109 | DocumentDB扩展RUM索引访问方法 |
| 9100 | {{< alias "orafce" >}} | 4.16.3 | 模拟 Oracle RDBMS 的一部分函数和包的函数和运算符 |
| 9110 | {{< alias "pgtt" >}} | 4.4 | 类似Oracle的全局临时表功能 |
| 9120 | {{< alias "session_variable" >}} | 3.4 | Oracle兼容的会话变量/常量操作函数 |
| 9130 | {{< alias "pg_statement_rollback" >}} | 1.5 | 在服务端提供类似Oracle/DB2的语句级回滚能力 |
| 9240 | {{< alias "pg_dbms_metadata" >}} | 1.0.0 | 添加 Oracle DBMS_METADATA 兼容性支持的扩展 |
| 9250 | {{< alias "pg_dbms_lock" >}} | 1.0 | 为PG添加对 Oracle DBMS_LOCK 的完整兼容性支持 |
| 9260 | {{< alias "pg_dbms_job" >}} | 1.5 | 添加 Oracle DBMS_JOB 兼容性支持的扩展 |
| 9270 | {{< alias "pg_dbms_errlog" >}} | 2.2 | 模仿 Oracle DBMS_ERRLOG 模块来记录特定表的DML错误 |
| 9300 | {{< alias "babelfishpg_common" >}} | 3.3.3 | SQL Server 数据类型兼容扩展 |
| 9310 | {{< alias "babelfishpg_tsql" >}} | 3.3.1 | SQL Server SQL语法兼容性扩展 |
| 9320 | {{< alias "babelfishpg_tds" >}} | 1.0.0 | SQL Server TDS线缆协议兼容扩展 |
| 9330 | {{< alias "babelfishpg_money" >}} | 1.1.0 | SQL Server 货币数据类型兼容扩展 |
| 9400 | {{< alias "spat" >}} | 0.1.0a4 | 在PG中嵌入Redis风格的内存数据库 |
| 9410 | {{< alias "pgmemcache" >}} | 2.3.0 | 为PG提供memcached兼容接口 |

## ETL

数据复制扩展：逻辑复制，逻辑解码，DDL复制，JSON/BSON/Protobuf 变更抽取，数据迁移，数据导入，数据比对等

| ID | 扩展/包 | 版本 | 描述 |
|:---:|:---|:---|:---|
| 9500 | {{< alias "pglogical" >}} | 2.4.6 | PostgreSQL逻辑复制：三方扩展实现 |
| 9501 | {{< alias "pglogical_origin" "pglogical" >}} | 2.4.6 | 用于从 Postgres 9.4 升级时的兼容性虚拟扩展 |
| 9510 | {{< alias "pglogical_ticker" >}} | 1.4.1 | pglogical复制延迟以秒计的精确视图 |
| 9520 | {{< alias "pgl_ddl_deploy" >}} | 2.2.1 | 使用 pglogical 执行自动 DDL 部署 |
| 9530 | {{< alias "pg_failover_slots" >}} | 1.2.0 | 在Failover过程中保留复制槽 |
| 9540 | {{< alias "db_migrator" >}} | 1.0.0 | 使用FDW从其他DBMS迁移到PostgreSQL |
| 9550 | {{< alias "pgactive" >}} | 2.1.7 | PostgreSQL多主逻辑复制 |
| 9630 | {{< alias "wal2json" >}} | 2.6 | 用逻辑解码捕获 JSON 格式的 CDC 变更 |
| 9640 | {{< alias "wal2mongo" >}} | 1.0.7 | 使用逻辑解码捕获MongoDB JSON格式的CDC变更 |
| 9650 | {{< alias "decoderbufs" >}} | 3.4.0 | 将WAL逻辑解码为ProtocolBuffer协议的消息 |
| 9660 | {{< alias "decoder_raw" >}} | 1.0 | 逻辑复制解码输出插件：RAW SQL格式 |
| 9700 | {{< alias "mimeo" >}} | 1.5.1 | 在PostgreSQL实例间进行表级复制 |
| 9710 | {{< alias "repmgr" >}} | 5.5.0 | PostgreSQL复制管理组件 |
| 9820 | {{< alias "pg_fact_loader" >}} | 2.0.1 | 在 Postgres 中构建事实表 |
| 9830 | {{< alias "pg_bulkload" >}} | 3.1.23 | 向 PostgreSQL 中高速加载数据 |
| 9970 | {{< alias "test_decoding" >}} | - | 基于SQL的WAL逻辑解码样例 |
| 9980 | {{< alias "pgoutput" >}} | - | PG内置的逻辑解码输出插件 |

