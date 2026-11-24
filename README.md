# Expedia Scraper

[![Oxylabs promo code](https://raw.githubusercontent.com/oxylabs/product-integrations/refs/heads/master/Affiliate-Universal-1090x275.png)](https://oxylabs.io/pages/gitoxy?utm_source=877&utm_medium=affiliate&groupid=877&utm_content=expedia-scraper-github&transaction_id=102f49063ab94276ae8f116d224b67)

<p align="left">
  <a href="https://dashboard.oxylabs.io"><img src="https://img.shields.io/badge/%F0%9F%9A%80_Get_Started-FF6B35?style=for-the-badge" alt="Get started"></a>
  <a href="https://developers.oxylabs.io/products/web-api"><img src="https://img.shields.io/badge/%F0%9F%93%9A_Documentation-4A90E2?style=for-the-badge" alt="Documentation"></a>
  <a href="https://dashboard.oxylabs.io"><img src="https://img.shields.io/badge/%F0%9F%93%8A_Dashboard-2EA44F?style=for-the-badge" alt="Dashboard"></a>
  <a href="https://discord.gg/Pds3gBmKMH"><img src="https://img.shields.io/badge/%F0%9F%92%AC_Discord-5865F2?style=for-the-badge" alt="Discord"></a>
  <a href="https://www.youtube.com/@oxylabs"><img src="https://img.shields.io/badge/%E2%96%B6%EF%B8%8F_YouTube-FF0000?style=for-the-badge" alt="YouTube"></a>
</p>

[Expedia Scraper](https://oxylabs.io/products/scraper-api/web/expedia) lets you gather flight, destination, price, and other data from Expedia’s web pages interruption-free and at a large scale. 

Follow this guide to scrape Expedia using our [Scraper API](https://oxylabs.io/products/scraper-api). 

### How it works

You can get Expedia results by providing URLs to our service. The API returns an HTML of any Expedia page you request.

#### Python code example

The following example shows how to get the result of an Expedia page in HTML format.

```python
import requests
from pprint import pprint

# Structure payload.
payload = {
   'source': 'universal',
   'url': 'https://www.expedia.de/Hotel-Search?adults=2&d1=2025-07-04&d2=2025-07-06&destination=Frankfurt%2C%20Deutschland%20%28FRA-Frankfurt%20Intl.%29&endDate=2025-07-06&flexibility=7_DAY&latLong=50.050978%2C8.571705&regionId=4280902&rooms=1&semdtl=&sort=RECOMMENDED&startDate=2025-07-04&theme=&useRewards=false&userIntent=',
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
```

Code samples for other programming languages are [here](https://github.com/oxylabs/expedia-scraper/tree/main/code%20samples), and Oxylabs documentation can be found [here](https://developers.oxylabs.io/products/web-scraper-api).

#### Output example

```json
{
  "results": [
    {
      "content": "<!doctype html>\n<html lang=\"en\">\n<head>...</script></body>\n</html>\n",
      "created_at": "2023-06-28 07:52:07",
      "updated_at": "2023-06-28 07:52:11",
      "page": 1,
      "url": "https://www.expedia.de/Hotel-Search?adults=2&d1=2025-07-04&d2=2025-07-06&destination=Frankfurt%2C%20Deutschland%20%28FRA-Frankfurt%20Intl.%29&endDate=2025-07-06&flexibility=7_DAY&latLong=50.050978%2C8.571705&regionId=4280902&rooms=1&semdtl=&sort=RECOMMENDED&startDate=2025-07-04&theme=&useRewards=false&userIntent=",
      "job_id": "7079728156646597633",
      "status_code": 200
    }
  ]
}
```

With Oxylabs' Expedia Scraper, you can entrust us with infrastructure management and web unblocking, allowing you to concentrate on the pivotal aspects – data and its analysis.

If you have questions, please contact us at support@oxylabs.io or via live chat on our website.
