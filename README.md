# Democracy
 The system looks for pull requests on its own code and then merges those requests and reboots with the new code if the PRs recieve enough upvotes. 
 To add excitement, this system controls a cryptocurrency wallet I have funded. Successfully merged PRs will determine how this money is spent.

 This project is meant to be a more purposeful reimpementation of the popular [ChaosBot](https://github.com/Chaosthebot/Chaos).

## High level goals and concepts

### The Corporation
All assets are controlled by a democratically ran corporation with shares distributed evenly through citizenship. The corporate structure of the project allows it to run in subordination and cooperation of existing democratic authorities and jurisdictions.

### Shares
Every citizen of the corporation gets 1 share. However, subsidiaries of the corporation can have a more traditional ownership model where citizens are allocated shares based on how much they have contributed to the revenue stream.

### The Constitution
The constitution is a module of code that govern's vital resources of the corporation. It does the following:
- Controls citizenship of the corporation (naturalization, renuciation, and exile)
- Ratifies ammendments to the constitution through a Pull Request voting system for citizens
- Grants wallet perimissions to other modules
- Specifies how modules are added to the construct
- Distributes a Freedom Dividend to all citizens

### Constructs
A Construct is a computer system that runs the Constitution in such a way that the Democracy Repository can not be changed except through consitutionally approved means (PR voting by citizens). A Construct also contains private keys (for the crypto wallets and other resources) that only it has access too. 

A Construct will have a RSA key called a Construct Key. A Construct Key will be used to encrypt secrets that can then be uploaded to the Democracy Repository for backup. The Construct will generate a private and public RSA key for each of the citizens and will use the citizens private keys to encrypt the Construct Key in such a way that it requires 99% of the citizens' private keys to decrypt it. This encrypted Construct Key will also be uploaded to the Democracy repository for backup. 

In the event that a new Construct needs to be instanciated, a currently running Construct will simply spin up a new Construct, generate its RSA keys, and encrypt all the secrets with the newly generated Construct Key. A Construct can only be SSH'ed into by using a Construct Key. In the event that all Constructs die, a supermajority of 99% of the citizens can decrypt one of the dead Construct's keys and instanciate a new Construct instance with it.

### The wallet
A cryptocurrency wallet that holds the corporations financial resources. Modules that are approved by the constitution will have access to the wallet and can programatically purchase products and services using the wallet.

### Citizenship
Citizens of the corporation:
- participate in democratic functions of the corporation. These include financial allocations, module acceptance, constitutional ammendment ratification, and citizenship control
- recieve a UBI

# Citizens
- @kuberlog
