import pika
# Connect to RabbitMQ and create a channel, also declare the queue
credentials = pika.PlainCredentials(username="chat_app", password="chat_app_password")
connection = pika.BlockingConnection(pika.ConnectionParameters(host="localhost", credentials=credentials))
channel = connection.channel()
channel.queue_declare(queue="py_sender")
my_message = input("you: ")
channel.basic_publish(exchange='', routing_key="py_sender", body=f'Amir: {my_message}')

connection.close()