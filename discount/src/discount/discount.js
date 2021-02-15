const db = require('../database/db')
const Percentage = require('./percentage')

const getPercentage = async (requestID, productID, userID) => {
    console.log("requestID: " + requestID + " | productID: " + productID, "| userID: ", userID)
    try {
        let user = await db.selectUser(userID);
        if (user == null){
            console.log("requestID: " + requestID + " | no user was found")
            return 0
        }

        const percentage = new Percentage(requestID);
        percentage
            .applyBlackfriday()
            .applyBirthday(user.date_of_birth)

        console.log("requestID: " + requestID + " | applied discount percentage: " + percentage.value)
        return percentage.value
    } catch (error) {
        console.log("requestID: " + requestID + " | failed to calculate the percentage: " + error)
        return 0
    }
};

module.exports = { getPercentage };