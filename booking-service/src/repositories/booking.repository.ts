import { IdempotencyKey, Prisma } from '../../prisma/generated/prisma/client';
import { validate as isValidUuid } from 'uuid';
import { prisma } from '../lib/prisma';
import { BadRequestError, NotFoundError } from '../utils/errors/app.error';

export async function createBooking(bookingInput: Prisma.BookingCreateInput) {
  const bookingCreated = prisma.booking.create({
    data: bookingInput,
  });
  return bookingCreated;
}

export async function createIdempotentKey(key: string, bookingId: number) {
  const idempotencyKeyCreated = await prisma.idempotencyKey.create({
    data: {
      key,
      booking: {
        connect: {
          id: bookingId,
        },
      },
    },
  });
  return idempotencyKeyCreated;
}

export async function getIdempotentKeyWithLock(tx: Prisma.TransactionClient,key: string) {
  if (!isValidUuid(key)) {
    throw new BadRequestError('Key is not a valid uuid');
  }
  const idempotencyKey =
    await tx.$queryRaw<IdempotencyKey[]>`SELECT * FROM \`IdempotencyKey\` WHERE \`key\`=${key} FOR UPDATE`;
  console.log("idempotencyKeyAtLockFunction",idempotencyKey)
  if (!idempotencyKey) {
    throw new NotFoundError('Idempotency key not found');
  }
  return idempotencyKey[0];
}

export async function getBookingById(bookingId: number) {
  const booking = await prisma.booking.findUnique({
    where: {
      id: bookingId,
    },
  });
  return booking;
}

export async function confirmBooking(tx:Prisma.TransactionClient,bookingId: number) {
  if (!bookingId) {
    throw new BadRequestError('bookingId is required to confirm a booking');
  }
  const bookingConfirmed = await tx.booking.update({
    where: {
      id: bookingId,
    },
    data: {
      status: 'CONFIRMED',
    },
  });
  return bookingConfirmed;
}
export async function cancelBooking(bookingId: number) {
  const bookingConfirmed = await prisma.booking.update({
    where: {
      id: bookingId,
    },
    data: {
      status: 'CANCELLED',
    },
  });
  return bookingConfirmed;
}

export async function finalizeIdempotentKey(tx:Prisma.TransactionClient,key: string) {
  const idempotencyKeyFinalized = await tx.idempotencyKey.update({
    where: {
      key,
    },
    data: {
      finalize: true,
    },
  });
  return idempotencyKeyFinalized;
}
