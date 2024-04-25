import json
from random import randrange
import requests
from locust import HttpUser, between, task

class readFile():
    def __init__(self):
        self.data = []

    def getData(self): 
        size = len(self.data) 
        if size > 0:
            index = randrange(0, size - 1) if size > 1 else 0
            return self.data.pop(index)
        else:
            print("size -> 0")
            return None
    
    def loadFile(self):
        print("Loading ...")
        try:
            response = requests.get("https://my.api.mockaroo.com/Songs.json?key=bcdcf1a0")
            if response.status_code == 200:
                self.data = response.json()
                # Modifica el campo "rank" para cada registro
                for item in self.data:  
                    item["year"] = str(item["year"])
                    item["rank"] = str(item["rank"])  
                
                print("Data loaded:")
                print(json.dumps(self.data, indent=4))
            else:
                print("Error loading data from Mockaroo. Status:", response.status_code)
        except Exception as e:
            print(f'Error: {e}')

class trafficData(HttpUser):
    wait_time = between(0.1, 0.9)
    reader = readFile()
    reader.loadFile()

    def on_start(self):
        print("On Start")
    
    @task
    def sendMessage(self):
        data = self.reader.getData() 
        if data is not None:
            res = self.client.post("/insert", json=data)
            response = res.json()
            print(response)
        else:
            print("Empty") 
            self.stop(True)

    @task
    def getMessage(self):
        self.client.get("/receive")
