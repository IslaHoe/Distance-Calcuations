import json
from math import sin, cos, sqrt, atan2, radians
# Using readline() 
file1 = open('myfile.txt', 'r') 
count = 0

while True: 
    count += 1
  
    # Get next line from file 
    line = file1.readline() 
  
    # if line is empty 
    # end of file is reached 
    if not line: 
        break
    y = json.loads(line)


    x = float(y["latitude"])
    z = x*10
    print(z)
    # the result is a Python dictionary:
    print(y["latitude"])
    
file1.close() 

# approximate radius of earth in km
R = 6373.0

lat1 = radians(52.2296756)
lon1 = radians(21.0122287)
lat2 = radians(52.406374)
lon2 = radians(16.9251681)

dlon = lon2 - lon1
dlat = lat2 - lat1

a = sin(dlat / 2)**2 + cos(lat1) * cos(lat2) * sin(dlon / 2)**2
c = 2 * atan2(sqrt(a), sqrt(1 - a))

distance = R * c

print("Result:", distance)
print("Should be:", 278.546, "km")