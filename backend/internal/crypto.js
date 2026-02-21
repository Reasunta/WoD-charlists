const config = require("config")
const crypto = require('crypto');

module.exports.fromBase64 = (s) => Buffer.from(s, 'base64').toString('utf8')
module.exports.sha256 = (s) => crypto.createHmac('sha256', config.HMAC_PRIVATE_KEY).update(s).digest('hex');
module.exports.isHashCorrect = (s, hash) => crypto.createHmac('sha256', config.HMAC_PRIVATE_KEY).update(s).digest('hex') === hash;
