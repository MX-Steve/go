package main

// use channel and goroutine to complete write log data
// 1. send log data into channel
//   1. determine the format of log data
//   2. use buffer in channel, how much,
//   3. Logs that are too late to write are discarded
// 2. run one goroutine in the backend to write log data into file
