/**
 * @author Daren Ileleji
 * @version 1.0.0
 * @date 2025-12-10
 * @fileoverview This program will ask the user for bank, interest rate, and post secondary cost info to determine how many years it will take
 */

package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {

// Variables
  var bankAccountString string = ""
  var bankAccountNumber float64 = 0

  var interestRateString string = ""
  var interestRateNumber float64 = 0

  var postSecondaryString string = ""
  var postSecondaryNumber float64 = 0

  var years int = 0

// Input
  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Enter the starting bank amount please: ")
  bankAccountString, _ = reader.ReadString('\n')
  bankAccountString = strings.TrimSpace(bankAccountString)
  bankAccountNumber, _ = strconv.ParseFloat(bankAccountString, 64)

  fmt.Print("Enter the yearly interest rate please: ")
  interestRateString, _ = reader.ReadString('\n')
  interestRateString = strings.TrimSpace(interestRateString)
  interestRateNumber, _ = strconv.ParseFloat(interestRateString, 64)

  fmt.Print("Enter the amount needed for post secondary please: ")
  postSecondaryString, _ = reader.ReadString('\n')
  postSecondaryString = strings.TrimSpace(postSecondaryString)
  postSecondaryNumber, _ = strconv.ParseFloat(postSecondaryString, 64)

// Processing
  for bankAccountNumber < postSecondaryNumber {
    bankAccountNumber = bankAccountNumber + (bankAccountNumber * (interestRateNumber / 100))
    years++
    }

// Output
  fmt.Println("It will take", years, "years for the starting bank account to reach the required amount for post-secondary education.")
  fmt.Println("\nDone.")
}
