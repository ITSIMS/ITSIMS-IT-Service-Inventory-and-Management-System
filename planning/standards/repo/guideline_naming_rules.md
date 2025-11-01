# Правила наименования файлов
## 1. Общие принципы
1. Все имена файлов пишутся **на английском языке**.
2. Используется **snake_case** — слова разделяются нижним подчёркиванием.
3. В именах **не используются**:
    - пробелы;
    - кириллица;
    - специальные символы (`!@#$%^&*()+=[]{};'",<>/?`);
4. Название файла должно **отражать его содержание**, а не формат или дату.
5. **Версии и даты не указываются** в названии — их отслеживает Git (история коммитов).
6. Расширение файла — стандартное: `.md`, `.docx`, `.pdf`, `.tsx`, `.png`, `.xlsx` и т.п.
## 2. Структура имени файла

```php
<content_type>_<short_description>.<extension>
```
где:
- `content_type` — тип документа (например, `report`, `diagram`, `spec`);
- `short_description` — краткое описание содержимого;
- `<extension>` — формат файла.
### Примеры

spec_requirements.docx
diagram_er_model.png
report_team_feedback.pdf
component_question_editor.tsx

## 3. Примеры типов документов
| Тип           | Назначение / содержание                                | Примеры файлов                                                 |
| ------------- | ------------------------------------------------------ | -------------------------------------------------------------- |
| **spec**      | Спецификация, техническое задание, требования          | `spec_requirements.docx`, `spec_functional.md`                 |
| **report**    | Отчёты, аналитические документы, результаты проверки   | `report_verification_results.pdf`, `report_analysis.docx`      |
| **diagram**   | Диаграммы, схемы, графические материалы                | `diagram_er_model.png`, `diagram_architecture.drawio`          |
| **model**     | Формальные модели (данных, классов, процессов)         | `model_restaurant_domain.md`, `model_user_roles.png`           |
| **plan**      | Планы, дорожные карты, расписания                      | `plan_development.docx`, `plan_verification.xlsx`              |
| **note**      | Небольшие пояснительные документы, внутренние заметки  | `note_discussion_summary.md`                                   |
| **component** | Элементы исходного кода (фронтенд, backend, утилиты)   | `component_question_editor.tsx`, `component_drag_handler.ts`   |
| **data**      | Исходные данные, тестовые примеры, CSV- или JSON-файлы | `data_test_cases.json`, `data_menu_items.csv`                  |
| **form**      | Шаблоны и формы для заполнения (инспекции, отчёты)     | `form_inspection_checklist.docx`                               |
| **guideline** | Руководства, стандарты, правила оформления             | `guideline_coding_style.md`, `guideline_document_structure.md` |
