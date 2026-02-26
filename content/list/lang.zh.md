---
title: "按语言"
description: "按实现语言组织的 PostgreSQL 扩展"
excludeSearch: true
weight: 200
---



| {{< language "c" >}}       | {{< language "c++" >}}       | {{< language "rust" >}}      | {{< language "java" >}}        | {{< language "python" >}}      | {{< language "sql" >}}         | {{< language "data" >}} |
|----------------------------|------------------------------|------------------------------|--------------------------------|--------------------------------|--------------------------------|-------------------------|

## 概览

| 语言 | 数量 | 描述 |
|:-------:|:-----:|:--------------|
| {{< language "C" >}} | 355 | 传统的 PostgreSQL 扩展开发语言 |
| {{< language "SQL" >}} | 40 | 纯 SQL 扩展和函数 |
| {{< language "Rust" >}} | 39 | 使用 pgrx 框架用 Rust 编写的扩展 |
| {{< language "Data" >}} | 10 | 仅包含数据的扩展 |
| {{< language "C++" >}} | 7 | 使用 C++ 特性和库的扩展 |
| {{< language "Python" >}} | 2 | 使用 Python 编写的扩展 |
| {{< language "Java" >}} | 1 | 在 JVM 上运行的扩展 |


## C

{{< language "C" >}} {{< badge content="355 个扩展" color="gray" icon="cube" >}}

传统的 PostgreSQL 扩展开发语言

| ID | 扩展 | 描述 |
|:---:|:---|:---|
| 1000 | {{< alias "timescaledb" >}} | 时序数据库扩展插件 |
| 1030 | {{< alias "periods" >}} | 为 PERIODs 和 SYSTEM VERSIONING 提供标准 SQL 功能 |
| 1040 | {{< alias "temporal_tables" >}} | 时态表功能支持 |
| 1070 | {{< alias "pg_cron" >}} | 定时任务调度器 |
| 1080 | {{< alias "pg_task" >}} | 在特定时间点在后台执行SQL命令 |
| 1100 | {{< alias "pg_background" >}} | 在后台运行 SQL 查询 |
| 1500 | {{< alias "postgis" >}} | PostGIS 几何和地理空间扩展 |
| 1501 | {{< alias "postgis_topology" "postgis" >}} | PostGIS 拓扑空间类型和函数 |
| 1502 | {{< alias "postgis_raster" "postgis" >}} | PostGIS 光栅类型和函数 |
| 1503 | {{< alias "postgis_sfcgal" "postgis" >}} | PostGIS SFCGAL 函数 |
| 1504 | {{< alias "postgis_tiger_geocoder" "postgis" >}} | PostGIS tiger 地理编码器和反向地理编码器 |
| 1505 | {{< alias "address_standardizer" "postgis" >}} | 地址标准化函数。 |
| 1506 | {{< alias "address_standardizer_data_us" "postgis" >}} | 地址标准化函数：美国数据集示例 |
| 1520 | {{< alias "pointcloud" >}} | 提供激光雷达点云数据类型支持 |
| 1521 | {{< alias "pointcloud_postgis" "pointcloud" >}} | 将激光雷达点云与PostGIS几何类型相集成 |
| 1530 | {{< alias "h3" "pg_h3" >}} | H3六边形层级索引支持 |
| 1531 | {{< alias "h3_postgis" "pg_h3" >}} | H3与PostGIS集成的扩展插件 |
| 1540 | {{< alias "q3c" >}} | Q3C天空索引插件 |
| 1550 | {{< alias "ogr_fdw" >}} | GIS 数据外部数据源包装器 |
| 1590 | {{< alias "pg_geohash" >}} | 使用GeoHash处理空间坐标的函数包 |
| 1650 | {{< alias "mobilitydb" >}} | MobilityDB地理空间投影数据管理分析平台 |
| 1690 | {{< alias "earthdistance" >}} | 计算地球表面上的大圆距离 |
| 1800 | {{< alias "vector" "pgvector" >}} | 向量数据类型和 ivfflat / hnsw 访问方法 |
| 1840 | {{< alias "pg_similarity" >}} | 提供17种距离度量函数 |
| 1850 | {{< alias "smlar" >}} | 高效的相似度搜索函数 |
| 1880 | {{< alias "pg4ml" >}} | PG4ML是一个机器学习框架 |
| 2110 | {{< alias "pgroonga" >}} | 使用Groonga，面向所有语言的高速全文检索平台 |
| 2111 | {{< alias "pgroonga_database" "pgroonga" >}} | PGGroonga 数据库管理模块 |
| 2120 | {{< alias "pg_bigm" >}} | 基于二字组的多语言全文检索扩展 |
| 2130 | {{< alias "zhparser" >}} | 中文分词，全文搜索解析器 |
| 2170 | {{< alias "biscuit" "pg_biscuit" >}} | 使用IAM的高性能文本模式匹配 |
| 2180 | {{< alias "pg_textsearch" >}} | 带有BM25排序的全文搜索扩展 |
| 2380 | {{< alias "fuzzystrmatch" >}} | 确定字符串之间的相似性和距离 |
| 2390 | {{< alias "pg_trgm" >}} | 文本相似度测量函数与模糊检索 |
| 2400 | {{< alias "citus" >}} | Citus 分布式数据库 |
| 2401 | {{< alias "citus_columnar" "citus" >}} | Citus 列式存储引擎 |
| 2410 | {{< alias "columnar" "hydra" >}} | 开源列式存储扩展 |
| 2470 | {{< alias "duckdb_fdw" >}} | DuckDB 外部数据源包装器 |
| 2510 | {{< alias "pg_partman" >}} | 用于按时间或 ID 管理分区表的扩展 |
| 2520 | {{< alias "plproxy" >}} | 作为过程语言实现的数据库分区 |
| 2530 | {{< alias "pg_strom" >}} | 使用GPU与NVMe加速大数据处理 |
| 2590 | {{< alias "tablefunc" >}} | 交叉表函数 |
| 2700 | {{< alias "age" >}} | Apache AGE，图数据库扩展 （Deb可用） |
| 2720 | {{< alias "rum" >}} | RUM 索引访问方法 |
| 2740 | {{< alias "pg_ttl_index" >}} | 基于TTL索引的自动数据过期清理 |
| 2770 | {{< alias "jsquery" >}} | 用于内省 JSONB 数据类型的查询类型 |
| 2780 | {{< alias "pg_hint_plan" >}} | 添加强制指定执行计划的能力 |
| 2790 | {{< alias "hypopg" >}} | 假设索引，用于创建一个虚拟索引检验执行计划 |
| 2810 | {{< alias "plan_filter" "pg_plan_filter" >}} | 使用执行计划代价过滤阻止特定查询语句 |
| 2830 | {{< alias "imgsmlr" >}} | 使用Haar小波分析计算图片相似度 |
| 2840 | {{< alias "pg_ivm" >}} | 增量维护的物化视图 |
| 2850 | {{< alias "pg_incremental" >}} | 增量处理流式事件 |
| 2890 | {{< alias "pgq" >}} | 通用队列的PG实现 |
| 2910 | {{< alias "orioledb" >}} | OrioleDB，下一代事务处理引擎 |
| 2940 | {{< alias "omni" "omnigres" >}} | PostgreSQL即平台，Omnigres主扩展与加载器 |
| 2941 | {{< alias "omni_auth" "omnigres" >}} | Omnigres 基础会话认证管理模块 |
| 2942 | {{< alias "omni_aws" "omnigres" >}} | Omnigres AWS S3 API封装 |
| 2943 | {{< alias "omni_cloudevents" "omnigres" >}} | Omnigres CloudEvents 支持 |
| 2944 | {{< alias "omni_containers" "omnigres" >}} | Omnigres Docker容器管理模块 |
| 2945 | {{< alias "omni_credentials" "omnigres" >}} | Omnigres 应用密钥管理模块 |
| 2946 | {{< alias "omni_csv" >}} | Omnigres CSV 工具箱 |
| 2947 | {{< alias "omni_datasets" >}} | Omnigres 数据库置备工具 |
| 2948 | {{< alias "omni_email" "omnigres" >}} | Omnigres Email 框架 |
| 2949 | {{< alias "omni_http" "omnigres" >}} | Omnigres 基本HTTP类型 |
| 2950 | {{< alias "omni_httpc" "omnigres" >}} | Omnigres HTTP客户端 |
| 2951 | {{< alias "omni_httpd" "omnigres" >}} | Omnigres HTTP服务器 |
| 2952 | {{< alias "omni_id" "omnigres" >}} | Omnigres ID身份数据类型 |
| 2953 | {{< alias "omni_json" "omnigres" >}} | Omnigres JSON工具箱 |
| 2954 | {{< alias "omni_kube" "omnigres" >}} | Omnigres Kubernetes集成模块 |
| 2955 | {{< alias "omni_ledger" "omnigres" >}} | Omnigres 金融账本模块 |
| 2956 | {{< alias "omni_manifest" "omnigres" >}} | Omnigres 包管理清单模块 |
| 2957 | {{< alias "omni_mimetypes" "omnigres" >}} | Omnigres MIME数据类型 |
| 2958 | {{< alias "omni_os" "omnigres" >}} | Omnigres 操作系统集成模块 |
| 2959 | {{< alias "omni_polyfill" "omnigres" >}} | Omnigres Postgres多态API |
| 2960 | {{< alias "omni_python" "omnigres" >}} | Omnigres 第一类Python支持模块 |
| 2961 | {{< alias "omni_regex" "omnigres" >}} | Omnigres PCRE兼容正则表达式模块 |
| 2962 | {{< alias "omni_rest" "omnigres" >}} | Omnigres REST API 工具包 |
| 2963 | {{< alias "omni_schema" "omnigres" >}} | Omnigres 高级模式管理组件 |
| 2964 | {{< alias "omni_seq" "omnigres" >}} | Omnigres 分布式整型序列号 |
| 2965 | {{< alias "omni_service" "omnigres" >}} | Omnigres 服务管理器 |
| 2966 | {{< alias "omni_session" "omnigres" >}} | Omnigres 会话管理器 |
| 2967 | {{< alias "omni_shmem" >}} | Omnigres 共享内存管理 |
| 2968 | {{< alias "omni_sql" "omnigres" >}} | Omnigres SQL编程组件 |
| 2969 | {{< alias "omni_sqlite" "omnigres" >}} | Omnigres 嵌入的SQLite支持 |
| 2970 | {{< alias "omni_test" "omnigres" >}} | Omnigres 测试框架 |
| 2971 | {{< alias "omni_txn" "omnigres" >}} | Omnigres 事务管理器模块 |
| 2972 | {{< alias "omni_types" "omnigres" >}} | Omnigres 高级数据类型模块 |
| 2973 | {{< alias "omni_var" "omnigres" >}} | Omnigres 局部变量模块 |
| 2974 | {{< alias "omni_vfs" "omnigres" >}} | Omnigres 虚拟文件系统 |
| 2975 | {{< alias "omni_vfs_types_v1" "omnigres" >}} | Omnigres 虚拟文件系统（v1） |
| 2976 | {{< alias "omni_web" "omnigres" >}} | Omnigres Web工具箱 |
| 2977 | {{< alias "omni_worker" "omnigres" >}} | Omnigres 通用Worker池 |
| 2978 | {{< alias "omni_xml" "omnigres" >}} | Omnigres XML工具包 |
| 2979 | {{< alias "omni_yaml" "omnigres" >}} | Omnigres YAML工具包 |
| 2990 | {{< alias "bloom" >}} | bloom 索引-基于指纹的索引 |
| 3000 | {{< alias "pg_tle" >}} | AWS 可信语言扩展 |
| 3011 | {{< alias "pljs" >}} | PL/JS 可信过程程序语言 |
| 3020 | {{< alias "pllua" >}} | Lua 程序语言 |
| 3021 | {{< alias "hstore_pllua" "pllua" >}} | Lua 程序语言的Hstore适配扩展 |
| 3030 | {{< alias "plluau" "pllua" >}} | Lua 程序语言（不受信任的） |
| 3031 | {{< alias "hstore_plluau" "pllua" >}} | Lua 程序语言的Hstore适配扩展（不受信任的） |
| 3050 | {{< alias "pldbgapi" "pldebugger" >}} | 用于调试 PL/pgSQL 函数的服务器端支持 |
| 3060 | {{< alias "plpgsql_check" >}} | 对 plpgsql 函数进行扩展检查 |
| 3070 | {{< alias "plprofiler" >}} | 剖析 PL/pgSQL 函数 |
| 3080 | {{< alias "plsh" >}} | PL/sh 程序语言 |
| 3100 | {{< alias "plr" >}} | 从数据库中加载R语言解释器并执行R脚本 |
| 3110 | {{< alias "plxslt" >}} | XSLT 存储过程语言 |
| 3200 | {{< alias "pgtap" >}} | PostgreSQL单元测试框架 |
| 3220 | {{< alias "dbt2" >}} | OSDL-DBT-2 测试组件 |
| 3240 | {{< alias "pltcl" >}} | PL/TCL 存储过程语言 |
| 3250 | {{< alias "pltclu" "pltcl" >}} | PL/TCL 存储过程语言（未受信/高权限） |
| 3260 | {{< alias "plperl" >}} | PL/Perl 存储过程语言 |
| 3261 | {{< alias "bool_plperl" "plperl" >}} | 在 bool 和 plperl 之间转换 |
| 3262 | {{< alias "hstore_plperl" "plperl" >}} | 在 hstore 和 plperl 之间转换适配类型 |
| 3263 | {{< alias "jsonb_plperl" "plperl" >}} | 在 jsonb 和 plperl 之间转换 |
| 3270 | {{< alias "plperlu" >}} | PL/PerlU 存储过程语言（未受信/高权限） |
| 3271 | {{< alias "bool_plperlu" "plperlu" >}} | 在 bool 和 plperlu 之间转换 |
| 3272 | {{< alias "jsonb_plperlu" "plperlu" >}} | 在 jsonb 和 plperlu 之间转换 |
| 3273 | {{< alias "hstore_plperlu" "plperlu" >}} | 在 hstore 和 plperlu 之间转换适配类型 |
| 3280 | {{< alias "plpgsql" >}} | PL/pgSQL 程序设计语言 |
| 3290 | {{< alias "plpython3u" >}} | PL/Python3 存储过程语言（未受信/高权限） |
| 3291 | {{< alias "jsonb_plpython3u" "plpython3u" >}} | 在 jsonb 和 plpython3u 之间转换 |
| 3292 | {{< alias "ltree_plpython3u" "plpython3u" >}} | 在 ltree 和 plpython3u 之间转换 |
| 3293 | {{< alias "hstore_plpython3u" "plpython3u" >}} | 在 hstore 和 plpython3u 之间转换 |
| 3500 | {{< alias "prefix" "pg_prefix" >}} | 前缀树数据类型 |
| 3510 | {{< alias "semver" "pg_semver" >}} | 语义版本号数据类型 |
| 3520 | {{< alias "unit" "pgunit" >}} | SI 国标单位扩展 |
| 3530 | {{< alias "pgpdf" >}} | PDF数据类型，管理函数与全文检索 |
| 3550 | {{< alias "md5hash" >}} | 提供128位MD5的原生数据类型 |
| 3560 | {{< alias "asn1oid" >}} | ASN1OID数据类型支持 |
| 3570 | {{< alias "roaringbitmap" "pg_roaringbitmap" >}} | 支持RoaringBitmap数据类型 |
| 3590 | {{< alias "pg_sphere" "pgsphere" >}} | 球面对象函数、运算符与索引支持 |
| 3600 | {{< alias "country" "pg_country" >}} | 国家代码数据类型，遵循ISO 3166-1标准 |
| 3620 | {{< alias "currency" "pg_currency" >}} | 使用1字节表示的货币数据类型 |
| 3630 | {{< alias "collection" "pgcollection" >}} | 在PlPGSQL中使用的内存优化高性能集合数据结构 |
| 3700 | {{< alias "pgmp" >}} | 多精度算术扩展 |
| 3710 | {{< alias "numeral" >}} | 数值类型扩展 |
| 3720 | {{< alias "pg_rational" >}} | 使用BIGINT表示的有理数数据类型 |
| 3730 | {{< alias "uint" "pguint" >}} | 无符号整型数据类型 |
| 3740 | {{< alias "uint128" "pg_uint128" >}} | 原生128位无符号整型数据类型 |
| 3750 | {{< alias "hashtypes" >}} | 包括SHA1，MD5在内的多种哈希数据类型 |
| 3820 | {{< alias "ip4r" >}} | PostgreSQL 的 IPv4/v6 和 IPv4/v6 范围索引类型 |
| 3830 | {{< alias "pg_duration" >}} | 用于表示时间段的强化数据类型 |
| 3840 | {{< alias "uri" "pg_uri" >}} | URI数据类型 |
| 3850 | {{< alias "emailaddr" "pg_emailaddr" >}} | Email地址数据类型 |
| 3860 | {{< alias "acl" "pg_acl" >}} | ACL数据类型 |
| 3880 | {{< alias "pg_rrule" >}} | 日历重复规则RRULE数据类型 |
| 3890 | {{< alias "timestamp9" >}} | 纳秒分辨率时间戳 |
| 3920 | {{< alias "chkpass" >}} | 数据类型：自动加密的密码 |
| 3930 | {{< alias "isn" >}} | 用于国际产品编号标准的数据类型 |
| 3940 | {{< alias "seg" >}} | 表示线段或浮点间隔的数据类型 |
| 3950 | {{< alias "cube" >}} | 用于存储多维立方体的数据类型 |
| 3960 | {{< alias "ltree" >}} | 用于表示分层树状结构的数据类型 |
| 3970 | {{< alias "hstore" >}} | 用于存储（键，值）对集合的数据类型 |
| 3980 | {{< alias "citext" >}} | 提供大小写不敏感的字符串类型 |
| 3990 | {{< alias "xml2" >}} | XPath 查询和 XSLT |
| 4010 | {{< alias "gzip" "pg_gzip" >}} | 使用SQL执行Gzip压缩与解压缩 |
| 4020 | {{< alias "bzip" "pg_bzip" >}} | BZIP压缩解压缩函数包 |
| 4030 | {{< alias "zstd" "pg_zstd" >}} | ZSTD压缩解压缩函数包 |
| 4070 | {{< alias "http" "pg_http" >}} | HTTP客户端，允许在数据库内收发HTTP请求 (supabase) |
| 4080 | {{< alias "pg_net" >}} | 用 SQL 进行异步非阻塞HTTP/HTTPS 请求的扩展 (supabase) |
| 4090 | {{< alias "pg_curl" >}} | 封装CURL，执行各种用URL传输数据的操作 |
| 4100 | {{< alias "pg_retry" >}} | 在临时错误中使用指数退避重试语句 |
| 4150 | {{< alias "pgjq" >}} | 在Postgres中使用jq查询JSON |
| 4190 | {{< alias "url_encode" >}} | 提供URL编码解码函数 |
| 4230 | {{< alias "pgpcre" >}} | PCRE/Perl风格的正则表达式支持 |
| 4240 | {{< alias "icu_ext" >}} | 访问ICU库提供的函数 |
| 4250 | {{< alias "pgqr" >}} | 从数据库中直接生成QR二维码 |
| 4260 | {{< alias "pg_protobuf" >}} | 提供Protobuf函数支持 |
| 4270 | {{< alias "envvar" "pg_envvar" >}} | 获取环境变量的函数 |
| 4280 | {{< alias "floatfile" >}} | 将浮点数组存储到文件中而不是堆表中 |
| 4300 | {{< alias "pg_readme" >}} | 为模式与扩展生成Markdown文档 |
| 4301 | {{< alias "pg_readme_test_extension" "pg_readme" >}} | 为模式与扩展生成Markdown文档 |
| 4400 | {{< alias "hashlib" "pg_hashlib" >}} | 稳定哈希函数包 |
| 4430 | {{< alias "xxhash" "pg_xxhash" >}} | xxhash哈希函数包 |
| 4440 | {{< alias "shacrypt" >}} | 实现SHA256-CRYPT与SHA512-CRYPT密钥加密算法 |
| 4450 | {{< alias "cryptint" >}} | 加密INT与BIGINT类型 |
| 4460 | {{< alias "pguecc" "pg_ecdsa" >}} | PostgreSQL的uECC绑定，椭圆曲线加解密函数包 |
| 4540 | {{< alias "pg_uuidv7" >}} | UUIDv7 支持 |
| 4550 | {{< alias "permuteseq" >}} | 伪随机数ID置换生成器 |
| 4560 | {{< alias "pg_hashids" >}} | 加盐将整型ID转为短字符串ID |
| 4570 | {{< alias "sequential_uuids" >}} | 生成连续生成的UUID |
| 4590 | {{< alias "snowflake" >}} | Snowflake 风格 64 位 ID 生成与序列工具 |
| 4600 | {{< alias "topn" >}} | top-n JSONB 的类型 |
| 4610 | {{< alias "quantile" >}} | Quantile聚合函数 |
| 4620 | {{< alias "lower_quantile" >}} | Lower Quantile 聚合函数 |
| 4630 | {{< alias "count_distinct" >}} | COUNT(DISTINCT …) 聚合的替代方案 |
| 4640 | {{< alias "omnisketch" >}} | 实现OmniSketch数据结构，实现近似摘要聚合 |
| 4650 | {{< alias "ddsketch" >}} | 实现DDSketch数据结构，实现在线的Quantile聚合 |
| 4660 | {{< alias "vasco" >}} | 使用MIC发现数据中隐含的关联 |
| 4670 | {{< alias "xicor" "pgxicor" >}} | 在PG中计算XI相关系数 |
| 4680 | {{< alias "weighted_statistics" "pg_weighted_statistics" >}} | 针对稀疏数据的高性能加权统计量计算 |
| 4700 | {{< alias "tdigest" >}} | tdigest 聚合函数 |
| 4710 | {{< alias "first_last_agg" >}} | first() 与 last() 聚合函数 |
| 4720 | {{< alias "extra_window_functions" >}} | 额外的窗口函数 |
| 4730 | {{< alias "floatvec" >}} | 数组类型数学运算扩展 |
| 4740 | {{< alias "aggs_for_vecs" >}} | 针对数组类型的聚合函数集合扩展 |
| 4750 | {{< alias "aggs_for_arrays" >}} | 计算数组聚合统计值的函数包 |
| 4760 | {{< alias "pg_csv" >}} | 灵活的CSV聚合处理函数 |
| 4770 | {{< alias "arraymath" "pg_arraymath" >}} | 数组逐元素数学运算符包 |
| 4780 | {{< alias "pg_math" >}} | 使用GSL库的数学统计函数 |
| 4790 | {{< alias "random" "pg_random" >}} | 随机数生成器 |
| 4800 | {{< alias "base36" "pg_base36" >}} | Base36编码解码扩展 |
| 4810 | {{< alias "base62" "pg_base62" >}} | Base62编码解码扩展 |
| 4840 | {{< alias "financial" "pg_financial" >}} | 金融领域聚合函数 |
| 4880 | {{< alias "refint" >}} | 实现引用完整性的函数 |
| 4881 | {{< alias "autoinc" >}} | 用于自动递增字段的函数 |
| 4882 | {{< alias "insert_username" >}} | 用于跟踪谁更改了表的函数 |
| 4883 | {{< alias "moddatetime" >}} | 跟踪最后修改时间 |
| 4890 | {{< alias "tsm_system_time" >}} | 接受毫秒数限制的 TABLESAMPLE 方法 |
| 4900 | {{< alias "dict_xsyn" >}} | 用于扩展同义词处理的文本搜索字典模板 |
| 4910 | {{< alias "tsm_system_rows" >}} | 接受行数限制的 TABLESAMPLE 方法 |
| 4920 | {{< alias "tcn" >}} | 用触发器通知变更 |
| 4930 | {{< alias "uuid-ossp" >}} | 生成通用唯一标识符（UUIDs） |
| 4940 | {{< alias "btree_gist" >}} | 用GiST索引常见数据类型 |
| 4950 | {{< alias "btree_gin" >}} | 用GIN索引常见数据类型 |
| 4960 | {{< alias "intarray" >}} | 1维整数数组的额外函数、运算符和索引支持 |
| 4970 | {{< alias "intagg" >}} | 整数聚合器和枚举器（过时） |
| 4980 | {{< alias "dict_int" >}} | 用于整数的文本搜索字典模板 |
| 4990 | {{< alias "unaccent" >}} | 删除重音的文本搜索字典 |
| 5010 | {{< alias "pg_repack" >}} | 在线垃圾清理与表膨胀治理 |
| 5020 | {{< alias "pg_rewrite" >}} | 在线重写整表，不阻塞读写 |
| 5040 | {{< alias "pg_squeeze" >}} | 从关系中删除未使用空间 |
| 5050 | {{< alias "pg_dirtyread" >}} | 从表中读取尚未垃圾回收的行 |
| 5060 | {{< alias "pgfincore" >}} | 检查和管理操作系统缓冲区缓存 |
| 5070 | {{< alias "pg_cooldown" >}} | 从缓冲区中移除特定关系的页面 |
| 5100 | {{< alias "prioritize" "pg_prioritize" >}} | 获取和设置 PostgreSQL 后端的优先级 |
| 5110 | {{< alias "pg_checksums" >}} | 在离线模式下激活/启用/禁用数据库集群的校验和功能 |
| 5120 | {{< alias "pg_readonly" >}} | 将集群设置为只读 |
| 5150 | {{< alias "pgautofailover" >}} | PG 自动故障迁移 |
| 5160 | {{< alias "pg_catcheck" >}} | 用于诊断系统目录是否损坏的工具 |
| 5170 | {{< alias "pre_prepare" "preprepare" >}} | 在服务端预先准备好PreparedStatement备用 |
| 5200 | {{< alias "pg_orphaned" >}} | 处理孤儿文件的扩展插件 |
| 5210 | {{< alias "pg_crash" >}} | 向数据库进程随机发送信号模拟故障 |
| 5220 | {{< alias "pg_cheat_funcs" >}} | 一些超级实用的作弊函数 |
| 5230 | {{< alias "fio" "pg_fio" >}} | PostgreSQL文件IO函数包 |
| 5810 | {{< alias "pg_savior" >}} | 阻止不带条件的全表更新以避免意外事故 |
| 5820 | {{< alias "safeupdate" >}} | 强制在 UPDATE 和 DELETE 时提供 Where 条件 |
| 5860 | {{< alias "table_log" >}} | 记录某张表的修改日志并做表/行级时间点恢复 |
| 5880 | {{< alias "pgagent" >}} | PostgreSQL任务调度工具，与PGADMIN配合使用 |
| 5890 | {{< alias "pg_prewarm" >}} | 预热关系数据 |
| 5900 | {{< alias "pgpool_adm" "pgpool" >}} | PGPool 管理函数 |
| 5910 | {{< alias "pgpool_recovery" "pgpool" >}} | PGPool辅助扩展，从v4.3提供的恢复函数 |
| 5920 | {{< alias "pgpool_regclass" "pgpool" >}} | PGPool辅助扩展，RegClass替代 |
| 5930 | {{< alias "lo" >}} | 大对象维护 |
| 5940 | {{< alias "basic_archive" >}} | 归档模块样例 |
| 5950 | {{< alias "basebackup_to_shell" >}} | 添加一种备份到Shell终端到基础备份方式 |
| 5960 | {{< alias "old_snapshot" >}} | 支持 old_snapshot_threshold 的实用程序 |
| 5970 | {{< alias "adminpack" >}} | PostgreSQL 管理函数集合 |
| 5980 | {{< alias "amcheck" >}} | 校验关系完整性 |
| 5990 | {{< alias "pg_surgery" >}} | 对损坏的关系进行手术 |
| 6000 | {{< alias "pg_profile" >}} | PostgreSQL 数据库负载记录与AWR报表工具 |
| 6010 | {{< alias "pg_tracing" >}} | PostgreSQL分布式Tracing |
| 6210 | {{< alias "pg_show_plans" >}} | 打印所有当前正在运行查询的执行计划 |
| 6220 | {{< alias "pg_stat_kcache" >}} | 内核统计信息收集 |
| 6230 | {{< alias "pg_stat_monitor" >}} | 提供查询聚合统计、客户端信息、执行计划详细信息和直方图 |
| 6240 | {{< alias "pg_qualstats" >}} | 收集有关 quals 的统计信息的扩展 |
| 6250 | {{< alias "pg_store_plans" >}} | 跟踪所有执行的 SQL 语句的计划统计信息 |
| 6270 | {{< alias "pg_track_optimizer" >}} | 跟踪规划器决策与实际执行的差距 |
| 6280 | {{< alias "pg_wait_sampling" >}} | 基于采样的等待事件统计 |
| 6410 | {{< alias "pgsentinel" >}} | 活跃会话历史 |
| 6420 | {{< alias "system_stats" >}} | PostgreSQL 的系统统计函数 |
| 6440 | {{< alias "pgnodemx" >}} | 使用SQL查询获取操作系统指标 |
| 6450 | {{< alias "pg_proctab" "pgnodemx" >}} | 通过SQL接口访问操作系统进程表 |
| 6510 | {{< alias "bgw_replstatus" >}} | 用于汇报本机主从状态的后台工作进程 |
| 6520 | {{< alias "pgmeminfo" >}} | 显示内存使用情况 |
| 6530 | {{< alias "toastinfo" >}} | 显示TOAST字段的详细信息 |
| 6850 | {{< alias "pg_relusage" >}} | 打印查询引用的表与列 |
| 6880 | {{< alias "pg_overexplain" >}} | 允许 EXPLAIN 转储更多详细 |
| 6890 | {{< alias "pg_logicalinspect" >}} | 检视逻辑解码组件详情 |
| 6900 | {{< alias "pageinspect" >}} | 检查数据库页面二进制内容 |
| 6910 | {{< alias "pgrowlocks" >}} | 显示行级锁信息 |
| 6920 | {{< alias "sslinfo" >}} | 关于 SSL 证书的信息 |
| 6930 | {{< alias "pg_buffercache" >}} | 检查共享缓冲区缓存 |
| 6940 | {{< alias "pg_walinspect" >}} | 用于检查 PostgreSQL WAL 日志内容的函数 |
| 6950 | {{< alias "pg_freespacemap" >}} | 检查自由空间映射的内容（FSM） |
| 6960 | {{< alias "pg_visibility" >}} | 检查可见性图（VM）和页面级可见性信息 |
| 6970 | {{< alias "pgstattuple" >}} | 显示元组级统计信息 |
| 6980 | {{< alias "auto_explain" >}} | 提供一种自动记录执行计划的手段 |
| 6990 | {{< alias "pg_stat_statements" >}} | 跟踪所有执行的 SQL 语句的计划和执行统计信息 |
| 7000 | {{< alias "passwordcheck_cracklib" >}} | 使用cracklib加固PG用户密码 |
| 7010 | {{< alias "supautils" >}} | 用于在云环境中确保数据库集群的安全 |
| 7020 | {{< alias "pgsodium" >}} | 表数据加密存储 TDE |
| 7030 | {{< alias "supabase_vault" "pg_vault" >}} | 在 Vault 中存储加密凭证的扩展 (supabase) |
| 7100 | {{< alias "pgaudit" >}} | 提供审计功能 |
| 7120 | {{< alias "pgauditlogtofile" >}} | pgAudit 子扩展，将审计日志写入单独的文件中 |
| 7140 | {{< alias "logerrors" >}} | 用于收集日志文件中消息统计信息的函数 |
| 7150 | {{< alias "pg_auth_mon" >}} | 监控每个用户的连接尝试 |
| 7310 | {{< alias "credcheck" >}} | 明文凭证检查器 |
| 7320 | {{< alias "pgcryptokey" >}} | PG密钥管理 |
| 7330 | {{< alias "pg_pwhash" >}} | PostgreSQL 高级密码哈希扩展（Argon2/scrypt/yescrypt） |
| 7360 | {{< alias "login_hook" >}} | 在用户登陆时执行login_hook.login()函数 |
| 7370 | {{< alias "set_user" >}} | 增加了日志记录的 SET ROLE |
| 7380 | {{< alias "pg_snakeoil" >}} | PostgreSQL动态链接库反病毒功能 |
| 7390 | {{< alias "pgextwlist" >}} | PostgreSQL扩展白名单功能 |
| 7410 | {{< alias "sslutils" >}} | 使用SQL管理SSL证书 |
| 7420 | {{< alias "noset" "pg_noset" >}} | 阻止非超级用户使用SET/RESET设置变量 |
| 7500 | {{< alias "pg_tde" >}} | Percona加密存储引擎 |
| 7960 | {{< alias "sepgsql" >}} | 基于SELinux标签的强制访问控制 |
| 7970 | {{< alias "auth_delay" >}} | 在返回认证失败前暂停一会，避免爆破 |
| 7980 | {{< alias "pgcrypto" >}} | 实用加解密函数 |
| 7990 | {{< alias "passwordcheck" >}} | 用于强制拒绝修改弱密码的扩展 |
| 8510 | {{< alias "multicorn" >}} | 用Python编写自定义的外部数据源包装器 |
| 8520 | {{< alias "odbc_fdw" >}} | 访问ODBC可访问的任何外部数据源 |
| 8530 | {{< alias "jdbc_fdw" >}} | 访问JDBC可访问的任何外部数据源 |
| 8540 | {{< alias "pgspider_ext" >}} | 使用多种FDW访问远程数据库服务器 |
| 8600 | {{< alias "mysql_fdw" >}} | MySQL外部数据包装器 |
| 8610 | {{< alias "oracle_fdw" >}} | 提供对Oracle的外部数据源包装器 |
| 8620 | {{< alias "tds_fdw" >}} | TDS 数据库（Sybase/SQL Server）外部数据包装器 |
| 8630 | {{< alias "db2_fdw" >}} | 提供对DB2的外部数据源包装器 |
| 8640 | {{< alias "sqlite_fdw" >}} | SQLite 外部数据包装器 |
| 8670 | {{< alias "informix_fdw" >}} | Informix 外部数据包装器 |
| 8680 | {{< alias "nominatim_fdw" >}} | Nominatim 地理编码接口的 FDW 扩展 |
| 8700 | {{< alias "mongo_fdw" >}} | MongoDB 外部数据包装器 |
| 8710 | {{< alias "redis_fdw" >}} | 查询外部Redis数据源 |
| 8720 | {{< alias "redis" "pg_redis_pubsub" >}} | 从PG向Redis发送Pub/Sub消息 |
| 8730 | {{< alias "kafka_fdw" >}} | Kafka外部数据源包装器 |
| 8740 | {{< alias "hdfs_fdw" >}} | hdfs 外部数据包装器 |
| 8750 | {{< alias "firebird_fdw" >}} | Firebird外部数据源包装器 |
| 8810 | {{< alias "log_fdw" >}} | 访问PostgreSQL日志文件的FDW |
| 8970 | {{< alias "dblink" >}} | 从数据库内连接到其他 PostgreSQL 数据库 |
| 8980 | {{< alias "file_fdw" >}} | 访问外部文件的外部数据包装器 |
| 8990 | {{< alias "postgres_fdw" >}} | 用于远程 PostgreSQL 服务器的外部数据包装器 |
| 9000 | {{< alias "documentdb" >}} | 微软DocumentDB的API层 |
| 9010 | {{< alias "documentdb_core" "documentdb" >}} | 微软DocumentDB的核心API层实现 |
| 9020 | {{< alias "documentdb_distributed" "documentdb" >}} | DocumentDB多节点模式的API层 |
| 9030 | {{< alias "documentdb_extended_rum" "documentdb" >}} | DocumentDB扩展RUM索引访问方法 |
| 9100 | {{< alias "orafce" >}} | 模拟 Oracle RDBMS 的一部分函数和包的函数和运算符 |
| 9110 | {{< alias "pgtt" >}} | 类似Oracle的全局临时表功能 |
| 9120 | {{< alias "session_variable" >}} | Oracle兼容的会话变量/常量操作函数 |
| 9130 | {{< alias "pg_statement_rollback" >}} | 在服务端提供类似Oracle/DB2的语句级回滚能力 |
| 9270 | {{< alias "pg_dbms_errlog" >}} | 模仿 Oracle DBMS_ERRLOG 模块来记录特定表的DML错误 |
| 9300 | {{< alias "babelfishpg_common" >}} | SQL Server 数据类型兼容扩展 |
| 9310 | {{< alias "babelfishpg_tsql" >}} | SQL Server SQL语法兼容性扩展 |
| 9320 | {{< alias "babelfishpg_tds" >}} | SQL Server TDS线缆协议兼容扩展 |
| 9330 | {{< alias "babelfishpg_money" >}} | SQL Server 货币数据类型兼容扩展 |
| 9400 | {{< alias "spat" >}} | 在PG中嵌入Redis风格的内存数据库 |
| 9410 | {{< alias "pgmemcache" >}} | 为PG提供memcached兼容接口 |
| 9500 | {{< alias "pglogical" >}} | PostgreSQL逻辑复制：三方扩展实现 |
| 9501 | {{< alias "pglogical_origin" "pglogical" >}} | 用于从 Postgres 9.4 升级时的兼容性虚拟扩展 |
| 9510 | {{< alias "pglogical_ticker" >}} | pglogical复制延迟以秒计的精确视图 |
| 9520 | {{< alias "pgl_ddl_deploy" >}} | 使用 pglogical 执行自动 DDL 部署 |
| 9530 | {{< alias "pg_failover_slots" >}} | 在Failover过程中保留复制槽 |
| 9550 | {{< alias "pgactive" >}} | PostgreSQL多主逻辑复制 |
| 9560 | {{< alias "spock" >}} | PostgreSQL 多主逻辑复制扩展 |
| 9570 | {{< alias "lolor" >}} | 让 PostgreSQL 大对象兼容逻辑复制的扩展 |
| 9630 | {{< alias "wal2json" >}} | 用逻辑解码捕获 JSON 格式的 CDC 变更 |
| 9640 | {{< alias "wal2mongo" >}} | 使用逻辑解码捕获MongoDB JSON格式的CDC变更 |
| 9650 | {{< alias "decoderbufs" >}} | 将WAL逻辑解码为ProtocolBuffer协议的消息 |
| 9660 | {{< alias "decoder_raw" >}} | 逻辑复制解码输出插件：RAW SQL格式 |
| 9710 | {{< alias "repmgr" >}} | PostgreSQL复制管理组件 |
| 9820 | {{< alias "pg_fact_loader" >}} | 在 Postgres 中构建事实表 |
| 9830 | {{< alias "pg_bulkload" >}} | 向 PostgreSQL 中高速加载数据 |
| 9970 | {{< alias "test_decoding" >}} | 基于SQL的WAL逻辑解码样例 |
| 9980 | {{< alias "pgoutput" >}} | PG内置的逻辑解码输出插件 |

## SQL

{{< language "SQL" >}} {{< badge content="40 个扩展" color="gray" icon="cube" >}}

纯 SQL 扩展和函数

| ID | 扩展 | 描述 |
|:---:|:---|:---|
| 1020 | {{< alias "timeseries" "pg_timeseries" >}} | 时序数据API封装 |
| 1050 | {{< alias "emaj" >}} | 让数据库的子集具有细粒度日志和时间旅行功能 |
| 1060 | {{< alias "table_version" >}} | PostgreSQL 版本控制表扩展 |
| 1560 | {{< alias "geoip" >}} | IP 地理位置扩展（围绕 MaxMind GeoLite 数据集的包装器） |
| 1651 | {{< alias "mobilitydb_datagen" "mobilitydb" >}} | MobilityDB随机数据生成函数 |
| 2500 | {{< alias "pg_fkpart" >}} | 按外键实用程序进行表分区的扩展 |
| 2800 | {{< alias "index_advisor" >}} | 查询索引建议器 |
| 2870 | {{< alias "pgmb" >}} | 一个简单的PostgreSQL消息代理系统 |
| 2880 | {{< alias "pgmq" >}} | 基于Postgres实现类似AWS SQS/RSMQ的消息队列 |
| 3580 | {{< alias "pgfaceting" >}} | 使用倒排索引的高速切面查询 |
| 3610 | {{< alias "pg_xenophile" >}} | PostgreSQL i8n与l10n工具包 |
| 3611 | {{< alias "l10n_table_dependent_extension" "pg_xenophile" >}} | PostgreSQL l10n 工具包 |
| 3870 | {{< alias "debversion" >}} | Debian版本号数据类型 |
| 4160 | {{< alias "pgjwt" >}} | JSON Web Token API 的PG实现 (supabase) |
| 4180 | {{< alias "pg_html5_email_address" >}} | 验证Email是否符合HTML5规范的扩展 |
| 4200 | {{< alias "pgsql_tweaks" >}} | 一些日常会用到的便利函数与视图 |
| 4220 | {{< alias "pg_extra_time" >}} | 一些关于日期与时间的扩展函数 |
| 4310 | {{< alias "ddl_historization" >}} | 用SQL将所有DDL变更写入到数据库表中 |
| 4320 | {{< alias "data_historization" >}} | 用SQL将数据变更历史保存到分区表中 |
| 4330 | {{< alias "schedoc" "pg_schedoc" >}} | 在Django与DBT之间通过注释文档交换元数据 |
| 4470 | {{< alias "sparql" "pgsparql" >}} | 使用SQL查询SPARQL数据源 |
| 5080 | {{< alias "ddlx" "pg_ddlx" >}} | 提取数据库对象的DDL |
| 5140 | {{< alias "pg_permissions" >}} | 查看对象权限并将其与期望状态进行比较 |
| 5180 | {{< alias "pg_upless" >}} | 检测表上的无用UPDATE |
| 5190 | {{< alias "pgcozy" >}} | 根据先前的pg_buffercache快照预热内存缓冲区 |
| 5850 | {{< alias "pg_drop_events" >}} | 记录删表删列删视图的事务号，辅助PITR确定时间点 |
| 6260 | {{< alias "pg_track_settings" >}} | 跟踪设置更改 |
| 6430 | {{< alias "meta" "pg_meta" >}} | 标准化，更友好的PostgreSQL系统目录视图 |
| 6500 | {{< alias "pg_sqlog" >}} | 提供访问PostgreSQL日志的SQL接口 |
| 6860 | {{< alias "pagevis" >}} | 使用ASCII字符可视化数据库物理页面布局 |
| 7130 | {{< alias "pg_auditor" >}} | 审计数据变更并提供闪回能力 |
| 7160 | {{< alias "pg_jobmon" >}} | 记录和监控函数 |
| 8650 | {{< alias "pgbouncer_fdw" >}} | 用SQL查询pgbouncer统计信息，并执行pgbouncer命令 |
| 8800 | {{< alias "aws_s3" >}} | 从S3导入导出数据的外部数据源包装器 |
| 9240 | {{< alias "pg_dbms_metadata" >}} | 添加 Oracle DBMS_METADATA 兼容性支持的扩展 |
| 9250 | {{< alias "pg_dbms_lock" >}} | 为PG添加对 Oracle DBMS_LOCK 的完整兼容性支持 |
| 9260 | {{< alias "pg_dbms_job" >}} | 添加 Oracle DBMS_JOB 兼容性支持的扩展 |
| 9290 | {{< alias "pg_utl_smtp" >}} | Oracle UTL_SMTP 兼容扩展（基于 plperlu） |
| 9540 | {{< alias "db_migrator" >}} | 使用FDW从其他DBMS迁移到PostgreSQL |
| 9700 | {{< alias "mimeo" >}} | 在PostgreSQL实例间进行表级复制 |

## Rust

{{< language "Rust" >}} {{< badge content="39 个扩展" color="gray" icon="cube" >}}

使用 pgrx 框架用 Rust 编写的扩展

| ID | 扩展 | 描述 |
|:---:|:---|:---|
| 1010 | {{< alias "timescaledb_toolkit" >}} | 超表分析查询，时间序列流式处理，以及其他SQL工具 |
| 1090 | {{< alias "pg_later" >}} | 执行查询，并在稍后异步获取查询结果 |
| 1570 | {{< alias "pg_polyline" >}} | Google快速Polyline编码解码扩展 |
| 1680 | {{< alias "tzf" "pg_tzf" >}} | 快速根据GPS经纬度坐标查找时区 |
| 1810 | {{< alias "vchord" >}} | 使用Rust重写的高性能向量扩展 |
| 1820 | {{< alias "vectorscale" "pgvectorscale" >}} | 使用DiskANN算法对向量进行高效索引 |
| 1830 | {{< alias "vectorize" "pg_vectorize" >}} | 在PostgreSQL中封装RAG向量检索服务 |
| 1860 | {{< alias "pg_summarize" >}} | 使用LLM对文本字段进行总结 |
| 1870 | {{< alias "pg_tiktoken" >}} | 在PostgreSQL中计算OpenAI使用的Token数 |
| 1890 | {{< alias "pgml" >}} | PostgresML：用SQL运行机器学习算法并训练模型 |
| 2100 | {{< alias "pg_search" >}} | ParadeDB BM25算法全文检索插件，ES全文检索 |
| 2140 | {{< alias "pg_bestmatch" >}} | 在数据库内生成BM25稀疏向量 |
| 2150 | {{< alias "vchord_bm25" >}} | BM25排序算法 |
| 2160 | {{< alias "pg_tokenizer" >}} | 用于全文检索的分词器 |
| 2420 | {{< alias "pg_analytics" >}} | 由 DuckDB 驱动的数据分析引擎 |
| 2440 | {{< alias "pg_mooncake" >}} | PostgreSQL列式存储表 |
| 2480 | {{< alias "pg_parquet" >}} | 在PostgreSQL与本地/S3中的Parquet文件复制数据 |
| 2750 | {{< alias "pg_graphql" >}} | PG内的GraphQL支持 |
| 2760 | {{< alias "pg_jsonschema" >}} | 提供JSON Schema校验能力 |
| 2920 | {{< alias "pg_cardano" >}} | Cardano相关工具包：加密函数，地址编解码，区块链处理 |
| 3040 | {{< alias "plprql" >}} | 在PostgreSQL使用PRQL——管线式关系查询语言 |
| 3540 | {{< alias "pglite_fusion" >}} | 在PG表中嵌入SQLite数据库作为数据类型 |
| 4170 | {{< alias "pg_smtp_client" >}} | 使用SMTP从PostgreSQL内发送邮件的客户端扩展 |
| 4290 | {{< alias "pg_render" >}} | 使用SQL渲染HTML页面 |
| 4500 | {{< alias "pg_idkit" >}} | 生成各式各样的唯一标识符：UUIDv6, ULID, KSUID |
| 4510 | {{< alias "pgx_ulid" >}} | ULID数据类型与函数 |
| 4580 | {{< alias "typeid" "pg_typeid" >}} | PG原生TypeID类型与函数 |
| 4830 | {{< alias "pg_base58" >}} | Base58 编码/解码函数 |
| 4850 | {{< alias "convert" "pg_convert" >}} | 用于空间里程等的公英制转换函数 |
| 5090 | {{< alias "pglinter" >}} | PG数据库规则检查插件 |
| 5130 | {{< alias "pgdd" >}} | 提供通过标准SQL查询数据库目录集簇的能力 |
| 5830 | {{< alias "pg_strict" >}} | 防止不带WHERE条件的危险UPDATE和DELETE操作 |
| 6540 | {{< alias "explain_ui" "pg_explain_ui" >}} | 快速跳转至PEV查阅可视化执行计划 |
| 7040 | {{< alias "pg_session_jwt" >}} | 使用JWT进行会话认证 |
| 7050 | {{< alias "anon" "pg_anon" >}} | 数据匿名化处理工具 |
| 7060 | {{< alias "pgsmcrypto" >}} | 为PostgreSQL提供商密算法支持：SM2,SM3,SM4 |
| 7070 | {{< alias "pg_enigma" >}} | PostgreSQL 加密数据类型 |
| 8500 | {{< alias "wrappers" >}} | Supabase提供的外部数据源包装器捆绑包 |
| 8660 | {{< alias "etcd_fdw" >}} | etcd分布式键值存储外部数据包装器 |

## Data

{{< language "Data" >}} {{< badge content="10 个扩展" color="gray" icon="cube" >}}

仅包含数据的扩展

| ID | 扩展 | 描述 |
|:---:|:---|:---|
| 2270 | {{< alias "hunspell_cs_cz" >}} | Hunspell捷克语全文检索词典 |
| 2271 | {{< alias "hunspell_de_de" >}} | Hunspell德语全文检索词典 |
| 2272 | {{< alias "hunspell_en_us" >}} | Hunspell英语全文检索词典 |
| 2273 | {{< alias "hunspell_fr" >}} | Hunspell法语全文检索词典 |
| 2274 | {{< alias "hunspell_ne_np" >}} | Hunspell尼泊尔语全文检索词典 |
| 2275 | {{< alias "hunspell_nl_nl" >}} | Hunspell荷兰语全文检索词典 |
| 2276 | {{< alias "hunspell_nn_no" >}} | Hunspell挪威语全文检索词典 |
| 2277 | {{< alias "hunspell_pt_pt" >}} | Hunspell葡萄牙语全文检索词典 |
| 2278 | {{< alias "hunspell_ru_ru" >}} | Hunspell俄语全文检索词典 |
| 2279 | {{< alias "hunspell_ru_ru_aot" >}} | Hunspell俄语全文检索词典（来自AOT.ru小组） |

## C++

{{< language "C++" >}} {{< badge content="7 个扩展" color="gray" icon="cube" >}}

使用 C++ 特性和库的扩展

| ID | 扩展 | 描述 |
|:---:|:---|:---|
| 1510 | {{< alias "pgrouting" >}} | 提供寻路能力 |
| 2430 | {{< alias "pg_duckdb" >}} | 在PostgreSQL中的嵌入式DuckDB扩展 |
| 2460 | {{< alias "pg_clickhouse" >}} | 从PostgreSQL中查询ClickHouse的接口 |
| 2710 | {{< alias "hll" >}} | hyperloglog 数据类型 |
| 2730 | {{< alias "pg_ai_query" >}} | AI驱动的 Postgres SQL 查询生成 |
| 2930 | {{< alias "rdkit" >}} | 在PostgreSQL化学领域数据管理功能 |
| 3010 | {{< alias "plv8" >}} | PL/JavaScript (v8) 可信过程程序语言 |

## Python

{{< language "Python" >}} {{< badge content="2 个扩展" color="gray" icon="cube" >}}

使用 Python 编写的扩展

| ID | 扩展 | 描述 |
|:---:|:---|:---|
| 3210 | {{< alias "faker" >}} | 插入生成的测试伪造数据，Python库的包装 |
| 6870 | {{< alias "powa" >}} | PostgreSQL 工作负载分析器-核心 |

## Java

{{< language "Java" >}} {{< badge content="1 个扩展" color="gray" icon="cube" >}}

在 JVM 上运行的扩展

| ID | 扩展 | 描述 |
|:---:|:---|:---|
| 3090 | {{< alias "pljava" >}} | Java 程序语言 |

