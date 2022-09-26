Для рассмотрения Go concurrency patterns: Context создаётся пример http сервера, который:
1) принимает запрос /search?q=...&timeout=...s
2) фордирует запрос в Google Web Search API и отображает результат