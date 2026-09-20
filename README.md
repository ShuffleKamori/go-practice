Проект реализован в учебной цели изучения Go, и не использовался как либо

Для запуска desktop введите в cmd
cd "путь к проекту"
go run .

Для запуска сервера введите в cmd
cd "путь к проекту\server"
.\.venv\Scripts\python -m uvicorn main:app --reload --port 8010

desktop реализован без использования ИИ как либо
Серверная часть реализована через ИИ

Структура проекта
go practice/
├── go.mod                    # модуль github.com/ShuffleKamori/go-practice
├── go.sum
├── main.go                   # точка входа: backend.Start_backend() + gui.Start_gui()
├── main.exe                  # собранный бинарник (можно игнорировать)
│
├── backend/
│   └── backend.go            # HTTP-клиент к FastAPI: GET_CATALOG → GET /products → []Config.Buycard
│
├── GUI/
│   └── GUI.go                # весь интерфейс Fyne: окно, вкладки, каталог, пагинация,
│                             #   фильтры, кликабельные карточки, кнопки
│
├── config/
│   └── Config.go             # типы: WindowApp (настройки окна), Buycard (товар)
│
├── events/
│   └── events.go             # каналы Request/Result, SendRequest(), CheckResult()
│
├── background/
│   ├── background1.jpg       # фон главного окна
│   ├── background11.jpg
│   └── background12.jpg
│
├── BuyCard/
│   └── Recycle.jpg           # картинка товара на карточках
│
├── server/                   # бэкенд-сервер (Python)
│   ├── main.py               # FastAPI + SQLModel: модель Product, сид 16 товаров, CRUD
│   ├── requirements.txt      # fastapi, sqlmodel, uvicorn[standard]
│   ├── store.db              # SQLite БД (создаётся автоматически при старте)
│   ├── .venv/                # виртуальное окружение (не трогать)
│   └── __pycache__/
│
└── .git/                     # git-репозиторий
