# Домашнее задание к занятию "6.4. PostgreSQL"

## Задача 1

Используя docker поднимите инстанс PostgreSQL (версию 13). Данные БД сохраните в volume.

```yml
version: '3.2'

volumes:
  db_data:
  db_backup:

services:
  postgres:
    image: postgres:13
    environment:
      - POSTGRES_DB=postgres
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
    volumes:
      - db_data:/var/lib/postgresql/data
      - db_backup:/var/lib/postgresql/backups
    ports:
      - 5433:5432
```

Подключитесь к БД PostgreSQL используя `psql`.

```SQL
[root@homesrv 6-4]# docker exec -it 6-4_postgres_1 bash
root@c95b55a7c4f5:/# psql -d postgres -U postgres
psql (13.7 (Debian 13.7-1.pgdg110+1))
Type "help" for help.

postgres=#
```

Воспользуйтесь командой `\?` для вывода подсказки по имеющимся в `psql` управляющим командам.

**Найдите и приведите** управляющие команды для:
- вывода списка БД
```SQL
postgres=# \l+
                                                                   List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges   |  Size   | Tablespace |                Description
-----------+----------+----------+------------+------------+-----------------------+---------+------------+--------------------------------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |                       | 7901 kB | pg_default | default administrative connection database
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +| 7753 kB | pg_default | unmodifiable empty database
           |          |          |            |            | postgres=CTc/postgres |         |            |
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +| 7753 kB | pg_default | default template for new databases
           |          |          |            |            | postgres=CTc/postgres |         |            |
(3 rows)

postgres=# \l
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
(3 rows)

postgres=#
```
- подключения к БД
```SQL
postgres=# \c postgres
You are now connected to database "postgres" as user "postgres".
postgres=#
```
- вывода списка таблиц
```SQL
postgres=# \dt
Did not find any relations.
postgres=# \dtS+
                                        List of relations
   Schema   |          Name           | Type  |  Owner   | Persistence |    Size    | Description
------------+-------------------------+-------+----------+-------------+------------+-------------
 pg_catalog | pg_aggregate            | table | postgres | permanent   | 56 kB      |
 pg_catalog | pg_am                   | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_amop                 | table | postgres | permanent   | 80 kB      |
 pg_catalog | pg_amproc               | table | postgres | permanent   | 64 kB      |
 pg_catalog | pg_attrdef              | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_attribute            | table | postgres | permanent   | 456 kB     |
 pg_catalog | pg_auth_members         | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_authid               | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_cast                 | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_class                | table | postgres | permanent   | 136 kB     |
 pg_catalog | pg_collation            | table | postgres | permanent   | 240 kB     |
 pg_catalog | pg_constraint           | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_conversion           | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_database             | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_db_role_setting      | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_default_acl          | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_depend               | table | postgres | permanent   | 488 kB     |
 pg_catalog | pg_description          | table | postgres | permanent   | 376 kB     |
 pg_catalog | pg_enum                 | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_event_trigger        | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_extension            | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_foreign_data_wrapper | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_foreign_server       | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_foreign_table        | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_index                | table | postgres | permanent   | 64 kB      |
 pg_catalog | pg_inherits             | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_init_privs           | table | postgres | permanent   | 56 kB      |
 pg_catalog | pg_language             | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_largeobject          | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_largeobject_metadata | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_namespace            | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_opclass              | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_operator             | table | postgres | permanent   | 144 kB     |
 pg_catalog | pg_opfamily             | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_partitioned_table    | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_policy               | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_proc                 | table | postgres | permanent   | 688 kB     |
 pg_catalog | pg_publication          | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_publication_rel      | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_range                | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_replication_origin   | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_rewrite              | table | postgres | permanent   | 656 kB     |
 pg_catalog | pg_seclabel             | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_sequence             | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_shdepend             | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_shdescription        | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_shseclabel           | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_statistic            | table | postgres | permanent   | 248 kB     |
 pg_catalog | pg_statistic_ext        | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_statistic_ext_data   | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_subscription         | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_subscription_rel     | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_tablespace           | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_transform            | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_trigger              | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_ts_config            | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_ts_config_map        | table | postgres | permanent   | 56 kB      |
 pg_catalog | pg_ts_dict              | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_ts_parser            | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_ts_template          | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_type                 | table | postgres | permanent   | 120 kB     |
 pg_catalog | pg_user_mapping         | table | postgres | permanent   | 8192 bytes |
(62 rows)

postgres=#
```
- вывода описания содержимого таблиц
```SQL

postgres=# \dS pg_am
               Table "pg_catalog.pg_am"
  Column   |  Type   | Collation | Nullable | Default
-----------+---------+-----------+----------+---------
 oid       | oid     |           | not null |
 amname    | name    |           | not null |
 amhandler | regproc |           | not null |
 amtype    | "char"  |           | not null |
Indexes:
    "pg_am_name_index" UNIQUE, btree (amname)
    "pg_am_oid_index" UNIQUE, btree (oid)

postgres=#
```
- выхода из psql
```SQL
postgres=# \q
root@c95b55a7c4f5:/#
```

## Задача 2

Используя `psql` создайте БД `test_database`.
```SQL
[root@homesrv 6-4]# docker exec -it 6-4_postgres_1 bash
root@c95b55a7c4f5:/# psql -d postgres -U postgres
psql (13.7 (Debian 13.7-1.pgdg110+1))
Type "help" for help.

postgres=# CREATE DATABASE test_database;
CREATE DATABASE
postgres=#
```

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-04-postgresql/test_data).

Восстановите бэкап БД в `test_database`.
```SQL
[root@homesrv 6-4]# cat test_dump.sql | docker exec -i 6-4_postgres_1 psql -U postgres test_database
SET
SET
SET
SET
SET
 set_config
------------

(1 row)

SET
SET
SET
SET
SET
SET
CREATE TABLE
ALTER TABLE
CREATE SEQUENCE
ALTER TABLE
ALTER SEQUENCE
ALTER TABLE
COPY 8
 setval
--------
      8
(1 row)

ALTER TABLE
[root@homesrv 6-4]#
```
Перейдите в управляющую консоль `psql` внутри контейнера.

Подключитесь к восстановленной БД и проведите операцию ANALYZE для сбора статистики по таблице.
```SQL
test_database=# \dt
         List of relations
 Schema |  Name  | Type  |  Owner
--------+--------+-------+----------
 public | orders | table | postgres
(1 row)

test_database=# ANALYZE VERBOSE public.orders;
INFO:  analyzing "public.orders"
INFO:  "orders": scanned 1 of 1 pages, containing 8 live rows and 0 dead rows; 8 rows in sample, 8 estimated total rows
ANALYZE
test_database=#
```

Используя таблицу [pg_stats](https://postgrespro.ru/docs/postgresql/12/view-pg-stats), найдите столбец таблицы `orders`
с наибольшим средним значением размера элементов в байтах.

**Приведите в ответе** команду, которую вы использовали для вычисления и полученный результат.

```SQL
test_database=# select avg_width from pg_stats where tablename='orders';
 avg_width
-----------
         4
        16
         4
(3 rows)

test_database=#
```

## Задача 3

Архитектор и администратор БД выяснили, что ваша таблица orders разрослась до невиданных размеров и
поиск по ней занимает долгое время. Вам, как успешному выпускнику курсов DevOps в нетологии предложили
провести разбиение таблицы на 2 (шардировать на orders_1 - price>499 и orders_2 - price<=499).

Предложите SQL-транзакцию для проведения данной операции.

Можно ли было изначально исключить "ручное" разбиение при проектировании таблицы orders?

```SQL
test_database=# CREATE TABLE orders_1 (CHECK (price>499)) INHERITS (orders);
CREATE TABLE
test_database=# CREATE TABLE orders_2 (CHECK (price<=499)) INHERITS (orders);
CREATE TABLE
test_database=# \dt
          List of relations
 Schema |   Name   | Type  |  Owner
--------+----------+-------+----------
 public | orders   | table | postgres
 public | orders_1 | table | postgres
 public | orders_2 | table | postgres
(3 rows)

test_database=#
```

## Задача 4

Используя утилиту `pg_dump` создайте бекап БД `test_database`.

```SQL
[root@homesrv 6-4]# docker exec -t 6-4_postgres_1 pg_dumpall -c -U postgres  -p 5432 -h localhost -l test_database -f /var/lib/postgresql/backup/dump_`date +%d-%m-%Y"_"%H_%M_%S`.sql
[root@homesrv 6-4]# cd /var/lib/docker/volumes/6-4_db_backup/_data
[root@homesrv _data]# cat dump_22-05-2022_21_08_27.sql
--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Drop databases (except postgres and template1)
--

DROP DATABASE test_database;




--
-- Drop roles
--

DROP ROLE postgres;


--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'md53175bce1d3201d16594cebf9d7eb3f9d';






--
-- Databases
--

--
-- Database "template1" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.7 (Debian 13.7-1.pgdg110+1)
-- Dumped by pg_dump version 13.7 (Debian 13.7-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

UPDATE pg_catalog.pg_database SET datistemplate = false WHERE datname = 'template1';
DROP DATABASE template1;
--
-- Name: template1; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE template1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE template1 OWNER TO postgres;

\connect template1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE template1; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE template1 IS 'default template for new databases';


--
-- Name: template1; Type: DATABASE PROPERTIES; Schema: -; Owner: postgres
--

ALTER DATABASE template1 IS_TEMPLATE = true;


\connect template1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE template1; Type: ACL; Schema: -; Owner: postgres
--

REVOKE CONNECT,TEMPORARY ON DATABASE template1 FROM PUBLIC;
GRANT CONNECT ON DATABASE template1 TO PUBLIC;


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.7 (Debian 13.7-1.pgdg110+1)
-- Dumped by pg_dump version 13.7 (Debian 13.7-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO postgres;

\connect postgres

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- PostgreSQL database dump complete
--

--
-- Database "test_database" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.7 (Debian 13.7-1.pgdg110+1)
-- Dumped by pg_dump version 13.7 (Debian 13.7-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: test_database; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE test_database WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE test_database OWNER TO postgres;

\connect test_database

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id integer NOT NULL,
    title character varying(80) NOT NULL,
    price integer DEFAULT 0
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: orders_1; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders_1 (
    CONSTRAINT orders_1_price_check CHECK ((price > 499))
)
INHERITS (public.orders);


ALTER TABLE public.orders_1 OWNER TO postgres;

--
-- Name: orders_2; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders_2 (
    CONSTRAINT orders_2_price_check CHECK ((price <= 499))
)
INHERITS (public.orders);


ALTER TABLE public.orders_2 OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_id_seq OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: orders_1 id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders_1 ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: orders_1 price; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders_1 ALTER COLUMN price SET DEFAULT 0;


--
-- Name: orders_2 id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders_2 ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: orders_2 price; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders_2 ALTER COLUMN price SET DEFAULT 0;


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, title, price) FROM stdin;
1       War and peace   100
2       My little database      500
3       Adventure psql time     300
4       Server gravity falls    300
5       Log gossips     123
6       WAL never lies  900
7       Me and my bash-pet      499
8       Dbiezdmin       501
\.


--
-- Data for Name: orders_1; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders_1 (id, title, price) FROM stdin;
\.


--
-- Data for Name: orders_2; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders_2 (id, title, price) FROM stdin;
\.


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 8, true);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

[root@homesrv _data]#
```

Как бы вы доработали бэкап-файл, чтобы добавить уникальность значения столбца `title` для таблиц `test_database`?

Создание индекса по выражению lower(title), позволяющего эффективно выполнять регистронезависимый поиск:

CREATE INDEX ON orders ((lower(title)));
