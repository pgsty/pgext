## Usage

Sources:

- [Official README](https://github.com/nvorobev/pg_metricus_c/blob/085cf982da50211938cf014551cfe194ebc48b27/README.md)
- [Extension SQL for 1.0](https://github.com/nvorobev/pg_metricus_c/blob/085cf982da50211938cf014551cfe194ebc48b27/pg_metricus--1.0.sql)
- [UDP sender implementation](https://github.com/nvorobev/pg_metricus_c/blob/085cf982da50211938cf014551cfe194ebc48b27/pg_metricus.c)

`pg_metricus` sends application-formatted metrics from SQL or PL/pgSQL to an IPv4 UDP endpoint such as Brubeck or Graphite. It provides transport only: the caller must construct the exact line protocol expected by the receiver.

### Core Workflow

Configure the destination address and port, reload PostgreSQL configuration, and install the extension in a chosen schema.

```conf
pg_metricus.host = '10.9.5.164'
pg_metricus.port = 8124
```

```sql
CREATE SCHEMA metricus;
CREATE EXTENSION pg_metricus SCHEMA metricus;

SELECT metricus.send_metric(
  format(E'%s.%s:%s|%s\n', 'db.sql', 'latency', 12, 'ms')
);
```

The endpoint settings can be changed with a configuration reload; the extension does not require preloading.

### SQL Object and Delivery Semantics

- `send_metric(text) RETURNS void` sends the supplied bytes as one UDP datagram.

Version 1.0 accepts only a numeric IPv4 address because its C implementation uses `inet_pton` with `AF_INET`. UDP provides no acknowledgement or retry, and the function does not make delivery transactional: a datagram already sent remains sent even if the surrounding database transaction later aborts. Avoid placing secrets in metric text, and treat successful return as a local send attempt rather than confirmation that the collector stored the metric.
