package settings

import (
	"fmt"
)

var (
	BASE_URL      string = "https://moikrug.ru"
	COMPANIES_URL string = "https://moikrug.ru/companies?with_vacancies=1"
	
	RABBITMQ_USER string = "rabbitmq"
	RABBITMQ_PASSWORD string = "rabbitmq"
	RABBITMQ_HOST string = "rabbitmq"
	RABBITMQ_PORT string = "5672"

	COMPANIES_QUEUE = "companies"

	RABBITMQ_URL string = fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		RABBITMQ_USER,
		RABBITMQ_PASSWORD,
		RABBITMQ_HOST,
		RABBITMQ_PORT,
	)


)

