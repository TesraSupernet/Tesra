# Governance contract

common event format is as follows, including txhash, state, gasConsumed and notify, each native contract method have different notifies.

|key|description|
|:--|:--|
|TxHash|transaction hash|
|State|1 indicates success，0 indicates fail|
|GasConsumed|gas fee consumed by this transaction|
|Notify|Notify event|

#### RegisterCandidate

* Usage: Register to become a candidate node

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of the auth contract
    {
      "ContractAddress": "0600000000000000000000000000000000000000", //contract address of auth contract
      "States":[
        "verifyToken", //method name
        "0700000000000000000000000000000000000000", //governance contract address
        "ZGlk0m9uddpBVVhDSnM3NmlqWlUzOHNlUEg5MlNuVWFvZDdQNXRVbUV4", //invoker tstid
        "registerCandidate",// authorize function name
        true //status
      ]
    },
    //notify of tst transfer
    {
      "ContractAddress": "0100000000000000000000000000000000000000", //tst contract address
      "States":[
        "transfer",// method name
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //from address
        "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //to address
        100 //transfer amount
      ]
    },
    //notify of tsg transfer
    {
      "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
      "States":[
        "transfer",// method name
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //from address
        "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //to address
        100 //transfer amount
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

#### UnRegisterCandidate

* Usage: Cancel register candidate request

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
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

#### QuitNode

* Usage: Quit candidate node

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
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

#### AuthorizeForPeer

* Usage: Authorize tst to a candidate node

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of tst transfer
    {
      "ContractAddress":"0100000000000000000000000000000000000000", //tst contract address
      "State":[
        "transfer", //method name
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //from address
        "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //to address
        10000000 //transfer amount
      ]
    },
    //unbounded tsg transfer
    {
      "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
      "States":[
        "transfer", // method name
        "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV", //tst contract address
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker address
        10000000 //unbounded tsg amount
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

#### UnAuthorizeForPeer

* Usage: Cancel the authorize tst to a candidate node

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
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

* Usage: Withdraw deposit tst

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of tst transfer
    {
      "ContractAddress": "0100000000000000000000000000000000000000", //tsg contract address
      "States":[
        "transfer", // method name
        "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",// governance contract
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA",// invoker address
        10000000 // withdraw amount
      ]
    },
    //unbounded tsg transfer
    {
      "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
      "States":[
        "transfer", // method name
        "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV", //tst contract address
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker address
        10000000 //unbounded tsg amount
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

#### WithdrawTsg

* Usage: Withdraw unbounded tsg

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of tst transfer to trigger unbounded tsg
    {
      "ContractAddress": "0100000000000000000000000000000000000000",
      "States":[
        "transfer", //method name
        "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //governance contract address
        "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //governance contract address
        1 //fixed amount
      ]
    },
    //unbounded tsg transfer
    {
      "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
      "States":[
        "transfer", // method name
        "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV", //tst contract address
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker address
        10000000 //unbounded tsg amount
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

#### AddInitPos

* Usage: Add node's init pos

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
    //notify of tst transfer
    {
      "ContractAddress": "0100000000000000000000000000000000000000", //tst contract address
      "States":[
        "transfer", //method name
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker address
        "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", //governance address
        1000 //add init pos amount
      ]
    },
    //unbounded tsg transfer
    {
      "ContractAddress": "0200000000000000000000000000000000000000", //tsg contract address
      "States":[
        "transfer", // method name
        "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV", //tst contract address
        "AbPRaepcpBAFHz9zCj4619qch4Aq5hJARA", //invoker address
        10000000 //unbounded tsg amount
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

#### ReduceInitPos

* Usage: Reduce node's init pos

* Event and notify:
```
{
  "TxHash":"",
  "State":1,
  "GasConsumed":10000000,
  "Notify":[
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
