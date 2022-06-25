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
