# Kubernetes

[Wikipedia](https://ru.wikipedia.org/wiki/Kubernetes):

**Kubernetes** (от др. -греч. *κυβερνήτης* — «*кормчий*», «*рулевой*», часто 
также используется нумероним *k8s*) — открытое программное обеспечение для 
[*оркестровки*](https://ru.wikipedia.org/wiki/%D0%9E%D1%80%D0%BA%D0%B5%D1%81%D1%82%D1%80%D0%BE%D0%B2%D0%BA%D0%B0_(%D0%98%D0%A2)) 
контейнеризированных приложений — автоматизации их развёртывания, 
масштабирования и координации в условиях кластера.

Задачи **Kubernetes**:

- Деплой и управление приложениями (**контейнерами**)
- Масштабирование и сужение согласно текущим требованиям
- Деплой с нулевым временем простоя (*zero downtime*)
- Откаты
- И многое другое

## Источники

- [Mastering Kubernetes. Master k8s from A to Z](https://amigoscode.com/p/kubernetes)

## Оглавление

- [Kubernetes](#kubernetes)
  - [Источники](#источники)
  - [Оглавление](#оглавление)
  - [Понятие кластера (Cluster)](#понятие-кластера-cluster)
  - [Master-нода](#master-нода)
    - [Control Plane](#control-plane)
    - [API Server](#api-server)
    - [Cluster Store (etcd)](#cluster-store-etcd)
    - [Scheduler](#scheduler)
    - [Controller Manager](#controller-manager)
    - [Cloud Controller Manager](#cloud-controller-manager)
    - [Подытог сведений о master-ноде](#подытог-сведений-о-master-ноде)
  - [Worker-ноды](#worker-ноды)
    - [Kubelet](#kubelet)
    - [Среда выполнения контейнера (Container Runtime)](#среда-выполнения-контейнера-container-runtime)
    - [Kube Proxy](#kube-proxy)
  - [Запуск Kubernetes](#запуск-kubernetes)
    - [Что значит *управляемые* **Kubernetes**](#что-значит-управляемые-kubernetes)
      - [Пример на основе **Amazon EKS**](#пример-на-основе-amazon-eks)
    - [Запуск кластера локально](#запуск-кластера-локально)
  - [Minikube](#minikube)
    - [Запуск кластера Minikube](#запуск-кластера-minikube)
    - [Остановка кластера Minikube](#остановка-кластера-minikube)
    - [Удаление кластера Minikube](#удаление-кластера-minikube)
    - [Статус Minikube](#статус-minikube)
    - [IP-адрес master-ноды Minikube](#ip-адрес-master-ноды-minikube)
    - [Запуск Minikube с несколькими нодами](#запуск-minikube-с-несколькими-нодами)
    - [Просмотр логов Minikube](#просмотр-логов-minikube)
      - [Просмотр логов с текущего момента](#просмотр-логов-с-текущего-момента)
      - [Просмотр логов определённой ноды](#просмотр-логов-определённой-ноды)
  - [Kubectl](#kubectl)
    - [Создание и запуск **пода** напрямую (императивный путь)](#создание-и-запуск-пода-напрямую-императивный-путь)
    - [Запуск **пода** с помощью конфигурационного файла (декларативный путь)](#запуск-пода-с-помощью-конфигурационного-файла-декларативный-путь)
    - [Получение списка **подов**](#получение-списка-подов)
      - [Слежение за **подами**](#слежение-за-подами)
      - [Получение списка всех **подов**](#получение-списка-всех-подов)
      - [Получения списка **подов** определённого пространства имён (неймспейса)](#получения-списка-подов-определённого-пространства-имён-неймспейса)
    - [Получение списка ***всего***](#получение-списка-всего)
    - [Получение расширенной информации о **поде**](#получение-расширенной-информации-о-поде)
      - [Выбор формата вывода](#выбор-формата-вывода)
      - [Получение полной информации](#получение-полной-информации)
    - [Получения списка пространств имён (неймспейсов)](#получения-списка-пространств-имён-неймспейсов)
    - [Получение списка **Deployments**](#получение-списка-deployments)
      - [Получение полной информации](#получение-полной-информации-1)
    - [Получение списка **ReplicaSet**](#получение-списка-replicaset)
      - [Получение полной информации](#получение-полной-информации-2)
    - [Проброс порта на хост](#проброс-порта-на-хост)
    - [Удаление **пода**](#удаление-пода)
    - [Дебаг **подов**](#дебаг-подов)
      - [Просмотр логов](#просмотр-логов)
      - [Просмотр логов для определённого контейнера](#просмотр-логов-для-определённого-контейнера)
      - [`exec` в контейнерах](#exec-в-контейнерах)
    - [Просмотр ресурсов API](#просмотр-ресурсов-api)
    - [Исследование кластера](#исследование-кластера)
      - [Подключение к ноде по SSH](#подключение-к-ноде-по-ssh)
  - [**Поды** (**Pods**)](#поды-pods)
    - [Способы создания **подов**](#способы-создания-подов)
      - [Императивный способ](#императивный-способ)
      - [Декларативный способ](#декларативный-способ)
    - [Deployments](#deployments)
    - [ReplicaSet](#replicaset)
      - [Контуры управления (Control loops)](#контуры-управления-control-loops)
      - [Rolling Update](#rolling-update)
      - [Предыдущий **ReplicaSet** после *Rolling Update*](#предыдущий-replicaset-после-rolling-update)
      - [Откат версии](#откат-версии)
      - [Информация о ревизии (версии)](#информация-о-ревизии-версии)

## Понятие кластера (Cluster)

**Кластер** (**Cluster**) – набор **нод** (**node**).

**Нода** (**Node**) – виртуальная (**VM**) или физическая машина.

Это всё может запускаться на облачных платформах таких как *AWS*, *Azure*, 
*Google Cloud* и т.д.

![](images/cluster.png)

## Master-нода

### Control Plane

![](images/control-plane.png)

### API Server

**API Server** выступает фронтендом для Control Plane. Все коммуникации 
(внутренние и внешние) проходят через **API Server**.

> **API Server** Открывает порт 443 с RESTful API.
> 
> Также он выполняет проверки аутентификации и авторизации.

### Cluster Store (etcd)

**Cluster Store** хранит весь стейт (состояние) нашего приложения, а также 
конфигурацию. Выступает также распределённым хранилищем *ключ-значение* и 
отдельным [single source of truth](https://en.wikipedia.org/wiki/Single_source_of_truth).

[Официальный сайт etcd](https://etcd.io/).

### Scheduler

**Scheduler** следит за новой нагрузкой/pods и назначает их к нодам по 
отдельным факторам расписания:

- Здорова ли нода? (Is healthy?)
- Достаточно ли ей ресурсов?
- Доступен ли порт?
- Affinity- и Anti Affinity-правила
- и другое

### Controller Manager

**Controller Manager** – это демон, который управляет контуром управления 
(control loop). Это *контроллер контроллеров*.

> ![](images/node-controller.png)
> 
> Другие контроллеры:
> 
> - ReplicaSet. Контроллер, ответственный за обеспечение корректного числа 
>   запущенных подов (pods)
> - Endpoint. Назначает порты сервисам.
> - Namespace.
> - Service Account.

Каждый контроллер следит за изменениями через **API Server**.

Цель: следить за изменениями, которые не совпадают с ожидаемым стейтом 
(состоянием) (if desire state doesn't match the current state). Если не 
совпадают, то "пнуть" соответствующий контроллер, чтобы он попытался исправить 
это.

Контроллер, по-сути, это watch loop.

### Cloud Controller Manager

**Cloud Controller Manager** ответственен за взаимодействие с облачным 
провайдером.

### Подытог сведений о master-ноде

Master-нода запускает все сервисы Control Plane.

## Worker-ноды

**Worker-нода** – это виртуальная или физическая машина (часто под управлением 
 Linux). Предоставляет среду для запуска приложений.

![](images/worker-node.png)

### Kubelet

**Kubelet** – это основной агент, который запускается в каждой отдельной ноде. 
Он получает определение *подов* (*Pods*) с **API Server**. Взаимодействует со 
средой выполнения контейнера (Container Runtime), чтобы запускать контейнеры, 
связанные с подами. Отправляет отчёты нод и подов в Master.

### Среда выполнения контейнера (Container Runtime)

**Среда выполнения контейнера** (**Container Runtime**) ответственна за пулл 
образов (images) с реестров, таких как **Docker Hub**, **Google Container 
Registry**, **Amazon ECR**, **Azure Container Registry** и т.д. Также она 
отвечает за запуск и остановку контейнеров с этих образов, абстрагируя 
управление контейнерами для **Kubernetes**.

Среда также предоставляет **C**ontainer **R**untime **I**nterface (**CRI**), 
который является интерфейсом для *сторонних* (*3rd party*) сред выполнения 
контейнеров.

> **Kubernetes** с какого-то начал использовать по умолчанию **Containerd** 
> вместо Docker. [Подробнее](https://habr.com/ru/company/flant/blog/531120/).
> 
> [Официальный сайт Containerd](https://containerd.io/).

### Kube Proxy

**Kube Proxy** – это агент, который запускается на каждой отдельной ноде через 
**DaemonSets**.

Отвечает за:

- Локальную сеть кластера (Local cluster networking)
- - Каждая нода получает свой уникальный IP-адрес
- Перенаправляет (routing) сетевой траффик на сервисы-балансировщики нагрузки 
  (load balanced services)


## Запуск Kubernetes

Способы запуска Kubernetes:

- Запустить самостоятельно. *Очень сложно*
- Использовать *управляемые **Kubernetes*** (*Managed **Kubernetes***)
- - **EKS** – **E**lastic **K**ubernetes **S**ervice
- - **GKE** – **G**oogle **K**ubernetes **S**ervice
- - **AKS** – **A**zure **K**ubernetes
- - И другие облачные поставщики (провайдеры, providers)

### Что значит *управляемые* **Kubernetes**

*Управляемые **Kubernetes*** избавляют нас от лишней работы с Master-нодой и 
позволяют сфокусироваться на разработке непосредственно нашего приложения.

> **Control Plane** полностью настроена за нас и для нас.

#### Пример на основе **Amazon EKS**

![](images/amazon-eks-example.png)

### Запуск кластера локально

Для запуска кластера локально существуют три основных способа:

- **Minikube**
- **Kind**
- **Docker**

Они служат следующим целям:

- Изучение **Kubernetes**
- Локальная разработка или **CI**

**ВАЖНО!** **НЕ НУЖНО ИСПОЛЬЗОВАТЬ ИХ НА РАЗЛИЧНЫХ СРЕДАХ ОКРУЖЕНИЯ, ВКЛЮЧАЯ 
PRODUCTION!**

> Здесь будет использоваться **Minikube**.

## Minikube

[Документация на Minikube](https://minikube.sigs.k8s.io/docs/).

### Запуск кластера Minikube

```bash
minikube start
```

Теперь мы имеем локальный кластер из одной ноды – master-ноды. Внутри находится 
**Control Plane** со всеми компонентами, которые описаны выше. У этой ноды 
есть [*свой IP-адрес*](#ip-адрес-master-ноды-minikube).

### Остановка кластера Minikube

```bash
minikube stop
```

> Все настройки сохранятся.

### Удаление кластера Minikube

```bash
minikube delete
```

### Статус Minikube

```bash
minikube status
```

Примерный результат:

```
minikube
type: Control Plane
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured
```

> Результат `docker ps`:
> 
> ```
> CONTAINER ID   IMAGE                                 COMMAND                  CREATED          STATUS          PORTS                                                                                                                                  NAMES
> a0fec3e7dbae   gcr.io/k8s-minikube/kicbase:v0.0.30   "/usr/local/bin/entr…"   22 minutes ago   Up 22 minutes   127.0.0.1:60181->22/tcp, 127.0.0.1:60182->2376/tcp, 127.0.0.1:60184->5000/tcp, 127.0.0.1:60185->8443/tcp, 127.0.0.1:60183->32443/tcp   minikube
> ```

[Драйверы Minikube](https://minikube.sigs.k8s.io/docs/drivers/).

### IP-адрес master-ноды Minikube

```bash
minikube ip
```

### Запуск Minikube с несколькими нодами

Для начала [удалим кластер](#удаление-кластера-minikube).

Напишем `minikube start --help`:

```diff
...
+ -n, --nodes=1: The number of nodes to spin up. Defaults to 1.
...
```

Таким образом запустим кластер с двумя нодами:

```bash
minikube start --nodes=2
```

Результат `minikube status`:

```diff
minikube
+ type: Control Plane
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured

minikube-m02
+ type: Worker
host: Running
kubelet: Running
```

Результат [`minikube get nodes`](#исследование-кластера):

```
NAME           STATUS   ROLES                  AGE    VERSION
minikube       Ready    control-plane,master   5m9s   v1.23.3
minikube-m02   Ready    <none>                 3m1s   v1.23.3
```

Теперь мы имеем кластер с двумя нодами и у каждой ноды 
[свой IP-адрес](#ip-адрес-master-ноды-minikube).

### Просмотр логов Minikube

Команда для получения логов master-ноды:

```bash
minikube logs
```

(получает логи **с самого запуска**, их очень много)

#### Просмотр логов с текущего момента

```bash
minikube logs -f
```

#### Просмотр логов определённой ноды

[Получаем список нод](#исследование-кластера).

Далее к `minikube logs` добавляем параметр `--node`, например:

```bash
minikube logs --node='minikube m02' -f
```

## Kubectl

**Kubectl** – инструмент командной строки **Kubernetes**. Представляет собой 
клиент для взаимодействия с **API Server**.

Позволяет взаимодействовать с кластером (посылать ему команды):

- Деплой
- Инспектирование
- Редактирование ресурсов
- Дебаг
- Просмотр логов
- И другое

**Под** (**Pod**) – набор из одного или нескольких контейнеров.

### Создание и запуск **пода** напрямую (императивный путь)

Создадим и запустим **под** с названием `hello-world` на основе Docker-образа 
`amigoscode/kubernetes:hello-world`:

```bash
kubectl run hello-world --image=amigoscode/kubernetes:hello-world --port=80
```

Результат выполнения команды:

```
pod/hello-world created
```

### Запуск **пода** с помощью конфигурационного файла (декларативный путь)

```bash
kubectl apply -f <путь_к_конфигурационному_файлу>
```

> Хитрый запуск через `cat`:
> 
> ```bash
> cat <путь_к_конфигурационному_файлу> | kubectl apply -f -
> ```

### Получение списка **подов**

> `// TODO:`
> - [ ] Универсально описать способы получения сущностей.

Команда получения *текущих **подов*** в *текущем пространстве имён* 
(*неймспейсе*):

```bash
kubectl get pods
```

Примерный результат:

```
NAME          READY   STATUS    RESTARTS   AGE
hello-world   1/1     Running   0          79s
```

#### Слежение за **подами**

Параметр `-w` позволяет получать новые состояния **подов** в реальном времени:

```bash
kubernetes get pods -w
```

#### Получение списка всех **подов**

```bash
kubectl get pods -A
```

#### Получения списка **подов** определённого пространства имён (неймспейса)

```bash
kubectl get pod -n <название_неймспейса>
```

### Получение списка ***всего***

Команда для получения списка *всех* сущностей:

```bash
kubectl get all
```

Команда для получения списка *всех* сущностей *отовсюду*:

```bash
kubectl get all -A
```

Примерный результат:

```
NAMESPACE     NAME                                   READY   STATUS    RESTARTS       AGE
kube-system   pod/coredns-64897985d-f492w            1/1     Running   3 (26h ago)    34h
kube-system   pod/etcd-minikube                      1/1     Running   3 (26h ago)    34h
kube-system   pod/kindnet-dplwt                      1/1     Running   3 (26h ago)    34h
kube-system   pod/kindnet-hzskd                      1/1     Running   3 (26h ago)    34h
kube-system   pod/kube-apiserver-minikube            1/1     Running   3 (26h ago)    34h
kube-system   pod/kube-controller-manager-minikube   1/1     Running   3 (26h ago)    34h
kube-system   pod/kube-proxy-bcpqc                   1/1     Running   3 (26h ago)    34h
kube-system   pod/kube-proxy-ppgxg                   1/1     Running   3 (26h ago)    34h
kube-system   pod/kube-scheduler-minikube            1/1     Running   3 (26h ago)    34h
kube-system   pod/storage-provisioner                1/1     Running   10 (61s ago)   34h

NAMESPACE     NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
default       service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP                  34h
kube-system   service/kube-dns     ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   34h

NAMESPACE     NAME                        DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR
AGE
kube-system   daemonset.apps/kindnet      2         2         2       2            2           <none>
34h
kube-system   daemonset.apps/kube-proxy   2         2         2       2            2           kubernetes.io/os=linux   34h

NAMESPACE     NAME                      READY   UP-TO-DATE   AVAILABLE   AGE
kube-system   deployment.apps/coredns   1/1     1            1           34h

NAMESPACE     NAME                                DESIRED   CURRENT   READY   AGE
kube-system   replicaset.apps/coredns-64897985d   1         1         1       34h
```

> Мы получили **поды** (`pod`), **сервисы** (`service`) и другое.

### Получение расширенной информации о **поде**

#### Выбор формата вывода

```bash
kubectl get pod <название_пода> -o wide
```

для вывода в расширенном формате.

Примерный результат:

```
NAME          READY   STATUS    RESTARTS   AGE   IP           NODE           NOMINATED NODE   READINESS GATES
hello-world   1/1     Running   0          34s   10.244.1.2   minikube-m02   <none>           <none>
```

Или

```bash
kubectl get pod <название_пода> -o yaml
```

для вывода в формате **YAML**.

Так же можно указать **JSON**-формат.

#### Получение полной информации

```bash
kubectl describe pod <название_пода>
```

### Получения списка пространств имён (неймспейсов)

```bash
kubectl get namespaces
```

Или

```bash
kubectl get ns
```

### Получение списка **Deployments**

```bash
kubectl get deployments
```

#### Получение полной информации

```bash
kubectl describe deployment <название_deployment>
```

### Получение списка **ReplicaSet**

```bash
kubectl get replicaset
```

Или

```bash
kubectl get rs
```

#### Получение полной информации

```bash
kubectl describe rs <название_replicaset>
```

### Проброс порта на хост

На примере **пода** `hello-world`:

```bash
kubectl port-forward pod/hello-world 8080:80
```

Порт `8080` – порта хоста, `80` – порт контейнера.

### Удаление **пода**

```bash
kubectl delete pod <название_пода>
```

### Дебаг **подов**

#### Просмотр логов

```bash
kubectl logs <название_пода>
```

Также можно смотреть логи в реальном времени:

```bash
kubectl logs <название_пода> -f
```

#### Просмотр логов для определённого контейнера

> **Под** – это *набор* контейнеров.

```bash
kubectl logs <название_пода> -c <название_контейнера>
```

#### `exec` в контейнерах

Зайдём в `bash`-оболочку контейнера какого-то пода:

```
kubectl exec -it <название_пода> -c <название_контейнера>  -- bash
```

Иногда может потребоваться заменить `bash` на `sh`.

### Просмотр ресурсов API

```bash
kubectl api-resource
```

Это выведет то, что поддерживает **Kubectl**.

Для получения справки по использованию **Kubectl** можно воспользоваться 
следующей командой:

```bash
kubectl --help
```

> Флаг `--help` можно применять к любой команде **Kubectl**.

### Исследование кластера

Получим список нод:

```bash
kubectl get nodes
```

Результат:

```
NAME       STATUS   ROLES                  AGE   VERSION
minikube   Ready    control-plane,master   32h   v1.23.3
```

Следует обратить внимание на колонку `ROLES`. Она говорит нам, что нода с 
названием `minikube` – это master-нода. **Minikube** создаёт её при выполнении 
команды `minikube start`.

Но если мы попробуем получить список **подов**:

```bash
kubectl get pods
```

То увидим, что результат будет пустой (нет ресурсов в *пространстве имён* 
(*неймспейсе*) по умолчанию):

```
No resources found in default namespace.
```

Но если добавить параметр `-A` к `kubectl get pods` – `kubectl get pods -A`, то 
мы получим **все** **поды**:

```
NAMESPACE     NAME                               READY   STATUS    RESTARTS      AGE
kube-system   coredns-64897985d-mbg9z            1/1     Running   1 (38m ago)   32h
kube-system   etcd-minikube                      1/1     Running   1 (38m ago)   32h
kube-system   kube-apiserver-minikube            1/1     Running   1 (38m ago)   32h
kube-system   kube-controller-manager-minikube   1/1     Running   1 (38m ago)   32h
kube-system   kube-proxy-47whj                   1/1     Running   1 (38m ago)   32h
kube-system   kube-scheduler-minikube            1/1     Running   1 (38m ago)   32h
kube-system   storage-provisioner                1/1     Running   3 (37m ago)   32h
```

> Все эти **поды** в одном неймспейсе – `kube-system`. **Под** `hello-world`, 
> который мы запускали ранее, создавался в неймспейсе `default`.

> Здесь нет **Cloud Controller Manager**, т.к. кластер запущен локально и не 
> использует облака.

#### Подключение к ноде по SSH

Для подключения по SSH следует выполнить следующую команду:

```bash
minikube ssh
```

> Можно указать ноду для подключения с помощью параметра `-n`. По умолчанию 
> подключение происходит к **Control Plane**.

> Если внутри написать `docker ps`, можно увидеть все сервисы, необходимые для 
> работы кластера:
> 
> ```
> CONTAINER ID   IMAGE                  COMMAND                  CREATED          STATUS          PORTS     NAMES
> 7ad334ff415b   6e38f40d628d           "/storage-provisioner"   48 minutes ago   Up 48 minutes             k8s_storage-provisioner_storage-provisioner_kube-system_ad5f6345-857d-41d9-bb93-eab64406337a_3
> a2830f984723   a4ca41631cc7           "/coredns -conf /etc…"   48 minutes ago   Up 48 minutes             k8s_coredns_coredns-64897985d-mbg9z_kube-system_91f163ef-3177-4863-a3fa-67a071ce423a_1
> a85f78ffe8ca   9b7cc9982109           "/usr/local/bin/kube…"   48 minutes ago   Up 48 minutes             k8s_kube-proxy_kube-proxy-47whj_kube-system_ba9b4016-0926-4ab2-b92d-40d3f284f9a3_1
> 753eca6a0986   k8s.gcr.io/pause:3.6   "/pause"                 48 minutes ago   Up 48 minutes             k8s_POD_kube-proxy-47whj_kube-system_ba9b4016-0926-4ab2-b92d-40d3f284f9a3_1
> 651563cc2935   k8s.gcr.io/pause:3.6   "/pause"                 48 minutes ago   Up 48 minutes             k8s_POD_coredns-64897985d-mbg9z_kube-system_91f163ef-3177-4863-a3fa-67a071ce423a_1
> d288e11149c0   k8s.gcr.io/pause:3.6   "/pause"                 48 minutes ago   Up 48 minutes             k8s_POD_storage-provisioner_kube-system_ad5f6345-857d-41d9-bb93-eab64406337a_1
> 0d126d5ef473   f40be0088a83           "kube-apiserver --ad…"   48 minutes ago   Up 48 minutes             k8s_kube-apiserver_kube-apiserver-minikube_kube-system_cd6e47233d36a9715b0ab9632f871843_1
> b170e1734e0a   25f8c7f3da61           "etcd --advertise-cl…"   48 minutes ago   Up 48 minutes             k8s_etcd_etcd-minikube_kube-system_9d3d310935e5fabe942511eec3e2cd0c_1
> fe360fddd78c   b07520cd7ab7           "kube-controller-man…"   48 minutes ago   Up 48 minutes             k8s_kube-controller-manager_kube-controller-manager-minikube_kube-system_b965983ec05322d0973594a01d5e8245_1
> f82d12d94d6d   99a3486be4f2           "kube-scheduler --au…"   48 minutes ago   Up 48 minutes             k8s_kube-scheduler_kube-scheduler-minikube_kube-system_be132fe5c6572cb34d93f5e05ce2a540_1
> 1905754b6018   k8s.gcr.io/pause:3.6   "/pause"                 48 minutes ago   Up 48 minutes             k8s_POD_kube-scheduler-minikube_kube-system_be132fe5c6572cb34d93f5e05ce2a540_1
> 0c8b9902d454   k8s.gcr.io/pause:3.6   "/pause"                 48 minutes ago   Up 48 minutes             k8s_POD_etcd-minikube_kube-system_9d3d310935e5fabe942511eec3e2cd0c_1
> 3949b6cdebe8   k8s.gcr.io/pause:3.6   "/pause"                 48 minutes ago   Up 48 minutes             k8s_POD_kube-controller-manager-minikube_kube-system_b965983ec05322d0973594a01d5e8245_1
> b428e53a7268   k8s.gcr.io/pause:3.6   "/pause"                 48 minutes ago   Up 48 minutes             k8s_POD_kube-apiserver-minikube_kube-system_cd6e47233d36a9715b0ab9632f871843_1
> ```

## **Поды** (**Pods**)

**Под** – наименьший (контейнеры такими не являются) возможный юнит для деплоя 
(deployable) в **Kubernetes**.

> Наименьший возможный юнит для деплоя (Smallest deployable unit):
> 
> В **Docker** – **контейнер**.
> 
> В **Kubernetes** – **под**.

![](images/pod.png)

В итоге **под** – это:

- Группа из одного или более контейнеров.
- Представляет запущенный процесс.
- Разделяет одну сеть и одни Volumes.
- Не следует создавать **поды** как самостоятельные единицы. Для создания 
  *следует использовать контроллеры*.
- Недолговечный (ephemeral, эфемерный) и "одноразовый" (disposable). *Именно 
  поэтому предыдущий пункт имеет силу*.

Следует обратить внимание:

- Никогда не нужно деплоить **поды**, используя `kind: Pod` в 
  **YAML**-конфигурации.
- **Поды** не умеют самоисцеляться (self-heal). Если **под** упал – он сам не 
  встанет. **Под** может упасть из-за непредвиденной ошибки в вашем приложении.

### Способы создания **подов**

Способы создания **подов**:

- Императивная (<s>повелительная</s>) команда:
  
  > Например:
  > 
  > ```bash
  > kubectl run hello-world --image=amigoscode/kubernetes:hello-world --port=80
  > ```

  Используется для изучения, устранения проблем, экспериментирования.
- Декларативная конфигурация
  
  Использование конфигурационного файла.

  Воспроизводимость (можно взять одну конфигурацию и применить её для различных
  окружений). Best practices.

#### Императивный способ

[Создание и запуск пода с помощью Kubectl](#создание-и-запуск-пода).

#### Декларативный способ

Создадим каталог `pods` и внутри него в файле `pod.yml` (такое название 
необязательно) запишем следующее:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: hello-world
  labels:
    name: hello-world
spec:
  containers:
  - name: hello-world
    image: amigoscode/kubernetes:hello-world
    resources:
      limits:
        memory: "128Mi"
        cpu: "500m"
    ports:
      - containerPort: 80
```

Перейдём в каталог `pods` с помощью команды `cd pods/` и выполним следующую 
команду для запуска **пода**:

```bash
kubectl apply -f pod.yml
```

Результат `kubectl get pods`:

```
NAME          READY   STATUS    RESTARTS   AGE
hello-world   1/1     Running   0          26s
```

> Для того, чтобы **под** работал как задумано, нужно 
> [пробросить порты](#проброс-порта-на-хост):
> 
> ```bash
> kubectl port-forward pod/hello-world 8080:80
> ```

### Deployments

**Deployments** – это ресурс **Kubernetes**, который управляет релизами нового 
приложения. Предоставляет возможность деплоя с нулевым временем простоя 
(*zero downtime*). Создаёт **ReplicaSet**.

![](images/deployments.png)

Создадим файл `deployment.yml` (такое название необязательно) со следующим 
содержимым:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-world
spec:
  selector:
    matchLabels:
      app: hello-world
  template:
    metadata:
      labels:
        app: hello-world
    spec:
      containers:
      - name: hello-world
        image: amigoscode/kubernetes:hello-world
        resources:
          limits:
            memory: "128Mi"
            cpu: "500m"
        ports:
        - containerPort: 80
```

Применим эту конфигурацию:

```bash
kubectl apply -f deployment.yml
```

### ReplicaSet

**ReplicaSet** гарантирует, что необходимое число **подов** всегда запущено.

> Так же как и с **подами**, никогда не следует создавать **ReplicaSet** как 
> самостоятельные единицы. Вместо этого следует использовать **Deployments**.

Создадим файл `deployment-replicas.yml` (такое название необязательно) со 
следующим содержимым:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-world
spec:
  replicas: 3
  selector:
    matchLabels:
      app: hello-world
  template:
    metadata:
      labels:
        app: hello-world
    spec:
      containers:
      - name: hello-world
        image: amigoscode/kubernetes:hello-world
        resources:
          limits:
            memory: "128Mi"
            cpu: "500m"
        ports:
        - containerPort: 80
```

> Следует обратить внимание на `replicas: 3`. Это поле указывает, что 
> **ReplicaSet** должен обеспечить запуск и работу 3-х **подов**.

Применим эту конфигурацию:

```bash
kubectl apply -f deployment-replicas.yml
```

Результат `kubectl get pods`:

```
NAME                           READY   STATUS    RESTARTS   AGE
hello-world-576c9b7fcd-4gs56   1/1     Running   0          105s
hello-world-576c9b7fcd-9gpc8   1/1     Running   0          105s
hello-world-576c9b7fcd-hqbwc   1/1     Running   0          105s
```

Результат `kubectl get rs`:

```
NAME                     DESIRED   CURRENT   READY   AGE
hello-world-576c9b7fcd   3         3         3       109s
```

Результат `kubectl get deployment`:

```
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
hello-world   3/3     3            3           12m
```

#### Контуры управления (Control loops)

**ReplicaSet** реализует фоновый контур управления (control loop), который 
проверяет, что необходимое число **подов** всегда представлено в кластере.

#### Rolling Update

Осуществим *Rolling Update*. Для этого создадим вторую версию нашего 
**Deployment**.

Чтобы получить новую версию достаточно поменять поле `image` в конфигурационном 
файле **Deployment**. Например (файл `deployment-v2.yml`):

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-world
spec:
  replicas: 3
  selector:
    matchLabels:
      app: hello-world
  template:
    metadata:
      labels:
        app: hello-world
      annotations:
        kubernetes.io/change-cause: "amigoscode/kubernetes:hello-world-v2"
    spec:
      containers:
      - name: hello-world
        image: amigoscode/kubernetes:hello-world-v2
        resources:
          limits:
            memory: "128Mi"
            cpu: "500m"
        ports:
        - containerPort: 80
```

> Здесь также добавлено поле `kubernetes.io/change-cause` в `annotations`, 
> которое описывает причину изменения. Оно может содержать *любой* текст.

Применим эту конфигурацию, осуществив *Rolling Update*:

```bash
kubectl apply -f deployment-v2.yml
```

С помощью `kubectl get pods -w` можем просмотреть процесс *Rolling Update*:

```
NAME                           READY   STATUS              RESTARTS   AGE
hello-world-576c9b7fcd-cg59z   1/1     Running             0          26s
hello-world-576c9b7fcd-jmv8q   1/1     Running             0          26s
hello-world-576c9b7fcd-vbjpd   1/1     Running             0          45s
hello-world-6dd6685f99-wcz8g   0/1     ContainerCreating   0          2s
hello-world-6dd6685f99-wcz8g   1/1     Running             0          3s
hello-world-576c9b7fcd-cg59z   1/1     Terminating         0          28s
hello-world-6dd6685f99-t745d   0/1     Pending             0          0s
hello-world-6dd6685f99-t745d   0/1     Pending             0          0s
hello-world-6dd6685f99-t745d   0/1     ContainerCreating   0          0s
hello-world-576c9b7fcd-cg59z   0/1     Terminating         0          29s
hello-world-576c9b7fcd-cg59z   0/1     Terminating         0          29s
hello-world-576c9b7fcd-cg59z   0/1     Terminating         0          29s
hello-world-6dd6685f99-t745d   1/1     Running             0          1s
hello-world-576c9b7fcd-jmv8q   1/1     Terminating         0          29s
hello-world-6dd6685f99-z9rdw   0/1     Pending             0          0s
hello-world-6dd6685f99-z9rdw   0/1     Pending             0          0s
hello-world-6dd6685f99-z9rdw   0/1     ContainerCreating   0          0s
hello-world-576c9b7fcd-jmv8q   0/1     Terminating         0          30s
hello-world-576c9b7fcd-jmv8q   0/1     Terminating         0          30s
hello-world-576c9b7fcd-jmv8q   0/1     Terminating         0          30s
hello-world-6dd6685f99-z9rdw   1/1     Running             0          2s
hello-world-576c9b7fcd-vbjpd   1/1     Terminating         0          50s
hello-world-576c9b7fcd-vbjpd   0/1     Terminating         0          51s
hello-world-576c9b7fcd-vbjpd   0/1     Terminating         0          51s
hello-world-576c9b7fcd-vbjpd   0/1     Terminating         0          51s
```

И уже результат:

```
NAME                           READY   STATUS    RESTARTS   AGE
hello-world-6dd6685f99-t745d   1/1     Running   0          2m22s
hello-world-6dd6685f99-wcz8g   1/1     Running   0          2m26s
hello-world-6dd6685f99-z9rdw   1/1     Running   0          2m21s
```

Можно убедиться, что версия обновилась с помощью перенаправления портов. 
**ТАКОЙ СПОСОБ МОЖНО ИСПОЛЬЗОВАТЬ ТОЛЬКО НА ДЕБАГЕ**:

```bash
kubectl port-forward deployment/hello-world 8080:80
```

#### Предыдущий **ReplicaSet** после *Rolling Update*

*Rolling Update* оставляет предыдущий **ReplicaSet**.

Результат `kubectl get rs`:

```
NAME                     DESIRED   CURRENT   READY   AGE
hello-world-576c9b7fcd   0         0         0       20m
hello-world-6dd6685f99   3         3         3       19m
```

Выполним следующую команду для просмотра версий **Deployment**'а под названием 
`hello-world`:

```bash
kubectl rollout history deployment hello-world
```

Результат:

```
deployment.apps/hello-world
REVISION  CHANGE-CAUSE
1         <none>
2         amigoscode/kubernetes:hello-world-v2
```

> `CHANGE-CAUSE` берётся из `annotations` **YAML**-конфигурации.

#### Откат версии

```bash
kubectl rollout undo deployment hello-world --to-revision=<номер_ревизии>
```

> Если не указать номер ревизии, то откат произойдёт к предыдущей версии.

> По умолчанию **Kubernetes** хранит *только 10 версий*. Но в 
> **YAML**-конфигурации в `spec`->`revisionHistoryLimit` можно указать другое 
> значение.

#### Информация о ревизии (версии)

```bash
kubectl rollout history deployment <название_deployment> --revision=<номер_ревизии>
```
