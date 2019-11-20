# Cosmos-SDK Hands-on

This repo is target to give developers a hands on guide to quickly build up a blockchain based on Cosmos-SDK, and enable the cross-chain capability between your application chains.

- In section one, it introduces how to build up a name service(like DNS service) via using Cosmos-SDK. In this part, will mainly focus on the architectures of Cosmos-SDK and learning some basic skills/convention of Golang. It's powered by  [Cosmos-SDK-Tutorials](https://tutorials.cosmos.network/nameservice/tutorial/00-intro.html)
- In section two, we are going to build up multiple chains with Cosmos-SDK and try to let tokens go through chains, to implement the `token cross-chain` feature powered by Cosmos IBC.
  For now, section two hasn`t started yet, TBC soon.



# Secion One - Building a Blockchain by Cosmos SDK



## Introdution of Cosmos-SDK

- What's Cosmos-SDK

  The [Cosmos-SDK](https://github.com/cosmos/cosmos-sdk) is a framework for building multi-asset public Proof-of-Stake (PoS) blockchains, like the Cosmos Hub, as well as permissionned Proof-Of-Authority (PoA) blockchains.

  The goal of the Cosmos SDK is to allow developers to easily create custom blockchains from scratch that can natively interoperate with other blockchains. We envision the SDK as the npm-like framework to build secure blockchain applications on top of [Tendermint](https://github.com/tendermint/tendermint).

  More details please take a refer to [What's Cosmos](https://cosmos.network/intro)

- Structure of this demo

  ```
  ./nameservice
  ├── Makefile   #not important, defines how to run and compile this program
  ├── Makefile.ledger  #not important, ledger nano
  ├── app.go     #important, defines how to initiate your app 
  ├── cmd        #important, entry points
  │   ├── nscli
  │   │   └── main.go   #important, the main entry for users to interact with your app
  │   └── nsd
  │       └── main.go   #important, the main entry of your node
  ├── go.mod     #not important, dependencies management of Go
  ├── go.sum     #not important, dependencies log
  └── x          #important, Cosmos convention that to all app related files in
      └── nameservice     #root of nameservice
          ├── client      #important, package of commands routes
          │   ├── cli     #important, defines the commands and routes for client way
          │   │   ├── query.go     #query releated commands
          │   │   └── tx.go        #transaction related, like setName and buyName
          │   └── rest    #important, cammands and routes for REST way
          │       ├── query.go     #query related commands
          │       ├── rest.go      #routes between endpoints and commands
          │       └── tx.go        #tx related commands
          │── types   #important,like aggregate root
          │   ├── codec.go  #defines decoder 
          │   ├── key.go    #constant file, defines module name and key store
          │   ├── buyNameMsg.go  #important,defines what's the info will be wrapped in tx
          │   ├── setNameMsg.go  #important,defines what's the info will be wrapped in tx
          │   ├── querier.go # defines the result of query
          │   └── types.go #core objective/token definitions
          ├── genesis.go  #important,defines how to inital genesis state
          ├── handler.go  #important, denfines message routes
          │── keeper.go   #important, keeper uses db to maange data on chain
          │── querier.go  #top router of query
          └── module.go   #most important, the application module entry
  ```

Start by putting this repo under your GOPATH, and start the trip:

```bash
#put this repo under below path
$GOPATH/src/github.com/{ .Username }/nameservice
#then it will look like this
/Users/cwzeng/.go/src/github.com/arthaszeng/nameservice
#jump into that path
cd $GOPATH/src/github.com/{ .Username }/nameservice
git checkout develop
```

## DApp Context and Design
The goal of the application you are building is to let users buy names and to set a value these names resolve to. The owner of a given name will be the current highest bidder. In this section, you will learn how these simple requirements translate to application design.

Here are the modules you will need for the nameservice application:

- `auth`: This module defines accounts and fees and gives access to these functionalities to the rest of your application.
- `bank`: This module enables the application to create and manage tokens and token balances.
- `staking` : This module enables the application to have validators that people can delegate to.
- `distribution` : This module give a functional way to passively distribute rewards between validators and delegators.
- `slashing` : This module disincentivizes people with value staked in the network, ie. Validators.
- `supply` : This module holds the total supply of the chain.
- `nameservice`: This module does not exist yet! It will handle the core logic for the `nameservice` application you are building. It is the main piece of software you have to work on to build your application.

## Entry of Cosmos SDK - AppModule
Before we start build our application, we need to take a view of the appModule as entry which is Cosmos-SDK required.<br/>
Jump into `x/nameservice/module.go`<br/>
The code blow is to test our AppModule and AppModuleBasic has implemented the functions of the interfaces.
```
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)
```
Then you can see the functions we need to implement, such like `NewQuerierHandler` `NewHandler` and etc, those requirements will trigger the implementation, that's the reason why we need to build the structure of this demo like shown in previous section.<br/> 
This file is the entry of our nameservice app module, no need to care about the implementations for now, next, let's take a look at `GenesisState`.  

## GenesisState
The AppModule interface includes a number of functions for use in initializing and exporting GenesisState for the chain. The `ModuleBasicManager` calls these functions on each module when starting, stopping or exporting the chain. Here is a very basic implementation that you can expand upon.

Go to `x/nameservice/genesis.go` and see the work of this part.

A few notes about the above code:

- `ValidateGenesis()` validates the provided genesis state to ensure that expected invariants hold
- `DefaultGenesisState()` is used mostly for testing. This provides a minimal GenesisState.
- `InitGenesis()` is called on chain start, this function imports genesis state into the keeper.
- `ExportGenesis()` is called after stopping the chain, this function loads application state into a GenesisState stuct to later be exported to `genesis.json` alongside data from the other modules.

## Types 
Okay, we have already seen the requirements of our nameservice app module from Cosmos-SDK, let's think about our implementations from this section. The first thing we're going to do is define a struct that holds all the metadata of a name. We will call this struct Whois after the ICANN DNS terminology.

#### `types.go`

Begin by creating the file `./x/nameservice/types/types.go` to hold the customs types for your module. In Cosmos SDK applications, the convention is that modules live in the `./x/` folder.

#### Whois

Each name will have three pieces of data associated with it.

- Value - The value that a name resolves to. This is just an arbitrary string, but in the future you can modify this to require it fitting a specific format, such as an IP address, DNS Zone file, or blockchain address.
- Owner - The address of the current owner of the name
- Price - The price you will need to pay in order to buy the name


## Keeper
The main core of a Cosmos SDK module is a piece called the `Keeper`. It is what handles interaction with the store, has references to other keepers for cross-module interactions, and contains most of the core functionality of a module.

#### Keeper Struct

To start your SDK module, define your `nameservice.Keeper` in `./x/nameservice/keeper.go` file. <br/>

A couple of notes about the above code:
- 3 different `cosmos-sdk` packages are imported: - [`codec`](https://godoc.org/github.com/cosmos/cosmos-sdk/codec) - the `codec` provides tools to work with the Cosmos encoding format, [Amino](https://github.com/tendermint/go-amino). - [`bank`](https://godoc.org/github.com/cosmos/cosmos-sdk/x/bank) - the `bank` module controls accounts and coin transfers. - [`types`](https://godoc.org/github.com/cosmos/cosmos-sdk/types) - `types` contains commonly used types throughout the SDK.
- The `Keeper` struct. In this keeper there are a couple of key pieces: - [`bank.Keeper`](https://godoc.org/github.com/cosmos/cosmos-sdk/x/bank#Keeper) - This is a reference to the `Keeper` from the `bank` module. Including it allows code in this module to call functions from the `bank` module. The SDK uses an [object capabilities](https://en.wikipedia.org/wiki/Object-capability_model) approach to accessing sections of the application state. This is to allow developers to employ a least authority approach, limiting the capabilities of a faulty or malicious module from affecting parts of state it doesn't need access to. - [`*codec.Codec`](https://godoc.org/github.com/cosmos/cosmos-sdk/codec#Codec) - This is a pointer to the codec that is used by Amino to encode and decode binary structs. - [`sdk.StoreKey`](https://godoc.org/github.com/cosmos/cosmos-sdk/types#StoreKey) - This is a store key which gates access to a `sdk.KVStore` which persists the state of your application: the Whois struct that the name points to (i.e. `map[name]Whois`).

> _*NOTE*_: This function uses the [`sdk.Context`](https://godoc.org/github.com/cosmos/cosmos-sdk/types#Context). This object holds functions to access a number of important pieces of the state like `blockHeight` and `chainID`.

## Message
Now that we have the `Keeper` setup, it is time to build the `Msgs` and `Handlers` that actually allow users to buy names and set values for them.

`Msgs` trigger state transitions. `Msgs` are wrapped in [`Txs`](https://github.com/cosmos/cosmos-sdk/blob/master/types/tx_msg.go#L34-L38) that clients submit to the network. The Cosmos SDK wraps and unwraps `Msgs` from `Txs`, which means, as an app developer, you only have to define `Msgs`. `Msgs` must satisfy the following interface (we'll implement all of these in the next section):
```go
// Transactions messages must fulfill the Msg
type Msg interface {
	// Return the message type.
	// Must be alphanumeric or empty.
	Type() string

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	Route() string

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	ValidateBasic() Error

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	GetSigners() []AccAddress
}
```
In this demo, there are two msgs `x/nameservice/types/buyNameMsg.go` and `x/nameservice/types/setNameMsg.go`, both of msgs must implement above interface's functions. <br/>

Especially, there is another types of messages named `querier` which is essentially msg but just for querying instead of tx.

## Handler
`Handlers` define the action that needs to be taken (which stores need to get updated, how, and under what conditions) when a given `Msg` is received.

In this module you have three types of `Msgs` that users can send to interact with the application state: [`SetName`](set-name.md), [`BuyName`](./buy-name.md) and [`DeleteName`](./delete-name.md). They will each have an associated `Handler`.<br/>

As we mentioned before, tx handler and query handler are separated functions need to be implemented in Cosmos-SDK's appModule requirements, therefore, there is another handler `querier` at root path of `/x/nameservice` which is routes for querying and it's similar with `handler`.

## Codec

To [register your types with Amino](https://github.com/tendermint/go-amino#registering-types) so that they can be encoded/decoded, there is a bit of code that needs to be placed in `./x/nameservice/types/codec.go`. Any interface you create and any struct that implements an interface needs to be declared in the `RegisterCodec` function. In this module the three `Msg` implementations (`SetName`, `BuyName` and `DeleteName`) need to be registered, but your `Whois` query return type does not. In addition, we define a module specific codec for use later.

## Command
#### Cli
The Cosmos SDK uses the [`cobra`](https://github.com/spf13/cobra) library for CLI interactions. This library makes it easy for each module to expose its own commands. To get started defining the user's CLI interactions with the app module, create the following files:

- `./x/nameservice/client/cli/query.go`
- `./x/nameservice/client/cli/tx.go`

***Queries***

Start in `query.go`. Here, define `cobra.Command`s for each of the modules `Queriers` (`resolve`, and `whois`):

`./nameservice/x/nameservice/client/cli/query.go`

Notes on the above code:

- The CLI introduces a new `context`: [`CLIContext`](https://godoc.org/github.com/cosmos/cosmos-sdk/client/context#CLIContext). It carries data about user input and application configuration that are needed for CLI interactions.
- The `path` required for the `cliCtx.QueryWithData()` function maps directly to the names in your query router.
  - The first part of the path is used to differentiate the types of queries possible to SDK applications: `custom` is for `Queriers`.
  - The second piece (`nameservice`) is the name of the module to route the query to.
  - Finally there is the specific querier in the module that will be called.
  - In this example the fourth piece is the query. This works because the query parameter is a simple string. To enable more complex query inputs you need to use the second argument of the [`.QueryWithData()`](https://godoc.org/github.com/cosmos/cosmos-sdk/client/context#CLIContext.QueryWithData) function to pass in `data`. For an example of this see the [queriers in the Staking module](https://github.com/cosmos/cosmos-sdk/blob/develop/x/stake/querier/querier.go#L103).

***Transactions***

Now that the query interactions are defined, it is time to move on to transaction generation in `tx.go`:

`./nameservice/x/nameservice/client/cli/tx.go`

Notes on the above code:

- The `authcmd` package is used here. [The godocs have more information on usage](https://godoc.org/github.com/cosmos/cosmos-sdk/x/auth/client/cli#GetAccountDecoder). It provides access to accounts controlled by the CLI and facilitates signing.

### REST

App module can also expose a REST interface to allow programatic access to the module's functionality. To get started create a file to hold the HTTP handlers.<br/>

- ***Routes***

    First, define the REST client interface for your module in a `RegisterRoutes` function. Have the routes all start with your module name to prevent name space collisions with other modules' routes:
`./x/nameservice/client/rest/rest.go`

- ***Query Handlers***
    
    Create a `query.go` file to place all the querys in.

    Next, its time to define the handlers mentioned above. These will be very similar to the CLI methods defined earlier. Start with the queries `whois` and `resolve`:

    `/nameservice/x/nameservice/client/rest/query.go`

    Notes on the above code:
    
    - Notice we are using the same `cliCtx.QueryWithData` function to fetch the data
    - These functions are almost the same as the corresponding CLI functionality

- ***Tx Handlers***

    First define a `tx.go` file to hold all your tx rest endpoints.
    
    Now define the `buyName`, `setName` and `deleteName` transaction routes. Notice these aren't actually sending the transactions to buy, set and delete names. That would require sending a password along with the request which would be a security issue. Instead these endpoints build and return each specific transaction which can then be signed in a secure manner and afterwards broadcast to the network using a standard endpoint like `/txs`.
    
    `./nameservice/x/nameservice/client/rest/tx.go`
    
    Notes on the above code:
    
    - The [`BaseReq`](https://godoc.org/github.com/cosmos/cosmos-sdk/client/utils#BaseReq) contains the basic required fields for making a transaction (which key to use, how to decode it, which chain you are on, etc...) and is designed to be embedded as shown.
    - `baseReq.ValidateBasic` handles setting the response code for you and therefore you don't need to worry about handling errors or successes when using those functions

## App <-- Appbase

Now that app module is ready, it can be incorporated in the `./app.go` file, along with the other two modules [`auth`](https://godoc.org/github.com/cosmos/cosmos-sdk/x/auth) and [`bank`](https://godoc.org/github.com/cosmos/cosmos-sdk/x/bank). Let's look at how to adding new nameservice module to the imports `	"github.com/arthaszeng/nameservice/x/nameservice"`.<br/>

Basically, the `./app.go` file defines the `nameServiceApp` struct, and how to new and initiate a `nameServiceApp`, configuring the dependencies, as well as implementing the BaseApp interface. <br/>

In the struct of `nameServiceApp`, we need to let it 'extends' BaseApp, then add the stores' keys and the `Keepers` into it.<br/>

Beside that, when you are going to dive into this file, the best entry is the constructor `NewNameServiceApp`. <br/>

The constructor needs to:

- Instantiate required `Keepers` from each desired module.
- Generate `storeKeys` required by each `Keeper`.
- Register `Handler`s from each module. The `AddRoute()` method from `baseapp`'s `router` is used to this end.
- Register `Querier`s from each module. The `AddRoute()` method from `baseapp`'s `queryRouter` is used to this end.
- Mount `KVStore`s to the provided keys in the `baseApp` multistore.
- Set the `initChainer` for defining the initial application state.

> _*NOTE*_: The TransientStore mentioned above is an in-memory implementation of the KVStore for state that is not persisted.

> _*NOTE*_: Pay attention to how the modules are initiated: the order matters! Here the sequence goes Auth --> Bank --> Feecollection --> Staking --> Distribution --> Slashing, then the hooks were set for the staking module. This is because some of these modules depend on others existing before they can be used.

The `initChainer` defines how accounts in `genesis.json` are mapped into the application state on initial chain start. The `ExportAppStateAndValidators` function helps bootstrap the initial state for the application. You don't need to worry too much about either of these for now. We also need to add a few more methods to our app `BeginBlocker`, `EndBlocker` and `LoadHeight`.


## Go Entry Points

In Golang the convention is to place files that compile to a binary in the `./cmd` folder of a project. For your application there are 2 binaries need to create:

- `nsd`: This binary is similar to `bitcoind` or other cryptocurrency daemons in that it maintains p2p connections, propagates transactions, handles local storage and provides an RPC interface to interact with the network. In this case, Tendermint is used for networking and transaction ordering.
- `nscli`: This binary provides commands that allow users to interact with the application.

To get started create two files in your project directory that will instantiate these binaries:

- `./cmd/nsd/main.go`
    - Most of the code combines the CLI commands from Tendermint, Cosmos-SDK and the Nameservice app module.
- `./cmd/nscli/main.go`
    - The code combines the CLI commands from Tendermint, Cosmos-SDK and the Nameservice app module.
    - The [`cobra` CLI documentation](http://github.com/spf13/cobra) will help with understanding the above code.
    - You can see the `ModuleClient` defined earlier in action here.
    - Note how the routes are included in the `registerRoutes` function.

## Go Mod & Makefile

***go.mod***

Golang has a few dependency management tools. In this tutorial you will be using [`Go Modules`](https://github.com/golang/go/wiki/Modules). `Go Modules` uses a `go.mod` file in the root of the repository to define what dependencies the application needs. Cosmos SDK apps currently depend on specific versions of some libraries.

- `./go.mod`

***Makefile***

Help users build your application by writing a `./Makefile` in the root directory that includes common commands:

> _*NOTE*_: The below Makefile contains some of same commands as the Cosmos SDK and Tendermint Makefiles.

- `./nameservice/Makefile`
- `./nameservice/Makefile.ledger`

## Building the `nameservice` application

If you want to build the `nameservice` application in this repo to see the functionalities, **Go 1.13.0+** is required .

### Prepare Go environment

Add some parameters to environment is necessary if you have never used the `go mod` before.

```bash
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
echo "export GOBIN=\$GOPATH/bin" >> ~/.bash_profile
echo "export PATH=\$PATH:\$GOBIN" >> ~/.bash_profile
echo "export GO111MODULE=on" >> ~/.bash_profile
source ~/.bash_profile
```

### Building the app

```bash
# Install the app into your $GOBIN
make install

# Now you should be able to run the following commands:
nsd help
nscli help

# Or check the version of them
nsd version
nscli version
```


## Running the live network and using the commands

To initialize configuration and a `genesis.json` file for your application and an account for the transactions, start by running:

> _*NOTE*_: In the below commands addresses are pulled using terminal utilities. You can also just input the raw strings saved from creating keys, shown below. The commands require [`jq`](https://stedolan.github.io/jq/download/) to be installed on your machine.

> _*NOTE*_: If you have run the tutorial before, you can start from scratch with a `nsd unsafe-reset-all` or by deleting both of the home folders `rm -rf ~/.ns*`

> _*NOTE*_: If you have the Cosmos app for ledger and you want to use it, when you create the key with `nscli keys add jack` just add `--ledger` at the end. That's all you need. When you sign, `jack` will be recognized as a Ledger key and will require a device.

```bash
# Initialize configuration files and genesis file
  # moniker is the name of your node
nsd init <moniker> --chain-id namechain


# Copy the `Address` output here and save it for later use
# [optional] add "--ledger" at the end to use a Ledger Nano S
nscli keys add jack

# Copy the `Address` output here and save it for later use
nscli keys add alice

# Add both accounts, with coins to the genesis file
nsd add-genesis-account $(nscli keys show jack -a) 1000nametoken,100000000stake
nsd add-genesis-account $(nscli keys show alice -a) 1000nametoken,100000000stake

# Configure your CLI to eliminate need for chain-id flag
nscli config chain-id namechain
nscli config output json
nscli config indent true
nscli config trust-node true

nsd gentx --name jack <or your key_name>
```

After you have generated a genesis transcation, you will have to input the gentx into the genesis file, so that your nameservice chain is aware of the validators. To do so, run:

`nsd collect-gentxs`

and to make sure your genesis file is correct, run:

`nsd validate-genesis`

You can now start `nsd` by calling `nsd start`. You will see logs begin streaming that represent blocks being produced, this will take a couple of seconds.

You have run your first node successfully.

```bash
# First check the accounts to ensure they have funds
nscli query account $(nscli keys show jack -a)
nscli query account $(nscli keys show alice -a)

# Buy your first name using your coins from the genesis file
nscli tx nameservice buy-name jack.id 5nametoken --from jack

# Set the value for the name you just bought
nscli tx nameservice set-name jack.id 8.8.8.8 --from jack

# Try out a resolve query against the name you registered
nscli query nameservice resolve jack.id
# > 8.8.8.8

# Try out a whois query against the name you just registered
nscli query nameservice whois jack.id
# > {"value":"8.8.8.8","owner":"cosmos1l7k5tdt2qam0zecxrx78yuw447ga54dsmtpk2s","price":[{"denom":"nametoken","amount":"5"}]}

# Alice buys name from jack
nscli tx nameservice buy-name jack.id 10nametoken --from alice

# Try out a whois query against the name you just deleted
nscli query nameservice whois jack.id
# > {"value":"","owner":"","price":[{"denom":"nametoken","amount":"1"}]}
```
