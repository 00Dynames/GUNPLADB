import sys, re, json
from bs4 import BeautifulSoup


table = open(sys.argv[1])

result = []

for i in range(13):
    table.readline()

for line in table:
    number = line
    number = number[5:7] 
    #print("number: " + str(number))
    
    name = table.readline()
    name = re.split("[<>]", name)
    #print("name: " + str(name[-3]))
    
    series = table.readline()
    series = re.split("[<>]", series)
    #print("series: " + str(series[-3]))

    price = table.readline()
    price = price.rstrip().split(" ")
    #print("price: " + str(price[-1]))

    release = table.readline()
    release = release.rstrip().split(" ")
    #print("release: " + " ".join(release[-2:]))
    
    description = table.readline()
    soup = BeautifulSoup(description.rstrip(), features="html.parser")
    #print("description: " + soup.get_text())

    ms_joints = table.readline()
    ms_joints = ms_joints.rstrip().split("</td><td>")
    #print(ms_joints)

    included_hands = table.readline()
    included_hands = included_hands.rstrip().split("</td><td>")
    #print(included_hands)

    table.readline()
    table.readline()

    result.append({
        'name': str(name[-3]),
        'id': int(number) if re.match('[0-9]+', number) else 'N/A',
        'series': str(series[-3]),
        'price': str(price[-1]),
        'release': " ".join(release[-2:]),
        'description': soup.get_text(),
        'ms_joints': ms_joints[1] if len(ms_joints) > 1 else 'N/A',
        'included_hands': included_hands[1] if len(included_hands) > 1 else 'N/A',
        'grade': "RG"
    })

    #print()
    
table.close()

print(json.dumps(result))
