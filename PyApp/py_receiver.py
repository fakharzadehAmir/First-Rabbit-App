import pika
# Connect to RabbitMQ and create a channel, also declare the queue
credentials = pika.PlainCredentials(username="chat_app", password="chat_app_password")
connection = pika.BlockingConnection(pika.ConnectionParameters(host="localhost", credentials=credentials))
channel = connection.channel()
def callback(ch, method, properties, body):
    body = body.decode('utf-8')
    print(body)
    channel.stop_consuming()
    connection.close()
channel.basic_consume(queue='go_sender', on_message_callback=callback, auto_ack=True)
channel.start_consuming()