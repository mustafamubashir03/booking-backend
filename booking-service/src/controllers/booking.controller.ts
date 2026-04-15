import { Request, Response } from 'express';
import { confirmBookingService, createBookingService } from '../services/booking.service';

export async function createBookingController(req: Request, res: Response) {
  const bookingCreated = await createBookingService(req.body);
  res.status(201).json({
    success: true,
    message: 'booking created successfully',
    data: bookingCreated?.idempotencyKey,
  });
}

export async function confirmBookingController(req: Request, res: Response) {
  const bookingConfirmed = await confirmBookingService(req.params.idempotencyKey);
  res.status(201).json({
    success: true,
    message: 'booking confirmed successfully',
    data: bookingConfirmed,
  });
}
