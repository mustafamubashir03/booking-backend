import express from 'express';
import {
    startScheduler,
    stopScheduler,
    getSchedulerStatus
} from '../../controllers/scheduler.controller';

const schedulerRouter = express.Router();

schedulerRouter.post('/start', startScheduler);
schedulerRouter.post('/stop', stopScheduler);
schedulerRouter.get('/status', getSchedulerStatus);

export default schedulerRouter;
