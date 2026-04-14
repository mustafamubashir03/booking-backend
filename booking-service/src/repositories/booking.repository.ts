import { Prisma } from '../../prisma/generated/prisma/client';
import { prisma } from '../lib/prisma';

export async function createBooking(bookingInput: Prisma.BookingCreateInput) {
  const booking = prisma.booking.create({
    data: bookingInput,
  });
  return booking;
}
