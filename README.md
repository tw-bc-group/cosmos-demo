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


## Types 



## Keeper



## Message



## Handler



## Codec



## Command 

### Cli

### REST

## App <-- Appbase


## Go Entry Points



## Go Mod & Makefile



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
