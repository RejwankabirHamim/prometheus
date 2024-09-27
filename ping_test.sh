# !/bin/bash

# Loop to run the curl command 10 times to test the counter is increasing
for i in {1..10}
do
  # Print the command to the terminal
  echo "$i. Executing command: curl localhost:8181/ping"

  # Execute the curl command
  curl localhost:8181/ping

done