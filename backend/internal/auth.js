const crypto = require("./crypto");
const random = require("./random");
const db = require("./db");

module.exports.register = async function(name, login, password) {
    const decodedName = crypto.fromBase64(name);
    const decodedLogin = crypto.fromBase64(login);
    const decodedPassword = crypto.fromBase64(password);

    const salt = random.randomString(32);
    const password_hash = crypto.sha256(decodedPassword.concat(salt))
    const db_health = await db.health();

    return {
        name: decodedName,
        login: decodedLogin,
        password: decodedPassword,
        salt: salt,
        hash: password_hash,
        db: db_health
    }
}
