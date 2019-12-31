# Authentication contract

common event format is as follows, including txhash, state, gasConsumed and notify, each native contract method have different notifies.

|key|description|
|:--|:--|
|TxHash|transaction hash|
|State|1 indicates success，0 indicates fail|
|GasConsumed|gas fee consumed by this transaction|
|Notify|Notify event|

#### InitContractAdmin

* Usage: Init admin information of a certain contract through authentication contract

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of the method
    {
      "ContractAddress": "0600000000000000000000000000000000000000", //contract address of authentication contract
      "States":[
        "initContractAdmin", //method name
        "ea1e2adf8c19f5a7e877860264ebf326e8c3aa5a", //contract address of contract which want to achieve authentication control
        "did:tst:AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA" //admin tstid if above contract
      ]
    },
    //notify of gas fee transfer
    {
      "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
      "States":[
        "transfer", //method name
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker's address (from)
        "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //governance contract address (to)
        10000000 //gas fee amount(decimal: 9)
      ]
    }
  ]
}
```

#### Transfer

* Usage: Transfer admin to another tstid

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of the method
    {
      "ContractAddress": "0600000000000000000000000000000000000000", //contract address of authentication contract
      "States":[
        "transfer", //method name
        "ea1e2adf8c19f5a7e877860264ebf326e8c3aa5a", //contract address of contract which want to achieve authentication control
        true //status
      ]
    },
    //notify of gas fee transfer
    {
      "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
      "States":[
        "transfer", //method name
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker's address (from)
        "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //governance contract address (to)
        10000000 //gas fee amount(decimal: 9)
      ]
    }
  ]
}
```


#### AssignFuncsToRole

* Usage: Assign authentication of invoking a function in a certain contract to a role

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of the method
    {
      "ContractAddress": "0600000000000000000000000000000000000000", //contract address of authentication contract
      "States":[
        "assignFuncsToRole", //method name
        "ea1e2adf8c19f5a7e877860264ebf326e8c3aa5a", //contract address of contract which want to achieve authentication control
        true //status
      ]
    },
     //notify of gas fee transfer
     {
       "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
       "States":[
         "transfer", //method name
         "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker's address (from)
         "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //governance contract address (to)
         10000000 //gas fee amount(decimal: 9)
       ]
     }
  ]
}
```

#### AssignTstIDsToRole

* Usage: Assign a role to a certain tstid

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of the method
    {
      "ContractAddress": "0600000000000000000000000000000000000000", //contract address of authentication contract
      "States":[
        "assignTstIDsToRole", //method name
        "ea1e2adf8c19f5a7e877860264ebf326e8c3aa5a", //contract address of contract which want to achieve authentication control
        true //status
      ]
    },
     //notify of gas fee transfer
     {
       "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
       "States":[
         "transfer", //method name
         "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker's address (from)
         "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //governance contract address (to)
         10000000 //gas fee amount(decimal: 9)
       ]
     }
  ]
}
```

#### Delegate

* Usage: delegate authentication to another tstid

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of the method
    {
      "ContractAddress": "0600000000000000000000000000000000000000", //contract address of authentication contract
      "States":[
        "delegate",// method name
        "ea1e2adf8c19f5a7e877860264ebf326e8c3aa5a", //contract address of contract which want to achieve authentication control
        "did:tst:AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //from tstid
        "did:tst:AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //to tstid
        true //status
      ]
    },
     //notify of gas fee transfer
     {
       "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
       "States":[
         "transfer", //method name
         "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker's address (from)
         "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //governance contract address (to)
         10000000 //gas fee amount(decimal: 9)
       ]
     }
  ]
}
```

#### Withdraw

* Usage: Withdraw delegated authentication

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of the method
    {
      "ContractAddress": "0600000000000000000000000000000000000000", //contract address of authentication contract
      "States":[
        "withdraw",// method name
        "ea1e2adf8c19f5a7e877860264ebf326e8c3aa5a", //contract address of contract which want to achieve authentication control
        "did:tst:AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //from tstid
        "did:tst:AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //to tstid
        true //status
      ]
    },
     //notify of gas fee transfer
     {
       "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
       "States":[
         "transfer", //method name
         "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker's address (from)
         "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //governance contract address (to)
         10000000 //gas fee amount(decimal: 9)
       ]
     }
  ]
}
```

#### VerifyToken

* Usage: Verify authentication of tstid

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of the method
    {
      "ContractAddress": "0600000000000000000000000000000000000000", //contract address of authentication contract
      "States":[
        "verifyToken", // method name
        "0700000000000000000000000000000000000000", //contract address of contract which want to achieve authentication control
        "ZGlk0m9uddpBVVhDSnM3NmlqWlUzOHNlUEg5MlNuVWFvZDdQNXRVbUV4", //invoker tstid
        "registerCandidate",// function name want to verify authentication
        true //status
      ]
    },
    //notify of gas fee transfer
    {
       "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
       "States":[
         "transfer", //method name
         "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker's address (from)
         "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //governance contract address (to)
         10000000 //gas fee amount(decimal: 9)
       ]
    }
  ]
}
```