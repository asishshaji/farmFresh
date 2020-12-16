import requests
import pprint
import os
import unittest

BASE_URL = "http://0.0.0.0:9090/super/"
CREATE_ADMIN = BASE_URL + "create"  # post method
GET_ADMINS = BASE_URL + "admins"  # method get

SUPERADMIN_PASS = "adminpass"


class Admin:
    def get_admins(self):
        res = requests.get(GET_ADMINS, params={
            "sup_password": SUPERADMIN_PASS
        })
        # pprint.pprint(res.json())
        return res.json()

    def create_admin(self, username, password, filename):
        data = {
            "sup_password": SUPERADMIN_PASS,
            "username": username,
            "password": password,
        }
        dp = open(filename, "rb")
        files = {
            "profile_pic": dp
        }

        res = requests.post(CREATE_ADMIN, data=data, files=files)
        # pprint.pprint(res.json())
        dp.close()

        return res.json()


class TestingGoServer(unittest.TestCase):

    def test_createAdmin(self):
        a = Admin()
        message = a.create_admin("ASISH2323 ", "kingkai1253", "dp.jpg")
        self.assertEqual(message['message'], "Success")

    def test_getAdmins(self):
        a = Admin()
        res = a.get_admins()

        self.assertTrue(len(res) > 0,
                        str(len(res)) + " admins exists")


if __name__ == '__main__':
    unittest.main()
