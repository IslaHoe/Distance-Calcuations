# Distance-Calcuations

**Requirements**
If you dont have Go installed please install it from https://golang.org/doc/install or if you're on a mac and have brew installed pleas run `brew update&& brew install golang`

**Directory Structure**
1. `main.go` is the core file used to run the program 
2. `customers.go` conatins all functions required to calcuate the distance between each customer and the office location 
3. `customers.sh` is a shell script which contains the following commands 
    * Removes any versions of customers.txt which are already in the current directory to ensure the most up to date list of customers is taken from  https://s3.amazonaws.com/intercom-take-home-test/customers.txt >> customers.txt
    * Builds the go program in to an exacutable file named codingChallange
    * Opens the output of the the program in your default text editor. 
4. `customer_test.go` contains the test function for the programs, in this file the functions tested were the calucations functions. Error checking for functions to open the files etc were left in the main.go file to follow convention. 

**Running the Program**
1. Ensure Go is installed 
2. Download the zip file or clone 
3. Navigate to ~/< your path >/distanceCalcuations
4. From the command line simply run `./customers.sh` 
5. Alternativly run `go build` and then run the excutable file `./<currentDirName>`
6. To run the test file run `go test`

