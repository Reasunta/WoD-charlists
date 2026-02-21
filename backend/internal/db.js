const { Pool } = require('pg');
const config = require('config')

const pool = new Pool({
  user: config.PS_DB_USER,
  host: config.PS_DB_HOST,
  database: config.PS_DB_NAME,
  password: config.PS_DB_PASSWORD,
  port: 5432,
});

const callSQLSelect = async (query) => {
    try {
        const result = await pool.query(query);
        return {data: result.rows, error: null};
    } catch (err) {
        return {data: null, error: err};
    }
}

module.exports.health = async () => {
  return callSQLSelect('SELECT NOW()')
};