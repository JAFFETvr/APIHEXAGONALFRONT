package equipment

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

// Instancia global para mantener la conexión y el canal abiertos
var rabbitMQInstance *RabbitMQ

func init() {
	var err error
	rabbitMQInstance, err = NewRabbitMQ()
	if err != nil {
		log.Fatal("Error al conectar a RabbitMQ:", err)
	}
}

// Nueva función para verificar si el canal está abierto (por medio del error de canal cerrado)
func (r *RabbitMQ) checkChannel() error {
	if r.channel == nil || r.channel.IsClosed() {
		log.Println("El canal está cerrado. Intentando reabrir la conexión...")
		var err error
		rabbitMQInstance, err = NewRabbitMQ() // Reabrir conexión
		if err != nil {
			log.Println("No se pudo reabrir la conexión:", err)
			return err
		}
	}
	return nil
}

func NewRabbitMQ() (*RabbitMQ, error) {
	// Cargar las variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env", err)
		return nil, err
	}

	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	url := "amqp://" + user + ":" + password + "@" + host + ":" + port
	log.Println("Conectando a RabbitMQ en:", url)

	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Println("Error al conectar a RabbitMQ:", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Error al abrir un canal:", err)
		return nil, err
	}

	return &RabbitMQ{conn: conn, channel: ch}, nil
}

// Método para enviar un mensaje
func (r *RabbitMQ) SendMessage(queue string, message string) error {
	// Verificar si el canal está cerrado
	err := r.checkChannel()
	if err != nil {
		return err
	}

	// Declarar la cola
	_, err = r.channel.QueueDeclare(
		queue,
		true,  // Durabilidad
		false, // Excluir cuando no haya consumidores
		false, // Exclusividad
		false, // Auto-borrado
		nil,   // Parámetros adicionales
	)
	if err != nil {
		log.Println("Error al declarar la cola:", err)
		return err
	}

	// Convertir mensaje a bytes
	body := []byte(message)

	// Enviar mensaje
	err = r.channel.Publish(
		"",       // Exchange por defecto
		queue,    // Nombre de la cola
		false,    // Confirmación
		false,    // Persistencia
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Error al enviar el mensaje:", err)
		return err
	}

	log.Println("Mensaje enviado a la cola:", queue)
	return nil
}
