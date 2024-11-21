const locationData = [
  {
    "country": "Argentina",
    "cities": [
      "buenos_aires",
      "la_plata",
      "san_isidro"
    ]
  },
  {
    "country": "Australia",
    "cities": [
      "brisbane",
      "burswood",
      "melbourne",
      "new_south_wales",
      "queensland",
      "sydney",
      "victoria",
      "west_melbourne"
    ]
  },
  {
    "country": "Austria",
    "cities": [
      "klagenfurt",
      "nickelsdorf",
      "vienna"
    ]
  },
  {
    "country": "Belarus",
    "cities": [
      "minsk"
    ]
  },
  {
    "country": "Belgium",
    "cities": [
      "antwerp",
      "rotselaar",
      "werchter"
    ]
  },
  {
    "country": "Brazil",
    "cities": [
      "belo_horizonte",
      "brasilia",
      "porto_alegre",
      "recife",
      "rio_de_janeiro",
      "sao_paulo"
    ]
  },
  {
    "country": "Canada",
    "cities": [
      "montreal",
      "quebec",
      "toronto",
      "vancouver",
      "windsor"
    ]
  },
  {
    "country": "Chile",
    "cities": [
      "santiago"
    ]
  },
  {
    "country": "China",
    "cities": [
      "changzhou",
      "hong_kong",
      "huizhou",
      "sanya"
    ]
  },
  {
    "country": "Colombia",
    "cities": [
      "bogota"
    ]
  },
  {
    "country": "Costa Rica",
    "cities": [
      "san_jose"
    ]
  },
  {
    "country": "Czechia",
    "cities": [
      "ostrava",
      "prague"
    ]
  },
  {
    "country": "Denmark",
    "cities": [
      "aalborg",
      "aarhus",
      "copenhagen",
      "skanderborg"
    ]
  },
  {
    "country": "Finland",
    "cities": [
      "oulu",
      "turku"
    ]
  },
  {
    "country": "France",
    "cities": [
      "arras",
      "boulogne_billancourt",
      "freyming_merlebach",
      "lyon",
      "nimes",
      "pagney_derriere_barine",
      "paris",
      "sochaux"
    ]
  },
  {
    "country": "French Polynesia",
    "cities": [
      "papeete"
    ]
  },
  {
    "country": "Germany",
    "cities": [
      "berlin",
      "cuxhaven",
      "dusseldorf",
      "frankfurt",
      "hamburg",
      "leipzig",
      "mannheim",
      "merkers",
      "monchengladbach",
      "munich",
      "salem",
      "scheessel"
    ]
  },
  {
    "country": "Greece",
    "cities": [
      "athens"
    ]
  },
  {
    "country": "Hungary",
    "cities": [
      "budapest"
    ]
  },
  {
    "country": "India",
    "cities": [
      "mumbai"
    ]
  },
  {
    "country": "Indonesia",
    "cities": [
      "jakarta",
      "yogyakarta"
    ]
  },
  {
    "country": "Ireland",
    "cities": [
      "dublin"
    ]
  },
  {
    "country": "Italy",
    "cities": [
      "florence",
      "imola",
      "milan"
    ]
  },
  {
    "country": "Japan",
    "cities": [
      "nagoya",
      "osaka",
      "saitama"
    ]
  },
  {
    "country": "Mexico",
    "cities": [
      "mexico_city",
      "monterrey",
      "playa_del_carmen"
    ]
  },
  {
    "country": "Netherlands",
    "cities": [
      "amsterdam",
      "eindhoven",
      "groningen",
      "landgraaf"
    ]
  },
  {
    "country": "Netherlands Antilles",
    "cities": [
      "willemstad"
    ]
  },
  {
    "country": "New Caledonia",
    "cities": [
      "noumea"
    ]
  },
  {
    "country": "New Zealand",
    "cities": [
      "auckland",
      "dunedin",
      "penrose",
      "wellington"
    ]
  },
  {
    "country": "Norway",
    "cities": [
      "oslo"
    ]
  },
  {
    "country": "Peru",
    "cities": [
      "lima"
    ]
  },
  {
    "country": "Philippines",
    "cities": [
      "manila"
    ]
  },
  {
    "country": "Poland",
    "cities": [
      "gdynia",
      "krakow",
      "warsaw"
    ]
  },
  {
    "country": "Portugal",
    "cities": [
      "lisbon"
    ]
  },
  {
    "country": "Qatar",
    "cities": [
      "doha"
    ]
  },
  {
    "country": "Romania",
    "cities": [
      "napoca"
    ]
  },
  {
    "country": "Saudi Arabia",
    "cities": [
      "riyadh"
    ]
  },
  {
    "country": "Slovakia",
    "cities": [
      "bratislava"
    ]
  },
  {
    "country": "South Korea",
    "cities": [
      "seoul"
    ]
  },
  {
    "country": "Spain",
    "cities": [
      "barcelona",
      "bilbao",
      "burriana",
      "madrid",
      "seville",
      "zaragoza"
    ]
  },
  {
    "country": "Sweden",
    "cities": [
      "gothenburg",
      "stockholm"
    ]
  },
  {
    "country": "Switzerland",
    "cities": [
      "frauenfeld",
      "lausanne",
      "sion",
      "st_gallen",
      "zurich"
    ]
  },
  {
    "country": "Taiwan",
    "cities": [
      "taipei"
    ]
  },
  {
    "country": "Thailand",
    "cities": [
      "bangkok"
    ]
  },
  {
    "country": "United Arab Emirates",
    "cities": [
      "abu_dhabi",
      "dubai"
    ]
  },
  {
    "country": "United Kingdom",
    "cities": [
      "aberdeen",
      "birmingham",
      "brixton",
      "cardiff",
      "glasgow",
      "london",
      "manchester",
      "westcliff_on_sea"
    ]
  },
  {
    "country": "USA",
    "cities": [
      "alabama",
      "amityville",
      "anaheim",
      "arizona",
      "berwyn",
      "boston",
      "brooklyn",
      "california",
      "canton",
      "charlotte",
      "chicago",
      "cincinnati",
      "cleveland",
      "colorado",
      "dallas",
      "del_mar",
      "detroit",
      "florida",
      "georgia",
      "grand_rapids",
      "hershey",
      "houston",
      "illinois",
      "indianapolis",
      "inglewood",
      "kansas_city",
      "las_vegas",
      "los_angeles",
      "madison",
      "maine",
      "massachusetts",
      "michigan",
      "minneapolis",
      "minnesota",
      "missouri",
      "nevada",
      "new_hampshire",
      "new_orleans",
      "new_york",
      "newark",
      "oakland",
      "oklahoma",
      "omaha",
      "oregon",
      "pennsylvania",
      "philadelphia",
      "pico_rivera",
      "pittsburgh",
      "rosemont",
      "san_francisco",
      "seattle",
      "st_louis",
      "texas",
      "uniondale",
      "utah",
      "washington"
    ]
  }
];