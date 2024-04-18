# prog2005-assignment1
A rest API written in Go that finds basic info about universities, and the countries they are based in.
This service uses the [hipolabs university domain API](https://github.com/Hipo/university-domains-list/) and [Rest countries API](https://gitlab.com/restcountries/restcountries).

## Endpoints
The api has three main endpoints:
* `/unisearcher/v1/uniinfo/{name_component}/`
* `/unisearcher/v1/neighbourunis/{country}/{name_component}[?limit=n]`
* `/unisearcher/v1/diag/`

### uniinfo endpoint
Returns a list of all universities with a given name component.

* Example: https://prog2005-assignment1-qgfb.onrender.com/unisearcher/v1/uniinfo/norwegian%20university%20of
* Body:
```json
[
    {
        "name": "Norwegian University of Sport and Physical Education",
        "country": "Norway",
        "isocode": "NO",
        "webpages": [
            "http://www.nih.no/"
        ],
        "languages": {
            "nno": "Norwegian Nynorsk",
            "nob": "Norwegian Bokmål",
            "smi": "Sami"
        },
        "map": "https://www.openstreetmap.org/relation/2978650"
    },
    ...
]
```

### neighbourunis endpoint
Returns a list of universities with a given {name_component} in neighbouring countries to {country}.
The maximum number of items to retrieve can be specified with the limit query.

* Example: https://prog2005-assignment1-qgfb.onrender.com/unisearcher/v1/neighbourunis/norway/tech
* Body:
```json
[
    {
        "name": "Tampere University of Technology",
        "country": "Finland",
        "isocode": "FI",
        "webpages": [
            "http://www.bth.se/"
        ],
        "languages": {
            "fin": "Finnish",
            "swe": "Swedish"
        },
        "map": "openstreetmap.org/relation/54224"
    },
    {
        "name": "Helsinki University of Technology",
        "country": "Finland",
        "isocode": "FI",
        "webpages": [
            "http://www.chalmers.se/"
        ],
        "languages": {
            "fin": "Finnish",
            "swe": "Swedish"
        },
        "map": "openstreetmap.org/relation/54224"
    },
    ...
]
```

### diag endpoint
Provides the status of this service and the remote services it depends on.

* Example: https://prog2005-assignment1-qgfb.onrender.com/unisearcher/v1/diag/
* Body:
```json
{
    "universitiesapi": "200 OK",
    "countriesapi": "200 OK",
    "version": "v1",
    "uptime": "287s"
}
```
