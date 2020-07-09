from math import sin, cos, sqrt, atan2, radians

# approximate radius of earth in km
R = 6373.0

lat1 = radians(53.339428)
lon1 = radians(-6.257664)
lat2 = radians(52.986375)
lon2 = radians(-6.043701)

print(lat1)

dlon = lon2 - lon1
dlat = lat2 - lat1

a = sin(dlat / 2)**2 + cos(lat1) * cos(lat2) * sin(dlon / 2)**2
c = 2 * atan2(sqrt(a), sqrt(1 - a))


print(c)
distance = R * c

print("Result:", distance)
print("Should be:", 278.546, "km")