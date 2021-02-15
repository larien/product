const MAX_PERCENTAGE = 10

// This file is used to create discounts to be applied in the percentage.
// Usage:
// const percentage = new Percentage()
// percentage
//  .applyBlackfriday()
//  .applyBirthday()
//  .applyAnotherDiscount()
// This implementation stil isn't ideal. All of the methods in the chain will be called even
// if the MAX_PERCENTAGE is reached. Ideally there should be a break mechanism to avoid
// unecessary process.

module.exports = class Percentage {
    constructor(requestID){
      this.requestID = requestID
      this.value = 0;
    }
    isMax() {
      if (this.value > MAX_PERCENTAGE) {
        this.value = MAX_PERCENTAGE;
      }
    }
    applyBlackfriday() {
        let today = new Date()
        let isBF = today.getDay() == 25 && today.getMonth() == 11
        if (isBF) {
            console.log("requestID: " + this.requestID + " | blackfriday discount applied")
            this.value += 10
        }
        this.isMax()
        return this
    }
    applyBirthday(dateOfBirth) {
        let today = new Date()
        let date = new Date(dateOfBirth)
        let isBirthday = (date.getDay() == today.getDay() && date.getMonth() == today.getMonth())
        if (isBirthday){
            console.log("requestID: " + this.requestID + " | user birthday discount applied")
            this.value += 5
        }
        this.isMax()
        return this
    }
  }