import { Job, Worker } from 'bullmq';
import { NotificationDto } from '../dto/notification.dto';
import { MAILER_QUEUE } from '../queues/mailer.queue';
import { getRedisConnection } from '../config/redis.config';

export const setupMailerWorker = () => {

  const emailProcessor = new Worker<NotificationDto>(
    MAILER_QUEUE,
    async (job: Job) => {
      if (job.name !== MAILER_QUEUE) throw new Error('Invalid job name');
      //call service layer
    },
    {
      connection: getRedisConnection(),
    },
  );

  emailProcessor.on('failed', () => {
    console.log('Email processing failed');
  });

  emailProcessor.on('completed', () => {
    console.log('Email processing completed');
  });
};
