import express from 'express';
import pingRouter from './ping.router';
import hotelRouter from './hotel.router';
import roomRouter from './room.router';
import schedulerRouter from './scheduler.router';

const v1Router = express.Router();

v1Router.use('/ping', pingRouter);
v1Router.use('/hotels', hotelRouter);
v1Router.use('/rooms', roomRouter);
v1Router.use('/scheduler', schedulerRouter);

export default v1Router;
