[fork]: images/fork.png "Fork"
[permissions]: images/permissions.png "Permissions"
[fork_done]: images/fork_done.png "Fork Done"
[members_access]: images/members_access.png "Members Access"
[members_adding]: images/members_adding.png "Members Adding"
[settings]: images/settings.png "Settings"
[full_name]: images/full_name.png "Full Name"
[ssh_keys_1]: images/ssh_keys_1.png "SSH Keys 1"
[ssh_keys_2]: images/ssh_keys_2.png "SSH Keys 2"
[clone]: images/clone.png "Clone"
[merge_request_1]: images/merge_request_1.png "Merge Request 1"
[merge_request_2]: images/merge_request_2.png "Merge Request 2"

# Программирование на Go

[TOC]

## Подготовка
Необходимо форкнуть репу, сделать ее приватной и добавить следующих пользователей

> @vbabayan @The_Psina @borenko.nastya

с правами **Developer**.

### Fork
Все задания домашних работ лежат в общем открытом репозитории. Для сдачи дз необходимо завести его копию на свой аккаунт.
В этой копии исходного репозитория вы и будете выполнять ДЗ. Для этого на основной странице репозитория необходимо нажать кнопку "Fork":

![alt text][fork]

### Permissions
Название и описание проекта можно оставить по-умолчанию.
1. "Project URL" - **обязательно** имя вашего пользователя GitLab;
2. "Branches to include" - **обязательно** "All branches";
3. "Visibility level" - **обязательно** "Private";
4. Нажатие "Fork project" приведет к созданию личной копии репозитория на вашем аккаунте.

![alt text][permissions]
![alt text][fork_done]

### Members:
Необходимо добавить менторам бессрочный доступ к репозиторию с правами "Developer".

![alt text][members_access]

1. Необходимо добавить в список пользователей, указанных в разделе [Подготовка](#подготовка);
2. "Select a role" - **обязательно** "Developer";
3. "Access expiration date" - **обязательно** должно быть пустым;
4. Жмяк на "Invite"

![alt text][members_adding]

### Настройка аккаунта GitLab
Для того чтобы нам было понятно, кто является автором ДЗ, необходимо установить свое имя в настройках аккаунта GitLab.

![alt text][settings]
![alt text][full_name]

### Настройка доступа в репозиторий на GitLab

Для того чтобы иметь возможность сохранять изменения кода на вашем компьютере на GitLab, нужно
"сообщить" гитлабу каким образом идентифицировать ваш компьютер.

Для этого существует понятие ssh-ключей.

ssh-ключ служит одновременно для идентификации и установления безопасного соединения. Любой ssh-ключ состоит из двух частей публичной и приватной:

- Публичный ключ свободно распространяется и с помощью него шифруются сообщения
- Приватный ключ хранится закрыто, и с помощью него получатель расшифровывает сообщения.

Для генерации пары открытый/закрытый ssh-ключ исполните следующие команды.

[Как использовать команды, подобные той, что ниже](#использование-кодовых-вставок)

```bash
$ mkdir -p ~/.ssh
$ chmod 700 ~/.ssh
$ ssh-keygen -t rsa
Generating public/private rsa key pair.
Enter file in which to save the key (/Users/p.bereznoy/.ssh/id_rsa):
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /tmp/test.
Your public key has been saved in /tmp/test.pub.
The key fingerprint is:
SHA256:9Yf53peICFiHtsvlLTU+lnMidP9Vd7U7aWtSE9GWyIs p.bereznoy@msk-wifi-21fap7-p_berezhnoy-noname.mail.msk
The keys randomart image is:
+---[RSA 2048]----+
|            . . o|
|             o oo|
|        . . . ..o|
|       + o E + .o|
|      + S   + ..=|
|     . o o + o o*|
|      . * * = +=+|
|       o = X *.+=|
|          + = =oo|
+----[SHA256]-----+
```

В результате работы сгенерируются файлы с публичным и приватным ключами:

Если вы не уверены, что готовы потратить пару часов на разбирательство "почему оно не работает", **не меняйте
путь к ключу**, который предлагает утилита (строка "Enter file in which to save the key"), просто нажмите Enter.

Если вам не нужен пароль для ключа (а большинству он не понадобится), тоже просто нажмите Enter.

В результате у вас были сгенерированы 2 файла:
```bash
$ ls ~/.ssh
id_rsa id_rsa.pub
```

Необходимо скопировать содержимое файла **id__rsa.pub** на гитлаб:
```bash
$ cat ~/.ssh/id_rsa.pub
ssh-rsa
AAAAB3NzaC1yc2EAAAADAQABAAABAQD9nu0UpP/5txdI9CkOkIj3N4A0wdQ8Skm1s6mISmPmq6efOLJH5JEJ3oEOWvFBZOGMzR0QfJ9UOWy02/+YEXAJ9hMKoenaKHovTXhL6i9T99bD9TDouWh9kR4XbDht2pcmEzRkvgKh+xSwqDt7IwShdQtBr93j9H/z5pL38mKOz98TLGEBXDJMOH0QGHk/FPRiVGQl6HxNOa7wGzYR1fMgWMK5qX6S/81dRMOWjgm3QvpUiNwk3POhkLcO5YOV+H3zxb65KzDXixScQBRBWGUqKzc2qoyoG84m7LirGHc5moH+q5Ieo+nC5l0NOd3sKqq5XL5L2ZmNoErM2WVQZKnz
<username>@<host>
```

Далее необходимо скопировать ключ на гитлаб:

![alt text][ssh_keys_1]

1. Необходимо поместить **публичный** ключ в поле "Key";
2. Название ключа будет проставлено автоматически, его можно не менять
3. "Usage Type" - **обязательно** "Authentication & Signing"
4. Если вы хотите, чтобы у ключа не было времени протухания, оставьте "Expiration Date пустым"
5. Жмяк "Add key"

![alt text][ssh_keys_2]

### Клонирование репозитория

Для того чтобы получить копию своего репозитория на компьютер, сначала скопируйте его URL (Clone with SSH не потребует ввода имени пользователя и пароля GitLab, в отличие от Clone with HTTPS):

![alt text][clone]

Затем воспользуйтесь командой:
```bash
$ git clone "URL"
```

### Локальные настройки GIT'а

Просим вас исполнить следующие команды для конфигурирования вашего локального гита.

Настройки принадлежности ваших коммитов:
```bash
$ git config --global user.name "Your Name Surname"
$ git config --global user.email "your@email"
```

Если указана опция --global, то эти настройки достаточно сделать только один раз, поскольку в этом случае Git будет
использовать эти данные для всего, что вы делаете в этой системе. Если для каких-то отдельных проектов вы хотите указать
другое имя или электронную почту, можно выполнить эту же команду без параметра --global в каталоге с нужным проектом.

Следующая команда необходима для того, чтобы пушить на гитлаб только текущую ветку.
Необходимо **ВСЕМ**.
```bash
$ git config --global push.default current
```

Некоторые цветовые настройки. Будет красиво:
```bash
$ git config --global color.branch true
$ git config --global color.interactive true
$ git config --global color.diff true
$ git config --global color.grep true
$ git config --global color.status true
```

## Выполнение
Каждое задание находится в отдельной ветке с именем **hw-\<num\>**, где _num_ - номер задания.
Описание задания находится в каждой ветке в *README.md*: его нужно выполнить, закоммитить и отправить на сервер.
На каждый пуш GitLab запустит CI, который: проверит стиль кода, соберет проект, прогонит различные тесты.

**Note: все линтеры, сборщики и тесты находятся в той же ветке, что и задание. Их можно и нужно запустить локально,
перед отправкой на сервер**

Перед тем как начать делать задание, вам необходимо создать новую ветку с любым удобным именем
(например, **making-hw-\<num\>**, _num_ - номер задания) от ветки **hw-\<num\>**:
```bash
$ git checkout hw-1
$ git push
$ git checkout -b making-hw-1
```

После этого нужно делать домашнее задание в этой ветке.

## Защита
После успешного прогона CI вы можете отправить задание на проверку.
Для этого нужно сделать **Merge Request** вашей ветки (**making-hw-\<num\>** в примере)
в ветку **hw-\<num\>** (_num_ - номер задания).
Через какое-то время свободный ментор возьмет его на проверку. Если этого не произошло - сообщите нам.

**Note: Merge Request необходимо делать в свой форк!**

На все замечания будут открыты дискуссии и выставлена предварительная оценка. Как будут запушены все исправления - пишите в комменты.

**Note: выполненным и защищенным будет считаться задание со смерженным ментором MR'ом и закрытыми дискуссиями.**
Если вы сами закроете дискуссии или сами вмержите код, дз оценка не будет учтена. Ее все равно проставляет ментор.

### Merge Request:
![alt text][merge_request_1]
![alt text][merge_request_2]


## Получение следующих дз

Для того чтобы новые дз появлялись в вашем форке, воспользуйтесь ручным зеркалированием.

Добавьте общий репозиторий в список удаленных репозиториев (делается один раз):
```bash
$ git remote add upstream "URL"
```

Для добавления всех изменений в локальный репозиторий:
```bash
$ git fetch --all
```

Для добавления ветки в ваш репозиторий:
```bash
$ git checkout hw-num
$ git push origin
```

## Дополнения

### Использование кодовых вставок
Все кодовые вставки в данном документе подчиняются следующим правилам:

#### Если перед строкой стоит знак **_$_**, значит это команда, которая запускается в терминале. Пример:
```bash
$ ls
$ ./test
```

Здесь выполняются 2 команды: **_ls_** и **_./test_**.
Чтобы выполнить их у себя на компьютере, необходимо скопировать все, что находится за символом **_$_**.

#### Если перед строкой не стоит символа **_$_**, значит это вывод программы, которая была выполнена выше. Пример:
```bash
$ ls
test 123 somedir
```

Здесь была выполнена команда **_ls_**, которая выдала в терминал текст **_test 123 somedir_**. Эти строки даны для
ознакомления с примером вывода программ.

### Полезные ссылки
- [Статья о том, как пользоваться командной строкой linux](https://selectel.ru/blog/basic-linux-commands/)
- [Git Book](https://git-scm.com/book/en/v2)
