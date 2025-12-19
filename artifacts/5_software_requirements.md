# Требования к программному обеспечению

На Автоматизированную информационную систему учёта и управления IT-сервисами (ITSIMS)

| Дата       | Номер версии | Автор     | Ссылка на задачу                                                          | Комментарий             |
| ---------- | ------------ | --------- | ------------------------------------------------------------------------- | ----------------------- |
| 2025-12-13 | v1.0         | Вальковец | [2025-12-13_Valkovets.md](../task-tracker/3_done/2025-12-13_Valkovets.md) | Первая версия документа |



## Оглавление

- [Требования к программному обеспечению](#требования-к-программному-обеспечению)
  - [Оглавление](#оглавление)
  - [1. Цель документа](#1-цель-документа)
  - [2. Термины, определения и сокращения](#2-термины-определения-и-сокращения)
    - [2.1. Базовые термины](#21-базовые-термины)
    - [2.2. Технические термины](#22-технические-термины)
    - [2.3. Сокращения и аббревиатуры](#23-сокращения-и-аббревиатуры)
  - [3. Архитектура программного обеспечения](#3-архитектура-программного-обеспечения)
    - [3.1. Компоненты системы](#31-компоненты-системы)
    - [3.2. Взаимодействие компонентов](#32-взаимодействие-компонентов)
    - [3.3. Технологический стек](#33-технологический-стек)
  - [4. Требования к Backend](#4-требования-к-backend)
    - [4.1. API модуля "Аутентификация" (auth)](#41-api-модуля-аутентификация-auth)
      - [Назначение](#назначение)
      - [Endpoints](#endpoints)
        - [POST /api/v1/auth/login — Вход в систему](#post-apiv1authlogin--вход-в-систему)
        - [POST /api/v1/auth/refresh — Обновление токена](#post-apiv1authrefresh--обновление-токена)
        - [POST /api/v1/auth/logout — Выход из системы](#post-apiv1authlogout--выход-из-системы)
        - [POST /api/v1/auth/password-reset — Запрос сброса пароля](#post-apiv1authpassword-reset--запрос-сброса-пароля)
        - [POST /api/v1/auth/password-reset/confirm — Подтверждение сброса пароля](#post-apiv1authpassword-resetconfirm--подтверждение-сброса-пароля)
      - [Безопасность](#безопасность)
    - [4.2. API модуля "Пользователи" (users)](#42-api-модуля-пользователи-users)
      - [Назначение](#назначение-1)
      - [Endpoints](#endpoints-1)
        - [GET /api/v1/users — Список пользователей](#get-apiv1users--список-пользователей)
        - [POST /api/v1/users — Создание пользователя](#post-apiv1users--создание-пользователя)
        - [GET /api/v1/users/{id} — Получение пользователя](#get-apiv1usersid--получение-пользователя)
        - [PUT /api/v1/users/{id} — Обновление пользователя](#put-apiv1usersid--обновление-пользователя)
        - [POST /api/v1/users/{id}/block — Блокировка пользователя](#post-apiv1usersidblock--блокировка-пользователя)
        - [POST /api/v1/users/{id}/unblock — Разблокировка пользователя](#post-apiv1usersidunblock--разблокировка-пользователя)
      - [Валидация](#валидация)
    - [4.3. API модуля "IT-сервисы" (services)](#43-api-модуля-it-сервисы-services)
      - [Назначение](#назначение-2)
      - [Endpoints](#endpoints-2)
        - [GET /api/v1/services — Список сервисов](#get-apiv1services--список-сервисов)
        - [POST /api/v1/services — Создание сервиса](#post-apiv1services--создание-сервиса)
        - [GET /api/v1/services/{id} — Получение сервиса](#get-apiv1servicesid--получение-сервиса)
        - [PUT /api/v1/services/{id} — Обновление сервиса](#put-apiv1servicesid--обновление-сервиса)
        - [DELETE /api/v1/services/{id} — Удаление сервиса](#delete-apiv1servicesid--удаление-сервиса)
        - [PATCH /api/v1/services/{id}/status — Изменение статуса](#patch-apiv1servicesidstatus--изменение-статуса)
      - [Валидация](#валидация-1)
      - [Права доступа](#права-доступа)
        - [GET /api/v1/services/{id}/permissions — Получение прав доступа](#get-apiv1servicesidpermissions--получение-прав-доступа)
        - [POST /api/v1/services/{id}/permissions — Назначение прав](#post-apiv1servicesidpermissions--назначение-прав)
        - [DELETE /api/v1/services/{id}/permissions/{user\_id} — Отзыв прав](#delete-apiv1servicesidpermissionsuser_id--отзыв-прав)
    - [4.4. API модуля "Категории" (categories)](#44-api-модуля-категории-categories)
      - [Назначение](#назначение-3)
      - [Endpoints](#endpoints-3)
        - [GET /api/v1/categories — Список категорий](#get-apiv1categories--список-категорий)
        - [POST /api/v1/categories — Создание категории](#post-apiv1categories--создание-категории)
        - [PUT /api/v1/categories/{id} — Обновление категории](#put-apiv1categoriesid--обновление-категории)
        - [DELETE /api/v1/categories/{id} — Удаление категории](#delete-apiv1categoriesid--удаление-категории)
      - [Шаблоны категорий](#шаблоны-категорий)
        - [GET /api/v1/categories/{id}/templates — Список шаблонов](#get-apiv1categoriesidtemplates--список-шаблонов)
        - [POST /api/v1/categories/{id}/templates — Создание шаблона](#post-apiv1categoriesidtemplates--создание-шаблона)
        - [PUT /api/v1/categories/{id}/templates/{template\_id} — Обновление шаблона](#put-apiv1categoriesidtemplatestemplate_id--обновление-шаблона)
      - [Валидация](#валидация-2)
    - [4.5. API модуля "Документация" (documents)](#45-api-модуля-документация-documents)
      - [Назначение](#назначение-4)
      - [Endpoints](#endpoints-4)
        - [GET /api/v1/services/{service\_id}/documents — Структура каталога](#get-apiv1servicesservice_iddocuments--структура-каталога)
        - [POST /api/v1/services/{service\_id}/documents/folders — Создание папки](#post-apiv1servicesservice_iddocumentsfolders--создание-папки)
        - [DELETE /api/v1/services/{service\_id}/documents/folders/{folder\_id} — Удаление папки](#delete-apiv1servicesservice_iddocumentsfoldersfolder_id--удаление-папки)
        - [POST /api/v1/services/{service\_id}/documents/files — Загрузка файла](#post-apiv1servicesservice_iddocumentsfiles--загрузка-файла)
        - [GET /api/v1/services/{service\_id}/documents/files/{file\_id} — Скачивание файла](#get-apiv1servicesservice_iddocumentsfilesfile_id--скачивание-файла)
        - [DELETE /api/v1/services/{service\_id}/documents/files/{file\_id} — Удаление файла](#delete-apiv1servicesservice_iddocumentsfilesfile_id--удаление-файла)
        - [PATCH /api/v1/services/{service\_id}/documents/files/{file\_id}/type — Изменение типа документа](#patch-apiv1servicesservice_iddocumentsfilesfile_idtype--изменение-типа-документа)
      - [Безопасность файлов](#безопасность-файлов)
    - [4.6. API модуля "Обслуживание" (maintenance)](#46-api-модуля-обслуживание-maintenance)
      - [Назначение](#назначение-5)
      - [Endpoints](#endpoints-5)
        - [GET /api/v1/services/{service\_id}/maintenance — История обслуживания](#get-apiv1servicesservice_idmaintenance--история-обслуживания)
        - [POST /api/v1/services/{service\_id}/maintenance — Создание записи](#post-apiv1servicesservice_idmaintenance--создание-записи)
      - [Валидация](#валидация-3)
    - [4.7. API модуля "Связи сервисов" (dependencies)](#47-api-модуля-связи-сервисов-dependencies)
      - [Назначение](#назначение-6)
      - [Endpoints](#endpoints-6)
        - [GET /api/v1/services/{service\_id}/dependencies — Связи сервиса](#get-apiv1servicesservice_iddependencies--связи-сервиса)
        - [POST /api/v1/services/{service\_id}/dependencies — Создание связи](#post-apiv1servicesservice_iddependencies--создание-связи)
        - [DELETE /api/v1/services/{service\_id}/dependencies/{dependency\_id} — Удаление связи](#delete-apiv1servicesservice_iddependenciesdependency_id--удаление-связи)
        - [GET /api/v1/dependencies/graph — Полный граф зависимостей](#get-apiv1dependenciesgraph--полный-граф-зависимостей)
    - [4.8. API модуля "Аудит" (audit)](#48-api-модуля-аудит-audit)
      - [Назначение](#назначение-7)
      - [Endpoints](#endpoints-7)
        - [GET /api/v1/audit — Журнал действий](#get-apiv1audit--журнал-действий)
      - [Логируемые события](#логируемые-события)
      - [Безопасность журнала](#безопасность-журнала)
  - [5. Требования к Frontend](#5-требования-к-frontend)
    - [5.1. Структура компонентов](#51-структура-компонентов)
    - [5.2. Управление состоянием](#52-управление-состоянием)
    - [5.3. Требования к UI компонентам](#53-требования-к-ui-компонентам)
      - [Навигация и layout](#навигация-и-layout)
      - [Реестр сервисов](#реестр-сервисов)
      - [Детальная страница сервиса](#детальная-страница-сервиса)
      - [Файловый каталог](#файловый-каталог)
      - [Граф зависимостей](#граф-зависимостей)
      - [Обратная связь](#обратная-связь)
      - [Доступность](#доступность)
  - [6. Требования к данным](#6-требования-к-данным)
    - [6.1. Схема базы данных](#61-схема-базы-данных)
      - [Таблица users](#таблица-users)
      - [Таблица refresh\_tokens](#таблица-refresh_tokens)
      - [Таблица categories](#таблица-categories)
      - [Таблица category\_templates](#таблица-category_templates)
      - [Таблица services](#таблица-services)
      - [Таблица service\_permissions](#таблица-service_permissions)
      - [Таблица document\_folders](#таблица-document_folders)
      - [Таблица document\_files](#таблица-document_files)
      - [Таблица maintenance\_records](#таблица-maintenance_records)
      - [Таблица service\_dependencies](#таблица-service_dependencies)
      - [Таблица audit\_log](#таблица-audit_log)
    - [6.2. Миграции и хранение](#62-миграции-и-хранение)
  - [7. Требования к инфраструктуре](#7-требования-к-инфраструктуре)
    - [7.1. Развёртывание](#71-развёртывание)
    - [7.2. Конфигурация](#72-конфигурация)
    - [7.3. Логирование и мониторинг](#73-логирование-и-мониторинг)
  - [8. Требования к тестированию](#8-требования-к-тестированию)
  - [9. Связанные документы](#9-связанные-документы)
    - [Артефакты проекта](#артефакты-проекта)
    - [Трассировка требований](#трассировка-требований)
    - [Правила оформления](#правила-оформления)
    - [Валидация и проверка качества](#валидация-и-проверка-качества)



## 1. Цель документа

Данный документ определяет требования к программному обеспечению системы ITSIMS (IT Service Inventory and Management System).

Документ предназначен для разработчиков, тестировщиков, DevOps-инженеров и всех участников команды, работающих над реализацией системы.

Из документа читатель узнает о технической реализации системы: архитектуре, API, структуре базы данных, требованиях к инфраструктуре и тестированию.

**Принцип:** Требования к программному обеспечению отвечают на вопрос "**КАК** технически будет реализована система", описывая конкретные технологии, протоколы, структуры данных и архитектурные решения.



## 2. Термины, определения и сокращения

### 2.1. Базовые термины

- **Заказчик** — АО «БАРС Груп» в лице представителя Суставова Сергея Александровича

- **Система** — Автоматизированная информационная система учёта и управления IT-сервисами (ITSIMS), разрабатываемая в рамках данного проекта

- **Пользователь** — сотрудник организации-заказчика, использующий систему в рамках своих должностных обязанностей и предоставленных прав доступа

- **Администратор системы** — пользователь с полным контролем над системой, включая управление всеми сервисами, категориями, правами доступа и настройками системы

- **IT-сервис** — единица учёта в системе, представляющая собой программное обеспечение с определённым набором характеристик и ассоциированной документацией

### 2.2. Технические термины

- **Endpoint** — конечная точка API, представляющая собой URL-адрес для выполнения определённой операции

- **JWT (JSON Web Token)** — стандарт токена доступа, используемый для аутентификации и передачи данных между сторонами в формате JSON

- **Access Token** — краткосрочный токен для аутентификации запросов к API

- **Refresh Token** — долгосрочный токен для обновления access token без повторного ввода учётных данных

- **ORM (Object-Relational Mapping)** — технология программирования, связывающая базы данных с объектно-ориентированными языками программирования

- **Middleware** — промежуточное программное обеспечение, обрабатывающее запросы между клиентом и сервером

- **CORS (Cross-Origin Resource Sharing)** — механизм, позволяющий веб-страницам запрашивать ресурсы с другого домена

- **S3 (Simple Storage Service)** — протокол объектного хранилища, стандарт де-факто для хранения файлов

- **MIME-тип** — идентификатор формата файла (например, application/pdf, image/png)

- **Магические байты (File Signature)** — последовательность байтов в начале файла, идентифицирующая его формат

- **Хеш-функция** — криптографическая функция для преобразования данных в строку фиксированной длины

- **bcrypt** — адаптивная хеш-функция для безопасного хранения паролей

### 2.3. Сокращения и аббревиатуры

- **API (Application Programming Interface, англ.)** — программный интерфейс приложения

- **REST (Representational State Transfer, англ.)** — архитектурный стиль взаимодействия компонентов распределённого приложения

- **SWR (Software Requirements, англ.)** — требования к программному обеспечению

- **SR (System Requirements, англ.)** — системные требования

- **CRUD (Create, Read, Update, Delete, англ.)** — базовые операции с данными

- **UTC (Coordinated Universal Time, англ.)** — всемирное координированное время

- **JSON (JavaScript Object Notation, англ.)** — текстовый формат обмена данными

- **SQL (Structured Query Language, англ.)** — язык структурированных запросов

- **HTTP (HyperText Transfer Protocol, англ.)** — протокол передачи гипертекста

- **HTTPS (HTTP Secure, англ.)** — расширение HTTP с шифрованием



## 3. Архитектура программного обеспечения

### 3.1. Компоненты системы

Система ITSIMS состоит из следующих программных компонентов:

| Компонент          | Назначение                                            | Технологии                                     |
| ------------------ | ----------------------------------------------------- | ---------------------------------------------- |
| **Frontend (SPA)** | Клиентское веб-приложение, пользовательский интерфейс | React 18, TypeScript 5, Chakra UI 2, Zustand 5 |
| **Backend (API)**  | Серверная часть, бизнес-логика, REST API              | Python 3.12, FastAPI 0.115, SQLAlchemy 2.0     |
| **Database**       | Хранилище структурированных данных                    | PostgreSQL 16                                  |
| **File Storage**   | Объектное хранилище файлов документации               | MinIO (S3-совместимое)                         |
| **Mail Server**    | Отправка email (восстановление пароля)                | Mailhog (dev), SMTP (prod)                     |

### 3.2. Взаимодействие компонентов

Компоненты системы взаимодействуют по следующей схеме:

```
┌─────────────┐     HTTPS/JSON      ┌─────────────┐
│   Frontend  │ ◄─────────────────► │   Backend   │
│   (React)   │     REST API        │  (FastAPI)  │
└─────────────┘                     └──────┬──────┘
                                           │
                    ┌──────────────────────┼──────────────────────┐
                    │                      │                      │
                    ▼                      ▼                      ▼
            ┌──────────────┐      ┌──────────────┐      ┌──────────────┐
            │  PostgreSQL  │      │    MinIO     │      │   Mailhog    │
            │   Database   │      │  S3 Storage  │      │    SMTP      │
            └──────────────┘      └──────────────┘      └──────────────┘
```

**Протоколы и форматы:**

- **Frontend ↔ Backend:** REST API через HTTPS, формат данных JSON
- **Backend ↔ PostgreSQL:** SQL через драйвер asyncpg
- **Backend ↔ MinIO:** S3 API через библиотеку boto3
- **Backend ↔ SMTP:** протокол SMTP для отправки email

**Формат даты и времени:** ISO 8601 в UTC (например, `2025-12-13T14:30:00Z`)

### 3.3. Технологический стек

| Категория          | Технология            | Версия | Назначение            |
| ------------------ | --------------------- | ------ | --------------------- |
| **Backend**        | Python                | 3.12   | Язык программирования |
| Backend            | FastAPI               | 0.115  | Web-фреймворк         |
| Backend            | SQLAlchemy            | 2.0    | ORM                   |
| Backend            | Alembic               | 1.13   | Миграции БД           |
| Backend            | Pydantic              | 2.x    | Валидация данных      |
| Backend            | python-jose           | 3.x    | JWT токены            |
| Backend            | passlib[bcrypt]       | 1.7    | Хеширование паролей   |
| Backend            | boto3                 | 1.x    | S3 клиент             |
| Backend            | structlog             | 24.x   | Логирование           |
| **Frontend**       | React                 | 18.x   | UI фреймворк          |
| Frontend           | TypeScript            | 5.x    | Типизация             |
| Frontend           | Vite                  | 5.x    | Сборщик               |
| Frontend           | Zustand               | 5.x    | State management      |
| Frontend           | Chakra UI             | 2.x    | UI библиотека         |
| Frontend           | React Router          | 6.x    | Роутинг               |
| Frontend           | Axios                 | 1.x    | HTTP клиент           |
| Frontend           | React Flow            | 11.x   | Визуализация графа    |
| **Database**       | PostgreSQL            | 16     | СУБД                  |
| **Storage**        | MinIO                 | latest | S3-хранилище          |
| **Infrastructure** | Docker                | 24.x   | Контейнеризация       |
| Infrastructure     | Docker Compose        | 2.x    | Оркестрация           |
| **Testing**        | pytest                | 8.x    | Тесты Backend         |
| Testing            | pytest-cov            | 4.x    | Покрытие кода         |
| Testing            | Jest                  | 29.x   | Тесты Frontend        |
| Testing            | React Testing Library | 14.x   | Тесты компонентов     |



## 4. Требования к Backend

### 4.1. API модуля "Аутентификация" (auth)

#### Назначение

Модуль отвечает за аутентификацию пользователей, управление JWT-токенами и сессиями.

#### Endpoints

##### POST /api/v1/auth/login — Вход в систему

- <a id="swr-api001"></a>**SWR-API001:** Модуль должен предоставлять endpoint `POST /api/v1/auth/login` для аутентификации пользователя с JSON payload: `{ username: string (required), password: string (required) }` [↑](./4_SR_to_SWR.md#swr-api001-up)

- <a id="swr-api002"></a>**SWR-API002:** При успешной аутентификации API должен возвращать HTTP 200 с телом: `{ access_token: string, refresh_token: string, token_type: "bearer", expires_in: integer }` [↑](./4_SR_to_SWR.md#swr-api002-up)

- <a id="swr-api003"></a>**SWR-API003:** При неверных учётных данных API должен возвращать HTTP 401 с телом: `{ detail: "Неверный логин или пароль" }` [↑](./4_SR_to_SWR.md#swr-api003-up)

- <a id="swr-api004"></a>**SWR-API004:** При попытке входа заблокированного пользователя API должен возвращать HTTP 403 с телом: `{ detail: "Учётная запись заблокирована" }` [↑](./4_SR_to_SWR.md#swr-api004-up)

##### POST /api/v1/auth/refresh — Обновление токена

- <a id="swr-api005"></a>**SWR-API005:** Модуль должен предоставлять endpoint `POST /api/v1/auth/refresh` для обновления access token с JSON payload: `{ refresh_token: string (required) }` [↑](./4_SR_to_SWR.md#swr-api005-up)

- <a id="swr-api006"></a>**SWR-API006:** При валидном refresh token API должен возвращать HTTP 200 с новой парой токенов [↑](./4_SR_to_SWR.md#swr-api006-up)

- <a id="swr-api007"></a>**SWR-API007:** При невалидном или истёкшем refresh token API должен возвращать HTTP 401 [↑](./4_SR_to_SWR.md#swr-api007-up)

##### POST /api/v1/auth/logout — Выход из системы

- <a id="swr-api008"></a>**SWR-API008:** Модуль должен предоставлять endpoint `POST /api/v1/auth/logout` для завершения сессии [↑](./4_SR_to_SWR.md#swr-api008-up)

- <a id="swr-api009"></a>**SWR-API009:** При выходе модуль должен инвалидировать refresh token в базе данных [↑](./4_SR_to_SWR.md#swr-api009-up)

##### POST /api/v1/auth/password-reset — Запрос сброса пароля

- <a id="swr-api010"></a>**SWR-API010:** Модуль должен предоставлять endpoint `POST /api/v1/auth/password-reset` для запроса сброса пароля с JSON payload: `{ email: string (required) }` [↑](./4_SR_to_SWR.md#swr-api010-up)

- <a id="swr-api011"></a>**SWR-API011:** Модуль должен отправлять email со ссылкой для сброса пароля, содержащей одноразовый токен со сроком действия 1 час [↑](./4_SR_to_SWR.md#swr-api011-up)

##### POST /api/v1/auth/password-reset/confirm — Подтверждение сброса пароля

- <a id="swr-api012"></a>**SWR-API012:** Модуль должен предоставлять endpoint `POST /api/v1/auth/password-reset/confirm` с JSON payload: `{ token: string (required), new_password: string (required) }` [↑](./4_SR_to_SWR.md#swr-api012-up)

#### Безопасность

- <a id="swr-sec001"></a>**SWR-SEC001:** Access token должен иметь срок жизни 30 минут [↑](./4_SR_to_SWR.md#swr-sec001-up)

- <a id="swr-sec002"></a>**SWR-SEC002:** Refresh token должен иметь срок жизни 7 дней [↑](./4_SR_to_SWR.md#swr-sec002-up)

- <a id="swr-sec003"></a>**SWR-SEC003:** При неактивности пользователя более 12 часов refresh token должен быть инвалидирован [↑](./4_SR_to_SWR.md#swr-sec003-up)

- <a id="swr-sec004"></a>**SWR-SEC004:** JWT токены должны подписываться алгоритмом HS256 с использованием секретного ключа из переменной окружения JWT_SECRET [↑](./4_SR_to_SWR.md#swr-sec004-up)

- <a id="swr-sec005"></a>**SWR-SEC005:** Пароли пользователей должны храниться в виде хеша bcrypt с cost factor 12 [↑](./4_SR_to_SWR.md#swr-sec005-up)

- <a id="swr-sec006"></a>**SWR-SEC006:** Модуль должен поддерживать одновременную работу пользователя с нескольких устройств (множественные активные refresh tokens) [↑](./4_SR_to_SWR.md#swr-sec006-up)


### 4.2. API модуля "Пользователи" (users)

#### Назначение

Модуль отвечает за CRUD операции с пользователями и управление ролями.

#### Endpoints

##### GET /api/v1/users — Список пользователей

- <a id="swr-api013"></a>**SWR-API013:** Модуль должен предоставлять endpoint `GET /api/v1/users` для получения списка пользователей с поддержкой пагинации: `?page=1&per_page=20` [↑](./4_SR_to_SWR.md#swr-api013-up)

- <a id="swr-api014"></a>**SWR-API014:** Endpoint должен быть доступен только пользователям с ролью "Администратор системы" [↑](./4_SR_to_SWR.md#swr-api014-up)

##### POST /api/v1/users — Создание пользователя

- <a id="swr-api015"></a>**SWR-API015:** Модуль должен предоставлять endpoint `POST /api/v1/users` для создания пользователя с JSON payload: `{ username: string (required, unique), email: string (required, unique), password: string (required), first_name: string (required), last_name: string (required), position: string (required), department: string (required) }` [↑](./4_SR_to_SWR.md#swr-api015-up)

- <a id="swr-api016"></a>**SWR-API016:** При успешном создании API должен возвращать HTTP 201 с данными созданного пользователя (без пароля) [↑](./4_SR_to_SWR.md#swr-api016-up)

- <a id="swr-api017"></a>**SWR-API017:** Модуль должен автоматически назначать роль "Администратор системы" пользователю с должностью "Технический директор" [↑](./4_SR_to_SWR.md#swr-api017-up)

##### GET /api/v1/users/{id} — Получение пользователя

- <a id="swr-api018"></a>**SWR-API018:** Модуль должен предоставлять endpoint `GET /api/v1/users/{id}` для получения информации о пользователе [↑](./4_SR_to_SWR.md#swr-api018-up)

##### PUT /api/v1/users/{id} — Обновление пользователя

- <a id="swr-api019"></a>**SWR-API019:** Модуль должен предоставлять endpoint `PUT /api/v1/users/{id}` для обновления данных пользователя [↑](./4_SR_to_SWR.md#swr-api019-up)

##### POST /api/v1/users/{id}/block — Блокировка пользователя

- <a id="swr-api020"></a>**SWR-API020:** Модуль должен предоставлять endpoint `POST /api/v1/users/{id}/block` для блокировки пользователя [↑](./4_SR_to_SWR.md#swr-api020-up)

- <a id="swr-api021"></a>**SWR-API021:** При блокировке модуль должен инвалидировать все активные refresh tokens пользователя [↑](./4_SR_to_SWR.md#swr-api021-up)

##### POST /api/v1/users/{id}/unblock — Разблокировка пользователя

- <a id="swr-api022"></a>**SWR-API022:** Модуль должен предоставлять endpoint `POST /api/v1/users/{id}/unblock` для разблокировки пользователя [↑](./4_SR_to_SWR.md#swr-api022-up)

#### Валидация

- <a id="swr-f001"></a>**SWR-F001:** Модуль должен валидировать username: обязательное, уникальное, 3-50 символов, только латинские буквы, цифры и символ подчёркивания [↑](./4_SR_to_SWR.md#swr-f001-up)

- <a id="swr-f002"></a>**SWR-F002:** Модуль должен валидировать email: обязательное, уникальное, корректный формат email [↑](./4_SR_to_SWR.md#swr-f002-up)

- <a id="swr-f003"></a>**SWR-F003:** Модуль должен валидировать password: минимум 8 символов, минимум 1 заглавная буква, 1 строчная буква, 1 цифра [↑](./4_SR_to_SWR.md#swr-f003-up)


### 4.3. API модуля "IT-сервисы" (services)

#### Назначение

Модуль отвечает за CRUD операции с IT-сервисами.

#### Endpoints

##### GET /api/v1/services — Список сервисов

- <a id="swr-api023"></a>**SWR-API023:** Модуль должен предоставлять endpoint `GET /api/v1/services` для получения списка IT-сервисов, доступных текущему пользователю [↑](./4_SR_to_SWR.md#swr-api023-up)

- <a id="swr-api024"></a>**SWR-API024:** Endpoint должен поддерживать пагинацию: `?page=1&per_page=20` с возможными значениями per_page: 20, 50, 100 [↑](./4_SR_to_SWR.md#swr-api024-up)

- <a id="swr-api025"></a>**SWR-API025:** Endpoint должен поддерживать поиск по названию: `?search=crm` с частичным совпадением (ILIKE) [↑](./4_SR_to_SWR.md#swr-api025-up)

- <a id="swr-api026"></a>**SWR-API026:** Endpoint должен поддерживать фильтрацию по категории: `?category_id=1` [↑](./4_SR_to_SWR.md#swr-api026-up)

- <a id="swr-api027"></a>**SWR-API027:** Endpoint должен поддерживать фильтрацию по статусу: `?status=active` [↑](./4_SR_to_SWR.md#swr-api027-up)

- <a id="swr-api028"></a>**SWR-API028:** Endpoint должен поддерживать сортировку: `?sort_by=name&order=asc` с допустимыми полями: name, category, status, version, updated_at [↑](./4_SR_to_SWR.md#swr-api028-up)

- <a id="swr-api029"></a>**SWR-API029:** Ответ должен содержать метаданные пагинации: `{ items: [...], total: integer, page: integer, per_page: integer, pages: integer }` [↑](./4_SR_to_SWR.md#swr-api029-up)

##### POST /api/v1/services — Создание сервиса

- <a id="swr-api030"></a>**SWR-API030:** Модуль должен предоставлять endpoint `POST /api/v1/services` для создания IT-сервиса с JSON payload: `{ name: string (required), description: string (required), category_id: integer (required), status: enum (required), version: string (optional), hardware_ids: array[string] (optional) }` [↑](./4_SR_to_SWR.md#swr-api030-up)

- <a id="swr-api031"></a>**SWR-API031:** При успешном создании API должен возвращать HTTP 201 с данными созданного сервиса, включая автоматически присвоенный id [↑](./4_SR_to_SWR.md#swr-api031-up)

- <a id="swr-api032"></a>**SWR-API032:** Модуль должен автоматически создавать файловый каталог в MinIO для документации сервиса при создании [↑](./4_SR_to_SWR.md#swr-api032-up)

- <a id="swr-api033"></a>**SWR-API033:** Модуль должен применять шаблон структуры папок категории при создании каталога, если шаблон настроен [↑](./4_SR_to_SWR.md#swr-api033-up)

- <a id="swr-api034"></a>**SWR-API034:** При дублировании названия API должен возвращать HTTP 409 с телом: `{ detail: "Сервис с таким названием уже существует" }` [↑](./4_SR_to_SWR.md#swr-api034-up)

##### GET /api/v1/services/{id} — Получение сервиса

- <a id="swr-api035"></a>**SWR-API035:** Модуль должен предоставлять endpoint `GET /api/v1/services/{id}` для получения детальной информации о сервисе [↑](./4_SR_to_SWR.md#swr-api035-up)

- <a id="swr-api036"></a>**SWR-API036:** Ответ должен включать все атрибуты сервиса, включая дату последнего обслуживания для пользователей с соответствующими правами [↑](./4_SR_to_SWR.md#swr-api036-up)

##### PUT /api/v1/services/{id} — Обновление сервиса

- <a id="swr-api037"></a>**SWR-API037:** Модуль должен предоставлять endpoint `PUT /api/v1/services/{id}` для обновления данных сервиса [↑](./4_SR_to_SWR.md#swr-api037-up)

- <a id="swr-api038"></a>**SWR-API038:** При конфликте редактирования (данные изменены другим пользователем) API должен возвращать HTTP 409 с телом: `{ detail: "Данные были изменены другим пользователем", current_version: object }` [↑](./4_SR_to_SWR.md#swr-api038-up)

- <a id="swr-api039"></a>**SWR-API039:** Модуль должен использовать поле version (optimistic locking) для обнаружения конфликтов редактирования [↑](./4_SR_to_SWR.md#swr-api039-up)

##### DELETE /api/v1/services/{id} — Удаление сервиса

- <a id="swr-api040"></a>**SWR-API040:** Модуль должен предоставлять endpoint `DELETE /api/v1/services/{id}` для полного удаления сервиса [↑](./4_SR_to_SWR.md#swr-api040-up)

- <a id="swr-api041"></a>**SWR-API041:** При удалении модуль должен удалять все связанные данные: документацию из MinIO, записи обслуживания, права доступа, связи с другими сервисами [↑](./4_SR_to_SWR.md#swr-api041-up)

##### PATCH /api/v1/services/{id}/status — Изменение статуса

- <a id="swr-api042"></a>**SWR-API042:** Модуль должен предоставлять endpoint `PATCH /api/v1/services/{id}/status` для изменения статуса сервиса с JSON payload: `{ status: enum (required) }` [↑](./4_SR_to_SWR.md#swr-api042-up)

- <a id="swr-api043"></a>**SWR-API043:** Допустимые значения статуса: `active`, `archived`, `maintenance`, `development`, `testing` [↑](./4_SR_to_SWR.md#swr-api043-up)

#### Валидация

- <a id="swr-f004"></a>**SWR-F004:** Модуль должен валидировать name: обязательное, уникальное, максимум 100 символов, допустимы латиница, кириллица, цифры, символы #, №, _, - [↑](./4_SR_to_SWR.md#swr-f004-up)

- <a id="swr-f005"></a>**SWR-F005:** Модуль должен валидировать description: обязательное, максимум 1000 символов [↑](./4_SR_to_SWR.md#swr-f005-up)

- <a id="swr-f006"></a>**SWR-F006:** Модуль должен валидировать version: опциональное, формат семантического версионирования (regex: `^\d+\.\d+\.\d+$`) [↑](./4_SR_to_SWR.md#swr-f006-up)

#### Права доступа

##### GET /api/v1/services/{id}/permissions — Получение прав доступа

- <a id="swr-api044"></a>**SWR-API044:** Модуль должен предоставлять endpoint `GET /api/v1/services/{id}/permissions` для получения списка пользователей и их ролей для данного сервиса [↑](./4_SR_to_SWR.md#swr-api044-up)

##### POST /api/v1/services/{id}/permissions — Назначение прав

- <a id="swr-api045"></a>**SWR-API045:** Модуль должен предоставлять endpoint `POST /api/v1/services/{id}/permissions` для назначения прав доступа с JSON payload: `{ user_id: integer (required), role: enum (required) }` [↑](./4_SR_to_SWR.md#swr-api045-up)

- <a id="swr-api046"></a>**SWR-API046:** Допустимые роли в контексте сервиса: `service_admin`, `maintenance_specialist`, `regular_user` [↑](./4_SR_to_SWR.md#swr-api046-up)

##### DELETE /api/v1/services/{id}/permissions/{user_id} — Отзыв прав

- <a id="swr-api047"></a>**SWR-API047:** Модуль должен предоставлять endpoint `DELETE /api/v1/services/{id}/permissions/{user_id}` для отзыва прав доступа [↑](./4_SR_to_SWR.md#swr-api047-up)


### 4.4. API модуля "Категории" (categories)

#### Назначение

Модуль отвечает за управление справочником категорий и шаблонами файловых каталогов.

#### Endpoints

##### GET /api/v1/categories — Список категорий

- <a id="swr-api048"></a>**SWR-API048:** Модуль должен предоставлять endpoint `GET /api/v1/categories` для получения списка категорий [↑](./4_SR_to_SWR.md#swr-api048-up)

##### POST /api/v1/categories — Создание категории

- <a id="swr-api049"></a>**SWR-API049:** Модуль должен предоставлять endpoint `POST /api/v1/categories` для создания категории с JSON payload: `{ name: string (required) }` [↑](./4_SR_to_SWR.md#swr-api049-up)

##### PUT /api/v1/categories/{id} — Обновление категории

- <a id="swr-api050"></a>**SWR-API050:** Модуль должен предоставлять endpoint `PUT /api/v1/categories/{id}` для обновления названия категории [↑](./4_SR_to_SWR.md#swr-api050-up)

##### DELETE /api/v1/categories/{id} — Удаление категории

- <a id="swr-api051"></a>**SWR-API051:** Модуль должен предоставлять endpoint `DELETE /api/v1/categories/{id}` для удаления категории [↑](./4_SR_to_SWR.md#swr-api051-up)

- <a id="swr-api052"></a>**SWR-API052:** При удалении категории существующие сервисы должны сохранять ссылку на категорию (soft reference) [↑](./4_SR_to_SWR.md#swr-api052-up)

#### Шаблоны категорий

##### GET /api/v1/categories/{id}/templates — Список шаблонов

- <a id="swr-api053"></a>**SWR-API053:** Модуль должен предоставлять endpoint `GET /api/v1/categories/{id}/templates` для получения шаблонов категории [↑](./4_SR_to_SWR.md#swr-api053-up)

##### POST /api/v1/categories/{id}/templates — Создание шаблона

- <a id="swr-api054"></a>**SWR-API054:** Модуль должен предоставлять endpoint `POST /api/v1/categories/{id}/templates` для создания шаблона с JSON payload: `{ name: string (required), folder_structure: array[object] (required) }` [↑](./4_SR_to_SWR.md#swr-api054-up)

- <a id="swr-api055"></a>**SWR-API055:** Структура папки в шаблоне должна иметь формат: `{ name: string, type: "user_docs" | "maintenance_docs", children: array[object] (optional) }` [↑](./4_SR_to_SWR.md#swr-api055-up)

##### PUT /api/v1/categories/{id}/templates/{template_id} — Обновление шаблона

- <a id="swr-api056"></a>**SWR-API056:** Модуль должен предоставлять endpoint `PUT /api/v1/categories/{id}/templates/{template_id}` для обновления шаблона [↑](./4_SR_to_SWR.md#swr-api056-up)

- <a id="swr-api057"></a>**SWR-API057:** Изменение шаблона не должно влиять на существующие файловые каталоги сервисов [↑](./4_SR_to_SWR.md#swr-api057-up)

#### Валидация

- <a id="swr-f007"></a>**SWR-F007:** Модуль должен валидировать название категории: обязательное, максимум 50 символов, допустимы латиница, кириллица и пробел [↑](./4_SR_to_SWR.md#swr-f007-up)


### 4.5. API модуля "Документация" (documents)

#### Назначение

Модуль отвечает за загрузку, хранение и управление файлами документации в MinIO.

#### Endpoints

##### GET /api/v1/services/{service_id}/documents — Структура каталога

- <a id="swr-api058"></a>**SWR-API058:** Модуль должен предоставлять endpoint `GET /api/v1/services/{service_id}/documents` для получения структуры файлового каталога сервиса [↑](./4_SR_to_SWR.md#swr-api058-up)

- <a id="swr-api059"></a>**SWR-API059:** Для обычных сотрудников модуль должен возвращать только папки и файлы с типом `user_docs` [↑](./4_SR_to_SWR.md#swr-api059-up)

- <a id="swr-api060"></a>**SWR-API060:** Для специалистов по обслуживанию и администраторов модуль должен возвращать все папки и файлы [↑](./4_SR_to_SWR.md#swr-api060-up)

- <a id="swr-api061"></a>**SWR-API061:** Ответ должен содержать древовидную структуру: `{ folders: array[{ id, name, type, children: array }], files: array[{ id, name, type, size, uploaded_at }] }` [↑](./4_SR_to_SWR.md#swr-api061-up)

##### POST /api/v1/services/{service_id}/documents/folders — Создание папки

- <a id="swr-api062"></a>**SWR-API062:** Модуль должен предоставлять endpoint `POST /api/v1/services/{service_id}/documents/folders` для создания папки с JSON payload: `{ name: string (required), parent_id: integer (optional), type: enum (required) }` [↑](./4_SR_to_SWR.md#swr-api062-up)

- <a id="swr-api063"></a>**SWR-API063:** Тип папки должен быть: `user_docs` (для пользователей) или `maintenance_docs` (для обслуживания) [↑](./4_SR_to_SWR.md#swr-api063-up)

##### DELETE /api/v1/services/{service_id}/documents/folders/{folder_id} — Удаление папки

- <a id="swr-api064"></a>**SWR-API064:** Модуль должен предоставлять endpoint `DELETE /api/v1/services/{service_id}/documents/folders/{folder_id}` для удаления папки [↑](./4_SR_to_SWR.md#swr-api064-up)

- <a id="swr-api065"></a>**SWR-API065:** При удалении папки модуль должен рекурсивно удалять все вложенные папки и файлы из MinIO [↑](./4_SR_to_SWR.md#swr-api065-up)

##### POST /api/v1/services/{service_id}/documents/files — Загрузка файла

- <a id="swr-api066"></a>**SWR-API066:** Модуль должен предоставлять endpoint `POST /api/v1/services/{service_id}/documents/files` для загрузки файла (multipart/form-data) с полями: `file: binary (required), folder_id: integer (optional), type: enum (required)` [↑](./4_SR_to_SWR.md#swr-api066-up)

- <a id="swr-api067"></a>**SWR-API067:** Модуль должен ограничивать максимальный размер загружаемого файла значением 20 МБ [↑](./4_SR_to_SWR.md#swr-api067-up)

- <a id="swr-api068"></a>**SWR-API068:** При превышении размера API должен возвращать HTTP 413 с телом: `{ detail: "Размер файла превышает допустимый лимит 20 МБ" }` [↑](./4_SR_to_SWR.md#swr-api068-up)

- <a id="swr-api069"></a>**SWR-API069:** Модуль должен ограничивать длину имени файла значением 50 символов [↑](./4_SR_to_SWR.md#swr-api069-up)

- <a id="swr-api070"></a>**SWR-API070:** Модуль должен автоматически заменять пробелы в имени файла на символ подчёркивания [↑](./4_SR_to_SWR.md#swr-api070-up)

##### GET /api/v1/services/{service_id}/documents/files/{file_id} — Скачивание файла

- <a id="swr-api071"></a>**SWR-API071:** Модуль должен предоставлять endpoint `GET /api/v1/services/{service_id}/documents/files/{file_id}` для скачивания файла [↑](./4_SR_to_SWR.md#swr-api071-up)

- <a id="swr-api072"></a>**SWR-API072:** Модуль должен генерировать presigned URL для скачивания файла из MinIO со сроком действия 15 минут [↑](./4_SR_to_SWR.md#swr-api072-up)

##### DELETE /api/v1/services/{service_id}/documents/files/{file_id} — Удаление файла

- <a id="swr-api073"></a>**SWR-API073:** Модуль должен предоставлять endpoint `DELETE /api/v1/services/{service_id}/documents/files/{file_id}` для удаления файла [↑](./4_SR_to_SWR.md#swr-api073-up)

##### PATCH /api/v1/services/{service_id}/documents/files/{file_id}/type — Изменение типа документа

- <a id="swr-api074"></a>**SWR-API074:** Модуль должен предоставлять endpoint `PATCH /api/v1/services/{service_id}/documents/files/{file_id}/type` для изменения типа документа [↑](./4_SR_to_SWR.md#swr-api074-up)

#### Безопасность файлов

- <a id="swr-sec007"></a>**SWR-SEC007:** Модуль должен проверять MIME-тип загружаемого файла на соответствие расширению [↑](./4_SR_to_SWR.md#swr-sec007-up)

- <a id="swr-sec008"></a>**SWR-SEC008:** Модуль должен проверять магические байты (file signature) загружаемого файла [↑](./4_SR_to_SWR.md#swr-sec008-up)

- <a id="swr-sec009"></a>**SWR-SEC009:** Модуль должен блокировать загрузку исполняемых файлов: .exe, .bat, .cmd, .sh, .ps1, .msi, .dll, .so [↑](./4_SR_to_SWR.md#swr-sec009-up)

- <a id="swr-sec010"></a>**SWR-SEC010:** При обнаружении потенциальной угрозы API должен возвращать HTTP 400 с телом: `{ detail: "Файл заблокирован по соображениям безопасности" }` [↑](./4_SR_to_SWR.md#swr-sec010-up)

- <a id="swr-sec011"></a>**SWR-SEC011:** Модуль должен позволять администратору сервиса настраивать список разрешённых форматов файлов для каталога сервиса [↑](./4_SR_to_SWR.md#swr-sec011-up)


### 4.6. API модуля "Обслуживание" (maintenance)

#### Назначение

Модуль отвечает за регистрацию и хранение записей обслуживания IT-сервисов.

#### Endpoints

##### GET /api/v1/services/{service_id}/maintenance — История обслуживания

- <a id="swr-api075"></a>**SWR-API075:** Модуль должен предоставлять endpoint `GET /api/v1/services/{service_id}/maintenance` для получения истории обслуживания с сортировкой от новых к старым [↑](./4_SR_to_SWR.md#swr-api075-up)

- <a id="swr-api076"></a>**SWR-API076:** Endpoint должен поддерживать пагинацию: `?page=1&per_page=20` [↑](./4_SR_to_SWR.md#swr-api076-up)

##### POST /api/v1/services/{service_id}/maintenance — Создание записи

- <a id="swr-api077"></a>**SWR-API077:** Модуль должен предоставлять endpoint `POST /api/v1/services/{service_id}/maintenance` для создания записи обслуживания с JSON payload: `{ performed_at: datetime (required), description: string (required), performers: array[string] (required) }` [↑](./4_SR_to_SWR.md#swr-api077-up)

- <a id="swr-api078"></a>**SWR-API078:** Модуль должен автоматически фиксировать пользователя, создавшего запись, и время создания [↑](./4_SR_to_SWR.md#swr-api078-up)

- <a id="swr-api079"></a>**SWR-API079:** Модуль должен автоматически обновлять поле last_maintenance_at сервиса при добавлении записи [↑](./4_SR_to_SWR.md#swr-api079-up)

#### Валидация

- <a id="swr-f008"></a>**SWR-F008:** Модуль должен валидировать description: обязательное, максимум 10000 символов [↑](./4_SR_to_SWR.md#swr-f008-up)

- <a id="swr-f009"></a>**SWR-F009:** Модуль должен валидировать performers: обязательное, массив строк, минимум 1 элемент [↑](./4_SR_to_SWR.md#swr-f009-up)


### 4.7. API модуля "Связи сервисов" (dependencies)

#### Назначение

Модуль отвечает за управление связями между IT-сервисами и визуализацию графа зависимостей.

#### Endpoints

##### GET /api/v1/services/{service_id}/dependencies — Связи сервиса

- <a id="swr-api080"></a>**SWR-API080:** Модуль должен предоставлять endpoint `GET /api/v1/services/{service_id}/dependencies` для получения связей сервиса [↑](./4_SR_to_SWR.md#swr-api080-up)

- <a id="swr-api081"></a>**SWR-API081:** Для специалистов по обслуживанию модуль должен возвращать связи с глубиной 1 уровень [↑](./4_SR_to_SWR.md#swr-api081-up)

- <a id="swr-api082"></a>**SWR-API082:** Для администраторов модуль должен возвращать полный граф связей [↑](./4_SR_to_SWR.md#swr-api082-up)

- <a id="swr-api083"></a>**SWR-API083:** Ответ должен содержать структуру: `{ depends_on: array[{ id, name }], used_by: array[{ id, name }] }` [↑](./4_SR_to_SWR.md#swr-api083-up)

##### POST /api/v1/services/{service_id}/dependencies — Создание связи

- <a id="swr-api084"></a>**SWR-API084:** Модуль должен предоставлять endpoint `POST /api/v1/services/{service_id}/dependencies` для создания связи с JSON payload: `{ depends_on_id: integer (required) }` [↑](./4_SR_to_SWR.md#swr-api084-up)

- <a id="swr-api085"></a>**SWR-API085:** Модуль должен блокировать создание циклических зависимостей и возвращать HTTP 400 [↑](./4_SR_to_SWR.md#swr-api085-up)

##### DELETE /api/v1/services/{service_id}/dependencies/{dependency_id} — Удаление связи

- <a id="swr-api086"></a>**SWR-API086:** Модуль должен предоставлять endpoint `DELETE /api/v1/services/{service_id}/dependencies/{dependency_id}` для удаления связи [↑](./4_SR_to_SWR.md#swr-api086-up)

##### GET /api/v1/dependencies/graph — Полный граф зависимостей

- <a id="swr-api087"></a>**SWR-API087:** Модуль должен предоставлять endpoint `GET /api/v1/dependencies/graph` для получения полного графа зависимостей в формате, совместимом с React Flow: `{ nodes: array[{ id, data, position }], edges: array[{ id, source, target }] }` [↑](./4_SR_to_SWR.md#swr-api087-up)


### 4.8. API модуля "Аудит" (audit)

#### Назначение

Модуль отвечает за ведение журнала действий пользователей.

#### Endpoints

##### GET /api/v1/audit — Журнал действий

- <a id="swr-api088"></a>**SWR-API088:** Модуль должен предоставлять endpoint `GET /api/v1/audit` для получения журнала действий [↑](./4_SR_to_SWR.md#swr-api088-up)

- <a id="swr-api089"></a>**SWR-API089:** Для администраторов системы endpoint должен возвращать все записи журнала [↑](./4_SR_to_SWR.md#swr-api089-up)

- <a id="swr-api090"></a>**SWR-API090:** Для администраторов сервиса endpoint должен возвращать только записи, связанные с их сервисами [↑](./4_SR_to_SWR.md#swr-api090-up)

- <a id="swr-api091"></a>**SWR-API091:** Endpoint должен поддерживать фильтрацию: `?user_id=1&action=create&service_id=5&from=2025-01-01&to=2025-12-31` [↑](./4_SR_to_SWR.md#swr-api091-up)

#### Логируемые события

- <a id="swr-log001"></a>**SWR-LOG001:** Модуль должен регистрировать события входа пользователей: `{ action: "login", user_id, ip, user_agent, timestamp }` [↑](./4_SR_to_SWR.md#swr-log001-up)

- <a id="swr-log002"></a>**SWR-LOG002:** Модуль должен регистрировать операции с сервисами: `{ action: "create|update|delete|archive", entity: "service", entity_id, user_id, changes, timestamp }` [↑](./4_SR_to_SWR.md#swr-log002-up)

- <a id="swr-log003"></a>**SWR-LOG003:** Модуль должен регистрировать операции с документацией: `{ action: "upload|delete", entity: "document", entity_id, service_id, user_id, filename, timestamp }` [↑](./4_SR_to_SWR.md#swr-log003-up)

- <a id="swr-log004"></a>**SWR-LOG004:** Модуль должен регистрировать изменения прав доступа: `{ action: "grant|revoke", entity: "permission", user_id, target_user_id, service_id, role, timestamp }` [↑](./4_SR_to_SWR.md#swr-log004-up)

#### Безопасность журнала

- <a id="swr-sec012"></a>**SWR-SEC012:** Записи журнала должны быть защищены от изменения и удаления (append-only) [↑](./4_SR_to_SWR.md#swr-sec012-up)

- <a id="swr-sec013"></a>**SWR-SEC013:** Таблица журнала не должна иметь endpoint для удаления записей [↑](./4_SR_to_SWR.md#swr-sec013-up)



## 5. Требования к Frontend

### 5.1. Структура компонентов

- <a id="swr-ui001"></a>**SWR-UI001:** Frontend должен быть реализован на React 18.x с TypeScript 5.x [↑](./4_SR_to_SWR.md#swr-ui001-up)

- <a id="swr-ui002"></a>**SWR-UI002:** Сборка проекта должна выполняться с помощью Vite 5.x [↑](./4_SR_to_SWR.md#swr-ui002-up)

- <a id="swr-ui003"></a>**SWR-UI003:** UI компоненты должны использовать библиотеку Chakra UI 2.x [↑](./4_SR_to_SWR.md#swr-ui003-up)

- <a id="swr-ui004"></a>**SWR-UI004:** Роутинг должен быть реализован с помощью React Router 6.x [↑](./4_SR_to_SWR.md#swr-ui004-up)

- <a id="swr-ui005"></a>**SWR-UI005:** HTTP-запросы должны выполняться через Axios 1.x с настроенными interceptors для JWT [↑](./4_SR_to_SWR.md#swr-ui005-up)

### 5.2. Управление состоянием

- <a id="swr-ui006"></a>**SWR-UI006:** Глобальное состояние должно управляться через Zustand 5.x [↑](./4_SR_to_SWR.md#swr-ui006-up)

- <a id="swr-ui007"></a>**SWR-UI007:** Состояние авторизации должно храниться в отдельном store: `{ user, tokens, isAuthenticated, login, logout, refresh }` [↑](./4_SR_to_SWR.md#swr-ui007-up)

- <a id="swr-ui008"></a>**SWR-UI008:** Access token должен храниться в памяти, refresh token — в httpOnly cookie или localStorage [↑](./4_SR_to_SWR.md#swr-ui008-up)

- <a id="swr-ui009"></a>**SWR-UI009:** При получении HTTP 401 приложение должно автоматически пытаться обновить токен через refresh endpoint [↑](./4_SR_to_SWR.md#swr-ui009-up)

### 5.3. Требования к UI компонентам

#### Навигация и layout

- <a id="swr-ui010"></a>**SWR-UI010:** Компонент Layout должен содержать боковую панель навигации с разделами: Реестр сервисов, Категории (для админов), Пользователи (для админов), Журнал действий [↑](./4_SR_to_SWR.md#swr-ui010-up)

- <a id="swr-ui011"></a>**SWR-UI011:** Компонент Breadcrumbs должен отображать текущее местоположение пользователя в иерархии страниц [↑](./4_SR_to_SWR.md#swr-ui011-up)

- <a id="swr-ui012"></a>**SWR-UI012:** Навигация должна скрывать разделы, недоступные пользователю согласно его роли [↑](./4_SR_to_SWR.md#swr-ui012-up)

#### Реестр сервисов

- <a id="swr-ui013"></a>**SWR-UI013:** Компонент ServiceList должен отображать таблицу сервисов с колонками: название, категория, статус, версия, дата изменения [↑](./4_SR_to_SWR.md#swr-ui013-up)

- <a id="swr-ui014"></a>**SWR-UI014:** Компонент должен поддерживать выбор количества записей на странице: 20, 50, 100 через Chakra UI Select [↑](./4_SR_to_SWR.md#swr-ui014-up)

- <a id="swr-ui015"></a>**SWR-UI015:** Компонент должен предоставлять поле поиска с debounce 300ms для фильтрации по названию [↑](./4_SR_to_SWR.md#swr-ui015-up)

- <a id="swr-ui016"></a>**SWR-UI016:** Компонент должен предоставлять фильтры по категории и статусу через Chakra UI Select [↑](./4_SR_to_SWR.md#swr-ui016-up)

- <a id="swr-ui017"></a>**SWR-UI017:** Компонент должен поддерживать сортировку по клику на заголовок колонки [↑](./4_SR_to_SWR.md#swr-ui017-up)

- <a id="swr-ui018"></a>**SWR-UI018:** Состояние фильтров, сортировки и пагинации должно сохраняться в URL query parameters [↑](./4_SR_to_SWR.md#swr-ui018-up)

#### Детальная страница сервиса

- <a id="swr-ui019"></a>**SWR-UI019:** Компонент ServiceDetail должен отображать все атрибуты сервиса с возможностью редактирования для пользователей с правами [↑](./4_SR_to_SWR.md#swr-ui019-up)

- <a id="swr-ui020"></a>**SWR-UI020:** Компонент должен отображать вкладки: Информация, Документация, Обслуживание, Связи, Права доступа [↑](./4_SR_to_SWR.md#swr-ui020-up)

- <a id="swr-ui021"></a>**SWR-UI021:** Компонент должен скрывать кнопки редактирования для пользователей без прав на редактирование [↑](./4_SR_to_SWR.md#swr-ui021-up)

#### Файловый каталог

- <a id="swr-ui022"></a>**SWR-UI022:** Компонент DocumentTree должен отображать структуру каталога в виде дерева с иконками папок и файлов [↑](./4_SR_to_SWR.md#swr-ui022-up)

- <a id="swr-ui023"></a>**SWR-UI023:** Компонент должен поддерживать drag-and-drop загрузку файлов [↑](./4_SR_to_SWR.md#swr-ui023-up)

- <a id="swr-ui024"></a>**SWR-UI024:** Компонент должен отображать прогресс загрузки файла через Chakra UI Progress [↑](./4_SR_to_SWR.md#swr-ui024-up)

- <a id="swr-ui025"></a>**SWR-UI025:** Компонент должен визуально различать типы документов: для пользователей (синий) и для обслуживания (оранжевый) [↑](./4_SR_to_SWR.md#swr-ui025-up)

#### Граф зависимостей

- <a id="swr-ui026"></a>**SWR-UI026:** Компонент DependencyGraph должен визуализировать связи между сервисами с помощью React Flow [↑](./4_SR_to_SWR.md#swr-ui026-up)

- <a id="swr-ui027"></a>**SWR-UI027:** Компонент должен поддерживать zoom, pan и автоматическое расположение узлов [↑](./4_SR_to_SWR.md#swr-ui027-up)

- <a id="swr-ui028"></a>**SWR-UI028:** При клике на узел графа должен происходить переход на страницу соответствующего сервиса [↑](./4_SR_to_SWR.md#swr-ui028-up)

#### Обратная связь

- <a id="swr-ui029"></a>**SWR-UI029:** При успешном выполнении операции должно отображаться toast-уведомление (Chakra UI Toast) зелёного цвета [↑](./4_SR_to_SWR.md#swr-ui029-up)

- <a id="swr-ui030"></a>**SWR-UI030:** При ошибке должно отображаться toast-уведомление красного цвета с описанием ошибки [↑](./4_SR_to_SWR.md#swr-ui030-up)

- <a id="swr-ui031"></a>**SWR-UI031:** Поля формы с ошибками валидации должны подсвечиваться красной рамкой с текстом ошибки под полем [↑](./4_SR_to_SWR.md#swr-ui031-up)

- <a id="swr-ui032"></a>**SWR-UI032:** Перед удалением сервиса или файла должен отображаться модальный диалог подтверждения (Chakra UI AlertDialog) [↑](./4_SR_to_SWR.md#swr-ui032-up)

#### Доступность

- <a id="swr-ui033"></a>**SWR-UI033:** Все интерактивные элементы должны быть доступны с клавиатуры (tab navigation) [↑](./4_SR_to_SWR.md#swr-ui033-up)

- <a id="swr-ui034"></a>**SWR-UI034:** Доступ к любой функции должен быть возможен не более чем за 3 клика от главной страницы [↑](./4_SR_to_SWR.md#swr-ui034-up)



## 6. Требования к данным

### 6.1. Схема базы данных

#### Таблица users

- <a id="swr-db001"></a>**SWR-DB001:** Таблица `users` должна иметь следующую структуру:

| Поле            | Тип          | Ограничения             |
| --------------- | ------------ | ----------------------- |
| id              | serial       | primary key             |
| username        | varchar(50)  | not null, unique        |
| email           | varchar(255) | not null, unique        |
| password_hash   | varchar(255) | not null                |
| first_name      | varchar(100) | not null                |
| last_name       | varchar(100) | not null                |
| position        | varchar(100) | not null                |
| department      | varchar(100) | not null                |
| is_system_admin | boolean      | not null, default false |
| is_blocked      | boolean      | not null, default false |
| created_at      | timestamp    | not null, default now() |
| updated_at      | timestamp    | not null, default now() |

[↑](./4_SR_to_SWR.md#swr-db001-up)

#### Таблица refresh_tokens

- <a id="swr-db002"></a>**SWR-DB002:** Таблица `refresh_tokens` должна иметь следующую структуру:

| Поле         | Тип          | Ограничения                    |
| ------------ | ------------ | ------------------------------ |
| id           | serial       | primary key                    |
| user_id      | integer      | not null, foreign key users.id |
| token_hash   | varchar(255) | not null, unique               |
| expires_at   | timestamp    | not null                       |
| last_used_at | timestamp    | not null                       |
| created_at   | timestamp    | not null, default now()        |

[↑](./4_SR_to_SWR.md#swr-db002-up)

#### Таблица categories

- <a id="swr-db003"></a>**SWR-DB003:** Таблица `categories` должна иметь следующую структуру:

| Поле       | Тип         | Ограничения             |
| ---------- | ----------- | ----------------------- |
| id         | serial      | primary key             |
| name       | varchar(50) | not null, unique        |
| created_at | timestamp   | not null, default now() |

[↑](./4_SR_to_SWR.md#swr-db003-up)

#### Таблица category_templates

- <a id="swr-db004"></a>**SWR-DB004:** Таблица `category_templates` должна иметь следующую структуру:

| Поле             | Тип          | Ограничения                         |
| ---------------- | ------------ | ----------------------------------- |
| id               | serial       | primary key                         |
| category_id      | integer      | not null, foreign key categories.id |
| name             | varchar(100) | not null                            |
| folder_structure | jsonb        | not null                            |
| created_at       | timestamp    | not null, default now()             |

[↑](./4_SR_to_SWR.md#swr-db004-up)

#### Таблица services

- <a id="swr-db005"></a>**SWR-DB005:** Таблица `services` должна иметь следующую структуру:

| Поле                | Тип           | Ограничения                                                              |
| ------------------- | ------------- | ------------------------------------------------------------------------ |
| id                  | serial        | primary key                                                              |
| name                | varchar(100)  | not null, unique                                                         |
| description         | varchar(1000) | not null                                                                 |
| category_id         | integer       | foreign key categories.id                                                |
| status              | varchar(20)   | not null, check in (active, archived, maintenance, development, testing) |
| version             | varchar(20)   | nullable                                                                 |
| hardware_ids        | text[]        | nullable                                                                 |
| storage_bucket      | varchar(255)  | not null                                                                 |
| allowed_file_types  | text[]        | nullable                                                                 |
| last_maintenance_at | timestamp     | nullable                                                                 |
| row_version         | integer       | not null, default 1                                                      |
| created_at          | timestamp     | not null, default now()                                                  |
| updated_at          | timestamp     | not null, default now()                                                  |

[↑](./4_SR_to_SWR.md#swr-db005-up)

- <a id="swr-db006"></a>**SWR-DB006:** Таблица `services` должна иметь индексы: `unique(name)`, `index(category_id)`, `index(status)` [↑](./4_SR_to_SWR.md#swr-db006-up)

#### Таблица service_permissions

- <a id="swr-db007"></a>**SWR-DB007:** Таблица `service_permissions` должна иметь следующую структуру:

| Поле       | Тип         | Ограничения                                                              |
| ---------- | ----------- | ------------------------------------------------------------------------ |
| id         | serial      | primary key                                                              |
| service_id | integer     | not null, foreign key services.id                                        |
| user_id    | integer     | not null, foreign key users.id                                           |
| role       | varchar(30) | not null, check in (service_admin, maintenance_specialist, regular_user) |
| created_at | timestamp   | not null, default now()                                                  |

[↑](./4_SR_to_SWR.md#swr-db007-up)

- <a id="swr-db008"></a>**SWR-DB008:** Таблица `service_permissions` должна иметь уникальный индекс: `unique(service_id, user_id)` [↑](./4_SR_to_SWR.md#swr-db008-up)

#### Таблица document_folders

- <a id="swr-db009"></a>**SWR-DB009:** Таблица `document_folders` должна иметь следующую структуру:

| Поле       | Тип          | Ограничения                                      |
| ---------- | ------------ | ------------------------------------------------ |
| id         | serial       | primary key                                      |
| service_id | integer      | not null, foreign key services.id                |
| parent_id  | integer      | nullable, foreign key document_folders.id        |
| name       | varchar(100) | not null                                         |
| type       | varchar(20)  | not null, check in (user_docs, maintenance_docs) |
| created_at | timestamp    | not null, default now()                          |

[↑](./4_SR_to_SWR.md#swr-db009-up)

#### Таблица document_files

- <a id="swr-db010"></a>**SWR-DB010:** Таблица `document_files` должна иметь следующую структуру:

| Поле        | Тип          | Ограничения                                      |
| ----------- | ------------ | ------------------------------------------------ |
| id          | serial       | primary key                                      |
| service_id  | integer      | not null, foreign key services.id                |
| folder_id   | integer      | nullable, foreign key document_folders.id        |
| name        | varchar(50)  | not null                                         |
| storage_key | varchar(500) | not null                                         |
| mime_type   | varchar(100) | not null                                         |
| size_bytes  | bigint       | not null                                         |
| type        | varchar(20)  | not null, check in (user_docs, maintenance_docs) |
| uploaded_by | integer      | not null, foreign key users.id                   |
| uploaded_at | timestamp    | not null, default now()                          |

[↑](./4_SR_to_SWR.md#swr-db010-up)

#### Таблица maintenance_records

- <a id="swr-db011"></a>**SWR-DB011:** Таблица `maintenance_records` должна иметь следующую структуру:

| Поле         | Тип       | Ограничения                       |
| ------------ | --------- | --------------------------------- |
| id           | serial    | primary key                       |
| service_id   | integer   | not null, foreign key services.id |
| performed_at | timestamp | not null                          |
| description  | text      | not null                          |
| performers   | text[]    | not null                          |
| created_by   | integer   | not null, foreign key users.id    |
| created_at   | timestamp | not null, default now()           |

[↑](./4_SR_to_SWR.md#swr-db011-up)

#### Таблица service_dependencies

- <a id="swr-db012"></a>**SWR-DB012:** Таблица `service_dependencies` должна иметь следующую структуру:

| Поле          | Тип       | Ограничения                       |
| ------------- | --------- | --------------------------------- |
| id            | serial    | primary key                       |
| service_id    | integer   | not null, foreign key services.id |
| depends_on_id | integer   | not null, foreign key services.id |
| created_at    | timestamp | not null, default now()           |

[↑](./4_SR_to_SWR.md#swr-db012-up)

- <a id="swr-db013"></a>**SWR-DB013:** Таблица `service_dependencies` должна иметь уникальный индекс: `unique(service_id, depends_on_id)` [↑](./4_SR_to_SWR.md#swr-db013-up)

#### Таблица audit_log

- <a id="swr-db014"></a>**SWR-DB014:** Таблица `audit_log` должна иметь следующую структуру:

| Поле        | Тип         | Ограничения                       |
| ----------- | ----------- | --------------------------------- |
| id          | serial      | primary key                       |
| user_id     | integer     | nullable, foreign key users.id    |
| action      | varchar(50) | not null                          |
| entity_type | varchar(50) | not null                          |
| entity_id   | integer     | nullable                          |
| service_id  | integer     | nullable, foreign key services.id |
| details     | jsonb       | nullable                          |
| ip_address  | inet        | nullable                          |
| user_agent  | text        | nullable                          |
| created_at  | timestamp   | not null, default now()           |

[↑](./4_SR_to_SWR.md#swr-db014-up)

- <a id="swr-db015"></a>**SWR-DB015:** Таблица `audit_log` должна иметь индексы: `index(user_id)`, `index(service_id)`, `index(action)`, `index(created_at)` [↑](./4_SR_to_SWR.md#swr-db015-up)

### 6.2. Миграции и хранение

- <a id="swr-db016"></a>**SWR-DB016:** В качестве СУБД должна использоваться PostgreSQL версии 16.x или выше [↑](./4_SR_to_SWR.md#swr-db016-up)

- <a id="swr-db017"></a>**SWR-DB017:** Миграции должны управляться через Alembic 1.13.x [↑](./4_SR_to_SWR.md#swr-db017-up)

- <a id="swr-db018"></a>**SWR-DB018:** Каждая миграция должна содержать методы upgrade() и downgrade() [↑](./4_SR_to_SWR.md#swr-db018-up)

- <a id="swr-db019"></a>**SWR-DB019:** Файлы документации должны храниться в MinIO с организацией по bucket: один bucket на сервис [↑](./4_SR_to_SWR.md#swr-db019-up)

- <a id="swr-db020"></a>**SWR-DB020:** Имя bucket должно формироваться по шаблону: `service-{service_id}` [↑](./4_SR_to_SWR.md#swr-db020-up)



## 7. Требования к инфраструктуре

### 7.1. Развёртывание

- <a id="swr-dep001"></a>**SWR-DEP001:** Все компоненты системы должны быть контейнеризированы с использованием Docker [↑](./4_SR_to_SWR.md#swr-dep001-up)

- <a id="swr-dep002"></a>**SWR-DEP002:** Для локальной разработки и развёртывания должен использоваться Docker Compose с сервисами: backend, frontend, postgres, minio, mailhog [↑](./4_SR_to_SWR.md#swr-dep002-up)

- <a id="swr-dep003"></a>**SWR-DEP003:** Docker Compose должен использовать volumes для персистентного хранения данных PostgreSQL и MinIO [↑](./4_SR_to_SWR.md#swr-dep003-up)

- <a id="swr-dep004"></a>**SWR-DEP004:** Frontend должен собираться в production-режиме и раздаваться через nginx [↑](./4_SR_to_SWR.md#swr-dep004-up)

### 7.2. Конфигурация

- <a id="swr-config001"></a>**SWR-CONFIG001:** Конфигурация Backend должна осуществляться через переменные окружения:

| Переменная                | Описание                       | Пример                                      |
| ------------------------- | ------------------------------ | ------------------------------------------- |
| DATABASE_URL              | URL подключения к PostgreSQL   | postgresql+asyncpg://user:pass@host:5432/db |
| JWT_SECRET                | Секретный ключ для подписи JWT | random-secret-string                        |
| JWT_ACCESS_EXPIRE_MINUTES | Время жизни access token       | 30                                          |
| JWT_REFRESH_EXPIRE_DAYS   | Время жизни refresh token      | 7                                           |
| MINIO_ENDPOINT            | Адрес MinIO                    | minio:9000                                  |
| MINIO_ACCESS_KEY          | Access key MinIO               | minioadmin                                  |
| MINIO_SECRET_KEY          | Secret key MinIO               | minioadmin                                  |
| SMTP_HOST                 | SMTP сервер                    | mailhog                                     |
| SMTP_PORT                 | SMTP порт                      | 1025                                        |
| CORS_ORIGINS              | Разрешённые origins            | http://localhost:3000                       |

[↑](./4_SR_to_SWR.md#swr-config001-up)

- <a id="swr-config002"></a>**SWR-CONFIG002:** Секреты (JWT_SECRET, DATABASE_URL, MINIO_SECRET_KEY) не должны храниться в репозитории [↑](./4_SR_to_SWR.md#swr-config002-up)

- <a id="swr-config003"></a>**SWR-CONFIG003:** Для локальной разработки должен использоваться файл .env (добавлен в .gitignore) с шаблоном .env.example [↑](./4_SR_to_SWR.md#swr-config003-up)

### 7.3. Логирование и мониторинг

- <a id="swr-log005"></a>**SWR-LOG005:** Логи должны выводиться в формате JSON с использованием библиотеки structlog [↑](./4_SR_to_SWR.md#swr-log005-up)

- <a id="swr-log006"></a>**SWR-LOG006:** Каждая запись лога должна содержать поля: `timestamp`, `level`, `message`, `request_id`, `user_id` (если авторизован) [↑](./4_SR_to_SWR.md#swr-log006-up)

- <a id="swr-log007"></a>**SWR-LOG007:** Должны поддерживаться уровни логирования: DEBUG, INFO, WARNING, ERROR, CRITICAL [↑](./4_SR_to_SWR.md#swr-log007-up)

- <a id="swr-log008"></a>**SWR-LOG008:** Уровень логирования должен настраиваться через переменную окружения LOG_LEVEL (по умолчанию INFO) [↑](./4_SR_to_SWR.md#swr-log008-up)

- <a id="swr-log009"></a>**SWR-LOG009:** Каждый HTTP-запрос должен логироваться с указанием метода, URL, статуса ответа и времени выполнения [↑](./4_SR_to_SWR.md#swr-log009-up)



## 8. Требования к тестированию

- <a id="swr-test001"></a>**SWR-TEST001:** Unit-тесты Backend должны выполняться с использованием pytest 8.x [↑](./4_SR_to_SWR.md#swr-test001-up)

- <a id="swr-test002"></a>**SWR-TEST002:** Покрытие кода Backend unit-тестами должно составлять не менее 60% [↑](./4_SR_to_SWR.md#swr-test002-up)

- <a id="swr-test003"></a>**SWR-TEST003:** Для тестирования API должна использоваться библиотека httpx с AsyncClient [↑](./4_SR_to_SWR.md#swr-test003-up)

- <a id="swr-test004"></a>**SWR-TEST004:** Unit-тесты Frontend должны выполняться с использованием Jest 29.x и React Testing Library 14.x [↑](./4_SR_to_SWR.md#swr-test004-up)

- <a id="swr-test005"></a>**SWR-TEST005:** Покрытие кода Frontend unit-тестами должно составлять не менее 50% [↑](./4_SR_to_SWR.md#swr-test005-up)

- <a id="swr-test006"></a>**SWR-TEST006:** Тесты должны использовать изолированную тестовую базу данных PostgreSQL [↑](./4_SR_to_SWR.md#swr-test006-up)

- <a id="swr-test007"></a>**SWR-TEST007:** Тесты должны использовать mock для MinIO через библиотеку moto [↑](./4_SR_to_SWR.md#swr-test007-up)



## 9. Связанные документы

### Артефакты проекта

- [Системные требования](./3_system_requirements.md) — исходный документ для трассировки требований
- [Техническое задание](./1_technical_specification.md) — требования заказчика на бизнес-уровне

### Трассировка требований

- [Трассировка SR → SWR](./4_SR_to_SWR.md) — связь Системных требований с Требованиями к ПО

### Правила оформления

- [Общие правила оформления артефактов](./_formatting/general_artifact_rules.md) — базовые требования к структуре и формату артефактов
- [Правила оформления Требований к ПО](./_formatting/3_software_requirements_rules.md) — специфичные правила для данного типа документа

### Валидация и проверка качества

- [README валидации](../validations/README.md) — организация проверок и формальных инспекций артефактов
