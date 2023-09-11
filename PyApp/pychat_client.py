import pika

# Connect to RabbitMQ and create a channel, also declare the queue
credentials = pika.PlainCredentials(username="chat_app", password="chat_app_password")
connection = pika.BlockingConnection(pika.ConnectionParameters(host="localhost", credentials=credentials))
channel = connection.channel()
channel.queue_declare(queue="py_sender")
# channel.queue_declare(queue="go_sender")
my_turn = True
your_name = input("What's your name? ")
while True:
    try:
        # if my_turn:
            my_message = input("you: ")
            channel.basic_publish(exchange='', routing_key="py_sender", body=f'{your_name}: {my_message}')
        # else:
        #     def callback(ch, method, properties, body):
        #         body = body.decode('utf-8')
        #         print(body)
        #     channel.start_consuming()
        #     channel.basic_consume(queue='go_sender', on_message_callback=callback, auto_ack=True)
        #     channel.stop_consuming()
        # my_turn = not my_turn
    except KeyboardInterrupt:
        connection.close()