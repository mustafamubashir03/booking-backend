require('ts-node/register');

const rawConfig = require('./db.config');

const config = rawConfig.default || rawConfig;

module.exports = config;
