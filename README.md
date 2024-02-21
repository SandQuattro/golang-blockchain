![image](https://github.com/SandQuattro/golang-blockchain/assets/31468131/83182959-5c74-49fa-bdb8-333d0e0a3c6a)

# golang-blockchain

mine method is the most interesting part

mine  is used to mine a block on a blockchain. 
The function takes in a parameter difficulty to determine the level of difficulty for mining. 
If the difficulty is 0, it sets the proof of work (pow) to 0 and generates the hash. 
Otherwise, it continues to generate hashes until the hash begins with the required number of leading zeros 
determined by the difficulty. 
The pow variable keeps track of the number of attempts made to obtain the hash of the current block.
# Future plans

- block corruption and validation