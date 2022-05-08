# Kubernetes

[Wikipedia](https://ru.wikipedia.org/wiki/Kubernetes):

**Kubernetes** (от др. -греч. *κυβερνήτης* — «*кормчий*», «*рулевой*», часто 
также используется нумероним *k8s*) — открытое программное обеспечение для 
[*оркестровки*](https://ru.wikipedia.org/wiki/%D0%9E%D1%80%D0%BA%D0%B5%D1%81%D1%82%D1%80%D0%BE%D0%B2%D0%BA%D0%B0_(%D0%98%D0%A2)) 
контейнеризированных приложений — автоматизации их развёртывания, 
масштабирования и координации в условиях кластера.

Задачи **Kubernetes**:

- Деплой и управление приложениями (**контейнерами**)
- Масштабирование и уменьшение согласно текущим требованиям
- Деплой с нулевым временем простоя (*zero downtime*)
- Откаты
- И многое другое

## Источники

- [Mastering Kubernetes. Master k8s from A to Z](https://amigoscode.com/p/kubernetes)


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

**Controller Manager** – это демон, который управляет контуром управления. Это 
контроллер контроллеров.

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

### Запуск Minikube

```bash
minikube start
```

Теперь мы имеем локальный кластер из одной ноды – master-ноды. Внутри находится 
**Control Plane** со всеми компонентами, которые описаны выше. У этой ноды 
есть [*свой IP-адрес*](#ip-адрес-master-ноды-minikube).

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

## Kubectl

**Kubectl** – инструмент командной строки **Kubernetes**.

Позволяет взаимодействовать с кластером (посылать ему команды):

- Деплой
- Инспектирование
- Редактирование ресурсов
- Дебаг
- Просмотр логов
- И другое

**Под** (**Pod**) – набор из одного или нескольких контейнеров.

### Создание и запуск пода

Создадим и запустим **под** с названием `hello-world` на основе Docker-образа 
`amigoscode/kubernetes:hello-world`:

```bash
kubectl run hello-world --image=amigoscode/kubernetes:hello-world --port=80
```

Результат выполнения команды:

```
pod/hello-world created
```

### Получение списка **подов**

```bash
kubectl get pods
```

Примерный результат:

```
NAME          READY   STATUS    RESTARTS   AGE
hello-world   1/1     Running   0          79s
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
