import express from 'express';
import { validateRequestBody } from '../../validators';
import { createBookingSchema } from '../../validators/booking.validator';
import {
  confirmBookingController,
  createBookingController,
} from '../../controllers/booking.controller';

const bookingRouter = express.Router();

bookingRouter.post('/', validateRequestBody(createBookingSchema), createBookingController);
bookingRouter.post('/confirm/:idempotencyKey', confirmBookingController);

export default bookingRouter;
