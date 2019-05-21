import sys, re
from bs4 import BeautifulSoup


table = open(sys.argv[1])

result = []

for i in range(11):
    table.readline()

for line in table:
    number = line
    number = number[5:8] 
    print("number: " + str(number))
    
    name = table.readline()
    name = re.split("[<>]", name)
    print("name: " + str(name[-3]))
    
    series = table.readline()
    series = re.split("[<>]", series)
    print("series: " + str(series[-3]))

    price = table.readline()
    price = price.rstrip().split(" ")
    print("price: " + str(price[-1]))

    release = table.readline()
    release = release.rstrip().split(" ")
    print("release: " + " ".join(release[-2:]))
    
    description = table.readline()
    soup = BeautifulSoup(description.rstrip())
    print("description: " + soup.get_text())

    table.readline()
    table.readline()

    print()
    
table.close()
