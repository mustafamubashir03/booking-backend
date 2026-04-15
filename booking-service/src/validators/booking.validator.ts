import { z } from 'zod';

export const createBookingSchema = z.object({
  userId: z.number({ message: 'userId must be present' }),
  hotelId: z.number({ message: 'hotelId must be present' }),
  totalGuests: z.number().min(1, { message: 'Guest must be atleast 1' }),
  bookingAmount: z.number().min(1, { message: 'booking amount must be atleast 1' }),
});
