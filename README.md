# ITSIMS - IT Service Inventory and Management System

Система учёта IT-сервисов, обеспечивающая хранение досье и инструкций по их использованию и сопровождению

---

## Описание проекта

ITSIMS - это система для централизованного учёта IT-сервисов организации, предоставляющая сотрудникам структурированную информацию о них. Система обеспечивает хранение досье сервисов, инструкций по их использованию и сопровождению, а также управление доступом к этой информации в соответствии с ролевой моделью.

## Структура репозитория

- `general_rules/` — общие правила и форматы документов
  - `unified_document_format.md` — единый формат документов
  - `decision_document_format.md` — решение о формате документов
- `artifacts/` — готовые артефакты проекта
  - `1_technical.md` — ТЗ, Техническая спецификация
  - `2_system_requirements.md` — Системные требования
  - `_formatting/` — правила и шаблоны для артефактов
    - `1_technical_rules.md` — правила оформления ТЗ
    - `1_technical_template.md` — шаблон ТЗ
- `planning/` — планирование и регламенты
  - `configuration_management_plan.md` — план конфигурационного управления
- `validations/` — отчеты о верификации и формальных инспекциях
  - `formal_inspection_regulation.md` — регламент формальной инспекции
  - `_formatting/` — шаблоны для формальных инспекций
  - `reports/` — отчеты инспекций
- `tools/` — скрипты и ресурсы для сборки/конвертации
  - `printing/` — инструменты печати
    - `instructions/` — инструкции по печати
    - `resources/` — ресурсы для конвертации

## Участники команды

- **Тимлид:** Баранов А.Т. (@retrobanner)
- **Ответственный за верификацию:** Зацепилин А.В. (@Anton Zatsepilin)
- **Ответственный за хранилище:** Гречко И.В. (@Ilya Grechko)
- **Ответственный за форматирование:** Когановский Г.И. (@Gregory K)
- **Автор документации:** Вальковец Д.И. (@Данила Вальковец)
- **Ответственный за план КУ:** Носов А.А. (@Never mind)
- **Специалист по поддержке:** Ловков К.И. (@Kirill_Vinill)

## Документация

Ключевые документы:

- ТЗ/Техническая спецификация: `artifacts/1_technical.md`
- Системные требования: `artifacts/2_system_requirements.md`
- План конфигурационного управления: `planning/configuration_management_plan.md`
- Единый формат документов: `general_rules/unified_document_format.md`
- Решение о формате документов: `general_rules/decision_document_format.md`
- Шаблон ТЗ: `artifacts/_formatting/1_technical_template.md`
- Правила оформления ТЗ: `artifacts/_formatting/1_technical_rules.md`
- Регламент формальной инспекции: `validations/formal_inspection_regulation.md`
- Шаблон отчёта формальной инспекции: `validations/_formatting/formal_inspection_report_template.pdf`
- Руководство по конвертации в PDF: `tools/printing/instructions/markdown_to_pdf_guide.md`

## Контакты

Для вопросов по проекту обращайтесь к тимлиду: Баранов А.Т. (@retrobanner)
