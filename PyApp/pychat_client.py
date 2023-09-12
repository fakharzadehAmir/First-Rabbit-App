import subprocess
import threading
your_name = input("What's your name? ")

def run_sender():
    subprocess.call(["/usr/bin/python3", "py_sender.py"])


def run_receiver():
    subprocess.call(["/usr/bin/python3", "py_receiver.py"])

my_turn = True
while True:
    if my_turn:
        run_sender()
    else:
        run_receiver()
    my_turn = not my_turn