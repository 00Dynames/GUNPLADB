from flask import Flask
from flaskext.mysql import MySQL

app = Flask(__name__)
mysql = MySQL()

# Config
app.config["MYSQL_DATABASE_USER"] = "root"
app.config["MYSQL_DATABASE_PASSWORD"] = "root"
app.config["MYSQL_DATABASE_DB"] = "gunpladb"
app.config["MYSQL_DATABASE_HOST"] = "localhost"
mysql.init_app(app)

from src import routes
