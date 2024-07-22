package logging

var (
	//ExecutionError ошибки, возникающие в процессе выполнения задачи или функции.
	ExecutionError = "ExecutionError"

	//ParsingError ошибки, связанные с разбором данных (например, при анализе JSON, XML и т.д.).
	ParsingError = "ParsingError"

	//DatabaseError проблемы, связанные с базой данных (например, ошибка подключения, запроса).
	DatabaseError = "DatabaseError"

	//NetworkError ошибки, связанные с сетевыми операциями (например, проблемы с подключением).
	NetworkError = "NetworkError"

	//IOError ошибки ввода-вывода, например, при работе с файловой системой.
	IOError = "IOError"

	//ConfigurationError ошибки, связанные с конфигурацией приложения.
	ConfigurationError = "ConfigurationError"

	//AuthenticationError ошибки аутентификации или авторизации.
	AuthenticationError = "AuthenticationError"

	//TimeoutError ошибки, вызванные превышением времени ожидания.
	TimeoutError = "TimeoutError"

	//APIError ошибки, связанные с внешними API.
	APIError = "APIError"
)
