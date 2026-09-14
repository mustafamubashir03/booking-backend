import { Job, Worker } from 'bullmq';
import { roomGenerationJobDTO } from '../dto/roomCategory.dto';
import { ROOM_GENERATION_QUEUE } from '../queues/roomGeneration.queue';
import { getRedisConnection } from '../config/redis.config';
import logger from '../config/logger.config';
import { generateRooms } from '../services/roomGeneration.service';


export const setupRoomGenerationWorker = () => {

    const roomGenerationProcessor = new Worker<roomGenerationJobDTO>(
        ROOM_GENERATION_QUEUE,
        async (job: Job) => {
            if (job.name !== ROOM_GENERATION_QUEUE) throw new Error('Invalid job name');
            //call service layer
            const payload = job.data;
            await generateRooms(payload);
            logger.info(`Room generation job completed successfully`);

        },
        {
            connection: getRedisConnection(),
        },
    );

    roomGenerationProcessor.on('failed', () => {
        logger.error('Room generation processing failed');
    });

    roomGenerationProcessor.on('completed', () => {
        logger.info('Room generation processing completed');
    });
};
