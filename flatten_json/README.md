# Flattening JSON
## Description
Flattening JSON takes a JSON event and flattens it so all nested structures are removed from it.  
## Supported event formats
JSON 
## Inputs:

N/A

## Test event 

### Inputs

N/A

### Event:

```json

{
    "dict": {
        "dict_innner_1":{
            "ice_cream": "is_good",
            "favorite_flavor": "vanilla",
            "favorite_topping": "caramel"
        },
        "dict_inner_2": {
            "beef_sticks": "Are yummy",
            "favorite_brand": "Klements Farm",
            "brand_location": "Wisconsin :)"
        }
    },
    "array": ["Food", "is", "yummy", ["strawberries", "grapes"], {
        "favorite_fruit": "black_grapes"
    }]
}

```

## Output to the test event

### Modified Event:
```json
{
    "array_0": "Food",
    "array_1": "is",
    "array_2": "yummy",
    "array_3_0": "strawberries",
    "array_3_1": "grapes",
    "beef_sticks": "Are yummy",
    "brand_location": "Wisconsin :)",
    "favorite_brand": "Klements Farm",
    "favorite_flavor": "vanilla",
    "favorite_fruit": "black_grapes",
    "favorite_topping": "caramel",
    "ice_cream": "is_good"
}
```