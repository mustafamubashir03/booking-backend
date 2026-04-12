import express from 'express';
import {
  createHotelController,
  deleteHotelByIdController,
  getHotelByIdController,
  getHotelController,
  updateHotelByIdController,
} from '../../controllers/hotel.controller';
import { validateRequestBody } from '../../validators';
import { hotelCreateSchema, hotelUpdateSchema } from '../../validators/hotel.validator';

const hotelRouter = express.Router();

hotelRouter.get('/', getHotelController);
hotelRouter.post('/', validateRequestBody(hotelCreateSchema), createHotelController);
hotelRouter.get('/:id', getHotelByIdController);
hotelRouter.put('/:id', validateRequestBody(hotelUpdateSchema), updateHotelByIdController);
hotelRouter.delete('/:id', deleteHotelByIdController);

export default hotelRouter;
