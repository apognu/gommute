# gommute

Get simple public transportation itineraries for [Navitia](http://www.navitia.io)-supported cities, through their API.

This was crafted for a specific need (very simple itineraries) and will not be updated to include additional features unless needed or requested.

## Usage

```go
from := gommute.Coordinates{Longitude: "2.3616223", Latitude: "48.8675065"}
to := gommute.Coordinates{Longitude: "2.2850278", Latitude: "48.8302232"}

gm := gommute.New("api_key", from, to)
gm.Time = gommute.NewTimeReference(time.Now().Add(3 * time.Hour), gommute.Departure)

journey := gm.Journey()
```

Now, ```journey``` can be JSON-marshalled to this:

```json
{
  "duration": 2413,
  "sections": [
    {
      "mode": "walking",
      "from": "14 Rue Meslay",
      "from_coordinates": {
        "lon": "2.3616223",
        "lat": "48.8675065"
      },
      "to": "République",
      "to_coordinates": {
        "lon": "2.363348",
        "lat": "48.867674"
      },
      "departure_time": "2016-03-03T13:38:41Z",
      "arrival_time": "2016-03-03T13:41:00Z",
      "duration": 139
    },
    {
      "mode": "metro",
      "name": "8",
      "color": "#C5A3CA",
      "direction": "Balard",
      "from": "République",
      "from_coordinates": {
        "lon": "2.363348",
        "lat": "48.867674"
      },
      "to": "Balard",
      "to_coordinates": {
        "lon": "2.278701",
        "lat": "48.836818"
      },
      "departure_time": "2016-03-03T13:41:00Z",
      "arrival_time": "2016-03-03T14:03:00Z",
      "duration": 1320
    },
    {
      "mode": "transfer",
      "departure_time": "2016-03-03T14:03:00Z",
      "arrival_time": "2016-03-03T14:05:48Z",
      "duration": 168
    },
    {
      "mode": "tramway",
      "name": "T3A",
      "color": "#DE8B53",
      "direction": "Porte De Vincennes",
      "from": "Balard",
      "from_coordinates": {
        "lon": "2.278948",
        "lat": "48.835722"
      },
      "to": "Porte De Versailles - Parc Des Expositions",
      "to_coordinates": {
        "lon": "2.288045",
        "lat": "48.832616"
      },
      "departure_time": "2016-03-03T14:09:00Z",
      "arrival_time": "2016-03-03T14:12:00Z",
      "duration": 180
    },
    {
      "mode": "walking",
      "from": "Porte De Versailles - Parc Des Expositions",
      "from_coordinates": {
        "lon": "2.288045",
        "lat": "48.832616"
      },
      "to": "Rue Du Quatre Septembre",
      "to_coordinates": {
        "lon": "2.2850278",
        "lat": "48.8302232"
      },
      "departure_time": "2016-03-03T14:12:00Z",
      "arrival_time": "2016-03-03T14:18:54Z",
      "duration": 414
    }
  ]
}
```
