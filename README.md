### PASH - Password Hasher

Simple password hashing and verification utility. Uses bcrypt.

There are 4 hashing options provided.
1. `HashMinCost()`: uses `bcrypt.MinCost` currently `4` _(at the time of writing)_
2. `Hash()`: uses `bcrypt.DefaultCost` currently `10`
3. `HashMaxCost()`: uses `bcrypt.MaxCost` currently `31`
4. `HashWithCost()`: accepts a variable cost factor. Use with care. The higher the cost factor, the more time it takes to
yield a result _(or crack)_.



#### Why?
Because I got tired of duplicating across pet projects.