import { JobsOptions } from 'bullmq';
import { roomGenerationJobDTO } from '../dto/roomCategory.dto';
import { ROOM_GENERATION_QUEUE, roomGenerationQueue } from '../queues/roomGeneration.queue';

export const addRoomGenerationJobToQueue = async (payload: roomGenerationJobDTO, options?: JobsOptions) => {
    try {
        await roomGenerationQueue.add(ROOM_GENERATION_QUEUE, payload, options);
        console.log('Room generation job added to queue', payload);
    } catch (error) {
        console.log('Failed to add room generation job to queue', error);
        throw error;
    }
};
