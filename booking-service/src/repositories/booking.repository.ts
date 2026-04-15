import { Prisma } from '../../prisma/generated/prisma/client';
import { prisma } from '../lib/prisma';

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

export async function getIdempotentKey(key: string) {
  const idempotencyKey = await prisma.idempotencyKey.findUnique({
    where: {
      key,
    },
  });
  return idempotencyKey;
}

export async function getBookingById(bookingId: number) {
  const booking = await prisma.booking.findUnique({
    where: {
      id: bookingId,
    },
  });
  return booking;
}

export async function confirmBooking(bookingId: number) {
  const bookingConfirmed = await prisma.booking.update({
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

export async function finalizeIdempotentKey(key: string) {
  const idempotencyKeyFinalized = await prisma.idempotencyKey.update({
    where: {
      key,
    },
    data: {
      finalize: true,
    },
  });
  return idempotencyKeyFinalized;
}
