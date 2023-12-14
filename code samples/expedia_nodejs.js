import fetch from 'node-fetch';

const username = 'YOUR_USERNAME';
const password = 'YOUR_PASSWORD';
const body = {
 'source': 'universal',
 'url': 'https://www.expedia.de/Hotel-Search?adults=2&d1=2024-07-04&d2=2024-07-06&destination=Frankfurt%2C%20Deutschland%20%28FRA-Frankfurt%20Intl.%29&endDate=2024-07-06&flexibility=7_DAY&latLong=50.050978%2C8.571705&regionId=4280902&rooms=1&semdtl=&sort=RECOMMENDED&startDate=2024-07-04&theme=&useRewards=false&userIntent=',
 'geo_location': 'Germany',
};
const response = await fetch('https://realtime.oxylabs.io/v1/queries', {
  method: 'post',
  body: JSON.stringify(body),
  headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Basic ' + Buffer.from(`${username}:${password}`).toString('base64'),
  }
});

console.log(await response.json());
