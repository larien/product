const pool = require('./pool')

pool.on('connect', () => {
    console.log('connected to the db');
  });

const selectUser = async (userID) => {
    const query = `
      SELECT first_name, last_name, TO_CHAR(date_of_birth :: DATE, 'yyyy-mm-dd') as date_of_birth
      FROM users
      WHERE id = $1
      LIMIT 1`;
    try {
      let result = await pool.query(query, [userID])
      return result.rows[0]
    } catch(e) {
      console.log(e)
    }
};
  
module.exports = { selectUser }