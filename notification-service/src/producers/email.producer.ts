import { NotificationDto } from '../dto/notification.dto';
import { MAILER_QUEUE, mailerQueue } from '../queues/mailer.queue';

export const addEmailToQueue = async (payload: NotificationDto) => {
  try {
    await mailerQueue.add(MAILER_QUEUE, payload);
    console.log('Email added to queue', payload);
  } catch (error) {
    console.log('Failed to add email to queue', error);
    throw error;
  }
};
