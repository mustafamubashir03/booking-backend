// This file contains all the basic configuration logic for the app server to work
import dotenv from 'dotenv';

type ServerConfig = {
  PORT: number;
  REDIS_SERVER_URL: string;
};
type dbConfig = {
  DB_USER: string;
  DB_NAME: string;
  DB_PASSWORD: string;
  DB_HOST: string;
};

type cronConfig = {
  CRON_SCHEDULE: string;
  FUTURE_HORIZON_DAYS: number;
}

function loadEnv() {
  dotenv.config();
  console.log(`Environment variables loaded`);
}

loadEnv();

export const serverConfig: ServerConfig = {
  PORT: Number(process.env.PORT) || 3001,
  REDIS_SERVER_URL: process.env.REDIS_SERVER_URL || 'redis://localhost:6379',
};

export const dbConfig: dbConfig = {
  DB_USER: process.env.DB_USER || '',
  DB_NAME: process.env.DB_NAME || '',
  DB_PASSWORD: process.env.DB_PASSWORD || '',
  DB_HOST: process.env.DB_HOST || '',
};

export const cronConfig: cronConfig = {
  CRON_SCHEDULE: process.env.CRON_SCHEDULE || '0 12 * * *',
  FUTURE_HORIZON_DAYS: Number(process.env.FUTURE_HORIZON_DAYS) || 30,
}
