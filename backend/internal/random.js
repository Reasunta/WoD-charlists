const { randomBytes } = require("node:crypto");

const randomString = (length) => {
  if (length % 2 !== 0) {
    length++;
  }

  return randomBytes(length).map(b => 32 + b % (126 - 32 + 1)).toString();
}

module.exports.randomString = randomString;