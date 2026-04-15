import {
  confirmBooking,
  createBooking,
  createIdempotentKey,
  finalizeIdempotentKey,
  getIdempotentKey,
} from '../repositories/booking.repository';
import { Prisma } from '../../prisma/generated/prisma/client';
import { generateIdempotencyKey } from '../utils/helpers/generateIdempotencyKey';
import { BadRequestError, NotFoundError } from '../utils/errors/app.error';

export async function createBookingService({
  userId,
  hotelId,
  totalGuests,
  bookingAmount,
}: Prisma.BookingCreateInput) {
  const bookingCreated = await createBooking({
    userId,
    hotelId,
    totalGuests,
    bookingAmount,
  });
  const idempotencyKey = generateIdempotencyKey();
  const idempotencyKeyCreated = await createIdempotentKey(idempotencyKey, bookingCreated?.id);
  return {
    bookingId: bookingCreated?.id,
    idempotencyKey: idempotencyKeyCreated?.key,
  };
}

export async function confirmBookingService(idempotencyKey: string) {
  const idempotentKey = await getIdempotentKey(idempotencyKey);
  if (!idempotentKey) {
    throw new NotFoundError('Idempotency key not found');
  }
  if (idempotentKey.finalize) {
    throw new BadRequestError('Idempotency key already finalized');
  }
  const bookingFinalized = await confirmBooking(idempotentKey?.bookingId);
  await finalizeIdempotentKey(idempotencyKey);
  return bookingFinalized;
}
