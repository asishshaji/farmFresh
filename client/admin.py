from random_username.generate import generate_username
import requests
import pprint
import os
import unittest
import time

ADMIN_BASE_URL = "http://0.0.0.0:9090/admin/"
ADMIN_LOGIN = ADMIN_BASE_URL + "login"
ADD_PRODUCT = ADMIN_BASE_URL + "product"
CHANGE_FARMER_STATE = ADMIN_BASE_URL + "farmer/state"
# GET_ADMINS = ADMIN_BASE_URL + "admins"  # method get


FARMER_BASE_URL = "http://0.0.0.0:9090/farmer/"
FARMER_SIGNUP = FARMER_BASE_URL + "signup"
FARMER_LOGIN = FARMER_BASE_URL + "login"


class Farmer:
    def __init__(self, firstname, lastname, age, dp, password):
        self.firstname = firstname
        self.lastname = lastname
        self.age = age
        self.dp = dp
        self.password = password

        self.username = ""

        self.token = ""

    def signUpFarmer(self):
        data = {
            "firstname": self.firstname,
            "lastname": self.lastname,
            "age": self.age,
            "password": self.password,
        }

        img = open(self.dp, "rb")

        files = {
            "profile_pic": img,
        }

        res = requests.post(FARMER_SIGNUP, data=data, files=files)
        self.username = res.json()['username']
        pprint.pprint(res.json())

    def loginFarmer(self):
        data = {
            "username": self.username,
            "password": self.password,
        }
        res = requests.post(FARMER_LOGIN, data=data)
        if res.status_code == 200:
            self.token = res.json()['token']
            return
        pprint.pprint(res.json())


class Admin:
    def __init__(self, username, password):
        self.username = username
        self.password = password
        self.token = ""

    def changeFarmerState(self, farmer_state, id):
        res = requests.post(CHANGE_FARMER_STATE, data={
            "farmer_id": id,
            "state": farmer_state
        },
            headers={
            'Authorization': "Bearer "+self.token}
        )
        print(res.json())

    def loginAdmin(self):
        data = {
            "username": self.username,
            "password": self.password
        }
        res = requests.post(ADMIN_LOGIN, data=data)
        self.token = res.json()['token']

    def addProduct(self, product_name, owner_id, price, market_price, about, category):
        data = {
            "name": product_name,
            "owner_id": owner_id,
            "price": price,
            "market_price": market_price,
            "about": about,
            "category": category,
        }
        res = requests.post(ADD_PRODUCT, data=data, headers={
                            'Authorization': "Bearer "+self.token})
        pprint.pprint(res.json())


def addProducts(admin, farmer):
    while 1:
        name = input("Enter product name :> ")
        cat = input("Enter category :> ")
        admin.addProduct(name,
                         farmer.username,
                         "34",
                         "77",
                         "Mollit sint cupidatat pariatur laboris nostrud reprehenderit ad laborum cillum adipisicing ullamco.",
                         cat)


random_first_name = generate_username()[0]
print("Generating random farmer with first name ", random_first_name)
farmer = Farmer(random_first_name, "AD", "23", "dp.jpg", "kingkai1253")
farmer.signUpFarmer()
farmer.loginFarmer()

admin = Admin("ASISH2323 ", "kingkai1253")
admin.loginAdmin()
admin.changeFarmerState("suspend", farmer.username)
addProducts(admin, farmer)
