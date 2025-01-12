# locustfile.py
from locust import HttpUser, TaskSet, between, task
import random
import string


token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzY3MTQ3OTcsImlhdCI6MTczNjYyODM5N30.EzWA_7KPkbiKEZBVlDbc6BMFINNapFU3XDGc1ZsUUE4"
constant_url = "https://www.example.com"

def generate_random_url():
    return f"https://www.example.com/{''.join(random.choices(string.ascii_letters + string.digits, k=10))}"

class MyTask(TaskSet):
    # 写请求
    @task
    def create_short_url(self):
        headers = {
            "Authorization": "Bearer " + token,
            "Content-Type": "application/json"
        }
        data = {
            "original_url": constant_url,
            "duration": 24
        }
        self.client.post(url="/api/url", headers=headers, json=data)
        
    # 读请求
    # @task
    # def redirect_short_url(self):
    #     self.client.get(url="/api/url/Ln7YA")

class HelloWorldUser(HttpUser):
    tasks = [MyTask]
    wait_time = between(1, 2)
