# Домашнее задание к занятию "6.5. Elasticsearch"

## Задача 1

В этом задании вы потренируетесь в:
- установке elasticsearch
- первоначальном конфигурировании elastcisearch
- запуске elasticsearch в docker

Используя докер образ [elasticsearch:7](https://hub.docker.com/_/elasticsearch) как базовый:

- составьте Dockerfile-манифест для elasticsearch
- соберите docker-образ и сделайте `push` в ваш docker.io репозиторий
- запустите контейнер из получившегося образа и выполните запрос пути `/` c хост-машины

Требования к `elasticsearch.yml`:
- данные `path` должны сохраняться в `/var/lib`
- имя ноды должно быть `netology_test`

В ответе приведите:
- текст Dockerfile манифеста
```yml
FROM elasticsearch:7.17.4
ADD elasticsearch.yml /usr/share/elasticsearch/config/

RUN mkdir /var/lib/logs \
    && chown elasticsearch:elasticsearch /var/lib/logs \
    && mkdir /var/lib/data \
    && chown elasticsearch:elasticsearch /var/lib/data
```
- ссылку на образ в репозитории dockerhub
[Ссылка на образ](https://hub.docker.com/repository/docker/sergej1024/elasticsearch)
- ответ `elasticsearch` на запрос пути `/` в json виде
```bash
[root@homesrv 6-5]# docker build -t sergej1024/elasticsearch:7.17.4 .
Sending build context to Docker daemon  6.656kB
Step 1/3 : FROM elasticsearch:7.17.4
7.17.4: Pulling from library/elasticsearch
d5fd17ec1767: Pull complete
3aceae0816c1: Pull complete
6f282e391d7d: Pull complete
e0d1c86ab271: Pull complete
1c2d02571b2b: Pull complete
25fb4b01f643: Pull complete
606786004049: Pull complete
28ec7712324b: Pull complete
7d5976c54116: Pull complete
Digest: sha256:529b3cfec4354beda158c6c7f2f8015cbdc9432a48c1d63e824d6fd728f30db2
Status: Downloaded newer image for elasticsearch:7.17.4
 ---> d64cccab426e
Step 2/3 : ADD elasticsearch.yml /usr/share/elasticsearch/config/
 ---> 874814dd84bc
Step 3/3 : RUN mkdir /var/lib/logs     && chown elasticsearch:elasticsearch /var/lib/logs     && mkdir /var/lib/data     && chown elasticsearch:elasticsearch /var/lib/data
 ---> Running in c8a72c274fb7
Removing intermediate container c8a72c274fb7
 ---> df61d77e995d
Successfully built df61d77e995d
Successfully tagged sergej1024/elasticsearch:7.17.4
[root@homesrv 6-5]# docker run --name es01 -p 9200:9200 -d sergej1024/elasticsearch:7.17.4
20a16ee75cb25e357b72973321a6c8a262441813108e91ede1a922f92595ef11
[root@homesrv 6-5]# docker ps -a
CONTAINER ID   IMAGE                             COMMAND                  CREATED          STATUS          PORTS                                                 NAMES
20a16ee75cb2   sergej1024/elasticsearch:7.17.4   "/bin/tini -- /usr/l…"   12 seconds ago   Up 11 seconds   0.0.0.0:9200->9200/tcp, :::9200->9200/tcp, 9300/tcp   es01
c95b55a7c4f5   postgres:13                       "docker-entrypoint.s…"   12 days ago      Up 11 minutes   0.0.0.0:5433->5432/tcp, :::5433->5432/tcp             6-4_postgres_1
[root@homesrv 6-5]# docker exec -it es01 /bin/sh
sh-5.0# curl -X GET "localhost:9200/"
{
  "name" : "netology_test",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "J4tdtOtCT1KcLrnCcrfrTw",
  "version" : {
    "number" : "7.17.4",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "79878662c54c886ae89206c685d9f1051a9d6411",
    "build_date" : "2022-05-18T18:04:20.964345128Z",
    "build_snapshot" : false,
    "lucene_version" : "8.11.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
sh-5.0#
```



Подсказки:
- при сетевых проблемах внимательно изучите кластерные и сетевые настройки в elasticsearch.yml
- при некоторых проблемах вам поможет docker директива ulimit
- elasticsearch в логах обычно описывает проблему и пути ее решения
- обратите внимание на настройки безопасности такие как `xpack.security.enabled`
- если докер образ не запускается и падает с ошибкой 137 в этом случае может помочь настройка `-e ES_HEAP_SIZE`
- при настройке `path` возможно потребуется настройка прав доступа на директорию

Далее мы будем работать с данным экземпляром elasticsearch.


## Задача 2

В этом задании вы научитесь:
- создавать и удалять индексы
- изучать состояние кластера
- обосновывать причину деградации доступности данных

Ознакомтесь с [документацией](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html)
и добавьте в `elasticsearch` 3 индекса, в соответствии со таблицей:

| Имя | Количество реплик | Количество шард |
|-----|-------------------|-----------------|
| ind-1| 0 | 1 |
| ind-2 | 1 | 2 |
| ind-3 | 2 | 4 |

Получите список индексов и их статусов, используя API и **приведите в ответе** на задание.

Получите состояние кластера `elasticsearch`, используя API.

Как вы думаете, почему часть индексов и кластер находится в состоянии yellow?

Удалите все индексы.

**Важно**

При проектировании кластера elasticsearch нужно корректно рассчитывать количество реплик и шард,
иначе возможна потеря данных индексов, вплоть до полной, при деградации системы.

```bash
[root@homesrv 6-5]# docker exec -it es01 /bin/sh
sh-5.0# curl -X PUT localhost:9200/ind-1 -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 1,  "number_of_replicas": 0 }}'
{"acknowledged":true,"shards_acknowledged":true,"index":"ind-1"}sh-5.0#
sh-5.0# curl -X PUT localhost:9200/ind-2 -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 2,  "number_of_replicas": 1 }}'
{"acknowledged":true,"shards_acknowledged":true,"index":"ind-2"}sh-5.0#
sh-5.0# curl -X PUT localhost:9200/ind-3 -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 4,  "number_of_replicas": 2 }}'
{"acknowledged":true,"shards_acknowledged":true,"index":"ind-3"}sh-5.0#
sh-5.0# curl -X GET 'http://localhost:9200/_cat/indices?v'
health status index uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   ind-1 uSUKA4RzQtGBBkp8uZmMtA   1   0          0            0       226b           226b
yellow open   ind-3 7HSdMyLlTOiKef2xXX2uqw   4   2          0            0       904b           904b
yellow open   ind-2 NmXU3uYFTl25k24PYukyBg   2   1          0            0       452b           452b
sh-5.0# curl -X GET 'http://localhost:9200/_cluster/health/ind-1?pretty'
{
  "cluster_name" : "elasticsearch",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 1,
  "active_shards" : 1,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}
sh-5.0# curl -X GET 'http://localhost:9200/_cluster/health/ind-2?pretty'
{
  "cluster_name" : "elasticsearch",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 2,
  "active_shards" : 2,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 2,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 47.368421052631575
}
sh-5.0# curl -X GET 'http://localhost:9200/_cluster/health/ind-3?pretty'
{
  "cluster_name" : "elasticsearch",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 4,
  "active_shards" : 4,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 8,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 47.368421052631575
}
sh-5.0# curl -XGET localhost:9200/_cluster/health/?pretty=true
{
  "cluster_name" : "elasticsearch",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 9,
  "active_shards" : 9,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 10,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 47.368421052631575
}
sh-5.0#curl -X DELETE 'http://localhost:9200/ind-1?pretty'
{
  "acknowledged" : true
}
sh-5.0# curl -X DELETE 'http://localhost:9200/ind-2?pretty'
{
  "acknowledged" : true
}
sh-5.0# curl -X DELETE 'http://localhost:9200/ind-3?pretty'
{
  "acknowledged" : true
}
sh-5.0#

```
    Индексы в статусе Yellow потому что у них указано число реплик, а по факту нет других серверов, соответсвено реплицировать некуда.

## Задача 3

В данном задании вы научитесь:
- создавать бэкапы данных
- восстанавливать индексы из бэкапов

Создайте директорию `{путь до корневой директории с elasticsearch в образе}/snapshots`.

Используя API [зарегистрируйте](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-register-repository.html#snapshots-register-repository)
данную директорию как `snapshot repository` c именем `netology_backup`.

**Приведите в ответе** запрос API и результат вызова API для создания репозитория.

```bash
[root@homesrv 6-5]# docker exec -it es01 /bin/sh
sh-5.0# curl -XPOST localhost:9200/_snapshot/netology_backup?pretty -H 'Content-Type: application/json' -d'{"type": "fs", "settings": { "location":"/var/lib/data/snapshots" }}'
{
  "acknowledged" : true
}
sh-5.0# curl -X GET 'http://localhost:9200/_snapshot/netology_backup?pretty'
{
  "netology_backup" : {
    "type" : "fs",
    "settings" : {
      "location" : "/var/lib/data/snapshots"
    }
  }
}
sh-5.0# curl -X PUT localhost:9200/_snapshot/netology_backup/elasticsearch?wait_for_completion=true
{"snapshot":{"snapshot":"elasticsearch","uuid":"p5VLjsw-Qbu8ErPuAT5Ufw","repository":"netology_backup","version_id":7170499,"version":"7.17.4","indices":[".ds-.logs-deprecation.elasticsearch-default-2022.06.03-000001","test",".ds-ilm-history-5-2022.06.03-000001"],"data_streams":["ilm-history-5",".logs-deprecation.elasticsearch-default"],"include_global_state":true,"state":"SUCCESS","start_time":"2022-06-03T16:47:58.587Z","start_time_in_millis":1654274878587,"end_time":"2022-06-03T16:47:58.787Z","end_time_in_millis":1654274878787,"duration_in_millis":200,"failures":[],"shards":{"total":3,"failed":0,"successful":3},"feature_states":[]}}sh-5.0#
sh-5.0# ls -l /var/lib/data/snapshots/
total 44
-rw-rw-r-- 1 elasticsearch root  1168 Jun  3 16:47 index-0
-rw-rw-r-- 1 elasticsearch root     8 Jun  3 16:47 index.latest
drwxrwxr-x 5 elasticsearch root    96 Jun  3 16:47 indices
-rw-rw-r-- 1 elasticsearch root 28823 Jun  3 16:47 meta-p5VLjsw-Qbu8ErPuAT5Ufw.dat
-rw-rw-r-- 1 elasticsearch root   625 Jun  3 16:47 snap-p5VLjsw-Qbu8ErPuAT5Ufw.dat
sh-5.0#
```

Создайте индекс `test` с 0 реплик и 1 шардом и **приведите в ответе** список индексов.

[Создайте `snapshot`](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-take-snapshot.html)
состояния кластера `elasticsearch`.

**Приведите в ответе** список файлов в директории со `snapshot`ами.

```bash
sh-5.0# curl -X PUT localhost:9200/test -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 1,  "number_of_replicas": 0 }}'
{"acknowledged":true,"shards_acknowledged":true,"index":"test"}sh-5.0#
sh-5.0# curl -X GET 'http://localhost:9200/_cat/indices?v'
health status index uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test  RXoVZDL0TWueNOBz7KWQeQ   1   0          0            0       226b           226b
sh-5.0#
```

Удалите индекс `test` и создайте индекс `test-2`. **Приведите в ответе** список индексов.

[Восстановите](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-restore-snapshot.html) состояние
кластера `elasticsearch` из `snapshot`, созданного ранее.

**Приведите в ответе** запрос к API восстановления и итоговый список индексов.

Подсказки:
- возможно вам понадобится доработать `elasticsearch.yml` в части директивы `path.repo` и перезапустить `elasticsearch`

```bash
sh-5.0# curl -X DELETE 'http://localhost:9200/test?pretty'
{
  "acknowledged" : true
}
sh-5.0# curl -X PUT localhost:9200/test-2?pretty -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 1,  "number_of_replicas": 0 }}'
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "test-2"
}
sh-5.0# curl -X GET http://localhost:9200/_cat/indices?v
health status index  uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test-2 hk-pMYi3SoWCfJ2IugFjtQ   1   0          0            0       226b           226b
sh-5.0# curl -X POST localhost:9200/_snapshot/netology_backup/elasticsearch/_restore?pretty -H 'Content-Type: application/json' -d'{"indices": "test"}'
{
  "accepted" : true
}
sh-5.0# curl -X GET http://localhost:9200/_cat/indices?v
health status index  uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test-2 hk-pMYi3SoWCfJ2IugFjtQ   1   0          0            0       226b           226b
green  open   test   wcH_La0_TpWSkx2W-_50FQ   1   0          0            0       226b           226b
sh-5.0#
```
