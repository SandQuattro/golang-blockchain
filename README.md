![image](https://github.com/SandQuattro/golang-blockchain/assets/31468131/83182959-5c74-49fa-bdb8-333d0e0a3c6a)

# golang-blockchain

mine method is the most interesting part

The mine function, implemented in the context of blockchain technology, performs a mining process based on a given 
difficulty level. 

Closer look at what happens at each stage of its work.

1. Function parameters
   difficulty - mining difficulty level. The higher the value, the more difficult it is to find a suitable hash.
2. Checking the complexity condition
   At the very beginning, the function checks whether the difficulty level is set (difficulty). If the difficulty is 
   less than or equal to zero (difficulty <= 0), that is, mining is not required or an incorrect difficulty level is set, the following actions are performed in the block:

The value of pow (Proof of Work) is set to 0.
The generateHash method is then called, which generates a hash for the current block based on its contents and pow value.
The function then exits using the return statement.
3. Generating a suitable hash
   If the difficulty condition is not met (that is, difficulty > 0), then the algorithm proceeds to the main mining cycle, the purpose of which is to find a block hash that satisfies a certain difficulty condition.

The cycle works like this:

An infinite loop (for) is used, which continues until a hash that meets the complexity condition is generated.
The difficulty condition is that the hash must start with a given number of zeros. This number of zeros is determined by the difficulty parameter.
At each iteration:
The pow counter is incremented (b.pow++). This change produces different hashes by changing the pow, which is part of the data from which the hash is generated.
The generateHash method is called to generate a new hash with an increased pow value.
The loop is terminated when a hash that satisfies the complexity condition (starting with the specified number of zeros) is found.
Thus, the mine function ensures mining is carried out by iteratively searching for a hash that matches a given difficulty level. This process is a key mechanism for ensuring security and decentralization in blockchain networks.

# Benchmark
BenchmarkMine
BenchmarkMine-8   	   14839	     80656 ns/op
PASS

# Future plans

- more tests with random payload and blockchain benchmark
- block corruption and validation
- gob serialization / deserialization
