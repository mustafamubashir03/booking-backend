import {
  confirmBooking,
  createBooking,
  createIdempotentKey,
  finalizeIdempotentKey,
  getIdempotentKeyWithLock,
} from '../repositories/booking.repository';
import { Prisma } from '../../prisma/generated/prisma/client';
import { generateIdempotencyKey } from '../utils/helpers/generateIdempotencyKey';
import { BadRequestError, NotFoundError } from '../utils/errors/app.error';
import { prisma } from '../lib/prisma';

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
  return await prisma.$transaction(async (tx) => {
    const idempotentKey = await getIdempotentKeyWithLock(tx,idempotencyKey);
    console.log("idempotentKey at confirmBooking Service Transaction",idempotentKey)
    if (!idempotentKey?.key) {
      throw new NotFoundError('Idempotency key not found');
    }
    if (idempotentKey.finalize) {
      throw new BadRequestError('Idempotency key already finalized');
    }
    const bookingFinalized = await confirmBooking(tx,idempotentKey?.bookingId);
    await finalizeIdempotentKey(tx,idempotencyKey);
    return bookingFinalized;
  });
}
