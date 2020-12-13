# Checksps

This is a test project to give Go a try. Concepts to explore: http calls, handling JSON responeses and csv files

I read species from a csv file, chcek if we have data abot it on the API and if not save those records on csv file 


The JSON I'm trying to prcess looks something like: 

```
[
  {
    "_id": "Zaretis crawfordhilli",
    "count": 180,
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