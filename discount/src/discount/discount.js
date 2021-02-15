const db = require('../database/db')
const Percentage = require('./percentage')

const getPercentage = async (productID, userID) => {
    console.log("productID: " + productID, "| userID: ", userID)
    try {
        let user = await db.selectUser(userID);
        if (user == null){
            console.log("no user was found")
            return 0
        }

        const percentage = new Percentage();
        percentage
            .applyBlackfriday()
            .applyBirthday(user.date_of_birth)

        console.log("applied discount percentage: " + percentage.value)
        return percentage.value
    } catch (error) {
        console.log("failed to calculate the percentage: " + error)
        return 0
    }
};

module.exports = { getPercentage };