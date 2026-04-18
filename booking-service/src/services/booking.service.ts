import {
  confirmBooking,
  createBooking,
  createIdempotentKey,
  finalizeIdempotentKey,
  getIdempotentKeyWithLock,
} from '../repositories/booking.repository';
import { Prisma } from '../../prisma/generated/prisma/client';
import { generateIdempotencyKey } from '../utils/helpers/generateIdempotencyKey';
import { BadRequestError, InternalServerError, NotFoundError } from '../utils/errors/app.error';
import { prisma } from '../lib/prisma';
import { redlock } from '../config/redisConfig';
import { serverConfig } from '../config';



export async function createBookingService({
  userId,
  hotelId,
  totalGuests,
  bookingAmount,
}: Prisma.BookingCreateInput) {
  let lock;
  const bookingResource = `hotel:${hotelId}`
  const ttl = serverConfig.LOCK_TTL
  try{
    lock = await redlock.acquire([bookingResource],ttl)
    console.log(lock)
    const bookingCreated = await createBooking({
            userId,
            hotelId,
            totalGuests,
            bookingAmount,
          });
    const idempotencyKey = generateIdempotencyKey();
  
    const idempotencyKeyCreated = await createIdempotentKey(
      idempotencyKey,
      bookingCreated?.id
    );

    return {
      bookingId: bookingCreated?.id,
      idempotencyKey: idempotencyKeyCreated?.key,
    };
  }catch(e){
    throw new InternalServerError("Resource has been locked temporary")
  }

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
