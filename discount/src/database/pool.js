const { Pool } = require('pg');
const databaseURL = process.env.POSTGRESQL_URL;

const databaseConfig = { connectionString:  databaseURL};
const pool = new Pool(databaseConfig);

module.exports = pool;