import cron, { ScheduledTask } from 'node-cron';
import logger from '../config/logger.config';
import { getRoomCategoryService } from '../services/roomCategory.service';
import { roomRepository } from '../services/roomGeneration.service';
import { addRoomGenerationJobToQueue } from '../producers/roomGeneration.producer';
import { cronConfig } from '../config';




let schedulerTask: ScheduledTask | null = null;

const schedulerJob = async () => {
    logger.info('Starting room generation scheduler job');
    try {
        const categories = await getRoomCategoryService();

        for (const category of categories) {
            const latestRoom = await roomRepository.findLatestByRoomCategoryId(category.id);

            let startDate = new Date();
            startDate.setHours(0, 0, 0, 0);
            startDate.setDate(startDate.getDate() + 1); // start from tomorrow

            if (latestRoom && latestRoom.dateOfAvailability) {
                const lastAvailableDate = new Date(latestRoom.dateOfAvailability);
                lastAvailableDate.setHours(0, 0, 0, 0);

                // If we have rooms available beyond today, start generating from the day after the last available
                if (lastAvailableDate >= startDate) {
                    startDate = new Date(lastAvailableDate);
                    startDate.setDate(startDate.getDate() + 1);
                }
            }

            const targetEndDate = new Date();
            targetEndDate.setHours(0, 0, 0, 0);
            targetEndDate.setDate(targetEndDate.getDate() + cronConfig.FUTURE_HORIZON_DAYS);

            // If the start date is before the target end date, we need to generate rooms
            if (startDate < targetEndDate) {
                const startDateString = startDate.toISOString().split('T')[0];
                const endDateString = targetEndDate.toISOString().split('T')[0];
                const jobId = `room-gen-${category.id}-${startDateString}-${endDateString}`;

                await addRoomGenerationJobToQueue({
                    roomCategoryId: category.id,
                    startDate: startDate.toISOString(),
                    endDate: targetEndDate.toISOString(),
                    batchSize: 10,
                }, { jobId });

                logger.info(`Enqueued room generation for category ${category.id} from ${startDateString} to ${endDateString}`);
            }
        }
        logger.info('Room generation scheduler job completed');
    } catch (error) {
        logger.error('Error in room generation scheduler', error);
    }
};

export const startRoomGenerationScheduler = () => {
    if (!schedulerTask) {
        schedulerTask = cron.schedule(cronConfig.CRON_SCHEDULE, schedulerJob);
        logger.info('Room generation scheduler started');
    } else {
        schedulerTask.start();
        logger.info('Room generation scheduler resumed');
    }
};

export const stopRoomGenerationScheduler = () => {
    if (schedulerTask) {
        schedulerTask.stop();
        logger.info('Room generation scheduler stopped');
    }
};

export const getRoomGenerationSchedulerStatus = () => {
    // node-cron ScheduledTask doesn't have a public status API, but we know it's stopped if we stop it.
    // Actually, node-cron tasks have an internal status, but a simple boolean flag or tracking if we stopped it is easier, 
    // but typically if we just rely on the task being created. Let's just return true if it's running.
    // However, node-cron tasks don't expose if they are stopped. 
    // We will just return a basic status.
    return {
        isConfigured: !!schedulerTask,
    };
};
