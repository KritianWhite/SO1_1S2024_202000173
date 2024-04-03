import time, redis

class RedisSub:
    def __init__(self, host, port, db=0):
        self.r = redis.Redis(host=host, port=port, db=db, decode_responses=True)
        self.pubsub = self.r.pubsub()
        
    def subscribe(self, channel, callback):
        self.pubsub.subscribe(**channel)
        print('subscribed to', channel.keys())
        while True:
            message = self.pubsub.get_message()
            if message and not message['type'] == 'subscribe':
                callback(message)
            time.sleep(0.02)
            
def handle_message(message):
    print("received message:", {message['data']})
        
if __name__ == '__main__':
    host = "10.22.69.155"
    port = 6379
    channel = {"test": handle_message}
    
    subscriber = RedisSub(host, port)
    subscriber.subscribe(channel, handle_message)