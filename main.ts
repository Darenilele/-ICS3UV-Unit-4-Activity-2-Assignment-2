/**
 * @author Daren Ileleji
 * @version 1.0.0
 * @date 2025-12-10
 * @fileoverview This program will ask the user for bank, intrest rate, and post secondary cost info to determine how many years it will take
 */


// Variables
let bankAccountString: string = "";
let bankAccountNumber: number = 0;

let intrestRateString: string = "";
let intrestRateNumber: number = 0;

let postSecondaryString: string = "";
let postSecondaryNumber: number = 0;

let years: number = 0;

// Input
bankAccountString = prompt("Enter the starting bank amount please:") || "this is an invalid number"
bankAccountNumber = parseInt(bankAccountString);

intrestRateString = prompt("Enter the yearly intrest rate please:") || "this is an invalid number"
intrestRateNumber = parseInt(intrestRateString);

postSecondaryString = prompt("Enter the amount needed for post secondary please:") || "this is an invalid number"
postSecondaryNumber = parseInt(postSecondaryString);

// Processing 
while (bankAccountNumber < postSecondaryNumber) {
  bankAccountNumber = bankAccountNumber + (bankAccountNumber * (intrestRateNumber / 100));
  years++;
}

//Output
console.log(`It will take ${years} years for the starting bank account to reach the required amount for post-secondary education.`)
console.log("\nDone.")