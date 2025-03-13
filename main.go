package main

import "fmt"

type bookmarkMap = map[string]string // Новый тип

func main() {

	fmt.Println("Приложения для закладок")
	
	bookmarks := bookmarkMap {} // Делаем  map

Menu: // Лейбел

	for {

		variant := getMenu()

		switch variant{ 

		case 1:
			printBookMarks(bookmarks)

		case 2:
			addBookMark(bookmarks)

		case 3:
			deleteBookMark(bookmarks)

		case 4:
			break Menu
			
		}

	}

}

func getMenu() int {

	var variant int

	fmt.Println(`
Меню :
1. Посмотреть закладки
2. Добавить закладку 
3. Удалить закладку 
4. Выход
				`) 

	fmt.Scan(&variant)
	return variant

}

func printBookMarks(bookmarks bookmarkMap) {
	
	if len(bookmarks) == 0 {

		fmt.Println("Пока нет закладок")

	}

	for key, value := range bookmarks {

		fmt.Println(key, ": ", value)

	}

}

func addBookMark(bookmarks bookmarkMap) { // Возвращает map[string]string

	var newBookmarkKey string
	var newBookmarkValue string

	fmt.Println("Введите название: ")
	fmt.Scan(&newBookmarkKey)

	fmt.Println("Введите ссылку: ")
	fmt.Scan(&newBookmarkValue)
	
	bookmarks[newBookmarkKey] = newBookmarkValue // Добавляем значение

}

func deleteBookMark(bookmarks bookmarkMap) {

	var newBookmarkKeyToDelete string
	
	fmt.Println("Введите название: ")
	fmt.Scan(&newBookmarkKeyToDelete)

	delete(bookmarks, newBookmarkKeyToDelete)

}