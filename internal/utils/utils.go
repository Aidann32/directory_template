package utils

import "os"

func DoesDirectoryExist(path string) (bool, error) {
	// Получаем информацию о пути
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // Директория не существует
		}
		return false, err // Ошибка, отличная от отсутствия файла
	}

	// Проверяем, является ли путь директорией
	if info.IsDir() {
		return true, nil // Директория существует
	}
	return false, nil // Путь существует, но это не директория
}
