#
# Makefile с правилами сборки приложения
#

# Создаем переменную с именем проекта
# Тут это не особо оправдано, но для практики можно
PROJECTNAME=$(shell basename "$(PWD)")

# Цель all - алиас для build
# Стандартная цель all используется в Unix-like системах для сборки всего проекта
all: build

# Цель test - тестирование
test:
		@echo "Testing..."
		go test

# Цель check - проверка исходников
check:
		@echo "Checking sources..."
		go vet
		golint

# Цель build - Сборка бинарного файла
# Т.к. в отличие от C/C++ контроль зависимостей файлов возложен go build,
# мы не используем эту возможность Makefile, а просто запускаем go build
build:
		@echo "Building binary..."
		go build -o $(PROJECTNAME)
