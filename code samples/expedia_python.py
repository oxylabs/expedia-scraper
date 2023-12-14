import requests
from pprint import pprint

# Structure payload.
payload = {
   'source': 'universal',
   'url': 'https://www.expedia.de/Hotel-Search?adults=2&d1=2024-07-04&d2=2024-07-06&destination=Frankfurt%2C%20Deutschland%20%28FRA-Frankfurt%20Intl.%29&endDate=2024-07-06&flexibility=7_DAY&latLong=50.050978%2C8.571705&regionId=4280902&rooms=1&semdtl=&sort=RECOMMENDED&startDate=2024-07-04&theme=&useRewards=false&userIntent=',
   'geo_location': 'Germany'
}

# Get response.
response = requests.request(
    'POST',
    'https://realtime.oxylabs.io/v1/queries',
    auth=('YOUR_USERNAME', 'YOUR_PASSWORD'), #Your credentials go here
    json=payload,
)

# Instead of response with job status and results url, this will return the
# JSON response with results.
pprint(response.json())
