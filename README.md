# Checksp

This is a test project to give Go a try. Concepts to explore: http calls, handling JSON responeses and csv files

The theme is around a series of housekeeping task for the Species Home Pages. Is cli tool, there are three subcommans: 
- `json` shows constructs a test json object. 
- `articles` checks on the cms database which articles have maps and if the maps have corresponding data in the API (species can change names).
- `list filename.csv` that does the later check on a `csv` file.  

### `list filname.csv`
I read species from a `csv` file, chcek if we have data abot it on the API and if not save those records on `csv` file 

### `json`

In order to understand unmarshalling and constructing the structs, it was easier to reverse engenieer the structs and marshall them first. That is [testJSON.go](testJson.go), with the struct correctly shape it was a breeze to handle the JSON response.

The JSON I'm trying to prcess looks something like: 

```JSON
[
  {
    "_id": "Zaretis crawfordhilli",
    "count": 2,
    "data": {
      "Quebrada Escondida": [
        {
          "_id": "5fcfb3133f2e48e107ac546f",
          "herbivoreSpecies": "Zaretis crawfordhilli",
          "collectionDate": "11/12/2001",
          "herbivoreFamily": "Nymphalidae",
          "latitude": "10.89928",
          "locality": "Quebrada Escondida",
          "longitude": "-85.27486",
          "voucher": "01-SRNP-23274"
        }
      ],
      "Mabea": [
        {
          "_id": "5fcfb3f83f2e48e107b4dc6c",
          "herbivoreSpecies": "Zaretis crawfordhilli",
          "collectionDate": "04/05/2018",
          "herbivoreFamily": "Nymphalidae",
          "latitude": "10.96084",
          "locality": "Mabea",
          "longitude": "-85.32205",
          "voucher": "18-SRNP-26343"
        }
      ],
    }
  }
]
```


### `articles`

This option requires a `.env` file with the following structutre: 

```
USER=db_user
PASSWORD=db_passwords
DATABASE=databae
HOST=host
```