## Usage

Sources:

- [README](https://github.com/sweatybridge/pg_ros2/blob/4886868056540ecaf8d389401a202260902e177e/README.md)
- [Control file](https://github.com/sweatybridge/pg_ros2/blob/4886868056540ecaf8d389401a202260902e177e/pg_ros2.control)
- [Cargo.toml](https://github.com/sweatybridge/pg_ros2/blob/4886868056540ecaf8d389401a202260902e177e/Cargo.toml)
- [src/lib.rs](https://github.com/sweatybridge/pg_ros2/blob/4886868056540ecaf8d389401a202260902e177e/src/lib.rs)

`pg_ros2` exposes ROS 2 graph snapshots in PostgreSQL tables and streams selected topic messages through notifications. Version 0.2.0 documents a Linux/Ubuntu 22.04, PostgreSQL 18, ROS 2 Humble and rclrs runtime.

### Graph Snapshots

The PostgreSQL service must inherit the ROS runtime environment. Add `pg_ros2` to `shared_preload_libraries`, set `pg_ros2.database` to an existing database, and restart. Install as a superuser.

```sql
CREATE EXTENSION pg_ros2;
SELECT * FROM nodes;
SELECT * FROM topics;
SELECT * FROM parameters;
SELECT * FROM worker_status;
SELECT * FROM parameter_status;
```

### Topic Subscription

In a listener session, subscribe to the channel matching the fully qualified ROS topic. In another autocommit session, start the long-running procedure:

```sql
LISTEN "/chatter";
```

```sql
SET statement_timeout = 0;
SET client_connection_check_interval = '1s';
CALL subscribe('/chatter');
```

The procedure can run independently of the graph worker. The optional `pg_durable` integration manages it as a workflow; it is not a hard dependency.

### Access and Delivery

Grant table reads and procedure execution only to appropriate roles. Remote parameter snapshots may contain secrets. Notification channels have no per-channel ACL: other database users can listen to a channel.

Graph snapshots are eventually refreshed and may become stale after worker or ROS failures. Topic delivery is transient and best-effort with bounded buffering and drops; it supplies no replay or durable consumer offset. Restart the procedure after a server failure and monitor worker status rather than interpreting an old snapshot as current.
