import Redis from 'ioredis';
import { serverConfig } from '.';

export const connectToRedis = () => {
  try {
    let connection: Redis;
    return () => {
      if (!connection) {
        connection = new Redis(serverConfig.REDIS_SERVER_URL, {
          maxRetriesPerRequest: null,
        });
        return connection;
      }
      return connection;
    };
  } catch (error) {
    console.log('Failed to connect redis');
    throw error;
  }
};

export const getRedisConnection = connectToRedis();
