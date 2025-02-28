# Домашнее задание к занятию "7.6. Написание собственных провайдеров для Terraform."

Бывает, что
* общедоступная документация по терраформ ресурсам не всегда достоверна,
* в документации не хватает каких-нибудь правил валидации или неточно описаны параметры,
* понадобиться использовать провайдер без официальной документации,
* может возникнуть необходимость написать свой провайдер для системы используемой в ваших проектах.   

## Задача 1.
Давайте потренируемся читать исходный код AWS провайдера, который можно склонировать от сюда:
[https://github.com/hashicorp/terraform-provider-aws.git](https://github.com/hashicorp/terraform-provider-aws.git).
Просто найдите нужные ресурсы в исходном коде и ответы на вопросы станут понятны.  


1. Найдите, где перечислены все доступные `resource` и `data_source`, приложите ссылку на эти строки в коде на
гитхабе.   

    > [resource](https://github.com/hashicorp/terraform-provider-aws/blob/caf5a742745561d36e6bd6c3032f7420e31f3518/internal/provider/provider.go#L909)
    >
    > [data_source](https://github.com/hashicorp/terraform-provider-aws/blob/caf5a742745561d36e6bd6c3032f7420e31f3518/internal/provider/provider.go#L425)

1. Для создания очереди сообщений SQS используется ресурс `aws_sqs_queue` у которого есть параметр `name`.
    * С каким другим параметром конфликтует `name`? Приложите строчку кода, в которой это указано.
        > [ConflictsWith: []string{"name_prefix"},](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/service/sqs/queue.go#L87)

    * Какая максимальная длина имени?
    * Какому регулярному выражению должно подчиняться имя?
        > [Максимальная длина имени с расширением ](https://github.com/hashicorp/terraform-provider-aws/blob/167536a0a72cd6294c7bd3eed85d36232e0d2ef5/internal/service/sqs/queue.go#L425)
        >
        > [Максимальная длина имени без раширения](https://github.com/hashicorp/terraform-provider-aws/blob/167536a0a72cd6294c7bd3eed85d36232e0d2ef5/internal/service/sqs/queue.go#L427)

## Задача 2. (Не обязательно)
В рамках вебинара и презентации мы разобрали как создать свой собственный провайдер на примере кофемашины.
Также вот официальная документация о создании провайдера:
[https://learn.hashicorp.com/collections/terraform/providers](https://learn.hashicorp.com/collections/terraform/providers).

1. Проделайте все шаги создания провайдера.
2. В виде результата приложение ссылку на исходный код.
3. Попробуйте скомпилировать провайдер, если получится то приложите снимок экрана с командой и результатом компиляции.   

```shell
[sergej@fedora GIT_SORE]$ git clone https://github.com/hashicorp/learn-terraform-hashicups-provider && cd learn-terraform-hashicups-provider
Клонирование в «learn-terraform-hashicups-provider»…
remote: Enumerating objects: 41, done.
remote: Counting objects: 100% (41/41), done.
remote: Compressing objects: 100% (32/32), done.
remote: Total 41 (delta 16), reused 25 (delta 8), pack-reused 0
Получение объектов: 100% (41/41), 6.52 КиБ | 6.52 МиБ/с, готово.
Определение изменений: 100% (16/16), готово.
[sergej@fedora learn-terraform-hashicups-provider]$ cd docker_compose && docker compose up -d
[+] Running 22/22
 ⠿ db Pulled                                                                                                                                                            11.3s
   ⠿ d599a449871e Pull complete                                                                                                                                          3.1s
   ⠿ 2ddf9fad2006 Pull complete                                                                                                                                          3.2s
   ⠿ a60892c15ec0 Pull complete                                                                                                                                          3.3s
   ⠿ 75e03bfe93e9 Pull complete                                                                                                                                          3.4s
   ⠿ 62558287f90b Pull complete                                                                                                                                          3.6s
   ⠿ f7f7c2564b8f Pull complete                                                                                                                                          3.8s
   ⠿ 43b55e177d25 Pull complete                                                                                                                                          4.2s
   ⠿ 6cdb6c6e946b Pull complete                                                                                                                                          4.6s
   ⠿ d210255df7e3 Pull complete                                                                                                                                          8.2s
   ⠿ fbb4c12b8149 Pull complete                                                                                                                                          8.3s
   ⠿ 69b1096dd930 Pull complete                                                                                                                                          8.3s
   ⠿ 73f3b8fa196b Pull complete                                                                                                                                          8.4s
   ⠿ 7720636b16fa Pull complete                                                                                                                                          8.4s
   ⠿ db3d7e34d278 Pull complete                                                                                                                                          8.5s
   ⠿ 3f3858a782f5 Pull complete                                                                                                                                          8.6s
   ⠿ bdc6e76bfff4 Pull complete                                                                                                                                          8.6s
 ⠿ api Pulled                                                                                                                                                            4.9s
   ⠿ cbdbe7a5bc2a Pull complete                                                                                                                                          1.3s
   ⠿ 8246d5ae59e2 Pull complete                                                                                                                                          1.3s
   ⠿ 3f1e85df6f1c Pull complete                                                                                                                                          1.4s
   ⠿ ff8d97847668 Pull complete                                                                                                                                          2.2s
[+] Running 3/3
 ⠿ Network docker_compose_default  Created                                                                                                                               0.2s
 ⠿ Container docker_compose-db-1   Started                                                                                                                               1.1s
 ⠿ Container docker_compose-api-1  Started                                                                                                                               1.2s
[sergej@fedora docker_compose]$ curl localhost:19090/health
ok
[sergej@fedora docker_compose]cd ..
[sergej@fedora learn-terraform-hashicups-provider]$ cd ..
[sergej@fedora GIT_SORE]$ curl -LO https://github.com/hashicorp/terraform-provider-hashicups/releases/download/v0.3.1/terraform-provider-hashicups_0.3.1_linux_amd64.zip
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
100 5060k  100 5060k    0     0  3321k      0  0:00:01  0:00:01 --:--:-- 12.1M
[sergej@fedora GIT_SORE]$ mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/linux_amd64
[sergej@fedora GIT_SORE]$ unzip terraform-provider-hashicups_0.3.1_linux_amd64.zip -d ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/linux_amd64
Archive:  terraform-provider-hashicups_0.3.1_linux_amd64.zip
  inflating: /home/sergej/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/linux_amd64/README.md  
  inflating: /home/sergej/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/linux_amd64/terraform-provider-hashicups_v0.3.1  
[sergej@fedora GIT_SORE]$ curl -X POST localhost:19090/signup -d '{"username":"education", "password":"test123"}'
{"UserID":1,"Username":"education","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTYzOTEzNDcsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.IDWa6BU-uROIdkQpaoClIZv1SaC9zv7SFGVJHVxs-N0"}
[sergej@fedora GIT_SORE]$ curl -X POST localhost:19090/signin -d '{"username":"education", "password":"test123"}'
{"UserID":1,"Username":"education","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTYzOTEzNjAsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.lfgFfJ1gmxaaNvyG4yKSHn5uwRzVrTYah4K7n3x0rEw"}
[sergej@fedora GIT_SORE]$ export HASHICUPS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTYzOTEzNjAsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.lfgFfJ1gmxaaNvyG4yKSHn5uwRzVrTYah4K7n3x0rEw
[sergej@fedora learn-terraform-hashicups-provider]$ chmod +x ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/linux_amd64/terraform-provider-hashicups_v0.3.1
[sergej@fedora learn-terraform-hashicups-provider]$ cd ~
[sergej@fedora ~]$ ll
итого 124
drwxr-xr-x.  2 sergej sergej  4096 ноя 12  2021  Downloads
drwxr-xr-x   3 sergej sergej  4096 июн 21 11:43  Games
drwxr-xr-x. 14 sergej sergej  4096 июн 27 09:41  GIT_SORE
drwxrwxr-x   4 sergej sergej  4096 июн 23 09:37  go
-rwxrwxr-x   1 sergej sergej  7658 мая 18 13:06  PortProton_1.0
drwxrwxr-x   3 sergej sergej  4096 мая 18 13:24  PortWINE
drwxr-xr-x.  2 sergej sergej  4096 мар 15 14:02  python
drwx------.  3 sergej sergej  4096 апр 21 14:28  snap
drwxrwxr-x.  2 sergej sergej  4096 июн 27 07:44  Telegram
drwx------.  3 sergej sergej  4096 апр 12 13:36 'VirtualBox VMs'
drwxrwxr-x.  4 sergej sergej  4096 апр 21 13:58  yandex-cloud
-rw-------.  1 sergej sergej  2602 июн  7 10:42  yes
-rw-r--r--.  1 sergej sergej   567 июн  7 10:42  yes.pub
drwxr-xr-x.  3 sergej sergej  4096 апр 19 15:55  Видео
drwxr-xr-x.  6 sergej sergej  4096 июн  7 16:19  Документы
drwxr-xr-x. 10 sergej sergej 20480 июн 23 09:58  Загрузки
drwxr-xr-x.  4 sergej sergej 24576 июн 22 13:57  Изображения
drwxr-xr-x.  2 sergej sergej  4096 апр 18 15:35  Музыка
drwxr-xr-x.  2 sergej sergej  4096 апр 18 15:35  Общедоступные
drwxr-xr-x.  4 sergej sergej  4096 июн 23 08:15 'Рабочий стол'
drwxr-xr-x.  2 sergej sergej  4096 апр 18 15:35  Шаблоны
[sergej@fedora ~]$ cd .terraform.d/
[sergej@fedora .terraform.d]$ ll
итого 16
-rw-r--r--  1 sergej docker  310 июн 27 09:46 checkpoint_cache
-rw-r--r--. 1 sergej sergej  394 апр 21 15:55 checkpoint_signature
-rw-------  1 sergej sergej  166 июн 21 12:48 credentials.tfrc.json
drwxr-xr-x  3 sergej docker 4096 июн 27 09:41 plugins
[sergej@fedora .terraform.d]$ curl -X POST localhost:19090/signin -d '{"username":"education", "password":"test123"}'
{"UserID":1,"Username":"education","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTYzOTI2NjUsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.5_jdxrkN5tTaRJg_8DsI5NWp6EcuyWrvZ2edd-W2rOo"}
[sergej@fedora .terraform.d]$ export HASHICUPS_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTYzOTEzNjAsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.lfgFfJ1gmxaaNvyG4yKSHn5uwRzVrTYah4K7n3x0rEw"
[sergej@fedora learn-terraform-hashicups-provider]$ cd ~/
[sergej@fedora ~]$ mkdir code
[sergej@fedora ~]$ cd ~/code
[sergej@fedora code]$ git clone https://github.com/hashicorp/terraform-provider-hashicups
Клонирование в «terraform-provider-hashicups»…
remote: Enumerating objects: 3502, done.
remote: Total 3502 (delta 0), reused 0 (delta 0), pack-reused 3502
Получение объектов: 100% (3502/3502), 71.46 МиБ | 20.99 МиБ/с, готово.
Определение изменений: 100% (886/886), готово.
[sergej@fedora code]$ cd terraform-provider-hashicups
[sergej@fedora terraform-provider-hashicups]$ go mod tidy
go: downloading github.com/hashicorp/terraform-plugin-docs v0.7.0
go: downloading github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.0-rc.2
go: downloading github.com/hashicorp-demoapp/hashicups-client-go v0.0.0-20200508203820-4c67e90efb8e
go: downloading github.com/mattn/go-colorable v0.1.12
go: downloading github.com/hashicorp/go-hclog v0.9.2
go: downloading github.com/hashicorp/go-plugin v1.3.0
go: downloading google.golang.org/grpc v1.27.1
go: downloading github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
go: downloading github.com/hashicorp/go-multierror v1.1.1
go: downloading github.com/hashicorp/go-uuid v1.0.1
go: downloading github.com/mitchellh/copystructure v1.2.0
go: downloading github.com/mitchellh/reflectwalk v1.0.2
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/hashicorp/terraform-json v0.13.0
go: downloading github.com/hashicorp/terraform-plugin-test v1.4.0
go: downloading github.com/mitchellh/go-testing-interface v1.0.4
go: downloading github.com/mitchellh/mapstructure v1.1.2
go: downloading github.com/mitchellh/cli v1.1.2
go: downloading golang.org/x/net v0.0.0-20210326060303-6b1517762897
go: downloading github.com/golang/protobuf v1.3.4
go: downloading github.com/mattn/go-isatty v0.0.14
go: downloading github.com/go-test/deep v1.0.3
go: downloading github.com/stretchr/testify v1.7.0
go: downloading github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d
go: downloading github.com/oklog/run v1.0.0
go: downloading github.com/jhump/protoreflect v1.6.0
go: downloading github.com/hashicorp/hcl/v2 v2.3.0
go: downloading github.com/zclconf/go-cty v1.10.0
go: downloading github.com/apparentlymart/go-dump v0.0.0-20190214190832-042adf3cf4a0
go: downloading github.com/hashicorp/errwrap v1.0.0
go: downloading golang.org/x/text v0.3.5
go: downloading github.com/hashicorp/logutils v1.0.0
go: downloading github.com/hashicorp/go-version v1.4.0
go: downloading github.com/hashicorp/go-getter v1.4.2-0.20200106182914-9813cbd4eb02
go: downloading github.com/hashicorp/hc-install v0.3.1
go: downloading github.com/hashicorp/terraform-exec v0.16.0
go: downloading github.com/Masterminds/sprig v2.22.0+incompatible
go: downloading github.com/armon/go-radix v0.0.0-20180808171621-7fddfc383310
go: downloading github.com/bgentry/speakeasy v0.1.0
go: downloading github.com/fatih/color v1.7.0
go: downloading github.com/posener/complete v1.1.1
go: downloading github.com/vmihailenco/msgpack v4.0.1+incompatible
go: downloading google.golang.org/genproto v0.0.0-20200310143817-43be25429f5a
go: downloading golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6
go: downloading github.com/agext/levenshtein v1.2.2
go: downloading github.com/apparentlymart/go-textseg v1.0.0
go: downloading github.com/mitchellh/go-wordwrap v1.0.0
go: downloading github.com/kr/pretty v0.2.1
go: downloading github.com/kylelemons/godebug v1.1.0
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
go: downloading cloud.google.com/go v0.45.1
go: downloading github.com/aws/aws-sdk-go v1.25.3
go: downloading github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d
go: downloading github.com/hashicorp/go-cleanhttp v0.5.2
go: downloading github.com/hashicorp/go-safetemp v1.0.0
go: downloading github.com/mitchellh/go-homedir v1.1.0
go: downloading github.com/ulikunitz/xz v0.5.5
go: downloading google.golang.org/api v0.9.0
go: downloading github.com/hashicorp/go-checkpoint v0.5.0
go: downloading github.com/russross/blackfriday v1.6.0
go: downloading github.com/Masterminds/goutils v1.1.0
go: downloading github.com/Masterminds/semver v1.5.0
go: downloading github.com/google/uuid v1.1.2
go: downloading github.com/huandu/xstrings v1.3.2
go: downloading github.com/imdario/mergo v0.3.12
go: downloading golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
go: downloading google.golang.org/appengine v1.6.5
go: downloading gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
go: downloading github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
go: downloading github.com/apparentlymart/go-textseg/v13 v13.0.0
go: downloading github.com/kr/text v0.2.0
go: downloading github.com/googleapis/gax-go/v2 v2.0.5
go: downloading golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
go: downloading gopkg.in/yaml.v2 v2.3.0
go: downloading go.opencensus.io v0.22.0
go: downloading github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
go: downloading github.com/go-git/go-git/v5 v5.4.2
go: downloading github.com/google/martian v2.1.0+incompatible
go: downloading github.com/hashicorp/golang-lru v0.5.1
go: downloading github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7
go: downloading github.com/go-git/go-billy/v5 v5.3.1
go: downloading github.com/sergi/go-diff v1.2.0
go: downloading github.com/emirpasic/gods v1.12.0
go: downloading github.com/acomagu/bufpipe v1.0.3
go: downloading github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99
go: downloading github.com/go-git/gcfg v1.5.0
go: downloading github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351
go: downloading github.com/xanzy/ssh-agent v0.3.0
go: downloading gopkg.in/warnings.v0 v0.1.2
go: downloading github.com/Microsoft/go-winio v0.4.16
[sergej@fedora terraform-provider-hashicups]$ go build -o terraform-provider-hashicups
[sergej@fedora terraform-provider-hashicups]$ mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/[OS_ARCH]
[sergej@fedora terraform-provider-hashicups]$ mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.3.1/[OS_ARCH]
[sergej@fedora terraform-provider-hashicups]$ cd ~/GIT_SORE/learn-terraform-hashicups-provider/
[sergej@fedora learn-terraform-hashicups-provider]$ terraform apply

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # hashicups_order.edu will be created
  + resource "hashicups_order" "edu" {
      + id           = (known after apply)
      + last_updated = (known after apply)

      + items {
          + quantity = 2

          + coffee {
              + description = (known after apply)
              + id          = 3
              + image       = (known after apply)
              + name        = (known after apply)
              + price       = (known after apply)
              + teaser      = (known after apply)
            }
        }
      + items {
          + quantity = 2

          + coffee {
              + description = (known after apply)
              + id          = 2
              + image       = (known after apply)
              + name        = (known after apply)
              + price       = (known after apply)
              + teaser      = (known after apply)
            }
        }
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + edu_order = {
      + id           = (known after apply)
      + items        = [
          + {
              + coffee   = [
                  + {
                      + description = (known after apply)
                      + id          = 3
                      + image       = (known after apply)
                      + name        = (known after apply)
                      + price       = (known after apply)
                      + teaser      = (known after apply)
                    },
                ]
              + quantity = 2
            },
          + {
              + coffee   = [
                  + {
                      + description = (known after apply)
                      + id          = 2
                      + image       = (known after apply)
                      + name        = (known after apply)
                      + price       = (known after apply)
                      + teaser      = (known after apply)
                    },
                ]
              + quantity = 2
            },
        ]
      + last_updated = (known after apply)
    }

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

hashicups_order.edu: Creating...
hashicups_order.edu: Creation complete after 0s [id=1]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.

Outputs:

edu_order = {
  "id" = "1"
  "items" = tolist([
    {
      "coffee" = tolist([
        {
          "description" = ""
          "id" = 3
          "image" = "/nomad.png"
          "name" = "Nomadicano"
          "price" = 150
          "teaser" = "Drink one today and you will want to schedule another"
        },
      ])
      "quantity" = 2
    },
    {
      "coffee" = tolist([
        {
          "description" = ""
          "id" = 2
          "image" = "/vault.png"
          "name" = "Vaulatte"
          "price" = 200
          "teaser" = "Nothing gives you a safe and secure feeling like a Vaulatte"
        },
      ])
      "quantity" = 2
    },
  ])
  "last_updated" = tostring(null)
}
[sergej@fedora learn-terraform-hashicups-provider]$ terraform state show hashicups_order.edu
# hashicups_order.edu:
resource "hashicups_order" "edu" {
    id = "1"

    items {
        quantity = 2

        coffee {
            id     = 3
            image  = "/nomad.png"
            name   = "Nomadicano"
            price  = 150
            teaser = "Drink one today and you will want to schedule another"
        }
    }
    items {
        quantity = 2

        coffee {
            id     = 2
            image  = "/vault.png"
            name   = "Vaulatte"
            price  = 200
            teaser = "Nothing gives you a safe and secure feeling like a Vaulatte"
        }
    }
}
[sergej@fedora learn-terraform-hashicups-provider]$ curl -X GET  -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/1
{"id":1,"items":[{"coffee":{"id":3,"name":"Nomadicano","teaser":"Drink one today and you will want to schedule another","description":"","price":150,"image":"/nomad.png","ingredients":null},"quantity":2},{"coffee":{"id":2,"name":"Vaulatte}}}
[sergej@fedora learn-terraform-hashicups-provider]$
```

<p align="center">
  <img width="1200" height="600" src="./img/scrin.png">
</p>
