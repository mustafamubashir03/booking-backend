import { Request, Response } from 'express';
import {
    startRoomGenerationScheduler,
    stopRoomGenerationScheduler,
    getRoomGenerationSchedulerStatus
} from '../schedulers/roomGeneration.scheduler';
import { StatusCodes } from 'http-status-codes';

export const startScheduler = (req: Request, res: Response) => {
    startRoomGenerationScheduler();
    res.status(StatusCodes.OK).json({
        success: true,
        message: 'Scheduler started/resumed successfully'
    });
};

export const stopScheduler = (req: Request, res: Response) => {
    stopRoomGenerationScheduler();
    res.status(StatusCodes.OK).json({
        success: true,
        message: 'Scheduler stopped successfully'
    });
};

export const getSchedulerStatus = (req: Request, res: Response) => {
    const status = getRoomGenerationSchedulerStatus();
    res.status(StatusCodes.OK).json({
        success: true,
        data: status
    });
};
