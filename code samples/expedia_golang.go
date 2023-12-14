package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    const Username = "YOUR_USERNAME"
    const Password = "YOUR_PASSWORD"

    payload := map[string]string{
      	"source": "universal",
	"url": "https://www.expedia.de/Hotel-Search?adults=2&d1=2024-07-04&d2=2024-07-06&destination=Frankfurt%2C%20Deutschland%20%28FRA-Frankfurt%20Intl.%29&endDate=2024-07-06&flexibility=7_DAY&latLong=50.050978%2C8.571705&regionId=4280902&rooms=1&semdtl=&sort=RECOMMENDED&startDate=2024-07-04&theme=&useRewards=false&userIntent=",
	"geo_location": "Germany",
          }

    jsonValue, _ := json.Marshal(payload)

    client := &http.Client{}
    request, _ := http.NewRequest("POST",
        "https://realtime.oxylabs.io/v1/queries",
        bytes.NewBuffer(jsonValue),
    )

    request.SetBasicAuth(Username, Password)
    response, _ := client.Do(request)

    responseText, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(responseText))
}
