# ITSIMS - IT Service Inventory and Management System

Система учёта IT-сервисов, обеспечивающая хранение досье и инструкций по их использованию и сопровождению

---

## Описание проекта

ITSIMS - это система для централизованного учёта IT-сервисов организации, предоставляющая сотрудникам структурированную информацию о них. Система обеспечивает хранение досье сервисов, инструкций по их использованию и сопровождению, а также управление доступом к этой информации в соответствии с ролевой моделью.

## Структура репозитория

- `planning/` — планирование и регламенты
  - `rule_configuration_management.md` — план конфигурационного управления
  - `formatting/` — стандарты форматирования и шаблоны
    - `unified_document_format.md` — единый формат документов
    - `decision_document_format.md` — решение о формате документов
    - `spec/` — правила и шаблоны для ТЗ
    - `formal_inspection/` — шаблоны для формальных инспекций
    - `printing/` — правила подготовки к печати
- `deliverables/` — готовые результаты
  - `specs/tech_spec/` — ТЗ, Технические спецификации
  - `requirements/` — Системные требования и др.
  - `code/` — исходный код
- `tools/` — скрипты и ресурсы для сборки/конвертации

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

- ТЗ/Техническая спецификация: `deliverables/specs/tech_spec/spec_technical.md`
- Системные требования (MD): `deliverables/requirements/spec_requirements.md`
- Системные требования (PDF): `deliverables/requirements/spec_system_requirements.pdf`
- План конфигурационного управления: `planning/rule_configuration_management.md`
- Единый формат документов: `planning/formatting/unified_document_format.md`
- Решение о формате документов: `planning/formatting/decision_document_format.md`
- Шаблон ТЗ: `planning/formatting/spec/spec_technical_template.md`
- Правила оформления ТЗ: `planning/formatting/spec/rules_spec_formatting.md`
- Шаблон формальной инспекции: `planning/formatting/formal_inspection/formal_inspection_report_template.pdf`
- Руководство по конвертации в PDF: `planning/formatting/printing/markdown_to_pdf_guide.md`

## Контакты

Для вопросов по проекту обращайтесь к тимлиду: Баранов А.Т. (@retrobanner)
