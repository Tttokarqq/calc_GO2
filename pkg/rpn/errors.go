package rpn

import (
	"errors"
)

var(
	err_skobk = errors.New("ошибка в записи скобок")
	err_symbl = errors.New("ошибка - непредвиденный сивол")
	err_znak = errors.New("ошибка в записи знаков")
	err_float = errors.New("ошибка при обработке дробных значений")
	Err_acc = errors.New("некорректное число точности. Необходимо целое из отрезка: [0;64]")
	Err_no_post = errors.New("нужен запрос типа post")
	Err_float_write = errors.New("ошибка в записи дробных чисел")
)