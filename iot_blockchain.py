import hashlib
import datetime
import random


class Block:
    def __init__(self, index, sensor_data, previous_hash):
        self.index = index
        self.sensor_data = sensor_data
        self.timestamp = str(datetime.datetime.now())
        self.previous_hash = previous_hash
        self.hash = self.generate_hash()

    def generate_hash(self):
        block_content = (
            str(self.index)
            + self.sensor_data
            + self.timestamp
            + self.previous_hash
        )
        return hashlib.sha256(block_content.encode()).hexdigest()


class Blockchain:
    def __init__(self):
        self.chain = []
        self.create_genesis_block()

    def create_genesis_block(self):
        genesis_block = Block(0, "Genesis Block", "0")
        self.chain.append(genesis_block)

    def add_block(self, sensor_data):
        previous_block = self.chain[-1]
        new_block = Block(len(self.chain), sensor_data, previous_block.hash)
        self.chain.append(new_block)

    def display_chain(self):
        for block in self.chain:
            print("\n--------------------------------")
            print("Block Number:", block.index)
            print("IoT Sensor Data:", block.sensor_data)
            print("Timestamp:", block.timestamp)
            print("Previous Hash:", block.previous_hash)
            print("Current Hash:", block.hash)

    def verify_chain(self):
        for i in range(1, len(self.chain)):
            current = self.chain[i]
            previous = self.chain[i - 1]
            if current.previous_hash != previous.hash:
                return False
            if current.hash != current.generate_hash():
                return False
        return True


def generate_sensor_reading():
    temperature = random.randint(20, 40)
    humidity = random.randint(30, 90)
    return f"Temperature={temperature}C, Humidity={humidity}%"


iot_chain = Blockchain()

for _ in range(3):
    iot_chain.add_block(generate_sensor_reading())

iot_chain.display_chain()

if iot_chain.verify_chain():
    print("\nIoT Blockchain Verified Successfully")
else:
    print("\nWarning! IoT Data Tampered")

# Tampering Test
iot_chain.chain[1].sensor_data = "Temperature=100C"

if iot_chain.verify_chain():
    print("\nIoT Blockchain Verified Successfully")
else:
    print("\nWarning! IoT Data Tampered")
