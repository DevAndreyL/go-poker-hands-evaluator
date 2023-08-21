## How to install and test:
To run this application you need to use the console, install modules and build the binary. 
To do this you need to `cd` into project root and run `go get` command. Then you need to go to `cmd/poker` directory inside this project
and run the next command: `go build`. After this, you will get **poker** binary inside the same folder, and you can run this
simply typed `./poker`. This command will run HTTP server that will listen 80 port on your localhost(127.0.0.1).
This microservice got only one endpoint - `http://127.0.0.1/evaluate-hand`. This endpoint will evaluate any valid combination
of 5 cards with suits and works only with the POST method. 

Also, you can set up simple ReactJS application for this evaluator. More details provided [here](https://github.com/DevAndreyL/react-poker-hands-evaluator).

**IMPORTANT** There are no validation for cards input at this moment. This evaluator works only for 5 cards in hand, and it must
be in uppercase with suit. E.g. - `["7S", "8S", "9S", "TS", "JS"]`. The example and more info provided below.

You can use this JSON data sample to make a POST request to `/evaluate-hand` endpoint:
```
{
    "hands": {
        "first": ["7S", "8S", "9S", "TS", "JS"],
        "second": ["TS", "JS", "QS", "KS", "AS"]
    }
}
```

Here some explanation for input above:
- `hands` are main object that contain inside _handName_ and cards with suits for this hand. E.g. `"first": ["7S", "8S", "9S", "TS", "JS"]`. 
- You can use any count of hands, but there can't be more than 5 cards in each hand.
- All hands must contain valid card names and suits in uppercase.
- Valid card names are: `2, 3, 4, 5, 6, 7, 8, 9, T, J, Q, K, A`
- Valid suits are: `S, D, H, C`


And you will get next response for this input:

```
{
    "result": {
        "first": {
            "handName": "first",
            "combinationName": "Straight Flush",
            "handWeight": 45,
            "combinationWeight": 9
        },
        "second": {
            "handName": "second",
            "combinationName": "Royal Flush",
            "handWeight": 60,
            "combinationWeight": 10
        }
    }
}
```

We can see next fields:
- `handName` - this is the hand name that was provided in request
- `combinationName` - this is the name of combination for provided cards in hand
- `handWeight` - this is the weight for all hand cards. The hire the card, the hire the weight
- `combinationWeight` - this is the weight for combination. This value is constant for each combination, and the highest combination get the highest weight.


### Algorithmic complexity described in `hand.go` file for each function.